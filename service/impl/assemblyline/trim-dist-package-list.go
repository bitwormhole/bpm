package assemblyline

import (
	"strconv"
	"strings"

	"github.com/bitwormhole/starter/collection"
)

////////////////////////////////////////////////////////////////////////////////

type distPackageListSegment struct {
	prefix    string // like format of 'type.id.'
	name      string // the package name
	id        string // the package id
	timestamp int64
	kvTable   map[string]string // the key-value table
}

func (inst *distPackageListSegment) init(segType, segID string) {
	inst.kvTable = make(map[string]string)
	inst.id = segID
	inst.prefix = segType + "." + segID + "."
	inst.name = "" // load lazy
}

func (inst *distPackageListSegment) put(key, value string) {
	inst.kvTable[key] = value
}

func (inst *distPackageListSegment) load() {

	date := inst.kvTable[inst.prefix+"date"]
	timestamp, _ := strconv.ParseInt(date, 10, 64)

	inst.name = inst.kvTable[inst.prefix+"name"]
	inst.timestamp = timestamp
}

func (inst *distPackageListSegment) writeTo(dst collection.Properties) {
	src := inst.kvTable
	for k, v := range src {
		dst.SetProperty(k, v)
	}
}

////////////////////////////////////////////////////////////////////////////////

type distPackageListTrimmer struct {
	segTableByName map[string]*distPackageListSegment
	segTableByID   map[string]*distPackageListSegment
}

func (inst *distPackageListTrimmer) reset() {
	inst.segTableByID = make(map[string]*distPackageListSegment)
	inst.segTableByName = make(map[string]*distPackageListSegment)
}

func (inst *distPackageListTrimmer) trim(src collection.Properties) collection.Properties {
	inst.reset()
	inst.loadItems(src)
	inst.trimDuplication()
	return inst.makeResult()
}

func (inst *distPackageListTrimmer) loadItems(src collection.Properties) {
	all := src.Export(nil)
	for k, v := range all {
		seg := inst.getSegmentByFullName(k, true)
		seg.put(k, v)
	}
	segments := inst.segTableByID
	for _, seg := range segments {
		seg.load()
	}
}

func (inst *distPackageListTrimmer) selectSegment(s1, s2 *distPackageListSegment) *distPackageListSegment {
	if s1 == nil {
		return s2
	} else if s2 == nil {
		return s1
	}
	if s1.timestamp < s2.timestamp {
		return s2
	}
	return s1
}

func (inst *distPackageListTrimmer) trimDuplication() {
	from := inst.segTableByID
	to := inst.segTableByName
	for _, seg := range from {
		pkgName := seg.name
		older := to[pkgName]
		to[pkgName] = inst.selectSegment(seg, older)
	}
}

func (inst *distPackageListTrimmer) makeResult() collection.Properties {
	src := inst.segTableByName
	dst := collection.CreateProperties()
	for _, seg := range src {
		seg.writeTo(dst)
	}
	return dst
}

func (inst *distPackageListTrimmer) getSegmentTypeAndID(fullname string) (segType, segID string) {
	array := strings.Split(fullname, ".")
	if len(array) > 2 {
		segType = array[0]
		segID = array[1]
	}
	return
}

func (inst *distPackageListTrimmer) getSegmentByFullName(fullname string, create bool) *distPackageListSegment {
	typ, id := inst.getSegmentTypeAndID(fullname)
	table := inst.segTableByID
	older := table[id]
	if older != nil {
		return older
	} else if create {
		seg := &distPackageListSegment{}
		seg.init(typ, id)
		table[id] = seg
		return seg
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

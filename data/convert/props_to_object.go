package convert

import (
	"crypto/sha256"
	"strconv"
	"strings"

	"github.com/bitwormhole/bpm/data/entity"
	"github.com/bitwormhole/bpm/data/po"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/util"
)

// LoadPackageManifest ...
func LoadPackageManifest(o *po.Manifest, props collection.Properties) error {
	const save = false
	list := make([]*entity.ManifestItem, 0)
	ids := listIds(props, "file.", ".id")
	for _, id := range ids {
		item := &entity.ManifestItem{}
		err := doManifestItemLS(save, id, item, props)
		if err != nil {
			return err
		}
		list = append(list, item)
	}
	o.Items = list
	return doManifestMetaLS(save, &o.Meta, props)
}

// LoadPackageSignature ...
func LoadPackageSignature(o *po.Signature, props collection.Properties) error {
	const save = false
	return doSignatureInfoLS(save, &o.Info, props)
}

// LoadMainConfig ...
func LoadMainConfig(o *po.AppMain, props collection.Properties) error {
	const save = false
	list := make([]*entity.MainScript, 0)
	ids := listIds(props, "script.", ".name")

	for _, id := range ids {
		item := &entity.MainScript{}
		err := doMainConfigScriptLS(save, id, item, props)
		if err != nil {
			return err
		}
		list = append(list, item)
	}

	o.Scripts = list
	return doMainConfigHeadLS(save, &o.Main, props)
}

// LoadPackageSourceList ...
func LoadPackageSourceList(file fs.Path) ([]*entity.PackSource, error) {
	text, err := file.GetIO().ReadText(nil)
	if err != nil {
		return nil, err
	}
	text = strings.ReplaceAll(text, "\r", "\n")
	list1 := strings.Split(text, "\n")
	list2 := make([]*entity.PackSource, 0)
	for i, row := range list1 {
		url := strings.TrimSpace(row)
		if len(url) < 6 {
			continue
		}
		ps := &entity.PackSource{}
		ps.URL = url
		ps.ID = strconv.Itoa(i)
		list2 = append(list2, ps)
	}
	return list2, nil
}

// LoadInstalledPackages ...
func LoadInstalledPackages(o *po.InstalledPackages, props collection.Properties) error {
	const save = false
	list := make([]*entity.InstalledPackageInfo, 0)
	ids := listIds(props, "installed.", ".id")
	for _, id := range ids {
		item := &entity.InstalledPackageInfo{}
		err := doInstalledPackagesItemLS(save, id, item, props)
		if err != nil {
			return err
		}
		list = append(list, item)
	}
	o.Packages = list
	return nil
}

// LoadAvailablePackages ...
func LoadAvailablePackages(o *po.AvailablePackages, props collection.Properties) error {
	const save = false
	list := make([]*entity.AvailablePackageInfo, 0)
	ids := listIds(props, "package.", ".id")
	for _, id := range ids {
		item := &entity.AvailablePackageInfo{}
		err := doAvailablePackagesItemLS(save, id, item, props)
		if err != nil {
			return err
		}
		list = append(list, item)
	}
	o.Packages = list
	return nil
}

// SaveInstalledPackages 保存已安装的包列表
func SaveInstalledPackages(o *po.InstalledPackages, props collection.Properties) error {
	const save = true
	all := o.Packages
	for _, item := range all {
		id := computeSHA256sumWithLength(10, item.Name)
		err := doInstalledPackagesItemLS(save, id, item, props)
		if err != nil {
			return err
		}
	}
	return nil
}

// SaveAvailablePackages 保存可安装的包列表
func SaveAvailablePackages(o *po.AvailablePackages, props collection.Properties) error {
	const save = true
	all := o.Packages
	for _, item := range all {
		id := computeSHA256sumWithLength(10, item.SHA256+item.Name)
		err := doAvailablePackagesItemLS(save, id, item, props)
		if err != nil {
			return err
		}
	}
	return nil
}

// SavePackageManifest 保存包项目清单
func SavePackageManifest(o *po.Manifest, props collection.Properties) error {
	const save = true
	all := o.Items
	for _, item := range all {
		id := computeSHA256sumWithLength(10, item.Path)
		err := doManifestItemLS(save, id, item, props)
		if err != nil {
			return err
		}
	}
	return doManifestMetaLS(save, &o.Meta, props)
}

// SavePackageSignature 保存包项目清单
func SavePackageSignature(o *po.Signature, props collection.Properties) error {
	const save = true
	return doSignatureInfoLS(save, &o.Info, props)
}

////////////////////////////////////////////////////////////////////////////////
// private

func doAvailablePackagesItemLS(save bool, id string, o *entity.AvailablePackageInfo, props collection.Properties) error {

	a := makeAdapterFor(save, props).Type("package").ID(id).Create()
	o.ID = id

	a.ForString(&o.Dependencies, "dependencies")
	a.ForString(&o.ID, "id")
	a.ForString(&o.Name, "name")
	a.ForString(&o.Platform, "platform")
	a.ForString(&o.SHA256, "sha256sum")
	a.ForString(&o.Type, "content-type")
	a.ForString(&o.URL, "url")
	a.ForString(&o.Version, "version")
	a.ForString(&o.DateString, "date-text")

	a.ForInt64(&o.Date, "date")
	a.ForInt64(&o.Size, "size")
	a.ForInt(&o.Revision, "revision")

	o.ID = id

	return nil
}

func doInstalledPackagesItemLS(save bool, id string, o *entity.InstalledPackageInfo, props collection.Properties) error {

	a := makeAdapterFor(save, props).Type("installed").ID(id).Create()
	o.ID = id

	a.ForString(&o.ID, "id")
	a.ForString(&o.File, "file")
	a.ForString(&o.Name, "name")
	a.ForString(&o.SHA256, "sha256sum")
	a.ForString(&o.Type, "content-type")
	a.ForString(&o.URL, "url")
	a.ForString(&o.Version, "version")
	a.ForString(&o.Dependencies, "dependencies")
	a.ForString(&o.Platform, "platform")
	a.ForString(&o.Main, "main")
	a.ForString(&o.DateString, "date-text")
	a.ForString(&o.Filename, "filename")

	a.ForInt(&o.Revision, "revision")
	a.ForBool(&o.AutoUpgrade, "auto-upgrade")
	a.ForInt64(&o.Size, "size")

	o.ID = id
	return nil
}

func doSignatureInfoLS(save bool, o *entity.SignatureInfo, props collection.Properties) error {

	const id = ""
	a := makeAdapterFor(save, props).Type("signature").ID(id).Create()
	o.ID = id

	a.ForString(&o.ID, "id")
	a.ForString(&o.Name, "name")
	a.ForString(&o.Type, "content-type")
	a.ForString(&o.URL, "url")
	a.ForString(&o.SHA256, "sha256sum")
	a.ForString(&o.Version, "version")
	a.ForString(&o.Dependencies, "dependencies")
	a.ForString(&o.Platform, "platform")
	a.ForString(&o.DateString, "date-text")
	a.ForString(&o.Filename, "filename")

	a.ForString(&o.Algorithm, "algorithm")
	a.ForString(&o.PublicFinger, "public-key-fingerprint")
	a.ForString(&o.Plain, "plain")
	a.ForString(&o.Secret, "secret")

	o.ID = id
	return nil
}

func doManifestMetaLS(save bool, o *entity.ManifestMeta, props collection.Properties) error {

	const id = ""
	a := makeAdapterFor(save, props).Type("package").ID(id).Create()
	o.ID = id

	a.ForString(&o.ID, "id")
	a.ForString(&o.Name, "name")
	a.ForString(&o.Type, "content-type")
	a.ForString(&o.URL, "url")
	a.ForString(&o.SHA256, "sha256sum")
	a.ForString(&o.Version, "version")
	a.ForString(&o.Dependencies, "dependencies")
	a.ForString(&o.Platform, "platform")
	a.ForString(&o.Main, "main")
	a.ForString(&o.DateString, "date-text")
	a.ForString(&o.Filename, "filename")

	a.ForInt(&o.Revision, "revision")
	a.ForInt64(&o.Size, "size")
	a.ForInt64(&o.Date, "date")

	a.ForString(&o.SignatureAlgorithm, "signature-algorithm")
	a.ForString(&o.SignaturePublicFinger, "signature-public-key-fingerprint")
	a.ForString(&o.SignaturePublicKey, "signature-public-key")
	a.ForString(&o.SignaturePrivateKey, "signature-private-key")

	o.ID = id
	return nil
}

func doManifestItemLS(save bool, id string, o *entity.ManifestItem, props collection.Properties) error {

	a := makeAdapterFor(save, props).Type("file").ID(id).Create()
	o.ID = id

	a.ForString(&o.ID, "id")
	a.ForString(&o.Name, "name")
	a.ForString(&o.Path, "path")
	a.ForString(&o.SHA256, "sha256sum")
	a.ForInt64(&o.Size, "size")
	a.ForBool(&o.IsDir, "isdir")
	a.ForBool(&o.IsOverride, "override")

	o.ID = id
	return nil
}

func doMainConfigHeadLS(save bool, o *entity.MainHead, props collection.Properties) error {

	a := makeAdapterFor(save, props).Type("main").ID("").Create()

	a.ForString(&o.Description, "description")
	a.ForString(&o.Name, "name")
	a.ForString(&o.Package, "package")
	a.ForString(&o.Script, "script")
	a.ForString(&o.Title, "title")

	return nil
}

func doMainConfigScriptLS(save bool, id string, o *entity.MainScript, props collection.Properties) error {

	a := makeAdapterFor(save, props).Type("script").ID(id).Create()
	o.Name = id

	a.ForString(&o.Arguments, "args")
	a.ForString(&o.Executable, "executable")
	a.ForString(&o.Name, "name")
	a.ForString(&o.WorkingDirectory, "working-directory")

	o.Name = id
	return nil
}

func makeAdapterFor(save bool, props collection.Properties) AdapterBuilder {
	ab := NewAdapterBuilder()
	if save {
		ab = ab.SetterFor(props)
	} else {
		ab = ab.GetterFor(props)
	}
	return ab
}

func listIds(props collection.Properties, prefix string, suffix string) []string {
	ids := make([]string, 0)
	all := props.Export(nil)
	for key := range all {
		if strings.HasPrefix(key, prefix) && strings.HasSuffix(key, suffix) {
			id := key[len(prefix) : len(key)-len(suffix)]
			ids = append(ids, id)
		}
	}
	return ids
}

func computeSHA256sumWithLength(length int, str string) string {
	data := []byte(str)
	sha256sum := sha256.Sum256(data)
	sum := util.StringifyBytes(sha256sum[:])
	if len(sum) > length {
		return sum[0:length]
	}
	return sum
}

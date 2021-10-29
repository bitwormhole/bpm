package tools

import (
	"archive/zip"
	"errors"
	"strings"

	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/util"
)

// Zip 生成zip文件，@to=zip-file,  @from=src-dir
func Zip(from, to fs.Path, override bool) error {

	if !from.IsDir() {
		return errors.New("the src path is not a dir, path=" + from.Path())
	}

	if to.Exists() {
		if override {
			to.Delete()
		} else {
			return errors.New("the dest zip file is exists: " + to.Path())
		}
	}

	opt := fs.Options{}
	opt.Create = true
	wtr1, err := to.GetIO().OpenWriter(&opt, true)
	if err != nil {
		return err
	}
	defer wtr1.Close()

	wtr2 := zip.NewWriter(wtr1)
	defer wtr2.Close()

	// wtr3, err := wtr2.Create("")
	// util.PumpStream(nil, wtr3, nil)

	walker := zipDirWalker{}
	walker.zipwtr = wtr2
	return walker.walk(from)
}

type zipDirWalker struct {
	zipwtr *zip.Writer
}

func (inst *zipDirWalker) walk(dir fs.Path) error {
	spath := dir.Name()
	return inst.walkInto(dir, spath, 99)
}

// spath:(short-path)
func (inst *zipDirWalker) walkInto(dir fs.Path, spath string, ttl int) error {
	if ttl < 0 {
		return errors.New("the path is too deep, path=" + dir.Path())
	}
	if !dir.IsDir() {
		return errors.New("the path is not a dir, path=" + dir.Path())
	}
	if spath != "" {
		if !strings.HasSuffix(spath, "/") {
			spath = spath + "/"
		}
	}
	items := dir.ListItems()
	for _, item := range items {
		name := item.Name()
		if item.IsDir() {
			err := inst.walkInto(item, spath+name+"/", ttl-1)
			if err != nil {
				return err
			}
		} else if item.IsFile() {
			err := inst.onFile(item, spath+name)
			if err != nil {
				return err
			}
		} else {
			// NOP
		}
	}
	return nil
}

func (inst *zipDirWalker) onFile(file fs.Path, spath string) error {
	in, err := file.GetIO().OpenReader(nil)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := inst.zipwtr.Create(spath)
	if err != nil {
		return err
	}
	return util.PumpStream(in, out, nil)
}

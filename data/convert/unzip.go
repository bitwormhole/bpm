package convert

import (
	"archive/zip"
	"errors"
	"strings"

	iofs "io/fs"

	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/util"
)

// Unzip 解压zip文件，@from=zip-file,  @to=dest-dir
func Unzip(from, to fs.Path) error {

	if !from.IsFile() {
		return errors.New("the path is not a file, path=" + from.Path())
	}

	if !to.IsDir() {
		return errors.New("the path is not a dir, path=" + to.Path())
	}

	reader, err := zip.OpenReader(from.Path())
	if err != nil {
		return err
	}

	defer reader.Close()
	files := reader.File

	for _, item := range files {
		info := item.FileInfo()
		target := to.GetChild("./" + item.Name)
		if info.IsDir() {
			target.Mkdirs()
			continue
		}
		err = unzipFileEntity(item, info, target)
		if err != nil {
			return err
		}
	}

	return nil
}

func unzipFileEntity(src *zip.File, info iofs.FileInfo, dst fs.Path) error {

	// vlog.Warn("unzip file ", src.Name)
	// vlog.Warn("  mod-time=", info.ModTime())
	// vlog.Warn("      name=", info.Name())
	// vlog.Warn("      size=", info.Size())
	// vlog.Warn("      mode=", info.Mode())
	// vlog.Warn("    is-dir=", info.IsDir())

	opt := fs.Options{}
	opt.Mode = info.Mode()
	opt.Create = true

	reader, err := src.Open()
	if err != nil {
		return err
	}
	defer reader.Close()

	writer, err := dst.GetIO().OpenWriter(&opt, true)
	if err != nil {
		return err
	}
	defer writer.Close()

	buffer := make([]byte, 1024*4)

	for {
		n, err := reader.Read(buffer)
		if n > 0 {
			_, err2 := writer.Write(buffer[0:n])
			if err2 != nil {
				return err2
			}
		}
		if err != nil {
			if strings.ToLower(err.Error()) == "eof" {
				break
			}
			return err
		}
	}

	writer.Close()

	// set mode
	dst.SetMode(info.Mode())

	// set time
	timeMod := util.TimeToInt64(info.ModTime())
	dst.SetLastModTime(timeMod)

	return nil
}

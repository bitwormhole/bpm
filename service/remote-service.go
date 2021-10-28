package service

import (
	"bytes"
	"errors"
	"io"

	"io/ioutil"
	"net/http"

	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
)

// HTTPGetService 提供HTTP下载服务
type HTTPGetService interface {
	Load(url string, dst io.Writer) error
	LoadText(url string) (string, error)
	LoadFile(url string, dst fs.Path) error
}

// HTTPGetServiceImpl 实现 HTTPGetService
type HTTPGetServiceImpl struct {
	markup.Component `id:"bpm-remote-service"`
}

func (inst *HTTPGetServiceImpl) _Impl() HTTPGetService {
	return inst
}

// Load 下载数据流
func (inst *HTTPGetServiceImpl) Load(url string, dst io.Writer) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	code := resp.StatusCode
	if code != http.StatusOK {
		msg := resp.Status
		return errors.New(msg)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	_, err = dst.Write(data)
	return err
}

// LoadText 下载文本
func (inst *HTTPGetServiceImpl) LoadText(url string) (string, error) {
	buffer := bytes.Buffer{}
	err := inst.Load(url, &buffer)
	if err != nil {
		return "", err
	}
	data := buffer.Bytes()
	text := string(data)
	return text, nil
}

// LoadFile 下载文件
func (inst *HTTPGetServiceImpl) LoadFile(url string, dst fs.Path) error {

	name := dst.Name()
	tmp := dst.GetChild("./../tmp-" + name + ".tmp~")

	// clear tmp
	defer func() {
		if tmp.Exists() {
			tmp.Delete()
		}
	}()

	// open tmp
	opt := fs.Options{}
	opt.Create = true
	writer, err := tmp.GetIO().OpenWriter(&opt, true)
	if err != nil {
		return err
	}
	defer writer.Close()

	// load to tmp
	err = inst.Load(url, writer)
	if err != nil {
		return err
	}
	writer.Close()

	// delete old dst
	if dst.Exists() {
		dst.Delete()
	}

	// tmp -> dst
	err = tmp.MoveTo(dst)
	if err != nil {
		return err
	}

	return nil
}

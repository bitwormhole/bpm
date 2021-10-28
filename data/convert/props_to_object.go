package convert

import (
	"crypto/sha256"
	"strings"

	"github.com/bitwormhole/bpm/data/entity"
	"github.com/bitwormhole/bpm/data/po"
	"github.com/bitwormhole/starter/collection"
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
	return nil
}

// LoadPackageSourceList ...
func LoadPackageSourceList(o *po.PackageSourceList, props collection.Properties) error {
	const save = false
	list := make([]*entity.PackSource, 0)
	ids := listIds(props, "source.", ".id")
	for _, id := range ids {
		item := &entity.PackSource{}
		err := doSourceListItemLS(save, id, item, props)
		if err != nil {
			return err
		}
		list = append(list, item)
	}
	o.Sources = list
	return nil
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

// SavePackageSourceList 保存软件源列表
func SavePackageSourceList(o *po.PackageSourceList, props collection.Properties) error {
	const save = true
	all := o.Sources
	for _, item := range all {
		id := computeSHA256sumWithLength(10, item.URL)
		err := doSourceListItemLS(save, id, item, props)
		if err != nil {
			return err
		}
	}
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
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// private

func doAvailablePackagesItemLS(save bool, id string, o *entity.AvailablePackageInfo, props collection.Properties) error {
	const (
		pkey         = "id"
		name         = "name"
		sha256sum    = "sha256sum"
		ctype        = "content-type"
		url          = "url"
		version      = "version"
		dependencies = "dependencies"
		revision     = "revision"
		size         = "size"
	)
	kp := "package." + id + "." // key-prefix
	if save {
		setter := props.Setter()
		setter.SetString(kp+pkey, id)

		setter.SetString(kp+name, o.Name)
		setter.SetString(kp+sha256sum, o.SHA256)
		setter.SetString(kp+ctype, o.Type)
		setter.SetString(kp+url, o.URL)
		setter.SetString(kp+version, o.Version)
		setter.SetString(kp+dependencies, o.Dependencies)

		setter.SetInt(kp+revision, o.Revision)
		setter.SetInt64(kp+size, o.Size)
	} else {
		getter := props.Getter()
		o.ID = id

		o.Name = getter.GetString(kp+name, "")
		o.SHA256 = getter.GetString(kp+sha256sum, "")
		o.Type = getter.GetString(kp+ctype, "")
		o.URL = getter.GetString(kp+url, "")
		o.Version = getter.GetString(kp+version, "")
		o.Dependencies = getter.GetString(kp+dependencies, "")

		o.Revision = getter.GetInt(kp+revision, 0)
		o.Size = getter.GetInt64(kp+size, 0)
	}
	return nil
}

func doInstalledPackagesItemLS(save bool, id string, o *entity.InstalledPackageInfo, props collection.Properties) error {
	const (
		pkey         = "id"
		file         = "file"
		name         = "name"
		sha256sum    = "sha256sum"
		ctype        = "content-type"
		url          = "url"
		version      = "version"
		revision     = "revision"
		dependencies = "dependencies"
		autoupgrade  = "auto-upgrade"
		size         = "size"
	)
	kp := "installed." + id + "." // key-prefix
	if save {
		setter := props.Setter()
		setter.SetString(kp+pkey, id)

		setter.SetString(kp+file, o.File)
		setter.SetString(kp+name, o.Name)
		setter.SetString(kp+sha256sum, o.SHA256)
		setter.SetString(kp+ctype, o.Type)
		setter.SetString(kp+url, o.URL)
		setter.SetString(kp+version, o.Version)
		setter.SetString(kp+dependencies, o.Dependencies)

		setter.SetInt(kp+revision, o.Revision)
		setter.SetBool(kp+autoupgrade, o.AutoUpgrade)
		setter.SetInt64(kp+size, o.Size)
	} else {
		getter := props.Getter()
		o.ID = id

		o.File = getter.GetString(kp+file, "")
		o.Name = getter.GetString(kp+name, "")
		o.SHA256 = getter.GetString(kp+sha256sum, "")
		o.Type = getter.GetString(kp+ctype, "")
		o.URL = getter.GetString(kp+url, "")
		o.Version = getter.GetString(kp+version, "")
		o.Dependencies = getter.GetString(kp+dependencies, "")

		o.Size = getter.GetInt64(kp+size, 0)
		o.Revision = getter.GetInt(kp+revision, 0)
		o.AutoUpgrade = getter.GetBool(kp+autoupgrade, false)
	}
	return nil
}

func doSourceListItemLS(save bool, id string, o *entity.PackSource, props collection.Properties) error {
	const (
		pkey = "id"
		url  = "url"
	)
	kp := "source." + id + "." // key-prefix
	if save {
		setter := props.Setter()
		setter.SetString(kp+pkey, id)
		setter.SetString(kp+url, o.URL)
	} else {
		getter := props.Getter()
		o.ID = id
		o.URL = getter.GetString(kp+url, "")
	}
	return nil
}

func doManifestItemLS(save bool, id string, o *entity.ManifestItem, props collection.Properties) error {
	const (
		pkey      = "id"
		name      = "name"
		path      = "path"
		sha256sum = "sha256sum"
		size      = "size"
		isdir     = "isdir"
	)
	kp := "file." + id + "." // key-prefix
	if save {
		setter := props.Setter()
		setter.SetString(kp+pkey, id)
		setter.SetString(kp+name, o.Name)
		setter.SetString(kp+path, o.Path)
		setter.SetString(kp+sha256sum, o.SHA256)
		setter.SetInt64(kp+size, o.Size)
		setter.SetBool(kp+isdir, o.IsDir)
	} else {
		getter := props.Getter()
		o.ID = id
		o.Name = getter.GetString(kp+name, "")
		o.Path = getter.GetString(kp+path, "")
		o.SHA256 = getter.GetString(kp+sha256sum, "")
		o.Size = getter.GetInt64(kp+size, 0)
		o.IsDir = getter.GetBool(kp+isdir, false)
	}
	return nil
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

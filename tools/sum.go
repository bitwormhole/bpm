package tools

import (
	"crypto/sha256"

	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/util"
)

// ComputeSHA256sum ...
func ComputeSHA256sum(file fs.Path) (string, error) {
	data, err := file.GetIO().ReadBinary(nil)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(data)
	str := util.StringifyBytes(sum[:])
	return str, nil
}

// ComputeSHA256sumForBytes ...
func ComputeSHA256sumForBytes(b []byte) string {
	sum := sha256.Sum256(b)
	return util.StringifyBytes(sum[:])
}

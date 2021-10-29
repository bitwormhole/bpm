package tools

import (
	"testing"

	"github.com/bitwormhole/starter/io/fs"
)

func TestZip(t *testing.T) {

	tmp := fs.Default().GetPath(t.TempDir())
	dir0 := tmp.GetChild(".dir")
	dir1 := dir0.GetChild("inner")
	file1 := dir0.GetChild("f1.txt")
	file2 := dir1.GetChild("f2.txt")
	dst := tmp.GetChild("test.zip")

	file1.GetIO().WriteText("h1", nil, true)
	file2.GetIO().WriteText("h2", nil, true)

	err := Zip(dir0, dst, false)
	if err != nil {
		t.Error(err)
	}

	t.Log("ok")
}

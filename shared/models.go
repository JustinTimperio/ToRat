package shared

import (
	"os"
	"reflect"

	"github.com/lu4p/shred"
)

// Make sure Models are never garbled.
var (
	_ = reflect.TypeOf(Void(0))
	_ = reflect.TypeOf(Cmd{})
	_ = reflect.TypeOf(Shred{})
	_ = reflect.TypeOf(File{})
	_ = reflect.TypeOf(Dir{})
	_ = reflect.TypeOf(EncAsym{})
	_ = reflect.TypeOf(Hardware{})
)

type Void int

type Cmd struct {
	Cmd        string
	Powershell bool
}

type Shred struct {
	Conf shred.Conf
	Path string
}

type File struct {
	Path    string
	Fpath   string
	Perm    os.FileMode
	Content []byte
}

type Dir struct {
	Path  string
	Files []string
}

type EncAsym struct {
	EncAesKey []byte
	EncData   []byte
}

type Hardware struct {
	OS     string
	CPU    string
	Cores  uint32
	RAM    string
	GPU    string
	Drives string
}

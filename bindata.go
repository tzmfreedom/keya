// Code generated by go-bindata.
// sources:
// data/members.toml
// data/songs.toml
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _dataMembersToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x98\xd1\x52\x1a\x49\x14\x86\xef\x79\x8a\x2e\x1e\x20\x05\xa3\xa0\x5c\xec\xc5\xee\x0b\xec\x03\xa4\x52\x5b\x26\xb1\xa2\x55\x51\xb7\x12\x6f\x72\x67\xf7\x22\x81\x41\x41\x23\x64\x34\x01\x47\x4d\x30\x8a\x06\x1d\x20\xab\x90\x1d\x78\x98\x43\xcf\x0c\x6f\xb1\xd5\x0d\x21\xce\x4c\xf7\xe8\x84\x9b\x5c\x24\xc4\xfa\x3c\xfc\xe7\x3f\xff\xdf\x8f\x1f\xaf\x2c\xae\x3c\x5d\x7c\xf5\xfa\xc9\x93\xc8\xea\xc2\xca\x22\xfa\x0d\x45\x6d\xbd\x65\x7d\x6a\x20\xe7\xa0\xe3\xa8\x46\x34\xb2\xb4\xfc\x6a\xe1\xc5\xc2\xea\x02\xfb\x27\xc0\x69\xc0\x1a\x10\x02\x44\x45\x80\x2f\x00\xef\x03\xce\x47\x23\x4f\x97\x5f\xad\x2f\x3d\x5f\x78\xc3\x3e\x13\x4f\xa5\xe6\x68\xa7\x9d\xb0\x2a\xd9\x39\x4b\xab\x45\x23\xcf\xd6\x56\x5f\xaf\x2f\xbe\x7c\xb9\xb0\xbe\xbc\xb6\x3a\xfa\x29\x2a\xe0\x0c\x60\x8d\x76\xbf\x44\x23\x4b\x8b\xcb\x2f\x96\xd6\xf9\xff\x4c\x2a\xcf\x56\xa2\x91\xa5\xb5\x11\x07\x6d\x34\xa8\x5e\x88\x46\x9e\xbe\x5c\x5b\x7b\xfe\xd7\xfa\x9b\xbf\xf9\xdf\xfe\x4e\x0f\xf3\xd1\x48\x44\x04\x3e\xf8\xae\x5a\xad\x1c\x1a\x98\xbb\xb4\x74\x25\x00\xef\xf1\x3f\x3f\x00\xee\x23\x20\x19\xc0\x69\x1f\xf8\x3c\xed\xb4\x53\x56\x25\x3b\x13\x93\x91\x9f\x02\xd9\x03\xdc\x02\xb2\xe7\x85\x4f\xcc\xde\x85\xb7\x6b\x87\xb4\x96\xa5\x37\x55\x2f\xff\x9f\x01\xfc\xb7\xaa\x55\xdd\x45\x4e\x21\xe7\x14\x2b\x5e\xfe\x0c\xe0\x2c\x90\x0d\x20\x39\x04\x44\x05\x5c\x17\x4e\x3d\x6e\x55\xb2\xb3\x12\x76\x92\x06\x5c\xf0\x51\x2b\x8f\x12\xae\xa1\x6f\xff\xe3\xec\xe6\x42\x40\xd3\xeb\xde\xf0\xfd\x09\xb2\x4e\xce\x86\xa7\x07\x5e\x68\x15\x70\x05\xf0\xf6\x88\x58\xac\x93\x78\x2c\x48\x28\x01\xe3\x4e\x4c\x3b\x6e\xbb\x53\xb2\x4b\xd7\x88\xd6\xb2\x43\xf2\x49\x40\x7e\x84\x00\xd7\x45\x83\x66\x2a\x49\x06\x0c\x1a\x7f\x03\xac\x03\x2e\xf9\xe4\x1d\x7f\x14\xbf\x0b\x3d\xac\x1e\xd0\xe6\x71\xa8\x59\x17\x2d\xe3\x08\xd9\xbd\x82\xd5\x3a\xf1\x12\xef\x71\x75\xef\x22\xc0\x7d\xce\xdd\x17\x72\xc7\x99\x42\xe2\x52\xf2\x32\xe0\x2a\x10\xd5\x37\xec\x94\x4b\x23\x9b\xdf\x68\xf7\xc2\xcb\xfd\x47\x30\xf7\xa1\x86\xec\x92\x31\xe8\x55\x05\xdc\xd7\x5c\x9b\x9a\x6c\x2b\x53\x3f\x64\xa2\xcc\x84\xd7\x49\x32\xe6\x42\xd7\xff\xb3\x0b\x3e\x79\x07\x78\xca\xd0\xcc\x3b\xda\x67\x44\x33\x97\x34\x5d\x17\xf8\x61\x99\x4f\xfd\x0c\x70\x06\xb1\xef\x9d\xfd\x02\x7e\x9d\xb3\xc1\x2b\x6c\xee\x09\x19\x7f\x9f\xfb\xd2\x16\x10\xec\x1b\xfd\xdc\x5d\x7e\xab\x6a\x0c\xba\x97\x61\x3c\xc5\xdc\x61\xfc\xce\xf9\xb9\xdd\x29\x09\xe0\xc7\xe4\x1a\x53\x3c\x51\x7d\xe4\xc9\x89\x64\x92\x21\x25\xe3\xf6\x72\x31\x77\xc0\xdc\x69\x5f\x63\xcb\x69\xa5\x3f\x0e\xcc\x96\x97\x5b\x1b\x2d\x67\x8f\xeb\x5c\x3c\x6e\x0e\x2d\xd7\x4b\x1a\xf0\xa9\x8f\xd8\xe5\x28\xd6\xfb\x6b\xcb\xd4\x43\x10\x3b\xc5\xcd\x41\xb7\x8c\x68\x31\x2f\x30\xc2\x03\xf6\xe5\xe2\xf4\x48\xe1\x42\x8d\x24\x26\xd0\xa9\x50\xd0\x33\x6e\xff\x96\x0c\x3a\x60\x39\x87\xd9\xb6\x55\xb9\xe4\xa6\x92\xce\x09\xb8\x3f\xb0\x73\x8f\xcf\xb8\xaf\x90\x8c\xd8\xc3\x99\xb8\xa5\xda\x16\x61\xbb\xdd\xdb\x4a\x7f\xb4\xf5\x5a\x28\xe8\xf2\x8d\xd5\x3c\x41\x4e\xb1\x02\x1b\x9b\x82\x79\xd7\xf9\xbc\xcb\x40\x8a\x63\x13\x07\xbc\x27\x44\x9f\x0d\x94\x89\x24\xa7\x24\xe6\x5d\x9e\x62\x18\xd4\x3c\x09\xa3\xed\x8a\x6e\x97\x74\x64\xef\x1e\xda\xe7\xc4\x4b\x7e\x05\xf8\xdf\x91\x87\xd7\x00\x7f\x17\x32\xcf\x05\x5d\xcc\x3c\xe0\x0b\x2f\xef\x5c\x7c\xba\x5d\x2c\xe8\x6c\x17\x9d\xdd\x6f\x7e\x58\x92\x1b\xed\x22\xe1\xc3\xf2\x78\xb7\x12\x8b\xc5\x7e\x31\x0a\x26\x92\xd3\xa9\x9a\x76\x5a\x56\x2e\xcf\xb6\xd1\x3a\x39\xa3\xb5\xac\x97\xbc\xc9\xc9\x4f\xf9\x42\x0a\xe2\x94\x12\x8b\xc5\xc7\x57\x5e\x91\x0a\x5b\x34\xea\x64\xec\x91\x72\xbf\xb4\x83\x6e\x7c\x23\x4b\x8d\x3c\x72\xb6\xbc\x09\x90\x87\x6e\x20\xe9\xd1\xb0\xf3\x80\xbf\x8a\x77\x91\xbb\xb5\x12\xd6\xad\x67\xa6\x4a\xde\xb6\xd1\x72\xae\x6f\x58\x94\x82\x8d\x4d\xab\x52\xf7\xa2\x67\x01\x7f\xe5\x99\xe8\x62\xb2\x8e\x7d\xa1\x54\x98\x91\x28\xd2\x74\x92\x01\xec\x8f\x26\xae\xfb\x48\x3f\x7f\x19\xee\xd7\x43\x44\x13\xeb\xf6\xd8\xe9\x75\x59\x7c\xa5\xea\x91\x17\xbb\xc8\x99\xeb\x80\x3b\xb2\x04\x9b\x18\xcb\x5b\x7e\x1e\x65\x55\xc7\x65\x80\xce\xd6\x99\x60\xe0\x01\x2a\xb1\x6e\x8f\x87\x24\x87\xec\x9d\xcc\xc0\xdc\xb9\x0f\xbb\x2c\xbc\x90\xcc\x45\x14\xe9\x5a\x6a\x22\xe6\x99\xa9\x98\x99\x67\xf7\x0d\xc4\xa4\x40\xf2\x62\xc3\xbe\x02\xdc\x9b\x7c\x40\xd2\xcd\xa4\xea\x50\x59\x9a\xb9\x2f\x40\x0d\xcb\x37\xb4\x1d\x86\x7a\xd0\x2d\xd3\xe2\x27\x64\x57\x0e\xed\x9b\x77\x82\x52\xb9\x03\xf8\x78\x94\x45\xfc\xe9\x69\x72\xd3\xe3\xd2\x3e\xf9\xa0\x7d\x94\x64\x91\x20\x59\x9b\x0d\x64\xb7\x35\x7b\x27\x23\xae\x92\x2c\xef\x21\x7e\x20\xe5\x55\x92\xef\xa2\x4c\xd7\xe2\x32\x39\x37\x6d\x27\xb3\xf4\x3e\x6d\x17\x90\xa3\x9a\xb4\xa1\x7b\xb9\xf3\x80\xb7\x01\xbf\x1f\x15\x4a\x2c\xb2\x90\xf8\x84\x3b\x54\x18\x49\x7a\x14\x32\x7c\x1b\xc6\xfc\xa8\x69\x50\xc3\x40\x34\x5d\x17\x44\xd5\x3c\xe0\x77\x7c\x58\x3d\x59\xf6\x1b\x63\x33\x17\x99\x0f\x7b\x24\x67\xa7\xea\x06\x54\x3d\xe2\xdd\xa6\xd8\xa4\xb7\x19\x01\xf8\xcf\x6e\xe0\xcf\x21\x3f\x2a\x8d\x74\xd0\x01\x8d\x26\x35\x55\x1a\x19\x9a\x2a\xa3\x1e\x74\x2f\xe9\x57\xaf\xf5\xb9\xeb\xd8\x36\x90\xb7\x7c\x70\xe2\xe8\x97\x0a\x0a\xad\x32\x33\x79\x48\xd8\x0e\x6c\x63\xfc\x40\x9e\xa1\x41\x67\xd3\xee\x15\x04\xf8\x65\x2e\x6f\x5c\x14\x15\xf8\xe4\xf8\x85\x47\xf9\x85\xa9\x27\x93\x53\x3d\xf3\xfc\x24\xb7\x7b\x05\xbb\xd0\x94\x93\xf7\x81\x6c\x49\x5b\x3c\x6f\xc1\xd2\xa3\x23\x2e\x38\xf7\x3f\x3d\x04\x1d\x9d\x8b\x7d\x7b\xe3\x92\xf5\x48\x41\x08\xd4\xb9\xcc\x3f\x8e\xab\xa4\xf4\xe2\x28\x52\xf7\x7e\xd0\xc9\x11\x67\x92\x20\x85\x5f\xec\xb3\x46\x46\xcd\x73\xc1\x83\x03\x67\xfe\xd1\xc8\x08\xb7\x96\xa0\x06\x1c\xf6\xee\xc4\x1e\xe4\xe0\x41\x97\xa7\x6a\xf0\x47\x4c\xd5\x14\xbc\xf5\x34\xf9\x95\xd7\x26\xef\x98\x58\xf6\x08\x3b\x7f\x5f\x2d\x13\x8e\xdd\x2d\x95\x5a\xd6\xc9\x5d\x49\xa4\xf2\x7f\x00\x00\x00\xff\xff\x53\x86\x08\x61\xf8\x16\x00\x00")

func dataMembersTomlBytes() ([]byte, error) {
	return bindataRead(
		_dataMembersToml,
		"data/members.toml",
	)
}

func dataMembersToml() (*asset, error) {
	bytes, err := dataMembersTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/members.toml", size: 5880, mode: os.FileMode(420), modTime: time.Unix(1488205092, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataSongsToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\xcb\x31\x0a\x82\x60\x14\xc0\xf1\xfd\x9d\xe2\xc3\x1b\x54\xb3\x27\x11\x87\x86\x68\x4a\x07\x4f\xf0\xde\x93\x08\x72\xad\x86\xa0\xa5\x20\xb0\xa0\x41\x02\xb1\xe8\x32\x7f\xa9\x6b\xb4\x15\xd4\x01\x7e\x49\x52\xe4\xd9\xb4\x48\x53\xc9\xf2\x10\x87\x81\x64\xe3\xd9\x24\xc4\x21\xc2\xae\xd8\x01\x3f\xe3\x0d\xbe\xc0\x77\x58\x8b\x1f\xf1\x1a\x9f\x63\x7b\xfc\x1e\x89\xfc\xf8\xe1\xc7\xf7\xed\xfa\xb5\xaa\xd0\x13\x7a\x79\x96\x5b\x74\x83\x2e\xd1\x1a\x2d\xff\xd9\xe8\xcb\xba\xaa\xef\x3a\xec\x86\x3d\xf0\x26\x12\x79\x07\x00\x00\xff\xff\x10\x4f\x04\x78\xa2\x00\x00\x00")

func dataSongsTomlBytes() ([]byte, error) {
	return bindataRead(
		_dataSongsToml,
		"data/songs.toml",
	)
}

func dataSongsToml() (*asset, error) {
	bytes, err := dataSongsTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/songs.toml", size: 162, mode: os.FileMode(420), modTime: time.Unix(1488377474, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"data/members.toml": dataMembersToml,
	"data/songs.toml": dataSongsToml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"data": &bintree{nil, map[string]*bintree{
		"members.toml": &bintree{dataMembersToml, map[string]*bintree{}},
		"songs.toml": &bintree{dataSongsToml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}


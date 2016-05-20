// Code generated by go-bindata.
// sources:
// assets/static/chat.tmpl
// DO NOT EDIT!

package assets

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

var _assetsStaticChatTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x58\x6d\x73\xd3\xba\x12\xfe\xde\x5f\x21\x04\x77\xc6\x19\x1a\xbb\x6f\x94\x92\x26\xbd\xc3\xd0\xde\xcb\x65\xe0\xd2\x99\xf6\xc0\x70\x80\x61\x14\x5b\xb1\x95\xca\x92\x91\xe4\xa4\x81\xe9\x7f\x3f\xbb\x7e\x8b\xe3\xa6\x4d\x7b\xe6\xf8\x43\x63\x4b\xbb\xfb\x3c\x5a\xad\x76\x57\x1d\x3e\x39\xfd\xf8\xe6\xf2\xcb\xf9\x19\x49\x5c\x2a\x4f\xb6\x86\xf8\x43\x24\x53\xf1\x88\x72\x45\x71\x80\xb3\x08\x7e\x9c\x70\x92\x9f\xbc\x49\x98\x23\x67\xd7\x2c\xcd\x24\x1f\x06\xe5\xd8\xd6\xd0\x86\x46\x64\x8e\x58\x13\x8e\x68\x10\xb0\x29\xbb\xf6\x63\xad\x63\xc9\x59\x26\xac\x1f\xea\xb4\x18\x0b\xa4\x18\xdb\x60\xfa\x33\xe7\x66\x11\xec\xf9\x3b\xfe\x7e\xf5\xe1\xa7\x42\xf9\x53\x4b\x4f\x86\x41\x69\x69\x69\xd2\x2d\x32\x3e\xa2\x8e\x5f\xbb\x60\xca\x66\xac\x1c\x05\x56\x04\x9e\x67\xde\x24\x57\xa1\x13\x5a\x79\x3d\xf2\xbb\x18\x9a\x31\x43\x42\xad\xd4\x71\xf3\x95\xda\x98\x8c\x40\x94\x3e\x85\x37\xda\x5b\x4e\xe4\x96\x1b\x29\xac\xab\x66\xeb\xcf\xb6\x88\xd4\xb5\x2e\xbc\xe1\x44\x31\x53\x83\x12\x96\x65\x5c\x45\xef\x75\xec\x81\xe9\x9a\x41\xad\x1b\x81\x26\x68\x7d\xdd\xf9\xbe\x3a\xac\x2f\x42\xa3\xa5\x84\xd9\xc8\xb7\xc5\xeb\xa5\xce\xc8\x68\xf9\xf9\x96\x8b\x38\x71\xa4\x0f\x03\xa1\x14\x5c\xb9\x72\xe0\xb8\x31\x03\x60\x7e\x09\x7d\xa9\x3d\x80\xe8\x35\x33\x62\x42\xbc\x1a\xa0\xcd\x07\x9f\x15\xb4\x87\x83\xdd\x6c\x95\x7f\x57\x57\x0e\x2e\xb6\xce\xe4\xa1\x83\xc5\x9f\x8a\x99\x17\x31\xc7\xda\x80\x05\x11\x18\xf3\x85\xfd\x91\xe6\x8e\x47\x5d\x36\x86\xbb\xdc\xa8\x36\x4c\xf3\x8a\x7b\x8d\x0c\x51\x1d\xdf\x8f\x97\x53\x91\x98\xbd\x91\xcc\x5a\x98\xa6\x74\xa9\x6c\xe7\xc2\x85\x49\x09\xc8\x67\xb0\x88\x1f\x18\x35\x5d\xc8\x90\x59\x4e\x76\x07\x2b\x63\x5d\xa3\x7c\xd6\x9f\x6a\xa1\x5a\xb6\x3b\xac\x68\xc2\x2c\x41\x11\x1e\xad\x11\x1a\x1b\xce\xae\x8e\x6f\xa3\xee\x6d\x44\xcd\x98\x71\x1b\x50\x25\x9f\x38\xe2\x51\xf2\x7c\xe9\x1a\x78\xa7\xbd\x47\xf0\xd8\xdf\xcc\xc3\x70\xcb\x55\xc8\xef\xe1\x22\x2c\x51\x7a\x4e\x56\x88\x3c\x9c\xc2\xc1\x46\x0a\x78\x4c\xef\x44\xff\x1b\x88\x2f\x36\x22\x72\x63\xee\x59\xef\x99\x31\xda\x0c\x1e\xb9\xde\x2a\x9e\xf1\xc4\xa3\xb3\x46\x44\xf1\x39\x39\x65\x8e\x7b\xad\xf4\xe2\x4c\x99\x5d\x86\xce\x04\x27\xb4\xe7\x33\xe7\x8c\x47\x43\x64\x46\xb7\x6b\x8e\x95\xbc\x33\xd5\x99\xf7\x0a\x85\x88\x14\x62\xa3\x6f\xd4\x26\x46\xa8\xab\x6f\xb4\xb0\x80\xdc\x3c\x00\xf4\x63\xee\xde\xea\xdc\x58\x48\x8c\x10\x23\x03\x24\x5f\x0d\x7f\x10\x0a\x0e\x24\x4c\xf4\x1e\x66\x99\x28\x11\x5e\x29\x96\xf2\x36\x44\xe1\x88\x89\xd1\xe9\x46\x23\xfc\x3a\x63\x2a\x6a\xeb\xe2\x9f\x5a\xad\x4c\x03\xa0\x7d\xbc\x3e\xcd\x24\x50\x86\xf8\x1f\x90\x99\xdf\x43\x66\xf6\x90\xc5\x36\x1e\x3e\xa1\xe2\x6e\xb6\x59\x33\x5c\xbb\x59\x8a\xca\xcd\xf0\x52\xd3\xc2\x4c\xdf\x8f\x8d\xce\xb3\xbe\x70\x3c\x5d\x71\x1f\xa0\xf4\x56\x37\xb8\xae\x0d\xf5\x1a\xa5\x68\x09\xdc\x10\x2e\x21\xcc\xee\xc0\x6d\x54\xc3\x44\xc8\xc8\x70\x28\x55\xfe\x44\x48\xc7\xcd\x4a\xf1\xaa\x1c\xf1\xcc\x73\x89\xb0\x15\x91\x1e\x56\x05\x64\x73\xd3\xa1\x23\x85\x6f\x78\xaa\x67\x4d\x30\x95\xbe\x6b\x79\x10\x4b\xd6\x44\x9b\x14\xd6\x64\xf3\x71\x2a\xdc\xed\x4a\x59\x3b\xee\x09\x96\xcb\xf5\xe9\x99\x4c\x18\xac\xac\x0b\xd1\x28\x62\x1d\x9a\x31\x09\x91\xf4\x48\x6d\x8c\x1e\x70\xcd\xaa\x12\x59\x66\xee\x01\x39\xd8\x5e\x9d\x43\x7f\x0c\x48\x03\xb8\xb4\xb9\x5a\x15\x71\x92\x36\x75\x1a\x1f\x5c\x9c\x6f\x71\xcf\xde\x5d\x7c\xfc\xbf\x0f\x45\x0b\x82\x44\x4c\x16\x65\xc5\x6a\xb9\xaf\xcd\xb9\xf4\x64\x6d\x26\x08\xc8\xff\x94\x70\x64\xce\xc7\x56\x87\x57\xdc\x59\x02\x11\x4d\xe0\xb0\xb2\x30\x21\x5a\xa5\xdc\x5a\x16\x73\x02\xb1\x1a\x49\x6e\xb6\x1a\xff\xcc\x85\x8a\xf4\xfc\x2b\xfd\xcc\xc7\x17\x85\x22\xfd\xde\xf6\xd4\xb2\x7b\xc0\xd8\x84\xe3\x7e\x32\x1c\x9f\x7c\x81\x43\x0b\xb9\x44\xcf\x21\x6c\xa0\x57\xe0\x98\x6b\xa1\xaf\xca\xb3\x4c\x1b\x47\x1a\x4b\xd6\x1f\x06\x63\xe8\x96\x50\x8b\xde\x5e\x46\x7d\x9a\xf0\x6f\x91\x82\xaa\x23\x0c\x4e\xcf\xe0\xd0\x66\xce\xa3\xe7\xd0\x9b\x41\xd4\x82\xcb\x01\x68\x81\xa8\x28\x40\xb7\x09\xc5\xe3\xb6\x5b\x77\x41\xe8\xbf\x2a\x7b\x35\xe0\x1e\x9d\xdb\x41\x10\xfc\xfe\xfd\xec\xe6\x26\x98\xdb\x7f\xa3\xe2\x88\x3e\xaf\x41\x5a\x9a\xbe\x56\xa1\xd4\x16\x71\x9b\xf8\xe3\x33\xb7\xd1\x0b\xa7\xc2\xa2\x3a\x0f\xa1\x73\x20\x98\x66\x08\x70\x9a\x71\xb3\xba\xea\xd6\x22\x2b\xb0\x7a\x2b\xee\x86\xab\x42\xaf\x08\x06\xa8\xb9\x96\xa3\x80\x5f\x04\x03\x6c\xf7\x3f\xd4\x51\x74\xd2\x56\x93\x2c\xb7\xa1\x65\xea\xa6\x16\x7c\x1e\xd1\x36\xdc\x69\xba\x08\xdc\xcd\xb6\x6f\xd6\x38\x7e\x6d\x33\xd7\x6b\xc7\x10\x9e\x85\x56\x6b\x2e\xb1\x2c\x18\x2e\x47\xd4\xba\x85\xe4\x36\xe1\x10\xdb\x24\x31\x7c\x32\xa2\x89\x73\x19\x46\x47\xca\xae\xc3\x48\xf9\x63\xad\x1d\x18\x67\x19\x7e\xe0\x3d\xa0\x19\x08\xf6\xfd\x7d\xff\x30\x08\xad\x5d\x8e\x15\xf7\x00\x18\xa1\x44\x40\x5c\xc6\x46\xb8\x05\x60\x24\x6c\xff\xe8\xa0\xbf\xfb\xf3\x28\xbd\x7c\xf7\xf1\xf5\xc5\xf5\xd1\x74\xf7\x75\xfe\x9c\xbd\xf8\x7c\xfa\x49\x9d\x8b\x3d\x79\xf5\x9f\xc9\x7c\x7e\xf6\x9a\x1d\x25\xa7\xa7\xd1\xf4\x4f\x99\xbd\xe7\xf1\x75\x32\xfd\xf4\xe1\x6c\x77\x12\x4f\x3f\x9f\xff\x37\xbd\xfa\x65\x5f\x52\x02\x0d\xaf\xb5\xda\x88\x58\xa8\x11\x65\x4a\xab\x45\xaa\x73\x8b\x37\x9c\x62\x19\xed\xab\x06\x72\x38\xd9\xda\x82\x40\x23\x7e\xd5\x13\x6e\xfb\x55\x9b\x56\x05\x41\xa8\x25\xb6\x05\x4f\x0f\x5f\xb2\x90\x1d\x1e\x6f\x41\xe6\x75\x6c\x8c\x66\x22\xbf\x39\x6f\xa5\xe8\x44\x2b\xd7\x9f\x17\x8d\xf5\x80\x8c\xb5\x8c\x56\xa5\xab\x3a\x5b\xca\xce\x13\xa8\x47\x7d\x9b\xb1\x90\x0f\xa0\x62\x83\x53\x40\xb6\x11\x2d\xab\x69\x2d\x2a\x22\x97\x0c\xc8\xab\x57\xff\x02\x11\xd8\x1e\x5c\x04\x2c\x26\xa8\xee\x6b\x63\x1d\x2d\xe0\x07\xd7\x50\x56\x3d\x0a\x1b\xed\x18\xb4\xae\xa6\x3f\x91\xb9\x88\xaa\x5b\x14\x3e\x6d\x29\x48\x3e\x94\x14\xc6\x60\x37\x2b\xd2\xaf\x76\x66\x49\x4b\xbc\xab\x02\xbe\xe8\xa7\x51\x7f\x77\xa7\x23\x53\xc8\x95\xe4\x45\x34\xa2\x78\x89\xaa\x55\x8a\xd1\x75\xe2\x41\x31\xd3\xc1\x2a\x8e\xfc\x26\xf8\xbd\x75\xe6\x72\x59\x40\x37\xb7\xbb\x5a\x67\xd9\x00\xac\x25\x91\xcb\x7b\x19\x74\x3f\x3b\xee\xdb\xec\xa9\xb5\x5c\xb1\x62\x17\x6c\x8b\xd2\x7d\x5b\xa0\x6b\x4c\xa8\x2c\xbf\x7b\x11\x8d\x46\x21\xd6\x8a\xee\xc6\x07\x08\xd3\xc7\xa0\x80\x9b\x20\x2d\x80\xb1\xf7\x26\x56\xfc\x02\xc9\xc3\x03\xe8\x89\xee\x36\xba\x9e\x46\x7f\xec\xd4\x3d\x54\x6e\xd1\x29\x9b\x93\x3b\x08\x41\x35\xcf\x41\xe6\x02\xd2\x14\x25\xf7\x71\xb9\x1d\x1d\x1b\xa6\x86\x01\x22\x3d\x74\x87\xab\x57\xa8\x3d\xe5\x89\x0a\xca\x7f\x94\xfc\x15\x00\x00\xff\xff\x30\x23\xfe\x57\x39\x11\x00\x00")

func assetsStaticChatTmplBytes() ([]byte, error) {
	return bindataRead(
		_assetsStaticChatTmpl,
		"assets/static/chat.tmpl",
	)
}

func assetsStaticChatTmpl() (*asset, error) {
	bytes, err := assetsStaticChatTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/static/chat.tmpl", size: 4409, mode: os.FileMode(420), modTime: time.Unix(1463722027, 0)}
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
	"assets/static/chat.tmpl": assetsStaticChatTmpl,
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
	"assets": &bintree{nil, map[string]*bintree{
		"static": &bintree{nil, map[string]*bintree{
			"chat.tmpl": &bintree{assetsStaticChatTmpl, map[string]*bintree{}},
		}},
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


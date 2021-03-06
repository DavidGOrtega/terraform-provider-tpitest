// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../cml/ami-test.json
// ../cml/ami.json
// ../cml/clone-amis.go
// ../cml/setup.sh
package iterative

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _CmlAmiTestJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x93\x51\x6b\xdb\x30\x10\xc7\xdf\xf3\x29\x84\x9e\x36\x88\xb7\x34\x0f\x5d\xe7\xb7\x10\xfa\x30\x68\xcb\xa0\x85\x31\xc2\x50\x15\xf9\x6a\x1f\x93\x25\xa3\x93\x9c\xb6\xc1\xdf\x7d\xc8\x4e\xd2\x59\xf1\x46\xdf\x42\xee\xe7\x9f\xee\x4e\x7f\xed\x67\x8c\x31\xc6\x5b\xe9\x50\x6e\x35\x10\x67\x39\x1b\xfe\xeb\xff\x47\x43\x5e\x1a\x05\xc2\xbf\x34\x10\x6b\x7c\xbf\x07\xd3\xb2\xc7\xd5\x8f\x7b\xf1\xed\xee\xfe\x61\x75\xb7\xbe\x16\x0f\x3f\xbf\x5f\x3f\x76\x1d\xef\xbf\xeb\xe6\x83\x72\x1b\x50\x17\xe0\x7a\xe3\xe6\x64\x7c\x73\xf7\xd0\x49\x2b\x6b\xf9\x6a\x4d\x06\x5b\xe2\x73\x36\x66\x24\x51\xa8\x41\x38\xab\x81\xe7\x89\xa0\x07\x62\x45\x48\x67\x78\xf4\x38\x93\xcb\x1d\xe5\x28\xeb\x3c\x5f\x5e\x2e\xbe\x5c\x2e\xae\xbe\x2e\xaf\x16\xcb\x3c\x52\x9f\x8b\x56\x65\xaa\xd6\x59\x23\xd5\x6f\x70\x7c\x7e\x2e\x23\x20\x42\x6b\x84\x91\x75\x3c\x8e\xbf\xd1\xd9\xa1\xc4\x47\x1f\x75\x63\x07\x77\x50\x46\x26\xce\x14\x28\xdb\x01\xf9\xec\x22\x39\x87\xcb\x1a\x07\x7f\xa4\xd0\x83\x93\x1e\x5b\xe8\x1b\xf3\x40\x7e\x0a\x2f\x80\x94\xc3\xc6\x1f\xdd\xeb\xdb\x1b\xf6\x61\x6d\x8d\x47\x63\x03\xb1\x5b\xa9\x2a\x34\xc0\x6e\x40\x3a\x83\xa6\xfc\x38\xe5\x28\x9d\x0d\x0d\xf1\x9c\x6d\xb8\xd4\x9a\xff\x4a\x90\x27\xeb\x14\x88\x02\xe2\x04\xe4\xc1\xc5\xe9\xbd\x0b\x90\xaa\x8e\x9c\x06\x0f\x82\x8c\x6c\xa8\xb2\xfe\x5f\x30\x51\x25\x02\x81\x3b\x8d\x1b\xb6\xc1\xf8\x90\x62\xe7\x39\x2b\x97\x9f\x96\xcf\x5a\xba\xf2\x5c\x69\x43\x6c\x20\x4e\xf4\x84\x7a\x68\x74\x22\x15\x43\x8d\x26\x8b\x31\xf1\xe8\x7c\x90\x1a\x5f\x65\xdc\x69\xd6\x9f\x9b\x33\x5e\xb5\xf5\x44\x2a\x18\xe3\xc7\x3c\x8c\xae\x6b\x1a\x75\xd6\xfa\xac\x80\x16\x15\x9c\xbc\x31\xd8\x67\x6c\x37\x91\x3f\xbb\x33\x43\xd7\x1b\xfe\x77\x7c\xd3\xdb\xea\xd9\xda\x92\x17\x0e\x14\x98\xb8\xff\xb8\xfe\xff\x47\x33\x18\xe1\x65\x99\xbc\xf1\x53\x79\x15\x7c\x65\xdd\x38\x93\x49\xd6\x67\xe3\x5f\x87\x9e\x78\xe3\x6c\x8b\xf1\x69\xbc\xeb\xb9\xe7\x8c\x53\x05\x3a\x5d\x1e\x47\xa3\xd1\xc4\xf2\x26\xe9\x8d\x9b\x16\x0b\x94\x19\xd5\x38\xee\x27\x4d\x30\x79\xe9\xe2\x42\xbc\x7b\x11\x1e\x6b\xb0\xa1\xcf\xe5\xc5\x22\xbd\x53\x0e\xcf\x0d\x28\x2f\x0a\x24\x65\x8d\x01\x75\xb6\xbf\xc3\x80\xb3\x6e\xf6\x27\x00\x00\xff\xff\x0e\xdc\x52\xd5\x22\x05\x00\x00")

func CmlAmiTestJsonBytes() ([]byte, error) {
	return bindataRead(
		_CmlAmiTestJson,
		"../cml/ami-test.json",
	)
}

func CmlAmiTestJson() (*asset, error) {
	bytes, err := CmlAmiTestJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../cml/ami-test.json", size: 1314, mode: os.FileMode(420), modTime: time.Unix(1615402156, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _CmlAmiJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x53\x4d\x8b\xdb\x30\x10\xbd\xe7\x57\x0c\x3a\xb5\x4b\x95\x78\x4d\xc8\x87\xa1\x87\xed\xee\x16\x0a\xbb\x29\x94\xf6\x54\x8a\x51\xec\x59\x5b\xd4\x96\xcc\x8c\xe4\xb4\xbb\xe4\xbf\x17\x39\xb1\x63\x13\xe8\x45\xe0\x99\xa7\xf7\x9e\x66\x9e\xdf\x66\x00\x62\xef\x75\x95\x23\xb1\x48\xe0\xe7\x0c\x00\xe0\xad\x3b\x01\x84\xfb\xdb\xa0\x48\x40\xa8\x5a\xbd\x5a\x23\x71\xcf\xe2\x43\xdf\x53\xcc\xbe\xc6\x94\x6c\x15\x20\xfd\x15\x00\x11\x2a\xa9\x22\xd3\x5d\x24\x93\xa8\x03\x27\x5a\xd5\x49\x12\xaf\xa2\xf5\x2a\xda\x6c\xe3\x4d\x14\x27\x01\xb5\xc8\xdb\x4c\x66\x75\x25\x1b\x95\xfd\x46\x1a\xb8\x01\x04\x23\xb3\xb6\x26\x35\xaa\xee\x1c\x5c\x50\xf2\xdc\x12\x67\xf0\x71\x70\x44\x58\x84\x7a\x02\xc2\xb3\x3c\x20\x3b\x79\x3b\xb2\x5b\xeb\x81\x4c\x3b\x24\xe5\x74\x8b\x41\x7c\x0a\xc9\x91\x33\xd2\x8d\x3b\x13\xdd\x3f\x3f\xc1\xbb\x7b\x6b\x9c\x36\xd6\x33\x3c\xab\xac\xd4\x06\xe1\x09\x15\x19\x6d\x8a\xf7\x73\xf8\xb1\xf7\xc6\x79\xb8\xdd\xcc\xa3\xe5\x94\xaa\x20\xeb\x9b\x6e\xa8\x42\x55\x95\xf8\x35\x34\x5f\x2c\x65\x98\xe6\x18\xfc\xb2\x43\x0a\x42\x8e\x3c\x8a\x2b\x44\x85\x0e\x53\x36\xaa\xe1\xd2\xba\x6b\x18\x73\x99\x7a\x46\xea\xdf\xe5\x3b\x2f\x97\xbe\x36\xec\x94\xc9\x30\xed\xf7\x58\xc4\xf3\xf8\x4f\xa5\xa8\x18\x93\x58\x1f\xc4\x82\xe3\x17\x5d\x9d\xec\x8c\xb6\x79\xaa\xf1\xa4\x08\x20\x5a\x4d\xce\xab\x4a\xbf\xaa\x30\x2a\xd9\x0b\x94\x6d\x3d\xda\x22\x80\x98\x5a\x5b\xe8\x5a\x15\xc8\x8b\x9b\xd3\xa7\xbc\x91\xdd\xdc\xa4\xaa\xf3\xd5\x52\x32\x52\x8b\x24\x6f\xa6\x0c\x64\xad\x93\x39\xb6\x3a\xc3\x41\x26\x04\x71\xc0\x1c\x47\xb1\xb1\x07\x73\x0e\xb2\x88\xb6\xdb\x75\x1c\xdd\x46\xdb\xe5\x7a\x7d\x19\x3e\x80\xa8\x2d\xbb\x94\x30\x43\x13\x46\x1a\x26\x7a\x9d\x24\x6f\x52\xa7\x8a\xe9\xab\xc5\x9d\x77\xa5\xa5\x49\x80\x86\x10\xce\xfa\xb3\x53\x12\x0d\xd9\x56\x87\x94\xfe\xf7\xb7\xe2\x12\xab\x51\xfe\xd0\xb4\x9a\xac\xa9\xd1\xb8\xb4\x55\xe7\x67\x3c\x3c\x7e\xfa\x72\xb7\x4b\x3f\x7f\xfb\xba\xfb\xfe\xb8\x7b\xf8\x68\xac\xd1\x26\xc8\x67\x9d\xfe\x25\x55\xa7\xdc\x06\xda\xf9\x82\xd1\xf9\x66\xce\xa5\xb8\xd8\x9a\x1d\x67\xff\x02\x00\x00\xff\xff\x47\x34\xb5\x2f\xed\x03\x00\x00")

func CmlAmiJsonBytes() ([]byte, error) {
	return bindataRead(
		_CmlAmiJson,
		"../cml/ami.json",
	)
}

func CmlAmiJson() (*asset, error) {
	bytes, err := CmlAmiJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../cml/ami.json", size: 1005, mode: os.FileMode(420), modTime: time.Unix(1615402156, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _CmlCloneAmisGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x55\x7f\x6b\xe3\xb6\x1b\xff\x5b\x7a\x15\xfa\x8a\xeb\x61\x97\xc4\x69\x73\x5f\x8e\x91\xa3\xb0\xd2\xe5\x46\x61\xbb\x95\x94\x6d\x7f\x74\xa5\xa8\xb2\xec\x88\x93\x25\x23\xc9\x49\x43\xc9\x7b\x1f\x92\x6c\x57\x76\x93\xae\x5c\xa1\x8a\xf5\xe8\xf9\xf9\x79\x1e\x7d\x34\x9b\x35\x46\xcf\x1e\xb9\x9c\x31\xb9\x41\xa5\x42\xba\x91\xe8\xc3\x19\xc2\x1f\x7e\xc6\x5f\x10\x7b\xe2\x16\xd6\x84\x7e\x27\x25\x43\x15\xe1\x12\x42\x5e\xd5\x4a\x5b\x94\x40\x80\x8b\xca\x62\x08\xb0\x50\xa5\xfb\x51\xc6\xad\x66\x27\x29\x86\x10\xe0\x92\xdb\x75\xf3\x98\x51\x55\xcd\xc8\xd6\xb8\xff\xa9\xc9\xbf\x4f\x4b\xe5\x3e\xf1\x7f\x29\xcc\x0c\x33\x86\x2b\xf9\xb6\xa2\x61\x7a\xc3\x29\x9b\x31\x3a\x1f\x29\xae\x89\x59\x73\xaa\x74\x3d\xb3\x4c\x6b\x52\x28\x5d\x4d\x6b\xd1\x94\x5c\x3a\xdb\xd9\x66\x3e\xcb\x39\x29\x31\x4c\x21\x2c\x1a\x49\x7d\x6d\x49\x8a\x9e\x21\xd0\xac\xe4\x4a\xa2\xc5\x05\xc2\x8d\x99\x6e\x99\xb1\xd3\x73\x0c\x01\xa9\xf8\x37\x52\x31\x2f\xe7\x96\x69\x62\xf9\x86\x4d\x69\x25\x70\x67\x62\xdc\xd9\xdd\xbd\xb1\x9a\xcb\xf2\x19\x02\xe0\xec\x19\x31\x76\x3a\xc7\x93\x78\x7b\xde\x6f\xbd\xf7\xf6\x94\xd4\x53\xa3\x1a\xbb\xee\x8e\x49\x3d\x95\x4a\xdb\xb5\x37\xf9\x74\x40\x36\xb2\x8b\x5d\x0f\x64\xf3\x03\xb6\xad\x1e\x25\x53\xca\xa4\xd5\x44\x74\x12\xd6\x1c\x90\xb4\x20\x0c\xb6\xf3\xe1\xf6\x53\xbf\xf5\x41\x3a\x6d\x43\xa2\x92\xf7\x10\x02\xd7\xd6\x09\x72\xeb\x52\x6b\xa5\x1d\x64\x6d\xa7\xb3\x6f\x6c\x7b\x1b\x3e\x93\x8f\x64\x6b\xb2\x2b\x25\x0b\xee\x81\x5c\x79\x7c\x17\xc8\x49\x6f\x3d\xbc\x49\x80\x3c\xdd\x4f\x20\x48\x21\xe0\x45\xe4\xf2\x7f\x17\x48\x72\xe1\x5a\x09\x84\x2a\xb3\x1b\xcd\xa5\x2d\x12\x7c\xb7\x5c\xad\xfe\x58\xdd\xa3\x13\x83\xa3\x04\x52\x08\x80\x32\xd9\xf2\x89\xdb\xe4\x3c\x6d\x73\xdc\x50\x97\x17\xa3\x73\x97\x53\xe2\x74\x53\xe8\x27\xe0\x86\x68\x52\xf9\x3e\x7f\x74\xa7\xbf\x30\x43\x35\x7f\x64\xd7\x15\x29\x99\xb9\x96\x75\x63\x5d\xd4\xaf\x5c\x58\xa6\xcd\x02\xdd\xdd\x9f\x3a\xb5\xb0\x77\x27\xc0\x2f\xc0\x0d\xd2\x02\xa1\xb8\x1e\x2c\x49\xc5\x70\x3a\xf1\xe7\x7f\x11\xd1\xb0\x60\xdf\x4e\x53\xa4\xd9\x0e\xa2\x2f\x1d\x80\xb0\x1e\x77\x4b\x34\x5d\x73\xcb\xa8\x6d\xf4\xfb\xdc\xe3\xa7\x9f\x3e\x3f\x7c\xfe\x3f\x8e\xfd\xef\x7d\xf3\x00\xf7\x65\xae\x98\x99\xa0\xf0\xb9\xd4\xa1\x81\x1b\x3a\x82\x22\xe9\xb1\x0a\xcd\x79\x51\x8f\x9a\xe3\x2e\x60\xf6\x55\xab\x6a\xa9\x75\xd2\x6b\xa4\x21\x54\x81\x04\x93\x49\x1f\x31\x0b\x7e\x53\x74\x71\x81\xce\x8e\xb6\x96\x54\x1c\x9d\x18\x24\x95\x45\x85\x6a\x64\x8e\x27\xa8\x43\xeb\x40\x9b\x9d\xf6\xe2\x02\x8d\x63\xdc\x9d\xdd\x43\x08\x36\x44\xa3\x6d\x89\x1c\x9b\x65\x7f\x13\x6e\x7f\xd5\xaa\xa9\x21\x28\x94\x46\x0f\x13\x94\x33\x63\xb9\x24\xb6\x25\x0a\x4d\x64\xc9\x50\xc7\x02\x2e\xbb\x6d\x99\x5d\xe6\xb9\x0f\x05\x4a\x85\x1c\xc5\x24\x91\x51\x18\x68\x14\xd0\xf7\xac\x03\x40\x51\xd9\xbe\xa0\x2b\xa1\x24\x97\x65\xc8\x0d\x59\xd5\xfa\x46\x27\x26\xcb\xb2\x7f\x24\x1e\x64\x10\x9c\xb9\x48\x0e\x36\x16\x7a\x42\x85\x92\xa1\x19\xed\x55\x39\x60\x32\x41\xa7\xa4\xe2\xa1\xec\xeb\xbc\xdd\x39\xb4\xda\xcf\xd0\xd3\xda\x19\xa4\x5f\xbc\xe3\xa8\x7b\xc3\x84\xc3\xcd\xa3\x47\xd3\x7e\x23\xe7\x17\x3f\x42\x26\x4c\xeb\xcc\xfb\x4a\xd2\x70\x18\x37\x0d\xb8\xc9\x18\xc6\xf5\xb9\x23\xd3\x50\xca\x8c\x29\x1a\x21\x76\xa1\xf2\xfc\x9d\xd1\xf7\x71\x57\xc2\xe8\x6d\x4b\xdf\xf0\x24\x85\xfb\xf6\x6d\x88\xb0\x34\xaa\xd1\x94\x0d\xda\x77\xc0\x73\x7f\xc2\x03\xb4\x4c\x5a\x5e\x70\xa6\x87\x72\xff\x9c\x0c\x24\x11\xe0\xfd\x6c\x30\x8f\xec\x73\xc7\x9c\x0f\x3f\xcc\x98\xaf\xeb\x0f\xe4\x79\x94\xef\xa8\xaa\x77\x2b\x66\x1a\x61\x27\xdd\x54\xb9\x9b\x7e\xa5\xea\x5d\x00\xc3\x73\x60\xbf\xed\xe9\xef\xd6\x63\xd4\x4e\xd5\x20\x85\x11\x1c\x9e\x8e\x6e\x23\x44\x17\x03\xee\x8a\xb1\xf6\xaa\x1d\xbf\xb5\x7f\x63\xc7\xfe\x9e\x3b\xbd\x08\xc6\xc5\x01\xbd\x78\xac\x1d\xb1\x05\x8a\x1a\x8d\xb7\x66\xb6\xd1\xd2\x49\xfb\x37\xc1\x8f\xc5\x9f\xd2\x72\xe1\x6b\x5b\x3e\x71\x63\x4d\xf2\xd6\x43\xd0\x62\x70\x8c\x6a\x4f\x5f\x10\xee\x2e\x61\x60\xdd\x1f\x7a\x40\x8c\x25\xf6\x9d\x14\x4f\x36\x84\x0b\xf2\x28\xd8\x6b\x96\x77\x8d\x7f\x98\xa0\x4a\xe5\xbc\xd8\x45\xfc\xfe\xbb\x17\xf8\x2c\x2f\xad\xd5\xfc\xb1\xb1\xed\x00\x1c\x3a\x19\x23\x30\x98\x82\x43\x65\xbb\xf8\xbf\x91\x46\xd2\xf5\x0d\xd3\x15\xf7\x33\xbd\x08\x8f\xec\x58\xec\xe3\x71\xea\x07\xd9\x78\x38\x2e\xf3\xbc\x87\x6a\xac\x1d\xf0\x0a\x2b\xf0\x24\xbe\x18\xbe\x8f\x42\x74\x98\xbd\x86\xc2\xcd\xc5\x0b\x12\x47\xa7\xa3\xdb\x4a\x2e\xe0\x1e\xfe\x1b\x00\x00\xff\xff\x0e\xc2\x59\xbc\x51\x0b\x00\x00")

func CmlCloneAmisGoBytes() ([]byte, error) {
	return bindataRead(
		_CmlCloneAmisGo,
		"../cml/clone-amis.go",
	)
}

func CmlCloneAmisGo() (*asset, error) {
	bytes, err := CmlCloneAmisGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../cml/clone-amis.go", size: 2897, mode: os.FileMode(493), modTime: time.Unix(1627376521, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _CmlSetupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x5b\x4f\xdc\x3c\x10\x7d\xcf\xaf\x98\x2f\x1f\x82\xf6\xc1\x09\x20\x54\xb5\x41\x54\xa2\xe2\x22\x54\x04\x55\xcb\x4b\x05\x68\xe5\xd8\x93\xc4\xdd\xc4\xb6\x3c\xe3\xa5\x91\xfa\xe3\xab\xec\x66\x6f\x2a\x45\xe5\x21\x2f\x9e\x33\xb7\x73\x4e\xe6\xff\xbc\x34\x36\xa7\x26\xb9\xb8\xba\x3e\x3f\xc9\x67\x32\xe4\xad\xab\x73\xd5\xb5\x13\x62\xa9\xa6\x59\xeb\xea\xc4\x54\x70\x0f\xff\x81\xa8\x20\xdd\x19\x70\x29\x3c\x1e\x03\x37\x68\x13\x80\xb3\xf3\x4f\x57\xa7\x37\x93\x8b\xaf\xb7\x37\x77\xe7\x37\x67\x27\xd6\x59\x63\x19\x83\x54\x6c\x66\x98\x00\xa0\x6a\x1c\xa4\xa7\x5f\xee\x8a\xe2\x12\xb9\x28\x4e\x89\x62\x87\xe2\x3b\x12\x3c\xa4\x1c\x22\x3e\xa4\xc7\x29\xfc\x02\x8a\xda\x01\x23\x82\x90\x90\x23\xab\x5c\x7a\x1e\xbe\x4c\x39\x5b\x65\x3a\xff\xb0\x2f\xe7\x99\x3d\x52\x92\xc0\x02\x2e\x3d\x43\xc0\xce\xcd\x10\xa2\x95\xcc\x68\x35\x6a\x11\x7d\x1d\xa4\x46\x1a\x50\x3d\x31\x76\x8a\x5b\xd0\x86\x64\xd9\xe2\x90\x22\xb4\x34\x6d\xbf\x84\x65\x84\x61\x66\x14\xae\x8b\x6a\x2d\x06\x54\x40\xef\xc8\xb0\x0b\x3d\x44\x6b\x66\x18\x08\x41\xf4\x2f\xa0\xbc\x97\x45\x6d\x58\x28\x17\x30\xf7\x5e\x6e\xa2\x3d\x43\xf4\x5a\x32\xc2\xee\xee\xea\x49\xd4\xc8\x60\x2c\xb1\x6c\x5b\x10\x3d\x90\xab\xf8\x49\x06\x14\x3e\x38\x8f\x81\x0d\x92\x50\xae\xeb\x9c\x85\x32\x9a\x56\x0b\x24\x42\xcb\x46\xb6\x50\x1b\x5e\x0d\xac\x62\x68\x41\x54\xf4\xed\x1a\x1a\x66\x4f\x45\x9e\xd7\xc8\x99\x76\x6a\x8a\x21\x53\xae\x03\xe1\xa0\x46\x16\xe3\x0b\x35\xab\x21\xa8\xf9\x23\x90\x00\x8c\x75\x23\x61\xe8\x9c\x06\x21\x2f\x61\x81\x80\x58\x46\xcb\x71\xd9\x98\x90\x2b\xa9\x5a\x10\xa2\x73\xda\x54\xfd\x3c\xa3\x58\x60\x8a\xf0\x04\x73\x3b\x85\x68\xf3\x65\x7d\xa7\xa6\xc3\xd4\xcf\x0c\x3c\x08\x1d\xb0\x45\x49\x48\x59\x23\xa9\x31\xca\x05\x3f\x0c\x9f\xd7\xbe\x5e\xda\x63\xe0\x6c\x8a\xfd\x40\x3e\x88\x0d\x6e\xc5\xa0\xc6\x86\x12\xa9\xc6\x12\xee\x65\x50\xcd\x89\xec\xf4\xbb\xa3\xc7\x7f\xe8\x03\x3b\x6f\x5a\x2a\x27\x63\x10\x84\xa2\xb7\xd0\x49\x63\xd3\xd7\x69\xc8\x18\x82\xac\x5c\xe8\xd6\x8b\xd2\x7a\x4d\x8d\x65\x66\x9d\x46\x72\x31\x28\x9c\xaf\x47\xc8\xd1\x4f\x0e\x0e\xb3\x9f\xcb\x2d\x4b\x49\xcd\xeb\x9a\x0e\x25\x7f\x6c\xff\x16\x1b\xd1\x85\x20\x42\x87\xb9\x89\x47\x47\x2d\xb1\xdb\x41\x90\x91\xdd\x98\xba\xb1\x00\x88\xf5\x0a\x76\x66\xb4\x91\x59\x6d\xb8\x89\x65\x66\xdc\xf8\x30\x7a\x68\x50\x6b\x50\xe8\x2f\x82\xbd\xbe\xde\x62\xbe\x83\xf7\xd9\xfe\xd1\x76\x24\x6b\x0d\xf1\xe6\xdd\x58\x1d\x8d\x05\xb9\x34\x07\x64\xfa\x99\xac\x97\xc9\xdd\x22\x76\x33\xf7\x70\xe5\xfb\xd5\x55\x09\x48\x2c\x03\x8f\xbf\xc7\x4a\x80\xf9\xcd\xdb\xbb\xfd\xbc\x07\x1f\xc7\x8b\x99\x54\x26\xf9\x1d\x00\x00\xff\xff\x54\xf5\x0f\x32\x6c\x05\x00\x00")

func CmlSetupShBytes() ([]byte, error) {
	return bindataRead(
		_CmlSetupSh,
		"../cml/setup.sh",
	)
}

func CmlSetupSh() (*asset, error) {
	bytes, err := CmlSetupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../cml/setup.sh", size: 1388, mode: os.FileMode(493), modTime: time.Unix(1627387284, 0)}
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
	"../cml/ami-test.json": CmlAmiTestJson,
	"../cml/ami.json":      CmlAmiJson,
	"../cml/clone-amis.go": CmlCloneAmisGo,
	"../cml/setup.sh":      CmlSetupSh,
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
	"..": &bintree{nil, map[string]*bintree{
		"cml": &bintree{nil, map[string]*bintree{
			"ami-test.json": &bintree{CmlAmiTestJson, map[string]*bintree{}},
			"ami.json":      &bintree{CmlAmiJson, map[string]*bintree{}},
			"clone-amis.go": &bintree{CmlCloneAmisGo, map[string]*bintree{}},
			"setup.sh":      &bintree{CmlSetupSh, map[string]*bintree{}},
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

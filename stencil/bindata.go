// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// fonts/bulbhead.flf
// fonts/drpepper.flf
package stencil

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

var _fontsBulbheadFlf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x56\x4d\x8f\xdb\x36\x13\xbe\xf3\x57\x3c\x30\x16\x88\x88\xd7\x5a\x6e\x92\x7d\x51\xb8\x08\x02\xb6\x40\x7b\xe8\xa9\x3f\x80\xe8\x84\xb6\x28\x5b\x8d\x2c\x05\x96\x8c\x6c\x0a\xfd\xf8\x82\x33\xa4\x2d\xab\xd9\x3d\xf5\xd0\xc3\x02\x4b\xad\xc9\xf9\x7e\xe6\x83\xac\xdb\xfa\x9d\xbf\xc3\x23\x1e\xb1\xd9\xe0\x01\x6f\x7f\x50\x3f\x9f\xdb\xed\x21\xf8\x0a\xdb\x6f\xf8\x2d\xd4\xf8\xbd\x1f\x3e\xfb\xee\xaf\x70\x5a\xe3\xdd\xfb\x3f\xcf\xdd\xe6\x51\xfd\xda\xec\xdb\x30\xe2\x14\xda\xe0\x87\x80\x77\xf7\x0f\x28\x4b\xfc\x74\xde\x9f\x87\x11\xff\x5f\xe3\xed\x66\xf3\x5e\x81\x88\x08\x98\x2f\xe4\x33\xf0\x92\xcd\x9c\xa4\x8a\x48\x70\x05\xa0\xf3\x02\xf2\x99\xa6\x22\x1e\x11\x91\x06\x0c\x91\xcb\x24\x05\x0d\xc2\x07\xe8\x82\x48\xe4\xa2\xc6\x74\x06\x8a\x5c\x9a\xa5\x4c\x64\x70\xcc\xa0\xa1\x55\x11\x4d\x1a\xfe\x46\xc1\xeb\xd7\x44\x7a\x31\x3b\xca\x8b\xc8\x28\xf5\xcb\xd3\x97\xd6\x77\x7e\x6c\xfa\x0e\x7d\x8d\xba\x39\x0d\x23\xda\xa6\x0b\x3f\xaa\x08\x27\x4a\xac\x8e\x7e\xdf\xec\xd0\x9d\x8f\xdb\x70\x5a\xa1\xee\x4f\xa8\x9b\x36\xa0\xa9\x42\x37\x36\x75\xb3\x63\x61\xe5\x63\xd8\x25\x86\x43\x7f\x6e\x2b\xf8\xf6\xab\xff\x36\x60\x1b\xf0\xc9\xbf\x59\xb3\x50\xd7\x7f\x55\x77\xc2\x34\x1e\x02\x56\x07\x7f\xaa\xb6\xad\xef\x3e\xaf\x22\xde\x5f\x4e\x4d\x37\x0e\xf0\x03\x3c\xf8\x74\x8d\xed\x79\xc4\xce\x77\x6f\xc6\xa8\x66\x38\x9e\x87\x43\xa8\xd4\xa3\x68\x38\x84\x66\x7f\x18\xa3\xc7\x1e\xbb\x83\x3f\xf9\xdd\x18\x4e\x2f\x12\xd7\xe8\xfa\x11\x4d\xb7\x6b\xcf\x55\xd3\xed\x51\x85\x61\x17\xba\x2a\x9c\x06\xb5\xd9\xb0\xd8\xd1\x3f\x71\xe4\x68\x43\xb7\x1f\x0f\x28\xc2\x53\x66\xde\xf5\xc7\x63\xe8\x04\x98\x41\xe3\x7f\xf0\xa8\xcf\xd5\x3e\xa0\xf6\xbb\xb1\x3f\xa9\x07\x31\x5c\x85\xda\x9f\xdb\x51\x9c\x3d\xf6\x55\xe0\xc0\xc7\x43\x33\xa0\xee\xbb\x11\x45\xdb\x7c\x0e\x58\x95\x47\x3c\xac\xd0\x77\xac\xd6\x77\x15\xab\xd5\xea\xed\x7b\x56\x22\x40\x47\xef\x6f\xac\x2a\x75\x77\x67\xe7\xcb\x2a\xe3\xac\xd2\x85\x55\xce\x58\x55\x68\x6b\xd5\x34\x59\x05\x5c\x96\x55\x98\x30\xc1\xaa\x72\x2a\xa7\xf2\xf2\x4f\x0e\xad\x02\x4d\x04\xab\x0c\x26\xd2\x56\x39\x9a\xe0\xac\x2a\x68\x22\xc3\x44\xae\xea\x78\xa0\x0d\x34\x2b\x34\x30\x51\x00\x05\xd7\x54\x64\x12\x16\x14\xc2\x60\x88\x3d\x21\x32\x2e\x7a\x67\x15\xd2\x9f\x70\x46\x0e\x68\xab\x0a\x14\xf1\xb7\x13\x15\x62\x04\x2e\x1e\x69\x26\x93\x11\x09\x58\xe5\x26\x63\x95\x99\x9c\x6c\xf9\x4b\x4c\x00\x4d\x60\xef\x0b\x6e\x3f\x1d\x49\x13\x4d\x89\x49\x56\xa1\x63\x70\x22\x25\x32\x24\x12\x24\xfc\xb8\xe1\x4e\x12\xc9\x86\x84\xab\xd9\x65\x03\x71\x2a\xc5\x41\x12\xb3\x89\xcd\xca\xae\x73\xff\x71\x40\x94\x5d\xa7\x14\x2d\x9f\xc7\x3e\x67\xd4\x52\xc0\x94\x70\x15\xf9\xa8\xc8\x24\xb7\xae\x0c\x42\xd7\x2c\x26\x69\x21\x49\x0b\xb2\xf9\x7b\xce\x6c\x0c\x5f\xa2\x89\x5e\x5c\xa5\x27\x70\x90\x13\x2d\xa5\x93\xef\x9a\x13\xcf\x44\x97\x89\xb7\x76\x17\x71\x27\x22\x98\xf8\x7d\xc9\xcb\x29\x60\xa4\x4a\x32\xbe\x85\x9e\xe1\x3b\xdb\xa8\x0b\x5c\x62\xd0\xaa\x0f\xf8\x20\xb5\xe1\xec\x3f\x12\x36\xdf\x59\x25\x38\x38\xc1\xf0\x23\x3e\x5a\x65\x92\xc1\x25\x7e\x46\xd0\xd7\x8b\xd4\xb8\x88\x52\xa1\x8d\x64\x2e\x57\x57\x2a\xaf\x34\x92\x23\xa7\x8c\x59\x56\xa8\x67\x39\x14\x38\x24\x64\x19\xcf\x29\x87\xe6\x52\x26\x92\x5c\x92\x92\x97\xbd\xa3\xb9\x86\x8b\x0a\x2e\xfe\x54\x47\xa2\x04\x37\x66\x12\x02\x3a\xce\xfd\x45\xa9\x3c\xc3\xa0\x67\xe5\x7a\xeb\x47\x79\x9f\x6b\x35\xb5\x1f\x44\x05\xdf\x48\x29\x96\x42\x1a\x1f\xc5\xad\x99\x5c\x6c\x11\x88\x1b\x3f\x32\x07\x32\xcb\x7d\xc9\x97\x97\xd4\x48\xae\xcc\x8b\x25\x93\x0c\x65\x3b\x79\x14\x08\xf8\x72\x51\x0a\x22\xb4\x08\x37\xe5\xaf\x00\x5c\xea\x2e\x1d\xfb\x58\xf4\x18\x67\xae\x43\x45\xd8\x5c\xf1\x9c\xa9\x59\x06\x85\x25\x5e\x91\xd9\xda\x33\x59\x66\xd0\x66\xe8\xbe\xac\xc5\x38\xf7\xfd\x5a\xe1\xfe\xb8\xf1\x45\xba\x47\xc6\xf0\x6d\xc7\x2e\xa1\x87\xc0\xca\xd3\x04\x4b\x4c\xf8\xbd\x90\x80\x9b\x05\xb3\xc4\x44\xe0\x77\xe0\x0e\x17\x24\x73\x29\x5c\x58\x9c\x59\xa2\x4b\xc6\x3d\xa3\x49\x5f\x13\xf0\xb2\xa9\x99\xd3\x97\xa8\xf2\xb0\x59\x66\x3a\x76\x27\xf2\x87\x28\x3b\x28\x77\x05\xb7\x92\xb4\x4d\xba\x47\xa4\xfc\x30\xa5\x0f\xd1\x64\x59\xaf\x74\xee\x1f\xd7\x89\x3f\x1b\xfc\xb3\xed\x72\xba\xb8\xdb\xcb\xeb\x75\x2e\xbc\xce\x85\xd7\xb9\xf0\x1f\x98\x0b\xeb\xd2\x2a\x8a\x0f\x1f\x7e\xfd\xe0\x53\x99\x9e\xc1\xf1\xe5\x1b\x97\x8b\xe0\x94\xf7\xcc\x40\x89\xab\x7c\x73\x79\x4f\x9a\xf8\x40\xcd\x6d\x6f\xaf\x3d\xae\xfe\xf5\xcd\xdf\x01\x00\x00\xff\xff\x10\x2f\x72\xbc\x8c\x0e\x00\x00")

func fontsBulbheadFlfBytes() ([]byte, error) {
	return bindataRead(
		_fontsBulbheadFlf,
		"fonts/bulbhead.flf",
	)
}

func fontsBulbheadFlf() (*asset, error) {
	bytes, err := fontsBulbheadFlfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fonts/bulbhead.flf", size: 3724, mode: os.FileMode(420), modTime: time.Unix(1667751253, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _fontsDrpepperFlf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x56\xdd\x6e\xe3\x36\x13\xbd\xe7\x53\x1c\xe4\x33\x3e\xed\xa2\x89\x98\x68\x93\xa0\x5d\x08\xaa\x8a\x76\x0b\xf4\xae\x40\x7b\x49\x60\xc2\x58\x94\x2d\x54\x96\x02\x89\xee\x6e\x00\x3e\x7c\x41\x52\x92\xf5\x67\xaf\x01\x5b\x7f\x3c\x1c\xce\x9c\x19\x1e\x4e\x5e\xe6\x91\xdc\xe0\x09\x8f\x88\xee\x71\x8f\x87\x67\xf6\x7b\x5d\x69\xe0\x33\x7e\x6b\x42\xfc\xa9\xde\xde\x54\x83\x0f\x32\xd7\xaa\x81\x44\x25\x0f\x0a\x45\x85\xba\x52\x68\x8b\x1d\x32\xfb\x50\x54\xd0\xfb\xa2\x45\xab\xdf\x4b\xf5\x31\x64\xbf\x1c\xf5\xbe\x6e\x3e\xe3\x8b\x6a\x6a\xfc\x2d\x0f\x87\xa2\x52\xd5\x2d\xf4\xc3\xd3\xfd\xa7\x87\xa7\x74\xbb\x0d\xf5\x51\x87\x79\x11\x32\xf6\xeb\x5e\x36\x72\xab\x55\xd3\x22\xf8\x5f\x00\x59\x65\x08\xfe\x1f\x40\x36\x0a\x65\x7d\x6c\xdf\xdd\x97\x3f\x82\x03\xaa\x5a\xe3\x5f\xd5\xbc\xa3\x95\xba\x68\xf3\x42\x65\xec\x6b\xa1\xf7\xd0\x7b\x85\x60\x13\xa0\x6e\x10\xe8\x20\x0c\x43\xfc\x75\xdc\xed\x54\xab\x8b\xba\x6a\x7f\x66\xec\xcb\xb7\xb7\x52\x56\xd2\xbe\xa2\xce\x91\x17\x4d\xab\x51\x16\x95\xfa\xcc\x6c\xe8\xb8\xc3\xcd\x41\xee\x8a\x2d\xaa\xe3\xe1\x55\x35\x37\xc8\xeb\x06\x79\x51\x2a\x14\x99\xaa\x74\x91\x17\x5b\x3b\x97\x49\x00\xb8\x43\xbb\xaf\x8f\x65\x06\x59\x7e\x95\xef\x2d\x5e\x15\x5e\x64\x70\xeb\xe6\x54\xf5\x57\xb6\xf1\x20\xeb\xd3\xcd\x5e\x36\xd9\x6b\x29\xab\x7f\x6e\x70\x77\x87\xb7\xa6\xa8\x74\x8b\x16\x12\xee\xe3\x2d\x5e\x8f\x1a\x5b\x59\x05\xda\x5a\x69\x0f\xc7\x76\xaf\x32\xf6\xe4\x0d\xec\x55\xb1\xdb\x6b\xeb\xaf\xc4\xb6\x67\x88\x3d\x5e\x18\xbc\x75\x04\x15\xd5\xb6\x3c\x66\x45\xb5\x43\xa6\xda\xad\xaa\x32\xd5\xb4\x2c\xba\xf7\xf3\x0e\xf2\x9b\x0b\x1c\xa5\xaa\x76\x7a\x8f\x0f\xea\x5b\x8f\xde\xd6\x87\x83\xaa\x3c\x2f\xed\x47\xfc\x00\xe4\xc7\x6c\xa7\x90\xcb\xad\xae\x1b\xd6\x19\xc8\x54\x2e\x8f\xa5\xf6\xce\x1e\xea\x4c\xb9\xb8\x5d\xe6\xf3\xba\xd2\xec\xe1\xd9\xc1\x3c\x91\xd6\xbf\x89\x59\xc6\x36\xe9\xf8\x97\x32\x10\x52\x66\x60\x52\x66\x88\xa7\x2c\xa6\x24\x65\x00\xfc\x08\xa5\xcc\x70\xc3\xdd\x17\x4c\xae\xfe\xe6\xde\x37\x64\xc8\xd0\x66\xf4\x00\x18\x98\x1e\x7a\x42\xa7\x0c\xc6\x90\x5d\xe3\x2e\x4e\x19\x27\xbb\x1e\x8c\xe9\x86\xc9\x8e\x24\x1c\xf6\x23\x07\x87\x45\xf0\x38\xe9\x8c\x78\x7f\xac\x8d\x18\x89\xbd\xf1\x50\xf0\x4d\xca\x04\x71\xb1\x19\x63\xac\xc7\xf6\x7d\xf8\xdb\x8b\xb5\x0d\x67\xda\x3a\xd6\x5f\x21\x48\xa4\x29\x23\xb2\x48\x01\xe1\x1c\xb4\x4c\xf8\x2b\x27\x7e\xf2\x9c\xb8\xa0\x94\x25\x40\x6c\xe7\xf1\x09\x0d\xde\x2f\x90\x81\x71\x64\xda\x77\x32\x8e\x06\x5a\xd0\xd0\xfd\x4f\x7e\xf6\xce\x33\x10\xf9\xe9\xe4\x27\x0f\x64\x77\x98\xfe\x62\x41\xe3\x34\x75\xe1\x61\x46\xdd\x64\x76\x67\x1a\x70\x89\x06\xb7\xb7\x17\x22\x0a\x26\xf4\xa6\x8c\xfb\x71\x57\x0d\xe6\x54\x08\xe4\xd7\x04\x92\x61\x85\x98\x88\x92\xe9\x0a\x2e\xb5\xe4\xbc\x88\x09\xc2\x63\xf8\x38\x06\x4f\x35\x78\xe8\xf8\xe7\x04\x84\x03\x4f\x3d\x4d\x63\x77\x1d\x13\x2f\xe4\x8c\x99\x99\xb1\x13\x26\xb1\xb7\xd0\x62\x5e\xd6\x31\xe4\xc2\xbe\x48\x4d\x8c\xd0\x06\xc7\xbd\x1d\x71\x66\xad\xd0\xf3\x06\xcf\x34\x71\xcc\xe8\xf3\x69\x99\x27\x68\x3a\x62\xba\x59\xb3\xc2\x8c\x11\xf7\x25\xc9\xe6\xee\xfb\x82\x58\x29\x8f\x59\xf1\x26\x2e\x04\x1a\x57\x67\x9f\xba\xd0\xa7\xce\x8f\xc5\x94\xac\x16\xc7\x8d\x4b\xbe\xe8\x48\x3f\xcf\x40\x5f\x47\x64\xc6\xae\x8c\x30\xa7\x8c\x9c\xcb\x1a\x7c\xd6\x62\x4b\xc0\xf9\xb5\x84\xdf\xaa\xe6\xbb\xd9\x77\x01\x99\x65\x49\xae\x60\xcc\x22\xfb\xae\xee\xbb\x61\xc4\xc4\xbb\x42\x1a\xd6\xeb\xc5\x10\x27\x6f\xce\x10\x30\x48\xe9\x7c\x07\xb9\x21\x3f\x97\xdc\x55\x0c\xe6\x3b\xe3\xe4\xf7\xa5\x93\x28\x97\x4d\x43\x43\x2d\x8c\xd4\xaf\x13\xd6\x41\x6a\x68\x96\x01\xf4\xfc\x0a\x78\xf6\x30\x72\x75\xc0\x4e\x43\x12\x93\x90\xc4\x99\x9c\x9a\x21\xfa\x85\x70\x4c\xf3\x05\x1b\xd9\x1a\xcf\x6b\x76\xc4\x05\x3b\x8e\x8c\x39\x0b\xe4\x65\xca\xe5\x54\xd0\xaa\xcc\x0c\xbb\x9e\x3a\x2d\x77\xdc\x9f\x24\x66\x25\x9d\xc1\xba\x20\x2e\x31\x86\xe6\xbb\x7e\x84\x1a\x62\xeb\x91\x27\xec\xb0\x59\xed\xf9\x28\xec\x09\xe2\x76\xbb\x4d\x93\x3f\x5c\x04\xce\x56\x9b\xf0\x54\x2c\x42\x70\x8a\x6b\x7a\x35\xea\xc4\x6d\x5e\x10\xfd\x76\x33\xeb\xb5\xe3\x9d\xea\x24\xa4\x77\xc9\x1a\x58\xa3\xbd\x17\xd2\xee\x94\xa4\x61\x67\x5a\x3b\xe0\x76\x62\xcc\x45\x72\xbe\x5d\xb8\x78\xd6\xa5\xcc\x7e\x12\x66\x71\x7a\x8f\x67\x58\xe9\x32\x3e\xe5\x66\x6d\x73\x8c\x2b\x68\x7e\xfa\x8c\xed\x70\x18\x04\xb6\x82\x0c\x85\x13\x0c\x61\x88\x8c\xfb\x6a\x15\xb3\xb5\x66\x76\xc2\xae\x12\x27\x76\x7a\xde\xdd\x22\x06\xe6\x6e\x45\x7a\xe6\x76\xfc\x5a\x61\x1f\x5e\xb0\x1a\xd7\x19\xd5\x89\xbd\xc4\xad\xaa\x8e\x3f\x7a\x46\x8d\x4d\xbc\x30\x3e\x12\x9f\x85\xea\x9c\x95\xb4\x2e\x80\xf0\xb4\x05\x82\x7e\x8f\x2c\xb5\x66\x84\x1d\xed\x25\x73\x91\xd8\xb5\x93\x78\x5a\x3e\x2b\x92\x73\x96\xd8\xbe\x74\x7d\xb1\xf6\x98\xde\x1f\xc7\xdf\x32\x41\xde\xcc\xac\x65\x9d\x35\x7e\x1b\xd7\xf8\x4d\xdb\xde\x95\xc6\x6f\xba\xe0\x20\x80\x6b\x04\x4c\x74\x67\x26\x38\x18\xa3\xa6\xca\xb3\x14\x1c\x0f\x1d\x7a\x04\x9e\x32\x3e\x69\x30\xd6\x9d\x9a\x96\x60\x97\xb8\x53\xa7\xc2\xfd\xf3\xb2\x85\xe1\x4e\x69\xc5\xb8\xb7\x36\x66\xfa\x5f\xf6\x2c\xc6\x66\x10\xc6\x1b\xe6\x63\xaf\xb8\xe0\x76\x16\x17\x1c\x8b\x76\x38\x4e\x28\x4e\xbe\xd3\x8f\xcc\x31\x6b\x67\xd7\x80\xb9\x74\x0e\xf8\xa5\x3b\xe8\xcb\xba\x00\x4d\x30\x17\x6b\x77\xba\xe4\xbc\x04\x56\x7b\x28\x84\xe3\xfa\x7e\xf8\xe9\xf9\xaa\xf8\xa3\x87\xc7\xab\x38\x88\xa2\xfb\xab\x78\x88\xa2\x4f\x57\xf9\x17\x45\x3f\x5e\xc5\x59\xf4\xf8\x7c\x15\x6f\xd1\x53\x74\x0d\x77\xff\x05\x00\x00\xff\xff\xe9\xac\x1c\x2e\x44\x11\x00\x00")

func fontsDrpepperFlfBytes() ([]byte, error) {
	return bindataRead(
		_fontsDrpepperFlf,
		"fonts/drpepper.flf",
	)
}

func fontsDrpepperFlf() (*asset, error) {
	bytes, err := fontsDrpepperFlfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fonts/drpepper.flf", size: 4420, mode: os.FileMode(420), modTime: time.Unix(1667013352, 0)}
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
	"fonts/bulbhead.flf": fontsBulbheadFlf,
	"fonts/drpepper.flf": fontsDrpepperFlf,
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
	"fonts": &bintree{nil, map[string]*bintree{
		"bulbhead.flf": &bintree{fontsBulbheadFlf, map[string]*bintree{}},
		"drpepper.flf": &bintree{fontsDrpepperFlf, map[string]*bintree{}},
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
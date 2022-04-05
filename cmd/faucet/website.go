// Code generated by go-bindata. DO NOT EDIT.
// sources:
// faucet.html (11.276kB)

package main

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _faucetHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x5a\x7b\x93\xdb\x36\x92\xff\x7b\xfc\x29\x3a\x3c\x7b\x25\x9d\x87\xa4\x66\xc6\xf6\xfa\x24\x52\x29\xaf\x37\xbb\xe7\xab\xbb\x24\x95\x38\x75\xb7\x95\x4d\x5d\x81\x64\x4b\x84\x07\x04\x18\x00\x94\x46\x99\xd2\x77\xbf\x6a\x80\xa4\xa8\xc7\x4c\xec\xb5\xaf\x6a\xfd\xc7\x98\xc4\xa3\xd1\x8f\x5f\xa3\x1f\x54\xf2\xd5\x9f\xbf\x7b\xfb\xfe\x6f\xdf\x7f\x03\xa5\xad\xc4\xe2\x49\x42\xff\x81\x60\x72\x95\x06\x28\x83\xc5\x93\x8b\xa4\x44\x56\x2c\x9e\x5c\x5c\x24\x15\x5a\x06\x79\xc9\xb4\x41\x9b\x06\x8d\x5d\x86\xaf\x83\xfd\x44\x69\x6d\x1d\xe2\xaf\x0d\x5f\xa7\xc1\xff\x84\x3f\xbd\x09\xdf\xaa\xaa\x66\x96\x67\x02\x03\xc8\x95\xb4\x28\x6d\x1a\xbc\xfb\x26\xc5\x62\x85\x83\x7d\x92\x55\x98\x06\x6b\x8e\x9b\x5a\x69\x3b\x58\xba\xe1\x85\x2d\xd3\x02\xd7\x3c\xc7\xd0\xbd\x5c\x02\x97\xdc\x72\x26\x42\x93\x33\x81\xe9\x55\xb0\x78\x42\x74\x2c\xb7\x02\x17\xf7\xf7\xd1\xb7\x68\x37\x4a\xdf\xee\x76\x33\x78\xd3\xd8\x12\xa5\xe5\x39\xb3\x58\xc0\x5f\x58\x93\xa3\x4d\x62\xbf\xd2\x6d\x12\x5c\xde\x42\xa9\x71\x99\x06\xc4\xba\x99\xc5\x71\x5e\xc8\x0f\x26\xca\x85\x6a\x8a\xa5\x60\x1a\xa3\x5c\x55\x31\xfb\xc0\xee\x62\xc1\x33\x13\xdb\x0d\xb7\x16\x75\x98\x29\x65\x8d\xd5\xac\x8e\x6f\xa2\x9b\xe8\x8f\x71\x6e\x4c\xdc\x8f\x45\x15\x97\x51\x6e\x4c\x00\x1a\x45\x1a\x18\xbb\x15\x68\x4a\x44\x1b\x40\xbc\xf8\xc7\xce\x5d\x2a\x69\x43\xb6\x41\xa3\x2a\x8c\x5f\x44\x7f\x8c\xa6\xee\xc8\xe1\xf0\xe3\xa7\xd2\xb1\x26\xd7\xbc\xb6\x60\x74\xfe\xd1\xe7\x7e\xf8\xb5\x41\xbd\x8d\x6f\xa2\xab\xe8\xaa\x7d\x71\xe7\x7c\x30\xc1\x22\x89\x3d\xc1\xc5\x67\xd1\x0e\xa5\xb2\xdb\xf8\x3a\x7a\x11\x5d\xc5\x35\xcb\x6f\xd9\x0a\x8b\xee\x24\x9a\x8a\xba\xc1\x2f\x76\xee\x43\x36\xfc\x70\x6c\xc2\x2f\x71\x58\xa5\x2a\x94\x36\xfa\x60\xe2\xeb\xe8\xea\x75\x34\xed\x06\x4e\xe9\xbb\x03\xc8\x68\x74\xd4\x45\xb4\x46\x4d\xc8\x15\x61\x8e\xd2\xa2\x86\x7b\x1a\xbd\xa8\xb8\x0c\x4b\xe4\xab\xd2\xce\xe0\x6a\x3a\x7d\x36\x3f\x37\xba\x2e\xfd\x70\xc1\x4d\x2d\xd8\x76\x06\x4b\x81\x77\x7e\x88\x09\xbe\x92\x21\xb7\x58\x99\x19\x78\xca\x6e\x62\xe7\xce\xac\xb5\x5a\x69\x34\xa6\x3d\xac\x56\x86\x5b\xae\xe4\x8c\x10\xc5\x2c\x5f\xe3\xb9\xb5\xa6\x66\xf2\x64\x03\xcb\x8c\x12\x8d\xc5\x23\x46\x32\xa1\xf2\x5b\x3f\xe6\xbc\x79\x28\x44\xae\x84\xd2\x33\xd8\x94\xbc\xdd\x06\xee\x20\xa8\x35\xb6\xe4\xa1\x66\x45\xc1\xe5\x6a\x06\xaf\xea\x56\x1e\xa8\x98\x5e\x71\x39\x83\xe9\x7e\x4b\x12\x77\x6a\x4c\x62\x7f\x71\x3d\xb9\x48\x32\x55\x6c\x9d\x0d\x0b\xbe\x86\x5c\x30\x63\xd2\xe0\x48\xc5\xee\x42\x3a\x58\x40\xf7\x10\xe3\xb2\x9b\x3a\x98\xd3\x6a\x13\x80\x3b\x28\x0d\x3c\x13\x61\xa6\xac\x55\xd5\x0c\xae\x88\xbd\x76\xcb\x11\x3d\x11\x8a\x55\x78\x75\xdd\x4d\x5e\x24\xe5\x55\x47\xc4\xe2\x9d\x0d\x9d\x7d\x7a\xcb\x04\x8b\x84\x77\x7b\x97\x0c\x96\x2c\xcc\x98\x2d\x03\x60\x9a\xb3\xb0\xe4\x45\x81\x32\x0d\xac\x6e\x90\x70\xc4\x17\x30\xbc\xfe\x1e\xb8\xfd\xca\xab\x8e\xaf\xb8\xe0\xeb\x56\xac\xc1\xe3\x91\x84\x0f\x0b\xf1\x1a\xda\x07\xb5\x5c\x1a\xb4\xe1\x40\xa6\xc1\x62\x2e\xeb\xc6\x86\x2b\xad\x9a\xba\x9f\xbf\x48\xdc\x28\xf0\x22\x0d\x1a\x2d\x82\xf6\xfa\x77\x8f\x76\x5b\xb7\xaa\x08\x7a\xc1\x95\xae\x42\xb2\x84\x56\x22\x80\x5a\xb0\x1c\x4b\x25\x0a\xd4\x69\xf0\xa3\xca\x39\x13\x20\xbd\xcc\xf0\xd3\x0f\xff\x09\xad\xc9\xb8\x5c\xc1\x56\x35\x1a\xbe\xb1\x25\x6a\x6c\x2a\x60\x45\x41\x70\x8d\xa2\x28\x88\xf7\x9c\x38\xf0\x9e\xf2\x1a\x66\x56\xee\xf9\xbd\x48\xb2\xc6\x5a\xd5\x2f\xcc\xac\x84\xcc\xca\xb0\xc0\x25\x6b\x84\x85\x42\xab\xba\x50\x1b\x19\x5a\xb5\x5a\x51\xa8\xf3\x52\xf8\x4d\x01\x14\xcc\xb2\x76\x2a\x0d\xba\xb5\x9d\x11\x99\xa9\x55\xdd\xd4\xad\x19\xfd\x20\xde\xd5\x4c\x16\x58\x90\xd1\x85\xc1\x60\xf1\x57\xbe\x46\xa8\xd0\x0b\x73\x71\x8c\x89\x9c\x69\xb4\xe1\x90\xe8\x09\x32\x92\xd8\x33\xe3\x45\x82\xf6\x5f\xd2\x88\x8e\x52\x2f\x42\x85\xb2\x81\x83\xb7\x50\xd3\xc5\x12\x2c\xee\xef\x35\x93\x2b\x84\xa7\xbc\xb8\xbb\x84\xa7\xac\x52\x8d\xb4\x30\x4b\x21\x7a\xe3\x1e\xcd\x6e\x77\x40\x1d\x20\x11\x7c\x91\xb0\xc7\xf0\x0d\x4a\xe6\x82\xe7\xb7\x69\x60\x39\xea\xf4\xfe\x9e\x88\xef\x76\x73\xb8\xbf\xe7\x4b\x78\x1a\xfd\x80\x39\xab\x6d\x5e\xb2\xdd\x6e\xa5\xbb\xe7\x08\xef\x30\x6f\x2c\x8e\x27\xf7\xf7\x28\x0c\xee\x76\xa6\xc9\x2a\x6e\xc7\xdd\x76\x1a\x97\xc5\x6e\x47\x3c\xb7\x7c\xee\x76\x10\x13\x51\x59\xe0\x1d\x3c\x8d\xbe\x47\xcd\x55\x61\xc0\xaf\x4f\x62\xb6\x48\x62\xc1\x17\xed\xbe\x43\x25\xc5\x8d\xd8\xe3\x25\x26\xc0\xf4\x40\x77\x7e\xe3\x58\x1d\x72\x7a\xc6\x0d\x56\x61\xcf\x7d\x8b\x07\xc3\x2d\xde\xe2\x36\x0d\xee\xef\x87\x7b\xdb\xd9\x9c\x09\x91\x31\xd2\x8b\x17\xad\xdf\xf4\x1b\x12\x4e\xd7\xdc\xb8\x9c\x6a\xd1\x71\xb0\x67\xfb\x23\xfd\xfa\xe8\xe6\xb2\xaa\x9e\xc1\xcd\xf5\xe0\xda\x3a\xe7\xf2\xaf\x8e\x5c\xfe\xe6\xec\xe2\x9a\x49\x14\xe0\xfe\x86\xa6\x62\xa2\x7b\x6e\xbd\x65\x70\x0d\x1c\x6f\x0a\xe9\x92\xee\x59\xeb\x2f\xfb\xe9\x1c\xd4\x1a\xf5\x52\xa8\xcd\x0c\x58\x63\xd5\x1c\x2a\x76\xd7\x07\xbc\x9b\xe9\x74\xc8\x37\xe5\x82\x2c\x13\xe8\xae\x17\x8d\xbf\x36\x68\xac\xe9\x2f\x13\x3f\xe5\xfe\xd2\x9d\x52\xa0\x34\x58\x1c\x69\x83\x4e\x24\xd5\xba\x55\x03\xd3\xf7\xca\x3c\xcb\xfb\x52\xa9\x3e\x86\x0c\xd9\x68\x49\x0f\xc2\x5d\xb0\x48\xac\xde\xaf\xbb\x48\x6c\xf1\x49\x31\x40\x53\x8e\xf7\x50\x08\xf0\x37\x1a\xc9\x5e\x23\x6a\x9f\x60\x10\x64\xc1\xbd\x26\xb1\x2d\x3e\xe3\x64\x02\x61\xc6\x0c\x7e\xcc\xf1\x2e\xd4\xef\x8f\x77\xaf\x9f\x7b\x7e\x89\x4c\xdb\x0c\x99\xfd\x18\x06\x96\x8d\x2c\x06\xf2\xbb\xbb\xf3\x73\x19\x68\x24\x5f\xa3\x36\xdc\x6e\x3f\x96\x03\x2c\xf6\x2c\xf8\xf7\x43\x16\x92\xd8\xea\xc7\xb1\x36\x7c\xf9\x42\xce\xfd\x7b\x39\xc9\xcd\xe2\xdf\xd5\x06\x0a\x85\x06\x6c\xc9\x0d\x50\x74\xfd\x3a\x89\xcb\x9b\x7e\x49\xbd\x78\x4f\x13\x4e\xa9\xb0\x74\xb9\x05\x70\x03\xba\x91\x2e\xf4\x2a\x09\xb6\xc4\xc3\x7c\xa4\x8d\xd2\x11\xbc\x57\x94\xd3\xad\x51\x5a\xa8\x98\xe0\x39\x57\x8d\x01\x96\x5b\xa5\x0d\x2c\xb5\xaa\x00\xef\x4a\xd6\x18\x4b\x84\xe8\xfa\x60\x6b\xc6\x85\xf3\x25\x67\x52\x50\x1a\x58\x9e\x37\x55\x43\x39\xa9\x5c\x01\x4a\xd5\xac\xca\x96\x17\xab\xc0\x07\x26\xa1\xe4\xaa\xe7\xc7\xd4\xac\x02\x66\x2d\xcb\x6f\xcd\x25\x74\xb7\x02\x30\x8d\x60\x39\x16\xb4\x2b\x57\x55\xa5\x24\xdc\xe8\x02\x6a\xa6\xed\x16\xcc\x61\x72\xc1\xf2\xdc\x45\xb9\x08\xde\xc8\xad\x92\x08\x25\x5b\x3b\x0e\xe1\xbd\xaf\x27\x88\xaf\xbf\xb0\x1c\x33\xa5\xfa\xd5\x50\xb1\x6d\x77\x5c\xcb\xfd\x86\xdb\x92\x7b\xf5\xd4\xa8\x2b\xda\x5a\x80\xe0\x15\xb7\x26\x4a\xe2\x7a\x7f\xa3\xee\x63\xb3\x08\x4b\xa5\xf9\x6f\x94\xd9\x88\xe1\xf5\x69\x8f\x2e\x97\xee\x6e\x74\x56\x17\xb8\xb4\x33\x78\xe1\xef\xc6\x63\x1c\xb7\x25\xd0\x39\x10\x77\x34\x5d\x69\x49\x01\x67\x06\x37\x3e\x9f\xf5\x89\x44\x61\x07\x1c\x14\x47\x50\xf3\x87\xbe\x7e\x5d\xdf\xf5\x7c\xf4\x49\xf1\xb4\x27\x42\x08\x38\x54\xca\x9a\xf7\x6a\xbc\x84\x8a\xdd\x22\x30\x48\xd8\x51\x89\xdc\x32\xed\x0a\x2c\xee\x1a\x04\xb1\xdd\x20\xda\xaf\xc9\x75\xd3\x1f\x3c\x41\x2e\x57\xcf\xae\xa7\x1e\x91\xf4\x40\xe4\x9f\x5d\x4f\xb9\xb4\xea\xd9\xf5\x74\x7a\x37\xfd\xc8\x7f\xcf\xae\xa7\x4a\x3e\xbb\x9e\xda\x12\x9f\x5d\x4f\x9f\x5d\xdf\x0c\xb1\xec\x47\xba\xd4\x92\x56\xa1\xa1\xd3\x3a\x88\x07\x60\x99\x5e\xa1\x4d\x83\xff\x65\x99\x6a\xec\x2c\x13\x4c\xde\x06\x0b\xc7\x2e\x65\x1b\x0e\x05\xe7\x13\x54\xa8\x99\x21\x48\x10\xc7\x0e\x25\x6d\x33\xc4\xc0\xd8\x34\x5a\xab\x46\x52\x54\x04\x92\xd9\x79\xa8\x1c\x11\xca\x48\x31\x93\x28\xc9\x74\xbc\x78\xab\xea\x6d\xe8\x88\xb8\xed\x27\x6a\x34\x4d\x5d\x2b\x6d\xa3\xa1\x3a\x19\x15\x42\x02\x4d\xfc\x7a\xfa\xf2\xf5\xab\x47\xd9\x37\x94\x66\x3b\x19\x7a\x0e\x59\xa6\xd6\x08\x3e\xa9\xcf\xd4\x1d\x30\x59\xc0\x92\x6b\x04\xb6\x61\xdb\xaf\x92\xb8\x70\x25\xd8\xe7\xa3\x76\xd9\x7a\xd7\x3f\x15\x6c\x3b\x97\xbf\x84\xba\xc9\x04\x37\x25\x30\x90\xb8\x81\xc4\x58\xad\xe4\x6a\xe1\x46\x73\xaa\x49\xdd\x2b\xd4\xca\xd8\xc7\xcc\x8f\x55\x86\x45\x71\x06\x00\x5f\xca\xfe\x9b\xcd\x26\xea\x34\xe9\x8c\x5f\xa2\xa8\x63\xba\xfe\x1a\xc9\xed\x36\xf6\x6e\xa4\x64\xfc\x35\x2f\xd2\xeb\xd7\xd7\xaf\x5e\x5d\xbf\xf8\xb7\xd7\x2f\x5f\x5e\xbf\x7e\xf1\xf2\x21\x64\x90\x50\x9f\x09\x0c\x9f\x46\x7f\xab\xa8\x6c\xed\x73\x68\x8f\x97\x2e\x77\xa3\x08\x5d\x50\x0d\xa2\x83\x7f\x18\x43\x8d\xa4\x44\x24\x64\xe2\x6c\x0e\xf1\x09\x28\x72\x30\x7a\x84\xb3\xcf\x84\x56\x07\x1f\x42\x8a\x6a\x2c\x49\xd8\x55\xf3\x5c\xc9\x1e\x4e\x97\x60\x78\x55\x8b\x2d\xe4\x7b\xab\x9f\xc7\xd5\x83\x46\xf9\x5d\x58\x1d\x9a\xcd\x83\xcc\x45\xff\x4a\x15\x48\x51\xdf\x34\x26\xc7\xda\xb5\x79\x29\x92\xfe\x69\xfb\x1b\x93\x96\x4b\xec\x22\x6e\x04\xdf\x49\xb1\x85\xc6\x20\x2c\x95\x86\x02\xb3\x66\xb5\x72\x69\x82\x86\x5a\xf3\x35\xb3\xd8\x85\x59\xd3\xa2\xa2\x07\xc5\xa0\xb2\xa1\x94\x47\x0c\x32\x90\xbf\xa9\x06\x72\x26\xc1\x6a\x96\xdf\x7a\x4f\x69\xb4\x26\x4f\xa9\xd1\x4b\xd3\x07\xfa\x0c\x85\xda\xb8\x25\x5e\xee\x25\x47\xe1\xa2\xbe\x41\x84\x52\x6d\xa0\x6a\x72\xe7\x90\x14\xd5\x9d\x10\x1b\xc6\x2d\x34\xd2\x72\xe1\xf5\x69\x1b\x2d\x29\x47\xc0\x83\x28\x7d\x52\xfb\x25\x58\x2d\xde\x97\x78\x26\x25\xea\xab\x36\xd0\xf8\xd6\x2f\x87\x5a\x2b\x8b\x39\x19\x14\xd8\x8a\x71\x69\xc8\x22\x2e\x0f\xc0\xea\x23\xaa\xba\xfe\xa9\x7d\xd8\xb7\x28\xdd\x74\x1c\xc3\x5f\x85\xca\x98\x80\x35\x21\x3d\x13\x94\xce\x29\x28\x15\x89\x3e\xd0\x96\xb1\xcc\x36\x06\xd4\xd2\x8d\x7a\xce\x69\xff\x9a\x69\xb2\x20\x56\xb5\x85\xb4\x6d\xb0\xd1\x98\x41\xbd\x6e\xdb\x86\xf4\x4a\x95\xfb\xc1\x7c\xaf\xf5\x14\x7e\xfe\x65\xfe\xa4\x65\xe5\xcf\xb8\x74\x90\x20\x7c\x7b\x91\x6d\xc9\x2c\xe4\x1a\x99\x45\x03\xb9\x50\xa6\xd1\x9e\xc3\x42\xab\x1a\x88\xcb\x8e\x52\x47\x99\x26\x6a\x77\x5a\x47\x64\x5c\x32\x53\x4e\xda\xfe\xa0\x46\x67\xa5\x7e\xae\x1b\xbf\x20\xd4\x8d\x89\x00\x4f\xa7\x73\xe0\x49\x47\x37\x12\x28\x57\xb6\x9c\x03\x7f\xfe\xbc\x5f\x7c\xc1\x97\x30\xee\x56\xfc\xcc\x7f\x89\xec\x5d\x44\xa7\x40\x9a\xc2\xf0\x34\x77\x60\x4b\xc7\xd4\x82\xe7\x38\xe6\x97\x70\x35\x99\x77\xb3\x99\x46\x76\xdb\xbd\xb5\x76\xf4\xff\xb9\xbf\xbb\xf9\xa1\x66\x9c\xf2\x0f\x74\xe3\x6b\x7f\x03\x0c\x56\xdc\x58\x68\xb4\x80\xd6\x87\xbd\x09\x7a\x83\xb8\x75\x43\xad\x9c\xe0\xb2\x7d\x68\x31\xd5\x89\xe0\xc9\x44\x06\x65\x31\xfe\x8f\x1f\xbf\xfb\x36\x32\x56\x73\xb9\xe2\xcb\xed\xf8\xbe\xd1\x62\x06\x4f\xc7\xc1\xbf\x34\x5a\x04\x93\x9f\xa7\xbf\x44\x6b\x26\x1a\xbc\x74\xf6\x9e\xb9\xbf\x27\xa7\x5c\x42\xfb\x38\x83\xc3\x03\x77\x93\xc9\xfc\x7c\x9f\x64\xd0\xd6\xd1\x68\xd0\x8e\x69\x61\x0f\xfc\x63\x1d\x31\xa8\xd0\x96\xca\xb9\xae\xc6\x5c\x49\x89\xb9\x85\xa6\x56\xb2\x55\x09\x08\x65\xcc\x1e\x88\xdd\x8a\xf4\x14\x14\xed\xfa\xd4\x05\xeb\xff\xc6\xec\x47\x95\xdf\xa2\x1d\x8f\xc7\x1b\x2e\x0b\xb5\x89\x84\xf2\x57\x6d\x44\x4e\xaa\x72\x25\x20\x4d\x53\x68\xa3\x68\x30\x81\xaf\x21\xd8\x18\x8a\xa7\x01\xcc\xe8\x91\x9e\x26\xf0\x1c\x8e\xb7\x97\x14\xef\x9f\x43\x10\xb3\x9a\x07\x13\xef\x0e\x9d\xe2\x95\xac\xd0\x18\xb6\xc2\x21\x83\xae\x32\xea\x41\x46\x72\x54\x66\x05\x29\x38\x03\xd5\x4c\x1b\xf4\x4b\x22\xaa\xc6\x3b\xb4\x11\x66\xdd\xb2\x34\x05\xd9\x08\xb1\x07\xa9\x77\x8a\x79\x07\xbf\x83\xe5\x91\x8f\x35\x5f\xa5\x29\x50\x69\x4a\x2a\x2e\xf6\x3b\xc9\xf8\xbe\x88\x9e\x44\x14\x17\xf6\x3b\x26\xf3\x21\x9a\x0f\xa8\x61\xf1\x7b\xe4\xb0\x38\xa6\x87\xc5\x03\x04\x5d\xcf\xe2\x31\x7a\xbe\xc7\x31\x20\xe7\x06\x1e\xa0\x26\x9b\x2a\x43\xfd\x18\x39\xdf\xb3\x68\xc9\x39\x55\xbf\x93\x76\xb0\xf7\x12\xae\x5e\x4d\x1e\xa0\x8e\x5a\xab\x07\x89\x4b\x65\xb7\xe3\x7b\xc1\xb6\x94\x33\xc1\xc8\xaa\xfa\xad\x6b\x31\x8c\x2e\x5d\xc4\x9d\x41\x4f\xe1\xd2\x35\x8f\x67\x30\x72\x6f\x34\xcf\x2b\x74\xbb\x5e\x4e\xa7\xd3\x4b\xe8\x3e\xbb\xfc\x89\x91\x13\xea\x06\x77\x0f\xf0\x63\x9a\x3c\xa7\xb8\xff\x39\x1c\xb5\x34\x7a\x9e\xda\xf7\xcf\xe0\xaa\x8f\x0d\x07\x6c\xc1\x1f\xfe\x00\x27\xb3\x87\x30\x8e\x63\xf8\x2f\x46\x65\xb8\x10\xae\x7b\xe0\x9a\x06\xfd\xfa\x8a\x1b\xe3\x8a\x71\x03\x85\x92\xd8\xee\xf9\xb4\x6b\xff\x84\xc7\x76\x19\x2c\x60\x7a\xcc\x20\x5d\x87\x83\xb0\x70\x26\x5a\x0c\xe8\x1e\x06\x82\x8b\xdd\xf0\xbc\x83\x9d\xbc\x42\xf8\x2a\x85\x20\x18\x6e\x3e\x59\x41\x0b\x7a\x62\x17\x06\xed\x7b\x6f\x8b\x71\x1b\x1d\xcf\xc5\xae\xc9\x25\xdc\x4c\xa7\xd3\xc9\x09\x13\xbb\xbd\x7a\xdf\xd4\x94\x36\x01\x93\x5b\x77\x25\xf6\xba\x75\x89\x23\xa5\x40\x74\xa5\x09\xc8\x95\x10\x3e\x67\x69\xb7\x92\x82\xdb\xe6\x49\x0a\xe1\xd5\xfc\x4c\x14\x1d\x68\x72\x20\xda\xb1\x79\xce\xe8\xfe\xd8\x44\x87\x3a\x3b\x5a\x1c\x5e\x1d\x18\xe5\xc0\x5e\xe7\x0d\x73\xd1\xf3\xcd\xf7\x1a\x3d\x32\xd7\xde\x5e\xc7\x3a\x1b\xf0\xef\xe9\x3c\xbf\xfa\x48\x31\xfa\xe9\xba\x31\xe5\xf8\x88\xd1\xc9\xfc\xd4\x36\xef\x2c\x6a\xca\x92\x15\x85\x2c\xb2\x05\x95\x02\x1a\x4f\x4c\xe2\x52\x75\x8d\xa1\x46\x59\xa0\xee\x52\x0a\x9f\xd9\x53\x02\x78\x60\x32\x5f\x55\x0e\xe1\x34\x90\xe8\x44\xb7\x73\xe0\xb0\xa0\x34\x0f\x78\x18\x0e\x64\x71\x79\x99\x92\x08\x00\x70\xe4\x09\x0e\xad\x07\x70\xa5\xc5\x28\x58\x6d\xb0\x80\x14\xfc\xa7\xf0\xf1\x24\x6a\x24\xbf\x1b\x4f\xc2\xf6\xfd\x98\x46\x37\x3f\xef\x6b\xc5\x8e\xf7\xe7\x29\x04\x89\xd5\xc0\x8b\x74\x14\xc0\xf3\x73\x7e\x48\xa1\x77\xb4\xd8\x73\x30\xdc\x0a\x90\xd8\x62\xe1\x9a\xa1\xbe\x68\xfb\x7b\x90\xb1\xfc\x76\xe5\xaa\xa1\x19\xe5\x5b\xe3\x13\xb2\x6c\xcd\x2c\xd3\x8e\xea\x64\x0e\xfb\xe5\x6d\xb5\x98\x93\x85\xe6\xe0\xcb\x52\xd7\x73\x85\xfe\x3b\x85\x7b\xcb\x94\x2e\x50\x87\x9a\x15\xbc\x31\x33\x78\x51\xdf\xcd\xff\xde\x7d\xc7\x71\x9d\xe1\x47\x59\xad\x35\x2e\x4e\x38\x6a\x5b\x8d\xcf\x21\x48\x62\x5a\xf0\x7b\x64\x7a\x61\x87\x9f\xe0\xe1\x4c\xff\x1b\xfa\x0f\xe4\xed\x78\xc5\x8b\x42\x20\x31\xbc\x27\x4f\x1e\x49\xf6\x1f\xfa\xd5\xe1\x91\xd0\x36\xbe\xf7\x7b\x76\x80\xc2\xe0\x23\x1b\xfa\x1e\xfa\x88\x00\x10\x92\xc8\xdc\xe9\xbc\xad\xb8\xdd\xb0\x1e\x39\x5d\xb4\x3f\xa8\x28\x1a\xed\x12\xae\x71\xd8\x02\xec\x12\x46\x86\x12\xc0\xc2\x8c\x26\x51\xd9\x54\x4c\xf2\xdf\x70\x4c\xc1\x69\xe2\x75\xe5\x9a\xf2\xc1\xe9\xbd\x7c\xc2\xcc\xbe\x5b\x3e\xea\x02\xdd\xa8\x55\xe2\xa8\xb3\xee\x8b\x7d\x81\x3f\x83\xe9\x7c\xf4\x89\x1a\x3a\x7f\x4a\x98\x31\x0d\xc3\x97\xb0\x8b\xc0\xa0\x15\x9d\xde\xcd\x65\x4c\x8f\x7c\x3b\xc3\x25\xe9\x52\x6d\xd2\xd1\xcd\xb4\x67\xd2\x1b\xda\xd9\x79\xd4\x62\xed\xc4\x18\xc4\x65\xe7\x9a\x0b\xb8\x99\x7e\x09\x6e\x7d\x4b\xe4\x48\x02\xab\x79\x8d\x05\xb0\xdc\xf2\x35\xfe\x3f\x08\xf2\x05\x94\xfc\xc9\x2c\x12\x0e\x3b\xe5\x39\x98\x1e\xf0\x4b\xb3\xbd\x6e\xff\x95\xfc\x0d\x62\xa7\xe1\xe7\x10\x9c\x15\xe4\x41\x24\x1e\x2d\x3c\x72\xed\x87\xfd\xde\x7d\x65\x0a\x8e\x03\x0b\xa5\xbc\xfd\x17\xd2\x49\x54\xda\x4a\x8c\x83\xc4\xba\x9f\xca\x10\xcf\x3d\x05\x47\xc0\x0f\x1f\xe6\x75\xbb\xc3\x6a\x86\x8a\x78\x3c\x2a\xb6\x60\x90\xa1\xf4\x05\x59\x97\x8e\xc0\x6e\xff\x8b\xa2\x38\x86\x1f\x2d\xd3\x16\x18\xfc\xf4\x0e\x9a\xba\x60\xd6\x7f\xcf\xa1\x20\xe9\xbf\x97\x74\x3f\x39\xca\x98\x36\xb0\x54\x7a\xc3\x74\xd1\x36\x69\x6c\x89\x5b\xf7\x3d\xa7\xcb\xff\x0c\xda\x77\x74\x8b\xad\x99\x18\x9f\x14\x7f\x4f\xc7\xa3\x68\x68\xf2\xd1\x24\x42\x96\x97\xa7\x0b\x5d\xc4\xea\xcf\x4d\xe1\x5b\x57\x07\x8c\x9f\x8e\x6d\xc9\xcd\x24\x62\xd6\xea\xf1\xe8\x00\x0c\xa3\x09\xd9\xf5\x6a\x50\x97\xf5\xdb\x93\x03\xb7\x7a\x8c\xc6\x3e\xa3\xee\xb3\x81\x6e\x79\x6e\xcc\xd8\xe3\x6a\x74\x39\xa0\x7d\x08\xab\xd1\xb3\x51\x6f\xa8\xbd\x7b\xef\xe5\x48\xcf\x72\x72\x40\x7a\x44\x5e\x36\x3a\x39\x9e\x15\xc5\x5b\xf2\x9f\x71\x70\xc6\xd3\x8f\xd1\x31\xe9\x95\xed\xef\xeb\x47\xb5\xec\x7f\x9b\xf1\x80\x8a\x79\x31\x9a\x44\xa6\xc9\x7c\x83\x62\xfc\xb2\xaf\xc2\xba\x65\x0e\xbc\xc7\xa1\xe0\x24\xa1\xa0\x23\x0e\x93\x8a\xf0\x28\x09\x79\x24\x6a\xb4\x47\x7a\xa9\x76\x97\xa4\xf0\xe9\xa4\xef\x6f\x7d\x63\x28\xc3\xf2\xfd\xff\x0d\x66\xc6\xb5\x13\xa0\xc5\xbb\x6b\xe9\xf8\xd6\xcd\x9b\xef\xdf\x0d\xda\x37\xbd\x47\x8c\x1d\xf5\xfe\xd7\x80\xe7\x9a\x25\x67\x7f\x7e\xb8\xd9\x6c\xa2\x95\x52\x2b\xe1\x7f\x78\xd8\x77\x53\x62\x56\xf3\xe8\x83\x09\x80\x99\xad\xcc\xa1\xc0\x25\xea\xc5\x80\x7c\xdb\x62\x49\x62\xff\xc3\xb8\x24\xf6\xbf\xfd\xfd\xbf\x00\x00\x00\xff\xff\xb2\x1e\x6f\x68\x0c\x2c\x00\x00")

func faucetHtmlBytes() ([]byte, error) {
	return bindataRead(
		_faucetHtml,
		"faucet.html",
	)
}

func faucetHtml() (*asset, error) {
	bytes, err := faucetHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "faucet.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc5, 0x8d, 0xb, 0x7a, 0xfd, 0x70, 0x68, 0x68, 0xd2, 0xd8, 0xf3, 0xf6, 0xac, 0x72, 0xed, 0xc2, 0x76, 0x18, 0x2d, 0x1, 0xe5, 0x3b, 0x55, 0xb, 0xce, 0xfc, 0xb6, 0xd5, 0x59, 0xc3, 0x94, 0x5b}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"faucet.html": faucetHtml,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"faucet.html": {faucetHtml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
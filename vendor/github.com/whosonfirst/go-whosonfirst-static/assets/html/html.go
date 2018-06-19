// Code generated by go-bindata.
// sources:
// templates/html/id.html
// DO NOT EDIT!

package html

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

var _templatesHtmlIdHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x5a\xfd\x72\xdb\xb8\x11\xff\xdb\x7e\x8a\x0d\x93\x69\x72\x33\xa1\x98\x5c\xae\x9d\xd6\x91\xd8\x0f\x27\x99\x7a\x9a\xeb\x65\xf2\xd1\x6b\xff\xba\x81\x88\x15\x89\x04\x04\x78\x00\x28\x59\x97\xfa\xb5\xfa\x00\x7d\xb2\x0e\x00\x7e\x80\x14\x65\xcb\x72\x9a\xcb\x4c\x2c\x12\xdc\xfd\xed\xee\x8f\x0b\xec\x02\xd2\xfc\xde\x8b\x1f\xce\xdf\xff\xeb\xcd\x4b\x28\x4c\xc9\xd3\xd3\xb9\xfd\x00\x4e\x44\xbe\x88\x50\x44\xe9\x29\xc0\xbc\x40\x42\xd3\xd3\x13\x80\xb9\x61\x86\x63\xfa\xf9\x33\xcc\xde\xbd\x79\x3b\xfb\x3b\x29\x11\xae\xae\xe0\xdf\xd0\x8e\x5c\x50\x7f\xff\x63\x21\x1f\x6a\xf8\x41\xc0\x2b\xa6\xb4\x99\x27\x5e\xcf\x41\x94\x68\x08\x14\xc6\x54\x31\xfe\x5c\xb3\xf5\x22\x3a\x97\xc2\xa0\x30\xf1\xfb\x6d\x85\x11\x64\xfe\x6e\x11\x19\xbc\x34\x89\x75\xe6\x39\x64\x05\x51\x1a\xcd\xe2\xc3\xfb\x57\xf1\xef\xa3\x00\x46\x90\x12\x17\x91\xc2\x15\x2a\x85\x2a\x50\x96\x8a\xe5\x4c\x44\x7b\x2c\xfe\x33\xfe\xf0\xe7\xf8\x5c\x96\x15\x31\x6c\xc9\x43\xa3\x17\x2f\x17\x7f\x88\x20\xd9\x31\x41\xaa\x8a\x63\x5c\xca\x25\xe3\x18\x6f\x70\x19\x93\xaa\x8a\x33\x52\x91\xa1\xfa\x16\xf5\xc1\xda\xda\x10\x53\xeb\x78\x49\x54\xac\xcd\x76\x00\xb3\xe4\x24\xfb\x34\x05\xf4\x57\x22\x68\x81\x9c\xbe\x52\x0c\x05\xe5\xdb\x90\x2e\x55\xe3\x94\xca\x9a\xe1\xa6\x92\xca\x04\xa2\x1b\x46\x4d\xb1\xa0\xb8\x66\x19\xc6\xee\xe6\x31\x30\xc1\x0c\x23\x3c\xd6\x19\xe1\xb8\x78\xfa\x18\x4a\x72\xc9\xca\xba\x0c\x06\x98\x18\x0e\xd4\x1a\x95\xbb\xb3\x24\x2c\x84\x74\xd6\x7b\xf3\x95\x92\x15\x2a\xb3\x5d\x44\xf6\x25\x9e\x71\xa2\x4d\x29\x29\x5b\x31\xa4\x81\x2f\x36\x71\x5e\x13\x6d\xbe\x6f\x1e\xc1\xd5\x55\x1b\xc5\x14\x94\xcc\xcf\xcc\x30\x4d\x88\x32\x2c\xe3\xa3\xd0\x07\x0a\x9a\x19\xfc\xc9\x92\x11\x68\x0d\xf3\xf3\x1a\x65\x97\xb8\x23\x7f\xc3\xd4\x67\x1a\x48\x97\xfd\x6f\x38\xc9\xd0\xfa\xe7\x9e\x88\x6e\xfc\x5c\xd6\xc2\xa8\x6d\x1f\xda\x84\x21\x8a\x3a\x53\xac\x32\x4c\x8a\x30\x1f\x8d\x1e\xcd\x25\xb8\x78\x61\x8d\x0e\x26\xdc\x35\xb0\xac\x24\x79\xe8\x7f\xc2\x68\x32\xd0\x9d\x55\x22\xff\xa3\x66\xbf\xe0\xe2\x52\x97\x7b\x5f\xa1\xd9\x30\x63\x50\x9d\x65\x44\xd1\x08\xd6\x84\xd7\xb8\x88\x74\x5d\x96\x44\x6d\xf7\x59\x6f\x75\x2c\xff\x81\x07\x7f\x22\x9c\xcb\x95\x29\xb0\xb2\x6c\xe9\x9b\xb4\x33\x85\xc4\x48\x75\x3c\x40\xad\xf8\x75\xf1\xdf\xa4\xfe\x75\x12\xa0\xb5\xf6\xff\xc8\x82\x16\xfb\x98\x54\x00\x8b\xc7\x99\xf8\x04\x0a\xf9\x22\x72\x0b\x95\x2e\x10\x4d\x04\x85\xc2\xd5\x22\x4a\x32\xad\x93\x92\x54\xbf\xa0\x98\x7d\xd4\xb3\x4c\xeb\x66\xd1\xf5\x71\x80\x56\xd9\x22\x4a\x3e\x92\x35\xf1\x03\xad\x6c\xc9\xac\x7c\x94\xce\x13\x3f\xee\xd2\x6e\xaf\x12\x97\x19\xe1\x2b\xa9\x48\x8e\xbb\x9a\x7e\xa9\xd8\xab\xab\x39\xab\xaa\x6d\x49\xaa\x59\xa6\xa4\xd6\x05\x61\x4a\xdf\xc2\xf4\xa6\x90\x5a\x8a\x95\x25\x7d\xb6\xd9\x6c\xc6\xa6\x0f\x55\x55\x28\x28\xaa\x91\xdd\x03\x75\x9b\x57\xc9\x70\xec\xf7\xe1\xe6\x33\x92\x15\x78\xac\xef\xb5\x62\xc7\x39\x2e\xd0\x1c\xa9\x48\x4a\xb6\xda\x0e\x75\xe1\x04\x4e\xe0\xe0\x57\x95\xa3\xfc\xa8\xe5\x4e\xa6\x1c\xaa\xce\x91\xac\x38\x9a\x59\x6d\x18\xdf\x65\xfd\x56\x10\x7e\xce\xdc\x0d\xa3\x20\x82\x72\x1c\xe7\x6d\xc7\x88\x9d\xa6\x37\xcf\xd2\x71\x26\xfb\xb9\x7a\x6b\xd5\x4c\x96\xa5\x14\x5e\xdb\x11\x7a\x7a\x5b\x04\xdb\xf9\xb0\x6c\xc6\x68\xb7\x5c\x1c\x44\x46\xaf\xb6\x77\xf6\xda\xa5\xb7\xe9\x1c\x7b\x10\xbf\x20\xc1\x86\x09\x2a\x37\x33\x42\xe9\xcb\x35\x0a\xf3\x9a\x69\x83\x02\xd5\xa3\x88\x4b\x42\xa3\xc7\xb0\xaa\x45\x66\xd7\x5d\xb0\xf7\x8f\xd0\xca\x7c\xf3\xd9\x69\x9e\x4c\x7b\x61\x5b\xa6\x47\xdf\x3c\x77\x22\x57\xee\xb3\xf7\x09\x60\x9e\xf8\xae\x19\x60\xbe\x94\x74\x0b\x94\x18\x12\x07\x40\x71\xad\x58\x8c\x82\x56\x92\xb5\x65\xe5\x05\x31\xe4\x65\x33\x62\xd7\x73\xab\x0c\xe0\xfe\xcc\x29\x5b\x43\xc6\x89\xd6\x8b\x28\x00\x89\x80\xd1\xc1\x40\xec\x6a\x62\xb4\x6b\xcd\xca\x8d\xaa\xc5\x8e\x4c\x45\x4c\xd1\x4b\xbd\x21\xa6\xd8\x27\xa7\x6c\xbf\x1e\x42\xbe\x71\x43\x21\x30\x27\x86\x99\x9a\x62\x2f\xf3\xba\x19\x09\x64\xa4\xc8\xc7\x42\xed\x50\x2f\x55\x32\x31\x81\xf6\x3d\x13\x2d\xe0\x50\x72\x17\xd3\x8a\x4e\xc0\x92\xcb\x29\x58\x72\xb9\x0b\x6b\x25\x27\x60\xc9\xe5\x00\xd6\x96\x4b\xe8\x5e\x96\x65\xa7\x24\x95\x4d\xbe\xf9\xbd\x38\x86\x39\x2b\xf3\x26\xbf\x77\xea\xae\x5e\xe7\x11\x14\xc8\xf2\xc2\x2c\xa2\x67\xbf\x7d\x62\x6b\x2f\xc4\x71\xda\xe2\x25\x94\xad\xdb\x9b\x76\xac\xf8\xd6\x22\xeb\x8a\x88\x89\xbc\x88\x5d\xbb\xeb\xf2\x7e\xd4\xaf\xcc\x13\xab\x92\x82\xd7\x9c\xcc\x9e\x98\x12\x83\x7e\x86\x39\x51\xd7\xe1\xec\xb5\x54\xb5\x2d\x4f\x94\x4e\x75\x41\x1d\x88\x18\x42\xf4\xd2\x5d\x6f\xe4\xea\x78\x23\x3f\x3b\xb5\xc1\x95\x84\xf3\x36\x08\xb6\x02\xfc\xb9\x21\x4d\x9f\xd7\xca\xe6\xdb\xec\x9d\x51\x4c\xe4\xaf\x38\xc9\x21\x7a\x1a\xc1\xd5\x95\x6f\x3d\xf6\xb8\xda\xec\xb8\x32\xaf\x1c\xa5\x3f\x22\x2c\x91\x33\x5c\x23\x98\x82\x69\x50\x98\x49\x45\xc1\x48\x58\x22\x34\x52\xb3\xc6\xa1\xc6\x0b\xe4\x1a\x6f\x76\xe5\xc9\xa1\xae\x08\x69\x06\xee\x08\x09\x36\xd1\x50\x1d\xe7\x98\x37\x6a\xef\x84\xcd\xab\x96\x50\x47\xe3\x49\x97\x4c\x36\x73\x9a\x9b\x93\x2e\x55\x43\xe7\x28\x1a\xc2\xb8\x8e\x26\xa9\x7f\x81\x95\xc2\x8c\x18\xa4\xc7\xb1\x4f\x3b\xfd\x28\x7d\x1f\x04\x57\x10\x0d\x4b\x44\x01\x25\x51\x9f\x90\x02\xd1\xd0\x8b\x8e\x83\x75\xe1\x9d\x9e\x4c\xf9\xf7\xae\xae\x50\x69\xa4\xfb\xfc\xb3\xff\x27\xed\xea\x4e\x11\x96\x5b\x30\x05\xc2\x4a\x72\x2e\x37\x4c\xe4\x8d\xac\x3e\x73\x6b\x3d\xcc\x6b\x3e\x39\x13\xec\xba\x8d\xaa\xe1\xcd\xfe\xfb\xfc\x19\x14\x11\x39\xc2\x03\xf6\x18\x1e\x68\x38\x5b\x78\x2f\x7b\x1f\xff\xb2\x6d\xdc\x72\xb8\x9c\xa5\x73\xd2\x56\x4d\xbf\x4a\x3c\xd0\x6e\xc5\x9a\x30\xc7\x28\x8c\x66\x3c\x5b\x6d\xf7\xaf\xfc\x1e\x28\x6d\xaf\xe6\x09\x49\xe7\x09\x67\xa1\xaf\x2d\xad\x3e\x4d\x6a\xde\x3c\x9b\xf9\x8f\x21\xf7\x96\xfc\x7d\xd4\x33\x91\x4f\x71\xef\x61\x42\xee\x3b\xca\xf5\x35\x7c\x1f\xc4\xf6\x8d\x4c\xeb\x36\x3d\x7f\x2d\x8e\x07\xfc\x76\xec\xce\x7c\x4a\x06\x0f\x9b\xc5\xbe\x9b\xae\x35\xdf\x99\x9f\x4b\xa2\x59\xe6\x02\xb7\xd1\xec\xdd\x1c\xee\x9d\x89\x8c\xf6\xab\xaf\x2b\x40\xed\x22\x4d\x04\x05\x66\xec\xeb\xb1\x85\x71\x8d\xf0\xe1\xed\x85\x43\x6a\x09\x2b\x8c\xa9\xf4\x59\x92\x58\x02\x66\x61\x57\x24\x55\x9e\xec\x34\x0e\x13\xa6\x15\x72\xdb\x65\x04\xb5\xc2\x0b\x5b\xaa\x66\x9e\xab\x71\x5e\x75\xad\x45\x14\x7f\xe7\x13\xc9\x46\x1d\xa4\xd1\x43\x0d\xbe\x23\x69\xe3\xce\x24\xc5\x34\xfe\x6e\x9e\xb8\x0b\xd8\x14\x2c\x2b\x80\x09\xca\xec\x5a\xa2\x81\x19\x37\xe9\xbb\xa0\xee\x17\x0c\x15\x51\x59\xb1\x8d\xd2\xb2\xe6\x86\x55\x1c\x1b\x44\x6d\xfd\x82\x25\x66\xa4\xd6\xb8\x4b\x43\xce\x4c\x51\x2f\x6d\x47\x1c\x36\xaa\xc9\x20\x4d\xbb\xbd\x5b\x62\x14\x62\x52\x12\x6d\x50\x25\xc1\xf0\x46\xae\xee\x7b\x6b\x3f\xd9\x17\xe3\xbc\xab\x2a\x14\x1a\xb4\x2c\xd1\xb0\x12\x75\xc0\xce\xc9\x44\x11\x0a\x18\x7a\x76\x1b\x86\x9e\x0d\x19\x2a\x91\x08\xc7\x4e\x69\x1b\x11\xe0\x98\x33\xc3\x4a\x62\x90\x6f\xa1\x20\x6b\x3c\x82\xb1\x0a\xa5\x7d\xf4\x55\x88\x23\x2a\xaf\x11\xc8\x52\xd6\xc6\x57\x4d\x2d\x95\x01\xb9\xb2\x37\x22\x3f\x9c\xc2\x6f\x6f\x43\xe1\xd3\x29\x0a\x37\xe8\xf8\x12\x0f\x0d\x2c\xa5\x29\x50\x21\x05\xa3\xb6\x76\x61\x33\x12\x2a\x26\x80\xca\x8d\x08\x4b\xfb\xc3\x3d\xe9\xc8\xca\x12\x29\x23\xa6\x65\xf7\xeb\xa7\x23\x67\x2b\xb4\xf1\x66\xb2\xac\xb8\x9b\x40\x34\x58\xd6\xae\x67\xf2\xe9\xcd\x4c\x5e\xec\x32\x69\x39\xb4\x54\x6d\x10\x88\xfa\x4a\x51\xd6\x22\x43\x65\x08\x13\x8e\x60\xff\x2e\x9b\x1a\xd5\xa4\xb5\x7f\x5b\x52\xe0\x38\x89\xda\x08\x2f\x8c\xa5\xc9\x4b\xfb\x16\x62\x5c\x66\x06\xf4\x7c\x99\x82\x33\x86\x4c\x77\x0b\xc1\xae\xd4\x78\x2a\x74\xcd\x54\x5b\x4d\x96\xb2\x16\xb6\x88\xc3\x52\x5e\x0e\x96\xff\xfb\x93\x3e\xe7\x68\xd7\x29\xb5\x85\xa9\xc1\xb8\x05\x5b\xca\xcb\x7e\xd5\x1f\x6c\xde\x1e\xc3\x9e\x8d\x5a\xf0\x20\xdc\x95\x0d\x87\x03\x79\xf7\xf2\xda\x1a\x56\x29\x26\x32\x56\x11\x0e\x19\x0a\xa3\x24\xa3\x77\x0e\xa5\x05\xea\xe3\x08\xb6\xb4\xbd\x57\x63\x97\xda\xba\xd6\xf6\xdf\xb6\xf4\x07\x3b\x45\x5b\xf1\xba\x36\x9c\xe3\xca\xc4\xba\x20\xea\x93\x2d\xef\x4d\x2b\x36\xd5\xa3\xf7\xb9\xdc\x05\xa2\x31\x33\x7d\xef\x79\x83\x56\x5c\x29\x34\x66\x6b\x77\x79\xd6\x87\xa6\x31\x3e\x44\x51\x91\x4d\xab\xd5\xb6\x8a\xfe\xba\xf9\x0c\xa2\x51\xb6\x90\x74\xe1\xb8\xc6\x67\x0a\xdc\xef\x0c\xc6\x71\x34\xe0\xc5\x33\x70\x47\x4a\xfe\x58\x27\x26\x9c\xe5\xe2\xcc\xd2\x04\xf7\x58\x59\x49\x65\x88\x30\xcf\xfd\x39\x48\x83\x93\xfa\xcf\xe0\x5d\x8f\x0c\x10\x91\x15\x52\x45\xe9\x7f\xff\xe3\x97\xb2\xe2\x59\xda\xb7\xb8\x00\x73\x43\x96\xcd\x97\x95\xee\xf6\x5e\x1c\x9f\x9e\xb8\x2f\x3f\x55\xea\x2e\xec\x65\x91\xfe\xe6\xfe\xd3\xdf\x3d\x79\x3e\x4f\x4c\x11\x8e\xbe\x73\xa6\xc6\xa3\xff\x20\xbc\xc6\xf1\xe0\xb9\x8d\x9f\xa2\xc8\x82\x27\xf3\xa4\x35\x62\xb7\xff\xad\x07\x7e\xcc\xea\xd8\x95\xd8\xef\xff\x1a\x9d\x56\x82\xee\xdf\x72\x05\xcd\xde\xd4\x6e\xd5\x26\xa9\xa1\xc7\x43\x5d\xe8\xf7\xaa\xc6\x2f\x00\xf3\x37\x61\xab\xe2\x18\xc7\x13\x12\x12\xd1\x5e\x7b\x36\x90\x68\x5b\x92\x8e\x22\xc3\xe9\x7e\x11\x2e\x3c\xd2\x5d\xa9\x68\x51\x0e\x65\xa2\x4d\x89\x7e\xa3\xdc\x66\xd2\xc1\x36\xa7\xb7\xf3\x9d\xe9\xa3\x80\xc6\x34\x58\x90\x7d\x92\xb7\x0d\xb5\xdf\xa5\xdf\x3a\xd4\xe9\x93\x81\x23\x42\x0d\x80\xa6\x42\x3d\x02\xe4\x58\x16\x5c\x63\x7b\x24\x0d\xa3\x5d\xfa\x1d\x42\xb0\x48\x77\x27\xc2\xa3\xec\x67\xe2\x04\xba\x22\x35\x4f\xfa\x25\xba\xab\x3a\x7b\x8a\x4b\xdf\x4d\xef\xad\x2f\x56\x25\x68\xba\xbb\xcb\xc9\x9e\xfc\xa0\x22\xd2\x96\xc4\xf6\x57\x0c\xc1\xe9\xfd\xa4\x75\x2d\x6b\x95\xa1\x2d\x5d\xfe\x22\xac\x5d\xcd\xa3\x43\x8b\x57\x77\x6c\x33\x3a\xeb\xb8\xd5\xd6\x3d\xbd\x8d\xf4\xce\x51\xd2\xa4\xdd\xe9\xb6\x3d\xb6\x06\x3a\xc0\xb7\x58\x49\xb8\xba\x4a\x96\x5c\x2e\xdb\xa6\x7d\x20\xb0\xe3\xdf\x97\x44\x1d\xc5\xd1\x9d\xd1\xf4\xbd\x4e\xd7\xe3\x0c\xce\xe5\xc7\xa7\xfe\x2b\x29\x0d\xaa\xfe\xe0\x7f\xcc\x44\xf3\x7d\xbc\x73\x7a\xf4\x53\x86\x74\xf4\xdb\x06\xd7\xd1\x5e\xf7\x45\x40\x30\xe2\xbe\x73\x5a\x4a\xba\x4d\x4f\xe7\x89\xff\x2d\xd7\xff\x02\x00\x00\xff\xff\x1e\xb1\x1c\x6e\xdc\x25\x00\x00")

func templatesHtmlIdHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlIdHtml,
		"templates/html/id.html",
	)
}

func templatesHtmlIdHtml() (*asset, error) {
	bytes, err := templatesHtmlIdHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/id.html", size: 9692, mode: os.FileMode(420), modTime: time.Unix(1524851647, 0)}
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
	"templates/html/id.html": templatesHtmlIdHtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"html": &bintree{nil, map[string]*bintree{
			"id.html": &bintree{templatesHtmlIdHtml, map[string]*bintree{}},
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

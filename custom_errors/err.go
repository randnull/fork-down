package custom_errors

import "errors"

var FileNotProvide = errors.New("path to file does not provide")
var ManifestNotProvide = errors.New("path to manifest does not provide")
var FileNotFound = errors.New("path to file does not exist")
var ManifestNotFound = errors.New("path to manifest does not exist")
var FileFormatError = errors.New("path file format != .bin")
var ManifestFormatError = errors.New("manifest file format != .bin")
var ErrorReadingManifest = errors.New("error reading manifest")
var ErrorWithReadFile = errors.New("error with reading file")
var ErrorDownloadChunk = errors.New("error download chunk")
var ErrorOpenFile = errors.New("error open output file")
var FatalError = errors.New("fatal error")
var ErrorWriteFile = errors.New("error to write file")

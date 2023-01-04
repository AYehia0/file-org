package fileorg

/*
This file contains all the different types of files and the directory corresponding to the file type.
*/
var DOC_FILES = []string{"txt", "pdf", "doc", "xml", "md", "xlsx", "csv", "py", "info", "log", "css", "srt"}
var VID_FILES = []string{"mp4", "avi", "flv", "mkv", "mov", "webm"}
var RAR_FILES = []string{"rar", "7z", "zip", "gz", "tar", "iso"}
var IMG_FILES = []string{"png", "jpg", "jpeg", "tif", "gif", "bmp", "webp"}
var MUSIC_FILES = []string{"mp3", "wav", "ogg", "aac"}

var EXT_FILES = map[string][]string{
	"docs":  DOC_FILES,
	"vids":  VID_FILES,
	"imgs":  IMG_FILES,
	"comps": RAR_FILES,
	"sound": MUSIC_FILES,
}

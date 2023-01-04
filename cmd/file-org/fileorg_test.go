package fileorg

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

var testDir = "./test_dir/"
var filesStr = `file file.mp3 file.mp4 file.MP4 file.srt file.py file.rar file.zip file.ZIP file.jpg file.jpeg file.png file.xlsx file.info file.`
var mappedFiles = map[string][]string{
	"info": {"file.info"},
	"jpeg": {"file.jpeg"},
	"jpg":  {"file.jpg"},
	"mp3":  {"file.mp3"},
	"mp4":  {"file.mp4", "file.MP4"},
	"png":  {"file.png"},
	"py":   {"file.py"},
	"rar":  {"file.rar"},
	"srt":  {"file.srt"},
	"xlsx": {"file.xlsx"},
	"zip":  {"file.ZIP", "file.zip"},
}

// Test the validation of paths
// calls isValidPath for a return bool valid path
func TestPath(t *testing.T) {
	paths := []string{
		".",
		"../file-org/",
		"/home/none/Downloads",
		"~/Downloads",
		"..",
	}

	for _, path := range paths {
		if !isValidPath(path) {
			t.Errorf("Expected %v to be valid!", path)
		}
	}
}

func flattenArray(arr [][]string) []string {
	tmpArr := make([]string, 0)
	for _, item := range arr {
		tmpArr = append(tmpArr, item...)
	}
	return tmpArr
}

// remove all the in the test dir and create new files then test.
func TestScanFiles(t *testing.T) {

	os.Mkdir(testDir, fs.ModePerm)

	for _, file := range strings.Split(filesStr, " ") {
		os.Create(testDir + file)
	}

	scannedFiles := scanFiles(testDir)

	// checking the keys
	scannedKeys := maps.Keys(scannedFiles)
	expectedKeys := maps.Keys(mappedFiles)

	sort.Strings(scannedKeys)
	sort.Strings(expectedKeys)

	if !reflect.DeepEqual(scannedKeys, expectedKeys) {
		t.Errorf("Expected scanned file structure to be %v, got %v", expectedKeys, scannedKeys)
	}

	// checking the values
	scannedValues := flattenArray(maps.Values(scannedFiles))
	expectedValues := flattenArray(maps.Values(mappedFiles))

	sort.Strings(scannedValues)
	sort.Strings(expectedValues)

	if !reflect.DeepEqual(scannedValues, expectedValues) {
		t.Errorf("Expected scanned file structure to be %v, got %v", expectedValues, scannedFiles)
	}

	os.RemoveAll(testDir)
}

func TestFolderCreation(t *testing.T) {

	os.Mkdir(testDir, fs.ModePerm)

	expectedDirs := maps.Keys(EXT_FILES)

	createTargetDirs(expectedDirs, testDir)

	dirs := make([]string, 0, len(expectedDirs))

	dirsFs, _ := ioutil.ReadDir(testDir)

	for _, dir := range dirsFs {
		dirs = append(dirs, dir.Name())
	}

	sort.Strings(dirs)
	sort.Strings(expectedDirs)

	if !reflect.DeepEqual(dirs, expectedDirs) {
		t.Errorf("Expected created directories to be %v, got %v", expectedDirs, dirs)
	}

	os.RemoveAll(testDir)
}

func TestFileMoving(t *testing.T) {

	os.Mkdir(testDir, fs.ModePerm)

	for _, file := range strings.Split(filesStr, " ") {
		os.Create(testDir + file)
	}

	createTargetDirs(maps.Keys(EXT_FILES), testDir)
	filesLog := moveFiles(mappedFiles, testDir, testDir)

	// check if all the files are moved correctly to the destination directory
	createdDirs, _ := ioutil.ReadDir(testDir)

	for _, dir := range createdDirs {
		if dir.IsDir() {
			dirName := dir.Name()
			insidePath := filepath.Join(testDir, dirName)
			expectedInsideFiles, _ := ioutil.ReadDir(insidePath)

			expectedFilePaths := filesLog[dirName]
			for _, file := range expectedInsideFiles {
				if !slices.Contains(expectedFilePaths, filepath.Join(testDir, dirName, file.Name())) {
					t.Errorf("Expected file : %s to be inside %s", file.Name(), filepath.Join(testDir, dirName))
				}
			}
		}
	}

	os.RemoveAll(testDir)
}

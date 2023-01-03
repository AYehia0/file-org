package fileorg

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/exp/maps"
)

// match the extension type
func matchExt(a string, exts map[string][]string) string {
	for k, v := range exts {
		for _, b := range v {
			if a == b {
				return k
			}
		}
	}
	return ""
}

func isValidPath(path string) bool {
	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		return true
	}
	return false
}

func scanFiles(pathToDir string) map[string][]string {

	// check if the pathToDir is valid one
	if !isValidPath(pathToDir) {
		log.Fatalf("Invalid Scanning Path! : %s", pathToDir)
	}

	filesMap := make(map[string][]string)
	files, err := os.ReadDir(pathToDir)
	if err != nil {
		log.Fatalf("Something went wrong while reading %s : Error %s", pathToDir, err)
	}

	for _, file := range files {
		// check if the file is a dir
		if file.IsDir() {
			continue
		}
		// check the file type
		for _, exts := range EXT_FILES {
			fileName := file.Name()
			for _, ext := range exts {
				// extType := strings.Split(fileName, ".")
				// fileExt := extType[len(extType)-1]
				fileExt := filepath.Ext(fileName)

				if fileExt == "" {
					continue
				}

				if strings.ToUpper(fileExt[1:]) == strings.ToUpper(ext) {
					filesMap[ext] = append(filesMap[ext], fileName)
				}
			}
		}
	}
	return filesMap
}

func createTargetDirs(dirs []string, orgDir string) {
	if !isValidPath(orgDir) {
		log.Fatalf("Invalid Target Path! : %s\n", orgDir)
	}
	for _, dir := range dirs {
		os.Mkdir(filepath.Join(orgDir, dir), os.ModePerm)
	}
}

func moveFiles(files map[string][]string, dirtyPath string, orgDir string) {
	for ext, files_ := range files {

		// match dir
		targetDir := matchExt(ext, EXT_FILES)

		if targetDir == "" {
			continue
		}

		for _, file := range files_ {
			currentPath := filepath.Join(dirtyPath, file)
			newPath := filepath.Join(orgDir, targetDir)
			err := os.Rename(currentPath, filepath.Join(newPath, file))

			if err != nil {
				log.Fatalf("Something went wrong while Moving %s : Error %s\n", currentPath, err)
			}
		}
	}
}

func Run(filesPath string, orgDir string) {
	files := scanFiles(filesPath)
	createTargetDirs(maps.Keys(EXT_FILES), orgDir)
	moveFiles(files, filesPath, orgDir)
}

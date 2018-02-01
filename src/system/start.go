package system

import (
	"path/filepath"
	"os"
	"io/ioutil"
	"strconv"
	"strings"
	"log"
)

var ROOT,pidFilename string

func init()  {
	ROOT = getParentDirectory(getCurrentDirectory())
	pidFilename = ROOT + "/pid/" + filepath.Base(os.Args[0]) + ".pid"
}


func savePid() {
	pid := os.Getpid()
	ioutil.WriteFile(pidFilename, []byte(strconv.Itoa(pid)), 0755)
}

func removePid()  {
	_ = os.Remove(pidFilename)
}


func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}


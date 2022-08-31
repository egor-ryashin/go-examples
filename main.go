package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	filepath.WalkDir(os.Args[1], func(path string, d fs.DirEntry, err error) error {
		fi, err := os.Stat(path)
		if !fi.IsDir() {
			dat, _ := os.ReadFile(path)
			sum := md5.Sum(dat)
			fmt.Printf("%s %s\n", hex.EncodeToString(sum[:]), path)
		}
		return nil
	})
}

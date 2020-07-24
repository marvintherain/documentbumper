package main 

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	"syscall"
	"log"
	"io"
	"github.com/schollz/progressbar"
)

func main() {
	files := WalkAndFilter(os.Args[1])
	
	lenFiles := int64(len(files))
	bar := progressbar.Default(lenFiles)
	bar.Add(1)
	for i := 0; i < len(files); i++ {
		suffix := "-bumped"
		src := files[i]
		dest := files[i] + suffix
		time.Sleep(40 * time.Millisecond)
		nBytes, err := copy(src, dest)
		bar.Add(1)
		if err != nil {
			fmt.Printf("Error copying file %s\n", files[i])
		} else {
			RemoveAndRename(src, dest)
			fmt.Printf("%s copied, %d bytes\n", files[i], nBytes)
		}
	}

}

// WalkAndFilter : walks the root directory filters for files older than 23 months
func WalkAndFilter(root string) []string {
	var files []string
	twentyThreeMonthsAgo := time.Now().AddDate(0, -23, 0)

    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() == false {
			d := info.Sys().(*syscall.Win32FileAttributeData)
			cTime := time.Unix(0, d.CreationTime.Nanoseconds())
			
			if cTime.Before(twentyThreeMonthsAgo) == true {
				files = append(files, path)

			}

		}
        return nil
    })
    if err != nil {
        panic(err)
	}
	return files 
}

// copy : copies the file from source to destination
func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
			return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
			log.Fatalf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
			return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
			return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	if err != nil {
		return 0, err
	 } 
	return nBytes, err
}

// RemoveAndRename : Renames destination to source
func RemoveAndRename(src, dest string) {
	err := os.Rename(dest, src)
	if err != nil {
		log.Fatalf("Could not rename due to: %s", err)
		} 
}
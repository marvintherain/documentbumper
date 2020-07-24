package main 

import (
    "fmt"
    "os"
	"path/filepath"
	"time"
	"syscall"
	"log"
	"io"
)

func main() {
	var files []string
	twentyThreeMonthsAgo := time.Now().AddDate(0, -2, 0)

    root := "test"
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
	
	for _, entry := range files {
		// fmt.Println(entry)
		suffix := "-bumped"
		src := entry
		dest := entry + suffix
		copy(src, dest)
		RemoveAndRename(src, dest)
	}

}

// copy : copies the file from source to destination
func copy(src, dst string) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
			log.Fatal(err)
	}

	if !sourceFileStat.Mode().IsRegular() {
			log.Fatalf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
			log.Fatal(err)
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
			log.Fatal(err)
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Successfully copied %d bytes", nBytes)
	}
}

// RemoveAndRename : Renames destination to source
func RemoveAndRename(src, dest string) {
	err := os.Rename(dest, src)
	if err != nil {
		log.Fatalf("Could not rename due to: %s", err)
		} 
}
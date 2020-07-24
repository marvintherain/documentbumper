package main 

import (
    "fmt"
    "os"
	"path/filepath"
	"time"
	"syscall"
)

func main() {
	var files []string
	twentyThreeMonthsAgo := time.Now().AddDate(0, -2, 0)
	// fileCreation := make(map[string]time.Time)

    root := "test"
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() == false {
			d := info.Sys().(*syscall.Win32FileAttributeData)
			cTime := time.Unix(0, d.CreationTime.Nanoseconds())
			
			if cTime.Before(twentyThreeMonthsAgo) == true {
				files = append(files, path)

			}
			// fileCreation[path] = cTime
		}
        return nil
    })
    if err != nil {
        panic(err)
	}
	
	for _, entry := range files {
		fmt.Println(entry)
	}

	// fmt.Println(fileCreation)

    // for path, entry := range fileCreation {
	// 	fmt.Println(path)
	// 	fmt.Println(entry)
    // }
}
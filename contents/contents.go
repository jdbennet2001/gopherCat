package contents

import (

	"fmt"
	"os"
	"path/filepath"
)

//Functions that start with capitals are exported
func Walk(rootpath string)[]string {

	var results []string

	err := filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {

		if filepath.Ext(path) == ".cbz"{
			results = append( results, path )
			fmt.Println(path)
		}

		return nil	// No error...
	})
	if err != nil {
		fmt.Printf("walk error [%v]\n", err)
	}

	return results

}
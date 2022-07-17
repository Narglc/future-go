package fuliba_sort

import (
	"log"
	"os"
)

func Save2File(fname string, ctx []byte) {
	var file *os.File
	_, err := os.Lstat(fname)
	if err != nil {
		file, err = os.Create(fname)
	} else {
		file, err = os.OpenFile(fname, os.O_APPEND, 7)
		if err != nil {
			log.Fatal(err)
		}
	}
	file.Write(ctx)
}

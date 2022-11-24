package main

import (
	"fmt"
	"os"
	"path/filepath"

	CmdYQ3 "github.com/mikefarah/yq/v3/cmd"
	CmdYQ4 "github.com/mikefarah/yq/v4/cmd"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	f, err := os.CreateTemp("", "sample")
	check(err)

	fmt.Println("Temp file name:", f.Name())

	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)

	yq4()

	yq3()
}
func yq4() {
	cmd := CmdYQ4.New()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func yq3() {
	cmd := CmdYQ3.New()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

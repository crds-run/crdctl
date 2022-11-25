package main

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//go:embed tusk.yaml
var fileByte []byte

//go:embed embed/*
var folder embed.FS

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	f, err := os.CreateTemp("", "sample")
	check(err)

	_, err = f.Write(fileByte)
	check(err)

	// fmt.Println(f.Name())

	// dname, err := os.MkdirTemp("", "sampledir")
	// check(err)
	// fmt.Println("Temp dir name:", dname)

	// defer os.RemoveAll(dname)

	// fname := filepath.Join(dname, "file1")
	// err = os.WriteFile(fname, fileByte, 0777)
	// check(err)

	content1, _ := folder.ReadFile("embed/init.sh")
	shell(string(content1))
	shell("go run github.com/rliebz/tusk@v0.6.4 --file " + f.Name() + " " + strings.Join(os.Args[1:], " "))
	defer os.Remove(f.Name())

	// content2, _ := folder.ReadFile("tusk.yaml")
	// print(string(content2))

	// print(string(fileByte))
}

func shell(theargs string) {
	// cmd := exec.Command(os.Getenv("SHELL"), "-c", " "+theargs)
	cmd := exec.Command("go", "run", "mvdan.cc/sh/v3/cmd/gosh@v3.5.1", "-c", " "+theargs)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		// os.Exit(1)
	}
}

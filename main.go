package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func main() {

	path, _ := os.Getwd()
    files, _ := ioutil.ReadDir(path)

	rand.Seed(time.Now().UnixNano())
	min := 0
    max := len(files)
	rand :=  rand.Intn(max - min) + min

	fmt.Println(files[rand].Name())
	
	var i string
    fmt.Scanln(&i)
}
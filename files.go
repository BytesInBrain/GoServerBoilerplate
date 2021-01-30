package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
)

func main() {
	fileBytes, err := ioutil.ReadFile("mol.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileBytes)
	fmt.Println(string(fileBytes))
	fmt.Println(reflect.TypeOf(fileBytes))

	s := "Fickk"
	errl := ioutil.WriteFile("exaj.txt", []byte(s), 0644)
	if errl != nil {
		log.Fatal(errl)
	}

	files, err3 := ioutil.ReadDir("/home")
	if err3 != nil {
		log.Fatal(err3)
	}
	for _, file := range files {
		fmt.Println(file.Mode(), file.Name())
	}
}

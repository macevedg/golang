package main

import (
	"encoding/json"
	"fmt"
	"unsafe"

	cha "github.com/macevedg/concurrency/channels"
)

type Person struct {
	name string
}

func main() {
	p :=
		Person{name: "macevedg"}

	if p == nil {
		fmt.Println("is nil")
	}
}

func mainold() {

	//numm := int64(1)

	ch := make(chan string, 2)

	go cha.Execute(ch)

	for msg := range ch {
		fmt.Println(msg)

		if msg == "quit" {
			close(ch)
			break
		}
	}

	/*
		loop := true
		for loop {
			select {
			case paramm := <-ch:
				{
					if paramm == "quit" {
						close(ch)
						loop = false
						break
					}
					fmt.Println(paramm)
					fmt.Printf("%v\n", ch)
				}
			}

		}*/

	mapBool := make(map[string]bool, 0)
	mapStruct := make(map[string]struct{}, 0)

	for i := 0; i < 1000; i++ {

		key := fmt.Sprintf("key%d", i)

		mapBool[key] = true
		mapStruct[key] = struct{}{}
	}

	fmt.Printf("bytes bool: %v\n", unsafe.Sizeof(mapBool))
	fmt.Printf("bytes struct: %v\n", unsafe.Sizeof(mapStruct))

	mapBoolBytes, _ := json.Marshal(mapBool)
	mapStructBytes, _ := json.Marshal(mapStruct)

	fmt.Printf("bytes bool: %v\n", unsafe.Sizeof(mapBoolBytes))
	fmt.Printf("bytes struct: %v\n", unsafe.Sizeof(mapStructBytes))

	boool := true
	structt := struct{}{}

	fmt.Printf("bytes bool: %v\n", unsafe.Sizeof(boool))
	fmt.Printf("bytes struct: %v\n", unsafe.Sizeof(structt))

	fmt.Println(checkMap("key"))
	fmt.Println(checkMap("key222"))
}

func checkMap(check string) bool {
	mm := make(map[string]bool, 0)

	mm["key"] = true

	return mm[check]
}

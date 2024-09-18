package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	result, err := devide(100.0, 10.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println(result)

}

func devide(x, y float32) (float32, error) {
	var result float32
	if y == 0 {
		return result, errors.New("cannot divide by zero")
	}
	result = x / y
	return result, nil
}

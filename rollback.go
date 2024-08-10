package main

import (
	"errors"
	"fmt"
)

var A int = 0

// there are two cases in this file
// simple case shows you how to use defer and same result of two patterns
// comples case shows you when to use this pattern
func main() {
	runSimpleCase()
	runComplexCase()
}

// setAComponent is assumed to have simple logic and be easy to set and rollback
// So in two cases, we do before setBComponent or setCComponent
func setAComponent() error {
	A = A + 1
	return nil
}

// rollbackAComponent does the exact opposite action what setAComponet does
func rollbackAComponent() error {
	A = A - 1
	return nil
}

// setBComponent is used in simpleCase.go which has just two options
func setBComponent() error {
	return errors.New("sample error")
	// return nil
}

// setCComponent is used in complexCase.go which has more than three options
func setCComponent() (Response, error) {
	// return Response{200, ""}, nil
	// return Response{200, ""}, errors.New("sample error")
	return Response{404, "not found error"}, nil
}

type Response struct {
	status  int
	message string
}

func printABefore() {
	fmt.Println("Before execution, A is ", A)
}

func printAAfter() {
	fmt.Println("After execution, A is ", A)
	fmt.Println("")
}
func resetA() {
	A = 0
}

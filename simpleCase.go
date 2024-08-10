package main

import "fmt"

func runSimpleCase() {
	resetA()
	fmt.Println("Running simple case")
	printABefore()
	simpleWithoutDefer()
	printAAfter()
	resetA()

	printABefore()
	simpleWithDefer()
	printAAfter()
	fmt.Println("Simple case complete")
}

func simpleWithoutDefer() error {
	err := setAComponent()
	if err != nil {
		return nil
	}
	err = setBComponent()
	if err != nil {
		fmt.Println("failed to set b, reason : ", err)
		rollbackAComponent()
		return err
	}

	return nil
}

func simpleWithDefer() (err error) {
	err = setAComponent()
	if err != nil {
		return err
	}

	defer func() {
		handleError1(err)
	}()
	return setBComponent()
}

func handleError1(err error) {
	if err != nil {
		fmt.Println("failed to set b, reason : ", err)
		rollbackAComponent()
	}
}

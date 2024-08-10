package main

import (
	"errors"
	"fmt"
)

func runComplexCase() {
	resetA()
	fmt.Println("Running complex case")
	printABefore()
	complexWithoutDefer()
	printAAfter()
	resetA()

	printABefore()
	complexWithDefer()
	printAAfter()
	fmt.Println("Complex case complete")
}

func complexWithoutDefer() error {
	err := setAComponent()
	if err != nil {
		return nil
	}
	resp, err := setCComponent()
	if err != nil {
		fmt.Println("failed to set c, reason : ", err)
		rollbackAComponent()
		return err
	}

	if resp.status != 200 {
		fmt.Println("failed to set c, reason : ", resp.message)
		rollbackAComponent()
		return errors.New(resp.message)
	}

	return nil
}

func complexWithDefer() (err error) {
	err = setAComponent()
	if err != nil {
		return err
	}

	defer func() {
		handleError2(err)
	}()
	resp, err := setCComponent()
	if err != nil {
		return err
	}

	if resp.status != 200 {
		err = errors.New(resp.message)
	}

	return err
}

func handleError2(err error) {
	if err != nil {
		fmt.Println("failed to set c, reason : ", err)
		rollbackAComponent()
	}
}

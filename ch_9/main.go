package main

import (
	"errors"
	"fmt"
	"os"
)

func cal(num , den int)(int ,int, error){
	if den == 0{
		return 0, 0, errors.New("DEN")
	}
	return num / den , num % den, nil
}

// Define sentinel error (package level)
var ErrNotFound = errors.New("not found")
var ErrInvalidInput = errors.New("invalid input")

func findUser(id int) (string, error) {
	if id <= 0 {
		return "", ErrInvalidInput
	}
	if id > 100 {
		return "", ErrNotFound
	}
	return "User", nil
}

type MyError struct{
	Code int
	Mess string
}

func (e MyError) Error() string{
	return fmt.Sprint("code %d : %s",e.Code,e.Mess)
}

func dosom() error {
	return MyError{Code: 404,Mess: "not found"}
}

func checkFile(name string) error {
    _, err := os.Open(name)
    if err != nil {
        return fmt.Errorf("checkFile: %w", err)  // wrapped!
    }
    return nil
}

type MnError struct {
    Code int
}

func (e *MnError) Erroras() string {
    return fmt.Sprintf("error code: %d", e.Code)
}

func doWork() error {
    return fmt.Errorf("wrapped: %w", &MnError{Code: 500})
}

func mustHaveValue(x *int) {
    if x == nil {
        panic("x cannot be nil!")
    }
    fmt.Println(*x)
}

func main() {
	err := errors.New("something went wrong")
	fmt.Println(err)
	fmt.Println(cal(0,0))

	// Check for specific error
	_, err = findUser(-1)
	if err == ErrInvalidInput {
		fmt.Println("Bad input!")
	}

	errr := dosom()
	fmt.Println(errr)

	erris := checkFile("missing.txt")

	// ❌ Won't work (error is wrapped!)
	if erris == os.ErrNotExist {
		// This won't match!
	}

	// ✅ Works! Checks inside wrapped errors
	if errors.Is(erris, os.ErrNotExist) {
		fmt.Println("File doesn't exist!")
	}

	mustHaveValue(nil)
	fmt.Println("hi")
}

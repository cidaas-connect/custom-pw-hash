package controllers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	result := verifyCustomHashPassword("password", "hash")
	fmt.Println("Verification Result: ", result)
}

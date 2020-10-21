package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	fmt.Printf("Start TestMain func with %v and %T\n", m, m)
	code := m.Run()
	fmt.Printf("After running TestMain code is %v and it is of type %T\n", code, code)
	os.Exit(code)
}

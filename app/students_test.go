package app

import (
	"fmt"
	"testing"
)

func TestStudentIsExist(t *testing.T) {
	a := StudentIsExist(9999)
	fmt.Println(a)
}
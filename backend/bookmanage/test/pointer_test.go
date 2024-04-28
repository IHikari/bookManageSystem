package test

import (
	"bookmanage/model"
	"fmt"
	"reflect"
	"testing"
)

func TestPointer(t *testing.T) {

	var book []model.Book

	fmt.Println(reflect.TypeOf(book))
	fmt.Println(reflect.TypeOf(&book))
}

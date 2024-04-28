package test

import (
	"bookmanage/common"
	"fmt"
	"testing"
)

func TestJwtToken(t *testing.T) {

	token, claims, err := common.ParseToken(
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJTY2hvb2xDYXJkIjoiMTExMTEiLCJleHAiOjE3MTExMzA5NjQsImlhdCI6MTcxMTEyMzc2NCwiaXNzIjoiaGlrYXJpIiwic3ViIjoidXNlciB0b2tlbiJ9.jD1fZ3JFoUGeiKapXlklFfc_w3Wf9aSkQ98tUmxEZVM")

	fmt.Println(token.Raw)
	fmt.Println(claims.ID)
	fmt.Println(err)
}

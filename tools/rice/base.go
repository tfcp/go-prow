package main

import (
	rice "github.com/GeertJohan/go.rice"
)

func init() {
	_, _ = rice.FindBox("../web/dist")
	return
}

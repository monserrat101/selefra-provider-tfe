package main

import (
	"github.com/selefra/selefra-provider-sdk/doc_gen"
	"github.com/selefra/selefra-provider-tfe/resources"

	"testing"
)

func Test(t *testing.T) {
	if err := doc_gen.New(resources.GetSelefraProvider(), "./tables").Run(); err != nil {
		panic(err)
	}
}

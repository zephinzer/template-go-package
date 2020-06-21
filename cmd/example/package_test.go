package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type Tests struct {
	suite.Suite
}

func TestExample(t *testing.T) {
	suite.Run(t, &Tests{})
}

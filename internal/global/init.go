package global

import (
	"time"

	"github.com/ptcar2009/ginkgo/internal/failer"
	"github.com/ptcar2009/ginkgo/internal/suite"
)

const DefaultTimeout = time.Duration(1 * time.Second)

var Suite *suite.Suite
var Failer *failer.Failer

func init() {
	InitializeGlobals()
}

func InitializeGlobals() {
	Failer = failer.New()
	Suite = suite.New(Failer)
}

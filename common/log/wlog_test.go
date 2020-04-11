package log

import (
	"fmt"
	"testing"
)

func TestWzlog(t *testing.T) {
	fmt.Sprintf("wayay,%v:", "abc")
	wLog := WLog{}
	wLog.Debug("helloworld")

	DEBUG("hehe")
}

package gomaskari

import (
	"context"
	"testing"
)

func Test_LetsGetShitDone(t *testing.T) {
	LetsGetShitDone(context.Background(), "test_scripts/main.go")
}

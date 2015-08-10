package main

import (
	"os"
	"testing"
	"time"
)

func TestExec(t *testing.T) {
	execWithTimeout("dummyproc", "10 a s d f g", nil, os.Stdout, 5*time.Second)
}
package example

import (
    "testing"
    . "./test_helpers"
)

func TestMain(t *testing.T) {
    output := CaptureStdout(Main)
    AssertEqual(t, output, "Hello, world.\n")
}

package test_helpers

import (
    "bytes"
    "io"
    "os"
    "testing"
)

// Test Helper
func CaptureStdout(block func()) string {
    old := os.Stdout // keep backup of the real stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    block()

    outC := make(chan string)
    // copy the output in a separate goroutine so printing can't block indefinitely
    go func() {
        var buf bytes.Buffer
        io.Copy(&buf, r)
        outC <- buf.String()
    }()

    // back to normal state
    w.Close()
    os.Stdout = old // restoring the real stdout
    out := <-outC

    // return our temp stdout
    return out
}

// - [In Go, how do I capture stdout of a function into a string? \- Stack Overflow]
//   (https://stackoverflow.com/questions/10473800/in-go-how-do-i-capture-stdout-of-a-function-into-a-string)

func AssertEqual(t *testing.T, a interface{}, b interface{}) {
  if a != b {
    t.Fatalf("%s != %s", a, b)
  }
}

// - [How to assert values in golang unit tests]
//   (https://gist.github.com/samalba/6059502)

package test_helpers

import (
    "bytes"
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "testing"
)

// Test Helper
func MockStdin(standardInput string, block func()) string {
    content := []byte(standardInput)

    tmpfile, err := ioutil.TempFile("", "example")
    if err != nil { fmt.Println(err) }

    defer func() {
        if err := tmpfile.Close(); err != nil { fmt.Println(err) }
    }()
    defer os.Remove(tmpfile.Name()) // clean up

    if _, err := tmpfile.Write(content); err != nil { fmt.Println(err) }
    if _, err := tmpfile.Seek(0, 0); err != nil { fmt.Println(err) }

    oldStdin := os.Stdin
    defer func() { os.Stdin = oldStdin }() // Restore original Stdin

    os.Stdin = tmpfile

    return CaptureStdout(block)
}

// - [testing \- Fill os\.Stdin for function that reads from it \- Stack Overflow]
//   (https://stackoverflow.com/questions/46365221/fill-os-stdin-for-function-that-reads-from-it/46365584#46365584)

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

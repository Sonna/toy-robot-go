package robot

import (
    "testing"
    . "test_helpers"
)

func TestRobotCLIMain(t *testing.T) {
    subject := "REPORT\nEXIT\n"
    osArgs := []string{"bin/main"}
    output := MockStdin(subject, func() { Main(osArgs) })

    AssertEqual(t, output, "0,0,NORTH\n")
}

func TestRobotCLIMainWithComplexInput(t *testing.T) {
    subject := "PLACE 3,2,SOUTH\nREPORT\nMOVE\nRIGHT\nMOVE\nREPORT\nEXIT\n"
    osArgs := []string{"bin/main"}
    output := MockStdin(subject, func() { Main(osArgs) })

    AssertEqual(t, output, "3,2,SOUTH\n2,1,WEST\n")
}

func TestRobotCLIMainWithExampleA(t *testing.T) {
    subject := []string{"bin/main", "../../examples/example_a.txt"}
    output := CaptureStdout(func() { Main(subject) })
    AssertEqual(t, output, "0,0,NORTH\n")
}

func TestRobotCLIMainWithExampleB(t *testing.T) {
    subject := []string{"bin/main", "../../examples/example_b.txt"}
    output := CaptureStdout(func() { Main(subject) })
    AssertEqual(t, output, "0,0,WEST\n")
}

func TestRobotCLIMainWithExampleC(t *testing.T) {
    subject := []string{"bin/main", "../../examples/example_c.txt"}
    output := CaptureStdout(func() { Main(subject) })
    AssertEqual(t, output, "3,3,NORTH\n")
}

package robot

import (
    "testing"
    . "test_helpers"
    // . "./test_helpers"
)

func TestRobotConstructor(t *testing.T) {
    subject := Robot{X: 0, Y: 0, Facing: "NORTH"}

    AssertEqual(t, subject.X, 0)
    AssertEqual(t, subject.Y, 0)
    AssertEqual(t, subject.Facing, "NORTH")
}

func TestRobotReport(t *testing.T) {
    subject := Robot{X: 0, Y: 0, Facing: "NORTH"}

    output := CaptureStdout(subject.report)
    AssertEqual(t, output, "0,0,NORTH\n")
}

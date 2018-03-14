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

func TestRobotLeft(t *testing.T) {
    subject := Robot{X: 0, Y: 0, Facing: "NORTH"}
    subject.left()
    AssertEqual(t, subject.Facing, "WEST")
}

func TestRobotLeftWhenFacingWest(t *testing.T) {
    subject := Robot{X: 0, Y: 0, Facing: "WEST"}
    subject.left()
    AssertEqual(t, subject.Facing, "SOUTH")
}

func TestRobotLeftWhenFacingSouth(t *testing.T) {
    subject := Robot{X: 0, Y: 0, Facing: "SOUTH"}
    subject.left()
    AssertEqual(t, subject.Facing, "EAST")
}

func TestRobotLeftWhenFacingEast(t *testing.T) {
    subject := Robot{X: 0, Y: 0, Facing: "EAST"}
    subject.left()
    AssertEqual(t, subject.Facing, "NORTH")
}

func TestRobotRight(t *testing.T) {
    subject := Robot{X: 0, Y: 0, Facing: "NORTH"}
    subject.right()
    AssertEqual(t, subject.Facing, "EAST")
}

func TestRobotRightWhenFacingEast(t *testing.T) {
    subject := Robot{X: 0, Y: 0, Facing: "EAST"}
    subject.right()
    AssertEqual(t, subject.Facing, "SOUTH")
}

func TestRobotRightWhenFacingSouth(t *testing.T) {
    subject := Robot{X: 0, Y: 0, Facing: "SOUTH"}
    subject.right()
    AssertEqual(t, subject.Facing, "WEST")
}

func TestRobotRightWhenFacingWest(t *testing.T) {
    subject := Robot{X: 0, Y: 0, Facing: "WEST"}
    subject.right()
    AssertEqual(t, subject.Facing, "NORTH")
}

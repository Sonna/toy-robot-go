package robot

import (
    "strings"
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

func TestRobotMove(t *testing.T) {
    subject := Robot{X: 0, Y: 0, Facing: "NORTH"}
    subject.move()
    AssertEqual(t, subject.X, 0)
    AssertEqual(t, subject.Y, 1)
    AssertEqual(t, subject.Facing, "NORTH")
}

func TestRobotMoveEast(t *testing.T) {
    subject := Robot{X: 0, Y: 0, Facing: "EAST"}
    subject.move()
    AssertEqual(t, subject.X, 1)
    AssertEqual(t, subject.Y, 0)
    AssertEqual(t, subject.Facing, "EAST")
}

func TestRobotMoveSouth(t *testing.T) {
    subject := Robot{X: 0, Y: 0, Facing: "SOUTH"}
    subject.move()
    AssertEqual(t, subject.X, 0)
    AssertEqual(t, subject.Y, 0)
    AssertEqual(t, subject.Facing, "SOUTH")
}

func TestRobotMoveWest(t *testing.T) {
    subject := Robot{X: 0, Y: 0, Facing: "WEST"}
    subject.move()
    AssertEqual(t, subject.X, 0)
    AssertEqual(t, subject.Y, 0)
    AssertEqual(t, subject.Facing, "WEST")
}

func TestRobotMoveDoesNotFallOffTableAt04West(t *testing.T) {
    subject := Robot{X: 0, Y: 4, Facing: "WEST"}
    subject.move()
    AssertEqual(t, subject.X, 0)
    AssertEqual(t, subject.Y, 4)
    AssertEqual(t, subject.Facing, "WEST")
}

func TestRobotMoveDoesNotFallOffTableAt04North(t *testing.T) {
    subject := Robot{X: 0, Y: 4, Facing: "NORTH"}
    subject.move()
    AssertEqual(t, subject.X, 0)
    AssertEqual(t, subject.Y, 4)
    AssertEqual(t, subject.Facing, "NORTH")
}

func TestRobotMoveDoesNotFallOffTableAt44North(t *testing.T) {
    subject := Robot{X: 4, Y: 4, Facing: "NORTH"}
    subject.move()
    AssertEqual(t, subject.X, 4)
    AssertEqual(t, subject.Y, 4)
    AssertEqual(t, subject.Facing, "NORTH")
}

func TestRobotMoveDoesNotFallOffTableAt44East(t *testing.T) {
    subject := Robot{X: 4, Y: 4, Facing: "EAST"}
    subject.move()
    AssertEqual(t, subject.X, 4)
    AssertEqual(t, subject.Y, 4)
    AssertEqual(t, subject.Facing, "EAST")
}

func TestRobotMoveDoesNotFallOffTableAt40East(t *testing.T) {
    subject := Robot{X: 4, Y: 0, Facing: "EAST"}
    subject.move()
    AssertEqual(t, subject.X, 4)
    AssertEqual(t, subject.Y, 0)
    AssertEqual(t, subject.Facing, "EAST")
}

func TestRobotMoveDoesNotFallOffTableAt40South(t *testing.T) {
    subject := Robot{X: 4, Y: 0, Facing: "SOUTH"}
    subject.move()
    AssertEqual(t, subject.X, 4)
    AssertEqual(t, subject.Y, 0)
    AssertEqual(t, subject.Facing, "SOUTH")
}

func TestRobotPlace(t *testing.T) {
    subject := Robot{X: 0, Y: 0, Facing: "NORTH"}
    subject.place("1,2,EAST")
    AssertEqual(t, subject.X, 1)
    AssertEqual(t, subject.Y, 2)
    AssertEqual(t, subject.Facing, "EAST")
}

func TestRobotPlaceHandleErrorX(t *testing.T) {
    subject := Robot{X: 0, Y: 0, Facing: "NORTH"}
    output := CaptureStdout(func () {
        subject.place("A,2,EAST")
    })
    AssertEqual(t,
        strings.Contains(output, "strconv.Atoi: parsing \"A\": invalid syntax"),
        true)
    AssertEqual(t, subject.X, 0)
    AssertEqual(t, subject.Y, 0)
    AssertEqual(t, subject.Facing, "NORTH")
}

func TestRobotPlaceHandleErrorY(t *testing.T) {
    subject := Robot{X: 0, Y: 0, Facing: "NORTH"}
    output := CaptureStdout(func () {
        subject.place("1,B,EAST")
    })
    AssertEqual(t,
        strings.Contains(output, "strconv.Atoi: parsing \"B\": invalid syntax"),
        true)
    AssertEqual(t, subject.X, 0)
    AssertEqual(t, subject.Y, 0)
    AssertEqual(t, subject.Facing, "NORTH")
}

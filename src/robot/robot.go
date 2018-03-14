package robot

import (
    "fmt"
    "strconv"
    "strings"
)

type Robot struct {
    X int
    Y int
    Facing string
}

var movement = func() map[string]map[rune]int {
  return map[string]map[rune]int {
    "NORTH": map[rune]int{'x': 0, 'y': 1},
    "SOUTH": map[rune]int{'x': 0, 'y': -1},
    "EAST": map[rune]int{'x': 1, 'y': 0},
    "WEST": map[rune]int{'x': -1, 'y': 0},
  }
}

var turn = func() map[string]map[string]string {
    return map[string]map[string]string {
        "NORTH": map[string]string{"LEFT": "WEST", "RIGHT": "EAST"},
        "SOUTH": map[string]string{"LEFT": "EAST", "RIGHT": "WEST"},
        "EAST": map[string]string{"LEFT": "NORTH", "RIGHT": "SOUTH"},
        "WEST": map[string]string{"LEFT": "SOUTH", "RIGHT": "NORTH"},
      }
}

func (r *Robot) report() {
    fmt.Printf("%d,%d,%s\n", r.X, r.Y, r.Facing)
}

func (r *Robot) left() {
    r.Facing = turn()[r.Facing]["LEFT"]
}

func (r *Robot) right() {
    r.Facing = turn()[r.Facing]["RIGHT"]
}

func (r *Robot) move() {
    if x := r.X + movement()[r.Facing]['x']; x > 0 && x < 4 {
        r.X = x
    }

    if y := r.Y + movement()[r.Facing]['y']; y > 0 && y < 4 {
        r.Y = y
    }
}

func (r *Robot) place(rawCoordinates string) {
    coordinates := strings.Split(rawCoordinates, ",")

    x, _xErr := strconv.Atoi(coordinates[0])
    if _xErr != nil { // handle error
        fmt.Println(_xErr)
        return
    }

    y, _yErr := strconv.Atoi(coordinates[1])
    if _yErr != nil { // handle error
        fmt.Println(_yErr)
        return
    }

    r.X = x
    r.Y = y
    r.Facing = coordinates[2]
}

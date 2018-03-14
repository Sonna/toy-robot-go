package robot

import "fmt"

type Robot struct {
    X int
    Y int
    Facing string
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

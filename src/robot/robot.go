package robot

import "fmt"

type Robot struct {
    X int
    Y int
    Facing string
}

func (r *Robot) report() {
    fmt.Printf("%d,%d,%s\n", r.X, r.Y, r.Facing)
}

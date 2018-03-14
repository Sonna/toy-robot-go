package robot

import (
    "bufio"
    "os"
    "strings"
)

func Main(args []string) {
    inputSource := os.Stdin
    robot := Robot{X: 0, Y: 0, Facing: "NORTH"}

    if len(args) > 1 {
        filename := args[1]
        file, _ := os.Open(filename)
        defer file.Close()

        inputSource = file
    }

    scanner := bufio.NewScanner(inputSource)
    for scanner.Scan() {
        line := scanner.Text()
        if line == "EXIT" {
            break
        }
        input := strings.Split(line, " ")
        robot.exec(input...)
    }
}

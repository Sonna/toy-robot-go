Toy Robot Simulator
===================

Description
-----------

- The application is a simulation of a toy robot moving on a square tabletop,
  of dimensions 5 units x 5 units.
- There are no other obstructions on the table surface.
- The robot is free to roam around the surface of the table, but must be
  prevented from falling to destruction. Any movement that would result in the
  robot falling from the table must be prevented, however further valid
  movement commands must still be allowed.

Create an application that can read in commands of the following form:

    PLACE X,Y,F
    MOVE
    LEFT
    RIGHT
    REPORT

- PLACE will put the toy robot on the table in position X,Y and facing NORTH,
  SOUTH, EAST or WEST.
- The origin (0,0) can be considered to be the SOUTH WEST most corner.
- The first valid command to the robot is a PLACE command, after that, any
  sequence of commands may be issued, in any order, including another PLACE
  command. The application should discard all commands in the sequence until
  a valid PLACE command has been executed.
- MOVE will move the toy robot one unit forward in the direction it is
  currently facing.
- LEFT and RIGHT will rotate the robot 90 degrees in the specified direction
  without changing the position of the robot.
- REPORT will announce the X,Y and F of the robot. This can be in any form,
  but standard output is sufficient.

- A robot that is not on the table can choose the ignore the MOVE, LEFT, RIGHT
  and REPORT commands.
- Input can be from a file, or from standard input, as the developer chooses.
- Provide test data to exercise the application.
- The application must be a command line application.

Constraints
-----------

- The toy robot must not fall off the table during movement. This also
  includes the initial placement of the toy robot.
- Any move that would cause the robot to fall must be ignored.

Example Input and Output
------------------------

### Example a

    PLACE 0,0,NORTH
    MOVE
    REPORT

Expected output:

    0,1,NORTH

### Example b

    PLACE 0,0,NORTH
    LEFT
    REPORT

Expected output:

    0,0,WEST

### Example c

    PLACE 1,2,EAST
    MOVE
    MOVE
    LEFT
    MOVE
    REPORT

Expected output

    3,3,NORTH


Deliverables
------------

Please provide your source code, and any test code/data you using in
developing your solution.

Please engineer your solution to a standard you consider suitable for
production. It is not required to provide any graphical output showing the
movement of the toy robot.

Please do not put your name in any of the submitted code since this makes it
harder for us to review your submission anonymously.

<!--

Installation
============

-->

Usage
-----

To use the `toy-robot` application, either execute the script and start typing
in commands; e.g.

```shell
    $ bin/main

    REPORT
    "0,0,NORTH"
    PLACE 1,2,SOUTH
    REPORT
    "1,2,SOUTH"
    MOVE
    MOVE
    RIGHT
    REPORT
    "1,0,WEST"
    EXIT
```

Or execute the script with the path to file that will be used as input; e.g.

```shell
    $ bin/main README.md

    "0,1,NORTH"
    "0,1,NORTH"
    "0,0,WEST"
    "3,3,NORTH"
    "0,0,NORTH"
    "1,2,SOUTH"
    "1,0,WEST"

```

All given User inputs should match what was specified in the Toy Robot
specification; meaning its case-sensitive, ignores invalid commands, reports
back to the user when asked `REPORT`, etc. However, it also implements a `EXIT`
command in order escape the application/script outside of the regular terminate
command `Ctrl + C`.

Caveats
-------

- The square tabletop, of dimensions 5 units x 5 units, assumes a Range of
  `(0...5)` not `(0..5)`; e.g.

  ```ruby
      (0...5).to_a  # => [0, 1, 2, 3, 4]
      (0...5).count # => 5

      (0..5).to_a  # => [0, 1, 2, 3, 4, 5]
      (0..5).count # => 6
  ```

  Since counting starts at `0` from the origin, it means the fifth value would
  be `4` not `5` (which in this case would be the sixth value and create a
  tabletop of 6 x 6 units, with 36 possible positions, rather than 25).

<!--

## References
- [How to Write Go Code \- The Go Programming Language]
  (https://golang.org/doc/code.html)

- [library/golang \- Docker Hub]
  (https://hub.docker.com/_/golang/)

- [Setting up a Go Development Environment]
  (https://skife.org/golang/2013/03/24/go_dev_env.html)

- [Five suggestions for setting up a Go project \| Dave Cheney]
  (https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project)

- [Working with Multiple Files in a Go Project \| Motion Express \| Ruby, Rails, Crystal & developers' techniques]
  (http://www.motion-express.com/blog/organizing-a-go-project)

- [golang\-standards/project\-layout: Standard Go Project Layout]
  (https://github.com/golang-standards/project-layout)

- [Go Hello World \- golangbot\.com]
  (https://golangbot.com/hello-world/)

- [Example using Go]
  (http://docs.drone.io/golang-example/)

- [How to assert values in golang unit tests]
  (https://gist.github.com/samalba/6059502)

- [go \- Import local package \- Stack Overflow]
  (https://stackoverflow.com/questions/30017777/import-local-package)

- [In Go, how do I capture stdout of a function into a string? \- Stack Overflow]
  (https://stackoverflow.com/questions/10473800/in-go-how-do-i-capture-stdout-of-a-function-into-a-string)

- [Testing in Go \| via @codeship]
  (https://blog.codeship.com/testing-in-go/)

- [52\-technologies\-in\-2016/README\.md at master · shekhargulati/52\-technologies\-in\-2016]
  (https://github.com/shekhargulati/52-technologies-in-2016/blob/master/29-go-unit-testing/README.md)

- [Testing — An Introduction to Programming in Go \| Go Resources]
  (https://www.golang-book.com/books/intro/12)

- [Testing in Golang – Thejas Babu – Medium]
  (https://medium.com/@thejasbabu/testing-in-golang-c378b351002d)

- [5 simple tips and tricks for writing unit tests in \#golang]
  (https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742)

- [Interfaces and Composition for Effective Unit Testing in Golang \| I care, I share, I'm Nathan LeClaire\.]
  (https://nathanleclaire.com/blog/2015/10/10/interfaces-and-composition-for-effective-unit-testing-in-golang/)

- [Correct directory structure for a Go Project? \- Stack Overflow]
  (https://stackoverflow.com/questions/19100191/correct-directory-structure-for-a-go-project)

- [directory \- Listing only directories using ls in bash: An examination \- Stack Overflow]
  (https://stackoverflow.com/questions/14352290/listing-only-directories-using-ls-in-bash-an-examination)

- [Golang basics \- writing unit tests]
  (https://blog.alexellis.io/golang-writing-unit-tests/)

-->

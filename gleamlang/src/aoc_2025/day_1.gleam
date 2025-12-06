import gleam/int
import gleam/list
import gleam/result
import gleam/string

type Direction {
  Left
  Right
}

pub opaque type Rotation {
  Rotation(direction: Direction, distance: Int)
}

fn parse_direction(input: String) -> Result(Direction, Nil) {
  case input {
    "L" -> Ok(Left)
    "R" -> Ok(Right)
    _ -> Error(Nil)
  }
}

fn parse_line(line: String) -> Rotation {
  // parse direction
  let assert Ok(direction_str) = string.first(line)
    as "Line should not be empty!"
  let assert Ok(direction) = parse_direction(direction_str)
    as "First character of line should be 'L' or 'R'!"

  // parse distance value
  let distance_str = string.drop_start(line, up_to: 1)
  let assert Ok(distance) = int.parse(distance_str)
    as "Distance string should be an integer!"
  Rotation(direction:, distance:)
}

pub fn parse(input: String) -> List(Rotation) {
  input
  |> string.trim
  |> string.split(on: "\n")
  |> list.map(parse_line)
}

pub fn pt_1(input: List(Rotation)) {
  todo as "part 1 not implemented"
}

pub fn pt_2(input: List(Rotation)) {
  todo as "part 2 not implemented"
}

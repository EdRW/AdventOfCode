import gleam/int
import gleam/list
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

pub fn parse(input: String) -> List(Rotation) {
  let lines =
    input
    |> string.trim
    |> string.split(on: "\n")

  use rotation_str <- list.map(lines)

  let assert Ok(direction_str) = string.first(rotation_str)
  let assert Ok(direction) = parse_direction(direction_str)

  let distance_str = string.drop_start(rotation_str, up_to: 1)
  let assert Ok(distance) = int.parse(distance_str)
  Rotation(direction:, distance:)
}

pub fn pt_1(input: List(Rotation)) {
  input
}

pub fn pt_2(input: List(Rotation)) {
  todo as "part 2 not implemented"
}

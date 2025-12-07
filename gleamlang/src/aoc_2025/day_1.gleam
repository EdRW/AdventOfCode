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

type Dial {
  Dial(position: Int, zero_count: Int)
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

pub fn normalize_position(position: Int) -> Int {
  let num_ticks = 100
  let abs_position = int.absolute_value(position)

  let abs_normalized_position = case abs_position >= num_ticks {
    True -> abs_position % num_ticks
    False -> abs_position
  }

  case position < 0 {
    True if abs_normalized_position != 0 -> num_ticks - abs_normalized_position
    _ -> abs_normalized_position
  }
}

fn turn_dial(dial: Dial, rotation: Rotation) -> Dial {
  let sign = case rotation.direction {
    Left -> -1
    Right -> 1
  }

  let new_raw_position = dial.position + sign * rotation.distance
  let position = normalize_position(new_raw_position)
  let zero_count =
    dial.zero_count
    + case position {
      0 -> 1
      _ -> 0
    }
  Dial(position:, zero_count:)
}

pub fn pt_1(input: List(Rotation)) {
  let dial =
    input
    |> list.fold(Dial(position: 50, zero_count: 0), turn_dial)

  dial.zero_count
}

fn turn_dial_2(dial: Dial, rotation: Rotation) -> Dial {
  let sign = case rotation.direction {
    Left -> -1
    Right -> 1
  }
  let distance_to_zero = case rotation.direction {
    Left -> dial.position
    Right -> 100 - dial.position
  }

  let new_raw_position = dial.position + sign * rotation.distance
  let position = normalize_position(new_raw_position)

  let num_zero_pass =
    rotation.distance
    / 100
    + case rotation.distance % 100 >= distance_to_zero {
      True if distance_to_zero != 0 -> 1
      _ -> 0
    }

  let zero_count = dial.zero_count + num_zero_pass
  Dial(position:, zero_count:)
}

pub fn pt_2(input: List(Rotation)) {
  let dial =
    input
    |> list.fold(Dial(position: 50, zero_count: 0), turn_dial_2)

  dial.zero_count
}

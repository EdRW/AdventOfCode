import gleam/int
import gleam/list
import gleam/string

pub opaque type Rotation {
  CounterClockwise(distance: Int)
  Clockwise(distance: Int)
}

type Dial {
  Dial(position: Int, zero_count: Int)
}

// --------------------------------------------------------------------------
//                                   Parser                                  
// --------------------------------------------------------------------------

pub fn parse(input: String) -> List(Rotation) {
  input
  |> string.trim
  |> string.split(on: "\n")
  |> list.map(parse_line)
}

fn parse_line(line: String) -> Rotation {
  // parse direction string
  let assert Ok(direction_str) = string.first(line)
    as "Line should not be empty!"

  // parse a partial Rotation with direction but no distance value
  let assert Ok(rotation) = parse_rotation(direction_str)
    as "First character of line should be 'L' or 'R'!"

  // parse distance value
  let distance_str = string.drop_start(line, up_to: 1)
  let assert Ok(distance) = int.parse(distance_str)
    as "Distance string should be an integer!"

  rotation(distance)
}

fn parse_rotation(l_or_r: String) {
  case l_or_r {
    "L" -> Ok(CounterClockwise)
    "R" -> Ok(Clockwise)
    _ -> Error(Nil)
  }
}

// --------------------------------------------------------------------------
//                                   Part 1                                  
// --------------------------------------------------------------------------

pub fn pt_1(input: List(Rotation)) {
  let dial =
    input
    |> list.fold(Dial(position: 50, zero_count: 0), turn_dial)

  dial.zero_count
}

fn turn_dial(dial: Dial, rotation: Rotation) -> Dial {
  let displacement = case rotation {
    CounterClockwise(distance) -> -1 * distance
    Clockwise(distance) -> distance
  }

  let new_raw_position = dial.position + displacement
  let position = normalize_position(new_raw_position)
  let zero_count =
    dial.zero_count
    + case position {
      0 -> 1
      _ -> 0
    }
  Dial(position:, zero_count:)
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

// --------------------------------------------------------------------------
//                                   Part 2                                  
// --------------------------------------------------------------------------

pub fn pt_2(input: List(Rotation)) {
  let dial =
    input
    |> list.fold(Dial(position: 50, zero_count: 0), turn_dial_2)

  dial.zero_count
}

fn turn_dial_2(dial: Dial, rotation: Rotation) -> Dial {
  let displacement = case rotation {
    CounterClockwise(distance) -> -1 * distance
    Clockwise(distance) -> distance
  }

  let new_raw_position = dial.position + displacement
  let position = normalize_position(new_raw_position)

  let zero_count = dial.zero_count + times_pass_zero(rotation, dial)
  Dial(position:, zero_count:)
}

fn times_pass_zero(rotation: Rotation, dial: Dial) -> Int {
  let num_ticks = 100
  let distance_to_zero = case rotation {
    CounterClockwise(_) -> dial.position
    Clockwise(_) -> num_ticks - dial.position
  }

  // add 1 extra time passed zero
  let is_passed_zero =
    distance_to_zero > 0 && rotation.distance % num_ticks >= distance_to_zero
  let maybe_extra_time = case is_passed_zero {
    False -> 0
    True -> 1
  }

  // we'll definitely pass zero with each full rotation
  let num_full_rotations = rotation.distance / num_ticks
  num_full_rotations + maybe_extra_time
}

import gleam/int
import gleam/list
import gleam/string

pub type IdRange =
  #(Int, Int)

// --------------------------------------------------------------------------
//                                   Parser                                  
// --------------------------------------------------------------------------

pub fn parse(input: String) -> List(IdRange) {
  input
  |> string.trim
  |> string.split(on: ",")
  |> list.map(parse_line)
}

fn parse_line(line: String) -> IdRange {
  let split_result = case string.split(line, on: "-") {
    [first_id_str, last_id_str, ..] -> Ok(#(first_id_str, last_id_str))
    _ -> Error(Nil)
  }

  let assert Ok(#(first_id_str, last_id_str)) = split_result
    as "Line should be values before and after a '-'"
  let assert Ok(first_id) = int.parse(first_id_str)
    as "First section should be an Int"
  let assert Ok(last_id) = int.parse(last_id_str)
    as "Second section should be an Int"

  #(first_id, last_id)
}

// --------------------------------------------------------------------------
//                                   Part 1                                  
// --------------------------------------------------------------------------

pub fn pt_1(input: List(IdRange)) {
  echo input
  todo as "part 1 not implemented"
}

// --------------------------------------------------------------------------
//                                   Part 2                                  
// --------------------------------------------------------------------------

pub fn pt_2(input: List(IdRange)) {
  todo as "part 2 not implemented"
}

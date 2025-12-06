import aoc_2025/day_1
import gleam/int
import gleam/list
import gleam/string
import gleeunit

pub fn main() -> Nil {
  gleeunit.main()
}

// gleeunit test functions end in `_test`
pub fn hello_world_test() {
  let name = "Joe"
  let greeting = "Hello, " <> name <> "!"

  assert greeting == "Hello, Joe!"
}

pub fn normalize_distance_test() {
  let cases = [
    #(0, 0),
    #(100, 0),
    #(100 * 2, 0),
  ]
  cases
  |> list.each(fn(test_case) {
    let #(input, result) = test_case
    assert day_1.normalize_position(input) == result
      as string.concat(["Input: ", int.to_string(input)])
    assert day_1.normalize_position(-1 * input) == result
      as string.concat(["Input: ", int.to_string(-1 * input)])
  })

  let cases = [
    #(1, 1),
    #(98, 98),
    #(99, 99),
    #(101, 1),
  ]

  cases
  |> list.each(fn(test_case) {
    let #(input, result) = test_case
    assert day_1.normalize_position(input) == result
      as string.concat(["Input: ", int.to_string(input)])
    assert day_1.normalize_position(-1 * input) == 100 - result
      as string.concat(["Input: ", int.to_string(-1 * input)])
  })
}

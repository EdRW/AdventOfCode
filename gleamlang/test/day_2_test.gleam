import aoc_2025/day_2
import gleam/int
import gleam/list
import gleam/string
import gleeunit

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn sum_invalids_in_i_test() {
  assert day_2.sum_invalids_for_i(1) == 495
  assert day_2.sum_invalids_for_i(2) == 495_900
  assert day_2.sum_invalids_for_i(3) == 495_540_450
}

pub fn calc_num_digits_test() {
  assert day_2.calc_num_digits(6) == 1
  assert day_2.calc_num_digits(25) == 2
  assert day_2.calc_num_digits(435_354) == 6
}

pub fn sum_invalid_ids_till_test() {
  assert day_2.sum_invalid_ids_till(11) == 11
  assert day_2.sum_invalid_ids_till(22) == 33
  assert day_2.sum_invalid_ids_till(99) == 495
  assert day_2.sum_invalid_ids_till(1111) == 2616
}

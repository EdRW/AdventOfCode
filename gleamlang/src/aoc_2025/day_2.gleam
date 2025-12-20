import gleam/float
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
  input
  |> list.map(sum_invalid_ids_in_range)
  |> list.fold(from: 0, with: fn(acc, invalid_id) { acc + invalid_id })
}

fn sum_invalid_ids_in_range(id_range: IdRange) -> Int {
  // an id is invalid if:
  //   when the digits of the number are treated as an array,
  //   and the array is split in half, 
  //   the values of each index of the arrays are the same
  //   eg. 123123 -> 123 , 123
  //   odd-digit ids are always valid
  //   cutting a number in half, and dividing the entire number by half the number,
  //   will result in a number of the form "1 n zeros 1"
  // 
  // But we also need to find candidate invalid ids in the id ranges
  // Algo
  // if the min value is greater than max value
  //   it is not an id
  // if the number of digits in the min and max are different
  // split the min number array in half
  //   find the smallest index where the two halves could match
  //   eg: 123123 - 123125 -> 12
  //   does the next index exist within the range?
  //     if not then we can't find an invalid id
  //   does the next index complete a valid id match? (bc it's the final index of the half)
  //   does the next index value match for both halves?
  //     ok let's increment the next digit (next index) and use as new min value
  // 
  // 11     : 11,22,33, ... 99
  // 101    : 1010 ... 9999
  // 1001   : 100100 ... 999999
  // 100001 : 10001000 ...99999999
  let first_sum = sum_invalid_ids_till(id_range.0 - 1)
  let second_sum = sum_invalid_ids_till(id_range.1)
  second_sum - first_sum
}

pub fn sum_invalid_ids_till(id: Int) -> Int {
  let id_float = int.to_float(id)
  let num_digits = calc_num_digits(id)
  let i = num_digits / 2
  case num_digits % 2 {
    0 -> {
      let i_float = int.to_float(i)
      let assert Ok(ten_to_i) = float.power(10.0, i_float)
      let assert Ok(ten_to_i_one) = float.power(10.0, i_float -. 1.0)
      let n_val = float.floor(id_float /. { ten_to_i +. 1.0 })
      let current_term =
        { ten_to_i +. 1.0 }
        /. 2.0
        *. {
          n_val *. { n_val +. 1.0 } -. ten_to_i_one *. { ten_to_i_one -. 1.0 }
        }
      float.round(current_term) + sum_invalids_for_i(i - 1)
    }
    _ -> sum_invalids_for_i(i)
  }
}

pub fn sum_invalids_for_i(i: Int) -> Int {
  case i < 1 {
    True -> 0
    False -> {
      let i_float = int.to_float(i)
      let assert Ok(term_1) = float.power(10.0, 3.0 *. i_float)
      let assert Ok(term_2) = float.power(10.0, 2.0 *. i_float)
      let assert Ok(term_3) = float.power(10.0, i_float)
      let result = 0.5 *. { 0.99 *. term_1 +. 0.09 *. term_2 -. 0.9 *. term_3 }
      float.round(result) + sum_invalids_for_i(i - 1)
    }
  }
}

pub fn calc_num_digits(num: Int) -> Int {
  let num_float = int.to_float(num)
  let assert Ok(ln_num) = float.logarithm(num_float)
  let assert Ok(ln_10) = float.logarithm(10.0)
  let log_num = ln_num /. ln_10
  float.round(float.floor(log_num)) + 1
}

// --------------------------------------------------------------------------
//                                   Part 2                                  
// --------------------------------------------------------------------------

pub fn pt_2(input: List(IdRange)) {
  todo as "part 2 not implemented"
}

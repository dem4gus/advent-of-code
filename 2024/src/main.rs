use std::fs;
use advent_of_code_2024::day02::count_safe;

fn main() {
    let day_one_path = "inputs/day02.txt";
    let part_one = match fs::read_to_string(day_one_path) {
        Ok(input) => count_safe(input),
        Err(error) => panic!("There was a problem calculating the answer: {}", error),
    };

    println!("Day 2: {}", part_one);
}

use advent_of_code_2024::day01::{calculate_distance, calculate_similarity};
use std::fs;

fn main() {
    let day_one_path = "inputs/day01.txt";
    let (part_one, part_two) = match fs::read_to_string(day_one_path) {
        Ok(input) => (
            calculate_distance(input.clone()),
            calculate_similarity(input),
        ),
        Err(error) => panic!("There was a problem calculating the answer: {}", error),
    };

    println!("The answers are {}, {}", part_one, part_two);
}

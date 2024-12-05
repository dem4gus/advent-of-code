use std::collections::HashMap;
use std::iter::zip;

pub fn calculate_distance(input: String) -> u64 {
    let mut first: Vec<i64> = vec![];
    let mut second: Vec<i64> = vec![];

    for line in input.lines() {
        let nums: Vec<_> = line
            .split_whitespace()
            .map(|n| n.parse().unwrap())
            .collect();
        first.push(nums[0]);
        second.push(nums[1]);
    }

    first.sort();
    second.sort();

    zip(first, second).map(|(x, y)| (x - y).abs() as u64).sum()
}

pub fn calculate_similarity(input: String) -> u64 {
    let mut first: Vec<String> = vec![];
    let mut second: HashMap<String, u64> = HashMap::new();

    for line in input.lines() {
        let nums: Vec<_> = line.split_whitespace().collect();
        first.push(nums[0].to_string());
        let key = nums[1].to_string();
        if let Some(val) = second.insert(key.clone(), 1) {
            second.insert(key, val + 1);
        };
    }

    first
        .iter()
        .map(|x| {
            return if let Some(&val) = second.get(x) {
                x.parse::<u64>().unwrap() * val
            } else {
                0u64
            };
        })
        .sum()
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "3   4
4   3
2   5
1   3
3   9
3   3";

    #[test]
    fn difference_passes() {
        let want = 11;
        let result = calculate_distance(INPUT.to_string());

        assert_eq!(result, want)
    }

    #[test]
    fn similarity_passes() {
        let want = 31;
        let result = calculate_similarity(INPUT.to_string());

        assert_eq!(result, want)
    }
}

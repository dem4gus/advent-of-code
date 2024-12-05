const MAX_STEP_SIZE: i64 = 3;

pub fn count_safe(reports: String) -> u64 {
    reports.lines().map(process_report).filter(|is_safe| *is_safe).count() as u64
}

fn process_report(report: &str) -> bool {
    let report = report
        .split_whitespace()
        .map(|n| n.parse::<i64>().unwrap())
        .collect::<Vec<i64>>();
    let mut rev = report.clone();
    rev.reverse();

    if !(report.is_sorted() || rev.is_sorted()) {
        return false;
    }

    let mut safe = true;
    report.iter().reduce(|a, b| {
        let step = (a-b).abs();
        safe = safe && step != 0 && step <= MAX_STEP_SIZE;
        b
    });

    safe
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9";
    #[test]
    fn get_safe_passes() {
        let want = 2;
        let result = count_safe(INPUT.to_string());
        assert_eq!(result, want)
    }
}

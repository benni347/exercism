pub fn is_leap_year(year: u64) -> bool {
    return if year % 4 == 0 && (year % 100 != 0 || year % 400 == 0) {
        true
    } else {
        false
    };
}

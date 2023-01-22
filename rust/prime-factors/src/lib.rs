pub fn factors(n: u64) -> Vec<u64> {
    let mut factors = vec![];
    let mut n = n;
    if n < 2 {
        return factors;
    }
    if n == 2 {
        factors.push(n);
        return factors;
    }
    for i in 2..=(n as f64).sqrt() as u64 {
        while n % i == 0 {
            factors.push(i);
            n /= i;
        }
    }
    if n > 2 {
        factors.push(n);
    }
    factors
}

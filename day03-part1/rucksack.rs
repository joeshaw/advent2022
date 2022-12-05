use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let f = File::open("input.txt").expect("file not found");
    let reader = BufReader::new(f);

    let mut sum = 0;
    for line in reader.lines() {
        let line = line.expect("unable to read line");
        let first_compartment = &line[..line.len()/2];
        let second_compartment = &line[line.len()/2..];

        for ch in first_compartment.chars() {
            if second_compartment.contains(ch) {
                if ch >= 'a' && ch <= 'z' {
                    sum += (ch as u8 - b'a' + 1) as i32;
                } else {
                    sum += (ch as u8 - b'A' + 27) as i32;
                }
                break;
            }
        }
    }

    println!("{}", sum);
}

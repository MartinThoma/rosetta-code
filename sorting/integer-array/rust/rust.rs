extern crate rand;
extern crate time;

use rand::Rng;
use time::PreciseTime;

fn main() {
    // Creates an array of 10000000 random integers in the range 0 - 1000000000
    //let mut array: [i32; 10000000] = [0; 10000000];
    let n = 10000000;
    let mut array = Vec::new();

    // Fill the array
    let mut rng = rand::thread_rng();
    for _ in 0..n {
        //array[i] = rng.gen::<i32>();
        array.push(rng.gen::<i32>());
    }

    // Sort
    let start = PreciseTime::now();
    array.sort();
    let end = PreciseTime::now();

    println!("{} seconds for sorting {} integers.", start.to(end), n);
}
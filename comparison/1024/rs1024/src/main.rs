// copy of https://zhuanlan.zhihu.com/p/30397623
extern crate rand;

use std::{thread, time};
use rand::Rng;

fn main() {
    let mut l_cnt = 0;
    let start = time::Instant::now();

    loop {
        let mut s1024 = String::new();
        let mut spawns = Vec::new();
        let mut cnt = 0;

        while cnt < 4 {
            spawns.push(thread::spawn(|| {
                let v = vec![1, 0, 2, 4];
                thread::sleep(time::Duration::from_millis(
                    rand::thread_rng().gen_range(10, 101)));
                v[rand::thread_rng().gen_range(0, 4)]
            }));

            cnt += 1;
        }

        for s in spawns {
            s1024 += &s.join().expect("Crack!").to_string();
        }

        l_cnt += 1;

        if s1024 == "1024" {
            println!("经历了 {} 次随机生成，花费 {} 秒。10.24 程序员节快乐！(*^▽^*)",
                l_cnt,
                time::Instant::now().duration_since(start).as_secs());

            break;
        }
    }
}

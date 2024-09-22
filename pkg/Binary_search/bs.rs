impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> i32 {
        let mut l = 0;
        let mut r = nums.len() as i32 -1;

        while l <= r {
            let mid = l + (r-l) /2;
            if nums[mid as usize] == target {
                return mid;
            } else if nums[mid as usize] < target {
                l = mid + 1;
            } else {
                r = mid - 1;
            }
        }

        return -1;
    }
}

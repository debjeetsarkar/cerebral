/**
You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check. Since each version is developed based on the previous version, all the versions after a bad version are also bad.

Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.

You are given an API bool isBadVersion(version) which will return whether version is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.

Example:

Given n = 5, and version = 4 is the first bad version.

call isBadVersion(3) -> false
call isBadVersion(5) -> true
call isBadVersion(4) -> true

Then 4 is the first bad version. 


2
2

low = 1
high = 2
mid = 3/2 = 1

low = 2
high = 2
mid = 4/2 = 2
**/

func firstBadVersion(n int) int {
    if n == 0 {
        return 0
    }

    low := 1
    high := n
    mid := low + (high - low) / 2

    for low < high {
        mid = low + (high - low) / 2
        if isBadVersion(mid) {
            high = mid
        } else {
            low = mid + 1
        }
    }

    return low
}
/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    l, r := 1, n
	for l < r {
		mid := (l + r) / 2
		if guess(mid) < 0 { // actual number is lower 
			r = mid - 1
		} else if guess(mid) > 0 { // actual number is higher
			l = mid + 1
		} else { //actual number 
			return mid
		}
	}
	return l
}

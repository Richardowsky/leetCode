package palindromeNumber

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	n := x
	big := 10
	small := 1
	for n != 0 {
		n = n / 10
		back := x % big / small
		front := x
		for front > big-1 {
			front = front / 10
		}
		front = front % 10
		if front != back {
			return false
		}
		big = big * 10
		small = small * 10
	}

	return true
}

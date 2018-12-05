// 1. The order of reaction does not matters
// 2. Can brute force(scan and scan), which takes O(n^2)
package main

import (
	"fmt"
	"os"
)

var uAAsc = "A"[0]
var uZAsc = "Z"[0]
var lAAsc = "a"[0]

// Can also handle the problem without format, just prefer this style.
func format(src string) []int {
	dst := make([]int, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] >= uAAsc && src[i] <= uZAsc {
			dst[i] = int(src[i]) - int(uAAsc)
		} else {
			dst[i] = -(int(src[i]) - int(lAAsc))
		}
	}
	return dst
}

func firstChallenge() {
	s := ""
	fmt.Scan(&s)

	arr := format(s)
	gone := make([]bool, len(arr))

	for i := 0; i < len(arr)/2; i++ {
		cur, continueLoop := -1, false
		for j := 0; j < len(arr); j++ {
			if gone[j] {
				continue
			}
			if cur == -1 {
				cur = j
				continue
			}
			if (arr[cur] == -arr[j]) && (s[cur] != s[j]) {
				gone[cur], gone[j], cur, continueLoop = true, true, -1, true
			} else {
				cur = j
			}
		}
		if !continueLoop {
			break
		}
	}
	sum := len(arr)
	for i := 0; i < len(arr); i++ {
		if gone[i] {
			sum--
		}
	}
	fmt.Println(sum)
}

func secondChallenge() {
	s := ""
	fmt.Scan(&s)

	arr := format(s)
	best := 50000
	for k := 0; k < 26; k++ {
		gone := make([]bool, len(arr))
		for i := 0; i < len(arr); i++ {
			if s[i] == (uAAsc+byte(k)) || s[i] == (lAAsc+byte(k)) {
				gone[i] = true
			}
		}

		for i := 0; i < (len(arr) / 2); i++ {
			cur, continueLoop := -1, false

			for j := 0; j < len(arr); j++ {
				if gone[j] {
					continue
				}
				if cur == -1 {
					cur = j
					continue
				}
				if (arr[cur] == -arr[j]) && (s[cur] != s[j]) {
					gone[cur], gone[j], cur, continueLoop = true, true, -1, true
				} else {
					cur = j
				}
			}

			if !continueLoop {
				break
			}
		}
		sum := len(arr)
		for i := 0; i < len(arr); i++ {
			if gone[i] {
				sum--
			}
		}
		if sum < best {
			best = sum
		}
	}
	fmt.Println(best)
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "2" {
		secondChallenge()
	} else {
		firstChallenge()
	}
}

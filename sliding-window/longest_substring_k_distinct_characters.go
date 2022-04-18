package main

/*
Given a string s and an integer k, return the length of the longest substring of s
	that contains at most k distinct characters.

MEDIUM
https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/
*/

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	var best int = 0
	var left int = 0

	// note there is no char type in go, instead there are runes, which are
	//	32 bit integers that represent unicode
	m := make(map[byte]int)

	for i := range s {
		m[s[i]] = i

		for {
			if len(m) > k {
				if m[s[left]] <= 1 {
					delete(m, s[left])
				} else {
					m[s[left]] = m[s[left]] - 1
				}

				left = left + 1
			} else {
				break
			}
		}

		if (i - left + 1) > best {
			best = i - left + 1
		}
	}
	return best
}

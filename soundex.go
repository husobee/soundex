// Package soundex - go soundex implementation
package soundex

import (
	"container/list"
	"errors"
	"strings"
)

//SoundexSubMap - Substitution map for soundex
/*

   b, f, p, v → 1
   c, g, j, k, q, s, x, z → 2
   d, t → 3
   l → 4
   m, n → 5
   r → 6

*/
var (
	//SoundexSubMap - runes to replace with numbers
	SoundexSubMap = map[rune]int{
		'b': 1, 'f': 1, 'p': 1, 'v': 1,
		'B': 1, 'F': 1, 'P': 1, 'V': 1,
		'c': 2, 'g': 2, 'j': 2, 'k': 2, 'q': 2, 's': 2, 'x': 2, 'z': 2,
		'C': 2, 'G': 2, 'J': 2, 'K': 2, 'Q': 2, 'S': 2, 'X': 2, 'Z': 2,
		'd': 3, 't': 3,
		'D': 3, 'T': 3,
		'l': 4,
		'L': 4,
		'm': 5, 'n': 5,
		'M': 5, 'N': 5,
		'r': 6,
		'R': 6,
	}
	//DropSubMap - runes to drop
	DropSubMap = map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'y': true, 'h': true, 'w': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true, 'Y': true, 'H': true, 'W': true,
	}
	IntToRuneMap = map[int]rune{
		1: '1', 2: '2', 3: '3', 4: '4', 5: '5', 6: '6',
	}
)

var (
	// ErrInvalidRune - a common error for invalid rune
	ErrInvalidRune = errors.New("Invalid rune, no soundex code found")
	// ErrFailedSoundex - a common error for invalid string
	ErrFailedSoundex = errors.New("Invalid string, can't make a soundex from it")
)

// RuneSoundex - get the soundex for a given rune
func RuneSoundex(r rune) (rune, error) {
	if v, ok := SoundexSubMap[r]; ok {
		return IntToRuneMap[v], nil
	}
	return rune('0'), ErrInvalidRune
}

//Soundex - function to get the soundex from a string
func Soundex(s string) (string, error) {
	var result []rune
	sList := list.New()
	rs := []rune(s)
	// turn rune slice into a list, for easier processing
	for _, r := range rs {
		sList.PushBack(r)
	}

	// deduplication
	for e := sList.Front(); e != nil; e = e.Next() {
		value, _ := e.Value.(rune)
		if e.Next() != nil {
			nextValue, _ := e.Next().Value.(rune)
			if _, ok := map[rune]bool{'h': true, 'w': true}[nextValue]; ok && e.Next().Next() != nil {
				//next codepoint is an h or w, check the next value
				nextNextValue, _ := e.Next().Next().Value.(rune)
				if dup := DuplicateRune(value, nextNextValue); dup {
					sList.Remove(e.Next().Next())
				}
			} else {
				if dup := DuplicateRune(value, nextValue); dup {
					sList.Remove(e.Next())
				}
			}
		}
		if len(result) > 3 {
			break
		}
		if len(result) > 0 {
			// apply substitution map
			if _, ok := DropSubMap[value]; !ok {
				soundexValue, _ := RuneSoundex(value)
				result = append(result, soundexValue)
			}
		} else {
			// put the first codepoint unmolested
			result = append(result, value)
		}
	}
	for i := len(result); i < 4; i++ {
		result = append(result, '0')
	}
	return strings.ToUpper(string(result)), nil
}

//DuplicateRune - remove duplicates in string adjacent
func DuplicateRune(r1, r2 rune) bool {
	rs1, _ := RuneSoundex(r1)
	rs2, _ := RuneSoundex(r2)
	return rs1 == rs2
}

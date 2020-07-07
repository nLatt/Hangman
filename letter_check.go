package main

import (
  "fmt"
  s "strings"
)

// checks for the occurence of the given letter in the word to find
func letter_check(word_to_find, letter string) []int {
  // array with length of the occurence of the letter to save the indexes
  indexes := make([]int, s.Count(word_to_find, letter))
  // get all indexes of the occurences of "letter"
  placeholder := 0
  for i := 0; i < len(indexes); i++ {
    if i > 0 {
      placeholder = indexes[i-1] + 1
      indexes[i] = s.Index(word_to_find[placeholder:], letter) + placeholder
    } else {
      indexes[i] = s.Index(word_to_find[placeholder:], letter)
    }
  }
  // result message
  if len(indexes) == 0 {
    fmt.Printf("No occurence of %s in word.\n", letter)
  } else {
    fmt.Printf("All indexes of %s in the word: %v\n", letter, indexes)
  }
  return indexes
}

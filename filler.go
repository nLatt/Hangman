package main

import ()

// place matching letters in the empty array
func filler(indexes []int, letter string, guess_results []string) []string {
  for i := range indexes {
    guess_results[indexes[i]] = letter
  }
  return guess_results
}

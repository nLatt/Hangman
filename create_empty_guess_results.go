package main

import ()

// create the string in which the guessed numbers will be stocked at their respective index
func create_empty_guess_results(word_to_find string) []string {
  empty_results := make([]string, len(word_to_find))
  for i := range empty_results {
    empty_results[i] = "_"
  }
  return empty_results
}

package main

import (
  "fmt"
  "os"
  rd "encoding/csv"
  b "bufio"
  "math/rand"
  s "strings"
  "time"
)

// returns a word of the chosen difficulty and the amount of guesses
func word_chooser() (string, int) {
  short, medium, long := word_grouper()
  rand.Seed(time.Now().UnixNano())

  for {
    fmt.Println("Please enter one of the following: short, medium, long")
    fmt.Print("-> ")
    scanner := b.NewScanner(os.Stdin)
    scanner.Scan()
    difficulty := s.ToLower(scanner.Text())
    switch difficulty{
    case "short":
      rand_int := rand.Intn(len(short))
      return short[rand_int], 8
    case "medium":
      rand_int := rand.Intn(len(medium))
      return medium[rand_int], 12
    case "long":
      rand_int := rand.Intn(len(long))
      return long[rand_int], 16
    }
  }
}

// groups words bz length into short, medium, and long
func word_grouper() ([]string, []string, []string) {

  file, err := os.Open("words_csv/dico.csv")
  error_catch(err)

  var short, medium, long []string

  reader := rd.NewReader(file)
  words, err := reader.ReadAll()
  error_catch(err)

  for i := 0; i < len(words); i++ {
    length := len(words[i][0])
    switch {
      case length <= 5:
        short = append(short, words[i][0])
      case length < 8 && length > 5:
        medium = append(medium, words[i][0])
      case length >= 8:
        long = append(long, words[i][0])
    }
  }

  return short, medium, long
}

func error_catch(err error) {
  if err != nil {
    panic(err)
    fmt.Println("An error has occured: ", err)
  }
}

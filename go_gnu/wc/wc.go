package main

import (
  "bufio"
  "flag"
  "fmt"
  "io"
  "os"
  "strings"
)

type fileInfo struct {
  name string
  lines int
  words int
  chars int
  bytes int
}

type flagInfo struct {
  line bool
  word bool
  char bool
  byte bool
}

func newFileInfo() fileInfo {
  f := fileInfo{}
  f.name = "-"
  f.lines = 0
  f.words = 0
  f.chars = 0
  f.bytes = 0
  return f
}

func isInputFromPipe() bool {
  fileInfo, _ := os.Stdin.Stat()
  return fileInfo.Mode() & os.ModeCharDevice == 0
}

func parseFile(r io.Reader, f fileInfo) fileInfo {
  scanner := bufio.NewScanner(bufio.NewReader(r))
  for scanner.Scan() {
    ln := scanner.Text()
    f.lines += 1
    f.words += len(strings.Fields(ln))
    f.chars += len([]rune(ln)) + 2
    f.bytes += len(scanner.Bytes()) + 2
  }
  return f
}

func output(i fileInfo) {
  s := strings.Repeat(" ", 3)
  fmt.Print(s)
  if !*lineFlag && !*wordFlag && !*charFlag && !*byteFlag {
    *lineFlag = true
    *wordFlag = true
    *byteFlag = true
  }
  if *lineFlag {
    fmt.Print(i.lines, s)
  }
  if *wordFlag {
    fmt.Print(i.words, s)
  }
  if *charFlag {
    fmt.Print(i.chars, s)
  }
  if *byteFlag {
    fmt.Print(i.bytes, s)
  }
  if i.name != "-" {
    fmt.Print(i.name)
  }
  fmt.Print("\n")
}

var (
  lineFlag *bool
  wordFlag *bool
  charFlag *bool
  byteFlag *bool
)

func init() {
  lineFlag = flag.Bool("l", false, "print the newline counts")
  wordFlag = flag.Bool("w", false, "print the word counts")
  charFlag = flag.Bool("m", false, "print the character counts")
  byteFlag = flag.Bool("c", false, "print the byte counts")
}

func main() {
  flag.Parse()
 
  var o fileInfo

  if isInputFromPipe() {
    o = parseFile(os.Stdin, newFileInfo())
    output(o)
  } else {
    for _, f := range(flag.Args()) {
        o = newFileInfo()
        o.name = f
        f, err := os.Open(f)
        if err != nil {
          panic(err)
        } else {
          o = parseFile(f, o)
          output(o)
        }
      }
    }
}

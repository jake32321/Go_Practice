/*
  Jacob Reed 2015

  This is meant solely to be a demonstration and learning purposes only.
  The code included is just an example of what can be done with GoLang.
  Please use with respect and courtesy.
*/

package main

//Required items for this program
import (
  "bytes"
  "fmt"
  "math/rand"
  "time"
)

//Creates a two dirmensional field
//Uses struct similar to C
type Field struct{
  //Data types declared after????
  s [][]bool
  w, h int
}

//This function will make a new empty field based of the determined width and height.
func NewField(w, h int) *Field {
    //Makes a two dimensional array for the field
    s := make([][]bool, h)
    for i := range s{
      //Sets the width of the created field
      s[i] = make([]bool, w)
    }
    return &Field{s: s, w: w, h: h}
}

//Sets variable f to a boolean value at a specific (x, y)
func (f *Field) Set(x, y int, b bool)  {
  f.s[y][x] = b
}

//Wrapper for those coordinates that end up outside the field boundaries
func (f *Field) Alive(x, y int) bool{
  //Wrapper for the width
  x += f.w
  x %= f.w
  //Wrapper for the height
  y += f.h
  y %= f.h
  return f.s[y][x]
}

//Gives state of a cell on the next step
func (f *Field) Next(x, y int) bool{
  //Counts cells surrounding it that are still alive on x and y axis
  alive := 0
  for i := -1; i <= 1; i++{
    for j := -1; j <= 1 j++{
      if(j != 0 || i != 0 && f.Alive(x+i, y+j)){
        alive++
      }
    }
  }
}

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

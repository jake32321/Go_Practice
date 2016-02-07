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
type Field struct {
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
func (f *Field) Set(x, y int, b bool) {
  f.s[y][x] = b
}

//Wrapper for those coordinates that end up outside the field boundaries
func (f *Field) Alive(x, y int) bool {
  //Wrapper for the width
  x += f.w
  x %= f.w
  //Wrapper for the height
  y += f.h
  y %= f.h
  return f.s[y][x]
}

//Gives state of a cell on the next step
func (f *Field) Next(x, y int) bool {
  //Counts cells surrounding it that are still alive on x and y axis
  alive := 0
  for i := -1; i <= 1; i++{
    for j := -1; j <= 1 j++{
      if(j != 0 || i != 0 && f.Alive(x+i, y+j)){
        alive++
      }
    }
  }
  //Returns the state based on the rules of the game.
  //State might be off as well.
  return alive == 3 || alive == 2 && f.Alive(x, y)
}

/*
Next bit of code will be the meat and bones behind the actual game and
its current state when it is running.  The structs and functions above
will be incorporated into the next bit.
*/

//Actual game state
type Life struct{
  a, b *Field
  w, h int
}

//Function starts a new game
func NewLife(w, h int) *Life{
  //Creates a new field for the game
  a := NewField(w, h)
  for i := 0; i < (w * h / 4); i++ {
    a.Set(rand.Intn(w), rand.Intn(h), true)
  }
  //Creates the new field from the randomly generated variables
  return &Life{
    a: a,
    b: NewField(w, h),
    w: w,
    h: h,
  }
}

//Function progresses the game one step, updating all the cells
func (l *Life) Step() {
  //Moves from state of the current cells to updated ones
  for y := 0; y < l.h; y++ {
    for x := 0; x < l.w; x++ {
      //Sets the updated field to be the next itteration of the old field.
      l.b.Set(x, y, l.a.Next(x, y))
    }
  }
  //Swap old with the new
  l.a, l.b = l.b, l.a
}

//Function visualizes the game using Strings
function (l *Life) String() string {
  var buf bytes.Buffer
  for y := 0; y <
}

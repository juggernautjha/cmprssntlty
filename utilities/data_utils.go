/*
For doing `things` with Data.
Author: Jha
The good thing about Go is that functions are under namespaces (package in this case)
and I can therefore keep splitting the code till the second coming of christ with little to
no consequence.
! Boy testing this is going to be fun.
*/
package utilities

import (
	"errors"
)

/*
Datablock -> wrapper around a list of symbols.
*/
type Datablock struct {
	Data []rune
}

func Get_Datablock(t []rune) Datablock {
	/*
		Utility function, converts string to a Datablock.
	*/
	return Datablock{t}
}

func Get_data(t *Datablock) []rune {
	return t.Data
}

func (t *Datablock) Get_and_count() ([]rune, map[rune]int64) {
	/*
		Gets a frequency count. Returns a list of symbols and a map of frequencies.
	*/
	var temp_map map[rune]int64 = make(map[rune]int64)
	for _, v := range t.Data {
		temp_map[v] += 1
	}
	var temp_list []rune
	for i := range temp_map {
		temp_list = append(temp_list, i)
	}
	return temp_list, temp_map
}

func (t *Datablock) Get_probability() Distribution {
	/*
		Like Arpit Bala said, dress well have sex.
	*/
	_, my_map := t.Get_and_count()
	return Get_fromfrequency(my_map)
}

/*
Blocks are boring. Real men use streams. Get ready for the worst use of interfaces you have ever seen.
Do not come at me, I am probably smarter than your mom.
I know that is not a high bar.
*/
type Stream interface {
	Goto(pos int) error
	Get_symbol() (rune, error)
	Get_block(blocksize int) (Datablock, error)
	Write_symbol(symbol rune) error
	Write_block(block Datablock) error
}

/*
Mostly for testing lol. I hope I do not have to compress lists. I do not have a problem with it,
it just sounds very redundant. Also, I'll proceed to fuck myself. I did not lookup for character type in Go.
*/
type ListStream struct {
	list     []rune
	position int
}

//! Generic functions, just copy paste these while defining new streams. You should be done with minimal damage to the goods.
/*
Good old library byllshit.
*/
func Get_ListStream(z []rune, position int) ListStream {
	return ListStream{list: z, position: position}
}

/*
Seeks. Lol.
*/
func (t *ListStream) Goto(pos int) error {
	if pos >= len(t.list) {
		return errors.New("Invalid Position")
	}
	t.position = pos
	return nil
}

// Gets the current symbol in the stream. Returns a big ol' error if that does not work
func (t *ListStream) Get_symbol() (rune, error) {
	if t.position >= len(t.list) {
		return 'a', errors.New("Invalid Position") //TODO: Check this at the 'client' side.
	}
	t.position += 1
	return t.list[t.position-1], nil
}

/*
Writes the rune to the list stream. Two ways: pos < len -> changes
pos >= len -> appends
*/
func (t *ListStream) Write_symbol(s rune) error {
	/*
		Writes the rune to the stream.
	*/
	l := len(t.list)
	if t.position < l {
		t.list[t.position] = s
		return nil
	}
	t.list = append(t.list, s)
	return nil
}

/*
Being the shameless copycat that I am, I am going to again basically translate
SCLs code. And copy their TODO: investigate faster ways of reading a block.
*/
func (t *ListStream) Get_block(blocksize int) (Datablock, error) {
	Datalist := make([]rune, 0)
	for i := 0; i < blocksize; i += 1 {
		s, err := t.Get_symbol()
		if err != nil {
			break
		}
		Datalist = append(Datalist, s)
	}
	if len(Datalist) == 0 {
		return Datablock{[]rune{}}, errors.New("Empty Block")
	}
	return Datablock{Datalist}, nil
}

/*
Write_block should be an easy function to implement, but what do I know?
*/
func (t *ListStream) Write_block(block Datablock) error {
	for _, v := range block.Data {
		t.Write_symbol(v)
	}
	return nil
}

// ! File stream yayy. Need to learn fileops in go, brb going to fuck myself.
//! Second thoughts: fuck it. We ball. Implementing algorithms now. Or maybe a nice wrapper around them.

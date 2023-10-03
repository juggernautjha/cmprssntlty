/*
The good thing about Go is that functions are under namespaces (package in this case)
and I can therefore keep splitting the code till the second coming of christ with little to
no consequence.
*/
package utilities

/*
Datablock -> wrapper around a list of symbols.
*/
type Datablock struct {
	data []string
}

func Get_datablock(t []string) Datablock {
	/*
		Utility function, converts string to a datablock.
	*/
	return Datablock{t}
}

func (t *Datablock) Get_and_count() ([]string, map[string]int64) {
	/*
		Gets a frequency count. Returns a list of symbols and a map of frequencies.
	*/
	var temp_map map[string]int64 = make(map[string]int64)
	for _, v := range t.data {
		temp_map[v] += 1
	}
	var temp_list []string
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

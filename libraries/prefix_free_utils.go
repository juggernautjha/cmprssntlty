/*
For doing `things` with Data.
Author: Jha
The good thing about Go is that functions are under namespaces (package in this case)
and I can therefore keep splitting the code till the second coming of christ with little to
no consequence.
! Boy testing this is going to be fun.
! Yes I copied the `docstring` from Data_utils.go.
*/
package libraries

import (
	"fmt"
	. "jcl/utilities"
	// . "jcl/utilities" //like modiji said, fuck namespace collisions.
)

type BitMap map[rune]string //!simple 1-1 correspondence between symbols and codewords
type MapBit map[string]rune //!simple 1-1 correspondence between codewords and symbols

/*Cannot make it simpler, takes in a map and a Datablock, encrypts it.*/
func Encode_block(block Datablock, b BitMap) Datablock {
	var new_block []rune
	for _, v := range Get_data(&block) {
		for _, k := range b[v] {
			new_block = append(new_block, k)
		}
		// new_block = append(new_block, []rune(b[v])...)
	}
	// Get_Datablock(new_block)
	return Datablock{Data: new_block}
}

/*
Decodes prefix free codes
*/
func Decode_block(block Datablock, b BitMap) Datablock {
	var inverse MapBit = make(MapBit)
	for i, j := range b {
		inverse[j] = i
	}
	var decoded []rune
	temp := ""
	for _, v := range Get_data(&block) {
		temp += string(v)
		val, ok := inverse[temp]
		if ok {
			decoded = append(decoded, val)
			temp = ""
		}
	}
	return Datablock{Data: decoded}
}

/*With basics out of the way, let us now focus on creating BitMaps*/
type Tree struct {
	symbol rune
	val    float64
	left   *Tree
	right  *Tree
}

/*Tree Based Methods Mashallah*/

/*Checks if the node is a leaf node*/
func Is_leaf(t *Tree) bool {
	return t.left == nil && t.right == nil
}

/*Traverses the tree, gets encoding for each symbol mashallah*/
func do_dfs(root *Tree, current string, mapping BitMap) {
	if Is_leaf(root) {
		mapping[root.symbol] = current
	}
	if root.left != nil {
		do_dfs(root.left, current+"0", mapping)
	}
	if root.right != nil {
		do_dfs(root.right, current+"1", mapping)
	}
}

/*Traverses the tree to get the encoding hashmap. */
func Get_encoding(root *Tree) BitMap {
	var mapping BitMap
	do_dfs(root, "", mapping)
	return mapping
}

/*How do we test the above function? Ofc by creating a method to create trees*/

/*Adds node in the right position*/
func add_node(root *Tree, code string, symbol rune) {
	curr_node := root
	for _, bit := range code {
		// curr_node.describe()
		if bit == '0' {
			if curr_node.left == nil {
				curr_node.left = &Tree{}
			}
			curr_node = curr_node.left
		}
		if bit == '1' {
			if curr_node.right == nil {
				curr_node.right = &Tree{}
			}
			curr_node = curr_node.right
		}

	}
	curr_node.symbol = symbol
}

/*DEBUG*/
func (t *Tree) describe() {
	println(t.symbol)
	println(t.left == nil)
	print(t.right == nil)
}

/*DEBUG*/
func (t *Tree) print_tree() {
	if t == nil {
		return
	}
	fmt.Printf("%c ", t.symbol)
	t.left.print_tree()
	t.right.print_tree()
}

func Create_tree(b BitMap) Tree {
	root_node := new(Tree)
	for k, v := range b {
		add_node(root_node, v, k)
	}
	root_node.print_tree()
	return *root_node
}

package node

import (
	"strconv"
	"strings"

	"github.com/GeoinovaDev/lower-resultys/io/file"
)

// Dump ...
func (node *Node) Dump(filename string) {
	output := &[]string{}

	_printNode(node, []int{0}, output)

	file.WriteContent(filename, strings.Join(*output, "\n"))
}

func _printNode(node *Node, indexes []int, output *[]string) {
	if node.Kind == "value" {
		_printValue(node, indexes, output)
	}

	if node.Childrens != nil {
		for i := 0; i < len(node.Childrens); i++ {
			if node.Childrens[i].Kind == "value" {
				// _checkIndex(node.Childrens[i], indexes, output)
				_printValue(node.Childrens[i], indexes, output)
			} else if node.Childrens[i].Kind == "array" {
				_printNode(node.Childrens[i], append(indexes, 0), output)
			}
			indexes[len(indexes)-1]++
		}
	}
}

func _checkIndex(node *Node, indexes []int, output *[]string) {
	arr := []string{}
	for i := 0; i < len(indexes); i++ {
		arr = append(arr, strconv.Itoa(indexes[i]))
	}

	r := "false"
	if strings.Join(arr, "") == node.Value {
		r = "true"
	}

	*output = append(*output, _printIndexes(indexes)+" = "+r)
}

func _printValue(node *Node, indexes []int, output *[]string) {
	*output = append(*output, _printIndexes(indexes)+" = "+node.Value)
}

func _printIndexes(indexes []int) string {
	arr := []string{}
	for i := 0; i < len(indexes); i++ {
		arr = append(arr, strconv.Itoa(indexes[i]))
	}

	return "[" + strings.Join(arr, ",") + "]"
}

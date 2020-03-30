# binary_search_tree

## Prerequisites:
You need to install - [go 1.13](https://golang.org/dl/)

## How to test:
1. Clone this repo: `git clone https://github.com/Richardowsky/binary_search_tree.git`
2. `make test`

## How to use:
```go
package example

import bst "binary_search_tree/src"

func example()  {
	tree := &bst.Node{
		Data: 3,
		 Left: &bst.Node{
			Data: 5,
		},
		 Right: &bst.Node{
			 Data: 2,
		},
	}
	
	bst.CheckBST(tree)
}

```
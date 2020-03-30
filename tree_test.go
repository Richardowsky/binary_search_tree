package binary_search_tree

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test1(t *testing.T) {
	tree := &Node{
		Data: 3,
		Left:  &Node{
			Data: 5,
			Left:  &Node{
				Data: 1,
			},
			Right: &Node{
				Data: 4,
			},
		},
		Right: &Node{
			Data: 2,
			Left:  &Node{
				Data: 6,
			},
		},
	}
	assert.Equal(t, CheckBST(tree), false)
}

func Test2(t *testing.T) {
	tree := &Node{
		Data: 10,
		Left:  &Node{
			Data: 8,
			Left:  &Node{
				Data: 2,
				Left:  &Node{
					Data: 1,
				},
				Right: &Node{
					Data: 6,
					Left:  &Node{
						Data: 4,
						Left:  &Node{
							Data: 3,
						},
						Right: &Node{
							Data: 5,
						},
					},
				},
			},
		},
		Right: &Node{
			Data: 11,
			Right: &Node{
				Data: 14,
				Left:  &Node{
					Data: 13,
				},
				Right: &Node{
					Data: 16,
				},
			},
		},
	}
	assert.Equal(t, CheckBST(tree), true)
}

func Test3(t *testing.T) {
	tree := &Node{
		Data: 15,
		Left:  &Node{
			Data: 15,
			Left:  &Node{
				Data: 15,
				Left:  &Node{
					Data: 15,
					Left:  &Node{
						Data: 15,
						Left:  &Node{
							Data: 15,
						},
					},
					Right: &Node{
						Data: 15,
						Left:  &Node{
							Data: 15,
						},
					},
				},
				Right: &Node{
					Data: 15,
					Left:  &Node{
						Data: 15,
						Left:  &Node{
							Data: 15,
						},
					},
					Right: &Node{
						Data: 15,
						Right: &Node{
							Data: 15,
						},
					},
				},
			},
			Right: &Node{
				Data: 15,
				Left:  &Node{
					Data: 15,
					Left:  &Node{
						Data: 15,
						Right: &Node{
							Data: 15,
						},
					},
					Right: &Node{
						Data: 15,
						Left:  &Node{
							Data: 15,
						},
					},
				},
				Right: &Node{
					Data: 15,
					Left:  &Node{
						Data: 15,
						Right: &Node{
							Data: 15,
						},
					},
					Right: &Node{
						Data: 15,
						Right: &Node{
							Data: 15,
						},
					},
				},
			},
		},
		Right: &Node{
			Data: 15,
			Left:  &Node{
				Data: 15,
				Left:  &Node{
					Data: 15,
					Left:  &Node{
						Data: 15,
						Left:  &Node{
							Data: 15,
						},
					},
					Right: &Node{
						Data: 15,
						Left:  &Node{
							Data: 15,
						},
					},
				},
				Right: &Node{
					Data: 15,
					Left:  &Node{
						Data: 15,
						Right: &Node{
							Data: 15,
						},
					},
					Right: &Node{
						Data: 15,
						Right: &Node{
							Data: 15,
						},
					},
				},
			},
			Right: &Node{
				Data: 15,
				Left:  &Node{
					Data: 15,
					Left:  &Node{
						Data: 15,
						Right: &Node{
							Data: 15,
						},
					},
					Right: &Node{
						Data: 15,
						Left:  &Node{
							Data: 15,
						},
					},
				},
				Right: &Node{
					Data: 15,
					Left:  &Node{
						Data: 15,
						Right: &Node{
							Data: 15,
						},
					},
					Right: &Node{
						Data: 15,
						Right: &Node{
							Data: 15,
						},
					},
				},
			},
		},
	}
	assert.Equal(t, CheckBST(tree), false)
}









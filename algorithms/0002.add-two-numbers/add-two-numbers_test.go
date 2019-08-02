package algs0002

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

type X struct {
    x1 *ListNode
    x2 *ListNode
}

type Y struct {
    y1 *ListNode
}

type Sample struct {
    x X
    y Y
}

func Test(t *testing.T) {
    _assert := assert.New(t)

    samples := []Sample{
        Sample {
            x: X{
                x1: makeListNode([]int{2,4,3}),
                x2: makeListNode([]int{5,6,4}),
            },
            y: Y{
                y1: makeListNode([]int{7, 0, 8}),
            },
        },
        Sample {
            x: X{
                x1: makeListNode([]int{1}),
                x2: makeListNode([]int{9}),
            },
            y: Y{
                y1: makeListNode([]int{0, 1}),
            },
        },
    }

    for _, sample := range samples {
        x, y := sample.x, sample.y
        _assert.Equal(y.y1, addTwoNumbers(x.x1, x.x2), "Sample: %v", x)
    }
}

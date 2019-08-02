package algs0004

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

type X struct {
    x1 []int
    x2 []int
}

type Y struct {
    y1 float64
}

type Sample struct {
    x X
    y Y
}

func Test(t *testing.T) {
    _assert := assert.New(t)
    samples := []Sample{
        Sample{
            x: X {
                x1: []int{1, 3},
                x2: []int{2},
            },
            y: Y{
                y1: 2.0,
            },

        },
        Sample{
            x: X {
                x1: []int{1, 2},
                x2: []int{3, 4},
            },
            y: Y{
                y1: 2.5,
            },

        },
    }

    for _, sample := range samples {
        x, y := sample.x, sample.y
        _assert.Equal(y.y1, findMedianSortedArrays(x.x1, x.x2), "sample: %v", x)
    }
}

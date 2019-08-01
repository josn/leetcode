package algs00

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

type X struct {
    x1 [] int
    x2 int
}

type Y struct {
    y1 int
}

type Sample struct {
    x X
    y Y
}

func Test(t *testing.T) {
    _assert := assert.New(t)

    samples := []Sample {
        Sample{
            x: X {
                x1: []int{1,2,3,5,6},
                x2: 5,
            },
            y: Y {
                y1: 3,
            },
        },
        Sample{
            x: X {
                x1: []int{2,4,6},
                x2: 5,
            },
            y: Y {
                y1: -1,
            },
        },
    }

    for _, sample := range samples {
        x, y := sample.x, sample.y
        _assert.Equal(y.y1, BinarySearch(x.x1, x.x2), "Sample: %v", sample)
    }
}

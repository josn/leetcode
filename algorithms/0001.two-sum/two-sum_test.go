package algs0001

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

type X struct {
    x1 []int
    x2 int
}

type Y struct {
    y1 []int
}

type Sample struct {
    x X
    y Y
}

func Test(t *testing.T) {
    _assert := assert.New(t)
    samples := []Sample {
        Sample {
            x: X {
                x1: []int {2, 7, 11, 15},
                x2: 9,
            },
            y: Y {
                y1: []int{0, 1},
            },
        },
        Sample {
            x: X {
                x1: []int {2, 7, 11, 15},
                x2: 18,
            },
            y: Y {
                y1: []int{1, 2},
            },
        },
    }

    for _, sample := range samples {
        x, y := sample.x, sample.y
        _assert.Equal(y.y1, twoSum(x.x1, x.x2), "Sample: %v", x)
    }
}

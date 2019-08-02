package algs0003

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

type X struct {
    x1 string
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
        Sample {
            x : X{
                x1: "abcabcbb",
            },
            y: Y {
                y1: 3,
            },
        },
    }

    for _, sample := range samples {
        x, y := sample.x, sample.y
        _assert.Equal(y.y1, lengthOfLongestSubstring(x.x1), "Sample: %v", x)
    }
}

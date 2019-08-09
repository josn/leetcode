package algs0009

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

type Sample struct {
    x int
    y bool
}

func Test(t *testing.T) {
    _assert := assert.New(t)
    samples := []Sample {
        Sample {
            x: 121,
            y: true,
        },
        Sample {
            x: -121,
            y: false,
        },
        Sample {
            x: 0,
            y: true,
        },
        Sample {
            x: 10,
            y: false,
        },
    }

    for _, sample := range samples {
        x, y := sample.x, sample.y
        _assert.Equal(y, isPalindrome(x), "Sample: %v", sample)
    }
}

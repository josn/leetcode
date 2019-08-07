package algs0005

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

type X struct {
    x1 string
}

type Y struct {
    y1 string
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
                x1: "a",
            },
            y: Y{
                y1: "a",
            },
        },
        Sample{
            x: X {
                x1: "ab",
            },
            y: Y{
                y1: "a",
            },
        },
        Sample{
            x: X {
                x1: "aba",
            },
            y: Y{
                y1: "aba",
            },
        },
        Sample{
            x: X {
                x1: "abba",
            },
            y: Y{
                y1: "abba",
            },
        },
        Sample{
            x: X {
                x1: "abbaba",
            },
            y: Y{
                y1: "abba",
            },
        },
    }

    for _, sample := range samples {
        x, y := sample.x, sample.y
        _assert.Equal(y.y1,longestPalindrome(x.x1), "Sample: %v", sample)
    }
}

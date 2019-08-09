package algs0008

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
        Sample{
            x: X {
                x1: "",
            },
            y: Y {
                y1: 0,
            },
        },
        Sample{
            x: X {
                x1: "123",
            },
            y: Y {
                y1: 123,
            },
        },
        Sample{
            x: X {
                x1: "-123",
            },
            y: Y {
                y1: -123,
            },
        },
        Sample{
            x: X {
                x1: "a -123",
            },
            y: Y {
                y1: 0,
            },
        },
        Sample{
            x: X {
                x1: "-123a",
            },
            y: Y {
                y1: -123,
            },
        },
        Sample{
            x: X {
                x1: "-91283472332",
            },
            y: Y {
                y1: -2147483648,
            },
        },
    }

    for _, sample := range samples {
        x, y := sample.x, sample.y
        _assert.Equal(y.y1, myAtoi(x.x1), "Sample :%v", sample)
    }
}

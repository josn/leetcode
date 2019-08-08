package algs0006

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

type X struct {
    x1 string
    x2 int
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
    samples := []Sample {
        Sample{
            x: X{
                x1: "PAYPALISHIRING",
                x2: 3,
            },
            y: Y {
                y1: "PAHNAPLSIIGYIR",
            },
        },
        Sample{
            x: X{
                x1: "PAYPALISHIRING",
                x2: 4,
            },
            y: Y {
                y1: "PINALSIGYAHRPI",
            },
        },
    }

    for _, sample := range samples {
        x, y := sample.x, sample.y
        _assert.Equal(y.y1, convert(x.x1, x.x2), "Sample: %v", sample)
    }
}

package algs0007

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

type X struct {
    x1 int
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
            x: X{
                x1: 123,
            },
            y: Y{
                y1: 321,
            },
        },
        Sample{
            x: X{
                x1: 123456789,
            },
            y: Y{
                y1: 987654321,
            },
        },
        Sample{
            x: X{
                x1: -120,
            },
            y: Y{
                y1: -21,
            },
        },
    }

    for _, sample := range samples {
        x, y := sample.x, sample.y
        _assert.Equal(y.y1, reverse(x.x1), "Sample: %v", sample)
    }
}

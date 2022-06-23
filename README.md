# fgo

A generic slice manipulation package for Go.

## Install

    go get github.com/jefflinse/fgo

## Import

    import (
        "github.com/jefflinse/fgo"
    )

## Use

### Filter

    s := []int{1, 2, 3, 4, 5}
    fs := fgo.Filter(s, func(x int) int {
        return x > 3
    })
    fmt.Println(mul) // 4, 5

### Map

    s := []int{1, 2, 3, 4, 5}
    ms := fgo.Map(s, func(x int) int {
        return x * 2
    })
    fmt.Println(ms) // 2, 4, 6, 8, 10

### Reduce

    s := []int{1, 2, 3, 4, 5}
    rs := fgo.Reduce(s, func(v string, x int) int {
        return v + string(x)
    }, "")
    fmt.Println(rs) // "12345"

## License

MIT

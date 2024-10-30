package main

import "testing"

func TestAddNegative(t *testing.T) {
    x := 1
    y := -3

    if result := Add(x, y); result != -2 {
        t.Fatalf(`Add(1, -3) = %d; want -2`, result)
    }
}
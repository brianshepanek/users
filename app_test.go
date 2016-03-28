package main 

import (
   "testing"
)


func TestTruth(t *testing.T) {
    if true != true {
        t.Error("everything I know is wrong")
    }
}

func TestTruth2(t *testing.T) {
    if true != true {
        t.Error("everything I know is wrong")
    }
}

func TestTruth3(t *testing.T) {
    if true != true {
        t.Error("everything I know is wrong")
    }
}
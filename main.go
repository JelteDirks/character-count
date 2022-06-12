package main

import (
//    "flag"    
    "character-count/accumulator"
)

func main() {
    dummy2 := &accumulator.Accumulator{
        Pattern: "/Users/jeltedirks/github.com/jeltedirks/motor-shield-control/*/*.rs",
        Cutoff: 30,
    }
    dummy2.ParseContent()
    dummy2.PrintData()
}

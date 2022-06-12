package accumulator

import (
    "path/filepath"
    "fmt"
    "io/ioutil"
)

type Accumulator struct {
    Pattern string
    Cutoff int
    Data map[string]int
}

func (a *Accumulator) ParseContent() {
    matches, err := filepath.Glob(a.Pattern)
    a.Data = make(map[string]int)

    if err != nil {
        fmt.Printf("%v", err)
    }
    
    for _, s := range matches {
        fmt.Printf("Parsing %v'\n",s)

        bytes, err := ioutil.ReadFile(s)
        if err != nil {
            panic(err)
        }

        for _, b := range string(bytes) {
            bs := string(b)
            _, err := a.Data[bs]
            if !err {
                a.Data[bs] = 0 
            }
            a.Data[bs]++
        }
    }
}

func (a *Accumulator) PrintData () {
}

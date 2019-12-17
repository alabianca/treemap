# Treemap
A sorted map implementation generator. 

### Generate a treemap implementation for any key and value type
Once installed in `GOBIN` run `treemap-gen -key={keytype} -value={valuetype} -name={treeName} -out{outfile.go}`
Example: `treemap-gen -key=string -value=int -name=String -out=treemap.go` will generate a map that is keyed by string and where values are integers.

### Usage
A treemap needs a comparator function. The comparator function is provided by the generator. It's argument types are the same as the map's keys.

```go
    type CompareFunc func(a string, b string) int // if -key=string
```

Example Treemap keyed by string
```go
    // keep the map sorted by string length ascending
    comp := func(a string, b string) int {
        return len(b) - len(a)
    }
    t := NewStringMap(comp)
    t.Insert("aaa", 1)
    t.Insert("a", 45)
    t.Insert("aa", 23)
    fmt.Println(t) // a(45)->aa(23)->aaa(1)
```

Keys can hold any type just like a regular map. `treemap-gen -key=Person -value=int -name=Person -out=treemap.go`
This will create a map keyed by Person struct. 
```go
    // keep the map sorted by age ascending
    comp := func(a Person, b Person) int {
        return a.Age - b.Age
    }
    p1 := Person{Age:34}
    p2 := Person{Age:21}
    p3 := Person{Age:8}

    t := NewStringMap(comp)
    t.Insert(p1, p1.Age)
    t.Insert(p2, p2.Age)
    t.Insert(p3, p3.Age)
    fmt.Println(t)
```

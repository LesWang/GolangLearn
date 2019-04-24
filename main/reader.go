package main

import(
    "fmt"
    "io"
)

type MyReader struct {

}

func (reader MyReader) Read(bytes []byte) (int, error) {
    for i := range bytes {
        bytes[i] = 'A'
    }
    return len(bytes), nil
}

type Rot13Reader struct {
    r io.Reader
}

func (reader Rot13Reader) Read(b []byte) (int, error) {
    n, err := reader.r.Read(b)
    if err != nil {
        for i:=0; i<n; i++{
            switch {
            case b[i] >= 'A' && b[i] <'N' :
                b[i] += 13
            case b[i]>='N' && b[i] <='Z':
                b[i] -=13
            case b[i] >= 'a' && b[i] <'n' :
                b[i] += 13
            case b[i]>='n' && b[i] <='z':
                b[i] -=13
            }
        }
    }
    return n, err
}

func main()  {
    var b = make([]byte, 10)
    var reader = MyReader{}
    for i := 0; i < 10; i++ {
        n, err := reader.Read(b)
        fmt.Println(n, err)
        fmt.Println(string(b[:n]))
    }
}

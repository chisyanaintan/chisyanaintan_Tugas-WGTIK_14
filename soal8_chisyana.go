package main
import "fmt"

func main() {
    var a,b bool
    fmt.Scan(&a, &b)
    if a == true && b == true {
        fmt.Print("keluar jalan-jalan")
    }else{
        fmt.Print("diam di rumah saja")
    }
}

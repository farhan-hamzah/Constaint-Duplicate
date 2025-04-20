package main
import "fmt"

func main(){
	const NMAX int = 100
	var A[NMAX]int
	var n, i, tebak, arr int
	var b bool
	fmt.Scan(&n, &tebak)
	for i =0; i <n; i++{
		fmt.Scan(&A[i])
		if i == tebak{
			arr = i
		}
	}
	if arr == tebak{
		b = true
	}else{
		b = false
	}
	fmt.Print(b)
}
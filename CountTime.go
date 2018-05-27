package main

import (
	"time"
	"net/http"
	"fmt"
)

func main() {

	start := time.Now()
	resp, err := http.Get("https://fb.com")
	if err!=nil {
		fmt.Println(resp)
	}


	/* soal no.1
	#error cannot sub time between time.Time datatype

	elapsed := time.Now() - start
	fmt.Println(elapsed)
	 */
	fmt.Println("soal no.1 : in-ccorect")
	fmt.Println("error cannot sub time between time.Time datatype")

	/* soal no.2
	correct
	*/
	fmt.Println("soal no.2 : correct")
	elapsed2 :=  time.Since(start).Seconds()
	fmt.Println(elapsed2)


	/* soal no.3 */
	fmt.Println("soal no.3 : in-correct")
	fmt.Println("return bool, to check if current time > start time")
	elapsed3 := time.Now().After(start)
	fmt.Println(elapsed3)


	/* soal no.4 */
	fmt.Println("soal no.4 : correct")
	elapsed4 := time.Now().Sub(start)
	fmt.Println(elapsed4)




}

# Pair Socks Problem

> Explain your solution and how to run the program in Markdown formatted.

```go
func sockMerchant(arr []int) int {
	// declare paris variable for save counter each element
	pairs := map[int]int{}

    // declare counter variable for save all counter
	counter := 0

	
	for i := 0; i < len(arr); i++ {
		// counter element untuk mengetahui banyaknya jumlah element yang muncul
		pairs[arr[i]]++
	}

	// loop all counter
	for _, v := range pairs {
		// hasil counter dibagi 2 untuk mencari jumlah pasangan, kalau ada sisa hasil bagi berarti element tersebut tidak dapat pasang
		// jumlah hasil bagi untuk mengetahui total yang mendapatkan pasangan
		counter += v / 2
	}

	return counter
}

```
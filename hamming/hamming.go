package hamming

import ("math/bits";
        //"fmt";
        "../xor")


func combinations(n, m int, p [][]int) [][]int {
	// For each combination of m elements out of n
	// call the function f passing a list of m integers in 0-n
	// without repetitions
	s := make([]int, m)
    var u [][]int
	last := m - 1
	var rc func(int, int)
	rc = func(i, next int) {
		for j := next; j < n; j++ {
			s[i] = j
			if i == last {
                t:=make([]int, len(s))
                copy(t, s)
                u=append(u,t)
			} else {
				rc(i+1, j+1)
			}
		}
		return
	}
	rc(0, 0)
    return u
}

func Distance(key1 []byte, key2 []byte) int {
    key1_bytes := key1
    key2_bytes := key2
    var total, size int
    total = 0
    if len(key1)==len(key2) {
        size = len(key1)
        var xored byte
        for x:=0;x<size;x++ {
        xored = xor.Xor(key1_bytes[x], key2_bytes[x])
        total = total + bits.OnesCount(uint(xored))
        }
        return total
    } else {
        println("Different key lengths!\n")
        return -1
    }
}


func Normalized_distance(ciphertext []byte, keysize int) float64 {
    var ndistances []float64
    var pairs [][]int
    var blocks [][]byte
    for i:=0;i+keysize<len(ciphertext);i+=keysize {
        blocks=append(blocks, ciphertext[i:i+keysize])
    }
    pairs = combinations(len(blocks), 2, pairs)
    for _,pair := range(pairs) {
        distance := Distance(blocks[pair[0]], blocks[pair[1]])
        var normalized float64
        normalized = float64(distance)/float64(keysize)
        ndistances = append(ndistances, normalized)
    }
    var sum float64
    sum = 0
    for _, ndist := range ndistances {
        sum += ndist
    }
    median_ndistances := sum/float64(len(ndistances))
    return median_ndistances
}

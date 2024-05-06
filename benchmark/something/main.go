package something

import "fmt"

func makeSomething(n int) []string {
	var r []string
	for i := 0; i < n; i++ {
		r = append(r, fmt.Sprintf("%05d 何か", i))
	}
	return r
}

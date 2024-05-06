package something

import "fmt"

func makeSomething(n int) []string {
	// NOTE: -countで6回以上実行する必要がある。
	r := make([]string, n)
	for i := 0; i < n; i++ {
		r[i] = fmt.Sprintf("%05d 何か", i)
	}
	return r
}

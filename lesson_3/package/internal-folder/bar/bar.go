package bar

import "github.com/kadirgonen/Go-Bootcamp/lesson_3/package/internal-folder/foo/internal"

func FailUseDoubler(i int) int {
	return internal.Doubler(i)
}

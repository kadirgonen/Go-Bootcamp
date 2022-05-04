package foo

import "github.com/kadirgonen/Go-Bootcamp/lesson_3/package/internal-folder/foo/internal"

func UseDoubler(i int) int {
	return internal.Doubler(i)
}

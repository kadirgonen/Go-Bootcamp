package sibling

import "github.com/kadirgonen/Go-Bootcamp/internal-folder/foo/internal"

func AlsoUseDoubler(i int) int {
	return internal.Doubler(i)
}

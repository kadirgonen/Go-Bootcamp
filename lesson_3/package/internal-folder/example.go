package internal_package_example

import "github.com/kadirgonen/Go-Bootcamp/internal-folder/foo/internal"

func FailUseDoubler(a int) int {
	return internal.Doubler(a)
}

package main

import "github.com/kadirgonen/Go-Bootcamp/lesson_3/package/internal-folder/foo/internal"

func FailUseDoubler(a int) int {
	return internal.Doubler(a)
}

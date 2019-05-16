package repository

import (
	"github.com/aditi-agarwal-tw/chimp/utility"
)

type ChimpRepository struct {
	workTree string
	gitDir   string
}

func (chimpRepo ChimpRepository) Initialize(path string) {
	utility.CreateDirectory(path, ".chimp", false)
}

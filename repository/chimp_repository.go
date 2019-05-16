package repository

import (
	"github.com/aditi-agarwal-tw/chimp/utility"
)

type ChimpRepository struct {
	workTree string
	gitDir   string
}

func (chimpRepo ChimpRepository) Initialize(path string) error {
	err := utility.CreateDirectory(path, ".chimp", false)
	return err
}

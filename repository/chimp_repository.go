package repository

type ChimpRepository struct {
	workTree string
	gitDir   string
}

func (chimpRepo ChimpRepository) Initialize(path string) {
}

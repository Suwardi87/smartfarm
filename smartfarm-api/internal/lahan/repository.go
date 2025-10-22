package lahan

type lahanRepository struct{}

func NewLahanRepository() LahanRepository {
	return &lahanRepository{}
}




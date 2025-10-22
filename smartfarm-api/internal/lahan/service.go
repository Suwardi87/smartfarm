package lahan


// 🔧 Implementasi Service
type LahanService struct {
	repo LahanRepository
}

func NewLahanService(repo LahanRepository) LahanService {
	return &LahanService{repo: repo}
}
package lahan


type lahanService struct {
	repo LahanRepository
}

func NewLahanService(repo LahanRepository) *lahanService {
	return &lahanService{repo: repo}
}


// 🧱 Interface: definisi kontrak layanan
type PetaniService interface {
	// CreatePetani(req CreatePetaniRequest) (*PetaniResponse, error)
	GetAllPetani() ([]LahanResponse, error)
// ✅ Tambahkan ini
	// ✅ Method baru pagination
	GetAllPetaniPaginated(page, limit int) ([]LahanResponse, int64, error)
	// GetPetaniByID(id uint) (*PetaniResponse, error)
	// UpdatePetani(id uint, req UpdatePetaniRequest) (*PetaniResponse, error)
	// DeletePetani(id uint) error
}
// Di file lahan_service.go
func (s *lahanService) GetAllLahanPaginated(page, limit int) ([]LahanResponse, int64, error) {
    // panggil repository
    lahans, total, err := s.repo.FindPaginated(page, limit)
    if err != nil {
        return nil, 0, err
    }

    // mapping ke DTO response
    var responses []LahanResponse
    for _, p := range lahans {
        responses = append(responses, LahanResponse{
            ID:         p.ID,
            PetaniID:   p.PetaniID,
            NamaPetani: p.Petani.Nama,
            NamaLahan:  p.NamaLahan,
            Lokasi:     p.Lokasi,
            Luas:       p.Luas,
            JenisTanah: p.JenisTanah,
            Koordinat:  p.Koordinat,
            CreatedAt:  p.CreatedAt,
            UpdatedAt:  p.UpdatedAt,
        })
    }

    return responses, total, nil
}



package petani

import "errors"

// 🔧 Implementasi Service
type petaniService struct {
	repo PetaniRepository
}

func NewPetaniService(repo PetaniRepository) PetaniService {
	return &petaniService{repo: repo}
}

// 🧱 Interface: definisi kontrak layanan
type PetaniService interface {
	CreatePetani(req CreatePetaniRequest) (*PetaniResponse, error)
	GetAllPetani() ([]PetaniResponse, error)
// ✅ Tambahkan ini
	// ✅ Method baru pagination
	GetAllPetaniPaginated(page, limit int) ([]PetaniResponse, int64, error)
	GetPetaniByID(id uint) (*PetaniResponse, error)
	UpdatePetani(id uint, req UpdatePetaniRequest) (*PetaniResponse, error)
	DeletePetani(id uint) error
}


// ✅ Create Petani
func (s *petaniService) CreatePetani(req CreatePetaniRequest) (*PetaniResponse, error) {
	petani := Petani{
		Nama:    req.Nama,
		Alamat: req.Alamat,
	}

	if err := s.repo.Create(&petani); err != nil {
		return nil, err
	}

	resp := PetaniResponse{
		ID:        petani.ID,
		Nama:      petani.Nama,
		Alamat:   petani.Alamat,
		CreatedAt: petani.CreatedAt,
		UpdatedAt: petani.UpdatedAt,
	}

	return &resp, nil
}

// ✅ Get All Petani
func (s *petaniService) GetAllPetani() ([]PetaniResponse, error) {
	petanis, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	// konversi model ke DTO response
	var responses []PetaniResponse
	for _, p := range petanis {
		responses = append(responses, PetaniResponse{
			ID:        p.ID,
			Nama:      p.Nama,
			Alamat:    p.Alamat,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}
	return responses, nil
}

// ✅ Get Petani by ID
func (s *petaniService) GetPetaniByID(id uint) (*PetaniResponse, error) {
	petani, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("petani tidak ditemukan")
	}

	resp := PetaniResponse{
		ID:        petani.ID,
		Nama:      petani.Nama,
		Alamat:    petani.Alamat,
		CreatedAt: petani.CreatedAt,
		UpdatedAt: petani.UpdatedAt,
	}
	return &resp, nil
}

func (s *petaniService) GetAllPetaniPaginated(page, limit int) ([]PetaniResponse, int64, error) {
    // panggil repository
    petanis, total, err := s.repo.FindPaginated(page, limit)
    if err != nil {
        return nil, 0, err
    }

    // mapping ke DTO response
    var responses []PetaniResponse
    for _, p := range petanis {
        responses = append(responses, PetaniResponse{
            ID:        p.ID,
            Nama:      p.Nama,
            Alamat:   p.Alamat,
            CreatedAt: p.CreatedAt,
            UpdatedAt: p.UpdatedAt,
        })
    }

    return responses, total, nil
}



// ✅ Update Petani
func (s *petaniService) UpdatePetani(id uint, req UpdatePetaniRequest) (*PetaniResponse, error) {
	petani, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("petani tidak ditemukan")
	}

	petani.Nama = req.Nama
	petani.Alamat = req.Alamat

	if err := s.repo.Update(petani); err != nil {
		return nil, err
	}

	resp := PetaniResponse{
		ID:        petani.ID,
		Nama:      petani.Nama,
		Alamat:    petani.Alamat,
		CreatedAt: petani.CreatedAt,
		UpdatedAt: petani.UpdatedAt,
	}
	return &resp, nil
}

// ✅ Delete Petani
func (s *petaniService) DeletePetani(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("petani tidak ditemukan")
	}
	return s.repo.Delete(id)
}

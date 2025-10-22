import { ref } from 'vue'
import { api } from '@/api/auth'

export function usePetani() {
  const daftarPetani = ref([])
  const petani = ref(null)
  const loading = ref(false)
  const error = ref(null)

  // GET - Ambil semua data
  const ambilSemuaPetani = async () => {
    loading.value = true
    try {
      const res = await api.get('/admin/petani')
      daftarPetani.value = res.data.data || res.data
    } catch (err) {
      error.value = err.response?.data?.message || err.message
    } finally {
      loading.value = false
    }
  }

  // GET - Ambil detail petani
  const ambilPetaniById = async (id) => {
    loading.value = true
    try {
      const res = await api.get(`/admin/petani/${id}`)
      petani.value = res.data.data || res.data
    } catch (err) {
      error.value = err.response?.data?.message || err.message
    } finally {
      loading.value = false
    }
  }

  // POST - Tambah petani
  const tambahPetani = async (dataBaru) => {
    try {
      const res = await api.post('/admin/petani', dataBaru)
      daftarPetani.value.push(res.data.data || res.data)
      return res
    } catch (err) {
      error.value = err.response?.data?.message || err.message
      throw err
    }
  }

  // PUT - Ubah petani
  const ubahPetani = async (id, dataBaru) => {
    try {
      const res = await api.put(`/admin/petani/${id}`, dataBaru)
      const index = daftarPetani.value.findIndex(p => p.id === id)
      if (index !== -1) daftarPetani.value[index] = res.data.data || res.data
      return res
    } catch (err) {
      error.value = err.response?.data?.message || err.message
      throw err
    }
  }

  // DELETE - Hapus petani
  const hapusPetani = async (id) => {
    try {
      await api.delete(`/admin/petani/${id}`)
      daftarPetani.value = daftarPetani.value.filter(p => p.id !== id)
    } catch (err) {
      error.value = err.response?.data?.message || err.message
      throw err
    }
  }

  return {
    daftarPetani,
    petani,
    loading,
    error,
    ambilSemuaPetani,
    ambilPetaniById,
    tambahPetani,
    ubahPetani,
    hapusPetani,
  }
}

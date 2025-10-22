import axios from 'axios'

// Buat instance axios utama
export const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8083/api',
  headers: {
    'Content-Type': 'application/json',
  },
})

// Contoh fungsi tambahan (bisa untuk auth)
export const login = async (credentials) => {
  return await api.post('/auth/login', credentials) 
}

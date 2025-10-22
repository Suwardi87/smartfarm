// src/stores/authStore.js
import { defineStore } from 'pinia'
import { login } from '@/api/auth'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
    token: null,
    loading: false,
    error: null,
  }),

  actions: {
    async loginUser(credentials) {
      this.loading = true
      this.error = null
      try {
        const { data } = await login(credentials)
        this.token = data.token
        this.user = data.user
        localStorage.setItem('token', data.token)
        localStorage.setItem('user', JSON.stringify(data.user))
        return true
      } catch (err) {
        this.error = err.response?.data?.message || 'Login gagal'
        return false
      } finally {
        this.loading = false
      }
    },

    logout() {
      this.user = null
      this.token = null
      localStorage.clear()
    },
  },
})

<script setup>
import { ref } from 'vue'
import { useAuthStore } from '@/store/useAuthStore'
import { useRouter } from 'vue-router'

const auth = useAuthStore()
const router = useRouter()

const username = ref('')
const password = ref('')
const errorMessage = ref('')

const handleLogin = async () => {
  const success = await auth.loginUser({
    username: username.value,
    password: password.value,
  })
  if (success) {
    router.push('/')
  } else {
    errorMessage.value = auth.error
  }
}
</script>

<template>
  <div class="w-full max-w-md mx-auto p-8 bg-white rounded-2xl shadow-md">
    <h1 class="text-2xl font-semibold mb-6 text-center">Masuk ke SmartFarm</h1>

    <div v-if="errorMessage" class="text-red-500 mb-4 text-sm">
      {{ errorMessage }}
    </div>

    <form @submit.prevent="handleLogin" class="space-y-4">
      <div>
        <label class="block mb-1 text-sm font-medium">username</label>
        <input
          v-model="username"
          type="username"
          class="w-full p-3 border rounded-md focus:ring focus:ring-green-200 outline-none"
          required
        />
      </div>

      <div>
        <label class="block mb-1 text-sm font-medium">Password</label>
        <input
          v-model="password"
          type="password"
          class="w-full p-3 border rounded-md focus:ring focus:ring-green-200 outline-none"
          required
        />
      </div>

      <button
        type="submit"
        class="w-full py-3 bg-green-600 hover:bg-green-700 text-white rounded-md font-semibold"
        :disabled="auth.loading"
      >
        {{ auth.loading ? 'Memproses...' : 'Login' }}
      </button>
    </form>
  </div>
</template>

<template>
  <section class="grid grid-cols-1 sm:grid-cols-3 gap-4">
    <!-- Total Petani -->
    <div class="bg-white shadow rounded-lg p-5 flex items-center justify-between">
      <div>
        <p class="text-sm text-gray-500">Total Petani</p>
        <h3 class="text-2xl font-bold text-gray-800">{{ totalPetani }}</h3>
      </div>
      <div class="bg-green-100 p-3 rounded-full">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-green-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
        </svg>
      </div>
    </div>

    <!-- Total Lahan -->
    <div class="bg-white shadow rounded-lg p-5 flex items-center justify-between">
      <div>
        <p class="text-sm text-gray-500">Total Lahan</p>
        <h3 class="text-2xl font-bold text-gray-800">{{ totalLahan }}</h3>
      </div>
      <div class="bg-blue-100 p-3 rounded-full">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-blue-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c1.657 0 3-1.343 3-3S13.657 2 12 2 9 3.343 9 5s1.343 3 3 3zM12 14v8m-4-4h8" />
        </svg>
      </div>
    </div>

    <!-- Rata-rata Panen -->
    <div class="bg-white shadow rounded-lg p-5 flex items-center justify-between">
      <div>
        <p class="text-sm text-gray-500">Rata-rata Panen</p>
        <h3 class="text-2xl font-bold text-gray-800">{{ rataPanen }} kg</h3>
      </div>
      <div class="bg-yellow-100 p-3 rounded-full">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-yellow-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3v18h18M9 9l3 3 6-6" />
        </svg>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { api } from '@/api/auth'

// state ringkasan
const totalPetani = ref(0)
const totalLahan = ref(0)
const rataPanen = ref(0)

// onMounted(() => {
//   ambilRingkasan()
// })


// async function ambilRingkasan() {
//   try {
//     // Ganti endpoint sesuai backend kamu, misal /api/petani/summary
//     const res = await api.get('/petani')
//     totalPetani.value = res.data.totalPetani || 0
//     totalLahan.value = res.data.totalLahan || 0
//     rataPanen.value = res.data.rataPanen || 0
//   } catch (err) {
//     console.error('Gagal memuat data ringkasan:', err)
//   }
// }
</script>

<template>
  <!-- Overlay -->
  <div class="fixed inset-0 z-50 flex items-center justify-center bg-black/40 backdrop-blur-sm">
    <div class="bg-white rounded-xl shadow-xl w-full max-w-md p-6 space-y-6 relative">

      <!-- Header -->
      <header class="flex items-center justify-between border-b pb-3">
        <h2 class="text-xl font-semibold text-gray-800">Detail Petani</h2>
        <button
          @click="$emit('tutup')"
          class="text-gray-500 hover:text-gray-700 transition"
        >
          ✕
        </button>
      </header>

      <!-- Konten Detail -->
      <div v-if="data" class="space-y-4">
        <div class="flex flex-col">
          <span class="text-sm text-gray-500">Nama Petani</span>
          <span class="text-base font-medium text-gray-800">{{ data.nama }}</span>
        </div>

        <div class="flex flex-col">
          <span class="text-sm text-gray-500">Alamat</span>
          <span class="text-base font-medium text-gray-800">{{ data.alamat }}</span>
        </div>

        <div class="flex flex-col" v-if="data.telepon">
          <span class="text-sm text-gray-500">No. Telepon</span>
          <span class="text-base font-medium text-gray-800">{{ data.telepon }}</span>
        </div>

        <div class="flex flex-col" v-if="data.created_at">
          <span class="text-sm text-gray-500">Tanggal Terdaftar</span>
          <span class="text-base font-medium text-gray-800">
            {{ formatTanggal(data.created_at) }}
          </span>
        </div>
      </div>

      <!-- Tombol Aksi -->
      <div class="flex justify-end border-t pt-4">
        <button
          @click="$emit('tutup')"
          class="px-4 py-2 rounded-lg bg-gray-200 text-gray-700 hover:bg-gray-300 transition"
        >
          Tutup
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { defineProps } from 'vue'

const props = defineProps({
  data: {
    type: Object,
    required: true
  }
})

// Fungsi bantu format tanggal
function formatTanggal(dateString) {
  const date = new Date(dateString)
  return date.toLocaleDateString('id-ID', {
    day: 'numeric',
    month: 'long',
    year: 'numeric'
  })
}
</script>

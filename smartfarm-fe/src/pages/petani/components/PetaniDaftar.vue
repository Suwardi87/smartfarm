<template>
  <section class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
    <!-- Jika loading -->
    <div v-if="loading" class="col-span-full flex justify-center items-center py-10 text-gray-500">
      Memuat data petani...
    </div>

    <!-- Jika tidak ada data -->
    <div v-else-if="!data || data.length === 0" class="col-span-full text-center text-gray-500 py-10">
      Belum ada data petani.
    </div>

    <!-- Daftar Petani -->
    <div
      v-else
      v-for="petani in data"
      :key="petani.id"
      class="bg-white shadow-md rounded-lg p-5 hover:shadow-lg transition flex flex-col justify-between"
    >
      <div>
        <h3 class="text-lg font-semibold text-gray-800">{{ petani.nama }}</h3>
        <p class="text-sm text-gray-600 mt-1">{{ petani.alamat }}</p>

        <div v-if="petani.telepon" class="flex items-center mt-2 text-gray-500 text-sm">
          <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5h2l3 7v5a2 2 0 002 2h6a2 2 0 002-2v-5l3-7h2" />
          </svg>
          {{ petani.telepon }}
        </div>
      </div>

      <!-- Tombol Aksi -->
      <div class="flex justify-end gap-2 mt-4 border-t pt-3">
        <button
          @click="$emit('lihat-detail', petani)"
          class="px-3 py-1 text-sm bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition"
        >
          Detail
        </button>
        <button
          @click="$emit('edit', petani)"
          class="px-3 py-1 text-sm bg-yellow-500 text-white rounded-lg hover:bg-yellow-600 transition"
        >
          Edit
        </button>
        <button
          @click="$emit('hapus', petani.id)"
          class="px-3 py-1 text-sm bg-red-600 text-white rounded-lg hover:bg-red-700 transition"
        >
          Hapus
        </button>
      </div>
    </div>
  </section>
</template>

<script setup>
defineProps({
  data: {
    type: Array,
    required: true
  },
  loading: {
    type: Boolean,
    default: false
  }
})

defineEmits(['lihat-detail', 'edit', 'hapus'])
</script>

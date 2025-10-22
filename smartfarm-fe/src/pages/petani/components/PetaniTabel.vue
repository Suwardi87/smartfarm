<template>
  <div class="bg-white shadow rounded-lg overflow-hidden my-4">
    <!-- Loading -->
    <div v-if="loading" class="p-6 text-center text-gray-500">
      Memuat data petani...
    </div>

    <!-- Tidak ada data -->
    <div v-else-if="!data || data.length === 0" class="p-6 text-center text-gray-500">
      Belum ada data petani.
    </div>

    <!-- Tabel Data -->
    <div v-else class="overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-100">
          <tr>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-600 uppercase tracking-wider">No</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-600 uppercase tracking-wider">Nama</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-600 uppercase tracking-wider">Alamat</th>
            <th class="px-4 py-3 text-left text-sm font-medium text-gray-600 uppercase tracking-wider">Aksi</th>
          </tr>
        </thead>

        <tbody class="bg-white divide-y divide-gray-100">
          <tr
            v-for="(petani, index) in data"
            :key="petani.id"
            class="hover:bg-gray-50 transition"
          >
            <td class="px-4 py-3 text-sm text-gray-700">{{ index + 1 }}</td>
            <td class="px-4 py-3 text-sm font-semibold text-gray-800">{{ petani.nama }}</td>
            <td class="px-4 py-3 text-sm text-gray-600">{{ petani.alamat }}</td>

            <!-- Tombol Aksi -->
            <td class="px-4 py-3 flex flex-wrap gap-2">
              <button
                @click="$emit('lihat-detail', petani)"
                class="px-3 py-1 bg-blue-500 text-white text-sm rounded-lg hover:bg-blue-600 transition"
              >
                Detail
              </button>
              <button
                @click="$emit('edit', petani)"
                class="px-3 py-1 bg-yellow-500 text-white text-sm rounded-lg hover:bg-yellow-600 transition"
              >
                Edit
              </button>
              <button
                @click="$emit('hapus', petani.id)"
                class="px-3 py-1 bg-red-600 text-white text-sm rounded-lg hover:bg-red-700 transition"
              >
                Hapus
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
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

<template>
  <!-- Background overlay -->
  <div class="fixed inset-0 z-50 flex items-center justify-center bg-black/40 backdrop-blur-sm">
    <div class="bg-white rounded-xl shadow-xl w-full max-w-lg p-6 space-y-6 relative">
      
      <!-- Header Modal -->
      <header class="flex items-center justify-between border-b pb-3">
        <h2 class="text-xl font-semibold text-gray-800">
          {{ editData ? 'Edit Petani' : 'Tambah Petani' }}
        </h2>
        <button
          @click="$emit('tutup')"
          class="text-gray-500 hover:text-gray-700 transition"
        >
          ✕
        </button>
      </header>

      <!-- Form -->
      <form @submit.prevent="kirimData" class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Nama Petani</label>
          <input
            v-model="form.nama"
            type="text"
            placeholder="Masukkan nama petani"
            class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:ring-2 focus:ring-green-500 focus:border-green-500 outline-none"
            required
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Alamat</label>
          <input
            v-model="form.alamat"
            type="text"
            placeholder="Masukkan alamat"
            class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:ring-2 focus:ring-green-500 focus:border-green-500 outline-none"
            required
          />
        </div>

        <!-- Tombol aksi -->
        <div class="flex justify-end space-x-3 pt-4 border-t">
          <button
            type="button"
            @click="$emit('tutup')"
            class="px-4 py-2 rounded-lg border border-gray-300 text-gray-600 hover:bg-gray-100 transition"
          >
            Batal
          </button>
          <button
            type="submit"
            class="px-4 py-2 rounded-lg bg-green-600 text-white hover:bg-green-700 transition"
          >
            {{ editData ? 'Simpan Perubahan' : 'Tambah Data' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, watch, reactive } from 'vue'

const props = defineProps({
  editData: {
    type: Object,
    default: null
  }
})

const emit = defineEmits(['tutup', 'simpan'])

const form = reactive({
  nama: '',
  alamat: ''
})

// isi form otomatis kalau mode edit
watch(
  () => props.editData,
  (val) => {
    if (val) {
      form.nama = val.nama
      form.alamat = val.alamat
    } else {
      form.nama = ''
      form.alamat = ''
    }
  },
  { immediate: true }
)

// kirim data ke parent
function kirimData() {
  emit('simpan', { ...form })
}
</script>

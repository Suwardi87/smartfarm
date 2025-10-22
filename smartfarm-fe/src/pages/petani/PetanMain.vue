<template>
  <div class="space-y-6">
    <!-- Header -->
    <header class="flex items-center justify-between py-8 my-4">
      <h1 class="text-2xl font-semibold text-gray-800">Manajemen Petani</h1>
      <button 
        @click="bukaModal"
        class="px-4 py-2  bg-green-600 text-white rounded-lg hover:bg-green-700"
      >
        Tambah Petani
      </button>
    </header>

    <!-- Ringkasan -->
    <PetaniRingkasan class="py-10"/>

    <!-- Tabel Daftar Petani -->
    <PetaniTabel
      :data="daftarPetani"
      :loading="loading"
      @lihat-detail="lihatDetail"
      @hapus="hapusPetani"
      @edit="editPetani"
    />

    <!-- Modal Form -->
    <PetaniFormModal
      v-if="modalTerbuka"
      :edit-data="petaniDipilih"
      @tutup="tutupModal"
      @simpan="simpanPetani"
    />

    <!-- Detail Petani -->
    <PetaniDetail
      v-if="petaniDetail"
      :data="petaniDetail"
      @tutup="tutupDetail"
    />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

// import composable CRUD
import { usePetani } from '@/composables/usePetani'

// import komponen UI
import PetaniTabel from './components/PetaniTabel.vue'
import PetaniFormModal from './components/PetaniFormModal.vue'
import PetaniRingkasan from './components/PetaniRingkasan.vue'
import PetaniDetail from './components/PetaniDetail.vue'

// state UI lokal
const modalTerbuka = ref(false)
const petaniDipilih = ref(null)
const petaniDetail = ref(null)

// ambil semua fungsi dari composable
const {
  daftarPetani,
  loading,
  ambilSemuaPetani,
  tambahPetani,
  ubahPetani,
  hapusPetani,
} = usePetani()

// lifecycle
onMounted(() => {
  ambilSemuaPetani()
})

// fungsi UI
function bukaModal() {
  petaniDipilih.value = null
  modalTerbuka.value = true
}

function tutupModal() {
  modalTerbuka.value = false
}

async function simpanPetani(data) {
  try {
    if (petaniDipilih.value) {
      await ubahPetani(petaniDipilih.value.id, data)
    } else {
      await tambahPetani(data)
    }
    tutupModal()
  } catch (err) {
    console.error('Gagal menyimpan petani:', err)
  }
}

function editPetani(petani) {
  petaniDipilih.value = petani
  modalTerbuka.value = true
}

function lihatDetail(petani) {
  petaniDetail.value = petani
}

function tutupDetail() {
  petaniDetail.value = null
}
</script>

<style scoped>
/* styling opsional */
</style>

<script setup lang="ts">
import { onMounted, ref, computed } from "vue"
import { useRoute, useRouter } from "vue-router"
import axios from "axios"
import { RotateCcw, Trash2 } from "lucide-vue-next"
import { Button } from "@/components/ui/button"
import AdminMainLayout from "@/components/layout/AdminMainLayout.vue"

// router composable
const route = useRoute()
const router = useRouter()

// reactive state
const berita = ref<any>(null)
const loading = ref(true)
const error = ref<string | null>(null)

// format helper
const fmt = (v: any, placeholder = "—") =>
  v === null || v === undefined || v === "" ? placeholder : v

// computed: info status berita
const statusInfo = computed(() => {
  const s = berita.value?.status?.toString().toLowerCase()
  const isPublish = s === "publish" || s === "published"
  return {
    label: isPublish ? "Publish" : "Draft",
    desc: isPublish ? "Dipublikasikan" : "Draft (belum dipublikasikan)",
    badgeClass: isPublish
      ? "bg-green-100 text-green-700 dark:bg-green-900/40 dark:text-green-300"
      : "bg-yellow-100 text-yellow-700 dark:bg-yellow-900/40 dark:text-yellow-300",
  }
})

// get berita detail by id
async function fetchBerita() {
  try {
    const id = route.params.id
    const { data } = await axios.get(`http://localhost:8083/api/berita/${id}`)
    berita.value = data
  } catch (err: any) {
    error.value = err.response?.data?.error || err.message
  } finally {
    loading.value = false
  }
}

// restore berita (PATCH)
async function restoreBerita(id: string) {
  if (!confirm("Yakin ingin memulihkan berita ini?")) return
  await axios.patch(`http://localhost:8083/api/berita/${id}/restore`)
  alert("✅ Berita berhasil dipulihkan")
  router.push("/berita")
}

// force delete berita (DELETE)
async function forceDeleteBerita(id: string) {
  if (!confirm("Hapus permanen berita ini? Tindakan ini tidak bisa dibatalkan!")) return
  await axios.delete(`http://localhost:8083/api/berita/${id}/force`)
  alert("🗑️ Berita berhasil dihapus permanen")
  router.push("/berita")
}

// load saat mount
onMounted(fetchBerita)
</script>

<template>
  <AdminMainLayout>
    <div class="flex justify-center px-4 py-10 md:px-6">
      <div
        class="w-full max-w-4xl rounded-2xl border border-gray-200 bg-white p-6 shadow-xl dark:border-gray-700 dark:bg-gray-800"
      >
        <!-- Loading -->
        <div v-if="loading" class="text-center py-10 text-gray-500 dark:text-gray-300">
          Memuat data berita...
        </div>

        <!-- Error -->
        <div v-else-if="error" class="text-center text-red-600 dark:text-red-400">
          {{ error }}
        </div>

        <!-- Berita detail -->
        <template v-else>
          <h1 class="mb-6 text-center text-3xl font-bold text-gray-900 dark:text-white">
            {{ berita.judul }}
          </h1>

          <!-- Label -->
          <div class="mb-6 flex flex-wrap justify-center gap-2">
            <span
              class="rounded-full bg-gray-100 px-3 py-1 text-xs font-semibold text-gray-700 dark:bg-gray-700 dark:text-gray-200"
            >
              Kategori: {{ fmt(berita.kategori, "Tidak ada") }}
            </span>
            <span
              :class="['rounded-full px-3 py-1 text-xs font-semibold', statusInfo.badgeClass]"
            >
              {{ statusInfo.label }}
            </span>
            <span
              v-if="berita.deleted_at"
              class="rounded-full bg-red-100 px-3 py-1 text-xs font-semibold text-red-700 dark:bg-red-900/40 dark:text-red-300"
            >
              Terhapus (soft delete)
            </span>
          </div>

          <!-- Thumbnail -->
          <div v-if="berita.thumbnail" class="mb-6 flex justify-center">
            <img
              :src="`/storage/${berita.thumbnail}`"
              :alt="`Thumbnail berita: ${berita.judul}`"
              class="max-h-96 rounded-xl object-cover shadow-md"
            />
          </div>

          <!-- Info Grid -->
          <h2 class="mb-3 text-lg font-semibold text-gray-900 dark:text-gray-200">Informasi</h2>
          <div class="mb-6 grid gap-6 md:grid-cols-2">
            <div class="space-y-2">
              <p><strong>Slug:</strong> {{ fmt(berita.slug, "Tidak ada") }}</p>
              <p><strong>Deskripsi:</strong> {{ fmt(berita.deskripsi, "Tidak ada") }}</p>
              <p><strong>Kategori:</strong> {{ fmt(berita.kategori, "Tidak ada") }}</p>
            </div>
            <div class="space-y-2">
              <p>
                <strong>Status:</strong>
                <span
                  :class="['rounded-full px-3 py-1 text-xs font-semibold', statusInfo.badgeClass]"
                >
                  {{ statusInfo.label }}
                </span>
                <span class="ml-2 text-sm text-gray-500 dark:text-gray-400">
                  — {{ statusInfo.desc }}
                </span>
              </p>
              <p>
                <strong>Tanggal Publish:</strong>
                {{
                  new Intl.DateTimeFormat("id-ID", {
                    year: "numeric",
                    month: "long",
                    day: "2-digit",
                  }).format(new Date(berita.tanggal_publish))
                }}
              </p>
              <p><strong>Dibuat:</strong> {{ fmt(berita.created_at, "—") }}</p>
              <p><strong>Diperbarui:</strong> {{ fmt(berita.updated_at, "—") }}</p>
              <p><strong>Dihapus:</strong> {{ fmt(berita.deleted_at, "—") }}</p>
            </div>
          </div>

          <!-- Deskripsi -->
          <p
            v-if="berita.deskripsi"
            class="mb-6 text-center text-lg text-gray-600 italic dark:text-gray-300"
          >
            “{{ berita.deskripsi }}”
          </p>

          <!-- Konten -->
          <h2 class="mb-3 text-lg font-semibold text-gray-900 dark:text-gray-200">Konten</h2>
          <div
            v-if="berita.konten"
            class="prose dark:prose-invert max-w-none leading-relaxed text-gray-800 dark:text-gray-200"
            v-html="berita.konten"
          ></div>
          <div v-else class="text-gray-500 italic dark:text-gray-400">Konten belum tersedia.</div>

          <!-- Tombol Aksi -->
          <div
            class="mt-10 flex justify-center gap-4 border-t border-gray-200 pt-6 dark:border-gray-700"
          >
            <Button @click="router.push('/berita')" class="bg-gray-600 hover:bg-gray-700"
              >Kembali</Button
            >

            <template v-if="berita.deleted_at">
              <Button
                class="bg-green-600 hover:bg-green-700"
                @click="restoreBerita(berita.id)"
              >
                <RotateCcw class="w-4 h-4 mr-2" /> Pulihkan
              </Button>

              <Button
                class="bg-red-600 hover:bg-red-700"
                @click="forceDeleteBerita(berita.id)"
              >
                <Trash2 class="w-4 h-4 mr-2" /> Hapus Permanen
              </Button>
            </template>
          </div>
        </template>
      </div>
    </div>
  </AdminMainLayout>
</template>

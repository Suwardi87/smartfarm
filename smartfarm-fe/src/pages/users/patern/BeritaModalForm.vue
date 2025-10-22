<script setup lang="ts">
import { useForm } from '@inertiajs/vue3'
import { ref, watch } from 'vue'
import { useCrudForm } from '@/composables/useCrudForm'
import { useSlug } from '@/composables/useSlug'

// === Props & Emits ===
const props = defineProps({
  show: Boolean,
  isEdit: Boolean,
  item: Object,
})
const emit = defineEmits(['close'])

// === Form data ===
const form = useForm({
  judul: '',
  slug: '',
  deskripsi: '',
  kategori: '',
  konten: '',
  thumbnail: null,
  tanggal_publish: '',
  status: 'draft',
})

// === Preview thumbnail ===
const previewThumbnail = ref(null)

// === Reset form ketika modal dibuka ===
watch(
  () => props.show,
  (show) => {
    if (!show) return
    form.clearErrors()
    if (props.isEdit && props.item) {
      form.reset()
      Object.assign(form, {
        judul: props.item.judul || '',
        slug: props.item.slug || '',
        deskripsi: props.item.deskripsi || '',
        kategori: props.item.kategori || '',
        konten: props.item.konten || '',
        thumbnail: null,
        tanggal_publish: props.item.tanggal_publish
          ? new Date(props.item.tanggal_publish).toISOString().slice(0, 16)
          : '',
        status: props.item.status || 'draft',
      })
      previewThumbnail.value = props.item.thumbnail
        ? `/storage/${props.item.thumbnail}`
        : null
    } else {
      form.reset()
      previewThumbnail.value = null
    }
  },
  { immediate: true }
)

// === Slug otomatis ===
useSlug(form)

// === Upload thumbnail ===
function handleFileUpload(e) {
  const file = e.target.files[0]
  form.thumbnail = file
  if (file) previewThumbnail.value = URL.createObjectURL(file)
}

// === Gunakan useCrudForm ===
const { submit } = useCrudForm(form, props, 'admin.berita', () => emit('close'))
</script>


<template>
    <div
        v-if="props.show"
        class="fixed inset-0 z-50 flex items-center justify-center bg-black/50 p-4"
    >
        <div
            class="relative max-h-[90vh] w-full max-w-2xl overflow-y-auto rounded-2xl bg-white p-6 shadow-xl dark:bg-gray-800"
        >
            <button
                @click="emit('close')"
                class="absolute top-3 right-3 text-xl font-bold text-gray-400 hover:text-gray-600"
            >
                ✕
            </button>

            <h3 class="mb-4 text-lg font-semibold">
                {{ props.isEdit ? 'Edit Berita' : 'Tambah Berita' }}
            </h3>

            <!-- ✅ Global validation block -->
            <div
                v-if="form.hasErrors"
                class="mb-4 rounded border border-red-400 bg-red-100 p-3 text-red-700"
            >
                <h4 class="font-bold">Terjadi Kesalahan:</h4>
                <ul class="list-disc pl-5">
                    <li
                        v-for="(messages, field) in form.errors"
                        :key="field"
                        class="text-sm"
                    >
                        {{ messages[0] }}
                    </li>
                </ul>
            </div>

            <!-- === FORM === -->
            <form @submit.prevent="submit" class="space-y-4">
                <!-- Judul -->
                <div>
                    <label class="mb-2 block text-sm font-medium">
                        Judul <span class="text-red-500">*</span>
                    </label>
                    <input
                        v-model="form.judul"
                        type="text"
                        placeholder="Masukkan judul berita"
                        class="w-full rounded-lg border px-3 py-2"
                        :class="form.errors.judul ? 'border-red-500' : ''"
                    />
                    <span
                        v-if="form.errors.judul"
                        class="text-sm text-red-500"
                    >
                        {{ form.errors.judul[0] }}
                    </span>
                </div>

                <!-- Slug -->
                <div>
                    <label class="mb-2 block text-sm font-medium">
                        Slug <span class="text-red-500">*</span>
                    </label>
                    <input
                        v-model="form.slug"
                        type="text"
                        placeholder="slug-otomatis"
                        class="w-full rounded-lg border px-3 py-2"
                        :class="form.errors.slug ? 'border-red-500' : ''"
                    />
                    <span
                        v-if="form.errors.slug"
                        class="text-sm text-red-500"
                    >
                        {{ form.errors.slug[0] }}
                    </span>
                </div>

                <!-- Deskripsi -->
                <div>
                    <label class="mb-2 block text-sm font-medium">Deskripsi</label>
                    <textarea
                        v-model="form.deskripsi"
                        rows="3"
                        class="w-full rounded-lg border px-3 py-2"
                        :class="form.errors.deskripsi ? 'border-red-500' : ''"
                    ></textarea>
                    <span
                        v-if="form.errors.deskripsi"
                        class="text-sm text-red-500"
                    >
                        {{ form.errors.deskripsi[0] }}
                    </span>
                </div>

                <!-- Konten -->
                <div>
                    <label class="mb-2 block text-sm font-medium">
                        Konten <span class="text-red-500">*</span>
                    </label>
                    <textarea
                        v-model="form.konten"
                        rows="6"
                        class="w-full rounded-lg border px-3 py-2"
                        :class="form.errors.konten ? 'border-red-500' : ''"
                    ></textarea>
                    <span
                        v-if="form.errors.konten"
                        class="text-sm text-red-500"
                    >
                        {{ form.errors.konten[0] }}
                    </span>
                </div>

                <!-- Kategori -->
                <div>
                    <label class="mb-2 block text-sm font-medium">Kategori</label>
                    <input
                        v-model="form.kategori"
                        type="text"
                        class="w-full rounded-lg border px-3 py-2"
                        :class="form.errors.kategori ? 'border-red-500' : ''"
                    />
                    <span
                        v-if="form.errors.kategori"
                        class="text-sm text-red-500"
                    >
                        {{ form.errors.kategori[0] }}
                    </span>
                </div>

                <!-- Thumbnail -->
                <div>
                    <label class="mb-2 block text-sm font-medium">Thumbnail</label>
                    <input
                        type="file"
                        @change="handleFileUpload"
                        accept="image/*"
                        class="w-full rounded-lg border px-3 py-2"
                        :class="form.errors.thumbnail ? 'border-red-500' : ''"
                    />
                    <span
                        v-if="form.errors.thumbnail"
                        class="text-sm text-red-500"
                    >
                        {{ form.errors.thumbnail[0] }}
                    </span>

                    <!-- ✅ Preview -->
                    <div v-if="previewThumbnail" class="mt-3">
                        <p class="mb-1 text-sm text-gray-600">
                            Preview Thumbnail:
                        </p>
                        <img
                            :src="previewThumbnail"
                            alt="Preview Thumbnail"
                            class="h-28 w-40 rounded-lg border object-cover"
                        />
                    </div>
                </div>

                <!-- Tanggal Publish -->
                <div>
                    <label class="mb-2 block text-sm font-medium">
                        Tanggal Publish
                    </label>
                    <input
                        v-model="form.tanggal_publish"
                        type="datetime-local"
                        class="w-full rounded-lg border px-3 py-2"
                        :class="form.errors.tanggal_publish ? 'border-red-500' : ''"
                    />
                    <span
                        v-if="form.errors.tanggal_publish"
                        class="text-sm text-red-500"
                    >
                        {{ form.errors.tanggal_publish[0] }}
                    </span>
                </div>

                <!-- Status -->
                <div>
                    <label class="mb-2 block text-sm font-medium">
                        Status <span class="text-red-500">*</span>
                    </label>
                    <select
                        v-model="form.status"
                        class="w-full rounded-lg border px-3 py-2"
                        :class="form.errors.status ? 'border-red-500' : ''"
                    >
                        <option value="draft">Draft</option>
                        <option value="publish">Publish</option>
                    </select>
                    <span
                        v-if="form.errors.status"
                        class="text-sm text-red-500"
                    >
                        {{ form.errors.status[0] }}
                    </span>
                </div>

                <!-- Tombol -->
                <div class="flex justify-end gap-3 border-t pt-4">
                    <button
                        type="button"
                        @click="emit('close')"
                        class="rounded-lg border px-4 py-2 hover:bg-gray-50"
                        :disabled="form.processing"
                    >
                        Batal
                    </button>
                    <button
                        type="submit"
                        :disabled="form.processing"
                        class="rounded-lg bg-blue-600 px-4 py-2 text-white hover:bg-blue-700 disabled:opacity-50"
                    >
                        {{
                            form.processing
                                ? props.isEdit
                                    ? 'Mengupdate...'
                                    : 'Menyimpan...'
                                : props.isEdit
                                ? 'Update'
                                : 'Simpan'
                        }}
                    </button>
                </div>
            </form>
        </div>
    </div>
</template>

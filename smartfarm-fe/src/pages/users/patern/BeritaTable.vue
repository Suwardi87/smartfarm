<script setup lang="ts">
import { Link } from '@inertiajs/vue3';
import { Eye, Pencil, Trash2 } from 'lucide-vue-next';

const props = defineProps({
  items: Array,
});
const emit = defineEmits(['edit']);

const fmt = (v, placeholder = '—') => (v ? v : placeholder);

const statusLabel = (s) => {
  const x = (s || '').toLowerCase();
  return x === 'publish' ? 'Publish' : 'Draft';
};
const statusClass = (s) => {
  const isPublish = (s || '').toLowerCase() === 'publish';
  return isPublish
    ? 'bg-green-100 text-green-700'
    : 'bg-yellow-100 text-yellow-700';
};
</script>

<template>
  <!-- Desktop: Table -->
  <div class="hidden overflow-x-auto rounded-2xl border bg-white shadow md:block">
    <table class="w-full text-left text-sm">
      <thead class="bg-gray-100">
        <tr>
          <th class="px-6 py-3">Judul</th>
          <th class="px-6 py-3">Slug</th>
          <th class="px-6 py-3">Kategori</th>
          <th class="px-6 py-3">Tanggal Publish</th>
          <th class="px-6 py-3">Status</th>
          <th class="px-6 py-3 text-right">Aksi</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in items" :key="item.id" class="border-t hover:bg-gray-50">
          <td class="px-6 py-4">{{ fmt(item.judul, 'Tanpa judul') }}</td>
          <td class="px-6 py-4">{{ fmt(item.slug, 'Tidak ada') }}</td>
          <td class="px-6 py-4">{{ fmt(item.kategori, 'Tidak ada') }}</td>
          <td class="px-6 py-4">
            {{
              item.status?.toLowerCase() === 'publish'
                ? new Intl.DateTimeFormat('id-ID', { year: 'numeric', month: 'long', day: 'numeric' }).format(new Date(item.tanggal_publish))
                : 'Belum dipublikasikan'
            }}
          </td>
          <td class="px-6 py-4">
            <span
              :class="[
                'rounded-full px-2 py-1 text-xs font-semibold',
                statusClass(item.status),
              ]"
            >
              {{ statusLabel(item.status) }}
            </span>
          </td>
          <td class="flex justify-end gap-2 px-6 py-4 text-right">
            <button
              v-if="!item.deleted_at"
              @click="emit('edit', item)"
              class="rounded-lg bg-blue-50 p-2 text-blue-600"
            >
              <Pencil class="h-4 w-4" />
            </button>
            <Link
              v-if="!item.deleted_at"
              :href="route('admin.berita.destroy', item.id)"
              method="delete"
              as="button"
              class="rounded-lg bg-red-50 p-2 text-red-600"
            >
              <Trash2 class="h-4 w-4" />
            </Link>

            <Link
              :href="route('admin.berita.show', { id: item.id })"
              class="rounded-lg bg-gray-50 p-2 text-gray-600"
            >
              <Eye class="h-4 w-4" />
            </Link>
          </td>
        </tr>
        <tr v-if="items.length === 0">
          <td
            colspan="6"
            class="px-6 py-10 text-center text-sm text-gray-500"
          >
            Tidak ada data berita.
          </td>
        </tr>
      </tbody>
    </table>
  </div>

  <!-- Mobile: Card List -->
  <div class="space-y-4 md:hidden">
    <div
      v-for="item in items"
      :key="item.id"
      class="rounded-xl border bg-white p-4 shadow"
    >
      <h3 class="font-semibold text-gray-800">
        {{ fmt(item.judul, 'Tanpa judul') }}
      </h3>
      <p class="text-xs text-gray-500 mb-2">Slug: {{ fmt(item.slug, '-') }}</p>
      <p class="text-sm text-gray-600">
        Kategori: {{ fmt(item.kategori, 'Tidak ada') }}
      </p>
      <p class="text-sm text-gray-600">
        Publish: {{
          item.status?.toLowerCase() === 'publish'
            ? fmt(item.tanggal_publish, 'Belum dipublikasikan')
            : 'Belum dipublikasikan'
        }}
      </p>
      <span
        :class="[
          'mt-2 inline-block rounded-full px-2 py-1 text-xs font-semibold',
          statusClass(item.status),
        ]"
      >
        {{ statusLabel(item.status) }}
      </span>

      <!-- Action Menu (di bawah) -->
      <div class="mt-4 flex justify-end gap-2 border-t pt-3">
        <button
          v-if="!item.deleted_at"
          @click="emit('edit', item)"
          class="rounded-lg bg-blue-50 p-2 text-blue-600"
        >
          <Pencil class="h-4 w-4" />
        </button>
        <Link
          v-if="!item.deleted_at"
          :href="route('admin.berita.destroy', item.id)"
          method="delete"
          as="button"
          class="rounded-lg bg-red-50 p-2 text-red-600"
        >
          <Trash2 class="h-4 w-4" />
        </Link>
        <Link
          :href="route('admin.berita.show', { id: item.id })"
          class="rounded-lg bg-gray-50 p-2 text-gray-600"
        >
          <Eye class="h-4 w-4" />
        </Link>
      </div>
    </div>

    <div v-if="items.length === 0" class="text-center text-sm text-gray-500">
      Tidak ada data berita.
    </div>
  </div>
</template>

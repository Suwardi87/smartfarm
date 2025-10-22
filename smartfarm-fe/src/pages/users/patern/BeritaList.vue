<script setup lang="ts">
import { Link } from '@inertiajs/vue3';
import { Eye, Pencil, Trash2 } from 'lucide-vue-next';

const props = defineProps({
    items: Array,
});
const emit = defineEmits(['edit']);
</script>

<template>
  <div class="space-y-3 md:hidden">
    <div v-for="item in items" :key="item.id"
      class="rounded-2xl border bg-white p-4 shadow">
      <div class="flex justify-between">
        <div>
          <h3 class="font-semibold">{{ item.judul || 'Tanpa judul' }}</h3>
          <p class="text-xs text-gray-500">Slug: {{ item.slug || 'Tidak ada' }}</p>
          <p class="text-xs text-gray-500">Publish: {{ item.tanggal_publish || 'Belum dipublikasikan' }}</p>
        </div>
        <div class="flex gap-2">
          <Link :href="route('admin.berita.show',{id:item.id})"
            class="rounded-lg bg-gray-50 p-2 text-gray-600">
            <Eye class="h-4 w-4"/>
          </Link>
          <button @click="emit('edit', item)" class="rounded-lg bg-blue-50 p-2 text-blue-600">
            <Pencil class="h-4 w-4"/>
          </button>
          <Link :href="route('admin.berita.destroy',{id:item.id})"
            method="delete" as="button" class="rounded-lg bg-red-50 p-2 text-red-600">
            <Trash2 class="h-4 w-4"/>
          </Link>
        </div>
      </div>
    </div>
    <div v-if="items.length === 0"
      class="rounded-2xl border-dashed border p-6 text-center text-sm text-gray-500">
      Tidak ada data berita.
    </div>
  </div>
</template>

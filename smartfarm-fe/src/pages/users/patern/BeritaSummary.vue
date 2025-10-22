<script setup lang="ts">
const props = defineProps({
    summary: Object,
});
</script>

<template>
  <div class="mb-4 flex flex-wrap gap-2 text-sm">
    <span class="rounded-full bg-gray-100 px-3 py-1">Total: {{ summary.total }}</span>
    <span class="rounded-full bg-green-100 px-3 py-1">Publish: {{ summary.publish }}</span>
    <span class="rounded-full bg-yellow-100 px-3 py-1">Draft: {{ summary.draft }}</span>
    <span class="rounded-full bg-red-100 px-3 py-1">Terhapus: {{ summary.deleted }}</span>
  </div>
</template>

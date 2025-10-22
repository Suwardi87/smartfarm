<template>
  <component
    :is="asChild ? 'span' : 'button'"
    class="flex items-center gap-3 rounded-md px-3 py-2 text-sm text-gray-700 hover:bg-gray-100 transition"
  >
    <slot />
  </component>
</template>

<script setup lang="ts">
defineProps({
  asChild: { type: Boolean, default: false },
})
</script>

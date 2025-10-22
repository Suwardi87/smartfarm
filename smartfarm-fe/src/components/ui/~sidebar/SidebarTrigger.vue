<template>
  <button
    @click="toggle"
    class="p-2 rounded-md hover:bg-gray-100 transition"
    :title="isOpen ? 'Tutup Sidebar' : 'Buka Sidebar'"
  >
    <!-- ❌ (X) saat open -->
    <svg
      v-if="isOpen"
      xmlns="http://www.w3.org/2000/svg"
      class="w-5 h-5"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
    </svg>

    <!-- ☰ (Menu) saat closed -->
    <svg
      v-else
      xmlns="http://www.w3.org/2000/svg"
      class="w-5 h-5"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h8m-8 6h16" />
    </svg>
  </button>
</template>

<script setup>
import { useSidebar } from './useSidebar'
const { isOpen, toggle } = useSidebar()
</script>

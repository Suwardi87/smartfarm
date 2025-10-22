<template>
  <header class="flex items-center justify-between px-4 py-3 border-b border-gray-200">
    <slot>
      <h2 class="font-semibold text-lg truncate">SmartFarm</h2>
    </slot>

    <!-- 🧠 Tombol Toggle harus di sini -->
    <SidebarTrigger />
  </header>
</template>

<script setup>
import SidebarTrigger from './SidebarTrigger.vue';

</script>

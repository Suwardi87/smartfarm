<script setup>
import AppSidebar from '@/components/layout/AppSidebar.vue'
import AppHeader from '@/components/layout/AppHeader.vue'
</script>

<template>
  <!-- <div class="flex h-screen"> -->
    <!-- 🧱 Sidebar di kiri -->
    <!-- <AppSidebar /> -->

    <!-- 🧩 Area kanan: header + konten -->
    <!-- <div class="flex-1 flex flex-col"> -->
      <!-- Header di atas -->
      <!-- <AppHeader /> -->

      <!-- Konten utama -->
      <!-- <main class="flex-1 p-6 bg-gray-50"> -->
        <!-- Tempat RouterView atau konten halaman -->
        <RouterView />
        <!-- <h1 class="text-2xl font-bold">Selamat datang di SmartFarm 🌿</h1> -->
      <!-- </main> -->
    <!-- </div> -->
  <!-- </div> -->
</template>

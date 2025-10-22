<script setup>
import AppSidebar from '@/components/layout/AppSidebar.vue'
import AppHeader from '@/components/layout/AppHeader.vue'
</script>

<template>
  <div class="flex min-h-screen">
    <AppSidebar />
    <div class="flex-1 flex flex-col">
      <AppHeader />
      <main class="p-6 flex-1 bg-gray-50">
        <RouterView />
      </main>
    </div>
  </div>
</template>

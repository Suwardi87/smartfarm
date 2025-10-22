<template>
  <aside
    :class="[
      'flex flex-col h-screen bg-white border-r border-gray-200 shadow transition-all duration-300 ease-in-out',
      isOpen ? 'w-64' : 'w-17',
    ]"
  >
    <slot />
  </aside>
</template>

<script setup lang="ts">
import { provideSidebar } from './useSidebar'

interface SidebarProps {
  side?: 'left' | 'right'
}
const props = withDefaults(defineProps<SidebarProps>(), { side: 'left' })

// 🧠 provide state to children
const { isOpen } = provideSidebar()
</script>

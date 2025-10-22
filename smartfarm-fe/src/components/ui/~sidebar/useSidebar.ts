import { ref, provide, inject } from "vue"

const SIDEBAR_KEY = Symbol("sidebar")

export function provideSidebar() {
  const isOpen = ref(true)
  const toggle = () => (isOpen.value = !isOpen.value)
  provide(SIDEBAR_KEY, { isOpen, toggle })
  return { isOpen, toggle }
}

export function useSidebar() {
  const ctx = inject<{ isOpen: any; toggle: () => void }>(SIDEBAR_KEY)
  if (!ctx) throw new Error("useSidebar() must be used within Sidebar")
  return ctx
}

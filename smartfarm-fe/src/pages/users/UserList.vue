<script setup lang="ts">
import { ref, onMounted } from "vue"
import axios from "axios"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import {
  Table,
  TableHeader,
  TableHead,
  TableRow,
  TableBody,
  TableCell,
} from "@/components/ui/table"
import {
  DropdownMenu,
  DropdownMenuTrigger,
  DropdownMenuContent,
  DropdownMenuItem,
} from "@/components/ui/dropdown-menu"
import { MoreHorizontal } from "lucide-vue-next"

// 🔹 State
const users = ref<any[]>([])
const filteredUsers = ref<any[]>([])
const search = ref("")
const loading = ref(false)
const error = ref<string | null>(null)

// 🔹 Fetch data dari backend
async function fetchUsers() {
  loading.value = true
  error.value = null
  try {
    const res = await axios.get("http://localhost:8083/api/auth/users")
    users.value = res.data
    filteredUsers.value = users.value
  } catch (err: any) {
    error.value = err.response?.data?.error || err.message
  } finally {
    loading.value = false
  }
}

// 🔹 Filter data
function filterUsers() {
  const q = search.value.toLowerCase()
  filteredUsers.value = users.value.filter(
    (u) =>
      u.nama_lengkap?.toLowerCase().includes(q) ||
      u.username?.toLowerCase().includes(q)
  )
}

// 🔹 Action
function addUser() {
  alert("Fitur tambah user (nanti modal form)")
}

function editUser(id: string) {
  alert("Edit user ID: " + id)
}

async function deleteUser(id: string) {
  if (confirm("Yakin hapus user ini?")) {
    try {
      await axios.delete(`http://localhost:8083/api/auth/users/${id}`)
      await fetchUsers()
      alert("✅ User berhasil dihapus")
    } catch (err: any) {
      alert("❌ Gagal menghapus user: " + (err.response?.data?.error || err.message))
    }
  }
}

// 🔹 Lifecycle
onMounted(fetchUsers)
</script>

<template>
  <div class="space-y-6">
    <!-- 🔹 Header Page -->
    <div class="flex justify-between items-center">
      <div>
        <h1 class="text-2xl font-semibold tracking-tight">Daftar Pengguna</h1>
        <p class="text-sm text-muted-foreground">Kelola akun pengguna sistem SmartFarm.</p>
      </div>
      <Button @click="addUser">+ Tambah Pengguna</Button>
    </div>

    <!-- 🔍 Toolbar -->
    <div class="flex items-center gap-3 ">
      <Input
        v-model="search"
        placeholder="Cari pengguna..."
        class="max-w-sm "
        @keyup.enter="filterUsers"
      />
      <Button variant="secondary" @click="filterUsers">Cari</Button>
    </div>

    <!-- 📋 State loading / error -->
    <div v-if="loading" class="text-gray-500 text-center py-6">Memuat data pengguna...</div>
    <div v-else-if="error" class="text-red-600 text-center py-6">{{ error }}</div>

    <!-- 📋 Tabel Data -->
    <div v-else class="rounded-md border bg-white dark:bg-gray-800">
      <Table >
        <TableHeader>
          <TableRow>
            <TableHead>Nama Lengkap</TableHead>
            <TableHead>Username</TableHead>
            <TableHead>Level</TableHead>
            <TableHead class="text-right w-[100px]">Aksi</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <TableRow v-for="user in filteredUsers" :key="user.id">
            <TableCell class="font-medium">{{ user.nama_lengkap }}</TableCell>
            <TableCell>{{ user.username }}</TableCell>
            <TableCell>{{ user.level }}</TableCell>
            <TableCell class="text-right">
              <DropdownMenu>
                <DropdownMenuTrigger as-child>
                  <Button variant="ghost" size="icon">
                    <MoreHorizontal class="h-4 w-4" />
                  </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent align="end">
                  <DropdownMenuItem @click="editUser(user.id)">Edit</DropdownMenuItem>
                  <DropdownMenuItem
                    class="text-red-600"
                    @click="deleteUser(user.id)"
                  >
                    Hapus
                  </DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>
            </TableCell>
          </TableRow>
          <TableRow v-if="filteredUsers.length === 0">
            <TableCell colspan="4" class="text-center text-sm text-gray-500">
              Tidak ada pengguna ditemukan.
            </TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </div>
  </div>
</template>

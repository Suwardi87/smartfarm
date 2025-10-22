import { createRouter, createWebHistory } from "vue-router"
import AuthLayout from "@/components/layout/AuthLayout.vue"
import DefaultLayout from "@/components/layout/DefaultLayout.vue"
import LoginPage from "@/pages/auth/Login.vue"
import Dashboard from "@/pages/dashboard/MainDashboard.vue"
import { useAuthStore } from "@/store/useAuthStore"

const routes = [
  {
    path: "/login",
    component: AuthLayout,
    children: [
      { path: "", name: "Login", component: LoginPage },
    ],
  },
  {
    path: "/",
    component: DefaultLayout,
    meta: { requiresAuth: true },
    children: [
      { path: "", name: "Dashboard", component: Dashboard },
      { path: "/user", name: "User", component: () => import("@/pages/users/UserList.vue") },
      { path: "/petani", name: "Petani", component: () => import("@/pages/petani/PetanMain.vue") },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach((to, from, next) => {
  const auth = useAuthStore()
  const token = auth.token || localStorage.getItem("token")
  if (to.meta.requiresAuth && !token) next("/login")
  else next()
})

export default router

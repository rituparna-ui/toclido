<template>
  <div>
    <nav class="flex justify-between items-center bg-[#393E46] p-4 text-white">
      <div class="text-2xl font-semibold">
        <NuxtLink to="/">TOCLIDO</NuxtLink>
      </div>
      <div class="flex gap-4" v-if="!authState.isAuth">
        <NuxtLink to="/auth/login">Login</NuxtLink>
        <NuxtLink to="/auth/signup">Signup</NuxtLink>
      </div>
      <div class="flex gap-4" v-else>
        <Button variant="secondary" @click="logout">Logout</Button>
      </div>
    </nav>
  </div>
</template>

<script setup>
import { useAuth } from "@/composables/useAuth";
import { Button } from "@/components/ui/button";

const { authState, logout } = useAuth();
</script>
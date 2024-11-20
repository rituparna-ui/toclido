<template>
  <div class="h-full flex flex-col">
    <Navbar />
    <NuxtPage />
    <Toaster position="top-right" />
  </div>
</template>

<script setup>
import { Toaster } from "@/components/ui/sonner";
import { useAuth } from "@/composables/useAuth";
</script>

<style>
html,
body,
#__nuxt {
  @apply h-full bg-[#222831] text-white;
}
</style>
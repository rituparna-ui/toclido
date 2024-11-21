<template>
  <div class="p-8">
    <h1 class="text-center text-2xl">Settings</h1>
    <div class="w-full lg:w-1/2 mx-auto"></div>
  </div>
</template>

<script setup>
definePageMeta({
  middleware: ["auth"],
});
</script>
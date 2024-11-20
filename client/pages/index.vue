<template>
  <div class="p-8">
    <h1>ToDos</h1>
  </div>
</template>

<script setup >
definePageMeta({
  middleware: ["auth"],
});
</script>

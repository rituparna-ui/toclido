<template>
  <div class="p-8">
    <h1 class="text-center text-2xl">My TODOs</h1>
    <div class="w-full lg:w-1/2 mx-auto">
      <TodosNew />
      <TodosItem v-for="todo in todos" :key="todo._id" :todo="todo" />
    </div>
  </div>
</template>

<script setup>
import { Badge } from "@/components/ui/badge";
const runtimeConfig = useRuntimeConfig();
const baseUrl = runtimeConfig.public.baseAPI;

definePageMeta({
  middleware: ["auth"],
});

const token = useCookie("token");
const { todos } = useTodos();
const { data } = await useFetch(baseUrl + "/todos", {
  method: "GET",
  headers: {
    Authorization: `Bearer ${token.value}`,
  },
});

todos.value = data.value.todos;
</script>

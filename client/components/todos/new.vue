<template>
  <div class="my-4">
    <form class="flex flex-col gap-4 text-black" @submit.prevent="createTodo">
      <Input type="text" placeholder="Title" required name="title" />
      <Input
        type="text"
        placeholder="Description"
        required
        name="description"
      />
      <Button type="submit">Add Todo</Button>
    </form>
  </div>
</template>

<script setup>
const runtimeConfig = useRuntimeConfig();
const baseUrl = runtimeConfig.public.baseAPI;

import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { toast } from "vue-sonner";

const { todos } = useTodos();

const token = useCookie("token");

const createTodo = async (e) => {
  const formData = new FormData(e.target);
  const data = Object.fromEntries(formData.entries());
  const { title, description } = data;
  try {
    const data = await $fetch(baseUrl + "/todos", {
      method: "POST",
      body: {
        title,
        description,
      },
      headers: {
        Authorization: `Bearer ${token.value}`,
      },
    });
    todos.value = [data.todo, ...todos.value];
    toast.success(data.message);
    e.target.reset();
  } catch (error) {
    toast.error(error.data.message);
  }
};
</script>
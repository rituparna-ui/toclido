<template>
  <div class="border rounded-md my-4 p-4">
    <div class="text-lg font-semibold flex justify-between items-center">
      <div :class="[todo.completed ? 'line-through text-yellow-400' : '']">
        {{ todo.title }}
      </div>
      <div>
        <Button variant="ghost" @click="toggleDesc">
          {{ descHidden ? "Show" : "Hide" }}
        </Button>
      </div>
    </div>
    <div v-if="!descHidden">
      <div class="border-y border-dotted py-2 my-2">
        {{ todo.description }}
      </div>
      <div class="flex justify-between items-center">
        <Button variant="secondary" @click="toggleStatus(todo._id)">
          {{ todo.completed ? "Mark Incomplete" : "Mark Completed" }}
        </Button>
        <Button variant="destructive" @click="deleteTodo(todo._id)">
          Delete
        </Button>
      </div>
    </div>
  </div>
</template>

<script setup>
const runtimeConfig = useRuntimeConfig();
const baseUrl = runtimeConfig.public.baseAPI;

import { Button } from "@/components/ui/button";
import { toast } from "vue-sonner";

const token = useCookie("token");
const { todos } = useTodos();

defineProps({
  todo: {
    title: String,
    description: String,
    completed: Boolean,
    _id: String,
    i: Number,
  },
});

const descHidden = ref(true);

const toggleDesc = () => {
  descHidden.value = !descHidden.value;
};

const toggleStatus = async (id) => {
  try {
    const data = await $fetch(baseUrl + "/todos/toggle-status/" + id, {
      method: "PATCH",
      headers: {
        Authorization: `Bearer ${token.value}`,
      },
    });
    toast.success(data.message);
    todos.value = todos.value.map((todo) => {
      if (todo._id === id) {
        todo.completed = !todo.completed;
      }
      return todo;
    });
  } catch (error) {
    toast.error(error.data.message);
  }
};

const deleteTodo = async (id) => {
  try {
    const data = await $fetch(baseUrl + "/todos/" + id, {
      method: "DELETE",
      headers: {
        Authorization: `Bearer ${token.value}`,
      },
    });
    toast.success(data.message);
    todos.value = todos.value.filter((todo) => todo._id !== id);
  } catch (error) {
    toast.error(error.data.message);
  }
};
</script>
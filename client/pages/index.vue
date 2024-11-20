<template>
  <div class="p-8">
    <h1 class="text-center text-2xl">My TODOs</h1>
    <div class="w-full lg:w-1/2 mx-auto">
      <Accordion type="single" collapsible>
        <AccordionItem v-for="todo in todos" :key="todo._id" :value="todo._id">
          <AccordionTrigger>
            {{ todo.title }}
          </AccordionTrigger>
          <AccordionContent class="flex justify-between items-center">
            <div>{{ todo.description }}</div>
            <div>
              <Badge :variant="todo.completed ? '' : 'destructive'">
                {{ todo.completed ? "Completed" : "Incomplete" }}
              </Badge>
            </div>
          </AccordionContent>
        </AccordionItem>
      </Accordion>
    </div>
  </div>
</template>

<script setup>
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";
import { Badge } from "@/components/ui/badge";
const runtimeConfig = useRuntimeConfig();
const baseUrl = runtimeConfig.public.baseAPI;

definePageMeta({
  middleware: ["auth"],
});

const token = useCookie("token");
const todos = useState("todos", () => []);

const { data } = await useFetch(baseUrl + "/todos", {
  method: "GET",
  headers: {
    Authorization: `Bearer ${token.value}`,
  },
});

todos.value = data.value.todos;
</script>

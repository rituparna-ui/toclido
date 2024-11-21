<template>
  <div class="p-8 flex-grow">
    <h1 class="text-center text-2xl">CLI Login Approval</h1>
    <div
      class="w-[400px] mx-auto mt-4 text-center flex items-center justify-center"
    >
      <Card class="p-4 h-[200px] flex flex-col justify-evenly">
        <div>
          Click the button below to login to the CLI app on your device.
        </div>
        <div class="mt-4">
          <Button class="w-full" @click="login" :disabled="approved">
            Login
          </Button>
        </div>
        <div v-if="approved" class="mt-4">
          CLI Login Succesful. Window will close in {{ timeout }} seconds.
        </div>
      </Card>
    </div>
  </div>
</template>

<script setup>
const runtimeConfig = useRuntimeConfig();
const baseUrl = runtimeConfig.public.baseAPI;

import { Button } from "@/components/ui/button";
import { toast } from "vue-sonner";

definePageMeta({
  middleware: ["auth"],
});

const token = useCookie("token");
const timeout = ref(3);
const approved = ref(false);

const login = async () => {
  try {
    const data = await $fetch(baseUrl + "/cli-tokens", {
      method: "POST",
      headers: {
        Authorization: `Bearer ${token.value}`,
      },
    });
    toast.success(data.message);
    try {
      await $fetch("http://localhost:1606", {
        method: "POST",
        body: data,
      });
    } catch (error) {
    } finally {
      approved.value = true;
      setTimeout(() => {
        window.close();
      }, 3000);
      setInterval(() => {
        timeout.value--;
      }, 1000);
    }
  } catch (error) {
    toast.error(error.data.message);
  }
};
</script>
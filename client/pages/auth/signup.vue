<template>
  <div class="flex-grow flex justify-center items-center">
    <Card class="w-96">
      <CardHeader>
        <CardTitle>Sign Up</CardTitle>
      </CardHeader>
      <CardContent>
        <form class="flex flex-col gap-4" @submit.prevent="signup">
          <Input type="text" placeholder="Name" required name="name" />
          <Input type="email" placeholder="Email" required name="email" />
          <Input
            type="password"
            placeholder="Password"
            required
            name="password"
          />
          <Button type="submit">Sign Up</Button>
        </form>
      </CardContent>
      <CardFooter class="flex justify-center gap-2">
        <div>Already have an account ?</div>
        <div><NuxtLink to="/auth/login">Login</NuxtLink></div>
      </CardFooter>
    </Card>
  </div>
</template>

<script setup>
definePageMeta({
  middleware: ["unauth"],
});
const runtimeConfig = useRuntimeConfig();
const baseUrl = runtimeConfig.public.baseAPI;

import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { toast } from "vue-sonner";

const signup = async (e) => {
  const formData = new FormData(e.target);
  const data = Object.fromEntries(formData.entries());
  const { name, email, password } = data;
  try {
    const data = await $fetch(baseUrl + "/auth/signup", {
      method: "POST",
      body: {
        name,
        email,
        password,
      },
    });
    toast.success(data.message);
    navigateTo("/auth/login");
  } catch (error) {
    toast.error(error.data.message);
  }
};
</script>


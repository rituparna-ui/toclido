<template>
  <div class="flex-grow flex justify-center items-center">
    <Card class="w-96">
      <CardHeader>
        <CardTitle>Login</CardTitle>
      </CardHeader>
      <CardContent>
        <form class="flex flex-col gap-4" @submit.prevent="login">
          <Input type="email" placeholder="Email" required name="email" />
          <Input
            type="password"
            placeholder="Password"
            required
            name="password"
          />
          <Button type="submit">Login</Button>
        </form>
      </CardContent>
      <CardFooter class="flex justify-center gap-2">
        <div>Don't have an account ?</div>
        <div><NuxtLink to="/auth/signup">Signup</NuxtLink></div>
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

import { useAuth } from "@/composables/useAuth";

const { authState } = useAuth();

const token = useCookie("token");

const login = async (e) => {
  const formData = new FormData(e.target);
  const data = Object.fromEntries(formData.entries());
  const { email, password } = data;
  try {
    const data = await $fetch(baseUrl + "/auth/login", {
      method: "POST",
      body: {
        email,
        password,
      },
    });
    toast.success(data.message);
    token.value = data.token;
    authState.value.isAuth = true;
    navigateTo("/");
  } catch (error) {
    toast.error(error.data.message);
  }
};
</script>


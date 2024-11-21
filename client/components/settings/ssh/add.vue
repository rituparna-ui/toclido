<template>
  <div class="my-8">
    <form class="flex flex-col gap-4" @submit.prevent="createSshKey">
      <Input
        type="text"
        class="text-black"
        placeholder="Name"
        required
        name="name"
      />
      <Input
        type="text"
        class="text-black"
        placeholder="Public Key"
        required
        name="publicKey"
      />
      <Button type="submit">Add SSH Key</Button>
    </form>
  </div>
</template>

<script setup>
const runtimeConfig = useRuntimeConfig();
const baseUrl = runtimeConfig.public.baseAPI;

import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { toast } from "vue-sonner";

const token = useCookie("token");
const { sshKeys } = useSshKeys();

const createSshKey = async (e) => {
  const formData = new FormData(e.target);
  const data = Object.fromEntries(formData.entries());
  const { name, publicKey } = data;
  try {
    const data = await $fetch(baseUrl + "/ssh-keys", {
      method: "POST",
      body: {
        name,
        publicKey,
      },
      headers: {
        Authorization: `Bearer ${token.value}`,
      },
    });
    toast.success(data.message);
    sshKeys.value = [data.sshKey, ...sshKeys.value];
    e.target.reset();
  } catch (error) {
    toast.error(error.data.message);
  }
};
</script>
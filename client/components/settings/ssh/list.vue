<template>
  <div>
    <div
      v-for="key of sshKeys"
      :key="key._id"
      class="border rounded-md my-4 p-4"
    >
      <div class="text-lg">{{ key.name }}</div>
      <div class="text-sm">{{ key.parsed.fingerprint }}</div>
      <div class="text-sm">
        Added on:
        {{
          new Date(key.createdAt).toLocaleString("en-US", {
            year: "numeric",
            month: "short",
            day: "numeric",
          })
        }}
      </div>
      <div class="text-sm">
        Last used on:
        {{
          new Date(key.lastUsedAt).toLocaleString("en-US", {
            year: "numeric",
            month: "short",
            day: "numeric",
          })
        }}
      </div>
      <div class="mt-2">
        <Button
          variant="destructive"
          class="w-full"
          @click="deleteSshKey(key._id)"
        >
          Delete
        </Button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { toast } from "vue-sonner";

const runtimeConfig = useRuntimeConfig();
const baseUrl = runtimeConfig.public.baseAPI;

const token = useCookie("token");

const { data } = await useFetch(baseUrl + "/ssh-keys", {
  method: "GET",
  headers: {
    Authorization: `Bearer ${token.value}`,
  },
});

const { sshKeys } = useSshKeys();
sshKeys.value = data.value.sshKeys;

const deleteSshKey = async (id) => {
  try {
    const data = await $fetch(baseUrl + "/ssh-keys/" + id, {
      method: "DELETE",
      headers: {
        Authorization: `Bearer ${token.value}`,
      },
    });
    toast.success(data.message);
    sshKeys.value = sshKeys.value.filter((key) => key._id !== id);
  } catch (error) {
    toast.error(error.data.message);
  }
};
</script>
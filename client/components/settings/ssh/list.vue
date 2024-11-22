<template>
  <div>
    <div
      v-for="key of sshKeys"
      :key="key._id"
      class="border rounded-md my-4 p-4 pl-8 flex items-center gap-8"
    >
      <div>
        <svg
          height="64"
          aria-hidden="true"
          viewBox="0 0 24 24"
          version="1.1"
          width="64"
          class="fill-white"
        >
          <path
            d="M16.75 8.5a1.25 1.25 0 1 0 0-2.5 1.25 1.25 0 0 0 0 2.5Z"
          ></path>
          <path
            d="M15.75 0a8.25 8.25 0 1 1-2.541 16.101l-1.636 1.636a1.744 1.744 0 0 1-1.237.513H9.25a.25.25 0 0 0-.25.25v1.448a.876.876 0 0 1-.256.619l-.214.213a.75.75 0 0 1-.545.22H5.25a.25.25 0 0 0-.25.25v1A1.75 1.75 0 0 1 3.25 24h-1.5A1.75 1.75 0 0 1 0 22.25v-2.836c0-.464.185-.908.513-1.236l7.386-7.388A8.249 8.249 0 0 1 15.75 0ZM9 8.25a6.733 6.733 0 0 0 .463 2.462.75.75 0 0 1-.168.804l-7.722 7.721a.25.25 0 0 0-.073.177v2.836c0 .138.112.25.25.25h1.5a.25.25 0 0 0 .25-.25v-1c0-.966.784-1.75 1.75-1.75H7.5v-1c0-.966.784-1.75 1.75-1.75h1.086a.25.25 0 0 0 .177-.073l1.971-1.972a.75.75 0 0 1 .804-.168A6.75 6.75 0 1 0 9 8.25Z"
          ></path>
        </svg>
      </div>
      <div class="w-full overflow-hidden">
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
        <div>
          <Button
            variant="destructive"
            class="mt-2"
            @click="deleteSshKey(key._id)"
          >
            Delete
          </Button>
        </div>
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
<template>
  <div>
    <div>
      <h1 class="text-center text-2xl mt-4">Authorized Tokens</h1>
      <div class="w-full mx-auto mt-4">
        <div
          v-for="token of tokens"
          :key="token.token"
          class="border rounded-md my-4 p-4 pl-8 flex justify-between items-center"
        >
          <div class="flex flex-col gap-3">
            <div class="text-bold text-xl">{{ token.token }}</div>
            <div>
              Created on:
              {{
                new Date(token.createdAt).toLocaleString("en-US", {
                  year: "numeric",
                  month: "short",
                  day: "numeric",
                })
              }}
            </div>
            <div>
              Last used on:
              {{
                new Date(token.updatedAt).toLocaleString("en-US", {
                  year: "numeric",
                  month: "short",
                  day: "numeric",
                })
              }}
            </div>
          </div>
          <div>
            <Button
              variant="destructive"
              class="mt-2"
              @click="deleteCliToken(token.token)"
            >
              Revoke
            </Button>
          </div>
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

const { data } = await useFetch(baseUrl + "/cli-tokens", {
  method: "GET",
  headers: {
    Authorization: `Bearer ${token.value}`,
  },
});

const tokens = useState("tokens-list", () => data.value.tokens);

const deleteCliToken = async (cliToken) => {
  try {
    const data = await $fetch(baseUrl + "/cli-tokens/revoke", {
      method: "POST",
      body: {
        token: cliToken,
      },
      headers: {
        Authorization: `Bearer ${token.value}`,
      },
    });
    tokens.value = tokens.value.filter((token) => token.token !== cliToken);
    toast.success(data.message);
  } catch (error) {
    toast.error(error.data.message);
  }
};
</script>
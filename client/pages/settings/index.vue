<template>
  <div class="p-8">
    <h1 class="text-center text-2xl">Settings</h1>
    <div class="w-full lg:w-1/2 mx-auto mt-4">
      <Tabs default-value="account" class="w-full">
        <TabsList class="w-full bg-[#393E46]">
          <TabsTrigger class="w-full text-lg" value="account">
            SSH Keys
          </TabsTrigger>
          <TabsTrigger class="w-full text-lg" value="password">
            CLI Tokens
          </TabsTrigger>
        </TabsList>
        <TabsContent value="account">
          <div>
            <SettingsSshAdd />
          </div>
        </TabsContent>
        <TabsContent value="password">
          <div>CLI Tokens</div>
        </TabsContent>
      </Tabs>
    </div>
  </div>
</template>

<script setup>
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

definePageMeta({
  middleware: ["auth"],
});
</script>
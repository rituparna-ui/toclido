<script setup>
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";

const { logout } = useAuth();
</script>

<template>
  <DropdownMenu>
    <DropdownMenuTrigger>
      <Avatar>
        <AvatarImage
          src="https://github.com/rituparna-ui.png"
          alt="@rituparna-ui"
        />
        <AvatarFallback>RW</AvatarFallback>
      </Avatar>
    </DropdownMenuTrigger>
    <DropdownMenuContent>
      <DropdownMenuLabel>My Account</DropdownMenuLabel>
      <DropdownMenuSeparator />
      <NuxtLink to="/">
        <DropdownMenuItem> Home </DropdownMenuItem>
      </NuxtLink>
      <NuxtLink to="/settings">
        <DropdownMenuItem> Settings </DropdownMenuItem>
      </NuxtLink>
      <DropdownMenuItem @click="logout">Logout</DropdownMenuItem>
    </DropdownMenuContent>
  </DropdownMenu>
</template>
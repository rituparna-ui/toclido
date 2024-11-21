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
      <DropdownMenuItem>Profile</DropdownMenuItem>
      <DropdownMenuItem>
        <NuxtLink to="/settings">Settings</NuxtLink>
      </DropdownMenuItem>
      <DropdownMenuItem @click="logout">Logout</DropdownMenuItem>
    </DropdownMenuContent>
  </DropdownMenu>
</template>
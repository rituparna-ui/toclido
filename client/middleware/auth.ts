export default defineNuxtRouteMiddleware((to, from) => {
  const token = useCookie("token");
  const { authState } = useAuth();
  const route = useRoute();
  console.log(route.path);
  if (token.value) {
    authState.value.isAuth = true;
    return true;
  }
  return navigateTo("/auth/login?redirect=" + route.path);
});

export default defineNuxtRouteMiddleware((to, from) => {
  const token = useCookie("token");
  const { authState } = useAuth();
  if (token.value) {
    authState.value.isAuth = true;
    return navigateTo("/");
  }
});

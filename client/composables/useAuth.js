export const useAuth = () => {
  const token = useCookie("token");
  const authState = useState("auth", () => ({
    isAuth: false,
  }));

  const logout = () => {
    authState.value.isAuth = false;
    token.value = null;
    navigateTo("/auth/login");
  };

  return { authState, logout };
};

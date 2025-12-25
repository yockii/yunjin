import type { UserInfo } from "@/types/user-info";
import { defineStore } from "pinia";
import { ref } from "vue";

export const useUserStore = defineStore("user", () => {
  const token = ref("");
  const userInfo = ref({} as UserInfo);

  const setToken = (newToken: string) => {
    token.value = newToken;
    sessionStorage.setItem("token", newToken);
  };

  const setUserInfo = (newUserInfo: UserInfo) => {
    userInfo.value = newUserInfo;
    sessionStorage.setItem("userInfo", JSON.stringify(newUserInfo));
  };

  const logout = () => {
    token.value = "";
    userInfo.value = {} as UserInfo;
    sessionStorage.removeItem("token");
    sessionStorage.removeItem("userInfo");
  };

  return {
    token,
    userInfo,
    setToken,
    setUserInfo,
    logout,
  };
});

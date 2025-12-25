import { defineStore } from "pinia";
import { ref } from "vue";
import type { Menu } from "@/types/menu";

export const useMenuStore = defineStore("menu", () => {
  const menu = ref<Menu[]>([]);
  const selectedMenu = ref<Menu | null>(null);

  const setMenu = (newMenu: Menu[]) => {
    menu.value = newMenu;
  };

  const setSelectedMenu = (newSelectedMenu: Menu | null) => {
    selectedMenu.value = newSelectedMenu;
  };

  return {
    menu,
    selectedMenu,
    setMenu,
    setSelectedMenu,
  };
});

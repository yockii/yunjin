<template>
    <div
        class="w-full h-full p-8 flex flex-col gap-2 justify-center items-center"
    >
        <FloatLabel variant="on">
            <InputText
                id="username"
                size="large"
                fluid
                v-model="username"
                autocomplete="off"
            />
            <label for="username">用户名</label>
        </FloatLabel>
        <FloatLabel variant="on">
            <Password
                size="large"
                fluid
                inputId="password"
                v-model="password"
            />
            <label for="password">密码</label>
        </FloatLabel>
        <ButtonGroup class="flex justify-center items-center">
            <Button
                label="登录"
                size="large"
                icon="pi pi-check"
                class="w-full"
                :loading="loading"
                @click="login"
            />
        </ButtonGroup>
    </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { FloatLabel, InputText, Password, ButtonGroup, Button } from "primevue";

const emit = defineEmits(["loginFailed", "loginSuccess"]);

const username = defineModel("username", {
    required: true,
    type: String,
});
const password = defineModel("password", {
    required: true,
    type: String,
});
const loading = ref(false);

const login = async () => {
    loading.value = true;

    // 模拟验证
    await new Promise((r) => setTimeout(r, 800));

    loading.value = false;

    if (username.value === "admin" && password.value === "123456") {
        emit("loginSuccess");
    } else {
        emit("loginFailed");
    }
};
</script>

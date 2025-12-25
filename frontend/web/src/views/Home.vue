<template>
    <div class="w-full h-screen flex flex-col items-center justify-center">
        <div class="relative perspective-distant">
            <div
                id="book"
                class="relative w-100 h-150 transform-3d rotate-x-45"
            >
                <div
                    id="book-decorations"
                    class="absolute bottom-0 left-0.5 w-full"
                >
                    <div class="relative">
                        <!-- Book Pages (thickness) -->
                        <div
                            id="book-pages"
                            class="absolute left-0 right-2 bottom-0.5 h-6 bg-[repeating-linear-gradient(to_right,#e8e5dc,#e0dccf_2px,#f5f2e9_4px)] rounded-l-lg shadow-inner"
                        ></div>
                        <!-- 书脊 -->
                        <div
                            id="book-spine"
                            class="absolute left-0 bottom-0 w-1 h-6 bg-black rounded-bl-lg -rotate-8"
                        ></div>

                        <!-- 书底部 -->
                        <div
                            class="absolute left-0.5 -bottom-0.5 right-1.5 h-1 bg-black rounded-b-lg"
                        ></div>
                    </div>
                </div>
                <!-- Book Cover -->
                <div
                    id="book-cover"
                    class="translate-z-px origin-left absolute left-0 top-0 right-0 bottom-5 transform-3d pl-1"
                >
                    <div
                        id="cover-front"
                        class="absolute inset-0 bg-linear-to-br from-black to-blue-300 rounded-lg flex items-center justify-center text-white text-4xl font-bold tracking-widest shadow-2xl backface-hidden"
                    >
                        云烬
                    </div>
                    <!-- 封面内页 -->
                    <div
                        id="cover-inner"
                        class="absolute inset-0 bg-[#f4f1ea] rounded-lg rotate-y-180 backface-hidden"
                    ></div>
                </div>

                <div
                    id="book-inner"
                    class="absolute left-0 top-1 right-2 bottom-6 bg-[#f6f2ea] rounded-md shadow-inner"
                >
                    <div
                        class="w-full p-8 absolute inset-0 transform-none flex flex-col gap-2 justify-center items-center"
                    >
                        <!-- 内页内容 -->
                        <FloatLabel variant="on">
                            <InputText
                                id="username"
                                size="large"
                                fluid
                                v-model="loginInfo.username"
                                autocomplete="off"
                            />
                            <label for="username">用户名</label>
                        </FloatLabel>
                        <FloatLabel variant="on">
                            <Password
                                size="large"
                                fluid
                                inputId="password"
                                v-model="loginInfo.password"
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
                </div>
                <!-- Book Shadow -->
                <div
                    class="absolute -bottom-10 left-10 right-10 h-24 bg-black/50 blur-3xl rounded-full"
                ></div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import gsap from "gsap";
import { FloatLabel, InputText, Password, ButtonGroup, Button } from "primevue";

const loginInfo = ref<LoginInfo>({
    username: "",
    password: "",
});
const loading = ref(false);

const login = () => {
    // 登录逻辑
};

const animate = () => {
    const tl = gsap.timeline({
        defaults: {
            ease: "power3.inOut",
        },
    });

    tl.to("#book", {
        duration: 1,
        x: 200,
    });
    tl.to(
        "#book-decorations",
        {
            duration: 1,
            // 由于3d视角会倾斜，所以整个装饰区域要纠正倾斜角度
            skewX: -8,
            x: -2,
            width: "-=4",
        },
        "<",
    );
    tl.to("#book-cover", {
        rotateY: -180,
        duration: 2,
    });
    // 视角从X-45变更为X-90
    tl.to("#book", {
        rotateX: 0,
        duration: 2,
    });
    // 同步需要把装饰压扁一点
    tl.to(
        "#book-decorations",
        {
            scaleY: 0.6,
            duration: 2,
            bottom: "+=10",
        },
        "<",
    );
};

onMounted(() => {
    animate();
});
</script>

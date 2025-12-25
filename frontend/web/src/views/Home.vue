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
                        class="absolute inset-0 bg-linear-to-br from-black to-blue-300 rounded-lg flex items-center justify-center text-white text-4xl font-bold tracking-widest shadow-2xl backface-hidden border-gray-300 border border-solid"
                    >
                        欢迎使用
                    </div>
                    <!-- 封面内页 -->
                    <div
                        id="cover-inner"
                        class="absolute inset-0 bg-[#f4f1ea] rounded-lg rotate-y-180 backface-hidden border-gray-300 border border-solid"
                    ></div>
                </div>

                <div
                    id="book-inner"
                    class="absolute left-0 top-1 right-2 bottom-6 bg-[#f6f2ea] rounded-md shadow-inner"
                >
                    <div
                        id="page-current"
                        class="z-10 w-full h-full absolute inset-0 transform-3d rounded-md border-gray-300 border border-solid"
                    >
                        <LoginForm
                            @login-failed="tearPage"
                            v-model:username="loginInfo.username"
                            v-model:password="loginInfo.password"
                        />
                    </div>

                    <div
                        id="page-tear"
                        class="z-0 w-full h-full absolute inset-0 transform-3d pointer-events-none backface-hidden bg-[#f6f2ea] rounded-md shadow-inner border-gray-300 border border-solid"
                    ></div>
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
import LoginForm from "@/components/LoginForm.vue";
import { LoginInfo } from "@/types/user-info";

const loginInfo = ref<LoginInfo>({
    username: "",
    password: "",
});

const mountedAnimate = () => {
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
            bottom: "+=12",
        },
        "<",
    );
};

const tearPage = () => {
    const tear = document.getElementById("page-tear");
    const current = document.getElementById("page-current");

    tear.innerHTML = current.innerHTML;
    loginInfo.value.username = "";
    loginInfo.value.password = "";

    gsap.set(tear, {
        zIndex: 20,
        opacity: 1,
        rotateY: 0,
        rotateZ: 0,
        x: 0,
        y: 0,
        skewX: 0,
        skewY: 0,
        transformOrigin: "0% 100%",
    });

    const tl = gsap.timeline({
        ease: "power2.inOut",
    });
    tl.to(tear, {
        rotateY: -28,
        skewY: -18,
        duration: 0.35,
    });
    tl.to(tear, {
        delay: 0.2,
        rotateY: -55,
        skewX: -10,
        skewY: -8,
        duration: 0.6,
    });
    tl.to(tear, {
        rotateY: -95,
        rotateZ: 8,
        x: 140,
        y: 90,
        opacity: 0,
        duration: 0.9,
    });

    // tl.call(() => {
    //     gsap.set(tear, { zIndex: 0, rotateY: 0, x: 0, y: 0 });
    // });
};

onMounted(() => {
    mountedAnimate();
});
</script>

<template>
    <v-app-bar app>
        <v-app-bar-title><v-btn to="/">Hackathon6</v-btn></v-app-bar-title>
        <v-spacer></v-spacer>
        <v-btn to="/help">ヘルプ</v-btn>
        <v-btn v-if="!isAuthenticated" to="/login">ログイン</v-btn>
        <v-btn v-else to="/mypage">マイページ</v-btn>
    </v-app-bar>
</template>

<script>
    export default {
        data: () => ({
            isAuthenticated: true,
        })
    }
</script>
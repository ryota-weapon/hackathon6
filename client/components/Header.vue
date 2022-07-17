<template>
    <v-app-bar app>
        <v-app-bar-title><v-btn width="" text to="/" class="text-h6">みんなの選手評価</v-btn></v-app-bar-title>
        <v-spacer></v-spacer>
        <v-btn text to="/help">ヘルプ</v-btn>
        
        <v-btn text  to="/login">ログイン</v-btn>
        
        <v-btn text v-if="!isAuthenticated" to="/login">ログイン</v-btn>
        <v-btn text v-else to="/mypage">マイページ</v-btn>
    </v-app-bar>
</template>

<script>
    // import { onAuthStateChanged } from '@firebase/auth';
    import { auth } from '~/plugins/firebase';
    
    // onAuthStateChanged = () => {
    //     return;
    // }

    export default {
        data: () => ({
            isAuthenticated: true,
        }),
        computed: {
        }
    }
</script>
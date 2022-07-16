<template>
    <v-container>
        <h1>マイページ</h1>
        <v-card class="mt-6">
            <v-row>
                <v-col>
                    <v-card-title>名前</v-card-title>
                </v-col>
                <v-col>
                    <v-card-title>{{user.name}}</v-card-title>
                </v-col>
            </v-row>

            <v-row>
                <v-col>
                    <v-card-title>アイコン</v-card-title>
                </v-col>
                <v-col>
                    <img style="border-radius: 50%;" src="https://lh3.googleusercontent.com/a/AItbvmlpWHw7gtsVkl2pf80cRJ9r3VerHDjI4IQzBY0bjg=s96-c" alt="hoge">
                    <!-- <v-img :source="user.icon"></v-img> -->
                </v-col>
            </v-row>

            <!-- <v-row>
                <v-col>
                    <v-card-title>名前</v-card-title>
                </v-col>
                <v-col>
                    <v-card-title>{{user.email}}</v-card-title>
                </v-col>
            </v-row> -->
        </v-card>
        <v-btn @click="signOut">サインアウト</v-btn>
    </v-container>
</template>

<script>
    import axios from "axios";
    import { auth } from "~/plugins/firebase";

    auth.onAuthStateChanged((user) => {
        if (!user) {
            // サインインしていない状態
            // サインイン画面に遷移する等
            // 例:
            console.log("not sign in");
            location.href = '/login';
        }else{
            console.log(user);
        }
    });

    export default {
        data: () => ({
            user: {displayName: "サッカー評論家", iconUrl: "https://", },
        }),
        created() {
            console.log(auth.currentUser);
            axios.get("http://localhost:1323/user/"+auth.currentUser.uid).then(res => {
                console.log(res);
                this.user = res.data;
                console.log(this.user);
            })
        },
        methods: {
            signOut: function() {
                auth.signOut();
                console.log(auth);
            }
        }
    }
</script>


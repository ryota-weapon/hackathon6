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
                    <img style="border-radius: 50%;" :src="user.icon" alt="">
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

    export default {
        data: () => ({
            user: {name: "サッカー評論家", icon: "https://", },
        }),
        methods: {
            signOut: function() {
                auth.signOut();
                this.$store.commit("auth/setAuthenticated", "");
            }
        },
        mounted(){
            const self = this;
            auth.onAuthStateChanged(function(user) {
                if (user) {
                    axios.get("http://localhost:1323/user/"+ user.uid ).then(res => {
                        console.log(res.data);
                        console.log(self);
                        self.user = res.data;
                    })
                } else {
                    location.href = '/login';
                }
            });     
        }
    }
</script>


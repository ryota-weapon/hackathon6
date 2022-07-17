<template>
    <v-container>
        <h1>ログイン</h1>
        <p>このアプリは、グーグルアカウントでのログインのみサポートしています。</p>
        <v-btn @click="logInWithGoogle">グーグルアカウントでログイン</v-btn>
    </v-container>
</template>

<script>
    import { auth } from '../plugins/firebase';
    import { signInWithPopup, GoogleAuthProvider} from "firebase/auth";
    import axios from "axios";
import { env } from 'process';


    export default {
        methods: {
            logInWithGoogle: function(){
                const provider = new GoogleAuthProvider();
                signInWithPopup(auth, provider).then( result => {
                    this.$store.commit("auth/setAuthenticated", result.user.uid);
                })
            }
        },
    }
</script>
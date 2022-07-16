<template>
    <v-container>
        <v-btn @click="signUpWithGoogle">グーグルアカウントで登録</v-btn>
    </v-container>
</template>

<script>
    import { auth } from '../plugins/firebase';
    import { signInWithPopup, GoogleAuthProvider} from "firebase/auth";
    import axios from "axios";
import { env } from 'process';


    export default {
        methods: {
            signUpWithGoogle: function() {
                const provider = new GoogleAuthProvider();
                signInWithPopup(auth, provider).then(result => {
                    console.log(result);
                    const info = result.user;
                    const postData = {
                        "id": info.uid,
                        "name": info.displayName,
                        "icon": info.photoURL,
                        "email": info.email,
                    };
                    axios.post("http://localhost:1323/", postData).then(
                        result => console.log(result)
                    )
                });
            }
        }
    }
</script>
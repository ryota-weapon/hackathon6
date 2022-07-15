<template>
    <v-container>
        <v-btn @click="signUpWithGoogle">Sign-in with google</v-btn>
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
                    axios.post("http://localhost:1323/", {"authId": result.user.uid, "name": result.user.displayName}).then(
                        result => console.log(result)
                    )
                });
            }
        }
    }
</script>
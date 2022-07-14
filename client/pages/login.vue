<template>
    <v-container>
        <h1>Login</h1>
    </v-container>
</template>

<script>
    import {auth} from '../plugins/firebase'
    import { signInWithPopup, GoogleAuthProvider} from "firebase/auth"

    const auth = firebase.auth();
        export default {
        methods: {
            signUpWithGoogle: function() {
                // const provider = new GoogleAuthProvider();
                // signInWithPopup(auth, provider).then(result => {
                //     console.log(result)
                // });
                window.alert("login");
            }
        }
    }
</script>
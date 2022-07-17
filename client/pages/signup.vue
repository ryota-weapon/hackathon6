<template>
    <v-container>
        <h1>登録</h1>
        <p>このアプリは、グーグルアカウントでの登録のみサポートしています。</p>
        <p>登録することで、以下の機能が使えるようになります。</p>
        <ul>
            <li>試合の作成</li>
            <li>選手の評価</li>
            <li>チャットの投稿</li>
        </ul>

        <v-container justify="center">
            <v-btn x-large @click="signUpWithGoogle">グーグルアカウントで登録</v-btn>
        </v-container>
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
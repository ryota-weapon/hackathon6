<template>
    <v-container>
        Make
        <v-form>
            <v-text-field v-model="team1" label="チーム1"></v-text-field>
            <v-text-field v-model="team2" label="チーム2"></v-text-field>
            <v-text-field v-model="startTime" label="開始時刻" type="date"></v-text-field>
            <v-btn @click="makeMatch">登録</v-btn>
        </v-form>
    </v-container>
</template>

<script>
    import { interfaceDeclaration } from "@babel/types";
import axios from "axios";

    export default {
        data: () => ({
            team1: "",
            team2: "",
            startTime: "2022-07-15",
            // startTime: Date.today,
        }),
        methods: {
            makeMatch: function() {
                const postData = {
                    startTime: this.startTime,
                    team1: parseInt(this.team1),
                    team2: parseInt(this.team2),
                }
                console.log(postData);
                axios.post("http://localhost:1323/makeMatch", postData).then(res => {
                    console.log(res);
                });
            },
        }
    }
</script>
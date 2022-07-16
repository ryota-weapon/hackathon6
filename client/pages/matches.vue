<template>
    <v-container>
        <v-list>
            <v-list-item-content v-for="match in matches" :key="match.id">
                {{match.id}}
                <router-link :to="{name: 'match', params: {id: match.id}}">
                    <v-list-item-title color="accent">{{match.team1}} x {{match.team2}}</v-list-item-title>
                    <v-list-item-subtitle>{{match.startTime}}</v-list-item-subtitle>
                </router-link>
            </v-list-item-content>
        </v-list>
    </v-container>
</template>

<script>
    import axios from 'axios';

    export default {
        data: () => ({
            matches: []
        }),
        created() {
            axios.get("http://localhost:1323/matches").then(res => {
                // this.matches = res.data;
                this.matches = [
                    {id: 1, team1: "日本", team2: "ウルグアイ", startTime: "2022-07-20"},
                    {id: 2, team1: "ブラジル", team2: "ウルグアイ", startTime: "2022-07-22"},
                    {id: 3, team1: "スペイン", team2: "フランス", startTime: "2022-07-24"},
                ]
            }) 
        },
        methods: {
            fetchAllMatches: function(){
                axios.get("http://localhost:1323/matches").then(res => {
                    console.log(res);
                    console.log(res.data)
                    this.matches = res.data;
                }) 
            }
        }
    }
</script>
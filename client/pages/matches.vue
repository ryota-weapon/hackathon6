<template>
    <v-container>
        <!-- <v-row>
            <v-col cols="10"></v-col>
            <v-col>
                並び順<v-select></v-select>
            </v-col>
        </v-row> -->

        <v-list>
            <v-card class="py-4 px-4" v-for="match in matches" :key="match.id">
                <router-link :to="{name: 'match', query: {id: encodeURIComponent(match.id)}}">
                    <v-list-item-title color="accent">{{match.team1}} x {{match.team2}}</v-list-item-title>
                    <v-list-item-subtitle>{{match.startTime}}</v-list-item-subtitle>
                </router-link>
            </v-card>
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
                    {id: "VjvLFbaaAO8WqAbv6d4o", team1: "日本", team2: "ウルグアイ", startTime: "2022-07-20"},
                    {id: "VjvLFbaaAO8WqAbv6d4o", team1: "ブラジル", team2: "ウルグアイ", startTime: "2022-07-22"},
                    {id: "VjvLFbaaAO8WqAbv6d4o", team1: "スペイン", team2: "フランス", startTime: "2022-07-24"},
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
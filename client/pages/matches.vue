<template>
    <v-container>
        <v-list>
            <v-list-item-content v-for="match in matches" :key="match.id">
                {{match.id}}
                <router-link :to="{name: 'match', params: {id: match.id}}">
                    <v-list-item-title>{{match.team1}} x {{match.team2}}</v-list-item-title>
                    <v-list-item-subtitle>{{match.startTime}}</v-list-item-subtitle>
                </router-link>
            </v-list-item-content>
        </v-list>
        <v-btn @click="fetchAllMatches">fetch</v-btn>
    </v-container>
</template>

<script>
    import axios from 'axios';

    export default {
        data: () => ({
            matches: []
        }),
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
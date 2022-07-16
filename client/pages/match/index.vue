<template>
    <v-container>
        <p>試合{{match.team1}} x {{match.team2}}</p>
        <p> from: {{match.startTime}}</p>

        <v-row>
            <v-col class="pr-0">
                <v-card>
                    <v-card-title class="justify-center">{{match.team1}}</v-card-title>
                    <v-divider></v-divider>
                    <v-card-text  v-for="player in [{name: '松木栗生', id: 1},{name: '本田圭佑', id: 2},{name: '川島英治', id: 3}]" :key="player.id">
                        {{player.name}}
                    </v-card-text >
                </v-card>
            </v-col >
            <v-col class="pl-0">
                <v-card>
                    <v-card-title class="justify-center">{{match.team2}}</v-card-title>
                    <v-divider></v-divider>
                    <v-card-text  v-for="player in [{name: '上村涼太', id: 1},{name: 'シュミット・ダニエル', id: 2},{name: '久保建英', id: 3}]" :key="player.id">
                        {{player.name}}
                        <v-card-action>
                            <v-btn>
                                <v-icon>mdi-thumb-up</v-icon>
                            </v-btn>
                        </v-card-action>
                        <v-card-action>
                            <v-btn>
                                <v-icon>mdi-thumb-down</v-icon>
                            </v-btn>
                        </v-card-action>
                    </v-card-text >
                </v-card>
            </v-col>
        </v-row>

        <h2>チャット</h2>
        <p>※不適切な発言は慎むよう、よろしくお願いします。</p>
        <v-card>
            <v-card-text>遠藤の守備がうますぎて相手泣いてる</v-card-text>
            <v-card-text>まじでうまかった。</v-card-text>
            <v-card-text>いなかったら２点は取られてたと思う。</v-card-text>
        </v-card>


    </v-container>
</template>

<script>
    import axios from "axios"
    export default {
        // async asyncData({}) {
        //     const match = await axios.get(
        //         "https://localhost:1323/match",
        //         {
        //             params: {id: this.$route.params.id},
        //         }
        //     );
        //     return { match };
        // },
        data: function(){
            return {
                matchId: this.$route.params.id,
                match: {},
            }
        },
        created() {
            axios.get("http://localhost:1323/match/"+this.matchId).then(res => {
                this.match = res.data;
                console.log(res);
            })
        },
        methods: {
            
        }
    }
</script>


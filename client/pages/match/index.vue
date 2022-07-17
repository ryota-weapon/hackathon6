<template>
    <v-container>
        <p>試合{{match.team1}} x {{match.team2}}</p>
        <p> from: {{match.startTime}}</p>

        <v-row>
            <v-col class="pr-0">
                <v-card>
                    <v-card-title class="justify-center">{{match.team1}}</v-card-title>
                    <v-divider></v-divider>

                    <v-card-text  v-for="player in [{name: '松木栗生', id: '1'},{name: '本田圭佑', id: '2'},{name: '川島英治', id: '3'}]" :key="player.id">
                        <v-row>
                            <v-col>
                                {{player.name}}
                            </v-col>
                            <v-spacer></v-spacer>
                            <v-col class="pr-0">
                                <v-btn text rounded>
                                    <v-icon @click="evaluatePlayer(player.id ,1)" v-if="evaluations[player.id]!=1">mdi-thumb-up-outline</v-icon>
                                    <v-icon @click="evaluatePlayer(player.id ,0)" color="grey" v-else >mdi-thumb-up</v-icon>
                                </v-btn>
                                <v-btn text rounded>
                                    <v-icon @click="evaluatePlayer(player.id ,-1)" v-if="evaluations[player.id]!=-1">mdi-thumb-down-outline</v-icon>
                                    <v-icon @click="evaluatePlayer(player.id ,0)"  color="grey" v-else >mdi-thumb-down</v-icon>
                                </v-btn>
                            </v-col>
                        </v-row>
                    </v-card-text >
                </v-card>
            </v-col >
            <v-col class="pl-0">
                <v-card>
                    <v-card-title class="justify-center">{{match.team2}}</v-card-title>
                    <v-divider></v-divider>
                    <v-card-text  v-for="player in [{name: '上村涼太', id: '4'},{name: 'シュミット・ダニエル', id: '5'},{name: '久保建英', id: '6'}]" :key="player.id">
                                                <v-row>
                            <v-col>
                                {{player.name}}
                            </v-col>
                            <v-spacer></v-spacer>
                            <v-col class="pr-0">
                                <v-btn text rounded>
                                    <v-icon @click="evaluatePlayer(player.id ,1)" v-if="evaluations[player.id]!=1">mdi-thumb-up-outline</v-icon>
                                    <v-icon @click="evaluatePlayer(player.id ,0)" color="grey" v-else >mdi-thumb-up</v-icon>
                                </v-btn>
                                <v-btn text rounded>
                                    <v-icon @click="evaluatePlayer(player.id ,-1)" v-if="evaluations[player.id]!=-1">mdi-thumb-down-outline</v-icon>
                                    <v-icon @click="evaluatePlayer(player.id ,0)"  color="grey" v-else >mdi-thumb-down</v-icon>
                                </v-btn>
                            </v-col>
                        </v-row>
                    </v-card-text >
                </v-card>
            </v-col>
        </v-row>

        <h2>チャット</h2>
        <p>選手や試合についてみんなで話しましょう！</p>
        <p style="font-size: small;">※不適切な発言は慎むよう、よろしくお願いします。</p>
        <v-card>
            <v-card-text>遠藤の守備がうますぎて相手泣いてる</v-card-text>
            <v-card-text>まじでうまかった。</v-card-text>
            <v-card-text>いなかったら２点は取られてたと思う。</v-card-text>
        </v-card>
    </v-container>
</template>

<script>
    import { auth } from "~/plugins/firebase"
    import axios from "axios"

    export default {
        data: function(){
            return {
                matchId: this.$route.query.id,
                matchId: "",
                match: {id: "", team1: "team1", team2: "team2"},
                evaluations: {"1": 1, "2": 0, "3": -1, "4": 1, "5": -1, "6": 0},   //id -> value（選手idからあるユーザーの評価の写像） 
                allEvaluations: {"1": 1, "2": 0, "3": -1, "4": 1, "5": -1, "6": 0}   //id -> value（選手idから全ユーザーの評価の写像） 
            }
        },
        created() {
            axios.get("http://localhost:1323/match/"+this.matchId).then(res => {
                console.log(res);
                this.match = res.data;
            })

            const postData = {
                matchId: "",
                playersId: [],
            }
            axios.get("http://localhost:1323/evaluation-all", postData).then(res => {
                console.log(res);
                this.evaluations = res.data;
            })
        },
        methods: {
            
            evaluatePlayer: function(playerId, value) {
                if (auth.currentUser == null) {
                    console.log("pleas login to evalueae players");
                    return;
                }

                const postData = {
                    matchId: this.$route.query.id,
                    playerId: playerId,
                    Value:    value,
                    userId:  auth.currentUser.uid,
                }

                console.log(postData);
                axios.post("http://localhost:1323/evaluate", postData).then(res => {
                    console.log(res);
                })

                this.evaluations[playerId] = value;
            }
        },

    }
</script>


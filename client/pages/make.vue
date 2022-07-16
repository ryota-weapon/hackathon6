<template>
    <v-container>
        <h1>試合を作成</h1>
        <v-form>
            <v-text-field v-model="team1" label="チーム1"></v-text-field>
            <v-text-field v-model="team2" label="チーム2"></v-text-field>
            <v-text-field v-model="startTime" label="開始時刻" type="date"></v-text-field>
            <p>選手追加</p>
            <v-row>
                <v-col cols="11">
                    <v-select
                        v-model="playerToAdd"
                        :items="players"
                        item-text="name"
                        :menu-props="{ top: true, offsetY: true }"
                    ></v-select>
                </v-col>
                <v-col cols="1">
                    <v-btn @click="addPlayer()">追加</v-btn>
                </v-col>
            </v-row>
            
            

            <v-row height="100px" class="mt-4">
                <v-hover v-slot:default="{ hover }">
                    <v-col :class="{'hover': hover, 'selected': selected==players1}" @click="selected = players1">
                            <v-card v-for="player in players1" >
                                <v-row class="mb-2">
                                    <v-col>
                                        <v-card-title>
                                            {{player.name}}
                                        </v-card-title>
                                    </v-col>

                                    <v-spacer></v-spacer>
                                    <v-col cols="2" class="pt-6">
                                        <v-card-action justify="center"><v-btn class="px-0">
                                            <v-icon >mdi-delete</v-icon>
                                        </v-btn></v-card-action>
                                    </v-col>
                                </v-row>
                            </v-card>
                    </v-col>
                </v-hover>
                <v-hover v-slot:default="{ hover }">
                    <v-col 
                        max-height="300px"
                        :class="{'hover': hover, 'selected': selected==players2}" 
                        @click="selected = players2" 
                    >
                        <v-card v-for="player in players2">
                            <v-row class="mb-2">
                                <v-col>
                                    <v-card-title>
                                        {{player.name}}
                                    </v-card-title>
                                </v-col>

                                <v-spacer></v-spacer>
                                <v-col cols="2" class="pt-6">
                                    <v-card-action justify="center"><v-btn class="px-0">
                                        <v-icon >mdi-delete</v-icon>
                                    </v-btn></v-card-action>
                                </v-col>
                            </v-row>
                        </v-card>
                    </v-col>
                </v-hover>
            </v-row>
        </v-form>
        
        <v-btn @click="makeMatch">登録</v-btn>
    </v-container>
</template>

<script>
    import { throwStatement } from "@babel/types";
import axios from "axios";

    const date = new Date();
    const y = date.getFullYear().toString();
    let m = (date.getMonth()+1).toString();
    let d = date.getDate().toString();
    // m = m<"10"? "0"+m : m;
    // d = d<"10"? "0"+d : d;
    if (m.length==1) m = "0"+m;
    if (d.length==1) d = "0"+d;
    const today = y+"-"+m+"-"+d;


    export default {
        data: () => ({
            team1: "",
            team2: "",
            startTime: today,
            players1: [
                {id: 1,  name: "酒井 宏樹"},
                {id: 2,  name: "本田 圭佑"},
                {id: 3,  name: "佐々木 朗希"},
                {id: 4,  name: "遠藤 航"},
                {id: 5,  name: "大迫 雄也"},
                {id: 6,  name: "向井 修"},
                {id: 7,  name: "林 圭佑"},
                {id: 8,  name: "大谷 二郎"},
                {id: 9,  name: "磯塚 体制"},
                {id: 10, name:  "井伊 直弼"},
                {id: 11, name:  "伊能 忠敬"},
            ],
            players2: [
                {id: 1,  name: "酒井 宏樹"},
                {id: 2,  name: "本田 圭佑"},
                {id: 3,  name: "佐々木 朗希"},
                {id: 4,  name: "遠藤 航"},
                {id: 5,  name: "大迫 雄也"},
                {id: 6,  name: "向井 修"},
                {id: 7,  name: "林 圭佑"},
                {id: 8,  name: "大谷 二郎"},
                {id: 9,  name: "磯塚 体制"},
                {id: 10, name:  "井伊 直弼"},
                {id: 11, name:  "伊藤　龍人"},
            ],
            playerToAdd: "",
            players: [
                {id: 1,  name: "酒井 宏樹"},
                {id: 2,  name: "本田 圭佑"},
                {id: 3,  name: "佐々木 朗希"},
                {id: 4,  name: "遠藤 航"},
                {id: 5,  name: "大迫 雄也"},
                {id: 6,  name: "向井 修"},
                {id: 7,  name: "林 圭佑"},
                {id: 8,  name: "大谷 二郎"},
                {id: 9,  name: "磯塚 体制"},
                {id: 10, name:  "井伊 直弼"},
            ],
            selected: "",
        }),
        methods: {
            addPlayer: function() {
                if (this.playerToAdd === "" || this.selected === "") {
                    return;
                }
                this.selected.push(this.playerToAdd);
                this.playerToAdd = "";
            },
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

<style>
    .hover{
        border: rgb(214, 214, 214) 2px solid;
        cursor: pointer;
    }
    .selected{
        border: black 2px solid;
    }
</style>
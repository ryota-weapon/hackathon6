<template>
    <v-footer app>
        <v-row justify="center" class="my-4">
            <!-- <v-col cols="4"></v-col> -->
            <v-col class="px-0" cols="1" align="center" dark v-for="share in shareItems"><v-icon>icon: {{share.icon}} </v-icon></v-col>
            <v-col cols="12" align="center">Copyright © 2022 みんなの選手評価</v-col>
        </v-row>
    </v-footer>
</template>

<script>
    export default {
        data: () => ({
            shareItems: [
                {icon: "mdi-twitter", link: ""},
                {icon: "mdi-facebook", link: ""},
                {icon: "mdi-instagram", link: ""},
            ]
        })
    }
</script>
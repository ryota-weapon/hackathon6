<template>
    <v-footer>
        <v-row justify="center">
            <!-- <v-col cols="4"></v-col> -->
            <v-col class="px-0" cols="1" align="center" dark v-for="share in shareItems"><v-icon>icon: {{share.icon}} </v-icon></v-col>
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
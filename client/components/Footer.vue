<template>
    <v-footer>
        <v-icon dark v-for="share in shareItems">icon: {{share.icon}} </v-icon>
    </v-footer>
</template>

<script>
    export default {
        data: () => ({
            shareItems: [
                {icon: "mdi-twitter", link: ""},
                {icon: "mdi-facebook", link: ""},
                {icon: "", link: ""},
            ]
        })
    }
</script>
<template>
    <v-container>
        <p>Match page</p>
        <p>matchId: {{matchId}}</p>
    </v-container>
</template>

<script>
    export default {
        data: function(){
            return {
                matchId: this.$route.params.id,
            }
        },
        methods: {
        }
    }
</script>


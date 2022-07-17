<template>
    <div class="landing">
        <!-- <v-card elevation="3" height="600px" class="mt-10" align-content="center"> -->
            <v-container>
                <h1 align="center">みんなの選手評価</h1>

                <!-- <v-card > -->
                    <v-row>
                        <v-col align="center">
                            <v-btn 
                                class="text-h2"
                                height="120px"
                                width="300px"
                                to="/app"
                            >
                                はじめる
                            </v-btn>
                        </v-col>
                        <v-col align="center">
                            <v-btn
                                class="text-h2"
                                height="120px"
                                width="300px"
                                to="/signup"
                            >
                                <v-icon>mdi-open-in-new</v-icon> 登録
                            </v-btn>
                        </v-col>
                    </v-row>
                <!-- </v-card> -->
            </v-container>

        <!-- </v-card> -->
    </div>

</template>

<style scoped>
h1{
    font-size: 100px;
}
    .landing{
        height: 100%;
        margin: 0px;
        padding-top: 200px;
        
        opacity: 0.6;
        background-size: cover;
        /* background-image: url("https://img.my-best.com/content_section/choice_component/sub_contents/bc422a80445258218882f28f720dca51.jpg?ixlib=rails-4.2.0&q=70&lossless=0&w=690&fit=max&s=ea9ee8368534b78c72904e677c7332d8"); */
    }
</style>

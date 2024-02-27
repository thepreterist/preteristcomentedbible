<template>
  <v-card>
    <v-layout>
      <!-- <v-system-bar color="deep-purple darken-3"></v-system-bar> -->

      <v-app-bar
        color="primary"
        prominent
      >
        <v-app-bar-nav-icon variant="text" @click.stop="drawer = !drawer"></v-app-bar-nav-icon>

        <v-toolbar-title>Matthew</v-toolbar-title>

        <v-spacer></v-spacer>

        <v-btn variant="text" icon="mdi-magnify"></v-btn>

        <v-btn variant="text" icon="mdi-filter"></v-btn>

        <v-btn variant="text" icon="mdi-dots-vertical"></v-btn>
      </v-app-bar>

      <v-navigation-drawer
        v-model="drawer"
        location="left"
        temporary
      >
        <v-list
          :items="items"
        ></v-list>
      </v-navigation-drawer>

      <v-main style="height: 100vh">

        <v-container fluid class="fill-height">

            <v-row no-gutters class="fill-height">
              <v-col cols="8" class="overflow-y-auto">
                <v-list lines="one" height="1000px">
                    <v-list-item
                      v-for="verse in versicles"
                      :key="verse.verse"
                      :title="verse.verse"
                      :subtitle="verse.text"
                    ></v-list-item>
                  </v-list>
              </v-col>

              <v-col cols="4" class="fill-height">
                <v-sheet class="pa-2 ma-2 fill-height">
                  .v-col-4
                </v-sheet>
              </v-col>
            </v-row>
          </v-container>


      </v-main>
    </v-layout>
  </v-card>
</template>

<script>

  export default {
    data: () => (
      {
      drawer: false,
      group: null,
      versicles: [],
      items: [
        {
          title: 'Matthew',
          value: 'foo',
        },
        {
          title: 'Luke',
          value: 'bar',
        },
        {
          title: 'Mark',
          value: 'fizz',
        },
        {
          title: 'John',
          value: 'buzz',
        },
      ]
    }),

    created() {

      this.getPosts()
      // // watch the params of the route to fetch the data again
      // this.$watch(
      //   // () => this.$route.params,
      //   () => {
      //     this.getPosts()
      //   },
      //   // fetch the data when the view is created and the data is
      //   // already being observed
      //   { immediate: true }
      // )
    },

    methods: {
      getPosts() {
        fetch('http://localhost:1323/kjv/Matthew/24', { mode: "cors"})
          .then(response => response.json())
          .then(data => this.versicles = data.verses)
      }
    },

    watch: {
      group () {
        this.drawer = false
      },
    },
  }
</script>

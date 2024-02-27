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
              <v-col cols="8" class="fill-height overflow-y-auto">
                <!-- <v-pagination :length="book.chapters" :start="chapter"
                v-model.lazy="chapter"></v-pagination> -->

                <!-- <div>{{ chapter }}</div> -->

              <v-bottom-sheet v-model="sheet">
                  <template v-slot:activator="{ props }">
                    <div class="text-center">
                      <v-btn
                        v-bind="props"
                        color="primary"
                        size="large"
                        :text="chapter + ''"
                      ></v-btn>
                    </div>
                  </template>

                  <v-list>
                    <v-list-subheader>Select Chapter</v-list-subheader>

                    <v-list-item
                      v-for="chapt in book.chapters"
                      :key="chapt"
                      :title="chapt"
                      @click="sheet = false; getPosts(chapt)"
                    ></v-list-item>
                  </v-list>

                  <!-- <v-btn
                  v-for="chapt in book.chapters"
                  @click="sheet = false"
                  >
                    {{ chapt }}
                  </v-btn> -->


                </v-bottom-sheet>

                <v-card v-for="verse in versicles" :subtitle="verse.verse" :text="verse.text"></v-card>
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
      book: {},
      chapter: 1,
      chaptersPanelTitle: "",
      sheet: false,
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

      getPosts(selectedChapter) {

        if(!selectedChapter) {
          selectedChapter = 1
        }

        fetch('http://localhost:1323/kjv/Matthew/' + selectedChapter, { mode: "cors"})
          .then(response => response.json())
          .then(data => {
            this.versicles = data.chapter.verses
            this.book = data
            this.chapter = parseInt(data.chapter.chapter)
            this.chaptersPanelTitle = "Chapter " + this.chapter
          })
      }
    },

    watch: {
      group () {
        this.drawer = false
      },
    },
  }
</script>

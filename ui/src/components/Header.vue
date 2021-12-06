<template>
    <v-app-bar
      app
      color="primary"
      title="ScraperUI"
      dark
    >
        <v-app-bar-title
            color="text"
        >
          ScraperUI
        </v-app-bar-title>
        <v-spacer></v-spacer>

        <v-menu
        bottom
        left
      >
        <template v-slot:activator="{ on, attrs }">
          <v-btn
            icon
            color="yellow"
            v-bind="attrs"
            v-on="on"
          >
            <v-icon>mdi-outline-keyboard-arrow-down</v-icon>
          </v-btn>
        </template>

        <v-list>
          <v-list-item
            v-for="(item, i) in items"
            :key="i"
          >
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </v-app-bar>
</template>

<script>
export default {
  name: "Header"
}
</script>

<style scoped>

</style>
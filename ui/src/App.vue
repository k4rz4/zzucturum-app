<template>
  <v-app>

    <Sidebar/>
    <v-main>
    <Header/>

    </v-main>
  </v-app>
</template>

<script>
import Header from './components/Header';
import Sidebar from './components/Sidebar';

export default {
  name: 'App',

  components: {
    Header,
    Sidebar
  },

  data: () => ({
    //
  }),
};
</script>

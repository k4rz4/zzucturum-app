import Vue from 'vue';
import Vuetify from 'vuetify/lib/framework';
import colors from 'vuetify/lib/util/colors'

Vue.use(Vuetify);

export default new Vuetify({
    theme: {
    themes: {
      light: {
        primary: colors.teal.darken3,
        secondary: colors.grey.darken3,
        secondary2: colors.grey.darken4,
        accent: colors.indigo.base,
        text: colors.shades.white
      },
    },
  },
});

import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import 'vuetify/styles'

export default createVuetify({
  components,
  directives,
  theme: {
    defaultTheme: 'light',
    themes: {
      dark: {
        dark: true,
        colors: {
          background: '#0a0a0f',
          surface: '#12121a',
          'surface-variant': '#1a1a26',
          primary: '#7c6af7',
          'primary-darken-1': '#5a4fd6',
          secondary: '#f7c26a',
          accent: '#f76a8a',
          error: '#ff5252',
          info: '#2196F3',
          success: '#4CAF50',
          warning: '#FB8C00',
        }
      },
      light: {
        dark: false,
        colors: {
          background: '#f0efff',
          surface: '#ffffff',
          'surface-variant': '#f5f4ff',
          primary: '#5c4df0',
          'primary-darken-1': '#4338ca',
          secondary: '#f59e0b',
          accent: '#ec4899',
          error: '#ef4444',
          success: '#22c55e',
        }
      }
    }
  },
  defaults: {
    VBtn: {
      rounded: 'lg',
      elevation: 0,
    },
    VTextField: {
      variant: 'outlined',
      density: 'comfortable',
    },
    VCard: {
      rounded: 'xl',
      elevation: 0,
    }
  }
})

/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./src/**/*.{html,templ,go}",
  ],
  theme: {
    extend: {
      colors: {
        "c-bg": "var(--c-bg)",
        "c-on-bg": "var(--c-on-bg)",
      }
    },
  },
  corePlugins: {
    container: false
  },
  plugins: [
    function({ addComponents }) {
      addComponents({
        '.container': {
          maxWidth: '95%',
          margin: "auto",
          '@screen sm': {
            maxWidth: '640px',
          },
          '@screen md': {
            maxWidth: '768px',
          },
          '@screen lg': {
            maxWidth: '1024px',
          },
          '@screen xl': {
            maxWidth: '1280px',
          },
        }
      })
    }
  ],
}


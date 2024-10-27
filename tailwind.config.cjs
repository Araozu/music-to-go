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
  plugins: [],
}


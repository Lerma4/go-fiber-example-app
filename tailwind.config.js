/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    'ui/html/*.templ',
  ],
  // darkMode: 'class',
  theme: {
    extend: {
      fontFamily: {
        mono: ['Courier Prime', 'monospace'],
      }
    },
  },
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["light", "dark", "cupcake"],
  },
}


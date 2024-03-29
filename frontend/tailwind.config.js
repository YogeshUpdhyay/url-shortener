/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'blue': '#2161E0',
        'red': '#E93266',
        'black': '#101044',
        'grey': '#54547A'
      },
      fontFamily: {
        themeFont: ['themeFont', 'sans-serif'],
      },
    },
  },
  plugins: [],
}

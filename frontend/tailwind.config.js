/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        background: '#0B0F19',
        surface: '#151C2C',
        primary: '#6366f1',
        primaryHover: '#4f46e5',
        textMain: '#E2E8F0',
        textMuted: '#94A3B8'
      }
    },
  },
  plugins: [],
}

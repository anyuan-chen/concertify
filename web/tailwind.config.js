/** @type {import('tailwindcss').Config} */
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}",
  ],
  prefix: "",
  important: false,
  theme: {
    fontFamily: {
      sans: ["Neue Montreal", "ui-sans-serif", "system-ui"],
    },
  },
  variants: {},
  plugins: [require("@tailwindcss/forms")],
};

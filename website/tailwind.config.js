/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{js,jsx,ts,tsx}"],
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

const defaultTheme = require("tailwindcss/defaultTheme");

module.exports = {
  content: ["./src/**/*.{html,ts}"],
  darkMode: "class",
  theme: {
    extend: {
      fontFamily: {
        inter: ["Inter", ...defaultTheme.fontFamily.sans],
        "inter-var": ["Inter var", ...defaultTheme.fontFamily.sans],
      },
    },
  },
  plugins: [],
};

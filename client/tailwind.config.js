/** @type {import('tailwindcss').Config} */
module.exports = {
  prefix: 'tw-',
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        "custom-primary": "#2FC4B2",
        "custom-primary-dark": "#259C8E",
        "custom-secondary": "#E7E7E7",
        "custom-secondary-dark": "#B8B8B8",
        "custom-danger":"#E83939",
        "custom-danger-dark":"#B92D2D",
        "custom-warning":"#FF9900",
        "custom-warning-dark":"#CC7A00",
        "custom-success":"#469F74",
        "custom-success-dark":"#387F5C",
      }
    },
  },
  plugins: [],
}

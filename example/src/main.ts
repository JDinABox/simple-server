import Alpine from "alpinejs";

import "./index.css";

window.Alpine = Alpine;

let theme = {
  init() {
    const dark = localStorage.getItem("theme-dark") || "";
    this.dark =
      dark === ""
        ? window.matchMedia("(prefers-color-scheme: dark)").matches
        : Boolean(JSON.parse(dark)).valueOf();
  },
  dark: true,
  toggle() {
    this.dark = !this.dark;
    localStorage.setItem("theme-dark", JSON.stringify(this.dark));
  },
};

Alpine.store("theme", theme);
Alpine.store("title", "Index");

Alpine.start();

import { defineConfig } from "vite";

import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [],
  build: {
    lib: {
      entry: path.resolve(__dirname, "lib/main.ts"),
      name: "simple-server",
      fileName: (format) => `ss.${format}.js`,
    },
    outDir: path.resolve(__dirname, "./assets"),
  },
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./lib"),
    },
  },
});

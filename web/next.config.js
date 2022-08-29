/** @type {import('next').NextConfig} */
module.exports = {
  reactStrictMode: true,
  async rewrites() {
    return [
      {
        source: "",
        destination: "http://localhost:8000/", // Proxy to Backend
      },
    ];
  },
};

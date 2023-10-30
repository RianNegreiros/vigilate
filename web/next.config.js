/** @type {import('next').NextConfig} */
const nextConfig = {
  Headers: [
    {
      key: 'Access-Control-Allow-Origin',
      value: 'https://vigilate.vercel.app'
    },
  ],
}

module.exports = nextConfig

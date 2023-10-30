/** @type {import('next').NextConfig} */
const nextConfig = {
  async headers() {
    return [
      {
        source: '/(.*)',
        headers: [
          {
            key: 'Access-Control-Allow-Origin',
            value: 'https://vigilate.vercel.app',
          },
        ],
      },
    ]
  }
}

module.exports = nextConfig

/** @type {import('next').NextConfig} */
const trustedOrigins = process.env.TRUSTED_ORIGINS
const nextConfig = {
  Headers: [
    {
      key: 'Access-Control-Allow-Origin',
      value: trustedOrigins,
    },
  ],
}

module.exports = nextConfig

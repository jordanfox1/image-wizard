import type { Metadata } from 'next'
import './globals.css'
import '@mantine/core/styles.css';
import { MantineProvider, ColorSchemeScript } from '@mantine/core';

export const metadata: Metadata = {
  title: 'Img Switch',
  description: 'Free online image conversion tool',
  authors: { name: 'Jordan Fox', url: 'http://jordan-fox-developer.s3-website-ap-southeast-2.amazonaws.com/' },
  icons: { icon: [{ url: '/images/wand.webp', href: '/images/wand.webp' }] },
}

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <head>
        <ColorSchemeScript />
      </head>
      <body>
        <MantineProvider defaultColorScheme="light" >{children}</MantineProvider>
      </body>
    </html>
  );
}

import type { Metadata } from 'next'
import './globals.css'
import '@mantine/core/styles.css';
import { MantineProvider, ColorSchemeScript, createTheme, rem  } from '@mantine/core';

export const metadata: Metadata = {
  title: 'Image Wizard',
  description: 'Free online image conversion tool',
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

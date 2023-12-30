import astroSingleFile from 'astro-single-file'
import { defineConfig } from 'astro/config'

export default defineConfig({
  integrations: [astroSingleFile()]
})

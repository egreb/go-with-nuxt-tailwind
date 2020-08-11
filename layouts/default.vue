<template>
  <Nuxt />
</template>
<script>
import { show, toggle } from '@/hooks/useMenu'
export default {
  setup() {
    return { show, toggle }
  },
}
</script>
<style>
html {
  font-family: 'Source Sans Pro', -apple-system, BlinkMacSystemFont, 'Segoe UI',
    Roboto, 'Helvetica Neue', Arial, sans-serif;
  font-size: 16px;
  word-spacing: 1px;
  -ms-text-size-adjust: 100%;
  -webkit-text-size-adjust: 100%;
  -moz-osx-font-smoothing: grayscale;
  -webkit-font-smoothing: antialiased;
  box-sizing: border-box;
}

*,
*::before,
*::after {
  box-sizing: border-box;
  margin: 0;
}

#app {
  height: 100vh;
  display: grid;
  grid-template-rows: auto 1fr auto;
}

main.content {
  display: grid;
  grid-template-columns: 1fr;
}

@media (min-width: 1024px) {
  main.content {
    grid-template-columns: minmax(150px, 15%) 1fr;
  }
  main.content.active {
    grid-template-columns: 1fr;
  }
}
</style>

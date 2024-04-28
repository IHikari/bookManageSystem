import { createPinia } from 'pinia'
import persist from 'pinia-plugin-persistedstate'

const pinia = createPinia()
pinia.use(persist)

export default pinia

// import { useUserStore } from './modules/user.js'
// export { useUserStore }
// import { useCountStore } from './modules/counter.js'
// export { useCountStore }

export * from './modules/user.js'
export * from './modules/counter.js'

<script>
export default {
  name: 'App',
  data() {
    return {
      currentEnv: 'in',
      environments: [],
      loading: true,
      error: null
    }
  },
  methods: {
    initializeEnvironment() {
      // Récupérer l'environnement depuis l'URL
      const pathParts = this.$route.path.split('/')
      if (pathParts.length > 1 && pathParts[1]) {
        this.currentEnv = pathParts[1]
      }
    }
  },
  async created() {
    this.initializeEnvironment()
    try {
      const response = await fetch('/api/environments')
      if (!response.ok) {
        throw new Error('Erreur lors de la récupération des environnements')
      }
      this.environments = await response.json()
    } catch (error) {
      this.error = error.message
      console.error('Erreur:', error)
    } finally {
      this.loading = false
    }
  },
  watch: {
    currentEnv(newEnv) {
      // Rediriger vers la même page avec le nouvel environnement
      const currentRoute = this.$route.name
      this.$router.push(`/${newEnv}/${currentRoute.toLowerCase()}`)
    },
    '$route': {
      immediate: true,
      handler(to) {
        const pathParts = to.path.split('/')
        if (pathParts.length > 1 && pathParts[1]) {
          this.currentEnv = pathParts[1]
        }
      }
    }
  }
}
</script>

<template>
  <div class="min-h-screen bg-gray-900">
    <!-- Main Navigation -->
    <nav class="fixed top-0 left-0 w-full z-50 bg-gray-800 border-b border-gray-700">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center">
          <div class="flex items-center space-x-8">
            <h1 class="text-xl font-bold text-white">ESB Console</h1>
            <router-link 
              :to="`/${currentEnv}/dashboard`"
              class="inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium"
              :class="[
                $route.name === 'Dashboard' 
                  ? 'border-indigo-500 text-white' 
                  : 'border-transparent text-gray-300 hover:border-gray-600 hover:text-white'
              ]"
            >
              <svg class="mr-2 h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
              </svg>
              Dashboard
            </router-link>

            <router-link 
              :to="`/${currentEnv}/jms`"
              class="inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium"
              :class="[
                $route.name === 'JMS' 
                  ? 'border-indigo-500 text-white' 
                  : 'border-transparent text-gray-300 hover:border-gray-600 hover:text-white'
              ]"
            >
              <svg class="mr-2 h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7v8a2 2 0 002 2h6M8 7V5a2 2 0 012-2h4.586a1 1 0 01.707.293l4.414 4.414a1 1 0 01.293.707V15a2 2 0 01-2 2h-2M8 7H6a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2v-2" />
              </svg>
              JMS
            </router-link>

            <router-link 
              :to="`/${currentEnv}/audit`"
              class="inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium"
              :class="[
                $route.name === 'Audit' 
                  ? 'border-indigo-500 text-white' 
                  : 'border-transparent text-gray-300 hover:border-gray-600 hover:text-white'
              ]"
            >
              <svg class="mr-2 h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
              </svg>
              Audit
            </router-link>

            <router-link 
              :to="`/${currentEnv}/admin`"
              class="inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium"
              :class="[
                $route.name === 'Admin' 
                  ? 'border-indigo-500 text-white' 
                  : 'border-transparent text-gray-300 hover:border-gray-600 hover:text-white'
              ]"
            >
              <svg class="mr-2 h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
              Admin
            </router-link>
          </div>

          <div class="flex items-center">
            <select
              id="env-select"
              v-model="currentEnv"
              class="block w-full pl-3 pr-10 py-2 text-base border-gray-600 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm rounded-md bg-gray-700 text-gray-200"
              :disabled="loading"
            >
              <option v-if="loading" value="">Chargement...</option>
              <option v-else-if="error" value="">Erreur de chargement</option>
              <option v-else v-for="env in environments" :key="env.env" :value="env.env">
                {{ env.title }}
              </option>
            </select>
          </div>
        </div>
      </div>
    </nav>

    <!-- Main Content -->
    <main class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8 pt-24">
      <div class="px-4 py-6 sm:px-0">
        <div class="bg-gray-800 rounded-lg shadow-sm p-6">
          <router-view></router-view>
        </div>
      </div>
    </main>
  </div>
</template>


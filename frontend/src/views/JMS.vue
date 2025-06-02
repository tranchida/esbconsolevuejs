<template>
  <div class="jms">
    <div class="mb-8">
      <h1 class="text-2xl font-bold text-white">
        JMS
      </h1>
      <p class="mt-2 text-sm text-gray-400">
        Environnement: <span class="font-medium text-indigo-400">{{ $route.params.env.toUpperCase() }}</span>
      </p>
    </div>

    <!-- Filtres -->
    <div class="bg-gray-700 shadow rounded-lg mb-6">
      <div class="px-4 py-5 sm:p-6">
        <div class="grid grid-cols-1 gap-4 sm:grid-cols-3">
          <div>
            <label for="queue" class="block text-sm font-medium text-gray-300">Queue</label>
            <select
              id="queue"
              class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-600 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm rounded-md bg-gray-800 text-gray-200"
            >
              <option>Toutes les queues</option>
            </select>
          </div>
          <div>
            <label for="status" class="block text-sm font-medium text-gray-300">Statut</label>
            <select
              id="status"
              class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-600 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm rounded-md bg-gray-800 text-gray-200"
            >
              <option>Tous les statuts</option>
            </select>
          </div>
          <div>
            <label for="search" class="block text-sm font-medium text-gray-300">Recherche</label>
            <input
              type="text"
              id="search"
              class="mt-1 block w-full border border-gray-600 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm bg-gray-800 text-gray-200"
              placeholder="Rechercher..."
            />
          </div>
        </div>
      </div>
    </div>

    <!-- Tableau des messages -->
    <div class="bg-gray-700 shadow rounded-lg overflow-hidden">
      <table class="min-w-full divide-y divide-gray-600">
        <thead class="bg-gray-800">
          <tr>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-300 uppercase tracking-wider">
              ID
            </th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-300 uppercase tracking-wider">
              Queue
            </th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-300 uppercase tracking-wider">
              Statut
            </th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-300 uppercase tracking-wider">
              Date
            </th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-300 uppercase tracking-wider">
              Actions
            </th>
          </tr>
        </thead>
        <tbody class="bg-gray-700 divide-y divide-gray-600">
          <tr>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-300">
              Aucun message
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-300">
              -
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-300">
              -
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-300">
              -
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-300">
              -
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  name: 'JMS'
}
</script> 
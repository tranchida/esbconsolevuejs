<template>
    <div class="admin">
        <h1 class="text-2xl font-bold text-gray-800 mb-4">
            Admin - Environnement: {{ $route.params.env.toUpperCase() }}
        </h1>
    </div>
</template>

<script>
export default {
    name: 'Admin'
}
</script>
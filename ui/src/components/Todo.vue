<template>
    <div>
        {{ todo.content }}
    </div>
</template>

<script>
    export default {
        name: 'Todo',
        props: [
            "todo",
        ]
    };
</script>

<style scoped>
</style>
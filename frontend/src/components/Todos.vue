<template>
    <div>
        <h2>TODO List</h2>
        <ul>
            <li v-bind:key="todo.id" v-for="todo in todos.slice().reverse()">
                <Todo v-bind:todo="todo" />
            </li>
        </ul>
    </div>
</template>

<script>
    import Todo from './Todo.vue';

    export default {
        name: 'Todos',
        components: {
            Todo
        },
        props: [
            "todos",
        ]
    };
</script>

<style scoped>
</style>
<template>
    <div>
        <Todos v-bind:todos="todos" />
        <AddTodo v-on:new-todo="newTodo" />
    </div>
</template>

<script>
    import Todos from './components/Todos.vue';
    import AddTodo from './components/AddTodo.vue';

    export default {
        name: 'app',
        components: {
            Todos,
            AddTodo
        },
        data() {
            return {
                todos: [],
            };
        },
        methods: {
            newTodo(todo) {
                this.todos = [...this.todos, todo];
            }
        }
    };
</script>
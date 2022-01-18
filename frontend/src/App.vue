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
            },
            fetchTodos() {
                fetch(`http://${window.location.hostname}:8000/todos`, {
                    method: 'GET'
                })
                .then(r => r.json())
                .then(d => this.todos = d.map(e => {
                    return {
                        id: e.Id,
                        content: e.Content
                    }
                }));
            }
        },
        created() {
            this.fetchTodos();
        }
    };
</script>
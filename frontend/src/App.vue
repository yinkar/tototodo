<template>
    <div class="container">
        <h1 data-test-header>Tototo Do</h1>
        <h2>TODO List</h2>
        <AddTodo v-on:new-todo="newTodo" />
        <Todos v-bind:todos="todos" />
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

<style lang="stylus">
    *
        margin 0
        padding 0

    body
        background-color #eee
        font-family sans-serif

    .container
        background-color #0b9
        width 50rem
        margin 60px auto
        box-shadow 3px 3px 10px rgba(0, 0, 0, .3)
        border-radius 40px
        padding 30px 60px
        box-sizing border-box

    .container > h2
        margin 10px 0
        text-align: center
</style>
<template>
    <div>
        <ul class="todos">
            <Todo 
                v-bind:todo="todo" 
                v-bind:key="todo.id" 
                v-for="todo in todos.slice().reverse()"
            />
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

<style lang="stylus">
    .todos
        list-style none
        margin 10px 0

</style>
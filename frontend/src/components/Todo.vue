<template>
    <div class="todo" :data-id="todo.id">
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

<style lang="stylus">
    .todo
        padding 10px 10px
        background-color #fff4
        color #020
        font-size .8em
        margin: 6px 0
        border-radius 2px
        transition: 100ms ease-in-out all

        &:hover
            background-color #fff8

</style>
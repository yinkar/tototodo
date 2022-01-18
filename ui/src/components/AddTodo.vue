<template>
    <div>
        <form @submit="addTodo">
            <input type="text" name="content" v-model="content" />
            <button type="submit">Add</button>
        </form>
    </div>
</template>

<script>
    import { v4 as uuidv4 } from 'uuid';

    export default {
        name: 'AddTodo',
        data() {
            return {
                content: ''
            }
        },
        methods: {
            addTodo(e) {
                e.preventDefault();

                if (this.content.length === 0) return;

                this.$emit('new-todo', {
                    id: uuidv4(),
                    content: this.content
                });
                this.content = '';
            }
        }
    };
</script>

<style scoped>
</style>
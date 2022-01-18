<template>
    <div>
        <form @submit="addTodo">
            <input type="text" name="content" id="content" v-model="content" />
            <button type="submit">Add</button>
        </form>
    </div>
</template>

<script>
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

                const todo = {
                    content: this.content
                };

                fetch(`http://${window.location.hostname}:8000/todos`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(todo)
                })
                .then(r => r.json())
                .then(d => {
                    this.$emit('new-todo', todo);
                    this.content = '';
                    document.querySelector('#content').classList.remove('error');
                })
                .catch(err => {
                    console.error(err);
                    document.querySelector('#content').classList.add('error');
                });
            }
        }
    };
</script>

<style scoped>
    #content.error {
        border: 2px solid red;
    }
</style>
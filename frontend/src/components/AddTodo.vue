<template>
    <div>
        <form class="add-todo" @submit="addTodo">
            <input type="text" name="content" id="content" v-model="content" />
            <button type="submit" id="add">Add</button>
        </form>
    </div>
</template>

<script>
export default {
    name: "AddTodo",
    data() {
        return {
            content: "",
        };
    },
    methods: {
        addTodo(e) {
            e.preventDefault();

            if (this.content.length === 0) return;

            const todo = {
                content: this.content,
            };

            fetch(`http://${window.location.hostname}:8000/todos`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(todo),
            })
                .then((r) => r.json())
                .then(() => {
                    this.$emit("new-todo", todo);
                    this.content = "";
                    document
                        .querySelector("#content")
                        .classList.remove("error");
                })
                .catch((err) => {
                    console.error(err);
                    document.querySelector("#content").classList.add("error");
                });
        },
    },
};
</script>

<style lang="stylus">
    .add-todo
        width 300px
        margin 30px auto
        text-align center

    #content
        border none
        background-color #fff8
        border-radius 20px 0 0 20px
        padding 10px
        height 30px
        box-sizing border-box
    
    #add
        border none
        padding 10px
        border-radius 0 20px 20px 0
        font-size .6em
        height 30px
        box-sizing border-box
        background-color #fffc
        cursor pointer


</style>
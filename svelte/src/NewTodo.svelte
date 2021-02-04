<script lang="ts">
    import type { TodosClient } from "./gen/TodosServiceClientPb";
    import type { Todo } from "./gen/todos_pb";
    import todos from "./store";
    import { newTodoRequest } from "./Client";

    import Button from "@smui/button";
    import Textfield from "@smui/textfield";

    export let client: TodosClient;
    let createResponse: Promise<Todo> = Promise.resolve(null);
    let value: string = "";

    const save = () => {
        createResponse = client.createTodo(newTodoRequest(value), null);
        createResponse.then((todo) => todos.addTodo(todo));
        value = "";
    };
</script>

<Textfield bind:value label="Enter A Todo" />
<p>{value}</p>
<Button on:click={save}>Save</Button>
{#await createResponse}
    <p>waiting</p>
{:catch error}
    <p>{error.message}</p>
{/await}

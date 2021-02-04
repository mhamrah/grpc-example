<script lang="ts">
    import DataTable, { Head, Body, Row, Cell } from "@smui/data-table";
    import App from "./App.svelte";
    import type { TodosClient } from "./gen/TodosServiceClientPb";
    import { ListTodosRequest } from "./gen/todos_pb";
    import NewTodo from "./NewTodo.svelte";
    import todos from "./store";

    export let client: TodosClient;
    const listResponse = client.listTodos(new ListTodosRequest(), null);
</script>

{#await listResponse}
    <p>waiting</p>
{:then result}
    <DataTable table$aria-label="Todos">
        <Head>
            <Row>
                <Cell>Id</Cell>
                <Cell>Title</Cell>
            </Row>
        </Head>
        <Body>
            {#each $todos.concat(result.getTodosList()) as item}
                <p />
                <Row>
                    <Cell>{item.getId()}</Cell>
                    <Cell>{item.getTitle()}</Cell>
                </Row>
            {/each}
        </Body>
    </DataTable>
{:catch error}
    <p>{error.message}</p>
{/await}

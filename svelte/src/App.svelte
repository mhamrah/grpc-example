<script lang="ts">
  import client from "./Client";
  import NewTodo from "./NewTodo.svelte";
  import Todos from "./Todos.svelte";
  import "./App.scss";
</script>

<main>
  <div class="mdc-layout-grid">
    <div class="mdc-layout-grid__inner">
      <div class="mdc-layout-grid__cell"><NewTodo {client} /></div>
      <div class="mdc-layout-grid__cell"><Todos {client} /></div>
    </div>
  </div>
</main>

<style>
</style>

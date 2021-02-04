import { writable } from 'svelte/store';
import type { Todo } from './gen/todos_pb';

const { subscribe, update } = writable<Todo[]>([]);

const addTodo = (todo: Todo) => {
    update(current => {
        return [todo, ...current]
    })
}

export default {
    subscribe,
    addTodo
}

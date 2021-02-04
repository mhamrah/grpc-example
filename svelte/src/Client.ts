import { TodosClient } from "./gen/TodosServiceClientPb";
import { CreateTodoRequest, ListTodosRequest, Todo } from "./gen/todos_pb";

export default new TodosClient("http://localhost:9000");


export const newTodoRequest = (title: string): CreateTodoRequest => {
    const todo = new Todo();
    todo.setTitle(title);
    const req = new CreateTodoRequest();
    req.setTodo(todo);
    return req;
};

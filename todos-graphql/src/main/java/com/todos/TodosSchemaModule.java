// Copyright 2017 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package com.todos.graphqlserver;

import com.google.api.graphql.rejoiner.Query;
import com.google.api.graphql.rejoiner.SchemaModule;
import todos.TodosGrpc;
import todos.TodosOuterClass.GetTodoRequest;
import todos.TodosOuterClass.ListTodosRequest;
import todos.TodosOuterClass.ListTodosResponse;
import todos.TodosOuterClass.Todo;

/** A GraphQL {@link SchemaModule} backed by a gRPC service. */
final class TodosSchemaModule extends SchemaModule {
  @Query("todo")
  Todo getTodo(GetTodoRequest request, TodosGrpc.TodosBlockingStub client) {
    return client.getTodo(request);
  }
  @Query("todos")
  ListTodosResponse listTodos(ListTodosRequest request, TodosGrpc.TodosBlockingStub client) {
    return client.listTodos(request);
  }
}

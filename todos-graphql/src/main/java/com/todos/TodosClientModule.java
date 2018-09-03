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

import com.google.inject.AbstractModule;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import todos.TodosGrpc;

/** Binds a stub for the {@link TodosGrpc} service. */
final class TodosClientModule extends AbstractModule {

  private static final String HOST = "todos";
  private static final int PORT = 50051;

  @Override
  protected void configure() {
    ManagedChannel channel =
        ManagedChannelBuilder.forAddress(HOST, PORT).usePlaintext(true).build();
    bind(TodosGrpc.TodosBlockingStub.class).toInstance(TodosGrpc.newBlockingStub(channel));
  }
}

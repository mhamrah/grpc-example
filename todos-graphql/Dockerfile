FROM gradle:4.10.0-jdk10-slim as build
USER root

ADD . .

RUN gradle --no-daemon installDist

FROM openjdk:10-jre-slim

WORKDIR /app

COPY --from=build /home/gradle/build/install/todos /app

ADD src/main/resources/index.html /app/src/main/resources/index.html

EXPOSE 8000

ENTRYPOINT [ "/app/bin/todo-graphqlserver" ]

# Описание
Проект написан в рамках изучения курса по gRPC.

# Настройка
Команды для установки зависимостей и генерации кода:
```bash
1. go get -u google.golang.org/grpc
2. go get -u google.golang.org/protobuf
3. protoc --go_out=plugins=grpc:. calc-api.proto
```
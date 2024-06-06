# Оператор марсохода

## Описание

Программа позволяет управлять имитацией марсохода через пользовательский ввод в консоле

- F – проехать на одну единицу вперёд,
- B – проехать на одну единицу назад,
- L – повернуть налево,
- R – повернуть направо.

пример ввода `FFLBFRLBBFFRRBBLFR`
после ввода команды на вывод подадуться координаты, или сообщение об ошибке (в файле [env](./.env_example) можно посмотреть вероятность возникновения ошибки `ROVER_ERR_PERCEN`)

Паралельно программе общения с пользователем, так же запускается cron-job для переодического пинга марсохода

## Запуск

Локальный

```sh
# сам запуск (для windows придется настроить git-bash)
make local-run
```

через docker-compose

- Создать файл `.env` (достаточно скопировать данные файла [.env_example](.env_example)), чтобы при поднятии контейнера прокидывались нужные переменные окружения
- выполинть команду
  ```sh
  make docker-run
  ```

### Структура

```sh
├── cmd -- точки входа в приложение
├── pkg
│   ├── application -- сборка компонентов приложения / ручной DI
│   ├── controller -- обработчик для взаимодействия с марсоходами
│   ├── clients -- клиенты для взаимодействия (пока только консоль)
│   ├── crons -- фоновые задачи
│   ├── config -- конфигурация
│   ├── logger
│   └── rovers -- описания и реализация (пока только тестового) марсохода
```

### Что надо доработать

- Не хватает доработки по оптимизации пути (строки)

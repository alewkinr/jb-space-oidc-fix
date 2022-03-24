# ODIC reverse-proxy


## Описание
Сервис работает, как reverse-proxy и подставляет access_token из query в `Authorization` заголовок для исправления
OIDC совместимости продукта Jet Brains Space

## Зависимости
- golang 1.16
- [Taskfile](https://taskfile.dev/#/)

## Запуск
- Для запуска откройте терминал в корне репоизтория и выполните команду `task build && ./bin/service`

## Сборка
- Для сборки вызовите команду `task build`
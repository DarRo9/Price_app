Price Recognition App

Приложение для распознавания цен на изображениях и расчета цены за килограмм.
## Пример

После локального запуска было загружено фото ценника по localhost:8080 в браузере. В результате были распознаны вес и цена, а также была расчитана цена за кг

<img width="703" height="594" alt="Screenshot from 2025-07-17 01-58-15" src="https://github.com/user-attachments/assets/495ddc89-302a-43dd-839d-070836cfdf96" />
<img width="538" height="341" alt="Screenshot from 2025-07-17 01-43-09" src="https://github.com/user-attachments/assets/5bfa76e7-035a-4fd7-92a9-a4a976c27f21" />


```
backend/
├── cmd/main.go              # Точка входа
├── internal/
│   ├── domain/              # Доменные модели
│   ├── usecase/             # Бизнес-логика
│   ├── repository/          # Интерфейсы и реализации репозиториев
│   └── delivery/            # HTTP handlers и роутинг
├── configs/                 # Конфигурация
├── web/                     # Веб-интерфейс
└── pkg/                     # Общие пакеты
```

## Возможности

- Распознавание текста на русском языке с изображений
- Автоматическое извлечение цены и веса
- Расчет цены за килограмм
- Оценка уверенности распознавания
- Современный веб-интерфейс
- Docker контейнеризация

## Установка и запуск

### С Docker (рекомендуется)

```bash
# Сборка и запуск
make build
make run

# Или напрямую
docker-compose up -d
```

### Локальная разработка

```bash
# Установка зависимостей
sudo apt-get install tesseract-ocr tesseract-ocr-rus libtesseract-dev libleptonica-dev

# Запуск
make dev
```

## Использование

1. Откройте http://localhost:8080
2. Загрузите изображение с ценником
3. Получите распознанный текст и расчет цены за кг

## API

- `POST /api/v1/upload` - Загрузка изображения
- `GET /api/v1/health` - Проверка здоровья сервиса

## Управление

```bash
make build    # Сборка
make run      # Запуск
make stop     # Остановка
make logs     # Просмотр логов
make clean    # Очистка
```

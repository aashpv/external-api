<div align="center">
  <h1 align="center">External API</h1>
  <h3>API for testing <a href="https://github.com/aashpv/song-lib">Song Library API</a></h3>
</div>

<br/>

External API — это API для тестирования API добавления в [Song Library API](https://github.com/aashpv/song-lib).

## Приступая к работе

### Предварительные требования

Вот что вам нужно для запуска:

- Go (version >= 18)
- PostgreSQL Database

### 1. Склонируйте репозиторий

```shell
git clone https://github.com/aashpv/external-api
```

### 2. Настройте конфигурацию

Измените local.yaml файл, используя [local.yaml](local.yaml) как шаблон. Укажите параметры базы данных и другие настройки.

### 3. Создайте БД и таблицу
```sql
CREATE DATABASE namebd;
```
```sql
CREATE TABLE IF NOT EXISTS songs_info (
    id SERIAL PRIMARY KEY,
    group_name VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    release_date VARCHAR(255) NOT NULL,
    text TEXT NOT NULL,
    link VARCHAR(255) NOT NULL
);
```

### 4. Добавьте несколько(по желанию) полей в БД

```sql
INSERT INTO songs_info(group_name, name, release_date, text, link)
VALUES (
  'Muse',
  'Supermassive Black Hole',
  '2006-07-16',
  'Ooh baby, don''t you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight',
  'https://www.youtube.com/watch?v=Xsp3_a-PMTw'
);
```

### 5. Запустите сервер

```shell
cd external-api\cmd
go run main.go
```

## Основные маршруты API

- **GET /info** - Получение дополнительной информации(release_date, text, link) о песне.


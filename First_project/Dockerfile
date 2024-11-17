# Используем базовый образ PostgreSQL версии 13
FROM postgres:13

# Устанавливаем переменные окружения для пользователя и пароля
ENV POSTGRES_USER=Denis
ENV POSTGRES_PASSWORD=1
ENV POSTGRES_DB=First_project

# Определяем тома для хранения данных
VOLUME ["/var/lib/postgresql/data", "/path/to/your/second/volume"]

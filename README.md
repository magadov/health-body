# Healthy Body

API для управления планами тренировок и питания, подписками и отзывами.

## Описание

Backend API для платформы здорового образа жизни. Позволяет пользователям приобретать планы тренировок и питания, оформлять подписки, оставлять отзывы и рассчитывать индекс массы тела.

## Стек технологий

![Go](https://img.shields.io/badge/Go-1.25.1-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-Web%20Framework-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![GORM](https://img.shields.io/badge/GORM-ORM-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Database-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![Swagger](https://img.shields.io/badge/Swagger-API%20Docs-85EA2D?style=for-the-badge&logo=swagger&logoColor=black)
![gomail](https://img.shields.io/badge/gomail-Email-FF6B6B?style=for-the-badge)

                    HTTP
┌────────┐ ───────────────────────────────┐
│ Client │                                │
└────────┘                                v

              ┌─────────────────────────────────┐
              │        Presentation Layer       │
              │        (Gin Handlers/API)       │
              └───────────────┬─────────────────┘
                              │
                              v
              ┌─────────────────────────────────┐
              │         Business Layer          │
              │                                 │
              │  User Service                   │
              │  Exercise Plan Service          │
              │  Meal Plan Service              │
              │  Subscription Service           │
              │  Reviews Service                │
              │  Category Service               │
              │  BMI Service                    │
              │  Notification Service           │
              └───────────────┬─────────────────┘
                              │
                              v
              ┌─────────────────────────────────┐
              │          Data Layer             │
              │     (Repositories / GORM)       │
              └───────────────┬─────────────────┘
                              │
                              v
              ┌─────────────────────────────────┐
              │           PostgreSQL            │
              │   Users, Plans, Reviews, etc.   │
              └─────────────────────────────────┘

```

## Основные возможности

- Управление пользователями
- Категории планов тренировок и питания
- Планы тренировок с упражнениями
- Планы питания с элементами
- Подписки
- Отзывы
- Расчет BMI (индекс массы тела)
- Email уведомления

## Разработчики

- [Висхан Магомадов](https://github.com/magadov)
- [Юсуп Альсиев](https://github.com/Alsiev)
- [Джабраил Техиев](https://github.com/wiwiieie011)

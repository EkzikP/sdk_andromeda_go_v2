# SDK для ПО "Центр охраны"

Данный SDK разработан для получения информации по API ПО "Центр охраны"
https://www.cnord.ru/security-center

В SDK реализованы методы:

- GetSites - Получить объект по номеру или идентификатору
- GetCustomers - Получить список ответственных лиц объекта
- GetCustomer - Получить ответственное лицо по идентификатору
- PostCheckPanic - Начать проверку КТС
- GetCheckPanic - Получить результат проверки КТС
- GetUsersMyAlarm - Получить список пользователей MyAlarm объекта
- PutChangeUserMyAlarm - Предоставить/забрать доступ пользователю к MyAlarm
- PutChangeKTSUserMyAlarm - Модифицировать право на использование виртуальной КТС
- GetUserObjectMyAlarm - Получить список объектов пользователя MyAlarm
- GetParts - Получить список разделов объекта
- GetZones - Получить список шлейфов объекта
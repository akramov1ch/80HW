# Uyga vazifa: PostgreSQL va Go yordamida asosiy REST API serverni implementatsiya qilish va muloyim to'xtatish (graceful shutdown)

## Maqsad
`PostgreSQL` ma'lumotlar bazasidan foydalanadigan, `CRUD` amallarini bajaradigan, va to'xtash signallarini (`SIGINT` va `SIGTERM`) qabul qilganda muloyim to'xtaydigan asosiy `REST API` serverni yaratish.

### Vazifalar

1. **Go-da REST API serverni yaratish:**
   - `items` kabi resurs uchun `CRUD` amallarini implementatsiya qilish.
   - `items`ni saqlash uchun `PostgreSQL` ma'lumotlar bazasidan foydalanish.
   - `SIGINT` va `SIGTERM` signallarini qabul qilganda muloyim to'xtashni amalga oshirish.

2. **Endpointlar:**
   - `GET /items`: Elementlar ro'yxatini qaytarish.
   - `POST /items`: Yangi element yaratish. 
   - `GET /items/{id}`: ID bo'yicha ma'lum bir elementni qaytarish.
   - `PUT /items/{id}`: ID bo'yicha ma'lum bir elementni yangilash.
   - `DELETE /items/{id}`: ID bo'yicha ma'lum bir elementni o'chirish.







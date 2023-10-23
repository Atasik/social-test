# social-test
Тестовый кейс

## Апи
Апи состоит из 5 методов
- GetPosts - /api/posts [GET]
- GetPostsByID /api/post/{postId} [GET]
- DeletePost /api/post/{postId} [DELETE]
- UpdatePost /api/post/{postId} [UPDATE]
- CreatePost /api/post/{postId} [POST]

Всё в общем-то крутится вокруг одной структуры - Post, котороя состоит из:
-Заголовка
-Текста

Сответственно в формате Json она подётся в методы CreatePost, DeletePost и UpdatePost.
В GetPostsById подаётся id в int формате, в GetPosts не подаётся ничего =)

Базу данных решил не прикручивать, всё сделал in-memory, тесты тоже не гонял.
Ну думаю, что по моему гитхабу можно понять, что я знаю, что такое юнит-тесты, так что я не со зла =)

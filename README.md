# 題目
todo-list 個人 side project

# 動機
覺得自己太廢物...連個後台管理系統都刻不出來...決定好好修煉一番...

# 目標
1. [x] 學習如何製作API(CRUD)
2. [x] 學習簡單前端做渲染(SSR)
3. [x] 透過專案學習DB技術
	1. [x] 學習設計DB Table
	2. [x] 學習SQL語法，盡量不要去仰賴ORM套件(使用DB pool設定，但是sql自己下)
	3. [x] 透過pagenation，做分頁
4. [x] 學習Casbin，做使用者行為限制
5. [ ] 建置Account機制，做使用者登入頁面
6. [ ] 使用 JWT 機制，將回傳 token 帶入使用者角色身份
7. [ ] 使用 go-redis 對 list 資料做緩存
8. [ ] 使用Websocket，做前端頁面推廣通知
9. [ ] 學習使用Lock(考慮多個使用者同時操作某一比留言訊息)


# 實現功能
1. 可以增加備忘錄
2. 可以讀取備忘錄內容(分頁)
3. 可以更新備忘錄狀態(批次)
4. 可以刪除備忘錄(批次)


# 基於 mysql 資料庫
- markdb
1. note_table
| Id |     created_at   |    updated_at    |    deleted_at    | title | content |
|----|:----------------:|-----------------:|-----------------:|------:|--------:|
| 1  | 2021/05/24 10:00 | 2021/05/25 10:00 | 2021/05/24 10:05 | study | ansible |
| 2  | 2021/05/25 10:00 | 2021/05/26 10:00 |                  | study | golang  |
| 3  | 2021/05/25 10:00 | 2021/05/27 10:00 |                  | study | python  |

2. auth_casbin_rule

# 快速啟動
1. 架設環境 mysql
> docker-compose up -d

2. 初始化 db 資料
> bash ./scripts/setup.sh

3. 運行服務
> go run main.go -config ./config/app.dev.ini

4. 打開網頁，確認備忘錄清單伺服器運行中
> http://127.0.0.1:8000


# Inspired:
- https://github.com/mohtamohit/go-todo-list 
### casbin
- https://github.com/casbin/casbin
### gin-web
- https://github.com/piupuer/gin-web
- https://github.com/piupuer/gin-web-vue
# 題目
todo-list 個人 side project

# 動機
覺得自己太廢物...連個後台管理系統都刻不出來...決定好好修煉一番...

# 目標
1. 學習如何製作API(CRUD)
2. 學習簡單前端做渲染(SSR)
3. 透過專案學習DB技術
	1. 學習設計DB Table
	2. 學習SQL語法，盡量不要去仰賴ORM套件(使用DB pool設定，但是sql自己下)
	3. 透過pagenation，做分頁
4. 學習Casbin，考慮做使用者登入管理
5. 學習使用Lock(考慮多個使用者同時操作某一比留言訊息)
6. 學習使用Raft，做水平Cluster擴展


# 實現功能
1. 可以增加備忘錄
2. 可以讀取備忘錄內容(分頁)
3. 可以更新備忘錄狀態(批次)
4. 可以刪除備忘錄(批次)


# 基於 mysql 資料庫
- markdb
1. content table
| Id |    Content    |   Created Time   |   Updated Time   | Fin |
|----|:-------------:|-----------------:|-----------------:|----:|
| 1  | play leetcode | 2021/05/24 10:00 | 2021/05/24 10:05 |  Y  |
| 2  | exercise      | 2021/05/24 11:00 | 2021/05/24 11:50 |  N  |
| 3  | learn golang  | 2021/05/24 12:00 | 2021/05/24 14:00 |  N  |

2. 

# 快速啟動
1. 架設環境 mysql
> docker-compose up -d

2. 初始化 db 資料
> bash ./scripts/setup.sh

3. 運行服務
> go run main.go -config ./config/app.dev.ini

4. 打開網頁，確認備忘錄清單伺服器運行中
> http://127.0.0.1:8000


# working history
- 2021/05/24:
```md
Implement:
後端 CRUD

TODO:
1. 考慮用 raw sql 而不是 gorm method
2. 把 gorm 替換成 database/sql
3. 做 pagenation
4. 針對讀資料做緩存
```


# Inspired:
- https://github.com/mohtamohit/go-todo-list 

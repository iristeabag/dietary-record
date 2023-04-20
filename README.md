# Dietary-record
簡單紀錄個人飲食、熱量、身體數據變化小工具

## 運行
### 前置作業
須先將執行 table 建置完畢。
(需要 migration 可以從上方的 migration 目錄底下找到)

#### 問題補充
如果 migration 發生以下錯誤時，請先執行此 psql 指令
```
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
```

### 啟動服務
```
go run main.go
```

## 架構介紹
基於 go-kit 實現同時支援 http 與 grpc 的微服務。
按照業務進行模組的拆分，每個業務是一個目錄。
於 eat 飲食紀錄中使用到呼叫 grpc 驗證 id 是否存在之功能。

### 整體目錄結構
```
├─food 食物
│  ├─endpoint
│      ├─grpc
│      └─http
│  ├─migration # table sql
│  ├─service # 主要邏輯
│  └─transport
│      ├─grpc
│      └─http
├─eat 飲食紀錄
│  ├─endpoint
│      ├─grpc
│      └─http
│  ├─migration # table sql
│  ├─service # 主要邏輯
│  └─transport
│      ├─grpc
│      └─http
├─body 身體數據
│  ├─endpoint
│      ├─grpc
│      └─http
│  ├─migration # table sql
│  ├─service # 主要邏輯
│  └─transport
│      ├─grpc
│      └─http
├─proto # proto 文件
└─server # http、grpc 服務
```
### DATABASE ERD 圖
![截圖 2023-04-19 下午12 43 09](https://user-images.githubusercontent.com/1109393/233096379-6df8e0c8-b211-4bcc-a717-1c8ebef681a3.png)

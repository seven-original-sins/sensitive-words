# sensitive-words
基于 aho-corasick 算法的敏感词检测 web 服务器

## 功能
- request auth token
- 查找内容中所有敏感词（HTTP）
- 监听字典数据源变化（新增/删除word），自动更新 trie
- 手动触发 reload（HTTP）

## 依赖
- go module
- [anknown/ahocorasick](github.com/anknown/ahocorasick)
- [fasthttp](github.com/valyala/fasthttp)

## 安装
```
git clone https://github.com/seven-original-sins/sensitive-words.git
```

## mysql 数据源表结构
```sql
CREATE TABLE `t_sensitive_word` (
  `iAutoId` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'auto increment id',
  `_iDeleteTime` int(11) unsigned NOT NULL DEFAULT 0 COMMENT 'soft delete time'
  `sWord` varchar(63) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '敏感词',
  PRIMARY KEY (`iAutoId`),
  KEY `idx_sWord` (`sWord`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin
```

## 环境变量
```
export SENSITIVE_TOKEN=yourtoken
export SENSITIVE_LISTEN_ADDR=:9999
export DB_SECURITY_READ_HOST=localhost
export DB_SECURITY_READ_PORT=3306
export DB_SECURITY_READ_DATABASE=yourdatabase
export DB_SECURITY_READ_USERNAME=username
export DB_SECURITY_READ_PASSWORD=password
```

## API
search 
```
curl -X POST \
  http://localhost:9999/words/search \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -d 'content=TMD'
```

reload
```
curl http://localhost:9999/words/reload
```
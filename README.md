# go-chatgpt
a simple chatgpt http server

# 如何配置

**configs/config.yml**

```yaml
env:
  name: my

gpt:
  authorization: "Bearer sk-*"
  port: "8080"
```

# 如何使用

```
curl http://localhost:8080/s?assistant=%s&content=%s&user=%s

assistant：role=assistant，非必填
user：userid，非必填
content：内容，必填

```


# future-go


## 2023新增
- UUID[go.uuid](https://github.com/satori/go.uuid)
- 雪花算法[snowflake](https://github.com/bwmarrin/snowflake)

> UUID **简单无序**十分适合生成 requestID; Snowflake 里面包含时间序列等，可以用于排序，效率都还可以
https://cloud.tencent.com/developer/article/1766264

- UUID 一般会使用它生成 Request ID 来标记单次请求
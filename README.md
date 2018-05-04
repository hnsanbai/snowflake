# snowflake
一个使用Go语言实现的snowflake包



# 获取包
go get github.com/hnsanbai/snowflake


# 使用
import github.com/hnsanbai/snowflake

id := snowflake.GetSnowFlakeID(10, 10)
// 参数1：数据中心标识，参数2：机器标识

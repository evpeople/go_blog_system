# go_blog_system
a blog system write with golang and vue

## 调研

gin  负责路由，实际上使用了go http 
gorm 负责数据模型
vue 负责前端

REST API

### REST API
1. URL 中，不应含有动词
1.1. 只应该使用GET POST PUT DELETE　PATCH 等http方式的动作，完成相应的操作，比如POST应该是往某个服务进行POST，然后在上下文中加入数据
1.2 　上文一个例子，　GET /articles  CRUD中的读取数据库
2. 避免多级URL，除了第一级以外，其他级别都用查询字符串表示

### 问题
后端的API怎么在前端调用（预计７月６日解决

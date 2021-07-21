# go_blog_system
a blog system write with golang and vue

## 使用
clone库后，运行server.go，然后在浏览器访问127.0.0.1:3456 或 localhost:3456 ，若出现404 错误，则运行成功
访问/api 可得到api列表

## 目前进度

### API 
[] 上传文章API
[] 展示单个文章API
[] 展示文章列表API
[] API上传的确权(以防被POST炸满数据库)
[] 按照分类展示的API
[] 按照标签展示的API

### 前端
[] 前端展示页面
[] 前端编辑页面

## 我可以做什么？
1. 为api写测试用例，推荐 使用 go test而不是goland ide的测试用例
2. 引入新的API
3. 写前端页面
4. 写自动化测试脚本 ，使用git action功能
5. 提出新的issue
6. 写词语图的生成函数
7. 尝试自己的路由mux的开发
8. 提出新的功能需求

## 我怎么开始做
1. fork此仓库
2. clone 自己fork的仓库到本地
3. 做出修改
4. pull到自己的仓库
5. 发起pull request
6. 请尽量遵循Angular提交信息规范,此规范的基本信息在commit.template里，可以配置成commit 的模板，或者采用其他插件
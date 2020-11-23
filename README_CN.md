# jenkinsapi
jenkins api for go

### 环境安装

docker run -p 8080:8080 -p 50000:50000 -v jenkins_data:/var/jenkins_home jenkinsci/blueocean

### 提交规范

用于说明git commit的类别，只允许使用下面的标识。

feat：新功能（feature）。

fix/to：修复bug，可以是QA发现的BUG，也可以是研发自己发现的BUG。

fix：产生diff并自动修复此问题。适合于一次提交直接修复问题

to：只产生diff不自动修复此问题。适合于多次提交。最终修复问题提交时使用fix

docs：文档（documentation）。

style：格式（不影响代码运行的变动）。

refactor：重构（即不是新增功能，也不是修改bug的代码变动）。

perf：优化相关，比如提升性能、体验。

test：增加测试。

chore：构建过程或辅助工具的变动。

revert：回滚到上一个版本。

merge：代码合并。

sync：同步主线或分支的Bug。

scope(可选):scope用于说明 commit 影响的范围，比如数据层、控制层、视图层等等，视项目不同而不同。

### 计划

12.6前完成除job相关的全部api对接（管理，展示为主）

12.13前完成主要的插件调研

12.31前完成job相关api接口和至少两个主要插件api接口对接
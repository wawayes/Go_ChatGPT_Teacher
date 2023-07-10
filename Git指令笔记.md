---
share: true  
---


# .gitignore失效问题

在.gitignore之前add过的文件，会维持track状态，如果中途在.gitignore文件中添加了该文件，那么就会失效（还是会被commit到远程）
此时可以使用命令
```shell
git rm [文件名,例如：.env这种包含敏感信息的文件] --cached
```
此时commit中的.env就已经变灰色了，push之后就会发现仓库里没有了.env文件


1.  设置 golang.org 镜像：glide mirror set https://golang.org/x/sys https://github.com/golang/sys --vcs git

2.  在 glide.yaml 中添加 package: `- package: golang.org/x/sys`

通过以上两步就可以解决 glide up 时候出现： `[ERROR] Error looking for golang.org/x/sys/unix: Cannot detect VCS` 的错误。

当然你可以不用 glide 用官方的 dep 就不存在这个问题？？

### Notice

**是用 dep 安装依赖包始终不成功！**

使用`govendor` 可以轻松成功！

使用`glide` 也可以成功！

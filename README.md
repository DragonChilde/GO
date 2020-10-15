# linux下安装GO

1. 解压文件到/opt/下

   ```bash
   tar -zxvf go1.15.3.linux-amd64.tar.gz -C /opt/
   ```

2. linux下配置Golang环境变量

   ```bash
   vim /etc/profilevim /etc/profile
   ```

3. 参数如下

   ```bash
   export GOROOT=/opt/go
   export PATH=$PATH:$GOROOT/bin
   export GOPATH=$HOME/goproject
   ```

4. 刷新环境变量

   ```bash
   source /etc/profile
   ```

   


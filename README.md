# astools

astools本意是Android shell tools，主要是Android adb shell下无法使用ssh和scp命令，所以写了这个cli tools，可以实现在Android adb shell环境下:

①通过ssh远程到服务器执行Linux command

②通过SCP实现本地与服务器之间互相上传和下载文件(暂不支持目录上传下载)。

得益于Golang的交叉编译，所以此cli tools不局限于Android端使用，同时也支持MacOS，Windows，Linux。

## Usage
### 查看astools帮助信息
```shell
hashqueue@hashqueue-pc:~$ adb shell
odin:/ $ cd /data/local/tmp
odin:/data/local/tmp $ ls -larth
total 385M
drwxr-x--x 7 root  root  3.3K 1970-03-30 04:49 ..
-rw-rw-rw- 1 shell shell 190M 2022-12-08 03:24 python3-android-3.9.0-arm64-21.tar
-rw-rw-rw- 1 shell shell   91 2022-12-08 03:32 test.py
-rwxrwxrwx 1 shell shell 190M 2023-01-02 04:35 python3-android-arm64.tar
-rwxrwxrwx 1 shell shell 4.7M 2023-01-02 05:25 astools_android_arm64
drwxrwx--x 2 shell shell 3.3K 2023-01-02 06:11 .
odin:/data/local/tmp $ ./astools_android_arm64 -h                                                                                                                           
Usage of ./astools_android_arm64:
  -cmd string
    	command to run.
  -ip string
    	machine ip address.
  -local-path string
    	local file path when use SCP.
  -pass string
    	ssh password or scp password.
  -port uint
    	ssh port number or scp port number. (default 22)
  -remote-path string
    	remote file path when use SCP.
  -scp-type string
    	upload: upload local file to remote server; download: download remote file to local. (default "upload")
  -timeout uint
    	execute command's timeout(s) when use ssh or file transfer's timeout(s) when use SCP. (0 means no timeout(s), not recommended!). (default 60)
  -type string
    	ssh: execute a command with ssh; scp: file transfer with SCP.
  -user string
    	ssh username or scp username.
```

### 查看astools提示信息
```shell
odin:/data/local/tmp $ ./astools_android_arm64                                                                                                                              
2023/01/01 22:15:50 [error] -type param must be `ssh` or `scp`, not null.
2023/01/01 22:15:50 [error] Parse parmas error, Please check your input params!
Welcome to astools, you can type ./astools_android_arm64 -h to show help message.
Usages:
1. To execute a command on remote server with ssh:
	./astools_android_arm64 -type ssh -ip 192.168.124.16 -user ubuntu -pass 123456 -port 22 -timeout 5 -cmd "pwd"
2. To upload local file to remote server:
	./astools_android_arm64 -type scp -ip 192.168.124.16 -user ubuntu -pass 123456 -port 22 -timeout 5 -scp-type upload -local-path ./demo.txt -remote-path /home/ubuntu/demo1.txt
3. To download remote file to local:
	./astools_android_arm64 -type scp -ip 192.168.124.16 -user ubuntu -pass 123456 -port 22 -timeout 5 -scp-type download -local-path ./demo2.txt -remote-path /home/ubuntu/demo1.txt
```

### 通过ssh向服务器发送一条命令执行完毕并退出
```shell
4|odin:/data/local/tmp $ ./astools_android_arm64 -type ssh -ip 192.168.124.16 -user ubuntu -pass 123456 -port 22 -timeout 5 -cmd "df -h"
2023/01/01 22:19:46 Start executing command [df -h]
2023/01/01 22:19:48 -----------------------stdout && stderr---------------------->
Filesystem      Size  Used Avail Use% Mounted on
tmpfs           781M  3.4M  778M   1% /run
/dev/sda2       440G  7.8G  415G   2% /
tmpfs           3.9G     0  3.9G   0% /dev/shm
tmpfs           5.0M     0  5.0M   0% /run/lock
/dev/sda1       253M  148M  105M  59% /boot/firmware
tmpfs           781M  4.0K  781M   1% /run/user/1000

-------------------------------------------end----------------------------------->
2023/01/01 22:19:48 Execute command [df -h] complete.
2023/01/01 22:19:48 Total use time: 2.256037 s
```

### 通过SCP上传本地文件到服务器
```shell
odin:/data/local/tmp $ ./astools_android_arm64 -type scp -ip 192.168.124.16 -user ubuntu -pass 123456 -port 22 -timeout 60 -scp-type upload -local-path ./python3-android-3.9.0-arm64-21.tar -remote-path /home/ubuntu/python3-android-arm64.tar
2023/01/01 22:24:32 Start copy local file[./python3-android-3.9.0-arm64-21.tar] to remote server[192.168.124.16:22 - /home/ubuntu/python3-android-arm64.tar], please wait...
2023/01/01 22:25:02 Copy local file[./python3-android-3.9.0-arm64-21.tar] to remote server[192.168.124.16:22 - /home/ubuntu/python3-android-arm64.tar] done.
2023/01/01 22:25:02 Total use time: 30.171892 s
```

### 通过SCP从服务器下载文件到本地
```shell
odin:/data/local/tmp $ ./astools_android_arm64 -type scp -ip 192.168.124.16 -user ubuntu -pass 123456 -port 22 -timeout 60 -scp-type download -local-path ./python3.tar -remote-path /home/ubuntu/python3-android-arm64.tar
2023/01/01 22:29:14 Start copy remote file[192.168.124.16:22 - /home/ubuntu/python3-android-arm64.tar] to local[./python3.tar], please wait...
2023/01/01 22:29:40 Copy remote file[192.168.124.16:22 - /home/ubuntu/python3-android-arm64.tar] to local[./python3.tar] done.
2023/01/01 22:29:40 Total use time: 26.007062 s
```

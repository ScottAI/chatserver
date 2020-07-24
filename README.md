Go语言实现简单聊天室

## 功能
#### 	用户可以连接到服务器
#### 	用户可以设定自己的用户名
#### 	用户可以向服务器发送消息，同时服务器会向其他用户广播该消息

### 协议包-protocol
我们可以把客户端与服务器的交互分为三种类型的命令：
##### 	发送命令(SEND)：客户端发送一个聊天信息
##### 	名字命令(NAME)：客户端发送自己的名字
##### 	信息命令(MESS)：服务器向客户端广播信息

### 服务器-server
服务端实现并不复杂，
因为我们把很多通信相关的功能在protocol包已经实现，
可以看出包的功能划分在项目中非常重要，划分清晰便于理清思路也便于代码实现。

##### 注意接口的使用。

现在我们来看一下server包文件的结构：
##### --chatserver
##### ----server
##### ------server.go
##### ------tcp_server.go
##### ------cmd
##### --------main.go


### 客户端-client
客户端主要完成的功能是：
##### 	连接服务器
##### 	使用用户名登录
##### 	发送消息
##### 	接收其他人发送的消息

现在我们来看一下客户端完成的代码结构：
##### --chatserver
##### ----client
##### ------client.go
##### ------tcp_client.go
##### ----gui
##### ------gui.go
##### ------cmd
##### --------main.go

## 运行效果

#### 登录界面
![](https://github.com/ScottAI/chatserver/blob/master/pictures/logon.png)

#### 信息发送界面
![](https://github.com/ScottAI/chatserver/blob/master/pictures/hello1.png)

#### 发送完成界面
![](https://github.com/ScottAI/chatserver/blob/master/pictures/send.png)

# bdntoy

- **名称含义**: baidu netdisk toy, 百度网盘玩具
- **背景**：截止2024.3.17，百度不支持用户检索文件库中文件的路径，所以不方便用户保存想要的部分分享目录，只能保存到自己网盘再查看，但内容太多，全部保存网盘空间不够。
- **用途**: 用于生成分享文件库的内容树
- **原理**: 百度网盘未对个人开发者提供查看分享文件库的接口，所以直接通过网页客户端找到相关的接口进行调用。

## TODO
- [x] 查看分享会话中的所有文件,并打印到控制台
- [ ] 查看部分目录并打印到控制台

## 使用

### 登录:
登录后会保存配置在目标路径，通过指定 BDNTOY_GO_CONFIG_DIR，可以设置配置保存的目录。 

配置名为 config.json,保存路径，默认 /HOME/.config/bdntoy 或者 /tmp



- 使用 cookie 登录
```shell
bdntoy login --cookies '你的cookie'
```

> 如何获取 cookie? 请参考文献 [4]

或者：
```shell
bdntoy login --bduss 'asdasd' --stoken '12312as'
```

> 如何获取登录凭证，请参考文献 [5]

### 获取会话列表

```shell
bdntoy hs
```

输出示例：
```text
+---+------------------------+------------+---------------------+
| # |          群名           | 邀请进群者  |      更新时间       |
+---+------------------------+------------+---------------------+
| 0 | 学习资料群               | 系统       | 2024-03-16 18:13:47 |
| 1 | 命运****之门、小惊***玩家  | 系统       | 2024-03-16 11:06:20 |
+---+------------------------+------------+---------------------+
```

选择会话
```shell
bdntoy hs 0
```

### 获取通讯录中的群组
```shell
bdntoy gs
```

选择群组:
```shell
bdntoy gs 0
```

```text
+---+------------------------+-------------+---------------------+
| # |          群名          |   创建者    |      创建时间       |
+---+------------------------+-------------+---------------------+
| 0 | 学习资料群               | tuh****o_go | 2017-09-16 08:18:50 |
| 1 | 命运****之门、小惊***玩家  | 小惊***玩家   | 2024-03-04 22:27:47 |
+---+------------------------+-------------+---------------------+
```

### 获取会话中的分享文件库

```shell
bdntoy fl
```

```text
+---+------------+------------------------+---------------------+
| # | 文件库名称   |          内容           |      创建时间       |
+---+------------+------------------------+---------------------+
| 0 | x**xx      | xxxxxxxx               | 2024-03-05 20:15:54 |
| 1 | x**xx      | xxxxxxxx               | 2024-03-05 20:15:52 |
+---+------------+------------------------+---------------------+
```

### 打印文件库中的所有文件
```shell
bdntoy tree 文件库序号
```
结果输出示例:
```text
+-----+--------------------------------------------------------------------+---------+---------------------+
|  #  |                                文件                                 | 大小/MB  |      创建时间       |
+-----+--------------------------------------------------------------------+---------+---------------------+
|   0 | 24.参考资源和后续课程预览.mp4                                          |     117 | 2018-12-25 15:48:26 |
|   1 | 23.下一代微服务安全架构.mp4                                            |     106 | 2018-12-25 15:48:18 |
|   2 | 22.OpenId Connect简介.mp4                                           |      35 | 2018-12-25 15:47:58 |
+-----+--------------------------------------------------------------------+---------+---------------------+
```

保存输出：
```shell
bdntoy tree 文件库序号 >> output.txt
```

# 参考文献
- [1] [Github仓库: 极客时间下载器](https://github.com/mmzou/geektime-dl)
- [2] [Github仓库: youtube下载器](https://github.com/kkdai/youtube) 
- [3] [百度网盘开放平台文档](https://pan.baidu.com/union/document/basic)
- [4] [如何抓包获取百度网盘网页版完整 Cookie](https://blog.imwcr.cn/2022/11/24/%E5%A6%82%E4%BD%95%E6%8A%93%E5%8C%85%E8%8E%B7%E5%8F%96%E7%99%BE%E5%BA%A6%E7%BD%91%E7%9B%98%E7%BD%91%E9%A1%B5%E7%89%88%E5%AE%8C%E6%95%B4-cookie/)
- [5] [使用cookie登录百度网盘账号](https://blog.csdn.net/weixin_39734304/article/details/102418180)

# 联系我

个人邮箱: haotu007@gmail.com

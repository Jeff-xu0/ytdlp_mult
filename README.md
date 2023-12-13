
引用了大佬的项目 https://github.com/yt-dlp/yt-dlp?tab=readme-ov-file#general-options

Quoting the mogul’s project https://github.com/yt-dlp/yt-dlp?tab=readme-ov-file#general-options

的成品

's finished product

# 中文版
## 开始
### 安装yt_dlp软件
本文基于windows，打开FilesNeeded，将下面的所有程序放到你自己的文件夹下[最好是新建一个]。

然后为这个文件夹添加一个环境变量，值为这个文件夹的路径。

然后将该编译好的程序放到你的文件夹下（下载的视频文件会统一的放到当前目录）

执行以下命令：
```shell
ytdlp_mult.exe -url https://www.bilibili.com/video/BV1PY411P7D1 
```
那么不出意外的话他就会在你的当前目录下下载视频文件了

然后，防止小白不会用，加一句说明：

我经常下B站的视频，B站视频的url的话满足以下需求：

-url 后的视频地址止步于 ‘?’ 之前 

什么意思呢？

比如我有一个地址是 https://www.bilibili.com/video/BV1PY411P7D1/?spm_id_from=333.337.search-card.all.click&vd_source=d4e22590a499d985a801b3d0ddab020f

很长，但是我只取 https://www.bilibili.com/video/BV1PY411P7D1

三千弱水，我只取一瓢淫

### 改进
代码里面有很多参数我都没写，大家有兴趣的话可以搜一下TODO写一下

yt_dlp的项目在：https://github.com/yt-dlp/yt-dlp?tab=readme-ov-file#general-options

进去之后搜索 ‘USAGE AND OPTIONS’

里面有很多参数，我就不一一写了，大家自己看吧

### 闲聊
然后呢，提一嘴，目前我只测过B站，但是yt_dlp是支持很多的视频网址，

包括那什么管，看名字就知道他是为那个而生的

不过代码里限制了，不能下，哪里限制了呢？

就是上面说的这个取网址那里限制了。有能力的改一改就好，没能力的就等我下次有空更新吧

然后那些连管都看不了的就不要考虑了,安心冲B站吧

### 最后
这是我随便写的一个demo[不要说我菜],嗯，随便写的。写的不是很好[这点我承认],希望各位大佬能加入进来帮小弟改进

欢迎各位的拉取请求，在此感谢，鞠躬

好吧，我估计是没人会看到这个demo的，等他雪藏在南极洲的海底吧

鞠躬

退场

...

# English
# start
## Install yt_dlp software
This article is based on windows. Open FilesNeeded and put all the following programs into your own folder [preferably create a new one].

Then add an environment variable to this folder with the value being the path to this folder.

Then put the compiled program in your folder (the downloaded video files will be uniformly placed in the current directory)

Execute the following command:
```shell
ytdlp_mult.exe -url https://www.bilibili.com/video/BV1PY411P7D1
```
Then if nothing else happens, it will download the video file in your current directory.

Then, just in case novices don’t know how to use it, add an explanation:

I often download videos from station B. The URL of the video from station B meets the following requirements:

The video address after -url ends before ‘?’

What does that mean?

For example, I have an address: https://www.bilibili.com/video/BV1PY411P7D1/?spm_id_from=333.337.search-card.all.click&vd_source=d4e22590a499d985a801b3d0ddab020f

It’s very long, but I only take https://www.bilibili.com/video/BV1PY411P7D1

Three thousand weak water, I only take one ladle

## Improve
There are many parameters in the code that I have not written. If you are interested, you can search the TODO and write them down.

The yt_dlp project is at: https://github.com/yt-dlp/yt-dlp?tab=readme-ov-file#general-options

Once inside, search for ‘USAGE AND OPTIONS’

There are many parameters in it, so I won’t write them down one by one. You can read it yourself.

## Chat
Then, just to mention, I have only tested site B so far, but yt_dlp supports many video URLs.

Including that tube, you can tell from the name that he was born for that

But there is a restriction in the code and it cannot be downloaded. Where is the restriction?

It’s the restriction on fetching the URL mentioned above. If you are able, just change it. If you are not, just wait until I have time to update it next time.

Then don’t consider those who can’t even watch it, just go to station B with peace of mind.

## at last
This is a demo I wrote casually [don’t call me bad], well, I wrote it casually. The writing is not very good [I admit this], I hope you guys can join in and help me improve.

Your pull requests are welcome, thank you and take a bow.

Well, I guess no one will see this demo. Let’s wait until it is hidden under the sea in Antarctica.

bow

Exit

...
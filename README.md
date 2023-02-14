# ExchangePE
Asset scanning by dictionary stitching Domain to identify Exchange Servers versions

## 1、前言

我们在内网中想要拿到域控，肯定会想到Exchange Service服务器，Exchange服务器的权限一般都是域管理员权限，所以拿下服务器的权限也就离域控权限不远了。这个工具主要是使用Go重构了ExchangeFinder工具，并做了一些更新，减少原工具匹配不全面的问题，以及实现了Go语言的高并发。

## 2、使用教程

可以自己搜集添加域名前缀，并加入到doamin.txt中,来增加爆破几率

直接上图

![test1.png](/Users/Monster/Desktop/RedTeam/picture/djvtZC1gsQVlEwR-20230214225813309.png)

![test111.png](/Users/Monster/Desktop/RedTeam/picture/AzncejgZRuMsOh7.png)

![testtt.png](/Users/Monster/Desktop/RedTeam/picture/Jirxq7LmQvSFYBo.png)

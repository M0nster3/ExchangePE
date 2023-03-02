# ExchangePE
Asset scanning by dictionary stitching Domain to identify Exchange Servers versions

## 1、前言

我们在内网中想要拿到域控，肯定会想到Exchange Service服务器，Exchange服务器的权限一般都是域管理员权限，所以拿下服务器的权限也就离域控权限不远了。这个工具主要是使用Go重构了ExchangeFinder工具，并做了一些更新，减少原工具匹配不全面的问题，以及实现了Go语言的高并发。

## 2、使用教程

可以自己搜集添加域名前缀，并加入到doamin.txt中,来增加爆破几率

修复了一些bug以及更新了规则，使用之前先输入./ExchangePE -update 1 更新一下版本，再进行使用，不再只依赖网络匹配

直接上图



[![ppFc1ds.png](https://s1.ax1x.com/2023/03/02/ppFc1ds.png)](https://imgse.com/i/ppFc1ds)

[![ppFcdL4.png](https://s1.ax1x.com/2023/03/02/ppFcdL4.png)](https://imgse.com/i/ppFcdL4)

[![ppFc0eJ.png](https://s1.ax1x.com/2023/03/02/ppFc0eJ.png)](https://imgse.com/i/ppFc0eJ)

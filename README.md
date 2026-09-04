# S-UI
**An Advanced Web Panel • Built on SagerNet/Sing-Box**

![](https://img.shields.io/github/v/release/alireza0/s-ui.svg)
![S-UI Docker pull](https://img.shields.io/docker/pulls/alireza7/s-ui.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/alireza0/s-ui)](https://goreportcard.com/report/github.com/alireza0/s-ui)
[![Downloads](https://img.shields.io/github/downloads/alireza0/s-ui/total.svg)](https://img.shields.io/github/downloads/alireza0/s-ui/total.svg)
[![License](https://img.shields.io/badge/license-GPL%20V3-blue.svg?longCache=true)](https://www.gnu.org/licenses/gpl-3.0.en.html)

> **Disclaimer:** This project is only for personal learning and communication, please do not use it for illegal purposes, please do not use it in a production environment

**If you think this project is helpful to you, you may wish to give a**:star2:

**Want to contribute?** See [CONTRIBUTING.md](CONTRIBUTING.md) for development setup, coding conventions, testing, and the pull request process.

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/alireza7)

<a href="https://nowpayments.io/donation/alireza7" target="_blank" rel="noreferrer noopener">
   <img src="https://nowpayments.io/images/embeds/donation-button-white.svg" alt="Crypto donation button by NOWPayments">
</a>

## Rulio723 定制版

本仓库以 [`alireza0/s-ui`](https://github.com/alireza0/s-ui) 官方最新版为基线维护，定制代码位于 `custom/s-ui` 分支。前端仍保持独立子模块结构，定制前端位于 [`Rulio723/s-ui-frontend`](https://github.com/Rulio723/s-ui-frontend) 的 `custom/s-ui-frontend` 分支。

### 新增功能

- 客户端独立保存 Email、Sub ID 和自动续期开关，同时保留 UUID、到期时间、流量上限及入站绑定能力。
- 支持按周期自动重置流量，并可同步顺延到期时间；自动续期时可重新启用已到期或超量客户端。
- 订阅地址优先使用不可预测的 Sub ID；仅对尚未设置 Sub ID 的旧客户端保留名称路径兼容。
- 增加 VLESS Reality Vision、Hysteria2、TUIC、Trojan 和 AnyTLS 快捷模板，支持端口/Tag 避让及失败补偿回滚。
- 增加导出中心，可导出分享链接、UTF-8 Base64、订阅地址及二维码。
- 支持按用户校准当前周期已用上传和下载流量，便于修复迁移或统计偏差。
- 增加中转管理，支持 SOCKS5 和分享链接转换、指定入站路由绑定、连通性测试及删除。
- 入站卡片增加二维码入口；多客户端入站可先选择目标客户端。
- 在官方明暗主题之外增加 Aurora、Mesh、DeepSea、Sunrise、Cyber、Mint 六套皮肤。
- 增加自动续期、Sub ID 查询、模板、中转、导出和皮肤相关测试。

### 与官方上游的区别

| 范围 | 官方上游 | Rulio723 定制分支 |
| --- | --- | --- |
| 客户端身份 | `name` 同时承担核心身份和订阅路径 | 独立 Email、Sub ID；`name` 继续用于 sing-box 身份和流量关联 |
| 周期管理 | 支持周期流量重置 | 可在重置流量时同步续期并重新启用客户端 |
| 流量校准 | 仅显示和自动统计客户端流量 | 可单独校准每个用户当前周期的上传与下载用量 |
| 订阅查询 | 根据客户端名称查询 | Sub ID 优先，旧数据仅在 Sub ID 为空时回退名称 |
| 入站创建 | 使用通用配置表单 | 额外提供五类常用协议快捷模板和批量创建 |
| 导出与二维码 | 提供基础的单客户端入口 | 提供全量导出中心和入站卡片二维码入口 |
| 中转 | 通过原始出站和路由配置管理 | 提供中转向导、入站绑定、测速和清理界面 |
| 外观 | 官方明暗主题 | 官方主题加六套独立皮肤 |
| 维护方式 | 官方主线开发 | 定制功能按独立提交维护，并持续跟随官方基线 |

本分支没有直接合并旧公开改版的完整历史，也没有默认加入多服务器中央代理、Web 一键更新、自动开放防火墙或批量破坏性删除等高风险功能。官方更新时应优先采用新的上游结构，再逐项保留和回归验证上述定制行为。

构建用于 Debian 等 glibc 系统的部署产物时，应先构建前端并写入后端 `web/html`，再使用 glibc 环境编译 Go 二进制；不要直接部署 Alpine/musl 构建产物。

## Quick Overview
| Features                               |      Enable?       |
| -------------------------------------- | :----------------: |
| Multi-Protocol                         | :heavy_check_mark: |
| Multi-Language                         | :heavy_check_mark: |
| Multi-Client/Inbound                   | :heavy_check_mark: |
| Advanced Traffic Routing Interface     | :heavy_check_mark: |
| Client & Traffic & System Status       | :heavy_check_mark: |
| Subscription Link (link/json/clash + info)| :heavy_check_mark: |
| Dark/Light Theme                       | :heavy_check_mark: |
| API Interface                          | :heavy_check_mark: |

## Supported Platforms
| Platform | Architecture | Status |
|----------|--------------|---------|
| Linux    | amd64, arm64, armv7, armv6, armv5, 386, s390x | ✅ Supported |
| Windows  | amd64, 386, arm64 | ✅ Supported |
| macOS    | amd64, arm64 | 🚧 Experimental |

## Screenshots

!["Main"](https://github.com/alireza0/s-ui-frontend/raw/main/media/main.png)

[Other UI Screenshots](https://github.com/alireza0/s-ui-frontend/blob/main/screenshots.md)

## API Documentation

[API-Documentation Wiki](https://github.com/alireza0/s-ui/wiki/API-Documentation)

## Default Installation Information
- Panel Port: 2095
- Panel Path: /app/
- Subscription Port: 2096
- Subscription Path: /sub/
- User/Password: admin

## Install & Upgrade to Latest Version

### Linux/macOS
```sh
bash <(curl -Ls https://raw.githubusercontent.com/alireza0/s-ui/master/install.sh)
```

#### Installer language

The installer is available in the same six languages as the panel: `en` (default), `fa`, `ru`, `vi`, `zhcn`, `zhtw`. Choose one with the `SUI_LANG` environment variable (when unset, your system `$LANG` is used as a hint):

```sh
SUI_LANG=fa bash <(curl -Ls https://raw.githubusercontent.com/alireza0/s-ui/master/install.sh)
```

### Alpine Linux
Alpine uses `apk` and OpenRC instead of `apt`/systemd. The install script detects Alpine automatically and sets up an OpenRC service. Since Alpine has no `bash` by default, install it first:

```sh
apk add bash
bash <(curl -Ls https://raw.githubusercontent.com/alireza0/s-ui/master/install.sh)
```

Manage the service with OpenRC: `rc-service s-ui start|stop|restart` and `rc-update add s-ui default`.

### Windows
1. Download the latest Windows release from [GitHub Releases](https://github.com/alireza0/s-ui/releases/latest)
2. Extract the ZIP file
3. Run `install-windows.bat` as Administrator
4. Follow the installation wizard

## Install legacy Version

**Step 1:** To install your desired legacy version, add the version to the end of the installation command. e.g., ver `v1.5.0`:

```sh
VERSION=v1.5.0 && bash <(curl -Ls https://raw.githubusercontent.com/alireza0/s-ui/$VERSION/install.sh) $VERSION
```

## Manual installation

### Linux/macOS
1. Get the latest version of S-UI based on your OS/Architecture from GitHub: [https://github.com/alireza0/s-ui/releases/latest](https://github.com/alireza0/s-ui/releases/latest)
2. **OPTIONAL** Get the latest version of `s-ui.sh` [https://raw.githubusercontent.com/alireza0/s-ui/master/s-ui.sh](https://raw.githubusercontent.com/alireza0/s-ui/master/s-ui.sh)
3. **OPTIONAL** Copy `s-ui.sh` to /usr/bin/ and run `chmod +x /usr/bin/s-ui`.
4. Extract s-ui tar.gz file to a directory of your choice and navigate to the directory where you extracted the tar.gz file.
5. Copy *.service files to /etc/systemd/system/ and run `systemctl daemon-reload`.
6. Enable autostart and start S-UI service using `systemctl enable s-ui --now`
7. Start sing-box service using `systemctl enable sing-box --now`

### Windows
1. Get the latest Windows version from GitHub: [https://github.com/alireza0/s-ui/releases/latest](https://github.com/alireza0/s-ui/releases/latest)
2. Download the appropriate Windows package (e.g., `s-ui-windows-amd64.zip`)
3. Extract the ZIP file to a directory of your choice
4. Run `install-windows.bat` as Administrator
5. Follow the installation wizard
6. Access the panel at http://localhost:2095/app

## Uninstall S-UI

### systemd
```sh
sudo -i

systemctl disable s-ui  --now

rm -f /etc/systemd/system/sing-box.service
systemctl daemon-reload

rm -fr /usr/local/s-ui
rm /usr/bin/s-ui
```

### Alpine (OpenRC)
```sh
rc-service s-ui stop
rc-update del s-ui default
rm -f /etc/init.d/s-ui

rm -fr /usr/local/s-ui
rm /usr/bin/s-ui
```

## Install using Docker

<details>
   <summary>Click for details</summary>

### Usage

**Step 1:** Install Docker

```shell
curl -fsSL https://get.docker.com | sh
```

**Step 2:** Install S-UI

> Docker compose method

```shell
mkdir s-ui && cd s-ui
wget -q https://raw.githubusercontent.com/alireza0/s-ui/master/docker-compose.yml
docker compose up -d
```

> Use docker

```shell
mkdir s-ui && cd s-ui
docker run -itd \
    -p 2095:2095 -p 2096:2096 -p 443:443 -p 80:80 \
    -v $PWD/db/:/app/db/ \
    -v $PWD/cert/:/root/cert/ \
    --name s-ui --restart=unless-stopped \
    alireza7/s-ui:latest
```

> Build your own image

```shell
git clone https://github.com/alireza0/s-ui
git submodule update --init --recursive
docker build -t s-ui .
```

</details>

## Manual run ( contribution )

<details>
   <summary>Click for details</summary>

### Build and run whole project
```shell
./runSUI.sh
```

### Clone the repository
```shell
# clone repository
git clone https://github.com/alireza0/s-ui
# clone submodules
git submodule update --init --recursive
```


### - Frontend

Visit [s-ui-frontend](https://github.com/alireza0/s-ui-frontend) for frontend code

### - Backend
> Please build frontend once before!

To build backend:
```shell
# remove old frontend compiled files
rm -fr web/html/*
# apply new frontend compiled files
cp -R frontend/dist/ web/html/
# build
go build -o sui main.go
```

To run backend (from root folder of repository):
```shell
./sui
```

</details>

## Languages

- English
- Farsi
- Vietnamese
- Chinese (Simplified)
- Chinese (Traditional)
- Russian

## Features

- Supported protocols:
  - General:  Mixed, SOCKS, HTTP, HTTPS, Direct, Redirect, TProxy
  - V2Ray based: VLESS, VMess, Trojan, Shadowsocks
  - Other protocols: ShadowTLS, Hysteria, Hysteria2, Naive, TUIC
- Supports XTLS protocols
- An advanced interface for routing traffic, incorporating PROXY Protocol, External, and Transparent Proxy, SSL Certificate, and Port
- An advanced interface for inbound and outbound configuration
- Clients’ traffic cap and expiration date
- Displays online clients, inbounds and outbounds with traffic statistics, and system status monitoring
- Subscription service with ability to add external links and subscription
- HTTPS for secure access to the web panel and subscription service (self-provided domain + SSL certificate)
- Dark/Light theme

## Environment Variables

<details>
  <summary>Click for details</summary>

### Usage

| Variable       |                      Type                      | Default       |
| -------------- | :--------------------------------------------: | :------------ |
| SUI_LOG_LEVEL  | `"debug"` \| `"info"` \| `"warn"` \| `"error"` | `"info"`      |
| SUI_DEBUG      |                   `boolean`                    | `false`       |
| SUI_BIN_FOLDER |                    `string`                    | `"bin"`       |
| SUI_DB_FOLDER  |                    `string`                    | `"db"`        |
| SINGBOX_API    |                    `string`                    | -             |

</details>

## SSL Certificate

<details>
  <summary>Click for details</summary>

### Certbot

```bash
snap install core; snap refresh core
snap install --classic certbot
ln -s /snap/bin/certbot /usr/bin/certbot

certbot certonly --standalone --register-unsafely-without-email --non-interactive --agree-tos -d <Your Domain Name>
```

</details>

## Third-party Projects

Community-made projects built around S-UI. These are not affiliated with or maintained by S-UI — use them at your own discretion:

- [itning/reset-s-ui-traffic](https://github.com/itning/reset-s-ui-traffic) — periodic traffic reset for all users
- [zqh2333/s-ui-traffic-reset](https://github.com/zqh2333/s-ui-traffic-reset) — traffic reset tool
- [Sownix21/SUI-Bot](https://github.com/Sownix21/SUI-Bot) - telegram bot

> Building something on top of S-UI (a Telegram bot, monitoring, automation, ...)? Open an issue/PR to get it listed here.

## Stargazers over Time
[![Stargazers over time](https://starchart.cc/alireza0/s-ui.svg)](https://starchart.cc/alireza0/s-ui)

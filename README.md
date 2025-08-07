<sub>[English README here](./README_EN.md)</sub>

IBus Lotus - Bộ gõ tiếng Việt cho Linux
===================================
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://opensource.org/licenses/GPL-3.0)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/LotusInputEngine/ibus-lotus/)

Ibus Lotus là bản fork của [ibus-bamboo](https://github.com/BambooEngine/ibus-bamboo/). Vì một số lý do mà ibus-bamboo không thể tiếp tục phát triển nữa, bộ gõ này ra đời với mục đích kế thừa và tiếp tục phát triển. Quản lý bởi [hien-ngo29](https://github.com/hien-ngo29).

## Installation
1. Cài đặt ibus:
```bash
# Debian/Ubuntu:
sudo apt-get install ibus

# Fedora
sudo dnf install ibus

# CentOS, RHEL, ...
sudo yum install ibus

# openSUSE Tumbleweed
sudo zypper install ibus

# FreeBSD
sudo/doas pkg install ibus
```
2. Tải file `ibus-lotus-<version>.zip` tại phần [Release của Repo](https://github.com/LotusInputEngine/ibus-lotus/releases/).

3. Giải nén file và install (thay `<version>` thành phiên bản đúng mà bạn tải, ví dụ `1.0.0`):
```bash
unzip ibus-lotus-<version>.zip
cd ibus-lotus-<version>

chmod +x ./install
sudo ./install
```

4. Chạy ibus
```
ibus start
```


Các bạn có thể cân nhắc [build từ source](./docs/building_instructions.md).

## Notes
Tuy bộ gõ được fork từ ibus-bamboo và đổi tên thành ibus-lotus nhưng một số đoạn code và script trong bộ gõ vẫn lấy tên ibus-bamboo. Rất tiếc đó là mình khá lười và không mấy hứng thú khi phải nghiên cứu lại hết code chỉ để đổi lại tên. Mình nghĩ sẽ tốt hơn nếu để lại phần việc này cho những thành viên của cộng đồng mã nguồn mở như các bạn. Nếu được hãy fork giúp mình và làm phần việc đó rồi mở Pull Request tại repo này nha <3.

## Xin cám ơn các contributor của IBus Bamboo <3

<a href="https://github.com/LotusInputEngine/ibus-lotus//graphs/contributors">
  <img src="https://contrib.rocks/image?repo=BambooEngine/ibus-bamboo" />
</a>

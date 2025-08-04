1. Cài `ibus`:
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

2. Cài đặt các gói phụ thuộc
- make
- golang
- libgtk-3-dev
- libx11-dev
- libxtst-dev

```bash
# Debian/Ubuntu:
sudo apt-get install make golang libx11-dev libxtst-dev libgtk-3-dev

# Fedora
sudo dnf install make go libX11-devel libXtst-devel gtk3-devel

# CentOS, RHEL, ...
sudo yum install make go libX11-devel libXtst-devel gtk3-devel

# openSUSE Tumbleweed
sudo zypper install make go libX11-devel libXtst-devel gtk3-devel

# FreeBSD
sudo/doas pkg install go pkgconf libX11 libXtst gtk3 bash
```
3. Download
```bash
wget https://github.com/LotusInputEngine/ibus-lotus/archive/master.zip -O ibus-lotus.zip
unzip ibus-lotus.zip

# hoặc clone từ github:
git clone https://github.com/LotusInputEngine/ibus-lotus.git
```
4. Build & install

**Linux**
```bash
cd ibus-lotus
sudo make install PREFIX=/usr

# Restart ibus
ibus restart
```
**FreeBSD**
```bash
cd ibus-lotus
sudo make install PREFIX=/usr/local

# Restart ibus
ibus restart
```

Gỡ cài đặt
======
**Linux**
```bash
sudo make uninstall PREFIX=/usr
ibus restart
```

**FreeBSD**
```bash
sudo make uninstall PREFIX=/usr/local
ibus restart
```
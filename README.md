IBus Lotus - Bộ gõ tiếng Việt cho Linux
===================================
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://opensource.org/licenses/GPL-3.0)
[![GitHub release](https://img.shields.io/github/release/LotusInputEngine/ibus-lotus.svg)](https://github.com/LotusInputEngine/ibus-lotus/releases/latest)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/LotusInputEngine/ibus-lotus/)

Ibus Lotus là bản fork của [ibus-bamboo](https://github.com/BambooEngine/ibus-bamboo/). Vì một số lý do mà ibus-bamboo không thể tiếp tục phát triển nữa, bộ gõ này ra đời với mục đích kế thừa và tiếp tục. Quản lý bởi [hien-ngo29](https://github.com/hien-ngo29).

## Sơ lược tính năng
* Hỗ trợ tất cả các bảng mã phổ biến:
  * Unicode, TCVN (ABC)
  * VIQR, VNI, VPS, VISCII, BK HCM1, BK HCM2,…
  * Unicode UTF-8, Unicode NCR - for Web editors.
* Các kiểu gõ thông dụng:
  * Telex, Telex W, Telex 2, Telex + VNI + VIQR
  * VNI, VIQR, Microsoft layout
* Nhiều tính năng hữu ích, dễ dàng tùy chỉnh:
  * Kiểm tra chính tả (sử dụng từ điển/luật ghép vần)
  * Dấu thanh chuẩn và dấu thanh kiểu mới
  * Bỏ dấu tự do, Gõ tắt,...
  * 2666 emojis từ [emojiOne](https://github.com/joypixels/emojione)
* Sử dụng phím tắt <kbd>Shift</kbd>+<kbd>~</kbd> để loại trừ ứng dụng không dùng bộ gõ, chuyển qua lại giữa các chế độ gõ:
  	* Pre-edit (default)
  	* Surrounding text, IBus ForwardKeyEvent,...
   ![ibus-lotus](./demo.gif)
* Khác với ibus-bamboo, ibus-lotus hiện tại đã hỗ trợ Wayland khá tốt trên 2 Desktop Environment chính đó là GNOME và KDE (Plasma). Chỉ cần yêu cầu các bạn cài một số extension và tool bên thứ ba (Xem hướng dẫn cài đặt bên dưới).

## Installation
Note: vì một số lý do mà mình không kham nổi việc publish ibus-lotus cho các kho package manager của từng distro nên hiện tại các bạn chỉ có 2 lựa chọn để install đó là cài từ phiên bản prebuilt hoặc build từ source

### Cài từ prebuilt
1. ibus-lotus/ibus-bamboo sử dụng IME ibus. Hãy nhớ rõ và chắc rằng bạn đã setup ibus đúng cách. Trên hầu hết các Distro và Desktop Environment thông dụng ibus thường đã được cài sẵn nên bạn có thể bỏ qua bước này.
2. Download file `ibus-lotus-<version>.zip` tại phần [Release của Repo](https://github.com/LotusInputEngine/ibus-lotus/releases/).

3. Giải nén file và install (thay `<version>` thành phiên bản đúng mà bạn đã download, ví dụ `1.0.0`):
```bash
unzip ibus-lotus-<version>.zip
cd ibus-lotus-<version>

chmod +x ./install
sudo ./install
```
Có thể bạn sẽ cần phải log out ra session của Desktop Environment bạn đang dùng và đăng nhập lại để ibus-lotus/ibus-bamboo xuất hiện trên Input Source trong Settings.

4. Thêm Input Source tại Settings của Desktop Environment bạn đang dùng. Bạn sẽ thấy một input source mang tên `Vietnamese (Bamboo)` tại phần Vietnamese (các bạn có thể Google cách thêm Input Source cho Desktop Environment bạn đang dùng).

### Cài từ source code
Xem hướng dẫn [build từ source](./docs/building_instructions.md).

### Note cho Wayland
Hãy cài những thứ này để tránh việc không mở được bảng chọn chế độ gõ trên Wayland.

**GNOME**: cài đặt extension [Window Call Extended](https://extensions.gnome.org/extension/4974/window-calls-extended/)

**KDE Plasma**: cài đặt `kdotool` từ package manager của distro.

**Nếu có thắc mắc hay trục trặc về việc cài đặt ibus-lotus hãy thoải mái [mở issue](https://github.com/LotusInputEngine/ibus-lotus/issues/new) trên repo này**

## Notes
Tuy bộ gõ được fork từ ibus-bamboo và đổi tên thành ibus-lotus nhưng một số đoạn code và script trong bộ gõ vẫn lấy tên ibus-bamboo. Rất tiếc đó là mình khá lười và không mấy hứng thú khi phải nghiên cứu lại hết code chỉ để đổi lại tên. Mình nghĩ sẽ tốt hơn nếu để lại phần việc này cho những thành viên của cộng đồng mã nguồn mở như các bạn. Nếu được hãy fork giúp mình và làm phần việc đó rồi mở Pull Request tại repo này nha <3.

## Xin cám ơn các contributor của IBus Bamboo <3

<a href="https://github.com/LotusInputEngine/ibus-lotus//graphs/contributors">
  <img src="https://contrib.rocks/image?repo=BambooEngine/ibus-bamboo" />
</a>

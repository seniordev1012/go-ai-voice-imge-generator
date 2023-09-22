#
# Copyright (c) 2023. Gundo Sifhufhi
#

go install github.com/fyne-io/fyne-cross@latest

fyne-cross linux --pull
fyne-cross android --pull
fyne-cross windows --pull
fyne-cross macos --pull
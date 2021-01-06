package converter

var km map[string]string
km = make(map[string]string)

km["KC_NO"]= nil
km["XXXXXXX"]= nil
km["KC_TRANSPARENT"] = "___"
km["KC_TRNS"] = "___"
km["_______ "] = "___"
km["KC_A"] = "A"
km["KC_B"] = "B"
km["KC_C"] = "C"
km["KC_D"] = "D"
km["KC_E"] = "E"
km["KC_F"] = "F"
km["KC_G"] = "G"
km["KC_H"] = "H"
km["KC_I"] = "I"
km["KC_J"] = "J"
km["KC_K"] = "K"
km["KC_L"] = "L"
km["KC_M"] = "M"
km["KC_N"] = "N"
km["KC_O"] = "O"
km["KC_P"] = "P"
km["KC_Q"] = "Q"
km["KC_R"] = "R"
km["KC_S"] = "S"
km["KC_T"] = "T"
km["KC_U"] = "U"
km["KC_V"] = "V"
km["KC_W"] = "W"
km["KC_X"] = "X"
km["KC_Y"] = "Y"
km["KC_Z"] = "Z"
km["KC_1"] = "1"
km["KC_2"] = "2"
km["KC_3"] = "3"
km["KC_4"] = "4"
km["KC_5"] = "5"
km["KC_6"] = "6"
km["KC_7"] = "7"
km["KC_8"] = "8"
km["KC_9"] = "9"
km["KC_0"] = "0"
km["KC_ENTER "] = "ENTER"
km["KC_ENT "] = "ENTER"
km["KC_ESCAPE "] = "ESC"
km["KC_ESC "] = "ESC"
km["KC_BSPACE "] = "BACKSPACE"
km["KC_BSPC  "] = "BACKSPACE"
km["KC_TAB "] = "TAB"
km["KC_SPACE "] = "SPACE"
km["KC_SPC "] = "SPACE"
km["KC_MINUS "] = "'-'"
km["KC_MINS "] = "'-'"
km["KC_EQUAL "] = "'='"
km["KC_EQL"] = "'='"
km["KC_LBRACKET "] = "'['"
km["KC_LBRC"] = "'['"
km["KC_RBRACKET"] = "']'"
km["KC_RBRC"] = "']'"
km["KC_BSLASH"] ="'\'"
km["KC_BSLS"] ='\'
/*
km["KC_NONUS_HASH
km["KC_NUHS
Non-US # and ~

km["KC_SCOLON
km["KC_SCLN
; and :

km["KC_QUOTE
km["KC_QUOT
' and "

km["KC_GRAVE
km["KC_GRV, km["KC_ZKHK
` and ~, JIS Zenkaku/Hankaku

km["KC_COMMA
km["KC_COMM
, and <

km["KC_DOT
​
. and >

km["KC_SLASH
km["KC_SLSH
/ and ?

km["KC_CAPSLOCK
km["KC_CLCK, km["KC_CAPS
Caps Lock

km["KC_F1
​
F1

km["KC_F2
​
F2

km["KC_F3
​
F3

km["KC_F4
​
F4

km["KC_F5
​
F5

km["KC_F6
​
F6

km["KC_F7
​
F7

km["KC_F8
​
F8

km["KC_F9
​
F9

km["KC_F10
​
F10

km["KC_F11
​
F11

km["KC_F12
​
F12

km["KC_PSCREEN
km["KC_PSCR
Print Screen
✔
✔2
✔
km["KC_SCROLLLOCK
km["KC_SLCK, km["KC_BRMD
Scroll Lock, Brightness Down (macOS)
✔
✔2
✔
km["KC_PAUSE
km["KC_PAUS, km["KC_BRK, km["KC_BRMU
Pause, Brightness Up (macOS)
✔
✔2
✔
km["KC_INSERT
km["KC_INS
Insert
✔
​
✔
km["KC_HOME
​
Home

km["KC_PGUP
​
Page Up

km["KC_DELETE
km["KC_DEL
Forward Delete

km["KC_END
​
End

km["KC_PGDOWN
km["KC_PGDN
Page Down

km["KC_RIGHT
km["KC_RGHT
Right Arrow

km["KC_LEFT
​
Left Arrow

km["KC_DOWN
​
Down Arrow

km["KC_UP
​
Up Arrow

km["KC_NUMLOCK
km["KC_NLCK
Keypad Num Lock and Clear

km["KC_KP_SLASH
km["KC_PSLS
Keypad /

km["KC_KP_ASTERISK
km["KC_PAST
Keypad *

km["KC_KP_MINUS
km["KC_PMNS
Keypad -

km["KC_KP_PLUS
km["KC_PPLS
Keypad +

km["KC_KP_ENTER
km["KC_PENT
Keypad Enter

km["KC_KP_1
km["KC_P1
Keypad 1 and End

km["KC_KP_2
km["KC_P2
Keypad 2 and Down Arrow

km["KC_KP_3
km["KC_P3
Keypad 3 and Page Down

km["KC_KP_4
km["KC_P4
Keypad 4 and Left Arrow

km["KC_KP_5
km["KC_P5
Keypad 5

km["KC_KP_6
km["KC_P6
Keypad 6 and Right Arrow

km["KC_KP_7
km["KC_P7
Keypad 7 and Home

km["KC_KP_8
km["KC_P8
Keypad 8 and Up Arrow

km["KC_KP_9
km["KC_P9
Keypad 9 and Page Up

km["KC_KP_0
km["KC_P0
Keypad 0 and Insert

km["KC_KP_DOT
km["KC_PDOT
Keypad . and Delete

km["KC_NONUS_BSLASH
km["KC_NUBS
Non-US \ and \|

km["KC_APPLICATION
km["KC_APP
Application (Windows Context Menu Key)
✔
​
✔
km["KC_POWER
​
System Power
​
✔3
✔
km["KC_KP_EQUAL
km["KC_PEQL
Keypad =

km["KC_F13
​
F13

km["KC_F14
​
F14

km["KC_F15
​
F15

km["KC_F16
​
F16

km["KC_F17
​
F17

km["KC_F18
​
F18

km["KC_F19
​
F19

km["KC_F20
​
F20
✔
​
✔
km["KC_F21
​
F21
✔
​
✔
km["KC_F22
​
F22
✔
​
✔
km["KC_F23
​
F23
✔
​
✔
km["KC_F24
​
F24
✔
​
✔
km["KC_EXECUTE
km["KC_EXEC
Execute
​
​
✔
km["KC_HELP
​
Help
​
​
✔
km["KC_MENU
​
Menu
​
​
✔
km["KC_SELECT
km["KC_SLCT
Select
​
​
✔
km["KC_STOP
​
Stop
​
​
✔
km["KC_AGAIN
km["KC_AGIN
Again
​
​
✔
km["KC_UNDO
​
Undo
​
​
✔
km["KC_CUT
​
Cut
​
​
✔
km["KC_COPY
​
Copy
​
​
✔
km["KC_PASTE
km["KC_PSTE
Paste
​
​
✔
km["KC_FIND
​
Find
​
​
✔
km["KC__MUTE
​
Mute
​
✔
✔
km["KC__VOLUP
​
Volume Up
​
✔
✔
km["KC__VOLDOWN
​
Volume Down
​
✔
✔
km["KC_LOCKING_CAPS
km["KC_LCAP
Locking Caps Lock
✔
✔
​
km["KC_LOCKING_NUM
km["KC_LNUM
Locking Num Lock
✔
✔
​
km["KC_LOCKING_SCROLL
km["KC_LSCR
Locking Scroll Lock
✔
✔
​
km["KC_KP_COMMA
km["KC_PCMM
Keypad ,
​
​
✔
km["KC_KP_EQUAL_AS400
​
Keypad = on AS/400 keyboards
​
​
​
km["KC_INT1
km["KC_RO
JIS \ and _
✔
​
✔
km["KC_INT2
km["KC_KANA
JIS Katakana/Hiragana
✔
​
✔
km["KC_INT3
km["KC_JYEN
JIS ¥ and \|
✔
​
✔
km["KC_INT4
km["KC_HENK
JIS Henkan
✔
​
✔
km["KC_INT5
km["KC_MHEN
JIS Muhenkan
✔
​
✔
km["KC_INT6
​
JIS Numpad ,
​
​
✔
km["KC_INT7
​
International 7
​
​
​
km["KC_INT8
​
International 8
​
​
​
km["KC_INT9
​
International 9
​
​
​
km["KC_LANG1
km["KC_HAEN
Hangul/English
​
​
✔
km["KC_LANG2
km["KC_HANJ
Hanja
​
​
✔
km["KC_LANG3
​
JIS Katakana
​
​
✔
km["KC_LANG4
​
JIS Hiragana
​
​
✔
km["KC_LANG5
​
JIS Zenkaku/Hankaku
​
​
✔
km["KC_LANG6
​
Language 6
​
​
​
km["KC_LANG7
​
Language 7
​
​
​
km["KC_LANG8
​
Language 8
​
​
​
km["KC_LANG9
​
Language 9
​
​
​
km["KC_ALT_ERASE
km["KC_ERAS
Alternate Erase
​
​
​
km["KC_SYSREQ
​
SysReq/Attention
​
​
​
km["KC_CANCEL
​
Cancel
​
​
​
km["KC_CLEAR
km["KC_CLR
Clear
​
​
✔
km["KC_PRIOR
​
Prior
​
​
​
km["KC_RETURN
​
Return
​
​
​
km["KC_SEPARATOR
​
Separator
​
​
​
km["KC_OUT
​
Out
​
​
​
km["KC_OPER
​
Oper
​
​
​
km["KC_CLEAR_AGAIN
​
Clear/Again
​
​
​
km["KC_CRSEL
​
CrSel/Props
​
​
​
km["KC_EXSEL
​
ExSel
​
​
​
km["KC_LCTRL
km["KC_LCTL
Left Control

km["KC_LSHIFT
km["KC_LSFT
Left Shift

km["KC_LALT
km["KC_LOPT
Left Alt (Option)

km["KC_LGUI
km["KC_LCMD, km["KC_LWIN
Left GUI (Windows/Command/Meta key)

km["KC_RCTRL
km["KC_RCTL
Right Control

km["KC_RSHIFT
km["KC_RSFT
Right Shift

km["KC_RALT
km["KC_ROPT, km["KC_ALGR
Right Alt (Option/AltGr)

km["KC_RGUI
km["KC_RCMD, km["KC_RWIN
Right GUI (Windows/Command/Meta key)

km["KC_SYSTEM_POWER
km["KC_PWR
System Power Down
✔
✔3
✔
km["KC_SYSTEM_SLEEP
km["KC_SLEP
System Sleep
✔
✔3
✔
km["KC_SYSTEM_WAKE
km["KC_WAKE
System Wake
​
✔3
✔
km["KC_AUDIO_MUTE
km["KC_MUTE
Mute

km["KC_AUDIO_VOL_UP
km["KC_VOLU
Volume Up
✔
✔4
✔
km["KC_AUDIO_VOL_DOWN
km["KC_VOLD
Volume Down
✔
✔4
✔
km["KC_MEDIA_NEXT_TRACK
km["KC_MNXT
Next Track
✔
✔5
✔
km["KC_MEDIA_PREV_TRACK
km["KC_MPRV
Previous Track
✔
✔5
✔
km["KC_MEDIA_STOP
km["KC_MSTP
Stop Track
✔
​
✔
km["KC_MEDIA_PLAY_PAUSE
km["KC_MPLY
Play/Pause Track

km["KC_MEDIA_SELECT
km["KC_MSEL
Launch Media Player
✔
​
✔
km["KC_MEDIA_EJECT
km["KC_EJCT
Eject
​
✔
✔
km["KC_MAIL
​
Launch Mail
✔
​
✔
km["KC_CALCULATOR
km["KC_CALC
Launch Calculator
✔
​
✔
km["KC_MY_COMPUTER
km["KC_MYCM
Launch My Computer
✔
​
✔
km["KC_WWW_SEARCH
km["KC_WSCH
Browser Search
✔
​
✔
km["KC_WWW_HOME
km["KC_WHOM
Browser Home
✔
​
✔
km["KC_WWW_BACK
km["KC_WBAK
Browser Back
✔
​
✔
km["KC_WWW_FORWARD
km["KC_WFWD
Browser Forward
✔
​
✔
km["KC_WWW_STOP
km["KC_WSTP
Browser Stop
✔
​
✔
km["KC_WWW_REFRESH
km["KC_WREF
Browser Refresh
✔
​
✔
km["KC_WWW_FAVORITES
km["KC_WFAV
Browser Favorites
✔
​
✔
km["KC_MEDIA_FAST_FORWARD
km["KC_MFFD
Next Track
✔
✔5
✔
km["KC_MEDIA_REWIND
km["KC_MRWD
Previous Track
✔6
✔5
✔
km["KC_BRIGHTNESS_UP
km["KC_BRIU
Brightness Up

km["KC_BRIGHTNESS_DOWN
km["KC_BRID
Brightness Down
/*
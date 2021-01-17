from PYKB import *
import time

keyboard = Keyboard()

# ESC(0)    1(1)   2(2)   3(3)   4(4)   5(5)   6(6)   7(7)   8(8)   9(9)   0(10)  -(11)  =(12)  BACKSPACE(13)
# TAB(27)   Q(26)  W(25)  E(24)  R(23)  T(22)  Y(21)  U(20)  I(19)  O(18)  P(17)  [(16)  ](15)   \(14)
# CAPS(28)  A(29)  S(30)  D(31)  F(32)  G(33)  H(34)  J(35)  K(36)  L(37)  ;(38)  "(39)      ENTER(40)
# LSHIFT(52) Z(51)  X(50)  C(49)  V(48)  B(47)  N(46)  M(45)  ,(44)  .(43)  /(42)            RSHIFT(41)
# LCTRL(53)  LGUI(54)  LALT(55)               SPACE(56)          RALT(57)  MENU(58)  Fn(59)  RCTRL(60)

{{.Keymap}}


def set_keyboard_default(dev):
    set_color(dev, 255, 64, 0)
    set_wasd(dev, 255, 255, 0)
    set_special_k(dev, 255, 0, 255)


def set_color(dev, r, g, b):
    for i in range(63):
        dev.backlight.pixel(i, r, g, b)
    dev.backlight.update()


def set_wasd(dev, r, g, b):
    for i in (25, 29, 30, 31):
        dev.backlight.pixel(i, r, g, b)
    dev.backlight.update()


def set_special_k(dev, r, g, b):
    for i in (0, 13, 27, 28, 40, 52, 41, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62):
        dev.backlight.pixel(i, r, g, b)
    dev.backlight.update()


{{.Macros}}
        elif n == 20:
            set_color(dev, 255, 0, 0)
            set_wasd(dev, 0, 255, 0)
        elif n == 21:
            set_color(dev, 0, 255, 0)
            set_wasd(dev, 255, 0, 0)
        elif n == 22:
            set_color(dev, 0, 0, 255)
            set_wasd(dev, 255, 255, 0)
        elif n == 23:
            set_color(dev, 0, 0, 0)
        elif n == 24:
            set_keyboard_default(dev)

        else:
            pass


def pairs_handler(dev, n):
    if n == 0:
        dev.send(GUI, ENTER)
        time.sleep(1)
        dev.send_text("setxkbmap -layout us -variant intl")
        dev.send(ENTER)
        dev.send_text("exit")
        dev.send(ENTER)
    else:
        dev.send(GUI, ENTER)
        time.sleep(1)
        dev.send_text("spotify")
        dev.send(ENTER)


keyboard.macro_handler = macro_handler
keyboard.pairs_handler = pairs_handler

# Pairs: J & K, U & I
keyboard.pairs = [{37, 36}, {20, 19}]

keyboard.verbose = False

set_keyboard_default(keyboard)

keyboard.run()

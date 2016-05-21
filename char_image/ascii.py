import argparse
from PIL import Image

ascii_char = list("$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\|()1{}[]?-_+~<>i!lI;:,\"^`'. ")


def get_char(r, g, b, alpha=256):
    if alpha == 0:
        return ' '
    length = len(ascii_char)
    gray = int(0.2126 * r + 0.7152 * g + 0.0722 * b)
    unit = (256.0 + 1) / length
    return ascii_char[int(gray / unit)]

IMG = 'ascii_dora.png'
im = Image.open(IMG)
im = im.resize((80, 80), Image.NEAREST)
txt = ''
for i in range(80):
    for j in range(80):
        txt += get_char(*im.getpixel((j,i)))
    txt += '\n'

print txt

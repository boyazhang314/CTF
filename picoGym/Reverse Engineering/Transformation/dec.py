enc = '灩捯䍔䙻ㄶ形楴獟楮獴㌴摟潦弸弰㑣〷㘰摽'

split = list(enc)

flag = ''
for s in split:
    flag += chr(ord(s) >> 8) # first character
    flag += chr(ord(s) % 2**8) # second character

print(flag)

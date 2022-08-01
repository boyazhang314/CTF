# Nice netcat...
**Points: 15**

## Description
There is a nice program that you can talk to by using this command in a shell: ```$ nc mercury.picoctf.net 43239```, but it doesn't speak English...

## Hints
<details>
<summary>Hint 1</summary>

> You can practice using netcat with this picoGym problem: [what's a netcat?](https://play.picoctf.org/practice/challenge/34)
</details>

<details>
<summary>Hint 2</summary>

> You can practice reading and writing ASCII with this picoGym problem: [Let's Warm Up](https://play.picoctf.org/practice/challenge/22)
</details>

## Solution
Netcat is a utility to read from and write to network connections. As promted, run the [netcat](https://linux.die.net/man/1/nc) command in a shell ```nc mercury.picoctf.net 43239``` to be greeted by a long list of numbers
```
112 
105 
99 
111 
67 
84 
70 
123 
103 
48 
48 
100 
95 
107 
49 
116 
116 
121 
33 
95 
110 
49 
99 
51 
95 
107 
49 
116 
116 
121 
33 
95 
55 
99 
48 
56 
50 
49 
102 
53 
125 
10 
```

At a glance, these numbers appear to be between roughly 30 and 128, implying that they most likely represent ASCII values. We can use any integer to ASCII converter to obtain the flag

## Flag 
```picoCTF{g00d_k1tty!_n1c3_k1tty!_7c0821f5}```
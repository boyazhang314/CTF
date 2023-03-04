# Missing Flag

## Description

> R-Team,
>
> We found an ancient Word document left off on one of the OmniFlags' computers. Does it contains any information useful to us?
>
> HQ

## Word

We are given a Word document, which opening reveals a long string of numbers:

`1090970103011201050101012301190104012109501000490100095049095048011205101100950116010404901150950109052099011404809508408708401250`

This appears to be a cipher or code of some sort. On careful inspection, we can see what appears to be ASCII numerical values, such as 109, 97, 103 etc. only they are separated by 0's. Note that ASCII numbers are at least 2 digits long.

Let's replace these values with spaces:

`109 97 103 112 105 101 123 119 104 121 95 100 49 100 95 49 95 48 112 51 110 95 116 104 49 115 95 109 52 99 114 48 95 84 87 84 125`

Convert this to ASCII to get the flag

## Flag

`magpie{why_d1d_1_0p3n_th1s_m4cr0_TWT}`

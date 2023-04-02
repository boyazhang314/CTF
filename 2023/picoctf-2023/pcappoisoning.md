# PcapPoisoning

## Description

> How about some hide and seek heh?
>
> Download this [file](https://artifacts.picoctf.net/c/371/trace.pcap) and find the flag.

## Get the Antidote

We get a `trace.pcap` file

Though we can open the file in Wireshark, let's first try using `string` to search through the contents

```
strings trace.pcap | grep "picoCTF"
```

## Flag

`picoCTF{P64P_4N4L7S1S_SU55355FUL_62558951}`

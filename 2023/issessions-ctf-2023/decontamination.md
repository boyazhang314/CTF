# Decontamination

## Description

> Time Corp is currently undergoing recovery from a major breach. In a last ditch effort to maintain access, the hackers dropped malicious files at various points in the network. While most of them were caught, the Incident Response team believes there may still be undiscovered files. Can you find the malicious file in the provided list? Flag Format: retroCTF{FILENAME\_HERE}

## Find the Trash

We have a `SuspiciousFiles.txt` file and a `DetectedSamplesList.txt` which has

```
Payroll_AF19FD5349.docm
PayrollE8ADFBBA4134.xlsm
Payroll_C32BADE1.xlsm
PayrollBE349BE24.docm
Payroll_ABCDEFABCD.xlsm
Payroll1234567890.docm
```

We can make the following observations about the malicious file:

* Ends with `.docm` or `.xlsm`
* Starts with `Payroll`
* `Payroll` could be followed by an underscore `_` and 8 - 12 characters `A-F0-9`

Can create a regular expression such as `^Payroll_?[A-F0-9]{8,12}.(docm|xlsm)$`

* `^` - Beginning of string
* `Payroll` - Starts with "Payroll"
* `_?` - Can have 0 or 1 "\_"
* `[A-F0-9]{8,12}` - Characters from "A" to "F" or "0" to "9" are used 8 to 12 times
* `.(docm|xlsm)` - Either ".docm" or ".xlsm"
* `$` - End of string

Using it to search, we find there is only one result: `Payroll_EA026F52BAF.xlsm`

<figure><img src="../../.gitbook/assets/image (18).png" alt=""><figcaption></figcaption></figure>

## Flag

`retroCTF{Payroll_EA026F52BAF.xlsm}`

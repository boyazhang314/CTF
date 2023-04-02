# two-sum

## Description

> Can you solve this?
>
> What two positive numbers can make this possible: `n1 > n1 + n2 OR n2 > n1 + n2`
>
> Enter them here `nc saturn.picoctf.net 61200`. [Source](https://artifacts.picoctf.net/c/456/flag.c)

## Hints

<details>

<summary>1</summary>

Integer overflow

</details>

<details>

<summary>2</summary>

Not necessarily a math problem

</details>

## Algebra

Of course `n1 > n1 + n2 OR n2 > n1 + n2` is impossible for positive integers `n1` and `n2` so let's look into the `flag.c` source file

First we have a function which adds two integers

```c
static int addIntOvf(int result, int a, int b) {
    result = a + b;
    if(a > 0 && b > 0 && result < 0)
        return -1;
    if(a < 0 && b < 0 && result > 0)
        return -1;
    return 0;
}
```

It returns -1 if `a` and `b` are negative but `a + b` is positive, or `a` and `b` are positive but `a + b` is negative

Then in the main function we have the following if statment which may print the flag

```c
if (scanf("%d", &num1) && scanf("%d", &num2)) {
    printf("You entered %d and %d\n", num1, num2);
    fflush(stdout);
    sum = num1 + num2;
    if (addIntOvf(sum, num1, num2) == 0) {
        printf("No overflow\n");
        fflush(stdout);
        exit(0);
    } else if (addIntOvf(sum, num1, num2) == -1) {
        printf("You have an integer overflow\n");
        fflush(stdout);
    }

    if (num1 > 0 || num2 > 0) {
        flag = fopen("flag.txt","r");
        if(flag == NULL){
            printf("flag not found: please run this on the server\n");
            fflush(stdout);
            exit(0);
        }
        char buf[60];
        fgets(buf, 59, flag);
        printf("YOUR FLAG IS: %s\n", buf);
        fflush(stdout);
        exit(0);
    }
}
```

* Reads in numbers `num1` and `num2` from input and gets the `sum = num1 + num2`
* If `addIntOvf(sum, num1, num2)` is 0, then there is "No overflow" and the program exits
* Otherwise `addIntOvf(sum, num1, num2)` is -1 and there's an integer overflow, so the program continues
* Then the program checks if at least one of `num1` and `num2` are positive, then it will print out the flag for us

So we need `addIntOvf` to return -1 for two positive integers. Of course, this isn't possible with normal integers, but we can try to remedy this with an _integer overflow_

<figure><img src="../../.gitbook/assets/image (9).png" alt=""><figcaption></figcaption></figure>

In C, the maximum possible value of an int is 2147483647

We can try passing in 2147483647 and 1, so the sum causes a buffer overflow and the result is negative, which will return -1 for `addIntOvf`

<figure><img src="../../.gitbook/assets/image (1).png" alt=""><figcaption></figcaption></figure>

## Flag

`picoCTF{Tw0_Sum_Integer_Bu773R_0v3rfl0w_ccd078bd}`

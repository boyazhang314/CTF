# Greatest Common Divisor

## Description

> The Greatest Common Divisor (GCD), sometimes known as the highest common factor, is the largest number which divides two positive integers (a,b).\
> \
> For `a = 12, b = 8` we can calculate the divisors of a: `{1,2,3,4,6,12}` and the divisors of b: `{1,2,4,8}`. Comparing these two, we see that `gcd(a,b) = 4`.\
> \
> Now imagine we take `a = 11, b = 17`. Both `a` and `b` are prime numbers. As a prime number has only itself and `1` as divisors, `gcd(a,b) = 1`.\
> \
> We say that for any two integers `a,b`, if `gcd(a,b) = 1` then `a` and `b` are coprime integers.\
> \
> If `a` and `b` are prime, they are also coprime. If `a` is prime and `b < a` then `a` and `b` are coprime.\
> \
> :high\_brightness: Think about the case for `a` prime and `b > a`, why are these not necessarily coprime?\
> \
> There are many tools to calculate the GCD of two integers, but for this task we recommend looking up [Euclid's Algorithm](https://en.wikipedia.org/wiki/Euclidean\_algorithm).\
> \
> Try coding it up; it's only a couple of lines. Use `a = 12, b = 8` to test it.\
> \
> Now calculate `gcd(a,b)` for `a = 66528, b = 52920` and enter it below.

## Mr. Euclid, Sir

To answer the first quiz question, if `b > a` and `a` is prime, then it is possible that `b` is a multiple of `a`, so the GCD of these two numbers is `a` and they are not coprime.&#x20;

Let's take a look at Euclid's Algorithm

### Euclidean Algorithm

We are given two positive integers `a` and `b`, assume that `a > b`

1. Divide `a / b` to get the remainder `r` and if `r = 0`, `b` is the GCD of `a, b`
2. Else, replace `a` with `b` and `b` with `r`, then repeat

<details>

<summary>Proof</summary>

We can prove this by induction, but intuitively this algorithm will eventually terminate as the remainder `r` will inevitably decrease and eventually reach 0

Equivalently, we have $$a = q_0b+r_0 \rightarrow b = q_1r_0 +r_1 \rightarrow r_0=q_2r_1+r_2 \cdots$$

Eventually, we get $$r_n=0$$ which implies $$r_{n-1}=gcd(a,b)$$

We wish to prove $$r_{n-1}=gcd(a,b)=g$$

1.  $$r_{n-1}\leq g$$

    Since $$r_n=0$$ we know $$r_{n-2}=q_nr_{n-1}$$ and $$r_{n-1}$$ divides $$r_{n-2}$$

    Thus we can also get $$r_{n-3}=q_{n-1}r_{n-2}+r_{n-1}=q_{n-1}q_nr_{n-1}+r_{n-1}$$ and $$r_{n-1}$$ also divides $$r_{n-3}$$

    We can repeat this to get that $$r_{n-1}$$ divides $$a$$ and $$b$$, and thus is a common divisor\
    It follows that  $$r_{n-1} \leq g$$, the greatest common divisor
2.  &#x20;$$r_{n-1}\geq g$$

    For any common divisor $$c$$ we have $$a=mc$$ and   $$b=nc$$

    Then  $$c$$ divides  $$r_0$$ since  $$r_0=a-q_0b=mc-q_0nc=(m-q_0n)c$$

    We can repeat this argument to get that  $$c$$ divides  $$r_{n-1}$$, so it follows that  $$g$$ divides  $$r_{n-1}$$ as it is also a common divisor\
    It follows that  $$r_{n-1}\geq g$$

From both of these statements, we conclude  $$r_{n-1}=g$$ &#x20;

</details>

### Implementation

```python
def gcd(a, b):
    # Only works if a > b
    if a < b:
        return gcd(b, a)

    # Euclidean algorithm
    while b != 0:
        # Calculate r = a % b
        # Set a = b
        # Set b = r
        a, b = b, a % b
    
    return a
    
print(gcd(66528, 52920))
```

Though of course, if you're lazy you can use an online tool

## Flag

1512

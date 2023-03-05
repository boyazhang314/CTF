# Rubis

## Description

> Hello C-Team,
>
> During our investigation of Omniflags, we came across a cluster of ships called Rubis that was used to store top-secret information randomly in one of the ships and would rotate this secret every few days to stop attackers from finding the secret. It has been long since Rubric was deprecated however, we found out Rubis still contains a key component that will assist in the rescue of the prisoners from the space prison. We managed to get the source code that was used to run Rubis and for better or for worse, Rubis is still running at `<ip>:<port>`. We need you to unlock the ship that contains the flag so we can save the prisoners.
>
> Godspeed, HQ

## Ship Ahoy!

Connecting to the server, we have a few options:

```
What would you like to do?
1. Get previous ships
2. Get public certificate
3. Get password hint
4. Unlock a ship
```

1. Prints out 5 seemingly random numbers
2. Prints out `n` and `e`
3. Prints out a hex string
4. Asks for a ship number and password, unlocks the ship if the password is correct

We are also given a `challenge.py` file which consists of the code running this server, so let's check it out

First, we have an `RSA` class, which I annotated with ###

```python
class RSA():
    def __init__(self, key_length): ### Takes in the key length
        self.e = 0x10001 ### e initialized to 65537
        phi = 0
        prime_length = key_length // 2
        ### Regenerate primes while gcd(e, phi) != 1
        while GCD(self.e, phi) != 1:
            x = getRandomNBitInteger(prime_length) ### Random x
            self.p = nextprime(x) ### p is the next prime after x
            self.q = prevprime(x) ### q is the previous prime before x
            phi = (self.p - 1) * (self.q - 1) ### Define phi = (p - 1)(q - 1)
            self.n = self.p * self.q ### Initialize n = pq
        self.d = inverse(self.e, phi) ### d = e^(-1) mod phi

    def encrypt(self, message):
        message = bytes_to_long(message)
        return pow(message, self.e, self.n) ### c = m^e mod n

    def decrypt(self, encrypted_message):
        message = pow(encrypted_message, self.d, self.n) ### m = c^d mod n
        return long_to_bytes(message)
```

For the most part, `__init__` describes the initialization of RSA, two large primes `p` and `q`, an integer `e` such that `gcd(e, phi) = 1` where `phi = (p - 1)(q - 1)`, `n = pq`, and the secret key `d` is the inverse of `e` in modular `n`

Encryption and decryption follow the same steps as normal RSA

What's peculiar is that instead of randomly generating `p`,`q` first, `e` is initialized and `p`,`q` are chosen until they satisfy the GCD requirement. Also, the two primes are not truly randomly generated, as they are both dependent on `x` and thus are relatively close primes. This means `n` is likely easy to factor

Secondly, there is a `Ship` class, which is straightforward

```python
class Ship():
    ### Ships have a password and data
    def __init__(self, password, data):
        self.password = password
        self.data = data
```

Finally, we have a `Cluster` class

```python
class Cluster():
    def __init__(self, rsa, password, flag):
        self.rsa = rsa ### Initialize RSA with the RSA class
        self.ships = {} ### Empty ships
        for i in range(1000000):
            ### Initialize 1000000 ships
            ### Each ship has password "pa$$w0rd" and data "Nothing to see here :( Better luck next time!"
            self.ships[i] = Ship(b"pa$$w0rd", b"Nothing to see here :( Better luck next time!")
        self.prevShips = [] ### Previous ships are empty
        self.currentShip = 0 ### Current ship is at index 0
        random.seed(round(time.time())) ### Set a random seed
        self.masterPassword = password ### Set the password
        self.setCurrentShip() ### Initialize previous ship and current ship
        self.populateSpecialShip(password, flag) ### Set current ship to the password and flag

    def setCurrentShip(self):
        ### Initialize previous ships to random values between 1 and 1000000
        for i in range(5):
            self.prevShips.append(random.randrange(1000000))
        ### Initialize current ship to the next random value between 1 and 1000000
        self.currentShip = random.randrange(1000000)

    def populateSpecialShip(self, password, flag):
        self.ships[self.currentShip].password = password
        self.ships[self.currentShip].data = flag
```

From this, we make a few observations

* All ships, except for the current ship, has password `pa$$w0rd` and data `Nothing to see here :( Better luck next time!`
* The current ship is the "special ship" we want, with a password we don't know and the data is the flag!
* The previous ships are 5 randomly generated values from 1 to 1000000 generated before generating the value for the current ship

We also have some class functions, which corresponds to the input options we have on the server

```python
def getPrevShips(self):
        return self.prevShips ### Returns previous ships
```

This is called by input 1. Get previous ships

```python
def getPasswordHint(self):
        return hex(self.rsa.encrypt(self.masterPassword))
```

This is called by input 3. Get password hint

* It encrypts the master password with RSA then turns it into a hex string

```python
def unlockShip(self, ship, password):
        if (ship < 0 or ship > 1000000):
            print(f"No ship exists with id: {ship}")
            exit(1)
        elif (self.ships[ship].password == password):
            print(
                f"You've managed to unlock ship {ship}! Here's what's inside: \n{self.ships[ship].data.decode()}")
        else:
            print("Err... couldn't open the ship! Maybe next time...\n")
            exit(1)
```

This is called by input 4. Unlock a ship

* It checks if the password is same as the ship's password, then prints the data if it is
* It exists if we fail to open the ship, so it is not possible to brute force

Now we have all the information we need, let's try to get the flag!

{% hint style="warning" %}
**Note** - Be careful when restarting the session, as the RSA values and previous ships are randomly generated on bootup and will be different each time. If we need to restart, we'll need to go through all the steps again!
{% endhint %}

We will need the password, so let's get that. Start a session than choose option 2 to get the value of `n` and `e`. If we want to decrypt the password, we need the secret key `d`

Let's try factoring `n` to get `p` and `q`

{% embed url="https://www.alpertron.com.ar/ECM.HTM" %}

It runs pretty fast due to the poor generation of `p` and `q`

We can then calculate `phi = (p - 1)(q - 1)` and `d = e^(-1) mod phi`

```python
from Crypto.Util.number import inverse

phi = (p - 1) * (q - 1)
d = inverse(e, phi)
```

Now we have the secret key for RSA, so let's get the password hint with input 3. Since it's a hex string, we need to convert it to its numerical value

```python
from Crypto.Util.number import long_to_bytes

hex = '{PASSWORD HINT}'
enc = int(hex, 0) ### Convert into number
key = pow(enc, d, n) ### Decrypt the password
print(long_to_bytes(key))
```

We get that the password is `f32m47'5_f4c702124710n_m37h0d_f02_7h3_w1n!!`

So we have the password, now we need to find the special ship to get the flag. How will we do that? Aren't the ships randomly generated? Also `unlockShip` will call `exit(1)` if we fail to unlock a ship, so we can't try them all until our password works :(

Now I had to think what option 1. Get the previous ships meant, since it was the only one we haven't used yet. It seemingly has nothing to do with the current ship as it is randomly and independently generating numbers for the previous ship, then the current ship. At least, in theory...

Consider `random.seed(round(time.time()))`

This sets a seed, but what does that mean? In reality, there is no such thing as true random. For the most part we use pseudorandom generators, which are a sequence of seemingly random numbers. For any one seed, there is a different sequence of numbers

But if we have the same seed, we will have the exact same sequence of numbers. This is why you need to change seeds on every use

We can try to exploit this fact and try to get the same seed that the server uses on bootup. Check the time when you run the server and keep track of it. Then, get the previous ships. We will compare different times as seeds until the first 5 numbers match the previous ships:

```python
import random, time
time = 1677361020 ### Roughly the time I ran the server
targets = [862117, 375147, 59052, 38858, 261625] ### Previous ships

while True:
    accept = True
    random.seed(time)
    for i in range(5): ### Check if the first 5 numbers are the same
        temp = random.randrange(1000000)
        if temp != targets[i]:
            accept = False
            break
    if accept:
        print(time)
        random.seed(time)
        for i in range(5):
            print(random.randrange(1000000))
        break
    time = time + 1 ### Test all possible times until it works

print("\n")
random.seed(time)
for i in range(6):
    print(random.randrange(1000000)) ### Get the 6th number
```

Now we have the 6th number generated by the seed, which must be our current ship! Select option 4, then pass in the ship number and the password to get the flag

## Flag

`magpie{17_w45n'7_72u1y_24nd0m_4f732_411}`

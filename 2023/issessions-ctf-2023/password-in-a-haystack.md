# Password in A Haystack

> One of our Time Machine developers decided to change their password manager. Unfortunately, during the exporting process, it seems the program messed up and didn't include the usernames and sites for each password. Now, they're locked out of the company GitHub and have too many passwords to reasonably check. All they remember is that it followed corporate policy. Can you help them recover their password?

## Needle

We have `credential_export.txt`, a list of 25000 possible passwords, and `PasswordPolicy.pdf`

The most notable part of the PDF is this table:

<figure><img src="../../.gitbook/assets/image (8).png" alt=""><figcaption></figcaption></figure>

Let us filter the passwords using these rules

* `cat credential_export.txt | grep -v ' ' > nospace`
* `cat nospace | grep -v ',' > nocomma`
* `cat nocomma | grep '^(.[^A-Za-z0-9]){3,}.$' > special`
* `cat special | grep '^(.[A-Z]){4,}.$' > capital`
* `cat capital | grep '^(.[a-z]){4,}.$' > lower`
* `cat lower | grep '^.{25,33}$' > length`
* `cat length | grep '^.[0-9].$' > number`

We end up with a single password `g72$89Hsu(!haHasi-zx89yJKn218sH(`

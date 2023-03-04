# No Password Here

## Description

> Hello B-Team,
>
> Here is one of the vaults that have a flag we need. There's a message associated with the vault, written as follows:
>
> > Don't even think about trying the password because I do not know the password either. I already informed my employees not to use "gets" when writing programs so what could possibly go wrong?
>
> Good luck,\
> HQ

## Yes Password There

We are given a `Code.c` source code file

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
void main()
{
	char flag[120];
	setvbuf ( stdout, NULL , _IONBF , 0 );
	char Test[20];

	//random number based on the current time.
	//YOU WILL NEVER GUESS THE PASSWORD. HAHAHAHAHHAH
    srand(time(0));
	sprintf(Test, "%d",rand());	
	
	
    char input[20];
    printf("Enter something?");
    scanf("%s",input);

	//Check password
	if (strncmp(Test,input,20) == 0)
	{
		FILE *f = fopen("flag.txt","r");
		
		fgets(flag,100,f);
		
		printf("Password is correct! Here is your flag: %s", flag);
	}
	

}
```

As cruelly pointed out by the comments, the password is quite random and likely not reasonably guessable, so let's try looking into the replacement function for "gets": `scanf`

We have an char `input` buffer of size 20, however what happens if our input exceeds this value?

Our goal is to get `strncmp(Test,input,20) == 0` to be true, so the program will open `flag.txt` and print it out for us. How can we do that? There's no way we can guess what `Test` is, so what else can we do?

Let's think about the memory stack for a moment, since we know `Test` was stored on it. First, we allocate a 120 character array for `flag`. Then, we allocate a 20 character array for `Test`. Finally, we allocate a 20 character array for `input`

Our stack looks something like this:

```
 0 ...            20 0 ...           20 0 ...              120  
| ----- input ----- | ----- Test ----- | ------- flag ------- |
```

If we send in some user input, `scanf` will consume the allocated memory like so:

```
 0 ...            20 0 ...           20 0 ...              120  
| Hello\0 ...       | ----- Test ----- | ------- flag ------- |
```

However, if our input exceeds the boundaries for `input` we will overwrite the password `Test`, and if we do enough, we can get `input` and `Test` to be the same!

For our payload, spam in a bunch of "A"s and the flag will print

## Flag

`magpie{5c4nf_n07_54f3}`

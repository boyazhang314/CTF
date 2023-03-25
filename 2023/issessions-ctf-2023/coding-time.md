# Coding Time

## Description

> This challenge will require you to code a line or two to get the flag.

## No, Not Code :(

We are given a Python file `passcracker.py`

```python
def main():
    password = "retroCTF{"
    passw0rd = "1br3j13dt7ro__Cp6T4F_53a5ytw4h07er5ydcg_45ck3r1v4gfc0ukqb33sr6n_lv6578axdorje2vrtso193e825"
    userInput = input("Enter the password: ")
    #DO NOT CHANGE THE CODE ABOVE THIS LINE
    
    #TODO: Write your code here
    

    #DO NOT CHANGE ANYTHING BELOW THIS LINE
    if userInput[:-1].__eq__(password):
        print("Correct password!")
    else:
        print("Wrong password!")

if __name__ == "__main__":
    main()


"""
NOTE:
1. The password is in the form of retroCTF{...}
2. You must implement code that can manipulate the string passw0rd... try to google some python functions that can help you
3. Once you figure out step 2, you must figure out a way to concatenate the string passw0rd to the string password
"""
```

As outlined by the `NOTE`, our goal is manipulate `passw0rd` to somehow to get the content in between the curly braces of the flag

Just from appearances, `passw0rd` looks too long to be the contents of the flag. It also doesn't look like a cipher or hash. The underscores seem to imply that some of these characters are the actual characters of the flag, so we can guess that we need to extract the characters for the flag

With Python, we can get every nth letter with `[::n]` at the end of the string. Testing some numbers, we find that we need every 3rd letter, which we can concatenate to make the flag

```python
print(password + passw0rd[::3] + '}')
```

## Flag

`retroCTF{1337_p455w0rd_cr4ck3r_68djvs98}`

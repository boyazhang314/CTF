# MatchTheRegex

## Description

> How about trying to match a regular expression

## Hints

<details>

<summary>1</summary>

any redirections?

</details>

## Pattern Matching

We have a website with a single input field

<figure><img src="../../.gitbook/assets/Untitled (4).png" alt=""><figcaption></figcaption></figure>

Any incorrect input gives us an angry alert

<figure><img src="../../.gitbook/assets/image (37).png" alt=""><figcaption></figcaption></figure>

Looking through the code, we find this interesting function which appears to validate the input

```javascript
function send_request() {
	let val = document.getElementById("name").value;
	// ^p.....F!?
	fetch(`/flag?input=${val}`)
		.then(res => res.text())
		.then(res => {
			const res_json = JSON.parse(res);
			alert(res_json.flag)
			return false;
		})
	return false;
}
```

The regex is `^p.....F!?`

* `^p` - String starts with "p"
* `.....` - 5 characters of any value
* `F` - The letter "F"
* `!?` - Matches 0 or 1 occurrences of "!"

Using this, we can pass in "picoCTF" or any other possible string that satisfies the regex

<figure><img src="../../.gitbook/assets/image (38).png" alt=""><figcaption></figcaption></figure>

## Flag

`picoCTF{succ3ssfully_matchtheregex_8ad436ed}`

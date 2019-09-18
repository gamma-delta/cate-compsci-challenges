# Challenge: Ledger (wip)

Hai~~ Good for you for snooping around!

Eventually we'll be posting some bigger challenges that last longer than a week. This is the code that generates the base for one such challenge. Feel free to look around; we'll be adding the fraudulent transaction (see below) by hand when we publish it for real.

The challenge is:

Attached is a ledger showing logs of transactions. For example,
```
"Frankie": [
    {
        "name": "Claudia",
        "amount": 19
    },
    {
        "name": "Devon",
        "amount": -8
    },
    {
        "name": "Alice",
        "amount": 57
    }
]
```
means Frankie recieved \$19 from Claudia, then paid \$8 to Devon, then recieved \$57 from Alice.

Of course, these transactions should be balanced. For example, if Claudia pays Alice \$20, Alice should recieve \$20 from Claudia and both sides of the transaction should be marked in the log.

There is one fraudulent transaction, where someone has marked a transaction in their log without the other person marking it in their log. Who is the fraud and what was the transdaction?
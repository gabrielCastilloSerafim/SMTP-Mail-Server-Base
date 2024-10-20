# Mail-Server

**Request example:**

**POST** `http://localhost:3001/v1/sendMail`

*Body:*

```
{
"addresses": ["example@gmail.com"], // NonOptional, must contain at least one email
"subject": "The mail subject", // Optional
"body": "The mail body" // NonOptional
}
```

# Mail-Server

## Request example:

**POST** `http://localhost:3001/v1/sendMail`

_Body:_

```
{
"addresses": ["example@gmail.com"], // NonOptional, must contain at least one email
"subject": "The mail subject", // Optional
"body": "The mail body" // NonOptional
"provider": 0 // The provider that we want to send the email from must match internal server enum
}
```

# Mail-Server

**POST** https://mail-server-production-2cba.up.railway.app/v1/sendMail

**Body:**

```
{
	`"addresses": ["sk8.andrade@gmail.com"], // NonOptional, must contain at least one email`
	`"subject": "The mail subject", // Optional`
	`"body": "Final test2" // NonOptional`
}
```

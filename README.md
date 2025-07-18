# Mail-Server

## Request example:

**POST** `http://localhost:3001/v1/sendMail`

_Headers:_
```
["apiKey": "1520AE6F-E377-41AE-BD90-9A4191F2F8CA"]
```

_Body:_

```
{
"addresses": ["example@gmail.com"], // NonOptional, must contain at least one email
"subject": "The mail subject", // Optional
"body": "The mail body", // NonOptional
"isHTML": false, // Optional, defaults to false
"provider": 0 // The provider that we want to send the email from must match internal server enum
}
```

## curl example:

```
curl --location 'http://localhost:3001/v1/sendMail' \
--header 'apiKey: 1520AE6F-E377-41AE-BD90-9A4191F2F8CA' \
--data-raw '{
    "addresses": [
        "sk8.andrade@gmail.com"
    ],
    "subject": "Test",
    "body": "<!DOCTYPE html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Link Example</title></head><body style=\"margin: 0; padding: 0; background-color: #f4f4f4;\"><table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" width=\"100%\"><tr><td style=\"padding: 20px 0;\"><table align=\"center\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" width=\"600\" style=\"border-collapse: collapse; background-color: #ffffff;\"><tr><td style=\"padding: 40px 30px; font-family: Arial, sans-serif; font-size: 16px; line-height: 24px;\"><p>Please visit this <a href=\"https://www.youtube.com/watch?v=8UFYvCysMEk\" style=\"color: #1a73e8;\">link</a>.</p></td></tr></table></td></tr></table></body></html>",
    "provider": 0,
    "isHTML": true
}'
```

## To send HTML you must send it as plain text and scaping the "" with \ like the following example:

```
"<!DOCTYPE html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Hello World</title><style>body{background-color:#121212;color:#e0e0e0;display:flex;justify-content:center;align-items:center;height:100vh;margin:0;font-family:sans-serif;}</style></head><body><h1>Hello, World!</h1></body></html>"
```

## How to add support for a nem Email:

To use Gmail for SMTP, you need your email address and a special App Password, not your regular Google account password.

You can only generate an App Password after enabling 2-Step Verification on your Google Account.

- Step 1: Enable 2-Step Verification
If you don't already have it enabled, you must turn on 2-Step Verification.

Go to the Google Account security page: myaccount.google.com/security

Click on 2-Step Verification (under "How you sign in to Google").

Follow the on-screen steps to enable it.

- Step 2: Generate an App Password
Once 2-Step Verification is on, you can create the password.

Return to the Google Account security page: myaccount.google.com/security

Click on App passwords. You may need to sign in again.

Under "Select app," choose Mail.

Under "Select device," choose Other (Custom name).

Enter a descriptive name for your application (e.g., "My App SMTP") and click Generate.

Google will display a 16-digit password. Copy this password immediately and save it somewhere secure. This is the password you will use in your application.

Gmail SMTP Settings
Use the following details in your application or script.

SMTP Server: smtp.gmail.com

Port: 587 (for TLS/STARTTLS)

Username: Your full Gmail address (e.g., example@gmail.com)

Password: The 16-digit App Password you just generated.

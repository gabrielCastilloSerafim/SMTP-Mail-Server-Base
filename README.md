# Mail-Server

## Request example:

**POST** `http://localhost:3001/v1/sendMail`

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

# To send HTML you must send it as text and scaping the "" with \ like the following example:
"<!DOCTYPE html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Hello World</title><style>body{background-color:#121212;color:#e0e0e0;display:flex;justify-content:center;align-items:center;height:100vh;margin:0;font-family:sans-serif;}</style></head><body><h1>Hello, World!</h1></body></html>"

# How to add support for a nem Email:

To use Gmail for SMTP, you need your email address and a special App Password, not your regular Google account password.

You can only generate an App Password after enabling 2-Step Verification on your Google Account.

Step 1: Enable 2-Step Verification
If you don't already have it enabled, you must turn on 2-Step Verification.

Go to the Google Account security page: myaccount.google.com/security

Click on 2-Step Verification (under "How you sign in to Google").

Follow the on-screen steps to enable it.

Step 2: Generate an App Password
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

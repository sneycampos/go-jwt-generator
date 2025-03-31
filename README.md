## JWT Generator and Validator

A simple Go application that demonstrates how to generate and validate JSON Web Tokens (JWTs) using the `github.com/golang-jwt/jwt/v5` package.


The program will:
1. Generate a JWT token with standard registered claims
2. Display the generated token
3. Validate the token
4. Display each claim value from the token

## Example Output

```
Generated JWT: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
Validating JWT...
Token is valid
Claims:
Issuer: my-app
Subject: user-id
Audience: [audience]
Issued At: 2023-01-01 12:00:00 +0000 UTC
Expires At: 2023-01-02 12:00:00 +0000 UTC
Not Before: 2023-01-01 11:00:00 +0000 UTC
ID: unique-jwt-id
```

## Security Note

This code uses a hardcoded secret key for demonstration purposes. In production, use a strong, securely stored secret key or asymmetric keys.

## License

MIT
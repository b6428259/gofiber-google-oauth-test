# gofiber-hex-google-oauth (Hexagonal Skeleton)

Minimal repo skeleton for **Go + Fiber** using **Hexagonal (Ports & Adapters)** to do **Google OAuth 2.0 / SSO**.
It only returns `name` and `email` as JSON after login.

## Structure

```
gofiber-hex-google-oauth/
├─ go.mod
├─ .env.example
├─ .gitignore
├─ README.md
├─ cmd/
│  └─ server/
│     └─ main.go
└─ internal/
   ├─ domain/
   │  └─ user.go
   ├─ ports/
   │  └─ oauth.go
   ├─ app/
   │  └─ auth_service.go
   └─ adapters/
      ├─ googleoauth/
      │  └─ client.go
      └─ http/
         └─ server.go
```

## Quickstart

1. **Clone & enter**
   ```bash
   git clone <this-repo> && cd gofiber-hex-google-oauth
   ```

2. **Create `.env` from example and fill values**
   ```bash
   cp .env.example .env
   # Edit GOOGLE_CLIENT_ID / GOOGLE_CLIENT_SECRET / BASE_URL
   ```

3. **Enable OAuth consent screen & add Redirect URI** in Google Cloud Console:
   - Authorized redirect URI: `http://localhost:3000/auth/google/callback`

4. **Install deps & run**
   ```bash
   go mod tidy
   go run ./cmd/server
   ```

5. **Login**
   - Open: `http://localhost:3000/auth/google/login`
   - After consent, callback returns:
     ```json
     { "name": "Your Google Name", "email": "you@gmail.com" }
     ```

> This skeleton stores `state` in a cookie (demo-level). For production, harden session/state handling and CSRF protections.

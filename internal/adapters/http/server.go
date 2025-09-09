package httpadapter

import (
    "crypto/rand"
    "encoding/base64"
    "time"

    "github.com/gofiber/fiber/v2"

    "gofiber-hex-google-oauth/internal/app"
)

type Server struct {
    app *app.AuthService
}

func New(app *app.AuthService) *Server {
    return &Server{app: app}
}

func (s *Server) Router() *fiber.App {
    r := fiber.New()

    r.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("OK: /auth/google/login → sign in with Google")
    })

    r.Get("/auth/google/login", func(c *fiber.Ctx) error {
        state := randomState()
        c.Cookie(&fiber.Cookie{
            Name:     "oauth_state",
            Value:    state,
            HTTPOnly: true,
            SameSite: "Lax",
            Expires:  time.Now().Add(10 * time.Minute),
        })
        url := s.app.OAuthLoginURL(state)
        return c.Redirect(url, fiber.StatusTemporaryRedirect)
    })

    r.Get("/auth/google/callback", func(c *fiber.Ctx) error {
        got := c.Query("state")
        want := c.Cookies("oauth_state")
        if got == "" || got != want {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid state"})
        }
        code := c.Query("code")
        if code == "" {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing code"})
        }

        user, err := s.app.LoginWithGoogle(c.Context(), code)
        if err != nil {
            return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
        }
        return c.JSON(fiber.Map{
            "name":  user.Name,
            "email": user.Email,
        })
    })

    return r
}

func randomState() string {
    b := make([]byte, 24)
    _, _ = rand.Read(b)
    return base64.RawURLEncoding.EncodeToString(b)
}

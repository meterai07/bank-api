package handlers

import (
	"bank/src/repository"
	"bank/src/utils"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	repo *repository.NasabahRepository
}

func NewAuthHandler(repo *repository.NasabahRepository) *AuthHandler {
	return &AuthHandler{repo: repo}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req struct {
		NIK      string `json:"nik"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, 400, "Format request tidak valid")
	}

	// Cari nasabah by NIK
	nasabah, err := h.repo.FindByNIK(req.NIK)
	if err != nil {
		return utils.Error(c, 401, "NIK atau password salah")
	}

	// Verifikasi password
	if err := bcrypt.CompareHashAndPassword([]byte(nasabah.Password), []byte(req.Password)); err != nil {
		return utils.Error(c, 401, "NIK atau password salah")
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"no_rekening": nasabah.NoRekening,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return utils.Error(c, 500, "Gagal membuat token")
	}

	return utils.Success(c, 200, fiber.Map{
		"token": tokenString,
	})
}

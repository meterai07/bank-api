package handlers

import (
	"bank/models"
	"bank/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type NasabahHandler struct {
	repo *repository.NasabahRepository
}

func NewNasabahHandler(repo *repository.NasabahRepository) *NasabahHandler {
	return &NasabahHandler{repo: repo}
}

func (h *NasabahHandler) Daftar(c *fiber.Ctx) error {
	var req struct {
		Nama string `json:"nama"`
		NIK  string `json:"nik"`
		NoHP string `json:"no_hp"`
	}

	if err := c.BodyParser(&req); err != nil {
		log.Error().Err(err).Msg("Invalid request format")
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	nasabah := models.Nasabah{
		Nama: req.Nama,
		NIK:  req.NIK,
		NoHP: req.NoHP,
	}

	if err := h.repo.Create(&nasabah); err != nil {
		if err == repository.ErrDuplicateData {
			return c.Status(400).JSON(fiber.Map{"error": "NIK atau No HP sudah terdaftar"})
		}
		log.Error().Err(err).Msg("Failed to create nasabah")
		return c.Status(500).JSON(fiber.Map{"error": "server error"})
	}

	log.Info().Str("no_rekening", nasabah.NoRekening).Msg("New registration")
	return c.Status(201).JSON(fiber.Map{"no_rekening": nasabah.NoRekening})
}

func (h *NasabahHandler) Tabung(c *fiber.Ctx) error {
	var req struct {
		NoRekening string `json:"no_rekening"`
		Nominal    int64  `json:"nominal"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"remark": "format request tidak valid"})
	}

	if req.Nominal <= 0 {
		return c.Status(400).JSON(fiber.Map{"remark": "nominal harus positif"})
	}

	saldo, err := h.repo.Tabung(req.NoRekening, req.Nominal)
	if err != nil {
		if err == repository.ErrAccountNotFound {
			return c.Status(400).JSON(fiber.Map{"remark": "nomor rekening tidak dikenali"})
		}
		return c.Status(500).JSON(fiber.Map{"remark": "kesalahan server"})
	}

	log.Info().Str("no_rekening", req.NoRekening).Int64("saldo", saldo).Msg("Saldo diperbarui")
	return c.Status(200).JSON(fiber.Map{"saldo": saldo})
}

func (h *NasabahHandler) Tarik(c *fiber.Ctx) error {
	var req struct {
		NoRekening string `json:"no_rekening"`
		Nominal    int64  `json:"nominal"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"remark": "format request tidak valid"})
	}

	if req.Nominal <= 0 {
		return c.Status(400).JSON(fiber.Map{"remark": "nominal harus positif"})
	}

	saldo, err := h.repo.Tarik(req.NoRekening, req.Nominal)
	if err != nil {
		switch err {
		case repository.ErrAccountNotFound:
			return c.Status(400).JSON(fiber.Map{"remark": "nomor rekening tidak dikenali"})
		case repository.ErrInsufficientFund:
			return c.Status(400).JSON(fiber.Map{"remark": "saldo tidak cukup"})
		default:
			return c.Status(500).JSON(fiber.Map{"remark": "kesalahan server"})
		}
	}

	log.Info().Str("no_rekening", req.NoRekening).Int64("saldo", saldo).Msg("Saldo diperbarui")
	return c.Status(200).JSON(fiber.Map{"saldo": saldo})
}

func (h *NasabahHandler) Saldo(c *fiber.Ctx) error {
	noRekening := c.Params("no_rekening")
	saldo, err := h.repo.GetSaldo(noRekening)
	if err != nil {
		if err == repository.ErrAccountNotFound {
			return c.Status(400).JSON(fiber.Map{"remark": "nomor rekening tidak dikenali"})
		}
		return c.Status(500).JSON(fiber.Map{"remark": "kesalahan server"})
	}
	return c.Status(200).JSON(fiber.Map{"saldo": saldo})
}

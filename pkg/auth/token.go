package auth

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/MamangRust/pointofsale-graphql-grpc/pkg/logger"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

var ErrTokenExpired = errors.New("token expired")

//go:generate mockgen -source=token.go -destination=mocks/token.go
type TokenManager interface {
	GenerateToken(userId int, audience string) (string, error)
	ValidateToken(accessToken string) (int, error)
}

type Manager struct {
	secretKey []byte
	logger    logger.LoggerInterface
}

func NewManager(secretKey string, logger logger.LoggerInterface) (*Manager, error) {
	if secretKey == "" {
		return nil, errors.New("empty secret key")
	}
	return &Manager{secretKey: []byte(secretKey), logger: logger}, nil
}

func (m *Manager) GenerateToken(userId int, audience string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(12 * time.Hour)

	m.logger.Debug("Generating access token",
		zap.Int("user_id", userId),
		zap.String("audience", audience),
		zap.Time("issued_at", nowTime),
		zap.Time("expires_at", expireTime),
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expireTime),
		Subject:   strconv.Itoa(userId),
		Audience:  []string{audience},
	})

	signed, err := token.SignedString([]byte(m.secretKey))
	if err != nil {
		m.logger.Error("Failed to sign token",
			zap.Int("user_id", userId),
			zap.Error(err),
		)
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	m.logger.Debug("Token successfully generated",
		zap.Int("user_id", userId),
	)

	return signed, nil
}

func (m *Manager) ValidateToken(accessToken string) (int, error) {
	m.logger.Debug("Validating access token")

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			m.logger.Error("Unexpected signing method",
				zap.Any("header_alg", token.Header["alg"]),
			)
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.secretKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			m.logger.Debug("Access token expired")
			return 0, ErrTokenExpired
		}

		m.logger.Error("Failed to parse token", zap.Error(err))
		return 0, fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		m.logger.Error("Failed to extract claims from token")
		return 0, fmt.Errorf("error get user claims from token")
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		m.logger.Error("Invalid subject claim type",
			zap.Any("sub", claims["sub"]),
		)
		return 0, fmt.Errorf("invalid subject claim type")
	}

	userId, err := strconv.Atoi(sub)
	if err != nil {
		m.logger.Error("Invalid user ID format in subject claim",
			zap.String("sub", sub),
			zap.Error(err),
		)
		return 0, fmt.Errorf("invalid user ID format: %w", err)
	}

	m.logger.Debug("Token validated successfully",
		zap.Int("user_id", userId),
	)

	return userId, nil
}

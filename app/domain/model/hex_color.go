package model

import (
	"errors"

	strings "github.com/umaidashi/go-cli-cobra/app/utils"
)

type HexColor string

func NewColor(color string) (HexColor, error) {
	if ok := strings.IsValidHexColor(color); !ok {
		return "", errors.New("color is invalid.")
	}

	return HexColor(color), nil
}

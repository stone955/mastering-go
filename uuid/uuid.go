package uuid

import (
	"github.com/google/uuid"
)

// GenGoogleUUID 生成google的uuid
func GenGoogleUUID() string {
	ID := uuid.New()
	return ID.String()
}

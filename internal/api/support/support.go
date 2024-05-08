package support

import (
	"MsLetoChat/internal/support/saveconvert"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetUserIDFromContext(ctx *gin.Context) (int64, error) {
	userID, isExist := ctx.Get("user_id")

	if !isExist {
		return 0, fmt.Errorf("invalid user_id in header")
	}

	ownerID, err := saveconvert.SafeConvertToInt64(userID)

	if err != nil {
		return 0, err
	}

	return ownerID, nil
}

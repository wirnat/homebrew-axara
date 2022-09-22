package branch_gorm

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

func getTx(ctx context.Context) (*gorm.DB, error) {
	db, ok := ctx.Value("tx").(*gorm.DB)
	if !ok {
		return nil, fmt.Errorf("without transaction")
	}

	return db, nil
}

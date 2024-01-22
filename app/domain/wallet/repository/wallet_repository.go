package repository

import (
	"github.com/cockroachdb/errors"
	coreModel "go-boilerplate/app/core/model"
	"go-boilerplate/app/domain/wallet/dto"
	"go-boilerplate/app/domain/wallet/model"
	"gorm.io/gorm"
)

type WalletRepository interface {
	GetWalletList(param dto.GetWalletListParam, data *[]model.GetWalletListModel, totalCount *int64) error
}

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return walletRepository{
		db: db,
	}
}

func (r walletRepository) GetWalletList(param dto.GetWalletListParam, data *[]model.GetWalletListModel, totalCount *int64) error {
	db := r.db.Model(&coreModel.WalletProduct{})

	db = db.Select("tb_wallet_product.*, w.*, b.*")

	db = db.Joins("INNER JOIN tb_wallet w ON tb_wallet_product.wallet_id = w.uid AND w.account_id = ?", param.AccountId)
	db = db.Joins("INNER JOIN tb_bill b ON tb_wallet_product.bill_id = b.uid")

	db = db.Count(totalCount).Offset(param.Offset).Limit(param.Limit).Order("tb_wallet_product." + param.OrderBy).Find(&data)
	if db.RowsAffected == 0 || len(*data) == 0 || *totalCount == 0 {
		return errors.New("wallet list data not found")
	}

	return errors.Wrap(db.Error, "failed to find all wallet list")
}

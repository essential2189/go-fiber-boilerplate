package service

import (
	"github.com/cockroachdb/errors"
	"go-boilerplate/app/core"
	"go-boilerplate/app/core/util"
	"go-boilerplate/app/domain/wallet/dto"
	"go-boilerplate/app/domain/wallet/model"
	"go-boilerplate/app/domain/wallet/repository"
)

type WalletService interface {
	GetWalletList(param dto.GetWalletListParam) (*dto.PaginationResponse, error)
}

type walletService struct {
	core core.Modules
	repo repository.WalletRepository
}

func NewWalletService(core core.Modules, repo repository.WalletRepository) WalletService {
	return walletService{
		core: core,
		repo: repo,
	}
}

func (s walletService) GetWalletList(param dto.GetWalletListParam) (*dto.PaginationResponse, error) {
	var data []model.GetWalletListModel
	var totalCount int64

	err := s.repo.GetWalletList(param, &data, &totalCount)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	currentPage, totalPages := util.CalculatePagination(int64(param.Limit), int64(param.Offset), totalCount)
	if currentPage > totalPages {
		return nil, errors.New("invalid page number")
	}

	var wallet []dto.GetWalletListRes
	for _, v := range data {
		wallet = append(wallet, dto.GetWalletListRes{
			WalletType:    v.WalletType,
			Priority:      v.Priority,
			Renew:         v.Renew,
			StartTs:       v.StartTs,
			EndTs:         v.EndTs,
			DownloadCount: v.DownloadCount,
			WalletStatus:  v.WalletStatus,
			CurrencyType:  v.CurrencyType,
			BillType:      v.BillType,
			BillStatus:    v.BillStatus,
			DeviceAgent:   v.DeviceAgent,
			TotalAmount:   v.TotalAmount,
		})
	}

	result := dto.PaginationResponse{CurrentPage: currentPage, TotalPage: totalPages, TotalCount: totalCount, ListCount: int64(len(wallet)), Data: wallet}

	return &result, nil
}

package usecase

import (
	"time"

	domain "github.com/railgun-0402/DI-Golang/app/domain/device"
	"github.com/railgun-0402/DI-Golang/app/utils"
)

type DeviceUsecase struct {
	Repo domain.DeviceRepository
}

func (uc *DeviceUsecase) Register(userID, token, platform string) error {
	d := domain.Device{
		ID: utils.GenerateUUID(),
		UserID: userID,
		Token: token,
		Platform: platform,
		CreatedAt: time.Now(),
	}
	return uc.Repo.Save(d)
}

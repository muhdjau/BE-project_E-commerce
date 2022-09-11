package usecase

import (
	"project/e-commerce/features/user"
)

type userUsecase struct {
	userData user.DataInterface
}

func New(data user.DataInterface) user.UsecaseInterface {
	return &userUsecase{
		userData: data,
	}
}

func (usecase *userUsecase) PostData(data user.Core) (int, error)

func (service *userUsecase) GetByToken(token int) (user.Core, error) {

	dataId, err := service.userData.SelectByToken(token)
	if err != nil {
		return user.Core{}, err
	}

	return dataId, nil

}

package usecase

import "github.com/vitor9/codepix/tree/master/codepix/domain/model"

// A diferenca do usecase para o repositorio, eh que o repositorio soh persiste, enquanto
// o usecase segue todo o fluxo abaixo para validar.
type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (pix *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error){
	account, err := pix.PixKeyRepository.FindAccount(accountId)

	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)
	if err != nil {
		return nil, err
	}

	pix.PixKeyRepository.RegisterKey(pixKey)
	if pixKey.ID == "" {
		return nil, err
	}

	return pixKey, nil
}

func (pix *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := pix.PixKeyRepository.FindKeyByKind(key, kind)

	if err != nil {
		return nil, err
	}

	return pixKey, nil
}
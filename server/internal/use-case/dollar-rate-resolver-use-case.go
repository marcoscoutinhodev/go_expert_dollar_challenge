package use_case

import (
	"encoding/json"

	"github.com/marcoscoutinhodev/go_expert_dollar_challenge/external/adapter"
	"github.com/marcoscoutinhodev/go_expert_dollar_challenge/external/repository"
	"github.com/marcoscoutinhodev/go_expert_dollar_challenge/internal/entity"
)

func DollarRateResolverUseCase() (*entity.DollarRateResponseEntity, error) {
	dollarRate := adapter.DollarRateAdapter()

	var dollarRateEntity entity.DollarRateEntity

	err := json.Unmarshal(dollarRate, &dollarRateEntity)

	if err != nil {
		return nil, err
	}

	repository.AddDollarRateRepository(&dollarRateEntity)

	dollarRateResponseEntity := entity.DollarRateResponseEntity{
		Dollar: dollarRateEntity.Usdbrl.Bid,
	}

	return &dollarRateResponseEntity, nil
}

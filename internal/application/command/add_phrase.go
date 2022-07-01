package command

import (
	"context"
)

type (
	AddPhrase struct {
		Phrase string `json:"phrase"`
	}

	AddPhraseHandler struct {
		PhrasePersistor interface {
			Persist(context.Context, string) error
		}
	}
)

func (a *AddPhraseHandler) Do(ctx context.Context, phrase string) error {
	return nil
}

//////
// Shared utils.
//////

package shared

import (
	"encoding/json"

	"github.com/thalesfsp/configurer/util"
	"github.com/thalesfsp/customerror"
)

// Marshal with custom error.
func Marshal(v any) ([]byte, error) {
	data, err := json.Marshal(&v)
	if err != nil {
		return nil, customerror.NewFailedToError("to marshal",
			customerror.WithError(err),
		)
	}

	return data, nil
}

// Process process the `default` -> `env` -> `id` struct's fields.
func Process(v any) error {
	if err := util.SetDefault(v); err != nil {
		return err
	}

	if err := util.SetEnv(v); err != nil {
		return err
	}

	if err := util.SetID(v); err != nil {
		return err
	}

	return nil
}

package validation

import (
	"context"
	"fmt"
)

type Validatable interface {
	Validate(ctx context.Context) map[string]string
}

func ExtractErrors(errs []error) []string {
	errsStr := make([]string, 0, len(errs))
	for _, err := range errs {
		if err != nil {
			errsStr = append(errsStr, err.Error())
		}
	}
	if len(errsStr) == 0 {
		return nil
	}
	return errsStr
}

type Rule func(key string, value interface{}) error

func ValidateLength(minL, maxL int) Rule {
	return func(key string, value interface{}) error {
		str, ok := value.(string)
		if !ok {
			return fmt.Errorf("%s is not a string", key)
		}
		if len(str) < minL {
			return fmt.Errorf("%s must be more than or equal to %d characters", key, minL)
		}
		if len(str) > maxL {
			return fmt.Errorf("%s must be less than or equal to %d characters", key, maxL)
		}
		return nil
	}
}

func ValidatePresence(key string, value interface{}) error {
	if value == "" {
		return fmt.Errorf("%s can't be blank", key)
	}
	return nil
}

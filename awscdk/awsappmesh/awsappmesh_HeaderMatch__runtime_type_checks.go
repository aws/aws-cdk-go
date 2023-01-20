//go:build !no_runtime_type_checking

package awsappmesh

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (h *jsiiProxy_HeaderMatch) validateBindParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateHeaderMatch_ValueDoesNotEndWithParameters(headerName *string, suffix *string) error {
	if headerName == nil {
		return fmt.Errorf("parameter headerName is required, but nil was provided")
	}

	if suffix == nil {
		return fmt.Errorf("parameter suffix is required, but nil was provided")
	}

	return nil
}

func validateHeaderMatch_ValueDoesNotMatchRegexParameters(headerName *string, regex *string) error {
	if headerName == nil {
		return fmt.Errorf("parameter headerName is required, but nil was provided")
	}

	if regex == nil {
		return fmt.Errorf("parameter regex is required, but nil was provided")
	}

	return nil
}

func validateHeaderMatch_ValueDoesNotStartWithParameters(headerName *string, prefix *string) error {
	if headerName == nil {
		return fmt.Errorf("parameter headerName is required, but nil was provided")
	}

	if prefix == nil {
		return fmt.Errorf("parameter prefix is required, but nil was provided")
	}

	return nil
}

func validateHeaderMatch_ValueEndsWithParameters(headerName *string, suffix *string) error {
	if headerName == nil {
		return fmt.Errorf("parameter headerName is required, but nil was provided")
	}

	if suffix == nil {
		return fmt.Errorf("parameter suffix is required, but nil was provided")
	}

	return nil
}

func validateHeaderMatch_ValueIsParameters(headerName *string, headerValue *string) error {
	if headerName == nil {
		return fmt.Errorf("parameter headerName is required, but nil was provided")
	}

	if headerValue == nil {
		return fmt.Errorf("parameter headerValue is required, but nil was provided")
	}

	return nil
}

func validateHeaderMatch_ValueIsNotParameters(headerName *string, headerValue *string) error {
	if headerName == nil {
		return fmt.Errorf("parameter headerName is required, but nil was provided")
	}

	if headerValue == nil {
		return fmt.Errorf("parameter headerValue is required, but nil was provided")
	}

	return nil
}

func validateHeaderMatch_ValueMatchesRegexParameters(headerName *string, regex *string) error {
	if headerName == nil {
		return fmt.Errorf("parameter headerName is required, but nil was provided")
	}

	if regex == nil {
		return fmt.Errorf("parameter regex is required, but nil was provided")
	}

	return nil
}

func validateHeaderMatch_ValuesIsInRangeParameters(headerName *string, start *float64, end *float64) error {
	if headerName == nil {
		return fmt.Errorf("parameter headerName is required, but nil was provided")
	}

	if start == nil {
		return fmt.Errorf("parameter start is required, but nil was provided")
	}

	if end == nil {
		return fmt.Errorf("parameter end is required, but nil was provided")
	}

	return nil
}

func validateHeaderMatch_ValuesIsNotInRangeParameters(headerName *string, start *float64, end *float64) error {
	if headerName == nil {
		return fmt.Errorf("parameter headerName is required, but nil was provided")
	}

	if start == nil {
		return fmt.Errorf("parameter start is required, but nil was provided")
	}

	if end == nil {
		return fmt.Errorf("parameter end is required, but nil was provided")
	}

	return nil
}

func validateHeaderMatch_ValueStartsWithParameters(headerName *string, prefix *string) error {
	if headerName == nil {
		return fmt.Errorf("parameter headerName is required, but nil was provided")
	}

	if prefix == nil {
		return fmt.Errorf("parameter prefix is required, but nil was provided")
	}

	return nil
}


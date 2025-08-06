//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HeaderMatch) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func validateHeaderMatch_ValueDoesNotEndWithParameters(headerName *string, suffix *string) error {
	return nil
}

func validateHeaderMatch_ValueDoesNotMatchRegexParameters(headerName *string, regex *string) error {
	return nil
}

func validateHeaderMatch_ValueDoesNotStartWithParameters(headerName *string, prefix *string) error {
	return nil
}

func validateHeaderMatch_ValueEndsWithParameters(headerName *string, suffix *string) error {
	return nil
}

func validateHeaderMatch_ValueIsParameters(headerName *string, headerValue *string) error {
	return nil
}

func validateHeaderMatch_ValueIsNotParameters(headerName *string, headerValue *string) error {
	return nil
}

func validateHeaderMatch_ValueMatchesRegexParameters(headerName *string, regex *string) error {
	return nil
}

func validateHeaderMatch_ValuesIsInRangeParameters(headerName *string, start *float64, end *float64) error {
	return nil
}

func validateHeaderMatch_ValuesIsNotInRangeParameters(headerName *string, start *float64, end *float64) error {
	return nil
}

func validateHeaderMatch_ValueStartsWithParameters(headerName *string, prefix *string) error {
	return nil
}


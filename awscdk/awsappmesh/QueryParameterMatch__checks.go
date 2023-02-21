//go:build !no_runtime_type_checking

package awsappmesh

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (q *jsiiProxy_QueryParameterMatch) validateBindParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateQueryParameterMatch_ValueIsParameters(queryParameterName *string, queryParameterValue *string) error {
	if queryParameterName == nil {
		return fmt.Errorf("parameter queryParameterName is required, but nil was provided")
	}

	if queryParameterValue == nil {
		return fmt.Errorf("parameter queryParameterValue is required, but nil was provided")
	}

	return nil
}


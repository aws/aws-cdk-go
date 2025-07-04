//go:build !no_runtime_type_checking

package assertions

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func (t *jsiiProxy_Tags) validateHasValuesParameters(tags interface{}) error {
	if tags == nil {
		return fmt.Errorf("parameter tags is required, but nil was provided")
	}

	return nil
}

func validateTags_FromStackParameters(stack awscdk.Stack) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	return nil
}


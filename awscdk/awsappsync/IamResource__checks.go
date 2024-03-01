//go:build !no_runtime_type_checking

package awsappsync

import (
	"fmt"
)

func (i *jsiiProxy_IamResource) validateResourceArnsParameters(api GraphqlApiBase) error {
	if api == nil {
		return fmt.Errorf("parameter api is required, but nil was provided")
	}

	return nil
}

func validateIamResource_OfTypeParameters(type_ *string) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	return nil
}


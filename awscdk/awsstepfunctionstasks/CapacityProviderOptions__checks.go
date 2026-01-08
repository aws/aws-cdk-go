//go:build !no_runtime_type_checking

package awsstepfunctionstasks

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

func validateCapacityProviderOptions_CustomParameters(capacityProviderStrategy *[]*awsecs.CapacityProviderStrategy) error {
	if capacityProviderStrategy == nil {
		return fmt.Errorf("parameter capacityProviderStrategy is required, but nil was provided")
	}
	for idx_172aa8, v := range *capacityProviderStrategy {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter capacityProviderStrategy[%#v]", idx_172aa8) }); err != nil {
			return err
		}
	}

	return nil
}


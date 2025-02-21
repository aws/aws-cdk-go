//go:build !no_runtime_type_checking

package awsecs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateNewPortMapParameters(networkmode NetworkMode, pm *PortMapping) error {
	if networkmode == "" {
		return fmt.Errorf("parameter networkmode is required, but nil was provided")
	}

	if pm == nil {
		return fmt.Errorf("parameter pm is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(pm, func() string { return "parameter pm" }); err != nil {
		return err
	}

	return nil
}


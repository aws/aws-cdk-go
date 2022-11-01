//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_ILayerVersion) validateAddPermissionParameters(id *string, permission *LayerVersionPermission) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if permission == nil {
		return fmt.Errorf("parameter permission is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(permission, func() string { return "parameter permission" }); err != nil {
		return err
	}

	return nil
}


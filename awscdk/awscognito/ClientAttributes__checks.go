//go:build !no_runtime_type_checking

package awscognito

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (c *jsiiProxy_ClientAttributes) validateWithStandardAttributesParameters(attributes *StandardAttributesMask) error {
	if attributes == nil {
		return fmt.Errorf("parameter attributes is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(attributes, func() string { return "parameter attributes" }); err != nil {
		return err
	}

	return nil
}


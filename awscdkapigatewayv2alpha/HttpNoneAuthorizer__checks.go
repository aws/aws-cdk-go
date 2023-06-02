//go:build !no_runtime_type_checking

package awscdkapigatewayv2alpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (h *jsiiProxy_HttpNoneAuthorizer) validateBindParameters(_options *HttpRouteAuthorizerBindOptions) error {
	if _options == nil {
		return fmt.Errorf("parameter _options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}


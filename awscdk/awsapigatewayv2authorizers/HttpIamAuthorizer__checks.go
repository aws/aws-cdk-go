//go:build !no_runtime_type_checking

package awsapigatewayv2authorizers

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
)

func (h *jsiiProxy_HttpIamAuthorizer) validateBindParameters(_options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) error {
	if _options == nil {
		return fmt.Errorf("parameter _options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}


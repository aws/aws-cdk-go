//go:build !no_runtime_type_checking

package awseventstargets

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsevents"
)

func (a *jsiiProxy_ApiGatewayV2) validateBindParameters(rule interfacesawsevents.IRuleRef) error {
	if rule == nil {
		return fmt.Errorf("parameter rule is required, but nil was provided")
	}

	return nil
}

func validateNewApiGatewayV2Parameters(httpApi awsapigatewayv2.IHttpApi, props *ApiGatewayProps) error {
	if httpApi == nil {
		return fmt.Errorf("parameter httpApi is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


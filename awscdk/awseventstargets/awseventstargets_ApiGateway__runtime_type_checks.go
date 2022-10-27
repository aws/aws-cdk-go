//go:build !no_runtime_type_checking

package awseventstargets

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
)

func (a *jsiiProxy_ApiGateway) validateBindParameters(rule awsevents.IRule) error {
	if rule == nil {
		return fmt.Errorf("parameter rule is required, but nil was provided")
	}

	return nil
}

func validateNewApiGatewayParameters(restApi awsapigateway.RestApi, props *ApiGatewayProps) error {
	if restApi == nil {
		return fmt.Errorf("parameter restApi is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


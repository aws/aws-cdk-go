//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsapigatewayv2integrations

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/awsservicediscovery"
)

func (h *jsiiProxy_HttpServiceDiscoveryIntegration) validateBindParameters(_arg *awsapigatewayv2.HttpRouteIntegrationBindOptions) error {
	if _arg == nil {
		return fmt.Errorf("parameter _arg is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_arg, func() string { return "parameter _arg" }); err != nil {
		return err
	}

	return nil
}

func (h *jsiiProxy_HttpServiceDiscoveryIntegration) validateCompleteBindParameters(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions) error {
	if _options == nil {
		return fmt.Errorf("parameter _options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) validateSetConnectionTypeParameters(val awsapigatewayv2.HttpConnectionType) error {
	if val == "" {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) validateSetHttpMethodParameters(val awsapigatewayv2.HttpMethod) error {
	if val == "" {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) validateSetIntegrationTypeParameters(val awsapigatewayv2.HttpIntegrationType) error {
	if val == "" {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) validateSetPayloadFormatVersionParameters(val awsapigatewayv2.PayloadFormatVersion) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewHttpServiceDiscoveryIntegrationParameters(id *string, service awsservicediscovery.IService, props *HttpServiceDiscoveryIntegrationProps) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if service == nil {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


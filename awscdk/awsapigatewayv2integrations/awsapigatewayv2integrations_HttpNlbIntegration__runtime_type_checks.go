//go:build !no_runtime_type_checking

package awsapigatewayv2integrations

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
)

func (h *jsiiProxy_HttpNlbIntegration) validateBindParameters(options *awsapigatewayv2.HttpRouteIntegrationBindOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (h *jsiiProxy_HttpNlbIntegration) validateCompleteBindParameters(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions) error {
	if _options == nil {
		return fmt.Errorf("parameter _options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_HttpNlbIntegration) validateSetConnectionTypeParameters(val awsapigatewayv2.HttpConnectionType) error {
	if val == "" {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_HttpNlbIntegration) validateSetHttpMethodParameters(val awsapigatewayv2.HttpMethod) error {
	if val == "" {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_HttpNlbIntegration) validateSetIntegrationTypeParameters(val awsapigatewayv2.HttpIntegrationType) error {
	if val == "" {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_HttpNlbIntegration) validateSetPayloadFormatVersionParameters(val awsapigatewayv2.PayloadFormatVersion) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewHttpNlbIntegrationParameters(id *string, listener awselasticloadbalancingv2.INetworkListener, props *HttpNlbIntegrationProps) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if listener == nil {
		return fmt.Errorf("parameter listener is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


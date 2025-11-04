//go:build !no_runtime_type_checking

package awscdkpipesenrichmentsalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
)

func (a *jsiiProxy_ApiGatewayEnrichment) validateBindParameters(pipe awscdkpipesalpha.IPipe) error {
	if pipe == nil {
		return fmt.Errorf("parameter pipe is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApiGatewayEnrichment) validateGrantInvokeParameters(grantee awsiam.IRole) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateNewApiGatewayEnrichmentParameters(restApi awsapigateway.IRestApi, props *ApiGatewayEnrichmentProps) error {
	if restApi == nil {
		return fmt.Errorf("parameter restApi is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


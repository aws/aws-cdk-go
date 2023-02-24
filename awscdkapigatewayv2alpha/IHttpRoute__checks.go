//go:build !no_runtime_type_checking

// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

func (i *jsiiProxy_IHttpRoute) validateGrantInvokeParameters(grantee awsiam.IGrantable, options *GrantInvokeOptions) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


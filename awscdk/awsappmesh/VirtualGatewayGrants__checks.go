//go:build !no_runtime_type_checking

package awsappmesh

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappmesh"
)

func (v *jsiiProxy_VirtualGatewayGrants) validateActionsParameters(grantee awsiam.IGrantable, actions *[]*string, options *awscdk.PermissionsOptions) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	if actions == nil {
		return fmt.Errorf("parameter actions is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (v *jsiiProxy_VirtualGatewayGrants) validateStreamAggregatedResourcesParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateVirtualGatewayGrants_FromVirtualGatewayParameters(resource interfacesawsappmesh.IVirtualGatewayRef) error {
	if resource == nil {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	return nil
}


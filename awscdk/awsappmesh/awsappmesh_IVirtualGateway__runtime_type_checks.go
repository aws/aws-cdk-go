//go:build !no_runtime_type_checking

package awsappmesh

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

func (i *jsiiProxy_IVirtualGateway) validateAddGatewayRouteParameters(id *string, route *GatewayRouteBaseProps) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if route == nil {
		return fmt.Errorf("parameter route is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(route, func() string { return "parameter route" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IVirtualGateway) validateGrantStreamAggregatedResourcesParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}


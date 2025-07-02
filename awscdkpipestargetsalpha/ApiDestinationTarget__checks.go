//go:build !no_runtime_type_checking

package awscdkpipestargetsalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
)

func (a *jsiiProxy_ApiDestinationTarget) validateBindParameters(pipe awscdkpipesalpha.IPipe) error {
	if pipe == nil {
		return fmt.Errorf("parameter pipe is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApiDestinationTarget) validateGrantPushParameters(grantee awsiam.IRole) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateNewApiDestinationTargetParameters(destination awsevents.IApiDestination, parameters *ApiDestinationTargetParameters) error {
	if destination == nil {
		return fmt.Errorf("parameter destination is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(parameters, func() string { return "parameter parameters" }); err != nil {
		return err
	}

	return nil
}


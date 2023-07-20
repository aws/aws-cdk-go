//go:build !no_runtime_type_checking

package awsstepfunctions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

func (d *jsiiProxy_DefinitionBody) validateBindParameters(scope constructs.Construct, sfnPrincipal awsiam.IPrincipal, sfnProps *StateMachineProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if sfnPrincipal == nil {
		return fmt.Errorf("parameter sfnPrincipal is required, but nil was provided")
	}

	if sfnProps == nil {
		return fmt.Errorf("parameter sfnProps is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(sfnProps, func() string { return "parameter sfnProps" }); err != nil {
		return err
	}

	return nil
}

func validateDefinitionBody_FromChainableParameters(chainable IChainable) error {
	if chainable == nil {
		return fmt.Errorf("parameter chainable is required, but nil was provided")
	}

	return nil
}

func validateDefinitionBody_FromFileParameters(path *string, options *awss3assets.AssetOptions) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateDefinitionBody_FromStringParameters(definition *string) error {
	if definition == nil {
		return fmt.Errorf("parameter definition is required, but nil was provided")
	}

	return nil
}


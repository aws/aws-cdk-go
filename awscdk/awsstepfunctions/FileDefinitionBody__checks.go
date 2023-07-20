//go:build !no_runtime_type_checking

package awsstepfunctions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

func (f *jsiiProxy_FileDefinitionBody) validateBindParameters(scope constructs.Construct, _sfnPrincipal awsiam.IPrincipal, _sfnProps *StateMachineProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if _sfnPrincipal == nil {
		return fmt.Errorf("parameter _sfnPrincipal is required, but nil was provided")
	}

	if _sfnProps == nil {
		return fmt.Errorf("parameter _sfnProps is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_sfnProps, func() string { return "parameter _sfnProps" }); err != nil {
		return err
	}

	return nil
}

func validateFileDefinitionBody_FromChainableParameters(chainable IChainable) error {
	if chainable == nil {
		return fmt.Errorf("parameter chainable is required, but nil was provided")
	}

	return nil
}

func validateFileDefinitionBody_FromFileParameters(path *string, options *awss3assets.AssetOptions) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateFileDefinitionBody_FromStringParameters(definition *string) error {
	if definition == nil {
		return fmt.Errorf("parameter definition is required, but nil was provided")
	}

	return nil
}

func validateNewFileDefinitionBodyParameters(path *string, options *awss3assets.AssetOptions) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


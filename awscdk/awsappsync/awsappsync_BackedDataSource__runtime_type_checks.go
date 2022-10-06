//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsappsync

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/constructs-go/constructs/v3"
)

func (b *jsiiProxy_BackedDataSource) validateCreateFunctionParameters(props *BaseAppsyncFunctionProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BackedDataSource) validateCreateResolverParameters(props *BaseResolverProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BackedDataSource) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BackedDataSource) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func validateBackedDataSource_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BackedDataSource) validateSetApiParameters(val IGraphqlApi) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewBackedDataSourceParameters(scope constructs.Construct, id *string, props *BackedDataSourceProps, extended *ExtendedDataSourceProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	if extended == nil {
		return fmt.Errorf("parameter extended is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(extended, func() string { return "parameter extended" }); err != nil {
		return err
	}

	return nil
}


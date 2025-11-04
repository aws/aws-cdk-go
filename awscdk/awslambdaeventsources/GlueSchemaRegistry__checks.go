//go:build !no_runtime_type_checking

package awslambdaeventsources

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

func (g *jsiiProxy_GlueSchemaRegistry) validateBindParameters(target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	if targetHandler == nil {
		return fmt.Errorf("parameter targetHandler is required, but nil was provided")
	}

	return nil
}

func validateNewGlueSchemaRegistryParameters(props *GlueSchemaRegistryProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


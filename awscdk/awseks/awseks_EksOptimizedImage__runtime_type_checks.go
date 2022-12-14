//go:build !no_runtime_type_checking

package awseks

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (e *jsiiProxy_EksOptimizedImage) validateGetImageParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateNewEksOptimizedImageParameters(props *EksOptimizedImageProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


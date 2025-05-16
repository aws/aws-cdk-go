//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (a *jsiiProxy_AdotLambdaLayerPythonSdkVersion) validateLayerArnParameters(scope constructs.IConstruct, architecture Architecture) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if architecture == nil {
		return fmt.Errorf("parameter architecture is required, but nil was provided")
	}

	return nil
}


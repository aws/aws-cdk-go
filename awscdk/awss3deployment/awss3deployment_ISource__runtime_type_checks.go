//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awss3deployment

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

func (i *jsiiProxy_ISource) validateBindParameters(scope awscdk.Construct, context *DeploymentSourceContext) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(context, func() string { return "parameter context" }); err != nil {
		return err
	}

	return nil
}


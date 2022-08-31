//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awscodebuild

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
)

func (i *jsiiProxy_IArtifacts) validateBindParameters(scope awscdk.Construct, project IProject) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	return nil
}


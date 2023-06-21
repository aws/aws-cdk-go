//go:build !no_runtime_type_checking

package awscodebuild

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (i *jsiiProxy_IFileSystemLocation) validateBindParameters(scope constructs.Construct, project IProject) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	return nil
}


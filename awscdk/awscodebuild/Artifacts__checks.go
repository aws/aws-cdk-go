//go:build !no_runtime_type_checking

package awscodebuild

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (a *jsiiProxy_Artifacts) validateBindParameters(_scope constructs.Construct, _project IProject) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if _project == nil {
		return fmt.Errorf("parameter _project is required, but nil was provided")
	}

	return nil
}

func validateArtifacts_S3Parameters(props *S3ArtifactsProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateNewArtifactsParameters(props *ArtifactsProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


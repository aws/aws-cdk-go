//go:build !no_runtime_type_checking

package previewawscodebuildevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodebuild"
)

func (p *jsiiProxy_ProjectEvents) validateCodeBuildBuildPhaseChangePatternParameters(options *ProjectEvents_CodeBuildBuildPhaseChange_CodeBuildBuildPhaseChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_ProjectEvents) validateCodeBuildBuildStateChangePatternParameters(options *ProjectEvents_CodeBuildBuildStateChange_CodeBuildBuildStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateProjectEvents_FromProjectParameters(projectRef interfacesawscodebuild.IProjectRef) error {
	if projectRef == nil {
		return fmt.Errorf("parameter projectRef is required, but nil was provided")
	}

	return nil
}


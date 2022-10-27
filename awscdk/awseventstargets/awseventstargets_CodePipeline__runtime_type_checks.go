//go:build !no_runtime_type_checking

package awseventstargets

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
)

func (c *jsiiProxy_CodePipeline) validateBindParameters(_rule awsevents.IRule) error {
	if _rule == nil {
		return fmt.Errorf("parameter _rule is required, but nil was provided")
	}

	return nil
}

func validateNewCodePipelineParameters(pipeline awscodepipeline.IPipeline, options *CodePipelineTargetOptions) error {
	if pipeline == nil {
		return fmt.Errorf("parameter pipeline is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


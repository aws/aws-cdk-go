package awsimagebuilderalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The step to run in a specific phase of the image build, which defines the step to execute to customize or test the instance.
//
// Example:
//   step := &ComponentDocumentStep{
//   	Name: jsii.String("configure-app"),
//   	Action: imagebuilder.ComponentAction_CREATE_FILE,
//   	Inputs: imagebuilder.ComponentStepInputs_FromObject(map[string]interface{}{
//   		"path": jsii.String("/etc/myapp/config.json"),
//   		"content": jsii.String("{\"env\": \"production\"}"),
//   	}),
//   }
//
// Experimental.
type ComponentDocumentStep struct {
	// The action to perform in the step.
	// Experimental.
	Action ComponentAction `field:"required" json:"action" yaml:"action"`
	// Contains parameters required by the action to run the step.
	// See: https://docs.aws.amazon.com/imagebuilder/latest/userguide/toe-action-modules.html
	//
	// Experimental.
	Inputs ComponentStepInputs `field:"required" json:"inputs" yaml:"inputs"`
	// The name of the step.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The condition to apply to the step.
	//
	// If the condition is false, then the step is skipped.
	// See: https://docs.aws.amazon.com/imagebuilder/latest/userguide/toe-comparison-operators.html
	//
	// Default: - no condition is applied to the step and it gets executed.
	//
	// Experimental.
	If ComponentStepIfCondition `field:"optional" json:"if" yaml:"if"`
	// A looping construct defining a repeated sequence of instructions.
	// Default: None.
	//
	// Experimental.
	Loop *ComponentDocumentLoop `field:"optional" json:"loop" yaml:"loop"`
	// Specifies what the step should do in case of failure.
	// Default: ComponentOnFailure.ABORT
	//
	// Experimental.
	OnFailure ComponentOnFailure `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The timeout of the step.
	// Default: - 120 minutes.
	//
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}


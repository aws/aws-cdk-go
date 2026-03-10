package previewawscloudformationevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.cloudformation@CloudFormationHookInvocationProgress.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cloudFormationHookInvocationProgress := awscdkmixinspreview.Events.NewCloudFormationHookInvocationProgress()
//
// Experimental.
type CloudFormationHookInvocationProgress interface {
}

// The jsii proxy struct for CloudFormationHookInvocationProgress
type jsiiProxy_CloudFormationHookInvocationProgress struct {
	_ byte // padding
}

// Experimental.
func NewCloudFormationHookInvocationProgress() CloudFormationHookInvocationProgress {
	_init_.Initialize()

	j := jsiiProxy_CloudFormationHookInvocationProgress{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudformation.events.CloudFormationHookInvocationProgress",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCloudFormationHookInvocationProgress_Override(c CloudFormationHookInvocationProgress) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudformation.events.CloudFormationHookInvocationProgress",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for CloudFormation Hook Invocation Progress.
// Experimental.
func CloudFormationHookInvocationProgress_CloudFormationHookInvocationProgressPattern(options *CloudFormationHookInvocationProgress_CloudFormationHookInvocationProgressProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateCloudFormationHookInvocationProgress_CloudFormationHookInvocationProgressPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudformation.events.CloudFormationHookInvocationProgress",
		"cloudFormationHookInvocationProgressPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


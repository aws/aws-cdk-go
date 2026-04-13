package previewawscodebuildevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.codebuild@CodeBuildBuildStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeBuildBuildStateChange := awscdkmixinspreview.Events.NewCodeBuildBuildStateChange()
//
// Experimental.
type CodeBuildBuildStateChange interface {
}

// The jsii proxy struct for CodeBuildBuildStateChange
type jsiiProxy_CodeBuildBuildStateChange struct {
	_ byte // padding
}

// Experimental.
func NewCodeBuildBuildStateChange() CodeBuildBuildStateChange {
	_init_.Initialize()

	j := jsiiProxy_CodeBuildBuildStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codebuild.events.CodeBuildBuildStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCodeBuildBuildStateChange_Override(c CodeBuildBuildStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codebuild.events.CodeBuildBuildStateChange",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for CodeBuild Build State Change.
// Experimental.
func CodeBuildBuildStateChange_EventPattern(options *CodeBuildBuildStateChange_CodeBuildBuildStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateCodeBuildBuildStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codebuild.events.CodeBuildBuildStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


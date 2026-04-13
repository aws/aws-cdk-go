package previewawsecrevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.ecr@ECRPullThroughCacheAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eCRPullThroughCacheAction := awscdkmixinspreview.Events.NewECRPullThroughCacheAction()
//
// Experimental.
type ECRPullThroughCacheAction interface {
}

// The jsii proxy struct for ECRPullThroughCacheAction
type jsiiProxy_ECRPullThroughCacheAction struct {
	_ byte // padding
}

// Experimental.
func NewECRPullThroughCacheAction() ECRPullThroughCacheAction {
	_init_.Initialize()

	j := jsiiProxy_ECRPullThroughCacheAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRPullThroughCacheAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewECRPullThroughCacheAction_Override(e ECRPullThroughCacheAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRPullThroughCacheAction",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for ECR Pull Through Cache Action.
// Experimental.
func ECRPullThroughCacheAction_EventPattern(options *ECRPullThroughCacheAction_ECRPullThroughCacheActionProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateECRPullThroughCacheAction_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRPullThroughCacheAction",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


package previewawsecrevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.ecr@ECRImageAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eCRImageAction := awscdkmixinspreview.Events.NewECRImageAction()
//
// Experimental.
type ECRImageAction interface {
}

// The jsii proxy struct for ECRImageAction
type jsiiProxy_ECRImageAction struct {
	_ byte // padding
}

// Experimental.
func NewECRImageAction() ECRImageAction {
	_init_.Initialize()

	j := jsiiProxy_ECRImageAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRImageAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewECRImageAction_Override(e ECRImageAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRImageAction",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for ECR Image Action.
// Experimental.
func ECRImageAction_EcrImageActionPattern(options *ECRImageAction_ECRImageActionProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateECRImageAction_EcrImageActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRImageAction",
		"ecrImageActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


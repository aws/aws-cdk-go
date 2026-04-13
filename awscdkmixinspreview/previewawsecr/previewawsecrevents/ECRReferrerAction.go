package previewawsecrevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.ecr@ECRReferrerAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eCRReferrerAction := awscdkmixinspreview.Events.NewECRReferrerAction()
//
// Experimental.
type ECRReferrerAction interface {
}

// The jsii proxy struct for ECRReferrerAction
type jsiiProxy_ECRReferrerAction struct {
	_ byte // padding
}

// Experimental.
func NewECRReferrerAction() ECRReferrerAction {
	_init_.Initialize()

	j := jsiiProxy_ECRReferrerAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRReferrerAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewECRReferrerAction_Override(e ECRReferrerAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRReferrerAction",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for ECR Referrer Action.
// Experimental.
func ECRReferrerAction_EventPattern(options *ECRReferrerAction_ECRReferrerActionProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateECRReferrerAction_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRReferrerAction",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


package previewawsecrevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.ecr@ECRReplicationAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eCRReplicationAction := awscdkmixinspreview.Events.NewECRReplicationAction()
//
// Experimental.
type ECRReplicationAction interface {
}

// The jsii proxy struct for ECRReplicationAction
type jsiiProxy_ECRReplicationAction struct {
	_ byte // padding
}

// Experimental.
func NewECRReplicationAction() ECRReplicationAction {
	_init_.Initialize()

	j := jsiiProxy_ECRReplicationAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRReplicationAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewECRReplicationAction_Override(e ECRReplicationAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRReplicationAction",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for ECR Replication Action.
// Experimental.
func ECRReplicationAction_EventPattern(options *ECRReplicationAction_ECRReplicationActionProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateECRReplicationAction_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRReplicationAction",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


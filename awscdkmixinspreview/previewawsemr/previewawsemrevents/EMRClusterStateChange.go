package previewawsemrevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.emr@EMRClusterStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eMRClusterStateChange := awscdkmixinspreview.Events.NewEMRClusterStateChange()
//
// Experimental.
type EMRClusterStateChange interface {
}

// The jsii proxy struct for EMRClusterStateChange
type jsiiProxy_EMRClusterStateChange struct {
	_ byte // padding
}

// Experimental.
func NewEMRClusterStateChange() EMRClusterStateChange {
	_init_.Initialize()

	j := jsiiProxy_EMRClusterStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emr.events.EMRClusterStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEMRClusterStateChange_Override(e EMRClusterStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emr.events.EMRClusterStateChange",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for EMR Cluster State Change.
// Experimental.
func EMRClusterStateChange_EventPattern(options *EMRClusterStateChange_EMRClusterStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateEMRClusterStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_emr.events.EMRClusterStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


package previewawsemrevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.emr@EMRAutoScalingPolicyStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eMRAutoScalingPolicyStateChange := awscdkmixinspreview.Events.NewEMRAutoScalingPolicyStateChange()
//
// Experimental.
type EMRAutoScalingPolicyStateChange interface {
}

// The jsii proxy struct for EMRAutoScalingPolicyStateChange
type jsiiProxy_EMRAutoScalingPolicyStateChange struct {
	_ byte // padding
}

// Experimental.
func NewEMRAutoScalingPolicyStateChange() EMRAutoScalingPolicyStateChange {
	_init_.Initialize()

	j := jsiiProxy_EMRAutoScalingPolicyStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emr.events.EMRAutoScalingPolicyStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEMRAutoScalingPolicyStateChange_Override(e EMRAutoScalingPolicyStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emr.events.EMRAutoScalingPolicyStateChange",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for EMR Auto Scaling Policy State Change.
// Experimental.
func EMRAutoScalingPolicyStateChange_EmrAutoScalingPolicyStateChangePattern(options *EMRAutoScalingPolicyStateChange_EMRAutoScalingPolicyStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateEMRAutoScalingPolicyStateChange_EmrAutoScalingPolicyStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_emr.events.EMRAutoScalingPolicyStateChange",
		"emrAutoScalingPolicyStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


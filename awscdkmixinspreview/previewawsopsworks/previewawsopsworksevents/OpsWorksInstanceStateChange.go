package previewawsopsworksevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.opsworks@OpsWorksInstanceStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   opsWorksInstanceStateChange := awscdkmixinspreview.Events.NewOpsWorksInstanceStateChange()
//
// Experimental.
type OpsWorksInstanceStateChange interface {
}

// The jsii proxy struct for OpsWorksInstanceStateChange
type jsiiProxy_OpsWorksInstanceStateChange struct {
	_ byte // padding
}

// Experimental.
func NewOpsWorksInstanceStateChange() OpsWorksInstanceStateChange {
	_init_.Initialize()

	j := jsiiProxy_OpsWorksInstanceStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opsworks.events.OpsWorksInstanceStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewOpsWorksInstanceStateChange_Override(o OpsWorksInstanceStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opsworks.events.OpsWorksInstanceStateChange",
		nil, // no parameters
		o,
	)
}

// EventBridge event pattern for OpsWorks Instance State Change.
// Experimental.
func OpsWorksInstanceStateChange_EventPattern(options *OpsWorksInstanceStateChange_OpsWorksInstanceStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateOpsWorksInstanceStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_opsworks.events.OpsWorksInstanceStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


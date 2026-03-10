package previewawsopsworksevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.opsworks@OpsWorksCommandStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   opsWorksCommandStateChange := awscdkmixinspreview.Events.NewOpsWorksCommandStateChange()
//
// Experimental.
type OpsWorksCommandStateChange interface {
}

// The jsii proxy struct for OpsWorksCommandStateChange
type jsiiProxy_OpsWorksCommandStateChange struct {
	_ byte // padding
}

// Experimental.
func NewOpsWorksCommandStateChange() OpsWorksCommandStateChange {
	_init_.Initialize()

	j := jsiiProxy_OpsWorksCommandStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opsworks.events.OpsWorksCommandStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewOpsWorksCommandStateChange_Override(o OpsWorksCommandStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opsworks.events.OpsWorksCommandStateChange",
		nil, // no parameters
		o,
	)
}

// EventBridge event pattern for OpsWorks Command State Change.
// Experimental.
func OpsWorksCommandStateChange_OpsWorksCommandStateChangePattern(options *OpsWorksCommandStateChange_OpsWorksCommandStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateOpsWorksCommandStateChange_OpsWorksCommandStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_opsworks.events.OpsWorksCommandStateChange",
		"opsWorksCommandStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


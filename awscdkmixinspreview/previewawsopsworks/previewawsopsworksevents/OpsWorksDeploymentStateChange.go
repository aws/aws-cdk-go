package previewawsopsworksevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.opsworks@OpsWorksDeploymentStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   opsWorksDeploymentStateChange := awscdkmixinspreview.Events.NewOpsWorksDeploymentStateChange()
//
// Experimental.
type OpsWorksDeploymentStateChange interface {
}

// The jsii proxy struct for OpsWorksDeploymentStateChange
type jsiiProxy_OpsWorksDeploymentStateChange struct {
	_ byte // padding
}

// Experimental.
func NewOpsWorksDeploymentStateChange() OpsWorksDeploymentStateChange {
	_init_.Initialize()

	j := jsiiProxy_OpsWorksDeploymentStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opsworks.events.OpsWorksDeploymentStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewOpsWorksDeploymentStateChange_Override(o OpsWorksDeploymentStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opsworks.events.OpsWorksDeploymentStateChange",
		nil, // no parameters
		o,
	)
}

// EventBridge event pattern for OpsWorks Deployment State Change.
// Experimental.
func OpsWorksDeploymentStateChange_EventPattern(options *OpsWorksDeploymentStateChange_OpsWorksDeploymentStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateOpsWorksDeploymentStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_opsworks.events.OpsWorksDeploymentStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


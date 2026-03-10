package previewawsopsworksevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.opsworks@OpsWorksAlert.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   opsWorksAlert := awscdkmixinspreview.Events.NewOpsWorksAlert()
//
// Experimental.
type OpsWorksAlert interface {
}

// The jsii proxy struct for OpsWorksAlert
type jsiiProxy_OpsWorksAlert struct {
	_ byte // padding
}

// Experimental.
func NewOpsWorksAlert() OpsWorksAlert {
	_init_.Initialize()

	j := jsiiProxy_OpsWorksAlert{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opsworks.events.OpsWorksAlert",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewOpsWorksAlert_Override(o OpsWorksAlert) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opsworks.events.OpsWorksAlert",
		nil, // no parameters
		o,
	)
}

// EventBridge event pattern for OpsWorks Alert.
// Experimental.
func OpsWorksAlert_OpsWorksAlertPattern(options *OpsWorksAlert_OpsWorksAlertProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateOpsWorksAlert_OpsWorksAlertPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_opsworks.events.OpsWorksAlert",
		"opsWorksAlertPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


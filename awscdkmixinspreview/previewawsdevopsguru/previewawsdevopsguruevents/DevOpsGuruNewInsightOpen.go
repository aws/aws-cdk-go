package previewawsdevopsguruevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.devopsguru@DevOpsGuruNewInsightOpen.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   devOpsGuruNewInsightOpen := awscdkmixinspreview.Events.NewDevOpsGuruNewInsightOpen()
//
// Experimental.
type DevOpsGuruNewInsightOpen interface {
}

// The jsii proxy struct for DevOpsGuruNewInsightOpen
type jsiiProxy_DevOpsGuruNewInsightOpen struct {
	_ byte // padding
}

// Experimental.
func NewDevOpsGuruNewInsightOpen() DevOpsGuruNewInsightOpen {
	_init_.Initialize()

	j := jsiiProxy_DevOpsGuruNewInsightOpen{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_devopsguru.events.DevOpsGuruNewInsightOpen",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDevOpsGuruNewInsightOpen_Override(d DevOpsGuruNewInsightOpen) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_devopsguru.events.DevOpsGuruNewInsightOpen",
		nil, // no parameters
		d,
	)
}

// EventBridge event pattern for DevOps Guru New Insight Open.
// Experimental.
func DevOpsGuruNewInsightOpen_EventPattern(options *DevOpsGuruNewInsightOpen_DevOpsGuruNewInsightOpenProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateDevOpsGuruNewInsightOpen_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_devopsguru.events.DevOpsGuruNewInsightOpen",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


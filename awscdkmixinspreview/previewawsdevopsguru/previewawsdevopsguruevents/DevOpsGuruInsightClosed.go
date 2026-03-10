package previewawsdevopsguruevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.devopsguru@DevOpsGuruInsightClosed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   devOpsGuruInsightClosed := awscdkmixinspreview.Events.NewDevOpsGuruInsightClosed()
//
// Experimental.
type DevOpsGuruInsightClosed interface {
}

// The jsii proxy struct for DevOpsGuruInsightClosed
type jsiiProxy_DevOpsGuruInsightClosed struct {
	_ byte // padding
}

// Experimental.
func NewDevOpsGuruInsightClosed() DevOpsGuruInsightClosed {
	_init_.Initialize()

	j := jsiiProxy_DevOpsGuruInsightClosed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_devopsguru.events.DevOpsGuruInsightClosed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDevOpsGuruInsightClosed_Override(d DevOpsGuruInsightClosed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_devopsguru.events.DevOpsGuruInsightClosed",
		nil, // no parameters
		d,
	)
}

// EventBridge event pattern for DevOps Guru Insight Closed.
// Experimental.
func DevOpsGuruInsightClosed_DevOpsGuruInsightClosedPattern(options *DevOpsGuruInsightClosed_DevOpsGuruInsightClosedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateDevOpsGuruInsightClosed_DevOpsGuruInsightClosedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_devopsguru.events.DevOpsGuruInsightClosed",
		"devOpsGuruInsightClosedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


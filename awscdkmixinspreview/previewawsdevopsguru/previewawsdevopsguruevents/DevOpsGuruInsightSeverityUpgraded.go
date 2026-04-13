package previewawsdevopsguruevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.devopsguru@DevOpsGuruInsightSeverityUpgraded.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   devOpsGuruInsightSeverityUpgraded := awscdkmixinspreview.Events.NewDevOpsGuruInsightSeverityUpgraded()
//
// Experimental.
type DevOpsGuruInsightSeverityUpgraded interface {
}

// The jsii proxy struct for DevOpsGuruInsightSeverityUpgraded
type jsiiProxy_DevOpsGuruInsightSeverityUpgraded struct {
	_ byte // padding
}

// Experimental.
func NewDevOpsGuruInsightSeverityUpgraded() DevOpsGuruInsightSeverityUpgraded {
	_init_.Initialize()

	j := jsiiProxy_DevOpsGuruInsightSeverityUpgraded{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_devopsguru.events.DevOpsGuruInsightSeverityUpgraded",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDevOpsGuruInsightSeverityUpgraded_Override(d DevOpsGuruInsightSeverityUpgraded) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_devopsguru.events.DevOpsGuruInsightSeverityUpgraded",
		nil, // no parameters
		d,
	)
}

// EventBridge event pattern for DevOps Guru Insight Severity Upgraded.
// Experimental.
func DevOpsGuruInsightSeverityUpgraded_EventPattern(options *DevOpsGuruInsightSeverityUpgraded_DevOpsGuruInsightSeverityUpgradedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateDevOpsGuruInsightSeverityUpgraded_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_devopsguru.events.DevOpsGuruInsightSeverityUpgraded",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


package previewawssecurityhubevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.securityhub@SecurityHubInsightResults.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   securityHubInsightResults := awscdkmixinspreview.Events.NewSecurityHubInsightResults()
//
// Experimental.
type SecurityHubInsightResults interface {
}

// The jsii proxy struct for SecurityHubInsightResults
type jsiiProxy_SecurityHubInsightResults struct {
	_ byte // padding
}

// Experimental.
func NewSecurityHubInsightResults() SecurityHubInsightResults {
	_init_.Initialize()

	j := jsiiProxy_SecurityHubInsightResults{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.events.SecurityHubInsightResults",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSecurityHubInsightResults_Override(s SecurityHubInsightResults) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.events.SecurityHubInsightResults",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for Security Hub Insight Results.
// Experimental.
func SecurityHubInsightResults_EventPattern(options *SecurityHubInsightResults_SecurityHubInsightResultsProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSecurityHubInsightResults_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_securityhub.events.SecurityHubInsightResults",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


package previewawssesevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.ses@AdvisorRecommendationStatusOpen.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   advisorRecommendationStatusOpen := awscdkmixinspreview.Events.NewAdvisorRecommendationStatusOpen()
//
// Experimental.
type AdvisorRecommendationStatusOpen interface {
}

// The jsii proxy struct for AdvisorRecommendationStatusOpen
type jsiiProxy_AdvisorRecommendationStatusOpen struct {
	_ byte // padding
}

// Experimental.
func NewAdvisorRecommendationStatusOpen() AdvisorRecommendationStatusOpen {
	_init_.Initialize()

	j := jsiiProxy_AdvisorRecommendationStatusOpen{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.events.AdvisorRecommendationStatusOpen",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAdvisorRecommendationStatusOpen_Override(a AdvisorRecommendationStatusOpen) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.events.AdvisorRecommendationStatusOpen",
		nil, // no parameters
		a,
	)
}

// EventBridge event pattern for Advisor Recommendation Status Open.
// Experimental.
func AdvisorRecommendationStatusOpen_EventPattern(options *AdvisorRecommendationStatusOpen_AdvisorRecommendationStatusOpenProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateAdvisorRecommendationStatusOpen_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ses.events.AdvisorRecommendationStatusOpen",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


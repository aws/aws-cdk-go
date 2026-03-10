package previewawssesevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.ses@AdvisorRecommendationStatusResolved.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   advisorRecommendationStatusResolved := awscdkmixinspreview.Events.NewAdvisorRecommendationStatusResolved()
//
// Experimental.
type AdvisorRecommendationStatusResolved interface {
}

// The jsii proxy struct for AdvisorRecommendationStatusResolved
type jsiiProxy_AdvisorRecommendationStatusResolved struct {
	_ byte // padding
}

// Experimental.
func NewAdvisorRecommendationStatusResolved() AdvisorRecommendationStatusResolved {
	_init_.Initialize()

	j := jsiiProxy_AdvisorRecommendationStatusResolved{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.events.AdvisorRecommendationStatusResolved",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAdvisorRecommendationStatusResolved_Override(a AdvisorRecommendationStatusResolved) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.events.AdvisorRecommendationStatusResolved",
		nil, // no parameters
		a,
	)
}

// EventBridge event pattern for Advisor Recommendation Status Resolved.
// Experimental.
func AdvisorRecommendationStatusResolved_AdvisorRecommendationStatusResolvedPattern(options *AdvisorRecommendationStatusResolved_AdvisorRecommendationStatusResolvedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateAdvisorRecommendationStatusResolved_AdvisorRecommendationStatusResolvedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ses.events.AdvisorRecommendationStatusResolved",
		"advisorRecommendationStatusResolvedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


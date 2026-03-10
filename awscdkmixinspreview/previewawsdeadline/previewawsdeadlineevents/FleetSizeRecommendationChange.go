package previewawsdeadlineevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.deadline@FleetSizeRecommendationChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fleetSizeRecommendationChange := awscdkmixinspreview.Events.NewFleetSizeRecommendationChange()
//
// Experimental.
type FleetSizeRecommendationChange interface {
}

// The jsii proxy struct for FleetSizeRecommendationChange
type jsiiProxy_FleetSizeRecommendationChange struct {
	_ byte // padding
}

// Experimental.
func NewFleetSizeRecommendationChange() FleetSizeRecommendationChange {
	_init_.Initialize()

	j := jsiiProxy_FleetSizeRecommendationChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_deadline.events.FleetSizeRecommendationChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFleetSizeRecommendationChange_Override(f FleetSizeRecommendationChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_deadline.events.FleetSizeRecommendationChange",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for Fleet Size Recommendation Change.
// Experimental.
func FleetSizeRecommendationChange_FleetSizeRecommendationChangePattern(options *FleetSizeRecommendationChange_FleetSizeRecommendationChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFleetSizeRecommendationChange_FleetSizeRecommendationChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_deadline.events.FleetSizeRecommendationChange",
		"fleetSizeRecommendationChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


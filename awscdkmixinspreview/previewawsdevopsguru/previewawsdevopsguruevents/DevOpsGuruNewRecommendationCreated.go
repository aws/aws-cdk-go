package previewawsdevopsguruevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.devopsguru@DevOpsGuruNewRecommendationCreated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   devOpsGuruNewRecommendationCreated := awscdkmixinspreview.Events.NewDevOpsGuruNewRecommendationCreated()
//
// Experimental.
type DevOpsGuruNewRecommendationCreated interface {
}

// The jsii proxy struct for DevOpsGuruNewRecommendationCreated
type jsiiProxy_DevOpsGuruNewRecommendationCreated struct {
	_ byte // padding
}

// Experimental.
func NewDevOpsGuruNewRecommendationCreated() DevOpsGuruNewRecommendationCreated {
	_init_.Initialize()

	j := jsiiProxy_DevOpsGuruNewRecommendationCreated{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_devopsguru.events.DevOpsGuruNewRecommendationCreated",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDevOpsGuruNewRecommendationCreated_Override(d DevOpsGuruNewRecommendationCreated) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_devopsguru.events.DevOpsGuruNewRecommendationCreated",
		nil, // no parameters
		d,
	)
}

// EventBridge event pattern for DevOps Guru New Recommendation Created.
// Experimental.
func DevOpsGuruNewRecommendationCreated_DevOpsGuruNewRecommendationCreatedPattern(options *DevOpsGuruNewRecommendationCreated_DevOpsGuruNewRecommendationCreatedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateDevOpsGuruNewRecommendationCreated_DevOpsGuruNewRecommendationCreatedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_devopsguru.events.DevOpsGuruNewRecommendationCreated",
		"devOpsGuruNewRecommendationCreatedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


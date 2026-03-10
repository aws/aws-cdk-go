package previewawssesevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.events.AdvisorRecommendationStatusOpen",
		reflect.TypeOf((*AdvisorRecommendationStatusOpen)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AdvisorRecommendationStatusOpen{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.events.AdvisorRecommendationStatusOpen.AdvisorRecommendationStatusOpenProps",
		reflect.TypeOf((*AdvisorRecommendationStatusOpen_AdvisorRecommendationStatusOpenProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.events.AdvisorRecommendationStatusResolved",
		reflect.TypeOf((*AdvisorRecommendationStatusResolved)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AdvisorRecommendationStatusResolved{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.events.AdvisorRecommendationStatusResolved.AdvisorRecommendationStatusResolvedProps",
		reflect.TypeOf((*AdvisorRecommendationStatusResolved_AdvisorRecommendationStatusResolvedProps)(nil)).Elem(),
	)
}

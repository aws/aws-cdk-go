package previewawscodeguruprofilerevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.codeguruprofiler@CodeGuruProfilerRecommendationStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeGuruProfilerRecommendationStateChange := awscdkmixinspreview.Events.NewCodeGuruProfilerRecommendationStateChange()
//
// Experimental.
type CodeGuruProfilerRecommendationStateChange interface {
}

// The jsii proxy struct for CodeGuruProfilerRecommendationStateChange
type jsiiProxy_CodeGuruProfilerRecommendationStateChange struct {
	_ byte // padding
}

// Experimental.
func NewCodeGuruProfilerRecommendationStateChange() CodeGuruProfilerRecommendationStateChange {
	_init_.Initialize()

	j := jsiiProxy_CodeGuruProfilerRecommendationStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codeguruprofiler.events.CodeGuruProfilerRecommendationStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCodeGuruProfilerRecommendationStateChange_Override(c CodeGuruProfilerRecommendationStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codeguruprofiler.events.CodeGuruProfilerRecommendationStateChange",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for CodeGuru Profiler Recommendations State Change.
// Experimental.
func CodeGuruProfilerRecommendationStateChange_CodeGuruProfilerRecommendationStateChangePattern(options *CodeGuruProfilerRecommendationStateChange_CodeGuruProfilerRecommendationStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateCodeGuruProfilerRecommendationStateChange_CodeGuruProfilerRecommendationStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codeguruprofiler.events.CodeGuruProfilerRecommendationStateChange",
		"codeGuruProfilerRecommendationStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


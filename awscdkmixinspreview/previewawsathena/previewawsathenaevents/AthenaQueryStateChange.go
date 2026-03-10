package previewawsathenaevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.athena@AthenaQueryStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   athenaQueryStateChange := awscdkmixinspreview.Events.NewAthenaQueryStateChange()
//
// Experimental.
type AthenaQueryStateChange interface {
}

// The jsii proxy struct for AthenaQueryStateChange
type jsiiProxy_AthenaQueryStateChange struct {
	_ byte // padding
}

// Experimental.
func NewAthenaQueryStateChange() AthenaQueryStateChange {
	_init_.Initialize()

	j := jsiiProxy_AthenaQueryStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_athena.events.AthenaQueryStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAthenaQueryStateChange_Override(a AthenaQueryStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_athena.events.AthenaQueryStateChange",
		nil, // no parameters
		a,
	)
}

// EventBridge event pattern for Athena Query State Change.
// Experimental.
func AthenaQueryStateChange_AthenaQueryStateChangePattern(options *AthenaQueryStateChange_AthenaQueryStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateAthenaQueryStateChange_AthenaQueryStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_athena.events.AthenaQueryStateChange",
		"athenaQueryStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


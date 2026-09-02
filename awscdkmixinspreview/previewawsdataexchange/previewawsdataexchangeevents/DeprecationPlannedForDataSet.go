package previewawsdataexchangeevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.dataexchange@DeprecationPlannedForDataSet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   deprecationPlannedForDataSet := awscdkmixinspreview.Events.NewDeprecationPlannedForDataSet()
//
// Experimental.
type DeprecationPlannedForDataSet interface {
}

// The jsii proxy struct for DeprecationPlannedForDataSet
type jsiiProxy_DeprecationPlannedForDataSet struct {
	_ byte // padding
}

// Experimental.
func NewDeprecationPlannedForDataSet() DeprecationPlannedForDataSet {
	_init_.Initialize()

	j := jsiiProxy_DeprecationPlannedForDataSet{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dataexchange.events.DeprecationPlannedForDataSet",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDeprecationPlannedForDataSet_Override(d DeprecationPlannedForDataSet) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dataexchange.events.DeprecationPlannedForDataSet",
		nil, // no parameters
		d,
	)
}

// EventBridge event pattern for Deprecation Planned for Data Set.
// Experimental.
func DeprecationPlannedForDataSet_EventPattern(options *DeprecationPlannedForDataSet_DeprecationPlannedForDataSetProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateDeprecationPlannedForDataSet_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_dataexchange.events.DeprecationPlannedForDataSet",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


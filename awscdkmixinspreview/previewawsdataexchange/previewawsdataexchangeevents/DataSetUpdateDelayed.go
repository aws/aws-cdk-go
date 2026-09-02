package previewawsdataexchangeevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.dataexchange@DataSetUpdateDelayed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSetUpdateDelayed := awscdkmixinspreview.Events.NewDataSetUpdateDelayed()
//
// Experimental.
type DataSetUpdateDelayed interface {
}

// The jsii proxy struct for DataSetUpdateDelayed
type jsiiProxy_DataSetUpdateDelayed struct {
	_ byte // padding
}

// Experimental.
func NewDataSetUpdateDelayed() DataSetUpdateDelayed {
	_init_.Initialize()

	j := jsiiProxy_DataSetUpdateDelayed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dataexchange.events.DataSetUpdateDelayed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDataSetUpdateDelayed_Override(d DataSetUpdateDelayed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dataexchange.events.DataSetUpdateDelayed",
		nil, // no parameters
		d,
	)
}

// EventBridge event pattern for Data Set Update Delayed.
// Experimental.
func DataSetUpdateDelayed_EventPattern(options *DataSetUpdateDelayed_DataSetUpdateDelayedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateDataSetUpdateDelayed_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_dataexchange.events.DataSetUpdateDelayed",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


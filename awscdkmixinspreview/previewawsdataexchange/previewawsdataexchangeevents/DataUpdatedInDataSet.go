package previewawsdataexchangeevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.dataexchange@DataUpdatedInDataSet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataUpdatedInDataSet := awscdkmixinspreview.Events.NewDataUpdatedInDataSet()
//
// Experimental.
type DataUpdatedInDataSet interface {
}

// The jsii proxy struct for DataUpdatedInDataSet
type jsiiProxy_DataUpdatedInDataSet struct {
	_ byte // padding
}

// Experimental.
func NewDataUpdatedInDataSet() DataUpdatedInDataSet {
	_init_.Initialize()

	j := jsiiProxy_DataUpdatedInDataSet{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dataexchange.events.DataUpdatedInDataSet",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDataUpdatedInDataSet_Override(d DataUpdatedInDataSet) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dataexchange.events.DataUpdatedInDataSet",
		nil, // no parameters
		d,
	)
}

// EventBridge event pattern for Data Updated in Data Set.
// Experimental.
func DataUpdatedInDataSet_EventPattern(options *DataUpdatedInDataSet_DataUpdatedInDataSetProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateDataUpdatedInDataSet_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_dataexchange.events.DataUpdatedInDataSet",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


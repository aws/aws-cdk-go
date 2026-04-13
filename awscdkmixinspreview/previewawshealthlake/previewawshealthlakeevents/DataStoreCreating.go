package previewawshealthlakeevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.healthlake@DataStoreCreating.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataStoreCreating := awscdkmixinspreview.Events.NewDataStoreCreating()
//
// Experimental.
type DataStoreCreating interface {
}

// The jsii proxy struct for DataStoreCreating
type jsiiProxy_DataStoreCreating struct {
	_ byte // padding
}

// Experimental.
func NewDataStoreCreating() DataStoreCreating {
	_init_.Initialize()

	j := jsiiProxy_DataStoreCreating{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthlake.events.DataStoreCreating",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDataStoreCreating_Override(d DataStoreCreating) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthlake.events.DataStoreCreating",
		nil, // no parameters
		d,
	)
}

// EventBridge event pattern for Data Store Creating.
// Experimental.
func DataStoreCreating_EventPattern(options *DataStoreCreating_DataStoreCreatingProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateDataStoreCreating_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthlake.events.DataStoreCreating",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


package previewawshealthlakeevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.healthlake@DataStoreActive.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataStoreActive := awscdkmixinspreview.Events.NewDataStoreActive()
//
// Experimental.
type DataStoreActive interface {
}

// The jsii proxy struct for DataStoreActive
type jsiiProxy_DataStoreActive struct {
	_ byte // padding
}

// Experimental.
func NewDataStoreActive() DataStoreActive {
	_init_.Initialize()

	j := jsiiProxy_DataStoreActive{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthlake.events.DataStoreActive",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDataStoreActive_Override(d DataStoreActive) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthlake.events.DataStoreActive",
		nil, // no parameters
		d,
	)
}

// EventBridge event pattern for Data Store Active.
// Experimental.
func DataStoreActive_EventPattern(options *DataStoreActive_DataStoreActiveProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateDataStoreActive_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthlake.events.DataStoreActive",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


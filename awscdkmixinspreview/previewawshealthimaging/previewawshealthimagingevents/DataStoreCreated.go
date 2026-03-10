package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.healthimaging@DataStoreCreated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataStoreCreated := awscdkmixinspreview.Events.NewDataStoreCreated()
//
// Experimental.
type DataStoreCreated interface {
}

// The jsii proxy struct for DataStoreCreated
type jsiiProxy_DataStoreCreated struct {
	_ byte // padding
}

// Experimental.
func NewDataStoreCreated() DataStoreCreated {
	_init_.Initialize()

	j := jsiiProxy_DataStoreCreated{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DataStoreCreated",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDataStoreCreated_Override(d DataStoreCreated) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DataStoreCreated",
		nil, // no parameters
		d,
	)
}

// EventBridge event pattern for Data Store Created.
// Experimental.
func DataStoreCreated_DataStoreCreatedPattern(options *DataStoreCreated_DataStoreCreatedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateDataStoreCreated_DataStoreCreatedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DataStoreCreated",
		"dataStoreCreatedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


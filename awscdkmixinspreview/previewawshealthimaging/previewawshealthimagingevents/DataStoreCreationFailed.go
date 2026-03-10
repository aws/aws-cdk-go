package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.healthimaging@DataStoreCreationFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataStoreCreationFailed := awscdkmixinspreview.Events.NewDataStoreCreationFailed()
//
// Experimental.
type DataStoreCreationFailed interface {
}

// The jsii proxy struct for DataStoreCreationFailed
type jsiiProxy_DataStoreCreationFailed struct {
	_ byte // padding
}

// Experimental.
func NewDataStoreCreationFailed() DataStoreCreationFailed {
	_init_.Initialize()

	j := jsiiProxy_DataStoreCreationFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DataStoreCreationFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDataStoreCreationFailed_Override(d DataStoreCreationFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DataStoreCreationFailed",
		nil, // no parameters
		d,
	)
}

// EventBridge event pattern for Data Store Creation Failed.
// Experimental.
func DataStoreCreationFailed_DataStoreCreationFailedPattern(options *DataStoreCreationFailed_DataStoreCreationFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateDataStoreCreationFailed_DataStoreCreationFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DataStoreCreationFailed",
		"dataStoreCreationFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


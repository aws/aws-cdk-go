package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.healthimaging@DataStoreDeleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataStoreDeleted := awscdkmixinspreview.Events.NewDataStoreDeleted()
//
// Experimental.
type DataStoreDeleted interface {
}

// The jsii proxy struct for DataStoreDeleted
type jsiiProxy_DataStoreDeleted struct {
	_ byte // padding
}

// Experimental.
func NewDataStoreDeleted() DataStoreDeleted {
	_init_.Initialize()

	j := jsiiProxy_DataStoreDeleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DataStoreDeleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDataStoreDeleted_Override(d DataStoreDeleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DataStoreDeleted",
		nil, // no parameters
		d,
	)
}

// EventBridge event pattern for Data Store Deleted.
// Experimental.
func DataStoreDeleted_DataStoreDeletedPattern(options *DataStoreDeleted_DataStoreDeletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateDataStoreDeleted_DataStoreDeletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DataStoreDeleted",
		"dataStoreDeletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


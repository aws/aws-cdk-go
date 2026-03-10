package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.healthimaging@DataStoreDeleting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataStoreDeleting := awscdkmixinspreview.Events.NewDataStoreDeleting()
//
// Experimental.
type DataStoreDeleting interface {
}

// The jsii proxy struct for DataStoreDeleting
type jsiiProxy_DataStoreDeleting struct {
	_ byte // padding
}

// Experimental.
func NewDataStoreDeleting() DataStoreDeleting {
	_init_.Initialize()

	j := jsiiProxy_DataStoreDeleting{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DataStoreDeleting",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDataStoreDeleting_Override(d DataStoreDeleting) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DataStoreDeleting",
		nil, // no parameters
		d,
	)
}

// EventBridge event pattern for Data Store Deleting.
// Experimental.
func DataStoreDeleting_DataStoreDeletingPattern(options *DataStoreDeleting_DataStoreDeletingProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateDataStoreDeleting_DataStoreDeletingPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DataStoreDeleting",
		"dataStoreDeletingPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


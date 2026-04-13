package previewawsdatasyncevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.datasync@DataSyncLocationStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSyncLocationStateChange := awscdkmixinspreview.Events.NewDataSyncLocationStateChange()
//
// Experimental.
type DataSyncLocationStateChange interface {
}

// The jsii proxy struct for DataSyncLocationStateChange
type jsiiProxy_DataSyncLocationStateChange struct {
	_ byte // padding
}

// Experimental.
func NewDataSyncLocationStateChange() DataSyncLocationStateChange {
	_init_.Initialize()

	j := jsiiProxy_DataSyncLocationStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datasync.events.DataSyncLocationStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDataSyncLocationStateChange_Override(d DataSyncLocationStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datasync.events.DataSyncLocationStateChange",
		nil, // no parameters
		d,
	)
}

// EventBridge event pattern for DataSync Location State Change.
// Experimental.
func DataSyncLocationStateChange_EventPattern(options *DataSyncLocationStateChange_DataSyncLocationStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateDataSyncLocationStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_datasync.events.DataSyncLocationStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


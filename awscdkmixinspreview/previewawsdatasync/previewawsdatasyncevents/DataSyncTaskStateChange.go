package previewawsdatasyncevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.datasync@DataSyncTaskStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSyncTaskStateChange := awscdkmixinspreview.Events.NewDataSyncTaskStateChange()
//
// Experimental.
type DataSyncTaskStateChange interface {
}

// The jsii proxy struct for DataSyncTaskStateChange
type jsiiProxy_DataSyncTaskStateChange struct {
	_ byte // padding
}

// Experimental.
func NewDataSyncTaskStateChange() DataSyncTaskStateChange {
	_init_.Initialize()

	j := jsiiProxy_DataSyncTaskStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datasync.events.DataSyncTaskStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDataSyncTaskStateChange_Override(d DataSyncTaskStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datasync.events.DataSyncTaskStateChange",
		nil, // no parameters
		d,
	)
}

// EventBridge event pattern for DataSync Task State Change.
// Experimental.
func DataSyncTaskStateChange_DataSyncTaskStateChangePattern(options *DataSyncTaskStateChange_DataSyncTaskStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateDataSyncTaskStateChange_DataSyncTaskStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_datasync.events.DataSyncTaskStateChange",
		"dataSyncTaskStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


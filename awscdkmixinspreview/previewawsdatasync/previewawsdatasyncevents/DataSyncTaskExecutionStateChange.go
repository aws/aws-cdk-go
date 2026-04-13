package previewawsdatasyncevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.datasync@DataSyncTaskExecutionStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSyncTaskExecutionStateChange := awscdkmixinspreview.Events.NewDataSyncTaskExecutionStateChange()
//
// Experimental.
type DataSyncTaskExecutionStateChange interface {
}

// The jsii proxy struct for DataSyncTaskExecutionStateChange
type jsiiProxy_DataSyncTaskExecutionStateChange struct {
	_ byte // padding
}

// Experimental.
func NewDataSyncTaskExecutionStateChange() DataSyncTaskExecutionStateChange {
	_init_.Initialize()

	j := jsiiProxy_DataSyncTaskExecutionStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datasync.events.DataSyncTaskExecutionStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDataSyncTaskExecutionStateChange_Override(d DataSyncTaskExecutionStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datasync.events.DataSyncTaskExecutionStateChange",
		nil, // no parameters
		d,
	)
}

// EventBridge event pattern for DataSync Task Execution State Change.
// Experimental.
func DataSyncTaskExecutionStateChange_EventPattern(options *DataSyncTaskExecutionStateChange_DataSyncTaskExecutionStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateDataSyncTaskExecutionStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_datasync.events.DataSyncTaskExecutionStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


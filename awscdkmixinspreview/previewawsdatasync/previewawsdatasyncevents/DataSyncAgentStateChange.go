package previewawsdatasyncevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.datasync@DataSyncAgentStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSyncAgentStateChange := awscdkmixinspreview.Events.NewDataSyncAgentStateChange()
//
// Experimental.
type DataSyncAgentStateChange interface {
}

// The jsii proxy struct for DataSyncAgentStateChange
type jsiiProxy_DataSyncAgentStateChange struct {
	_ byte // padding
}

// Experimental.
func NewDataSyncAgentStateChange() DataSyncAgentStateChange {
	_init_.Initialize()

	j := jsiiProxy_DataSyncAgentStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datasync.events.DataSyncAgentStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDataSyncAgentStateChange_Override(d DataSyncAgentStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datasync.events.DataSyncAgentStateChange",
		nil, // no parameters
		d,
	)
}

// EventBridge event pattern for DataSync Agent State Change.
// Experimental.
func DataSyncAgentStateChange_EventPattern(options *DataSyncAgentStateChange_DataSyncAgentStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateDataSyncAgentStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_datasync.events.DataSyncAgentStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


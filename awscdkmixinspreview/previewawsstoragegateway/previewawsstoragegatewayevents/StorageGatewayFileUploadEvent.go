package previewawsstoragegatewayevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.storagegateway@StorageGatewayFileUploadEvent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   storageGatewayFileUploadEvent := awscdkmixinspreview.Events.NewStorageGatewayFileUploadEvent()
//
// Experimental.
type StorageGatewayFileUploadEvent interface {
}

// The jsii proxy struct for StorageGatewayFileUploadEvent
type jsiiProxy_StorageGatewayFileUploadEvent struct {
	_ byte // padding
}

// Experimental.
func NewStorageGatewayFileUploadEvent() StorageGatewayFileUploadEvent {
	_init_.Initialize()

	j := jsiiProxy_StorageGatewayFileUploadEvent{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_storagegateway.events.StorageGatewayFileUploadEvent",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewStorageGatewayFileUploadEvent_Override(s StorageGatewayFileUploadEvent) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_storagegateway.events.StorageGatewayFileUploadEvent",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for Storage Gateway File Upload Event.
// Experimental.
func StorageGatewayFileUploadEvent_EventPattern(options *StorageGatewayFileUploadEvent_StorageGatewayFileUploadEventProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateStorageGatewayFileUploadEvent_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_storagegateway.events.StorageGatewayFileUploadEvent",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


package previewawsstoragegatewayevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.storagegateway@StorageGatewayObjectUploadEvent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   storageGatewayObjectUploadEvent := awscdkmixinspreview.Events.NewStorageGatewayObjectUploadEvent()
//
// Experimental.
type StorageGatewayObjectUploadEvent interface {
}

// The jsii proxy struct for StorageGatewayObjectUploadEvent
type jsiiProxy_StorageGatewayObjectUploadEvent struct {
	_ byte // padding
}

// Experimental.
func NewStorageGatewayObjectUploadEvent() StorageGatewayObjectUploadEvent {
	_init_.Initialize()

	j := jsiiProxy_StorageGatewayObjectUploadEvent{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_storagegateway.events.StorageGatewayObjectUploadEvent",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewStorageGatewayObjectUploadEvent_Override(s StorageGatewayObjectUploadEvent) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_storagegateway.events.StorageGatewayObjectUploadEvent",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for Storage Gateway Object Upload Event.
// Experimental.
func StorageGatewayObjectUploadEvent_EventPattern(options *StorageGatewayObjectUploadEvent_StorageGatewayObjectUploadEventProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateStorageGatewayObjectUploadEvent_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_storagegateway.events.StorageGatewayObjectUploadEvent",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


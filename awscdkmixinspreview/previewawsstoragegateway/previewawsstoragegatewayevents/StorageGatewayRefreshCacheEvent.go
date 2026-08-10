package previewawsstoragegatewayevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.storagegateway@StorageGatewayRefreshCacheEvent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   storageGatewayRefreshCacheEvent := awscdkmixinspreview.Events.NewStorageGatewayRefreshCacheEvent()
//
// Experimental.
type StorageGatewayRefreshCacheEvent interface {
}

// The jsii proxy struct for StorageGatewayRefreshCacheEvent
type jsiiProxy_StorageGatewayRefreshCacheEvent struct {
	_ byte // padding
}

// Experimental.
func NewStorageGatewayRefreshCacheEvent() StorageGatewayRefreshCacheEvent {
	_init_.Initialize()

	j := jsiiProxy_StorageGatewayRefreshCacheEvent{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_storagegateway.events.StorageGatewayRefreshCacheEvent",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewStorageGatewayRefreshCacheEvent_Override(s StorageGatewayRefreshCacheEvent) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_storagegateway.events.StorageGatewayRefreshCacheEvent",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for Storage Gateway Refresh Cache Event.
// Experimental.
func StorageGatewayRefreshCacheEvent_EventPattern(options *StorageGatewayRefreshCacheEvent_StorageGatewayRefreshCacheEventProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateStorageGatewayRefreshCacheEvent_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_storagegateway.events.StorageGatewayRefreshCacheEvent",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


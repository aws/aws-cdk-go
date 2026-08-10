package previewawsstoragegatewayevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_storagegateway.events.StorageGatewayFileUploadEvent",
		reflect.TypeOf((*StorageGatewayFileUploadEvent)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_StorageGatewayFileUploadEvent{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_storagegateway.events.StorageGatewayFileUploadEvent.StorageGatewayFileUploadEventProps",
		reflect.TypeOf((*StorageGatewayFileUploadEvent_StorageGatewayFileUploadEventProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_storagegateway.events.StorageGatewayObjectUploadEvent",
		reflect.TypeOf((*StorageGatewayObjectUploadEvent)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_StorageGatewayObjectUploadEvent{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_storagegateway.events.StorageGatewayObjectUploadEvent.StorageGatewayObjectUploadEventProps",
		reflect.TypeOf((*StorageGatewayObjectUploadEvent_StorageGatewayObjectUploadEventProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_storagegateway.events.StorageGatewayRefreshCacheEvent",
		reflect.TypeOf((*StorageGatewayRefreshCacheEvent)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_StorageGatewayRefreshCacheEvent{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_storagegateway.events.StorageGatewayRefreshCacheEvent.StorageGatewayRefreshCacheEventProps",
		reflect.TypeOf((*StorageGatewayRefreshCacheEvent_StorageGatewayRefreshCacheEventProps)(nil)).Elem(),
	)
}

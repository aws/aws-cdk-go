package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
)

// A CloudFront Key Value Store.
type IKeyValueStore interface {
	awscdk.IResource
	// The ARN of the Key Value Store.
	KeyValueStoreArn() *string
	// The Unique ID of the Key Value Store.
	KeyValueStoreId() *string
	// The status of the Key Value Store.
	KeyValueStoreStatus() *string
}

// The jsii proxy for IKeyValueStore
type jsiiProxy_IKeyValueStore struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IKeyValueStore) KeyValueStoreArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyValueStoreArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKeyValueStore) KeyValueStoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyValueStoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKeyValueStore) KeyValueStoreStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyValueStoreStatus",
		&returns,
	)
	return returns
}


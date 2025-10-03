package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFront Key Value Store.
type IKeyValueStore interface {
	IKeyValueStoreRef
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
	jsiiProxy_IKeyValueStoreRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IKeyValueStore) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
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

func (j *jsiiProxy_IKeyValueStore) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKeyValueStore) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKeyValueStore) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


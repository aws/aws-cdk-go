package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Interface for Event API.
type IEventApi interface {
	IApi
	// add a new channel namespace.
	//
	// Returns: the channel namespace.
	AddChannelNamespace(id *string, options *ChannelNamespaceOptions) ChannelNamespace
	// Adds an IAM policy statement associated with this Event API to an IAM principal's policy.
	Grant(grantee awsiam.IGrantable, resources AppSyncEventResource, actions ...*string) awsiam.Grant
	// Adds an IAM policy statement for EventConnect access to this EventApi to an IAM principal's policy.
	GrantConnect(grantee awsiam.IGrantable) awsiam.Grant
	// Adds an IAM policy statement for EventPublish access to this EventApi to an IAM principal's policy.
	GrantPublish(grantee awsiam.IGrantable) awsiam.Grant
	// Adds an IAM policy statement to publish and subscribe to this API for an IAM principal's policy.
	GrantPublishAndSubscribe(grantee awsiam.IGrantable) awsiam.Grant
	// Adds an IAM policy statement for EventSubscribe access to this EventApi to an IAM principal's policy.
	GrantSubscribe(grantee awsiam.IGrantable) awsiam.Grant
	// The Authorization Types for this Event Api.
	AuthProviderTypes() *[]AppSyncAuthorizationType
	// The domain name of the Api's HTTP endpoint.
	HttpDns() *string
	// The domain name of the Api's real-time endpoint.
	RealtimeDns() *string
}

// The jsii proxy for IEventApi
type jsiiProxy_IEventApi struct {
	jsiiProxy_IApi
}

func (i *jsiiProxy_IEventApi) AddChannelNamespace(id *string, options *ChannelNamespaceOptions) ChannelNamespace {
	if err := i.validateAddChannelNamespaceParameters(id, options); err != nil {
		panic(err)
	}
	var returns ChannelNamespace

	_jsii_.Invoke(
		i,
		"addChannelNamespace",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEventApi) Grant(grantee awsiam.IGrantable, resources AppSyncEventResource, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee, resources); err != nil {
		panic(err)
	}
	args := []interface{}{grantee, resources}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEventApi) GrantConnect(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantConnectParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantConnect",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEventApi) GrantPublish(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantPublishParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantPublish",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEventApi) GrantPublishAndSubscribe(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantPublishAndSubscribeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantPublishAndSubscribe",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEventApi) GrantSubscribe(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantSubscribeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantSubscribe",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IEventApi) AuthProviderTypes() *[]AppSyncAuthorizationType {
	var returns *[]AppSyncAuthorizationType
	_jsii_.Get(
		j,
		"authProviderTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventApi) HttpDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventApi) RealtimeDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeDns",
		&returns,
	)
	return returns
}


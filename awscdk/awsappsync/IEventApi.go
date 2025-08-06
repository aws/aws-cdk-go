package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchservice"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Interface for Event API.
type IEventApi interface {
	IApi
	// add a new channel namespace.
	//
	// Returns: the channel namespace.
	AddChannelNamespace(id *string, options *ChannelNamespaceOptions) ChannelNamespace
	// Add a new DynamoDB data source to this API.
	AddDynamoDbDataSource(id *string, table awsdynamodb.ITable, options *AppSyncDataSourceOptions) AppSyncDynamoDbDataSource
	// Add an EventBridge data source to this api.
	AddEventBridgeDataSource(id *string, eventBus awsevents.IEventBus, options *AppSyncDataSourceOptions) AppSyncEventBridgeDataSource
	// add a new http data source to this API.
	AddHttpDataSource(id *string, endpoint *string, options *AppSyncHttpDataSourceOptions) AppSyncHttpDataSource
	// add a new Lambda data source to this API.
	AddLambdaDataSource(id *string, lambdaFunction awslambda.IFunction, options *AppSyncDataSourceOptions) AppSyncLambdaDataSource
	// Add a new OpenSearch data source to this API.
	AddOpenSearchDataSource(id *string, domain awsopensearchservice.IDomain, options *AppSyncDataSourceOptions) AppSyncOpenSearchDataSource
	// add a new Rds data source to this API.
	AddRdsDataSource(id *string, serverlessCluster interface{}, secretStore awssecretsmanager.ISecret, databaseName *string, options *AppSyncDataSourceOptions) AppSyncRdsDataSource
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

func (i *jsiiProxy_IEventApi) AddDynamoDbDataSource(id *string, table awsdynamodb.ITable, options *AppSyncDataSourceOptions) AppSyncDynamoDbDataSource {
	if err := i.validateAddDynamoDbDataSourceParameters(id, table, options); err != nil {
		panic(err)
	}
	var returns AppSyncDynamoDbDataSource

	_jsii_.Invoke(
		i,
		"addDynamoDbDataSource",
		[]interface{}{id, table, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEventApi) AddEventBridgeDataSource(id *string, eventBus awsevents.IEventBus, options *AppSyncDataSourceOptions) AppSyncEventBridgeDataSource {
	if err := i.validateAddEventBridgeDataSourceParameters(id, eventBus, options); err != nil {
		panic(err)
	}
	var returns AppSyncEventBridgeDataSource

	_jsii_.Invoke(
		i,
		"addEventBridgeDataSource",
		[]interface{}{id, eventBus, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEventApi) AddHttpDataSource(id *string, endpoint *string, options *AppSyncHttpDataSourceOptions) AppSyncHttpDataSource {
	if err := i.validateAddHttpDataSourceParameters(id, endpoint, options); err != nil {
		panic(err)
	}
	var returns AppSyncHttpDataSource

	_jsii_.Invoke(
		i,
		"addHttpDataSource",
		[]interface{}{id, endpoint, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEventApi) AddLambdaDataSource(id *string, lambdaFunction awslambda.IFunction, options *AppSyncDataSourceOptions) AppSyncLambdaDataSource {
	if err := i.validateAddLambdaDataSourceParameters(id, lambdaFunction, options); err != nil {
		panic(err)
	}
	var returns AppSyncLambdaDataSource

	_jsii_.Invoke(
		i,
		"addLambdaDataSource",
		[]interface{}{id, lambdaFunction, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEventApi) AddOpenSearchDataSource(id *string, domain awsopensearchservice.IDomain, options *AppSyncDataSourceOptions) AppSyncOpenSearchDataSource {
	if err := i.validateAddOpenSearchDataSourceParameters(id, domain, options); err != nil {
		panic(err)
	}
	var returns AppSyncOpenSearchDataSource

	_jsii_.Invoke(
		i,
		"addOpenSearchDataSource",
		[]interface{}{id, domain, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEventApi) AddRdsDataSource(id *string, serverlessCluster interface{}, secretStore awssecretsmanager.ISecret, databaseName *string, options *AppSyncDataSourceOptions) AppSyncRdsDataSource {
	if err := i.validateAddRdsDataSourceParameters(id, serverlessCluster, secretStore, options); err != nil {
		panic(err)
	}
	var returns AppSyncRdsDataSource

	_jsii_.Invoke(
		i,
		"addRdsDataSource",
		[]interface{}{id, serverlessCluster, secretStore, databaseName, options},
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


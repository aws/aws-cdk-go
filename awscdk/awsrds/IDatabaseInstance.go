package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
)

// A database instance.
type IDatabaseInstance interface {
	awsec2.IConnectable
	awscdk.IResource
	awssecretsmanager.ISecretAttachmentTarget
	// Add a new db proxy to this instance.
	AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy
	// Grant the given identity connection access to the database.
	GrantConnect(grantee awsiam.IGrantable, dbUser *string) awsiam.Grant
	// Return the given named metric for this DBInstance.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The percentage of CPU utilization.
	//
	// Average over 5 minutes.
	MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of database connections in use.
	//
	// Average over 5 minutes.
	MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of available random access memory.
	//
	// Average over 5 minutes.
	MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of available storage space.
	//
	// Average over 5 minutes.
	MetricFreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The average number of disk write I/O operations per second.
	//
	// Average over 5 minutes.
	MetricReadIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The average number of disk read I/O operations per second.
	//
	// Average over 5 minutes.
	MetricWriteIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Defines a CloudWatch event rule which triggers for instance events.
	//
	// Use
	// `rule.addEventPattern(pattern)` to specify a filter.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// The instance endpoint address.
	DbInstanceEndpointAddress() *string
	// The instance endpoint port.
	DbInstanceEndpointPort() *string
	// The engine of this database Instance.
	//
	// May be not known for imported Instances if it wasn't provided explicitly,
	// or for read replicas.
	Engine() IInstanceEngine
	// The instance arn.
	InstanceArn() *string
	// The instance endpoint.
	InstanceEndpoint() Endpoint
	// The instance identifier.
	InstanceIdentifier() *string
	// The AWS Region-unique, immutable identifier for the DB instance.
	//
	// This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB instance is accessed.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#aws-resource-rds-dbinstance-return-values
	//
	InstanceResourceId() *string
}

// The jsii proxy for IDatabaseInstance
type jsiiProxy_IDatabaseInstance struct {
	internal.Type__awsec2IConnectable
	internal.Type__awscdkIResource
	internal.Type__awssecretsmanagerISecretAttachmentTarget
}

func (i *jsiiProxy_IDatabaseInstance) AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy {
	if err := i.validateAddProxyParameters(id, options); err != nil {
		panic(err)
	}
	var returns DatabaseProxy

	_jsii_.Invoke(
		i,
		"addProxy",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseInstance) GrantConnect(grantee awsiam.IGrantable, dbUser *string) awsiam.Grant {
	if err := i.validateGrantConnectParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantConnect",
		[]interface{}{grantee, dbUser},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseInstance) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseInstance) MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricCPUUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricCPUUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseInstance) MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricDatabaseConnectionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricDatabaseConnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseInstance) MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricFreeableMemoryParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricFreeableMemory",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseInstance) MetricFreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricFreeStorageSpaceParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricFreeStorageSpace",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseInstance) MetricReadIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricReadIOPSParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricReadIOPS",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseInstance) MetricWriteIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricWriteIOPSParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricWriteIOPS",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseInstance) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnEventParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IDatabaseInstance) AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps {
	var returns *awssecretsmanager.SecretAttachmentTargetProps

	_jsii_.Invoke(
		i,
		"asSecretAttachmentTarget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IDatabaseInstance) DbInstanceEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) DbInstanceEndpointPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) Engine() IInstanceEngine {
	var returns IInstanceEngine
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) InstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) InstanceEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"instanceEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) InstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) InstanceResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


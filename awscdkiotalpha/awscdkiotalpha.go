// The CDK Construct Library for AWS::IoT
package awscdkiotalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkiotalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot"
	"github.com/aws/aws-cdk-go/awscdkiotalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Properties for an topic rule action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import iot_alpha "github.com/aws/aws-cdk-go/awscdkiotalpha"
//
//   actionConfig := &actionConfig{
//   	configuration: &actionProperty{
//   		cloudwatchAlarm: &cloudwatchAlarmActionProperty{
//   			alarmName: jsii.String("alarmName"),
//   			roleArn: jsii.String("roleArn"),
//   			stateReason: jsii.String("stateReason"),
//   			stateValue: jsii.String("stateValue"),
//   		},
//   		cloudwatchLogs: &cloudwatchLogsActionProperty{
//   			logGroupName: jsii.String("logGroupName"),
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		cloudwatchMetric: &cloudwatchMetricActionProperty{
//   			metricName: jsii.String("metricName"),
//   			metricNamespace: jsii.String("metricNamespace"),
//   			metricUnit: jsii.String("metricUnit"),
//   			metricValue: jsii.String("metricValue"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			metricTimestamp: jsii.String("metricTimestamp"),
//   		},
//   		dynamoDb: &dynamoDBActionProperty{
//   			hashKeyField: jsii.String("hashKeyField"),
//   			hashKeyValue: jsii.String("hashKeyValue"),
//   			roleArn: jsii.String("roleArn"),
//   			tableName: jsii.String("tableName"),
//
//   			// the properties below are optional
//   			hashKeyType: jsii.String("hashKeyType"),
//   			payloadField: jsii.String("payloadField"),
//   			rangeKeyField: jsii.String("rangeKeyField"),
//   			rangeKeyType: jsii.String("rangeKeyType"),
//   			rangeKeyValue: jsii.String("rangeKeyValue"),
//   		},
//   		dynamoDBv2: &dynamoDBv2ActionProperty{
//   			putItem: &putItemInputProperty{
//   				tableName: jsii.String("tableName"),
//   			},
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		elasticsearch: &elasticsearchActionProperty{
//   			endpoint: jsii.String("endpoint"),
//   			id: jsii.String("id"),
//   			index: jsii.String("index"),
//   			roleArn: jsii.String("roleArn"),
//   			type: jsii.String("type"),
//   		},
//   		firehose: &firehoseActionProperty{
//   			deliveryStreamName: jsii.String("deliveryStreamName"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			batchMode: jsii.Boolean(false),
//   			separator: jsii.String("separator"),
//   		},
//   		http: &httpActionProperty{
//   			url: jsii.String("url"),
//
//   			// the properties below are optional
//   			auth: &httpAuthorizationProperty{
//   				sigv4: &sigV4AuthorizationProperty{
//   					roleArn: jsii.String("roleArn"),
//   					serviceName: jsii.String("serviceName"),
//   					signingRegion: jsii.String("signingRegion"),
//   				},
//   			},
//   			confirmationUrl: jsii.String("confirmationUrl"),
//   			headers: []interface{}{
//   				&httpActionHeaderProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		iotAnalytics: &iotAnalyticsActionProperty{
//   			channelName: jsii.String("channelName"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			batchMode: jsii.Boolean(false),
//   		},
//   		iotEvents: &iotEventsActionProperty{
//   			inputName: jsii.String("inputName"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			batchMode: jsii.Boolean(false),
//   			messageId: jsii.String("messageId"),
//   		},
//   		iotSiteWise: &iotSiteWiseActionProperty{
//   			putAssetPropertyValueEntries: []interface{}{
//   				&putAssetPropertyValueEntryProperty{
//   					propertyValues: []interface{}{
//   						&assetPropertyValueProperty{
//   							timestamp: &assetPropertyTimestampProperty{
//   								timeInSeconds: jsii.String("timeInSeconds"),
//
//   								// the properties below are optional
//   								offsetInNanos: jsii.String("offsetInNanos"),
//   							},
//   							value: &assetPropertyVariantProperty{
//   								booleanValue: jsii.String("booleanValue"),
//   								doubleValue: jsii.String("doubleValue"),
//   								integerValue: jsii.String("integerValue"),
//   								stringValue: jsii.String("stringValue"),
//   							},
//
//   							// the properties below are optional
//   							quality: jsii.String("quality"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					assetId: jsii.String("assetId"),
//   					entryId: jsii.String("entryId"),
//   					propertyAlias: jsii.String("propertyAlias"),
//   					propertyId: jsii.String("propertyId"),
//   				},
//   			},
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		kafka: &kafkaActionProperty{
//   			clientProperties: map[string]*string{
//   				"clientPropertiesKey": jsii.String("clientProperties"),
//   			},
//   			destinationArn: jsii.String("destinationArn"),
//   			topic: jsii.String("topic"),
//
//   			// the properties below are optional
//   			key: jsii.String("key"),
//   			partition: jsii.String("partition"),
//   		},
//   		kinesis: &kinesisActionProperty{
//   			roleArn: jsii.String("roleArn"),
//   			streamName: jsii.String("streamName"),
//
//   			// the properties below are optional
//   			partitionKey: jsii.String("partitionKey"),
//   		},
//   		lambda: &lambdaActionProperty{
//   			functionArn: jsii.String("functionArn"),
//   		},
//   		openSearch: &openSearchActionProperty{
//   			endpoint: jsii.String("endpoint"),
//   			id: jsii.String("id"),
//   			index: jsii.String("index"),
//   			roleArn: jsii.String("roleArn"),
//   			type: jsii.String("type"),
//   		},
//   		republish: &republishActionProperty{
//   			roleArn: jsii.String("roleArn"),
//   			topic: jsii.String("topic"),
//
//   			// the properties below are optional
//   			qos: jsii.Number(123),
//   		},
//   		s3: &s3ActionProperty{
//   			bucketName: jsii.String("bucketName"),
//   			key: jsii.String("key"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			cannedAcl: jsii.String("cannedAcl"),
//   		},
//   		sns: &snsActionProperty{
//   			roleArn: jsii.String("roleArn"),
//   			targetArn: jsii.String("targetArn"),
//
//   			// the properties below are optional
//   			messageFormat: jsii.String("messageFormat"),
//   		},
//   		sqs: &sqsActionProperty{
//   			queueUrl: jsii.String("queueUrl"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			useBase64: jsii.Boolean(false),
//   		},
//   		stepFunctions: &stepFunctionsActionProperty{
//   			roleArn: jsii.String("roleArn"),
//   			stateMachineName: jsii.String("stateMachineName"),
//
//   			// the properties below are optional
//   			executionNamePrefix: jsii.String("executionNamePrefix"),
//   		},
//   		timestream: &timestreamActionProperty{
//   			databaseName: jsii.String("databaseName"),
//   			dimensions: []interface{}{
//   				&timestreamDimensionProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			roleArn: jsii.String("roleArn"),
//   			tableName: jsii.String("tableName"),
//
//   			// the properties below are optional
//   			batchMode: jsii.Boolean(false),
//   			timestamp: &timestreamTimestampProperty{
//   				unit: jsii.String("unit"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type ActionConfig struct {
	// The configuration for this action.
	// Experimental.
	Configuration *awsiot.CfnTopicRule_ActionProperty `field:"required" json:"configuration" yaml:"configuration"`
}

// An abstract action for TopicRule.
// Experimental.
type IAction interface {
	// Returns the topic rule action specification.
	// Experimental.
	Bind(topicRule ITopicRule) *ActionConfig
}

// The jsii proxy for IAction
type jsiiProxy_IAction struct {
	_ byte // padding
}

func (i *jsiiProxy_IAction) Bind(topicRule ITopicRule) *ActionConfig {
	var returns *ActionConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{topicRule},
		&returns,
	)

	return returns
}

// Represents an AWS IoT Rule.
// Experimental.
type ITopicRule interface {
	awscdk.IResource
	// The value of the topic rule Amazon Resource Name (ARN), such as arn:aws:iot:us-east-2:123456789012:rule/rule_name.
	// Experimental.
	TopicRuleArn() *string
	// The name topic rule.
	// Experimental.
	TopicRuleName() *string
}

// The jsii proxy for ITopicRule
type jsiiProxy_ITopicRule struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ITopicRule) TopicRuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicRuleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITopicRule) TopicRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicRuleName",
		&returns,
	)
	return returns
}

// Defines AWS IoT SQL.
//
// Example:
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//
//   iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewS3PutObjectAction(bucket, &s3PutObjectActionProps{
//   			accessControl: s3.bucketAccessControl_PUBLIC_READ,
//   		}),
//   	},
//   })
//
// Experimental.
type IotSql interface {
	// Returns the IoT SQL configuration.
	// Experimental.
	Bind(scope constructs.Construct) *IotSqlConfig
}

// The jsii proxy struct for IotSql
type jsiiProxy_IotSql struct {
	_ byte // padding
}

// Experimental.
func NewIotSql_Override(i IotSql) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-alpha.IotSql",
		nil, // no parameters
		i,
	)
}

// Uses the original SQL version built on 2015-10-08.
//
// Returns: Instance of IotSql.
// Experimental.
func IotSql_FromStringAsVer20151008(sql *string) IotSql {
	_init_.Initialize()

	var returns IotSql

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iot-alpha.IotSql",
		"fromStringAsVer20151008",
		[]interface{}{sql},
		&returns,
	)

	return returns
}

// Uses the SQL version built on 2016-03-23.
//
// Returns: Instance of IotSql.
// Experimental.
func IotSql_FromStringAsVer20160323(sql *string) IotSql {
	_init_.Initialize()

	var returns IotSql

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iot-alpha.IotSql",
		"fromStringAsVer20160323",
		[]interface{}{sql},
		&returns,
	)

	return returns
}

// Uses the most recent beta SQL version.
//
// If you use this version, it might
// introduce breaking changes to your rules.
//
// Returns: Instance of IotSql.
// Experimental.
func IotSql_FromStringAsVerNewestUnstable(sql *string) IotSql {
	_init_.Initialize()

	var returns IotSql

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iot-alpha.IotSql",
		"fromStringAsVerNewestUnstable",
		[]interface{}{sql},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSql) Bind(scope constructs.Construct) *IotSqlConfig {
	var returns *IotSqlConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// The type returned from the `bind()` method in {@link IotSql}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import iot_alpha "github.com/aws/aws-cdk-go/awscdkiotalpha"
//
//   iotSqlConfig := &iotSqlConfig{
//   	awsIotSqlVersion: jsii.String("awsIotSqlVersion"),
//   	sql: jsii.String("sql"),
//   }
//
// Experimental.
type IotSqlConfig struct {
	// The version of the SQL rules engine to use when evaluating the rule.
	// Experimental.
	AwsIotSqlVersion *string `field:"required" json:"awsIotSqlVersion" yaml:"awsIotSqlVersion"`
	// The SQL statement used to query the topic.
	// Experimental.
	Sql *string `field:"required" json:"sql" yaml:"sql"`
}

// Defines an AWS IoT Rule in this stack.
//
// Example:
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//
//   iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewS3PutObjectAction(bucket, &s3PutObjectActionProps{
//   			accessControl: s3.bucketAccessControl_PUBLIC_READ,
//   		}),
//   	},
//   })
//
// Experimental.
type TopicRule interface {
	awscdk.Resource
	ITopicRule
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Arn of this topic rule.
	// Experimental.
	TopicRuleArn() *string
	// Name of this topic rule.
	// Experimental.
	TopicRuleName() *string
	// Add a action to the topic rule.
	// Experimental.
	AddAction(action IAction)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TopicRule
type jsiiProxy_TopicRule struct {
	internal.Type__awscdkResource
	jsiiProxy_ITopicRule
}

func (j *jsiiProxy_TopicRule) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicRule) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicRule) TopicRuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicRuleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicRule) TopicRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicRuleName",
		&returns,
	)
	return returns
}


// Experimental.
func NewTopicRule(scope constructs.Construct, id *string, props *TopicRuleProps) TopicRule {
	_init_.Initialize()

	j := jsiiProxy_TopicRule{}

	_jsii_.Create(
		"@aws-cdk/aws-iot-alpha.TopicRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewTopicRule_Override(t TopicRule, scope constructs.Construct, id *string, props *TopicRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-alpha.TopicRule",
		[]interface{}{scope, id, props},
		t,
	)
}

// Import an existing AWS IoT Rule provided an ARN.
// Experimental.
func TopicRule_FromTopicRuleArn(scope constructs.Construct, id *string, topicRuleArn *string) ITopicRule {
	_init_.Initialize()

	var returns ITopicRule

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iot-alpha.TopicRule",
		"fromTopicRuleArn",
		[]interface{}{scope, id, topicRuleArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func TopicRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iot-alpha.TopicRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func TopicRule_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iot-alpha.TopicRule",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicRule) AddAction(action IAction) {
	_jsii_.InvokeVoid(
		t,
		"addAction",
		[]interface{}{action},
	)
}

func (t *jsiiProxy_TopicRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		t,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (t *jsiiProxy_TopicRule) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicRule) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicRule) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for defining an AWS IoT Rule.
//
// Example:
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//
//   iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewS3PutObjectAction(bucket, &s3PutObjectActionProps{
//   			accessControl: s3.bucketAccessControl_PUBLIC_READ,
//   		}),
//   	},
//   })
//
// Experimental.
type TopicRuleProps struct {
	// A simplified SQL syntax to filter messages received on an MQTT topic and push the data elsewhere.
	// See: https://docs.aws.amazon.com/iot/latest/developerguide/iot-sql-reference.html
	//
	// Experimental.
	Sql IotSql `field:"required" json:"sql" yaml:"sql"`
	// The actions associated with the topic rule.
	// Experimental.
	Actions *[]IAction `field:"optional" json:"actions" yaml:"actions"`
	// A textual description of the topic rule.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether the rule is enabled.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The action AWS IoT performs when it is unable to perform a rule's action.
	// Experimental.
	ErrorAction IAction `field:"optional" json:"errorAction" yaml:"errorAction"`
	// The name of the topic rule.
	// Experimental.
	TopicRuleName *string `field:"optional" json:"topicRuleName" yaml:"topicRuleName"`
}


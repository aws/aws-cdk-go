package awskafkaconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskafkaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::KafkaConnect::Connector`.
//
// Creates a connector using the specified properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnector := awscdk.Aws_kafkaconnect.NewCfnConnector(this, jsii.String("MyCfnConnector"), &cfnConnectorProps{
//   	capacity: &capacityProperty{
//   		autoScaling: &autoScalingProperty{
//   			maxWorkerCount: jsii.Number(123),
//   			mcuCount: jsii.Number(123),
//   			minWorkerCount: jsii.Number(123),
//   			scaleInPolicy: &scaleInPolicyProperty{
//   				cpuUtilizationPercentage: jsii.Number(123),
//   			},
//   			scaleOutPolicy: &scaleOutPolicyProperty{
//   				cpuUtilizationPercentage: jsii.Number(123),
//   			},
//   		},
//   		provisionedCapacity: &provisionedCapacityProperty{
//   			workerCount: jsii.Number(123),
//
//   			// the properties below are optional
//   			mcuCount: jsii.Number(123),
//   		},
//   	},
//   	connectorConfiguration: map[string]*string{
//   		"connectorConfigurationKey": jsii.String("connectorConfiguration"),
//   	},
//   	connectorName: jsii.String("connectorName"),
//   	kafkaCluster: &kafkaClusterProperty{
//   		apacheKafkaCluster: &apacheKafkaClusterProperty{
//   			bootstrapServers: jsii.String("bootstrapServers"),
//   			vpc: &vpcProperty{
//   				securityGroups: []*string{
//   					jsii.String("securityGroups"),
//   				},
//   				subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//   			},
//   		},
//   	},
//   	kafkaClusterClientAuthentication: &kafkaClusterClientAuthenticationProperty{
//   		authenticationType: jsii.String("authenticationType"),
//   	},
//   	kafkaClusterEncryptionInTransit: &kafkaClusterEncryptionInTransitProperty{
//   		encryptionType: jsii.String("encryptionType"),
//   	},
//   	kafkaConnectVersion: jsii.String("kafkaConnectVersion"),
//   	plugins: []interface{}{
//   		&pluginProperty{
//   			customPlugin: &customPluginProperty{
//   				customPluginArn: jsii.String("customPluginArn"),
//   				revision: jsii.Number(123),
//   			},
//   		},
//   	},
//   	serviceExecutionRoleArn: jsii.String("serviceExecutionRoleArn"),
//
//   	// the properties below are optional
//   	connectorDescription: jsii.String("connectorDescription"),
//   	logDelivery: &logDeliveryProperty{
//   		workerLogDelivery: &workerLogDeliveryProperty{
//   			cloudWatchLogs: &cloudWatchLogsLogDeliveryProperty{
//   				enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				logGroup: jsii.String("logGroup"),
//   			},
//   			firehose: &firehoseLogDeliveryProperty{
//   				enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				deliveryStream: jsii.String("deliveryStream"),
//   			},
//   			s3: &s3LogDeliveryProperty{
//   				enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				bucket: jsii.String("bucket"),
//   				prefix: jsii.String("prefix"),
//   			},
//   		},
//   	},
//   	workerConfiguration: &workerConfigurationProperty{
//   		revision: jsii.Number(123),
//   		workerConfigurationArn: jsii.String("workerConfigurationArn"),
//   	},
//   })
//
type CfnConnector interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the newly created connector.
	AttrConnectorArn() *string
	// The connector's compute capacity settings.
	Capacity() interface{}
	SetCapacity(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The configuration of the connector.
	ConnectorConfiguration() interface{}
	SetConnectorConfiguration(val interface{})
	// The description of the connector.
	ConnectorDescription() *string
	SetConnectorDescription(val *string)
	// The name of the connector.
	ConnectorName() *string
	SetConnectorName(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The details of the Apache Kafka cluster to which the connector is connected.
	KafkaCluster() interface{}
	SetKafkaCluster(val interface{})
	// The type of client authentication used to connect to the Apache Kafka cluster.
	//
	// The value is NONE when no client authentication is used.
	KafkaClusterClientAuthentication() interface{}
	SetKafkaClusterClientAuthentication(val interface{})
	// Details of encryption in transit to the Apache Kafka cluster.
	KafkaClusterEncryptionInTransit() interface{}
	SetKafkaClusterEncryptionInTransit(val interface{})
	// The version of Kafka Connect.
	//
	// It has to be compatible with both the Apache Kafka cluster's version and the plugins.
	KafkaConnectVersion() *string
	SetKafkaConnectVersion(val *string)
	// The settings for delivering connector logs to Amazon CloudWatch Logs.
	LogDelivery() interface{}
	SetLogDelivery(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Specifies which plugin to use for the connector.
	//
	// You must specify a single-element list. Amazon MSK Connect does not currently support specifying multiple plugins.
	Plugins() interface{}
	SetPlugins(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The Amazon Resource Name (ARN) of the IAM role used by the connector to access Amazon Web Services resources.
	ServiceExecutionRoleArn() *string
	SetServiceExecutionRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// The worker configurations that are in use with the connector.
	WorkerConfiguration() interface{}
	SetWorkerConfiguration(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnConnector
type jsiiProxy_CfnConnector struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnConnector) AttrConnectorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrConnectorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) Capacity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) ConnectorConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) ConnectorDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) ConnectorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) KafkaCluster() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kafkaCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) KafkaClusterClientAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kafkaClusterClientAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) KafkaClusterEncryptionInTransit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kafkaClusterEncryptionInTransit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) KafkaConnectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kafkaConnectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) LogDelivery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) Plugins() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"plugins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) ServiceExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnector) WorkerConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workerConfiguration",
		&returns,
	)
	return returns
}


// Create a new `AWS::KafkaConnect::Connector`.
func NewCfnConnector(scope constructs.Construct, id *string, props *CfnConnectorProps) CfnConnector {
	_init_.Initialize()

	if err := validateNewCfnConnectorParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConnector{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::KafkaConnect::Connector`.
func NewCfnConnector_Override(c CfnConnector, scope constructs.Construct, id *string, props *CfnConnectorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnConnector)SetCapacity(val interface{}) {
	if err := j.validateSetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacity",
		val,
	)
}

func (j *jsiiProxy_CfnConnector)SetConnectorConfiguration(val interface{}) {
	if err := j.validateSetConnectorConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectorConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnConnector)SetConnectorDescription(val *string) {
	_jsii_.Set(
		j,
		"connectorDescription",
		val,
	)
}

func (j *jsiiProxy_CfnConnector)SetConnectorName(val *string) {
	if err := j.validateSetConnectorNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectorName",
		val,
	)
}

func (j *jsiiProxy_CfnConnector)SetKafkaCluster(val interface{}) {
	if err := j.validateSetKafkaClusterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kafkaCluster",
		val,
	)
}

func (j *jsiiProxy_CfnConnector)SetKafkaClusterClientAuthentication(val interface{}) {
	if err := j.validateSetKafkaClusterClientAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kafkaClusterClientAuthentication",
		val,
	)
}

func (j *jsiiProxy_CfnConnector)SetKafkaClusterEncryptionInTransit(val interface{}) {
	if err := j.validateSetKafkaClusterEncryptionInTransitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kafkaClusterEncryptionInTransit",
		val,
	)
}

func (j *jsiiProxy_CfnConnector)SetKafkaConnectVersion(val *string) {
	if err := j.validateSetKafkaConnectVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kafkaConnectVersion",
		val,
	)
}

func (j *jsiiProxy_CfnConnector)SetLogDelivery(val interface{}) {
	if err := j.validateSetLogDeliveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logDelivery",
		val,
	)
}

func (j *jsiiProxy_CfnConnector)SetPlugins(val interface{}) {
	if err := j.validateSetPluginsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"plugins",
		val,
	)
}

func (j *jsiiProxy_CfnConnector)SetServiceExecutionRoleArn(val *string) {
	if err := j.validateSetServiceExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceExecutionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnConnector)SetWorkerConfiguration(val interface{}) {
	if err := j.validateSetWorkerConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workerConfiguration",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnConnector_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConnector_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnConnector_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnConnector_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnConnector_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConnector_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConnector_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kafkaconnect.CfnConnector",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConnector) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnConnector) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnConnector) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnConnector) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnConnector) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnConnector) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnConnector) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnConnector) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnector) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnector) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnConnector) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnConnector) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnector) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnector) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnector) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}


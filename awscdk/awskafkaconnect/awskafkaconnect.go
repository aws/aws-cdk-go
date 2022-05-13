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
	AttrConnectorArn() *string
	// `AWS::KafkaConnect::Connector.Capacity`.
	Capacity() interface{}
	SetCapacity(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// `AWS::KafkaConnect::Connector.ConnectorConfiguration`.
	ConnectorConfiguration() interface{}
	SetConnectorConfiguration(val interface{})
	// `AWS::KafkaConnect::Connector.ConnectorDescription`.
	ConnectorDescription() *string
	SetConnectorDescription(val *string)
	// `AWS::KafkaConnect::Connector.ConnectorName`.
	ConnectorName() *string
	SetConnectorName(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// `AWS::KafkaConnect::Connector.KafkaCluster`.
	KafkaCluster() interface{}
	SetKafkaCluster(val interface{})
	// `AWS::KafkaConnect::Connector.KafkaClusterClientAuthentication`.
	KafkaClusterClientAuthentication() interface{}
	SetKafkaClusterClientAuthentication(val interface{})
	// `AWS::KafkaConnect::Connector.KafkaClusterEncryptionInTransit`.
	KafkaClusterEncryptionInTransit() interface{}
	SetKafkaClusterEncryptionInTransit(val interface{})
	// `AWS::KafkaConnect::Connector.KafkaConnectVersion`.
	KafkaConnectVersion() *string
	SetKafkaConnectVersion(val *string)
	// `AWS::KafkaConnect::Connector.LogDelivery`.
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
	// `AWS::KafkaConnect::Connector.Plugins`.
	Plugins() interface{}
	SetPlugins(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// `AWS::KafkaConnect::Connector.ServiceExecutionRoleArn`.
	ServiceExecutionRoleArn() *string
	SetServiceExecutionRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// `AWS::KafkaConnect::Connector.WorkerConfiguration`.
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

func (j *jsiiProxy_CfnConnector) SetCapacity(val interface{}) {
	_jsii_.Set(
		j,
		"capacity",
		val,
	)
}

func (j *jsiiProxy_CfnConnector) SetConnectorConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"connectorConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnConnector) SetConnectorDescription(val *string) {
	_jsii_.Set(
		j,
		"connectorDescription",
		val,
	)
}

func (j *jsiiProxy_CfnConnector) SetConnectorName(val *string) {
	_jsii_.Set(
		j,
		"connectorName",
		val,
	)
}

func (j *jsiiProxy_CfnConnector) SetKafkaCluster(val interface{}) {
	_jsii_.Set(
		j,
		"kafkaCluster",
		val,
	)
}

func (j *jsiiProxy_CfnConnector) SetKafkaClusterClientAuthentication(val interface{}) {
	_jsii_.Set(
		j,
		"kafkaClusterClientAuthentication",
		val,
	)
}

func (j *jsiiProxy_CfnConnector) SetKafkaClusterEncryptionInTransit(val interface{}) {
	_jsii_.Set(
		j,
		"kafkaClusterEncryptionInTransit",
		val,
	)
}

func (j *jsiiProxy_CfnConnector) SetKafkaConnectVersion(val *string) {
	_jsii_.Set(
		j,
		"kafkaConnectVersion",
		val,
	)
}

func (j *jsiiProxy_CfnConnector) SetLogDelivery(val interface{}) {
	_jsii_.Set(
		j,
		"logDelivery",
		val,
	)
}

func (j *jsiiProxy_CfnConnector) SetPlugins(val interface{}) {
	_jsii_.Set(
		j,
		"plugins",
		val,
	)
}

func (j *jsiiProxy_CfnConnector) SetServiceExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"serviceExecutionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnConnector) SetWorkerConfiguration(val interface{}) {
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnConnector_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

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
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnConnector) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnConnector) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnConnector) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnConnector) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnConnector) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnConnector) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnConnector) GetAtt(attributeName *string) awscdk.Reference {
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
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnConnector) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnConnector) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apacheKafkaClusterProperty := &apacheKafkaClusterProperty{
//   	bootstrapServers: jsii.String("bootstrapServers"),
//   	vpc: &vpcProperty{
//   		securityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }
//
type CfnConnector_ApacheKafkaClusterProperty struct {
	// `CfnConnector.ApacheKafkaClusterProperty.BootstrapServers`.
	BootstrapServers *string `field:"required" json:"bootstrapServers" yaml:"bootstrapServers"`
	// `CfnConnector.ApacheKafkaClusterProperty.Vpc`.
	Vpc interface{} `field:"required" json:"vpc" yaml:"vpc"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoScalingProperty := &autoScalingProperty{
//   	maxWorkerCount: jsii.Number(123),
//   	mcuCount: jsii.Number(123),
//   	minWorkerCount: jsii.Number(123),
//   	scaleInPolicy: &scaleInPolicyProperty{
//   		cpuUtilizationPercentage: jsii.Number(123),
//   	},
//   	scaleOutPolicy: &scaleOutPolicyProperty{
//   		cpuUtilizationPercentage: jsii.Number(123),
//   	},
//   }
//
type CfnConnector_AutoScalingProperty struct {
	// `CfnConnector.AutoScalingProperty.MaxWorkerCount`.
	MaxWorkerCount *float64 `field:"required" json:"maxWorkerCount" yaml:"maxWorkerCount"`
	// `CfnConnector.AutoScalingProperty.McuCount`.
	McuCount *float64 `field:"required" json:"mcuCount" yaml:"mcuCount"`
	// `CfnConnector.AutoScalingProperty.MinWorkerCount`.
	MinWorkerCount *float64 `field:"required" json:"minWorkerCount" yaml:"minWorkerCount"`
	// `CfnConnector.AutoScalingProperty.ScaleInPolicy`.
	ScaleInPolicy interface{} `field:"required" json:"scaleInPolicy" yaml:"scaleInPolicy"`
	// `CfnConnector.AutoScalingProperty.ScaleOutPolicy`.
	ScaleOutPolicy interface{} `field:"required" json:"scaleOutPolicy" yaml:"scaleOutPolicy"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityProperty := &capacityProperty{
//   	autoScaling: &autoScalingProperty{
//   		maxWorkerCount: jsii.Number(123),
//   		mcuCount: jsii.Number(123),
//   		minWorkerCount: jsii.Number(123),
//   		scaleInPolicy: &scaleInPolicyProperty{
//   			cpuUtilizationPercentage: jsii.Number(123),
//   		},
//   		scaleOutPolicy: &scaleOutPolicyProperty{
//   			cpuUtilizationPercentage: jsii.Number(123),
//   		},
//   	},
//   	provisionedCapacity: &provisionedCapacityProperty{
//   		workerCount: jsii.Number(123),
//
//   		// the properties below are optional
//   		mcuCount: jsii.Number(123),
//   	},
//   }
//
type CfnConnector_CapacityProperty struct {
	// `CfnConnector.CapacityProperty.AutoScaling`.
	AutoScaling interface{} `field:"optional" json:"autoScaling" yaml:"autoScaling"`
	// `CfnConnector.CapacityProperty.ProvisionedCapacity`.
	ProvisionedCapacity interface{} `field:"optional" json:"provisionedCapacity" yaml:"provisionedCapacity"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogsLogDeliveryProperty := &cloudWatchLogsLogDeliveryProperty{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	logGroup: jsii.String("logGroup"),
//   }
//
type CfnConnector_CloudWatchLogsLogDeliveryProperty struct {
	// `CfnConnector.CloudWatchLogsLogDeliveryProperty.Enabled`.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// `CfnConnector.CloudWatchLogsLogDeliveryProperty.LogGroup`.
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customPluginProperty := &customPluginProperty{
//   	customPluginArn: jsii.String("customPluginArn"),
//   	revision: jsii.Number(123),
//   }
//
type CfnConnector_CustomPluginProperty struct {
	// `CfnConnector.CustomPluginProperty.CustomPluginArn`.
	CustomPluginArn *string `field:"required" json:"customPluginArn" yaml:"customPluginArn"`
	// `CfnConnector.CustomPluginProperty.Revision`.
	Revision *float64 `field:"required" json:"revision" yaml:"revision"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firehoseLogDeliveryProperty := &firehoseLogDeliveryProperty{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	deliveryStream: jsii.String("deliveryStream"),
//   }
//
type CfnConnector_FirehoseLogDeliveryProperty struct {
	// `CfnConnector.FirehoseLogDeliveryProperty.Enabled`.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// `CfnConnector.FirehoseLogDeliveryProperty.DeliveryStream`.
	DeliveryStream *string `field:"optional" json:"deliveryStream" yaml:"deliveryStream"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kafkaClusterClientAuthenticationProperty := &kafkaClusterClientAuthenticationProperty{
//   	authenticationType: jsii.String("authenticationType"),
//   }
//
type CfnConnector_KafkaClusterClientAuthenticationProperty struct {
	// `CfnConnector.KafkaClusterClientAuthenticationProperty.AuthenticationType`.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kafkaClusterEncryptionInTransitProperty := &kafkaClusterEncryptionInTransitProperty{
//   	encryptionType: jsii.String("encryptionType"),
//   }
//
type CfnConnector_KafkaClusterEncryptionInTransitProperty struct {
	// `CfnConnector.KafkaClusterEncryptionInTransitProperty.EncryptionType`.
	EncryptionType *string `field:"required" json:"encryptionType" yaml:"encryptionType"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kafkaClusterProperty := &kafkaClusterProperty{
//   	apacheKafkaCluster: &apacheKafkaClusterProperty{
//   		bootstrapServers: jsii.String("bootstrapServers"),
//   		vpc: &vpcProperty{
//   			securityGroups: []*string{
//   				jsii.String("securityGroups"),
//   			},
//   			subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   }
//
type CfnConnector_KafkaClusterProperty struct {
	// `CfnConnector.KafkaClusterProperty.ApacheKafkaCluster`.
	ApacheKafkaCluster interface{} `field:"required" json:"apacheKafkaCluster" yaml:"apacheKafkaCluster"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logDeliveryProperty := &logDeliveryProperty{
//   	workerLogDelivery: &workerLogDeliveryProperty{
//   		cloudWatchLogs: &cloudWatchLogsLogDeliveryProperty{
//   			enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			logGroup: jsii.String("logGroup"),
//   		},
//   		firehose: &firehoseLogDeliveryProperty{
//   			enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			deliveryStream: jsii.String("deliveryStream"),
//   		},
//   		s3: &s3LogDeliveryProperty{
//   			enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			bucket: jsii.String("bucket"),
//   			prefix: jsii.String("prefix"),
//   		},
//   	},
//   }
//
type CfnConnector_LogDeliveryProperty struct {
	// `CfnConnector.LogDeliveryProperty.WorkerLogDelivery`.
	WorkerLogDelivery interface{} `field:"required" json:"workerLogDelivery" yaml:"workerLogDelivery"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pluginProperty := &pluginProperty{
//   	customPlugin: &customPluginProperty{
//   		customPluginArn: jsii.String("customPluginArn"),
//   		revision: jsii.Number(123),
//   	},
//   }
//
type CfnConnector_PluginProperty struct {
	// `CfnConnector.PluginProperty.CustomPlugin`.
	CustomPlugin interface{} `field:"required" json:"customPlugin" yaml:"customPlugin"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   provisionedCapacityProperty := &provisionedCapacityProperty{
//   	workerCount: jsii.Number(123),
//
//   	// the properties below are optional
//   	mcuCount: jsii.Number(123),
//   }
//
type CfnConnector_ProvisionedCapacityProperty struct {
	// `CfnConnector.ProvisionedCapacityProperty.WorkerCount`.
	WorkerCount *float64 `field:"required" json:"workerCount" yaml:"workerCount"`
	// `CfnConnector.ProvisionedCapacityProperty.McuCount`.
	McuCount *float64 `field:"optional" json:"mcuCount" yaml:"mcuCount"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LogDeliveryProperty := &s3LogDeliveryProperty{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	bucket: jsii.String("bucket"),
//   	prefix: jsii.String("prefix"),
//   }
//
type CfnConnector_S3LogDeliveryProperty struct {
	// `CfnConnector.S3LogDeliveryProperty.Enabled`.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// `CfnConnector.S3LogDeliveryProperty.Bucket`.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// `CfnConnector.S3LogDeliveryProperty.Prefix`.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scaleInPolicyProperty := &scaleInPolicyProperty{
//   	cpuUtilizationPercentage: jsii.Number(123),
//   }
//
type CfnConnector_ScaleInPolicyProperty struct {
	// `CfnConnector.ScaleInPolicyProperty.CpuUtilizationPercentage`.
	CpuUtilizationPercentage *float64 `field:"required" json:"cpuUtilizationPercentage" yaml:"cpuUtilizationPercentage"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scaleOutPolicyProperty := &scaleOutPolicyProperty{
//   	cpuUtilizationPercentage: jsii.Number(123),
//   }
//
type CfnConnector_ScaleOutPolicyProperty struct {
	// `CfnConnector.ScaleOutPolicyProperty.CpuUtilizationPercentage`.
	CpuUtilizationPercentage *float64 `field:"required" json:"cpuUtilizationPercentage" yaml:"cpuUtilizationPercentage"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcProperty := &vpcProperty{
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   }
//
type CfnConnector_VpcProperty struct {
	// `CfnConnector.VpcProperty.SecurityGroups`.
	SecurityGroups *[]*string `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// `CfnConnector.VpcProperty.Subnets`.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workerConfigurationProperty := &workerConfigurationProperty{
//   	revision: jsii.Number(123),
//   	workerConfigurationArn: jsii.String("workerConfigurationArn"),
//   }
//
type CfnConnector_WorkerConfigurationProperty struct {
	// `CfnConnector.WorkerConfigurationProperty.Revision`.
	Revision *float64 `field:"required" json:"revision" yaml:"revision"`
	// `CfnConnector.WorkerConfigurationProperty.WorkerConfigurationArn`.
	WorkerConfigurationArn *string `field:"required" json:"workerConfigurationArn" yaml:"workerConfigurationArn"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workerLogDeliveryProperty := &workerLogDeliveryProperty{
//   	cloudWatchLogs: &cloudWatchLogsLogDeliveryProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		logGroup: jsii.String("logGroup"),
//   	},
//   	firehose: &firehoseLogDeliveryProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		deliveryStream: jsii.String("deliveryStream"),
//   	},
//   	s3: &s3LogDeliveryProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		bucket: jsii.String("bucket"),
//   		prefix: jsii.String("prefix"),
//   	},
//   }
//
type CfnConnector_WorkerLogDeliveryProperty struct {
	// `CfnConnector.WorkerLogDeliveryProperty.CloudWatchLogs`.
	CloudWatchLogs interface{} `field:"optional" json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
	// `CfnConnector.WorkerLogDeliveryProperty.Firehose`.
	Firehose interface{} `field:"optional" json:"firehose" yaml:"firehose"`
	// `CfnConnector.WorkerLogDeliveryProperty.S3`.
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

// Properties for defining a `CfnConnector`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectorProps := &cfnConnectorProps{
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
//   }
//
type CfnConnectorProps struct {
	// `AWS::KafkaConnect::Connector.Capacity`.
	Capacity interface{} `field:"required" json:"capacity" yaml:"capacity"`
	// `AWS::KafkaConnect::Connector.ConnectorConfiguration`.
	ConnectorConfiguration interface{} `field:"required" json:"connectorConfiguration" yaml:"connectorConfiguration"`
	// `AWS::KafkaConnect::Connector.ConnectorName`.
	ConnectorName *string `field:"required" json:"connectorName" yaml:"connectorName"`
	// `AWS::KafkaConnect::Connector.KafkaCluster`.
	KafkaCluster interface{} `field:"required" json:"kafkaCluster" yaml:"kafkaCluster"`
	// `AWS::KafkaConnect::Connector.KafkaClusterClientAuthentication`.
	KafkaClusterClientAuthentication interface{} `field:"required" json:"kafkaClusterClientAuthentication" yaml:"kafkaClusterClientAuthentication"`
	// `AWS::KafkaConnect::Connector.KafkaClusterEncryptionInTransit`.
	KafkaClusterEncryptionInTransit interface{} `field:"required" json:"kafkaClusterEncryptionInTransit" yaml:"kafkaClusterEncryptionInTransit"`
	// `AWS::KafkaConnect::Connector.KafkaConnectVersion`.
	KafkaConnectVersion *string `field:"required" json:"kafkaConnectVersion" yaml:"kafkaConnectVersion"`
	// `AWS::KafkaConnect::Connector.Plugins`.
	Plugins interface{} `field:"required" json:"plugins" yaml:"plugins"`
	// `AWS::KafkaConnect::Connector.ServiceExecutionRoleArn`.
	ServiceExecutionRoleArn *string `field:"required" json:"serviceExecutionRoleArn" yaml:"serviceExecutionRoleArn"`
	// `AWS::KafkaConnect::Connector.ConnectorDescription`.
	ConnectorDescription *string `field:"optional" json:"connectorDescription" yaml:"connectorDescription"`
	// `AWS::KafkaConnect::Connector.LogDelivery`.
	LogDelivery interface{} `field:"optional" json:"logDelivery" yaml:"logDelivery"`
	// `AWS::KafkaConnect::Connector.WorkerConfiguration`.
	WorkerConfiguration interface{} `field:"optional" json:"workerConfiguration" yaml:"workerConfiguration"`
}


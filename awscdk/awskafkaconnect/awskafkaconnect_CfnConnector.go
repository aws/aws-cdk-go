package awskafkaconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awskafkaconnect/internal"
	"github.com/aws/constructs-go/constructs/v3"
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
//   cfnConnector := awscdk.Aws_kafkaconnect.NewCfnConnector(this, jsii.String("MyCfnConnector"), &CfnConnectorProps{
//   	Capacity: &CapacityProperty{
//   		AutoScaling: &AutoScalingProperty{
//   			MaxWorkerCount: jsii.Number(123),
//   			McuCount: jsii.Number(123),
//   			MinWorkerCount: jsii.Number(123),
//   			ScaleInPolicy: &ScaleInPolicyProperty{
//   				CpuUtilizationPercentage: jsii.Number(123),
//   			},
//   			ScaleOutPolicy: &ScaleOutPolicyProperty{
//   				CpuUtilizationPercentage: jsii.Number(123),
//   			},
//   		},
//   		ProvisionedCapacity: &ProvisionedCapacityProperty{
//   			WorkerCount: jsii.Number(123),
//
//   			// the properties below are optional
//   			McuCount: jsii.Number(123),
//   		},
//   	},
//   	ConnectorConfiguration: map[string]*string{
//   		"connectorConfigurationKey": jsii.String("connectorConfiguration"),
//   	},
//   	ConnectorName: jsii.String("connectorName"),
//   	KafkaCluster: &KafkaClusterProperty{
//   		ApacheKafkaCluster: &ApacheKafkaClusterProperty{
//   			BootstrapServers: jsii.String("bootstrapServers"),
//   			Vpc: &VpcProperty{
//   				SecurityGroups: []*string{
//   					jsii.String("securityGroups"),
//   				},
//   				Subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//   			},
//   		},
//   	},
//   	KafkaClusterClientAuthentication: &KafkaClusterClientAuthenticationProperty{
//   		AuthenticationType: jsii.String("authenticationType"),
//   	},
//   	KafkaClusterEncryptionInTransit: &KafkaClusterEncryptionInTransitProperty{
//   		EncryptionType: jsii.String("encryptionType"),
//   	},
//   	KafkaConnectVersion: jsii.String("kafkaConnectVersion"),
//   	Plugins: []interface{}{
//   		&PluginProperty{
//   			CustomPlugin: &CustomPluginProperty{
//   				CustomPluginArn: jsii.String("customPluginArn"),
//   				Revision: jsii.Number(123),
//   			},
//   		},
//   	},
//   	ServiceExecutionRoleArn: jsii.String("serviceExecutionRoleArn"),
//
//   	// the properties below are optional
//   	ConnectorDescription: jsii.String("connectorDescription"),
//   	LogDelivery: &LogDeliveryProperty{
//   		WorkerLogDelivery: &WorkerLogDeliveryProperty{
//   			CloudWatchLogs: &CloudWatchLogsLogDeliveryProperty{
//   				Enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				LogGroup: jsii.String("logGroup"),
//   			},
//   			Firehose: &FirehoseLogDeliveryProperty{
//   				Enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				DeliveryStream: jsii.String("deliveryStream"),
//   			},
//   			S3: &S3LogDeliveryProperty{
//   				Enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				Bucket: jsii.String("bucket"),
//   				Prefix: jsii.String("prefix"),
//   			},
//   		},
//   	},
//   	WorkerConfiguration: &WorkerConfigurationProperty{
//   		Revision: jsii.Number(123),
//   		WorkerConfigurationArn: jsii.String("workerConfigurationArn"),
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
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
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
	// Experimental.
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
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Specifies which plugin to use for the connector.
	//
	// You must specify a single-element list. Amazon MSK Connect does not currently support specifying multiple plugins.
	Plugins() interface{}
	SetPlugins(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The Amazon Resource Name (ARN) of the IAM role used by the connector to access Amazon Web Services resources.
	ServiceExecutionRoleArn() *string
	SetServiceExecutionRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The worker configurations that are in use with the connector.
	WorkerConfiguration() interface{}
	SetWorkerConfiguration(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
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

func (j *jsiiProxy_CfnConnector) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnConnector(scope awscdk.Construct, id *string, props *CfnConnectorProps) CfnConnector {
	_init_.Initialize()

	if err := validateNewCfnConnectorParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConnector{}

	_jsii_.Create(
		"monocdk.aws_kafkaconnect.CfnConnector",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::KafkaConnect::Connector`.
func NewCfnConnector_Override(c CfnConnector, scope awscdk.Construct, id *string, props *CfnConnectorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_kafkaconnect.CfnConnector",
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
// Experimental.
func CfnConnector_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConnector_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kafkaconnect.CfnConnector",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnConnector_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnConnector_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kafkaconnect.CfnConnector",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnConnector_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConnector_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kafkaconnect.CfnConnector",
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
		"monocdk.aws_kafkaconnect.CfnConnector",
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

func (c *jsiiProxy_CfnConnector) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnConnector) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnConnector) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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

func (c *jsiiProxy_CfnConnector) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnConnector) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

func (c *jsiiProxy_CfnConnector) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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


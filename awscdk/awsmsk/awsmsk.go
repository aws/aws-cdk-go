package awsmsk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/awsmsk/internal"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/constructs-go/constructs/v3"
)

// Configuration details related to broker logs.
// Experimental.
type BrokerLogging struct {
	// The CloudWatch Logs group that is the destination for broker logs.
	// Experimental.
	CloudwatchLogGroup awslogs.ILogGroup `json:"cloudwatchLogGroup"`
	// The Kinesis Data Firehose delivery stream that is the destination for broker logs.
	// Experimental.
	FirehoseDeliveryStreamName *string `json:"firehoseDeliveryStreamName"`
	// Details of the Amazon S3 destination for broker logs.
	// Experimental.
	S3 *S3LoggingConfiguration `json:"s3"`
}

// A CloudFormation `AWS::MSK::Cluster`.
type CfnCluster interface {
	awscdk.CfnResource
	awscdk.IInspectable
	BrokerNodeGroupInfo() interface{}
	SetBrokerNodeGroupInfo(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ClientAuthentication() interface{}
	SetClientAuthentication(val interface{})
	ClusterName() *string
	SetClusterName(val *string)
	ConfigurationInfo() interface{}
	SetConfigurationInfo(val interface{})
	CreationStack() *[]*string
	EncryptionInfo() interface{}
	SetEncryptionInfo(val interface{})
	EnhancedMonitoring() *string
	SetEnhancedMonitoring(val *string)
	KafkaVersion() *string
	SetKafkaVersion(val *string)
	LoggingInfo() interface{}
	SetLoggingInfo(val interface{})
	LogicalId() *string
	Node() awscdk.ConstructNode
	NumberOfBrokerNodes() *float64
	SetNumberOfBrokerNodes(val *float64)
	OpenMonitoring() interface{}
	SetOpenMonitoring(val interface{})
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnCluster
type jsiiProxy_CfnCluster struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCluster) BrokerNodeGroupInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"brokerNodeGroupInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ClientAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ConfigurationInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configurationInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) EncryptionInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) EnhancedMonitoring() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enhancedMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) KafkaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kafkaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) LoggingInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) NumberOfBrokerNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBrokerNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) OpenMonitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::MSK::Cluster`.
func NewCfnCluster(scope awscdk.Construct, id *string, props *CfnClusterProps) CfnCluster {
	_init_.Initialize()

	j := jsiiProxy_CfnCluster{}

	_jsii_.Create(
		"monocdk.aws_msk.CfnCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MSK::Cluster`.
func NewCfnCluster_Override(c CfnCluster, scope awscdk.Construct, id *string, props *CfnClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_msk.CfnCluster",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCluster) SetBrokerNodeGroupInfo(val interface{}) {
	_jsii_.Set(
		j,
		"brokerNodeGroupInfo",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetClientAuthentication(val interface{}) {
	_jsii_.Set(
		j,
		"clientAuthentication",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetClusterName(val *string) {
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetConfigurationInfo(val interface{}) {
	_jsii_.Set(
		j,
		"configurationInfo",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetEncryptionInfo(val interface{}) {
	_jsii_.Set(
		j,
		"encryptionInfo",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetEnhancedMonitoring(val *string) {
	_jsii_.Set(
		j,
		"enhancedMonitoring",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetKafkaVersion(val *string) {
	_jsii_.Set(
		j,
		"kafkaVersion",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetLoggingInfo(val interface{}) {
	_jsii_.Set(
		j,
		"loggingInfo",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetNumberOfBrokerNodes(val *float64) {
	_jsii_.Set(
		j,
		"numberOfBrokerNodes",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetOpenMonitoring(val interface{}) {
	_jsii_.Set(
		j,
		"openMonitoring",
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
func CfnCluster_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_msk.CfnCluster",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnCluster_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_msk.CfnCluster",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_msk.CfnCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCluster_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_msk.CfnCluster",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnCluster) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnCluster) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnCluster) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
// Experimental.
func (c *jsiiProxy_CfnCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnCluster) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnCluster) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnCluster) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnCluster) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnCluster) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnCluster) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnCluster) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnCluster) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnCluster) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnCluster) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnCluster) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnCluster) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnCluster) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnCluster) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnCluster_BrokerLogsProperty struct {
	// `CfnCluster.BrokerLogsProperty.CloudWatchLogs`.
	CloudWatchLogs interface{} `json:"cloudWatchLogs"`
	// `CfnCluster.BrokerLogsProperty.Firehose`.
	Firehose interface{} `json:"firehose"`
	// `CfnCluster.BrokerLogsProperty.S3`.
	S3 interface{} `json:"s3"`
}

type CfnCluster_BrokerNodeGroupInfoProperty struct {
	// `CfnCluster.BrokerNodeGroupInfoProperty.ClientSubnets`.
	ClientSubnets *[]*string `json:"clientSubnets"`
	// `CfnCluster.BrokerNodeGroupInfoProperty.InstanceType`.
	InstanceType *string `json:"instanceType"`
	// `CfnCluster.BrokerNodeGroupInfoProperty.BrokerAZDistribution`.
	BrokerAzDistribution *string `json:"brokerAzDistribution"`
	// `CfnCluster.BrokerNodeGroupInfoProperty.SecurityGroups`.
	SecurityGroups *[]*string `json:"securityGroups"`
	// `CfnCluster.BrokerNodeGroupInfoProperty.StorageInfo`.
	StorageInfo interface{} `json:"storageInfo"`
}

type CfnCluster_ClientAuthenticationProperty struct {
	// `CfnCluster.ClientAuthenticationProperty.Sasl`.
	Sasl interface{} `json:"sasl"`
	// `CfnCluster.ClientAuthenticationProperty.Tls`.
	Tls interface{} `json:"tls"`
}

type CfnCluster_CloudWatchLogsProperty struct {
	// `CfnCluster.CloudWatchLogsProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `CfnCluster.CloudWatchLogsProperty.LogGroup`.
	LogGroup *string `json:"logGroup"`
}

type CfnCluster_ConfigurationInfoProperty struct {
	// `CfnCluster.ConfigurationInfoProperty.Arn`.
	Arn *string `json:"arn"`
	// `CfnCluster.ConfigurationInfoProperty.Revision`.
	Revision *float64 `json:"revision"`
}

type CfnCluster_EBSStorageInfoProperty struct {
	// `CfnCluster.EBSStorageInfoProperty.VolumeSize`.
	VolumeSize *float64 `json:"volumeSize"`
}

type CfnCluster_EncryptionAtRestProperty struct {
	// `CfnCluster.EncryptionAtRestProperty.DataVolumeKMSKeyId`.
	DataVolumeKmsKeyId *string `json:"dataVolumeKmsKeyId"`
}

type CfnCluster_EncryptionInTransitProperty struct {
	// `CfnCluster.EncryptionInTransitProperty.ClientBroker`.
	ClientBroker *string `json:"clientBroker"`
	// `CfnCluster.EncryptionInTransitProperty.InCluster`.
	InCluster interface{} `json:"inCluster"`
}

type CfnCluster_EncryptionInfoProperty struct {
	// `CfnCluster.EncryptionInfoProperty.EncryptionAtRest`.
	EncryptionAtRest interface{} `json:"encryptionAtRest"`
	// `CfnCluster.EncryptionInfoProperty.EncryptionInTransit`.
	EncryptionInTransit interface{} `json:"encryptionInTransit"`
}

type CfnCluster_FirehoseProperty struct {
	// `CfnCluster.FirehoseProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `CfnCluster.FirehoseProperty.DeliveryStream`.
	DeliveryStream *string `json:"deliveryStream"`
}

type CfnCluster_IamProperty struct {
	// `CfnCluster.IamProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
}

type CfnCluster_JmxExporterProperty struct {
	// `CfnCluster.JmxExporterProperty.EnabledInBroker`.
	EnabledInBroker interface{} `json:"enabledInBroker"`
}

type CfnCluster_LoggingInfoProperty struct {
	// `CfnCluster.LoggingInfoProperty.BrokerLogs`.
	BrokerLogs interface{} `json:"brokerLogs"`
}

type CfnCluster_NodeExporterProperty struct {
	// `CfnCluster.NodeExporterProperty.EnabledInBroker`.
	EnabledInBroker interface{} `json:"enabledInBroker"`
}

type CfnCluster_OpenMonitoringProperty struct {
	// `CfnCluster.OpenMonitoringProperty.Prometheus`.
	Prometheus interface{} `json:"prometheus"`
}

type CfnCluster_PrometheusProperty struct {
	// `CfnCluster.PrometheusProperty.JmxExporter`.
	JmxExporter interface{} `json:"jmxExporter"`
	// `CfnCluster.PrometheusProperty.NodeExporter`.
	NodeExporter interface{} `json:"nodeExporter"`
}

type CfnCluster_S3Property struct {
	// `CfnCluster.S3Property.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `CfnCluster.S3Property.Bucket`.
	Bucket *string `json:"bucket"`
	// `CfnCluster.S3Property.Prefix`.
	Prefix *string `json:"prefix"`
}

type CfnCluster_SaslProperty struct {
	// `CfnCluster.SaslProperty.Iam`.
	Iam interface{} `json:"iam"`
	// `CfnCluster.SaslProperty.Scram`.
	Scram interface{} `json:"scram"`
}

type CfnCluster_ScramProperty struct {
	// `CfnCluster.ScramProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
}

type CfnCluster_StorageInfoProperty struct {
	// `CfnCluster.StorageInfoProperty.EBSStorageInfo`.
	EbsStorageInfo interface{} `json:"ebsStorageInfo"`
}

type CfnCluster_TlsProperty struct {
	// `CfnCluster.TlsProperty.CertificateAuthorityArnList`.
	CertificateAuthorityArnList *[]*string `json:"certificateAuthorityArnList"`
}

// Properties for defining a `AWS::MSK::Cluster`.
type CfnClusterProps struct {
	// `AWS::MSK::Cluster.BrokerNodeGroupInfo`.
	BrokerNodeGroupInfo interface{} `json:"brokerNodeGroupInfo"`
	// `AWS::MSK::Cluster.ClusterName`.
	ClusterName *string `json:"clusterName"`
	// `AWS::MSK::Cluster.KafkaVersion`.
	KafkaVersion *string `json:"kafkaVersion"`
	// `AWS::MSK::Cluster.NumberOfBrokerNodes`.
	NumberOfBrokerNodes *float64 `json:"numberOfBrokerNodes"`
	// `AWS::MSK::Cluster.ClientAuthentication`.
	ClientAuthentication interface{} `json:"clientAuthentication"`
	// `AWS::MSK::Cluster.ConfigurationInfo`.
	ConfigurationInfo interface{} `json:"configurationInfo"`
	// `AWS::MSK::Cluster.EncryptionInfo`.
	EncryptionInfo interface{} `json:"encryptionInfo"`
	// `AWS::MSK::Cluster.EnhancedMonitoring`.
	EnhancedMonitoring *string `json:"enhancedMonitoring"`
	// `AWS::MSK::Cluster.LoggingInfo`.
	LoggingInfo interface{} `json:"loggingInfo"`
	// `AWS::MSK::Cluster.OpenMonitoring`.
	OpenMonitoring interface{} `json:"openMonitoring"`
	// `AWS::MSK::Cluster.Tags`.
	Tags interface{} `json:"tags"`
}

// Configuration properties for client authentication.
// Experimental.
type ClientAuthentication interface {
	SaslProps() *SaslAuthProps
	TlsProps() *TlsAuthProps
}

// The jsii proxy struct for ClientAuthentication
type jsiiProxy_ClientAuthentication struct {
	_ byte // padding
}

func (j *jsiiProxy_ClientAuthentication) SaslProps() *SaslAuthProps {
	var returns *SaslAuthProps
	_jsii_.Get(
		j,
		"saslProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClientAuthentication) TlsProps() *TlsAuthProps {
	var returns *TlsAuthProps
	_jsii_.Get(
		j,
		"tlsProps",
		&returns,
	)
	return returns
}


// SASL authentication.
// Experimental.
func ClientAuthentication_Sasl(props *SaslAuthProps) ClientAuthentication {
	_init_.Initialize()

	var returns ClientAuthentication

	_jsii_.StaticInvoke(
		"monocdk.aws_msk.ClientAuthentication",
		"sasl",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// TLS authentication.
// Experimental.
func ClientAuthentication_Tls(props *TlsAuthProps) ClientAuthentication {
	_init_.Initialize()

	var returns ClientAuthentication

	_jsii_.StaticInvoke(
		"monocdk.aws_msk.ClientAuthentication",
		"tls",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Indicates the encryption setting for data in transit between clients and brokers.
// Experimental.
type ClientBrokerEncryption string

const (
	ClientBrokerEncryption_TLS ClientBrokerEncryption = "TLS"
	ClientBrokerEncryption_TLS_PLAINTEXT ClientBrokerEncryption = "TLS_PLAINTEXT"
	ClientBrokerEncryption_PLAINTEXT ClientBrokerEncryption = "PLAINTEXT"
)

// Create a MSK Cluster.
// Experimental.
type Cluster interface {
	awscdk.Resource
	ICluster
	BootstrapBrokers() *string
	BootstrapBrokersSaslScram() *string
	BootstrapBrokersTls() *string
	ClusterArn() *string
	ClusterName() *string
	Connections() awsec2.Connections
	Env() *awscdk.ResourceEnvironment
	Node() awscdk.ConstructNode
	PhysicalName() *string
	SaslScramAuthenticationKey() awskms.IKey
	Stack() awscdk.Stack
	ZookeeperConnectionString() *string
	ZookeeperConnectionStringTls() *string
	AddUser(usernames ...*string)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for Cluster
type jsiiProxy_Cluster struct {
	internal.Type__awscdkResource
	jsiiProxy_ICluster
}

func (j *jsiiProxy_Cluster) BootstrapBrokers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) BootstrapBrokersSaslScram() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersSaslScram",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) BootstrapBrokersTls() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) SaslScramAuthenticationKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"saslScramAuthenticationKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ZookeeperConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zookeeperConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ZookeeperConnectionStringTls() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zookeeperConnectionStringTls",
		&returns,
	)
	return returns
}


// Experimental.
func NewCluster(scope constructs.Construct, id *string, props *ClusterProps) Cluster {
	_init_.Initialize()

	j := jsiiProxy_Cluster{}

	_jsii_.Create(
		"monocdk.aws_msk.Cluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCluster_Override(c Cluster, scope constructs.Construct, id *string, props *ClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_msk.Cluster",
		[]interface{}{scope, id, props},
		c,
	)
}

// Reference an existing cluster, defined outside of the CDK code, by name.
// Experimental.
func Cluster_FromClusterArn(scope constructs.Construct, id *string, clusterArn *string) ICluster {
	_init_.Initialize()

	var returns ICluster

	_jsii_.StaticInvoke(
		"monocdk.aws_msk.Cluster",
		"fromClusterArn",
		[]interface{}{scope, id, clusterArn},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Cluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_msk.Cluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Cluster_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_msk.Cluster",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// A list of usersnames to register with the cluster.
//
// The password will automatically be generated using Secrets
// Manager and the { username, password } JSON object stored in Secrets Manager as `AmazonMSK_username`.
//
// Must be using the SASL/SCRAM authentication mechanism.
// Experimental.
func (c *jsiiProxy_Cluster) AddUser(usernames ...*string) {
	args := []interface{}{}
	for _, a := range usernames {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addUser",
		args,
	)
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (c *jsiiProxy_Cluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (c *jsiiProxy_Cluster) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (c *jsiiProxy_Cluster) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (c *jsiiProxy_Cluster) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_Cluster) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_Cluster) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_Cluster) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_Cluster) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_Cluster) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (c *jsiiProxy_Cluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_Cluster) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The Amazon MSK configuration to use for the cluster.
//
// Note: There is currently no Cloudformation Resource to create a Configuration
// Experimental.
type ClusterConfigurationInfo struct {
	// The Amazon Resource Name (ARN) of the MSK configuration to use.
	//
	// For example, arn:aws:kafka:us-east-1:123456789012:configuration/example-configuration-name/abcdabcd-1234-abcd-1234-abcd123e8e8e-1.
	// Experimental.
	Arn *string `json:"arn"`
	// The revision of the Amazon MSK configuration to use.
	// Experimental.
	Revision *float64 `json:"revision"`
}

// The level of monitoring for the MSK cluster.
// See: https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html#metrics-details
//
// Experimental.
type ClusterMonitoringLevel string

const (
	ClusterMonitoringLevel_DEFAULT ClusterMonitoringLevel = "DEFAULT"
	ClusterMonitoringLevel_PER_BROKER ClusterMonitoringLevel = "PER_BROKER"
	ClusterMonitoringLevel_PER_TOPIC_PER_BROKER ClusterMonitoringLevel = "PER_TOPIC_PER_BROKER"
	ClusterMonitoringLevel_PER_TOPIC_PER_PARTITION ClusterMonitoringLevel = "PER_TOPIC_PER_PARTITION"
)

// Properties for a MSK Cluster.
// Experimental.
type ClusterProps struct {
	// The physical name of the cluster.
	// Experimental.
	ClusterName *string `json:"clusterName"`
	// The version of Apache Kafka.
	// Experimental.
	KafkaVersion KafkaVersion `json:"kafkaVersion"`
	// Defines the virtual networking environment for this cluster.
	//
	// Must have at least 2 subnets in two different AZs.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// Configuration properties for client authentication.
	//
	// MSK supports using private TLS certificates or SASL/SCRAM to authenticate the identity of clients.
	// Experimental.
	ClientAuthentication ClientAuthentication `json:"clientAuthentication"`
	// The Amazon MSK configuration to use for the cluster.
	// Experimental.
	ConfigurationInfo *ClusterConfigurationInfo `json:"configurationInfo"`
	// Information about storage volumes attached to MSK broker nodes.
	// Experimental.
	EbsStorageInfo *EbsStorageInfo `json:"ebsStorageInfo"`
	// Config details for encryption in transit.
	// Experimental.
	EncryptionInTransit *EncryptionInTransitConfig `json:"encryptionInTransit"`
	// The EC2 instance type that you want Amazon MSK to use when it creates your brokers.
	// See: https://docs.aws.amazon.com/msk/latest/developerguide/msk-create-cluster.html#broker-instance-types
	//
	// Experimental.
	InstanceType awsec2.InstanceType `json:"instanceType"`
	// Configure your MSK cluster to send broker logs to different destination types.
	// Experimental.
	Logging *BrokerLogging `json:"logging"`
	// Cluster monitoring configuration.
	// Experimental.
	Monitoring *MonitoringConfiguration `json:"monitoring"`
	// Number of Apache Kafka brokers deployed in each Availability Zone.
	// Experimental.
	NumberOfBrokerNodes *float64 `json:"numberOfBrokerNodes"`
	// What to do when this resource is deleted from a stack.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
	// The AWS security groups to associate with the elastic network interfaces in order to specify who can connect to and communicate with the Amazon MSK cluster.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// Where to place the nodes within the VPC.
	//
	// Amazon MSK distributes the broker nodes evenly across the subnets that you specify.
	// The subnets that you specify must be in distinct Availability Zones.
	// Client subnets can't be in Availability Zone us-east-1e.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets"`
}

// EBS volume information.
// Experimental.
type EbsStorageInfo struct {
	// The AWS KMS key for encrypting data at rest.
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey"`
	// The size in GiB of the EBS volume for the data drive on each broker node.
	// Experimental.
	VolumeSize *float64 `json:"volumeSize"`
}

// The settings for encrypting data in transit.
// See: https://docs.aws.amazon.com/msk/latest/developerguide/msk-encryption.html#msk-encryption-in-transit
//
// Experimental.
type EncryptionInTransitConfig struct {
	// Indicates the encryption setting for data in transit between clients and brokers.
	// Experimental.
	ClientBroker ClientBrokerEncryption `json:"clientBroker"`
	// Indicates that data communication among the broker nodes of the cluster is encrypted.
	// Experimental.
	EnableInCluster *bool `json:"enableInCluster"`
}

// Represents a MSK Cluster.
// Experimental.
type ICluster interface {
	awsec2.IConnectable
	awscdk.IResource
	// The ARN of cluster.
	// Experimental.
	ClusterArn() *string
	// The physical name of the cluster.
	// Experimental.
	ClusterName() *string
}

// The jsii proxy for ICluster
type jsiiProxy_ICluster struct {
	internal.Type__awsec2IConnectable
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ICluster) ClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// Kafka cluster version.
// Experimental.
type KafkaVersion interface {
	Version() *string
}

// The jsii proxy struct for KafkaVersion
type jsiiProxy_KafkaVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_KafkaVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Custom cluster version.
// Experimental.
func KafkaVersion_Of(version *string) KafkaVersion {
	_init_.Initialize()

	var returns KafkaVersion

	_jsii_.StaticInvoke(
		"monocdk.aws_msk.KafkaVersion",
		"of",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func KafkaVersion_V1_1_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V1_1_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_2_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_2_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_3_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_3_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_4_1_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_4_1_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_5_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_5_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_6_0() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_6_0",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_6_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_6_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_7_0() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_7_0",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_8_0() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_8_0",
		&returns,
	)
	return returns
}

// Monitoring Configuration.
// Experimental.
type MonitoringConfiguration struct {
	// Specifies the level of monitoring for the MSK cluster.
	// Experimental.
	ClusterMonitoringLevel ClusterMonitoringLevel `json:"clusterMonitoringLevel"`
	// Indicates whether you want to enable or disable the JMX Exporter.
	// Experimental.
	EnablePrometheusJmxExporter *bool `json:"enablePrometheusJmxExporter"`
	// Indicates whether you want to enable or disable the Prometheus Node Exporter.
	//
	// You can use the Prometheus Node Exporter to get CPU and disk metrics for the broker nodes.
	// Experimental.
	EnablePrometheusNodeExporter *bool `json:"enablePrometheusNodeExporter"`
}

// Details of the Amazon S3 destination for broker logs.
// Experimental.
type S3LoggingConfiguration struct {
	// The S3 bucket that is the destination for broker logs.
	// Experimental.
	Bucket awss3.IBucket `json:"bucket"`
	// The S3 prefix that is the destination for broker logs.
	// Experimental.
	Prefix *string `json:"prefix"`
}

// SASL authentication properties.
// Experimental.
type SaslAuthProps struct {
	// KMS Key to encrypt SASL/SCRAM secrets.
	//
	// You must use a customer master key (CMK) when creating users in secrets manager.
	// You cannot use a Secret with Amazon MSK that uses the default Secrets Manager encryption key.
	// Experimental.
	Key awskms.IKey `json:"key"`
	// Enable SASL/SCRAM authentication.
	// Experimental.
	Scram *bool `json:"scram"`
}

// TLS authentication properties.
// Experimental.
type TlsAuthProps struct {
	// List of ACM Certificate Authorities to enable TLS authentication.
	// Experimental.
	CertificateAuthorityArns *[]*string `json:"certificateAuthorityArns"`
}


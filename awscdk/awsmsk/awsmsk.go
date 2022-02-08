package awsmsk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmsk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::MSK::Cluster`.
//
// The `AWS::MSK::Cluster` resource creates an Amazon MSK cluster . For more information, see [What Is Amazon MSK?](https://docs.aws.amazon.com/msk/latest/developerguide/what-is-msk.html) in the *Amazon MSK Developer Guide* .
//
// TODO: EXAMPLE
//
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
	Node() constructs.Node
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
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

func (j *jsiiProxy_CfnCluster) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnCluster(scope constructs.Construct, id *string, props *CfnClusterProps) CfnCluster {
	_init_.Initialize()

	j := jsiiProxy_CfnCluster{}

	_jsii_.Create(
		"aws-cdk-lib.aws_msk.CfnCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MSK::Cluster`.
func NewCfnCluster_Override(c CfnCluster, scope constructs.Construct, id *string, props *CfnClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_msk.CfnCluster",
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
func CfnCluster_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_msk.CfnCluster",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCluster_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_msk.CfnCluster",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_msk.CfnCluster",
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
		"aws-cdk-lib.aws_msk.CfnCluster",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
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
func (c *jsiiProxy_CfnCluster) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
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

func (c *jsiiProxy_CfnCluster) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// You can configure your Amazon MSK cluster to send broker logs to different destination types.
//
// This configuration specifies the details of these destinations.
//
// TODO: EXAMPLE
//
type CfnCluster_BrokerLogsProperty struct {
	// Details of the CloudWatch Logs destination for broker logs.
	CloudWatchLogs interface{} `json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
	// Details of the Kinesis Data Firehose delivery stream that is the destination for broker logs.
	Firehose interface{} `json:"firehose" yaml:"firehose"`
	// Details of the Amazon MSK destination for broker logs.
	S3 interface{} `json:"s3" yaml:"s3"`
}

// The setup to be used for brokers in the cluster.
//
// TODO: EXAMPLE
//
type CfnCluster_BrokerNodeGroupInfoProperty struct {
	// The list of subnets to connect to in the client virtual private cloud (VPC).
	//
	// Amazon creates elastic network interfaces inside these subnets. Client applications use elastic network interfaces to produce and consume data.
	//
	// Specify exactly two subnets if you are using the US West (N. California) Region. For other Regions where Amazon MSK is available, you can specify either two or three subnets. The subnets that you specify must be in distinct Availability Zones. When you create a cluster, Amazon MSK distributes the broker nodes evenly across the subnets that you specify.
	//
	// Client subnets can't be in Availability Zone us-east-1e.
	ClientSubnets *[]*string `json:"clientSubnets" yaml:"clientSubnets"`
	// The type of Amazon EC2 instances to use for brokers.
	//
	// The following instance types are allowed: kafka.m5.large, kafka.m5.xlarge, kafka.m5.2xlarge, kafka.m5.4xlarge, kafka.m5.8xlarge, kafka.m5.12xlarge, kafka.m5.16xlarge, and kafka.m5.24xlarge.
	InstanceType *string `json:"instanceType" yaml:"instanceType"`
	// This parameter is currently not in use.
	BrokerAzDistribution *string `json:"brokerAzDistribution" yaml:"brokerAzDistribution"`
	// Information about the cluster's connectivity setting.
	ConnectivityInfo interface{} `json:"connectivityInfo" yaml:"connectivityInfo"`
	// The security groups to associate with the elastic network interfaces in order to specify who can connect to and communicate with the Amazon MSK cluster.
	//
	// If you don't specify a security group, Amazon MSK uses the default security group associated with the VPC. If you specify security groups that were shared with you, you must ensure that you have permissions to them. Specifically, you need the `ec2:DescribeSecurityGroups` permission.
	SecurityGroups *[]*string `json:"securityGroups" yaml:"securityGroups"`
	// Contains information about storage volumes attached to MSK broker nodes.
	StorageInfo interface{} `json:"storageInfo" yaml:"storageInfo"`
}

// Includes information related to client authentication.
//
// TODO: EXAMPLE
//
type CfnCluster_ClientAuthenticationProperty struct {
	// Details for ClientAuthentication using SASL.
	Sasl interface{} `json:"sasl" yaml:"sasl"`
	// Details for client authentication using TLS.
	Tls interface{} `json:"tls" yaml:"tls"`
	// Details for ClientAuthentication using no authentication.
	Unauthenticated interface{} `json:"unauthenticated" yaml:"unauthenticated"`
}

// Details of the CloudWatch Logs destination for broker logs.
//
// TODO: EXAMPLE
//
type CfnCluster_CloudWatchLogsProperty struct {
	// Specifies whether broker logs get sent to the specified CloudWatch Logs destination.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// The CloudWatch Logs group that is the destination for broker logs.
	LogGroup *string `json:"logGroup" yaml:"logGroup"`
}

// Specifies the Amazon MSK configuration to use for the brokers.
//
// TODO: EXAMPLE
//
type CfnCluster_ConfigurationInfoProperty struct {
	// The Amazon Resource Name (ARN) of the MSK configuration to use.
	//
	// For example, `arn:aws:kafka:us-east-1:123456789012:configuration/example-configuration-name/abcdabcd-1234-abcd-1234-abcd123e8e8e-1` .
	Arn *string `json:"arn" yaml:"arn"`
	// The revision of the Amazon MSK configuration to use.
	Revision *float64 `json:"revision" yaml:"revision"`
}

// Specifies whether the cluster's brokers are publicly accessible.
//
// By default, they are not.
//
// TODO: EXAMPLE
//
type CfnCluster_ConnectivityInfoProperty struct {
	// Specifies whether the cluster's brokers are accessible from the internet.
	//
	// Public access is off by default.
	PublicAccess interface{} `json:"publicAccess" yaml:"publicAccess"`
}

// Contains information about the EBS storage volumes attached to brokers.
//
// TODO: EXAMPLE
//
type CfnCluster_EBSStorageInfoProperty struct {
	// `CfnCluster.EBSStorageInfoProperty.ProvisionedThroughput`.
	ProvisionedThroughput interface{} `json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// The size in GiB of the EBS volume for the data drive on each broker node.
	VolumeSize *float64 `json:"volumeSize" yaml:"volumeSize"`
}

// The data volume encryption details.
//
// TODO: EXAMPLE
//
type CfnCluster_EncryptionAtRestProperty struct {
	// The ARN of the Amazon KMS key for encrypting data at rest.
	//
	// If you don't specify a KMS key, MSK creates one for you and uses it on your behalf.
	DataVolumeKmsKeyId *string `json:"dataVolumeKmsKeyId" yaml:"dataVolumeKmsKeyId"`
}

// The settings for encrypting data in transit.
//
// TODO: EXAMPLE
//
type CfnCluster_EncryptionInTransitProperty struct {
	// Indicates the encryption setting for data in transit between clients and brokers. The following are the possible values.
	//
	// - `TLS` means that client-broker communication is enabled with TLS only.
	// - `TLS_PLAINTEXT` means that client-broker communication is enabled for both TLS-encrypted, as well as plain text data.
	// - `PLAINTEXT` means that client-broker communication is enabled in plain text only.
	//
	// The default value is `TLS` .
	ClientBroker *string `json:"clientBroker" yaml:"clientBroker"`
	// When set to true, it indicates that data communication among the broker nodes of the cluster is encrypted.
	//
	// When set to false, the communication happens in plain text. The default value is true.
	InCluster interface{} `json:"inCluster" yaml:"inCluster"`
}

// Includes encryption-related information, such as the Amazon KMS key used for encrypting data at rest and whether you want MSK to encrypt your data in transit.
//
// TODO: EXAMPLE
//
type CfnCluster_EncryptionInfoProperty struct {
	// The data-volume encryption details.
	EncryptionAtRest interface{} `json:"encryptionAtRest" yaml:"encryptionAtRest"`
	// The details for encryption in transit.
	EncryptionInTransit interface{} `json:"encryptionInTransit" yaml:"encryptionInTransit"`
}

// Details of the Kinesis Data Firehose delivery stream that is the destination for broker logs.
//
// TODO: EXAMPLE
//
type CfnCluster_FirehoseProperty struct {
	// Specifies whether broker logs get sent to the specified Kinesis Data Firehose delivery stream.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// The Kinesis Data Firehose delivery stream that is the destination for broker logs.
	DeliveryStream *string `json:"deliveryStream" yaml:"deliveryStream"`
}

// Details for IAM access control.
//
// TODO: EXAMPLE
//
type CfnCluster_IamProperty struct {
	// Whether IAM access control is enabled.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
}

// Indicates whether you want to enable or disable the JMX Exporter.
//
// TODO: EXAMPLE
//
type CfnCluster_JmxExporterProperty struct {
	// Indicates whether you want to enable or disable the JMX Exporter.
	EnabledInBroker interface{} `json:"enabledInBroker" yaml:"enabledInBroker"`
}

// You can configure your Amazon MSK cluster to send broker logs to different destination types.
//
// This is a container for the configuration details related to broker logs.
//
// TODO: EXAMPLE
//
type CfnCluster_LoggingInfoProperty struct {
	// You can configure your Amazon MSK cluster to send broker logs to different destination types.
	//
	// This configuration specifies the details of these destinations.
	BrokerLogs interface{} `json:"brokerLogs" yaml:"brokerLogs"`
}

// Indicates whether you want to enable or disable the Node Exporter.
//
// TODO: EXAMPLE
//
type CfnCluster_NodeExporterProperty struct {
	// Indicates whether you want to enable or disable the Node Exporter.
	EnabledInBroker interface{} `json:"enabledInBroker" yaml:"enabledInBroker"`
}

// JMX and Node monitoring for the MSK cluster.
//
// TODO: EXAMPLE
//
type CfnCluster_OpenMonitoringProperty struct {
	// Prometheus exporter settings.
	Prometheus interface{} `json:"prometheus" yaml:"prometheus"`
}

// Prometheus settings for open monitoring.
//
// TODO: EXAMPLE
//
type CfnCluster_PrometheusProperty struct {
	// Indicates whether you want to enable or disable the JMX Exporter.
	JmxExporter interface{} `json:"jmxExporter" yaml:"jmxExporter"`
	// Indicates whether you want to enable or disable the Node Exporter.
	NodeExporter interface{} `json:"nodeExporter" yaml:"nodeExporter"`
}

// TODO: EXAMPLE
//
type CfnCluster_ProvisionedThroughputProperty struct {
	// `CfnCluster.ProvisionedThroughputProperty.Enabled`.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// `CfnCluster.ProvisionedThroughputProperty.VolumeThroughput`.
	VolumeThroughput *float64 `json:"volumeThroughput" yaml:"volumeThroughput"`
}

// Specifies whether the cluster's brokers are accessible from the internet.
//
// Public access is off by default.
//
// TODO: EXAMPLE
//
type CfnCluster_PublicAccessProperty struct {
	// Set to `DISABLED` to turn off public access or to `SERVICE_PROVIDED_EIPS` to turn it on.
	//
	// Public access if off by default.
	Type *string `json:"type" yaml:"type"`
}

// The details of the Amazon S3 destination for broker logs.
//
// TODO: EXAMPLE
//
type CfnCluster_S3Property struct {
	// Specifies whether broker logs get sent to the specified Amazon S3 destination.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// The name of the S3 bucket that is the destination for broker logs.
	Bucket *string `json:"bucket" yaml:"bucket"`
	// The S3 prefix that is the destination for broker logs.
	Prefix *string `json:"prefix" yaml:"prefix"`
}

// Details for client authentication using SASL.
//
// To turn on SASL, you must also turn on `EncryptionInTransit` by setting `inCluster` to true. You must set `clientBroker` to either `TLS` or `TLS_PLAINTEXT` . If you choose `TLS_PLAINTEXT` , then you must also set `unauthenticated` to true.
//
// TODO: EXAMPLE
//
type CfnCluster_SaslProperty struct {
	// Details for IAM access control.
	Iam interface{} `json:"iam" yaml:"iam"`
	// Details for SASL/SCRAM client authentication.
	Scram interface{} `json:"scram" yaml:"scram"`
}

// Details for SASL/SCRAM client authentication.
//
// TODO: EXAMPLE
//
type CfnCluster_ScramProperty struct {
	// SASL/SCRAM authentication is enabled or not.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
}

// Contains information about storage volumes attached to MSK broker nodes.
//
// TODO: EXAMPLE
//
type CfnCluster_StorageInfoProperty struct {
	// EBS volume information.
	EbsStorageInfo interface{} `json:"ebsStorageInfo" yaml:"ebsStorageInfo"`
}

// Details for client authentication using TLS.
//
// TODO: EXAMPLE
//
type CfnCluster_TlsProperty struct {
	// List of ACM Certificate Authority ARNs.
	CertificateAuthorityArnList *[]*string `json:"certificateAuthorityArnList" yaml:"certificateAuthorityArnList"`
	// TLS authentication is enabled or not.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
}

// Details for allowing no client authentication.
//
// TODO: EXAMPLE
//
type CfnCluster_UnauthenticatedProperty struct {
	// Unauthenticated is enabled or not.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
}

// Properties for defining a `CfnCluster`.
//
// TODO: EXAMPLE
//
type CfnClusterProps struct {
	// The setup to be used for brokers in the cluster.
	BrokerNodeGroupInfo interface{} `json:"brokerNodeGroupInfo" yaml:"brokerNodeGroupInfo"`
	// The name of the cluster.
	ClusterName *string `json:"clusterName" yaml:"clusterName"`
	// The version of Apache Kafka.
	//
	// For more information, see [Supported Apache Kafka versions](https://docs.aws.amazon.com/msk/latest/developerguide/supported-kafka-versions.html) in the Amazon MSK Developer Guide.
	KafkaVersion *string `json:"kafkaVersion" yaml:"kafkaVersion"`
	// The number of broker nodes you want in the Amazon MSK cluster.
	//
	// You can submit an update to increase the number of broker nodes in a cluster.
	NumberOfBrokerNodes *float64 `json:"numberOfBrokerNodes" yaml:"numberOfBrokerNodes"`
	// Includes information related to client authentication.
	ClientAuthentication interface{} `json:"clientAuthentication" yaml:"clientAuthentication"`
	// The Amazon MSK configuration to use for the cluster.
	ConfigurationInfo interface{} `json:"configurationInfo" yaml:"configurationInfo"`
	// Includes all encryption-related information.
	EncryptionInfo interface{} `json:"encryptionInfo" yaml:"encryptionInfo"`
	// Specifies the level of monitoring for the MSK cluster.
	//
	// The possible values are `DEFAULT` , `PER_BROKER` , and `PER_TOPIC_PER_BROKER` .
	EnhancedMonitoring *string `json:"enhancedMonitoring" yaml:"enhancedMonitoring"`
	// You can configure your Amazon MSK cluster to send broker logs to different destination types.
	//
	// This is a container for the configuration details related to broker logs.
	LoggingInfo interface{} `json:"loggingInfo" yaml:"loggingInfo"`
	// The settings for open monitoring.
	OpenMonitoring interface{} `json:"openMonitoring" yaml:"openMonitoring"`
	// A map of key:value pairs to apply to this resource.
	//
	// Both key and value are of type String.
	Tags interface{} `json:"tags" yaml:"tags"`
}


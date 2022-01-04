package awsmsk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsacmpca"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/awsmsk/internal"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/constructs-go/constructs/v3"
)

// Configuration details related to broker logs.
//
// TODO: EXAMPLE
//
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
//
// The `AWS::MSK::Cluster` resource creates an Amazon MSK cluster. For more information, see [What Is Amazon MSK?](https://docs.aws.amazon.com/msk/latest/developerguide/what-is-msk.html) in the *Amazon MSK Developer Guide* .
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
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// You can configure your MSK cluster to send broker logs to different destination types.
//
// This configuration specifies the details of these destinations.
//
// TODO: EXAMPLE
//
type CfnCluster_BrokerLogsProperty struct {
	// Details of the CloudWatch Logs destination for broker logs.
	CloudWatchLogs interface{} `json:"cloudWatchLogs"`
	// Details of the Kinesis Data Firehose delivery stream that is the destination for broker logs.
	Firehose interface{} `json:"firehose"`
	// Details of the Amazon S3 destination for broker logs.
	S3 interface{} `json:"s3"`
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
	ClientSubnets *[]*string `json:"clientSubnets"`
	// The type of Amazon EC2 instances to use for brokers.
	//
	// The following instance types are allowed: kafka.m5.large, kafka.m5.xlarge, kafka.m5.2xlarge, kafka.m5.4xlarge, kafka.m5.8xlarge, kafka.m5.12xlarge, kafka.m5.16xlarge, and kafka.m5.24xlarge.
	InstanceType *string `json:"instanceType"`
	// This parameter is currently not in use.
	BrokerAzDistribution *string `json:"brokerAzDistribution"`
	// Information about the cluster's connectivity setting.
	ConnectivityInfo interface{} `json:"connectivityInfo"`
	// The security groups to associate with the elastic network interfaces in order to specify who can connect to and communicate with the Amazon MSK cluster.
	//
	// If you don't specify a security group, Amazon MSK uses the default security group associated with the VPC. If you specify security groups that were shared with you, you must ensure that you have permissions to them. Specifically, you need the `ec2:DescribeSecurityGroups` permission.
	SecurityGroups *[]*string `json:"securityGroups"`
	// Contains information about storage volumes attached to MSK broker nodes.
	StorageInfo interface{} `json:"storageInfo"`
}

// Includes information related to client authentication.
//
// TODO: EXAMPLE
//
type CfnCluster_ClientAuthenticationProperty struct {
	// Details for ClientAuthentication using SASL.
	Sasl interface{} `json:"sasl"`
	// Details for client authentication using TLS.
	Tls interface{} `json:"tls"`
	// Details for ClientAuthentication using no authentication.
	Unauthenticated interface{} `json:"unauthenticated"`
}

// Details of the CloudWatch Logs destination for broker logs.
//
// TODO: EXAMPLE
//
type CfnCluster_CloudWatchLogsProperty struct {
	// Specifies whether broker logs get sent to the specified CloudWatch Logs destination.
	Enabled interface{} `json:"enabled"`
	// The CloudWatch log group that is the destination for broker logs.
	LogGroup *string `json:"logGroup"`
}

// Specifies the Amazon MSK configuration to use for the brokers.
//
// TODO: EXAMPLE
//
type CfnCluster_ConfigurationInfoProperty struct {
	// The Amazon Resource Name (ARN) of the MSK configuration to use.
	//
	// For example, `arn:aws:kafka:us-east-1:123456789012:configuration/example-configuration-name/abcdabcd-1234-abcd-1234-abcd123e8e8e-1` .
	Arn *string `json:"arn"`
	// The revision of the Amazon MSK configuration to use.
	Revision *float64 `json:"revision"`
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
	PublicAccess interface{} `json:"publicAccess"`
}

// Contains information about the EBS storage volumes attached to brokers.
//
// TODO: EXAMPLE
//
type CfnCluster_EBSStorageInfoProperty struct {
	// The size in GiB of the EBS volume for the data drive on each broker node.
	VolumeSize *float64 `json:"volumeSize"`
}

// The data volume encryption details.
//
// TODO: EXAMPLE
//
type CfnCluster_EncryptionAtRestProperty struct {
	// The ARN of the Amazon KMS key for encrypting data at rest.
	//
	// If you don't specify a KMS key, MSK creates one for you and uses it on your behalf.
	DataVolumeKmsKeyId *string `json:"dataVolumeKmsKeyId"`
}

// The settings for encrypting data in transit.
//
// TODO: EXAMPLE
//
type CfnCluster_EncryptionInTransitProperty struct {
	// Indicates the encryption setting for data in transit between clients and brokers. The following are the possible values.
	//
	// - `TLS` means that client-broker communication is enabled with TLS only.
	// - `TLS_PLAINTEXT` means that client-broker communication is enabled for both TLS-encrypted, as well as plaintext data.
	// - `PLAINTEXT` means that client-broker communication is enabled in plaintext only.
	//
	// The default value is `TLS` .
	ClientBroker *string `json:"clientBroker"`
	// When set to true, it indicates that data communication among the broker nodes of the cluster is encrypted.
	//
	// When set to false, the communication happens in plaintext. The default value is true.
	InCluster interface{} `json:"inCluster"`
}

// Includes encryption-related information, such as the Amazon KMS key used for encrypting data at rest and whether you want MSK to encrypt your data in transit.
//
// TODO: EXAMPLE
//
type CfnCluster_EncryptionInfoProperty struct {
	// The data-volume encryption details.
	EncryptionAtRest interface{} `json:"encryptionAtRest"`
	// The details for encryption in transit.
	EncryptionInTransit interface{} `json:"encryptionInTransit"`
}

// Details of the Kinesis Data Firehose delivery stream that is the destination for broker logs.
//
// TODO: EXAMPLE
//
type CfnCluster_FirehoseProperty struct {
	// Specifies whether broker logs get sent to the specified Kinesis Data Firehose delivery stream.
	Enabled interface{} `json:"enabled"`
	// The Kinesis Data Firehose delivery stream that is the destination for broker logs.
	DeliveryStream *string `json:"deliveryStream"`
}

// Details for IAM access control.
//
// TODO: EXAMPLE
//
type CfnCluster_IamProperty struct {
	// Whether IAM access control is enabled.
	Enabled interface{} `json:"enabled"`
}

// Indicates whether you want to enable or disable the JMX Exporter.
//
// TODO: EXAMPLE
//
type CfnCluster_JmxExporterProperty struct {
	// Indicates whether you want to enable or disable the JMX Exporter.
	EnabledInBroker interface{} `json:"enabledInBroker"`
}

// You can configure your MSK cluster to send broker logs to different destination types.
//
// This is a container for the configuration details related to broker logs.
//
// TODO: EXAMPLE
//
type CfnCluster_LoggingInfoProperty struct {
	// You can configure your MSK cluster to send broker logs to different destination types.
	//
	// This configuration specifies the details of these destinations.
	BrokerLogs interface{} `json:"brokerLogs"`
}

// Indicates whether you want to enable or disable the Node Exporter.
//
// TODO: EXAMPLE
//
type CfnCluster_NodeExporterProperty struct {
	// Indicates whether you want to enable or disable the Node Exporter.
	EnabledInBroker interface{} `json:"enabledInBroker"`
}

// JMX and Node monitoring for the MSK cluster.
//
// TODO: EXAMPLE
//
type CfnCluster_OpenMonitoringProperty struct {
	// Prometheus exporter settings.
	Prometheus interface{} `json:"prometheus"`
}

// Prometheus settings for open monitoring.
//
// TODO: EXAMPLE
//
type CfnCluster_PrometheusProperty struct {
	// Indicates whether you want to enable or disable the JMX Exporter.
	JmxExporter interface{} `json:"jmxExporter"`
	// Indicates whether you want to enable or disable the Node Exporter.
	NodeExporter interface{} `json:"nodeExporter"`
}

// Specifies whether the cluster's brokers are accessible from the internet.
//
// Public access is off by default.
//
// TODO: EXAMPLE
//
type CfnCluster_PublicAccessProperty struct {
	// Set to DISABLED to turn off public access or to SERVICE_PROVIDED_EIPS to turn it on.
	//
	// Public access if off by default.
	Type *string `json:"type"`
}

// The details of the Amazon S3 destination for broker logs.
//
// TODO: EXAMPLE
//
type CfnCluster_S3Property struct {
	// Specifies whether broker logs get sent to the specified Amazon S3 destination.
	Enabled interface{} `json:"enabled"`
	// The name of the S3 bucket that is the destination for broker logs.
	Bucket *string `json:"bucket"`
	// The S3 prefix that is the destination for broker logs.
	Prefix *string `json:"prefix"`
}

// Details for client authentication using SASL.
//
// To turn on SASL, you must also turn on `EncryptionInTransit` by setting `inCluster` to true. You must set `clientBroker` to either `TLS` or `TLS_PLAINTEXT` . If you choose `TLS_PLAINTEXT` , then you must also set `unauthenticated` to true.
//
// TODO: EXAMPLE
//
type CfnCluster_SaslProperty struct {
	// Details for IAM access control.
	Iam interface{} `json:"iam"`
	// Details for SASL/SCRAM client authentication.
	Scram interface{} `json:"scram"`
}

// Details for SASL/SCRAM client authentication.
//
// TODO: EXAMPLE
//
type CfnCluster_ScramProperty struct {
	// SASL/SCRAM authentication is enabled or not.
	Enabled interface{} `json:"enabled"`
}

// Contains information about storage volumes attached to MSK broker nodes.
//
// TODO: EXAMPLE
//
type CfnCluster_StorageInfoProperty struct {
	// EBS volume information.
	EbsStorageInfo interface{} `json:"ebsStorageInfo"`
}

// Details for client authentication using TLS.
//
// TODO: EXAMPLE
//
type CfnCluster_TlsProperty struct {
	// List of ACM Certificate Authority ARNs.
	CertificateAuthorityArnList *[]*string `json:"certificateAuthorityArnList"`
	// TLS authentication is enabled or not.
	Enabled interface{} `json:"enabled"`
}

// Details for allowing no client authentication.
//
// TODO: EXAMPLE
//
type CfnCluster_UnauthenticatedProperty struct {
	// Unauthenticated is enabled or not.
	Enabled interface{} `json:"enabled"`
}

// Properties for defining a `CfnCluster`.
//
// TODO: EXAMPLE
//
type CfnClusterProps struct {
	// The setup to be used for brokers in the cluster.
	BrokerNodeGroupInfo interface{} `json:"brokerNodeGroupInfo"`
	// The name of the cluster.
	ClusterName *string `json:"clusterName"`
	// The version of Apache Kafka.
	//
	// You can use Amazon MSK to create clusters that use Apache Kafka versions 1.1.1 and 2.2.1.
	KafkaVersion *string `json:"kafkaVersion"`
	// The number of broker nodes you want in the Amazon MSK cluster.
	//
	// You can submit an update to increase the number of broker nodes in a cluster.
	NumberOfBrokerNodes *float64 `json:"numberOfBrokerNodes"`
	// Includes information related to client authentication.
	ClientAuthentication interface{} `json:"clientAuthentication"`
	// The Amazon MSK configuration to use for the cluster.
	ConfigurationInfo interface{} `json:"configurationInfo"`
	// Includes all encryption-related information.
	EncryptionInfo interface{} `json:"encryptionInfo"`
	// Specifies the level of monitoring for the MSK cluster.
	//
	// The possible values are `DEFAULT` , `PER_BROKER` , and `PER_TOPIC_PER_BROKER` .
	EnhancedMonitoring *string `json:"enhancedMonitoring"`
	// You can configure your MSK cluster to send broker logs to different destination types.
	//
	// This is a container for the configuration details related to broker logs.
	LoggingInfo interface{} `json:"loggingInfo"`
	// The settings for open monitoring.
	OpenMonitoring interface{} `json:"openMonitoring"`
	// A map of key:value pairs to apply to this resource.
	//
	// Both key and value are of type String.
	Tags interface{} `json:"tags"`
}

// Configuration properties for client authentication.
//
// TODO: EXAMPLE
//
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
//
// TODO: EXAMPLE
//
// Experimental.
type ClientBrokerEncryption string

const (
	ClientBrokerEncryption_TLS ClientBrokerEncryption = "TLS"
	ClientBrokerEncryption_TLS_PLAINTEXT ClientBrokerEncryption = "TLS_PLAINTEXT"
	ClientBrokerEncryption_PLAINTEXT ClientBrokerEncryption = "PLAINTEXT"
)

// Create a MSK Cluster.
//
// TODO: EXAMPLE
//
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
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
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
//
// TODO: EXAMPLE
//
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
//
// TODO: EXAMPLE
//
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
//
// TODO: EXAMPLE
//
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
//
// TODO: EXAMPLE
//
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

func (i *jsiiProxy_ICluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
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
//
// TODO: EXAMPLE
//
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

func KafkaVersion_V2_6_2() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_6_2",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_6_3() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_6_3",
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

func KafkaVersion_V2_7_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_7_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_7_2() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_7_2",
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

func KafkaVersion_V2_8_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"monocdk.aws_msk.KafkaVersion",
		"V2_8_1",
		&returns,
	)
	return returns
}

// Monitoring Configuration.
//
// TODO: EXAMPLE
//
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
//
// TODO: EXAMPLE
//
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
//
// TODO: EXAMPLE
//
// Experimental.
type SaslAuthProps struct {
	// Enable IAM access control.
	// Experimental.
	Iam *bool `json:"iam"`
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
//
// TODO: EXAMPLE
//
// Experimental.
type TlsAuthProps struct {
	// List of ACM Certificate Authorities to enable TLS authentication.
	// Experimental.
	CertificateAuthorities *[]awsacmpca.ICertificateAuthority `json:"certificateAuthorities"`
}


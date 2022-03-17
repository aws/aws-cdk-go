package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/awskinesisfirehose/internal"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::KinesisFirehose::DeliveryStream`.
//
// The `AWS::KinesisFirehose::DeliveryStream` resource specifies an Amazon Kinesis Data Firehose (Kinesis Data Firehose) delivery stream that delivers real-time streaming data to an Amazon Simple Storage Service (Amazon S3), Amazon Redshift, or Amazon Elasticsearch Service (Amazon ES) destination. For more information, see [Creating an Amazon Kinesis Data Firehose Delivery Stream](https://docs.aws.amazon.com/firehose/latest/dev/basic-create.html) in the *Amazon Kinesis Data Firehose Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnDeliveryStream interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AmazonopensearchserviceDestinationConfiguration() interface{}
	SetAmazonopensearchserviceDestinationConfiguration(val interface{})
	AttrArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DeliveryStreamEncryptionConfigurationInput() interface{}
	SetDeliveryStreamEncryptionConfigurationInput(val interface{})
	DeliveryStreamName() *string
	SetDeliveryStreamName(val *string)
	DeliveryStreamType() *string
	SetDeliveryStreamType(val *string)
	ElasticsearchDestinationConfiguration() interface{}
	SetElasticsearchDestinationConfiguration(val interface{})
	ExtendedS3DestinationConfiguration() interface{}
	SetExtendedS3DestinationConfiguration(val interface{})
	HttpEndpointDestinationConfiguration() interface{}
	SetHttpEndpointDestinationConfiguration(val interface{})
	KinesisStreamSourceConfiguration() interface{}
	SetKinesisStreamSourceConfiguration(val interface{})
	LogicalId() *string
	Node() awscdk.ConstructNode
	RedshiftDestinationConfiguration() interface{}
	SetRedshiftDestinationConfiguration(val interface{})
	Ref() *string
	S3DestinationConfiguration() interface{}
	SetS3DestinationConfiguration(val interface{})
	SplunkDestinationConfiguration() interface{}
	SetSplunkDestinationConfiguration(val interface{})
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

// The jsii proxy struct for CfnDeliveryStream
type jsiiProxy_CfnDeliveryStream struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDeliveryStream) AmazonopensearchserviceDestinationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonopensearchserviceDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) DeliveryStreamEncryptionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deliveryStreamEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) DeliveryStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) DeliveryStreamType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) ElasticsearchDestinationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) ExtendedS3DestinationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extendedS3DestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) HttpEndpointDestinationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpEndpointDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) KinesisStreamSourceConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisStreamSourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) RedshiftDestinationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redshiftDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) S3DestinationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3DestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) SplunkDestinationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splunkDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::KinesisFirehose::DeliveryStream`.
func NewCfnDeliveryStream(scope awscdk.Construct, id *string, props *CfnDeliveryStreamProps) CfnDeliveryStream {
	_init_.Initialize()

	j := jsiiProxy_CfnDeliveryStream{}

	_jsii_.Create(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::KinesisFirehose::DeliveryStream`.
func NewCfnDeliveryStream_Override(c CfnDeliveryStream, scope awscdk.Construct, id *string, props *CfnDeliveryStreamProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDeliveryStream) SetAmazonopensearchserviceDestinationConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"amazonopensearchserviceDestinationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream) SetDeliveryStreamEncryptionConfigurationInput(val interface{}) {
	_jsii_.Set(
		j,
		"deliveryStreamEncryptionConfigurationInput",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream) SetDeliveryStreamName(val *string) {
	_jsii_.Set(
		j,
		"deliveryStreamName",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream) SetDeliveryStreamType(val *string) {
	_jsii_.Set(
		j,
		"deliveryStreamType",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream) SetElasticsearchDestinationConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"elasticsearchDestinationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream) SetExtendedS3DestinationConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"extendedS3DestinationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream) SetHttpEndpointDestinationConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"httpEndpointDestinationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream) SetKinesisStreamSourceConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"kinesisStreamSourceConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream) SetRedshiftDestinationConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"redshiftDestinationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream) SetS3DestinationConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"s3DestinationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream) SetSplunkDestinationConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"splunkDestinationConfiguration",
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
func CfnDeliveryStream_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDeliveryStream_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDeliveryStream_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDeliveryStream_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnDeliveryStream) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnDeliveryStream) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnDeliveryStream) AddMetadata(key *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnDeliveryStream) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnDeliveryStream) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnDeliveryStream) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnDeliveryStream) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnDeliveryStream) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnDeliveryStream) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnDeliveryStream) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnDeliveryStream) OnPrepare() {
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
func (c *jsiiProxy_CfnDeliveryStream) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnDeliveryStream) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnDeliveryStream) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnDeliveryStream) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDeliveryStream) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnDeliveryStream) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnDeliveryStream) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnDeliveryStream) ToString() *string {
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
func (c *jsiiProxy_CfnDeliveryStream) Validate() *[]*string {
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
func (c *jsiiProxy_CfnDeliveryStream) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Describes the buffering to perform before delivering data to the Amazon OpenSearch Service destination.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_AmazonopensearchserviceBufferingHintsProperty struct {
	// Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination.
	//
	// The default value is 300 (5 minutes).
	IntervalInSeconds *float64 `json:"intervalInSeconds" yaml:"intervalInSeconds"`
	// Buffer incoming data to the specified size, in MBs, before delivering it to the destination.
	//
	// The default value is 5. We recommend setting this parameter to a value greater than the amount of data you typically ingest into the delivery stream in 10 seconds. For example, if you typically ingest data at 1 MB/sec, the value should be 10 MB or higher.
	SizeInMBs *float64 `json:"sizeInMBs" yaml:"sizeInMBs"`
}

// Describes the configuration of a destination in Amazon OpenSearch Service.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_AmazonopensearchserviceDestinationConfigurationProperty struct {
	// The Amazon OpenSearch Service index name.
	IndexName *string `json:"indexName" yaml:"indexName"`
	// The Amazon Resource Name (ARN) of the IAM role to be assumed by Kinesis Data Firehose for calling the Amazon OpenSearch Service Configuration API and for indexing documents.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// Describes the configuration of a destination in Amazon S3.
	S3Configuration interface{} `json:"s3Configuration" yaml:"s3Configuration"`
	// The buffering options.
	//
	// If no value is specified, the default values for AmazonopensearchserviceBufferingHints are used.
	BufferingHints interface{} `json:"bufferingHints" yaml:"bufferingHints"`
	// Describes the Amazon CloudWatch logging options for your delivery stream.
	CloudWatchLoggingOptions interface{} `json:"cloudWatchLoggingOptions" yaml:"cloudWatchLoggingOptions"`
	// The endpoint to use when communicating with the cluster.
	//
	// Specify either this ClusterEndpoint or the DomainARN field.
	ClusterEndpoint *string `json:"clusterEndpoint" yaml:"clusterEndpoint"`
	// The ARN of the Amazon OpenSearch Service domain.
	DomainArn *string `json:"domainArn" yaml:"domainArn"`
	// The Amazon OpenSearch Service index rotation period.
	//
	// Index rotation appends a timestamp to the IndexName to facilitate the expiration of old data.
	IndexRotationPeriod *string `json:"indexRotationPeriod" yaml:"indexRotationPeriod"`
	// Describes a data processing configuration.
	ProcessingConfiguration interface{} `json:"processingConfiguration" yaml:"processingConfiguration"`
	// The retry behavior in case Kinesis Data Firehose is unable to deliver documents to Amazon OpenSearch Service.
	//
	// The default value is 300 (5 minutes).
	RetryOptions interface{} `json:"retryOptions" yaml:"retryOptions"`
	// Defines how documents should be delivered to Amazon S3.
	S3BackupMode *string `json:"s3BackupMode" yaml:"s3BackupMode"`
	// The Amazon OpenSearch Service type name.
	TypeName *string `json:"typeName" yaml:"typeName"`
	// The details of the VPC of the Amazon OpenSearch Service destination.
	VpcConfiguration interface{} `json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

// Configures retry behavior in case Kinesis Data Firehose is unable to deliver documents to Amazon OpenSearch Service.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_AmazonopensearchserviceRetryOptionsProperty struct {
	// After an initial failure to deliver to Amazon OpenSearch Service, the total amount of time during which Kinesis Data Firehose retries delivery (including the first attempt).
	//
	// After this time has elapsed, the failed documents are written to Amazon S3. Default value is 300 seconds (5 minutes). A value of 0 (zero) results in no retries.
	DurationInSeconds *float64 `json:"durationInSeconds" yaml:"durationInSeconds"`
}

// The `BufferingHints` property type specifies how Amazon Kinesis Data Firehose (Kinesis Data Firehose) buffers incoming data before delivering it to the destination.
//
// The first buffer condition that is satisfied triggers Kinesis Data Firehose to deliver the data.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_BufferingHintsProperty struct {
	// The length of time, in seconds, that Kinesis Data Firehose buffers incoming data before delivering it to the destination.
	//
	// For valid values, see the `IntervalInSeconds` content for the [BufferingHints](https://docs.aws.amazon.com/firehose/latest/APIReference/API_BufferingHints.html) data type in the *Amazon Kinesis Data Firehose API Reference* .
	IntervalInSeconds *float64 `json:"intervalInSeconds" yaml:"intervalInSeconds"`
	// The size of the buffer, in MBs, that Kinesis Data Firehose uses for incoming data before delivering it to the destination.
	//
	// For valid values, see the `SizeInMBs` content for the [BufferingHints](https://docs.aws.amazon.com/firehose/latest/APIReference/API_BufferingHints.html) data type in the *Amazon Kinesis Data Firehose API Reference* .
	SizeInMBs *float64 `json:"sizeInMBs" yaml:"sizeInMBs"`
}

// The `CloudWatchLoggingOptions` property type specifies Amazon CloudWatch Logs (CloudWatch Logs) logging options that Amazon Kinesis Data Firehose (Kinesis Data Firehose) uses for the delivery stream.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_CloudWatchLoggingOptionsProperty struct {
	// Indicates whether CloudWatch Logs logging is enabled.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// The name of the CloudWatch Logs log group that contains the log stream that Kinesis Data Firehose will use.
	//
	// Conditional. If you enable logging, you must specify this property.
	LogGroupName *string `json:"logGroupName" yaml:"logGroupName"`
	// The name of the CloudWatch Logs log stream that Kinesis Data Firehose uses to send logs about data delivery.
	//
	// Conditional. If you enable logging, you must specify this property.
	LogStreamName *string `json:"logStreamName" yaml:"logStreamName"`
}

// The `CopyCommand` property type configures the Amazon Redshift `COPY` command that Amazon Kinesis Data Firehose (Kinesis Data Firehose) uses to load data into an Amazon Redshift cluster from an Amazon S3 bucket.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_CopyCommandProperty struct {
	// The name of the target table.
	//
	// The table must already exist in the database.
	DataTableName *string `json:"dataTableName" yaml:"dataTableName"`
	// Parameters to use with the Amazon Redshift `COPY` command.
	//
	// For examples, see the `CopyOptions` content for the [CopyCommand](https://docs.aws.amazon.com/firehose/latest/APIReference/API_CopyCommand.html) data type in the *Amazon Kinesis Data Firehose API Reference* .
	CopyOptions *string `json:"copyOptions" yaml:"copyOptions"`
	// A comma-separated list of column names.
	DataTableColumns *string `json:"dataTableColumns" yaml:"dataTableColumns"`
}

// Specifies that you want Kinesis Data Firehose to convert data from the JSON format to the Parquet or ORC format before writing it to Amazon S3.
//
// Kinesis Data Firehose uses the serializer and deserializer that you specify, in addition to the column information from the AWS Glue table, to deserialize your input data from JSON and then serialize it to the Parquet or ORC format. For more information, see [Kinesis Data Firehose Record Format Conversion](https://docs.aws.amazon.com/firehose/latest/dev/record-format-conversion.html) .
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_DataFormatConversionConfigurationProperty struct {
	// Defaults to `true` .
	//
	// Set it to `false` if you want to disable format conversion while preserving the configuration details.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// Specifies the deserializer that you want Kinesis Data Firehose to use to convert the format of your data from JSON.
	//
	// This parameter is required if `Enabled` is set to true.
	InputFormatConfiguration interface{} `json:"inputFormatConfiguration" yaml:"inputFormatConfiguration"`
	// Specifies the serializer that you want Kinesis Data Firehose to use to convert the format of your data to the Parquet or ORC format.
	//
	// This parameter is required if `Enabled` is set to true.
	OutputFormatConfiguration interface{} `json:"outputFormatConfiguration" yaml:"outputFormatConfiguration"`
	// Specifies the AWS Glue Data Catalog table that contains the column information.
	//
	// This parameter is required if `Enabled` is set to true.
	SchemaConfiguration interface{} `json:"schemaConfiguration" yaml:"schemaConfiguration"`
}

// Specifies the type and Amazon Resource Name (ARN) of the CMK to use for Server-Side Encryption (SSE).
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_DeliveryStreamEncryptionConfigurationInputProperty struct {
	// Indicates the type of customer master key (CMK) to use for encryption.
	//
	// The default setting is `AWS_OWNED_CMK` . For more information about CMKs, see [Customer Master Keys (CMKs)](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#master_keys) .
	//
	// You can use a CMK of type CUSTOMER_MANAGED_CMK to encrypt up to 500 delivery streams.
	//
	// > To encrypt your delivery stream, use symmetric CMKs. Kinesis Data Firehose doesn't support asymmetric CMKs. For information about symmetric and asymmetric CMKs, see [About Symmetric and Asymmetric CMKs](https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-concepts.html) in the AWS Key Management Service developer guide.
	KeyType *string `json:"keyType" yaml:"keyType"`
	// If you set `KeyType` to `CUSTOMER_MANAGED_CMK` , you must specify the Amazon Resource Name (ARN) of the CMK.
	//
	// If you set `KeyType` to `AWS _OWNED_CMK` , Kinesis Data Firehose uses a service-account CMK.
	KeyArn *string `json:"keyArn" yaml:"keyArn"`
}

// The deserializer you want Kinesis Data Firehose to use for converting the input data from JSON.
//
// Kinesis Data Firehose then serializes the data to its final format using the `Serializer` . Kinesis Data Firehose supports two types of deserializers: the [Apache Hive JSON SerDe](https://docs.aws.amazon.com/https://cwiki.apache.org/confluence/display/Hive/LanguageManual+DDL#LanguageManualDDL-JSON) and the [OpenX JSON SerDe](https://docs.aws.amazon.com/https://github.com/rcongiu/Hive-JSON-Serde) .
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_DeserializerProperty struct {
	// The native Hive / HCatalog JsonSerDe.
	//
	// Used by Kinesis Data Firehose for deserializing data, which means converting it from the JSON format in preparation for serializing it to the Parquet or ORC format. This is one of two deserializers you can choose, depending on which one offers the functionality you need. The other option is the OpenX SerDe.
	HiveJsonSerDe interface{} `json:"hiveJsonSerDe" yaml:"hiveJsonSerDe"`
	// The OpenX SerDe.
	//
	// Used by Kinesis Data Firehose for deserializing data, which means converting it from the JSON format in preparation for serializing it to the Parquet or ORC format. This is one of two deserializers you can choose, depending on which one offers the functionality you need. The other option is the native Hive / HCatalog JsonSerDe.
	OpenXJsonSerDe interface{} `json:"openXJsonSerDe" yaml:"openXJsonSerDe"`
}

// The `DynamicPartitioningConfiguration` property type specifies the configuration of the dynamic partitioning mechanism that creates targeted data sets from the streaming data by partitioning it based on partition keys.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_DynamicPartitioningConfigurationProperty struct {
	// Specifies whether dynamic partitioning is enabled for this Kinesis Data Firehose delivery stream.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// Specifies the retry behavior in case Kinesis Data Firehose is unable to deliver data to an Amazon S3 prefix.
	RetryOptions interface{} `json:"retryOptions" yaml:"retryOptions"`
}

// The `ElasticsearchBufferingHints` property type specifies how Amazon Kinesis Data Firehose (Kinesis Data Firehose) buffers incoming data while delivering it to the destination.
//
// The first buffer condition that is satisfied triggers Kinesis Data Firehose to deliver the data.
//
// ElasticsearchBufferingHints is the property type for the `BufferingHints` property of the [Amazon Kinesis Data Firehose DeliveryStream ElasticsearchDestinationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-elasticsearchdestinationconfiguration.html) property type.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_ElasticsearchBufferingHintsProperty struct {
	// The length of time, in seconds, that Kinesis Data Firehose buffers incoming data before delivering it to the destination.
	//
	// For valid values, see the `IntervalInSeconds` content for the [BufferingHints](https://docs.aws.amazon.com/firehose/latest/APIReference/API_BufferingHints.html) data type in the *Amazon Kinesis Data Firehose API Reference* .
	IntervalInSeconds *float64 `json:"intervalInSeconds" yaml:"intervalInSeconds"`
	// The size of the buffer, in MBs, that Kinesis Data Firehose uses for incoming data before delivering it to the destination.
	//
	// For valid values, see the `SizeInMBs` content for the [BufferingHints](https://docs.aws.amazon.com/firehose/latest/APIReference/API_BufferingHints.html) data type in the *Amazon Kinesis Data Firehose API Reference* .
	SizeInMBs *float64 `json:"sizeInMBs" yaml:"sizeInMBs"`
}

// The `ElasticsearchDestinationConfiguration` property type specifies an Amazon Elasticsearch Service (Amazon ES) domain that Amazon Kinesis Data Firehose (Kinesis Data Firehose) delivers data to.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_ElasticsearchDestinationConfigurationProperty struct {
	// The name of the Elasticsearch index to which Kinesis Data Firehose adds data for indexing.
	IndexName *string `json:"indexName" yaml:"indexName"`
	// The Amazon Resource Name (ARN) of the IAM role to be assumed by Kinesis Data Firehose for calling the Amazon ES Configuration API and for indexing documents.
	//
	// For more information, see [Controlling Access with Amazon Kinesis Data Firehose](https://docs.aws.amazon.com/firehose/latest/dev/controlling-access.html) .
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The S3 bucket where Kinesis Data Firehose backs up incoming data.
	S3Configuration interface{} `json:"s3Configuration" yaml:"s3Configuration"`
	// Configures how Kinesis Data Firehose buffers incoming data while delivering it to the Amazon ES domain.
	BufferingHints interface{} `json:"bufferingHints" yaml:"bufferingHints"`
	// The Amazon CloudWatch Logs logging options for the delivery stream.
	CloudWatchLoggingOptions interface{} `json:"cloudWatchLoggingOptions" yaml:"cloudWatchLoggingOptions"`
	// The endpoint to use when communicating with the cluster.
	//
	// Specify either this `ClusterEndpoint` or the `DomainARN` field.
	ClusterEndpoint *string `json:"clusterEndpoint" yaml:"clusterEndpoint"`
	// The ARN of the Amazon ES domain.
	//
	// The IAM role must have permissions for `DescribeElasticsearchDomain` , `DescribeElasticsearchDomains` , and `DescribeElasticsearchDomainConfig` after assuming the role specified in *RoleARN* .
	//
	// Specify either `ClusterEndpoint` or `DomainARN` .
	DomainArn *string `json:"domainArn" yaml:"domainArn"`
	// The frequency of Elasticsearch index rotation.
	//
	// If you enable index rotation, Kinesis Data Firehose appends a portion of the UTC arrival timestamp to the specified index name, and rotates the appended timestamp accordingly. For more information, see [Index Rotation for the Amazon ES Destination](https://docs.aws.amazon.com/firehose/latest/dev/basic-deliver.html#es-index-rotation) in the *Amazon Kinesis Data Firehose Developer Guide* .
	IndexRotationPeriod *string `json:"indexRotationPeriod" yaml:"indexRotationPeriod"`
	// The data processing configuration for the Kinesis Data Firehose delivery stream.
	ProcessingConfiguration interface{} `json:"processingConfiguration" yaml:"processingConfiguration"`
	// The retry behavior when Kinesis Data Firehose is unable to deliver data to Amazon ES.
	RetryOptions interface{} `json:"retryOptions" yaml:"retryOptions"`
	// The condition under which Kinesis Data Firehose delivers data to Amazon Simple Storage Service (Amazon S3).
	//
	// You can send Amazon S3 all documents (all data) or only the documents that Kinesis Data Firehose could not deliver to the Amazon ES destination. For more information and valid values, see the `S3BackupMode` content for the [ElasticsearchDestinationConfiguration](https://docs.aws.amazon.com/firehose/latest/APIReference/API_ElasticsearchDestinationConfiguration.html) data type in the *Amazon Kinesis Data Firehose API Reference* .
	S3BackupMode *string `json:"s3BackupMode" yaml:"s3BackupMode"`
	// The Elasticsearch type name that Amazon ES adds to documents when indexing data.
	TypeName *string `json:"typeName" yaml:"typeName"`
	// The details of the VPC of the Amazon ES destination.
	VpcConfiguration interface{} `json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

// The `ElasticsearchRetryOptions` property type configures the retry behavior for when Amazon Kinesis Data Firehose (Kinesis Data Firehose) can't deliver data to Amazon Elasticsearch Service (Amazon ES).
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_ElasticsearchRetryOptionsProperty struct {
	// After an initial failure to deliver to Amazon ES, the total amount of time during which Kinesis Data Firehose re-attempts delivery (including the first attempt).
	//
	// If Kinesis Data Firehose can't deliver the data within the specified time, it writes the data to the backup S3 bucket. For valid values, see the `DurationInSeconds` content for the [ElasticsearchRetryOptions](https://docs.aws.amazon.com/firehose/latest/APIReference/API_ElasticsearchRetryOptions.html) data type in the *Amazon Kinesis Data Firehose API Reference* .
	DurationInSeconds *float64 `json:"durationInSeconds" yaml:"durationInSeconds"`
}

// The `EncryptionConfiguration` property type specifies the encryption settings that Amazon Kinesis Data Firehose (Kinesis Data Firehose) uses when delivering data to Amazon Simple Storage Service (Amazon S3).
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_EncryptionConfigurationProperty struct {
	// The AWS Key Management Service ( AWS KMS) encryption key that Amazon S3 uses to encrypt your data.
	KmsEncryptionConfig interface{} `json:"kmsEncryptionConfig" yaml:"kmsEncryptionConfig"`
	// Disables encryption.
	//
	// For valid values, see the `NoEncryptionConfig` content for the [EncryptionConfiguration](https://docs.aws.amazon.com/firehose/latest/APIReference/API_EncryptionConfiguration.html) data type in the *Amazon Kinesis Data Firehose API Reference* .
	NoEncryptionConfig *string `json:"noEncryptionConfig" yaml:"noEncryptionConfig"`
}

// The `ExtendedS3DestinationConfiguration` property type configures an Amazon S3 destination for an Amazon Kinesis Data Firehose delivery stream.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_ExtendedS3DestinationConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon S3 bucket.
	//
	// For constraints, see [ExtendedS3DestinationConfiguration](https://docs.aws.amazon.com/firehose/latest/APIReference/API_ExtendedS3DestinationConfiguration.html) in the *Amazon Kinesis Data Firehose API Reference* .
	BucketArn *string `json:"bucketArn" yaml:"bucketArn"`
	// The Amazon Resource Name (ARN) of the AWS credentials.
	//
	// For constraints, see [ExtendedS3DestinationConfiguration](https://docs.aws.amazon.com/firehose/latest/APIReference/API_ExtendedS3DestinationConfiguration.html) in the *Amazon Kinesis Data Firehose API Reference* .
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The buffering option.
	BufferingHints interface{} `json:"bufferingHints" yaml:"bufferingHints"`
	// The Amazon CloudWatch logging options for your delivery stream.
	CloudWatchLoggingOptions interface{} `json:"cloudWatchLoggingOptions" yaml:"cloudWatchLoggingOptions"`
	// The compression format.
	//
	// If no value is specified, the default is `UNCOMPRESSED` .
	CompressionFormat *string `json:"compressionFormat" yaml:"compressionFormat"`
	// The serializer, deserializer, and schema for converting data from the JSON format to the Parquet or ORC format before writing it to Amazon S3.
	DataFormatConversionConfiguration interface{} `json:"dataFormatConversionConfiguration" yaml:"dataFormatConversionConfiguration"`
	// The configuration of the dynamic partitioning mechanism that creates targeted data sets from the streaming data by partitioning it based on partition keys.
	DynamicPartitioningConfiguration interface{} `json:"dynamicPartitioningConfiguration" yaml:"dynamicPartitioningConfiguration"`
	// The encryption configuration for the Kinesis Data Firehose delivery stream.
	//
	// The default value is `NoEncryption` .
	EncryptionConfiguration interface{} `json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3.
	//
	// This prefix appears immediately following the bucket name. For information about how to specify this prefix, see [Custom Prefixes for Amazon S3 Objects](https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html) .
	ErrorOutputPrefix *string `json:"errorOutputPrefix" yaml:"errorOutputPrefix"`
	// The `YYYY/MM/DD/HH` time format prefix is automatically used for delivered Amazon S3 files.
	//
	// For more information, see [ExtendedS3DestinationConfiguration](https://docs.aws.amazon.com/firehose/latest/APIReference/API_ExtendedS3DestinationConfiguration.html) in the *Amazon Kinesis Data Firehose API Reference* .
	Prefix *string `json:"prefix" yaml:"prefix"`
	// The data processing configuration for the Kinesis Data Firehose delivery stream.
	ProcessingConfiguration interface{} `json:"processingConfiguration" yaml:"processingConfiguration"`
	// The configuration for backup in Amazon S3.
	S3BackupConfiguration interface{} `json:"s3BackupConfiguration" yaml:"s3BackupConfiguration"`
	// The Amazon S3 backup mode.
	//
	// After you create a delivery stream, you can update it to enable Amazon S3 backup if it is disabled. If backup is enabled, you can't update the delivery stream to disable it.
	S3BackupMode *string `json:"s3BackupMode" yaml:"s3BackupMode"`
}

// The native Hive / HCatalog JsonSerDe.
//
// Used by Kinesis Data Firehose for deserializing data, which means converting it from the JSON format in preparation for serializing it to the Parquet or ORC format. This is one of two deserializers you can choose, depending on which one offers the functionality you need. The other option is the OpenX SerDe.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_HiveJsonSerDeProperty struct {
	// Indicates how you want Kinesis Data Firehose to parse the date and timestamps that may be present in your input data JSON.
	//
	// To specify these format strings, follow the pattern syntax of JodaTime's DateTimeFormat format strings. For more information, see [Class DateTimeFormat](https://docs.aws.amazon.com/https://www.joda.org/joda-time/apidocs/org/joda/time/format/DateTimeFormat.html) . You can also use the special value `millis` to parse timestamps in epoch milliseconds. If you don't specify a format, Kinesis Data Firehose uses `java.sql.Timestamp::valueOf` by default.
	TimestampFormats *[]*string `json:"timestampFormats" yaml:"timestampFormats"`
}

// Describes the metadata that's delivered to the specified HTTP endpoint destination.
//
// Kinesis Firehose supports any custom HTTP endpoint or HTTP endpoints owned by supported third-party service providers, including Datadog, MongoDB, and New Relic.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_HttpEndpointCommonAttributeProperty struct {
	// The name of the HTTP endpoint common attribute.
	AttributeName *string `json:"attributeName" yaml:"attributeName"`
	// The value of the HTTP endpoint common attribute.
	AttributeValue *string `json:"attributeValue" yaml:"attributeValue"`
}

// Describes the configuration of the HTTP endpoint to which Kinesis Firehose delivers data.
//
// Kinesis Firehose supports any custom HTTP endpoint or HTTP endpoints owned by supported third-party service providers, including Datadog, MongoDB, and New Relic.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_HttpEndpointConfigurationProperty struct {
	// The URL of the HTTP endpoint selected as the destination.
	Url *string `json:"url" yaml:"url"`
	// The access key required for Kinesis Firehose to authenticate with the HTTP endpoint selected as the destination.
	AccessKey *string `json:"accessKey" yaml:"accessKey"`
	// The name of the HTTP endpoint selected as the destination.
	Name *string `json:"name" yaml:"name"`
}

// Describes the configuration of the HTTP endpoint destination.
//
// Kinesis Firehose supports any custom HTTP endpoint or HTTP endpoints owned by supported third-party service providers, including Datadog, MongoDB, and New Relic.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_HttpEndpointDestinationConfigurationProperty struct {
	// The configuration of the HTTP endpoint selected as the destination.
	EndpointConfiguration interface{} `json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// Describes the configuration of a destination in Amazon S3.
	S3Configuration interface{} `json:"s3Configuration" yaml:"s3Configuration"`
	// The buffering options that can be used before data is delivered to the specified destination.
	//
	// Kinesis Data Firehose treats these options as hints, and it might choose to use more optimal values. The SizeInMBs and IntervalInSeconds parameters are optional. However, if you specify a value for one of them, you must also provide a value for the other.
	BufferingHints interface{} `json:"bufferingHints" yaml:"bufferingHints"`
	// Describes the Amazon CloudWatch logging options for your delivery stream.
	CloudWatchLoggingOptions interface{} `json:"cloudWatchLoggingOptions" yaml:"cloudWatchLoggingOptions"`
	// Describes the data processing configuration.
	ProcessingConfiguration interface{} `json:"processingConfiguration" yaml:"processingConfiguration"`
	// The configuration of the request sent to the HTTP endpoint specified as the destination.
	RequestConfiguration interface{} `json:"requestConfiguration" yaml:"requestConfiguration"`
	// Describes the retry behavior in case Kinesis Data Firehose is unable to deliver data to the specified HTTP endpoint destination, or if it doesn't receive a valid acknowledgment of receipt from the specified HTTP endpoint destination.
	RetryOptions interface{} `json:"retryOptions" yaml:"retryOptions"`
	// Kinesis Data Firehose uses this IAM role for all the permissions that the delivery stream needs.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// Describes the S3 bucket backup options for the data that Kinesis Data Firehose delivers to the HTTP endpoint destination.
	//
	// You can back up all documents (AllData) or only the documents that Kinesis Data Firehose could not deliver to the specified HTTP endpoint destination (FailedDataOnly).
	S3BackupMode *string `json:"s3BackupMode" yaml:"s3BackupMode"`
}

// The configuration of the HTTP endpoint request.
//
// Kinesis Firehose supports any custom HTTP endpoint or HTTP endpoints owned by supported third-party service providers, including Datadog, MongoDB, and New Relic.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_HttpEndpointRequestConfigurationProperty struct {
	// Describes the metadata sent to the HTTP endpoint destination.
	CommonAttributes interface{} `json:"commonAttributes" yaml:"commonAttributes"`
	// Kinesis Data Firehose uses the content encoding to compress the body of a request before sending the request to the destination.
	//
	// For more information, see Content-Encoding in MDN Web Docs, the official Mozilla documentation.
	ContentEncoding *string `json:"contentEncoding" yaml:"contentEncoding"`
}

// Specifies the deserializer you want to use to convert the format of the input data.
//
// This parameter is required if `Enabled` is set to true.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_InputFormatConfigurationProperty struct {
	// Specifies which deserializer to use.
	//
	// You can choose either the Apache Hive JSON SerDe or the OpenX JSON SerDe. If both are non-null, the server rejects the request.
	Deserializer interface{} `json:"deserializer" yaml:"deserializer"`
}

// The `KMSEncryptionConfig` property type specifies the AWS Key Management Service ( AWS KMS) encryption key that Amazon Simple Storage Service (Amazon S3) uses to encrypt data delivered by the Amazon Kinesis Data Firehose (Kinesis Data Firehose) stream.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_KMSEncryptionConfigProperty struct {
	// The Amazon Resource Name (ARN) of the AWS KMS encryption key that Amazon S3 uses to encrypt data delivered by the Kinesis Data Firehose stream.
	//
	// The key must belong to the same region as the destination S3 bucket.
	AwskmsKeyArn *string `json:"awskmsKeyArn" yaml:"awskmsKeyArn"`
}

// The `KinesisStreamSourceConfiguration` property type specifies the stream and role Amazon Resource Names (ARNs) for a Kinesis stream used as the source for a delivery stream.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_KinesisStreamSourceConfigurationProperty struct {
	// The ARN of the source Kinesis data stream.
	KinesisStreamArn *string `json:"kinesisStreamArn" yaml:"kinesisStreamArn"`
	// The ARN of the role that provides access to the source Kinesis data stream.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
}

// The OpenX SerDe.
//
// Used by Kinesis Data Firehose for deserializing data, which means converting it from the JSON format in preparation for serializing it to the Parquet or ORC format. This is one of two deserializers you can choose, depending on which one offers the functionality you need. The other option is the native Hive / HCatalog JsonSerDe.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_OpenXJsonSerDeProperty struct {
	// When set to `true` , which is the default, Kinesis Data Firehose converts JSON keys to lowercase before deserializing them.
	CaseInsensitive interface{} `json:"caseInsensitive" yaml:"caseInsensitive"`
	// Maps column names to JSON keys that aren't identical to the column names.
	//
	// This is useful when the JSON contains keys that are Hive keywords. For example, `timestamp` is a Hive keyword. If you have a JSON key named `timestamp` , set this parameter to `{"ts": "timestamp"}` to map this key to a column named `ts` .
	ColumnToJsonKeyMappings interface{} `json:"columnToJsonKeyMappings" yaml:"columnToJsonKeyMappings"`
	// When set to `true` , specifies that the names of the keys include dots and that you want Kinesis Data Firehose to replace them with underscores.
	//
	// This is useful because Apache Hive does not allow dots in column names. For example, if the JSON contains a key whose name is "a.b", you can define the column name to be "a_b" when using this option.
	//
	// The default is `false` .
	ConvertDotsInJsonKeysToUnderscores interface{} `json:"convertDotsInJsonKeysToUnderscores" yaml:"convertDotsInJsonKeysToUnderscores"`
}

// A serializer to use for converting data to the ORC format before storing it in Amazon S3.
//
// For more information, see [Apache ORC](https://docs.aws.amazon.com/https://orc.apache.org/docs/) .
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_OrcSerDeProperty struct {
	// The Hadoop Distributed File System (HDFS) block size.
	//
	// This is useful if you intend to copy the data from Amazon S3 to HDFS before querying. The default is 256 MiB and the minimum is 64 MiB. Kinesis Data Firehose uses this value for padding calculations.
	BlockSizeBytes *float64 `json:"blockSizeBytes" yaml:"blockSizeBytes"`
	// The column names for which you want Kinesis Data Firehose to create bloom filters.
	//
	// The default is `null` .
	BloomFilterColumns *[]*string `json:"bloomFilterColumns" yaml:"bloomFilterColumns"`
	// The Bloom filter false positive probability (FPP).
	//
	// The lower the FPP, the bigger the Bloom filter. The default value is 0.05, the minimum is 0, and the maximum is 1.
	BloomFilterFalsePositiveProbability *float64 `json:"bloomFilterFalsePositiveProbability" yaml:"bloomFilterFalsePositiveProbability"`
	// The compression code to use over data blocks.
	//
	// The default is `SNAPPY` .
	Compression *string `json:"compression" yaml:"compression"`
	// Represents the fraction of the total number of non-null rows.
	//
	// To turn off dictionary encoding, set this fraction to a number that is less than the number of distinct keys in a dictionary. To always use dictionary encoding, set this threshold to 1.
	DictionaryKeyThreshold *float64 `json:"dictionaryKeyThreshold" yaml:"dictionaryKeyThreshold"`
	// Set this to `true` to indicate that you want stripes to be padded to the HDFS block boundaries.
	//
	// This is useful if you intend to copy the data from Amazon S3 to HDFS before querying. The default is `false` .
	EnablePadding interface{} `json:"enablePadding" yaml:"enablePadding"`
	// The version of the file to write.
	//
	// The possible values are `V0_11` and `V0_12` . The default is `V0_12` .
	FormatVersion *string `json:"formatVersion" yaml:"formatVersion"`
	// A number between 0 and 1 that defines the tolerance for block padding as a decimal fraction of stripe size.
	//
	// The default value is 0.05, which means 5 percent of stripe size.
	//
	// For the default values of 64 MiB ORC stripes and 256 MiB HDFS blocks, the default block padding tolerance of 5 percent reserves a maximum of 3.2 MiB for padding within the 256 MiB block. In such a case, if the available size within the block is more than 3.2 MiB, a new, smaller stripe is inserted to fit within that space. This ensures that no stripe crosses block boundaries and causes remote reads within a node-local task.
	//
	// Kinesis Data Firehose ignores this parameter when `EnablePadding` is `false` .
	PaddingTolerance *float64 `json:"paddingTolerance" yaml:"paddingTolerance"`
	// The number of rows between index entries.
	//
	// The default is 10,000 and the minimum is 1,000.
	RowIndexStride *float64 `json:"rowIndexStride" yaml:"rowIndexStride"`
	// The number of bytes in each stripe.
	//
	// The default is 64 MiB and the minimum is 8 MiB.
	StripeSizeBytes *float64 `json:"stripeSizeBytes" yaml:"stripeSizeBytes"`
}

// Specifies the serializer that you want Kinesis Data Firehose to use to convert the format of your data before it writes it to Amazon S3.
//
// This parameter is required if `Enabled` is set to true.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_OutputFormatConfigurationProperty struct {
	// Specifies which serializer to use.
	//
	// You can choose either the ORC SerDe or the Parquet SerDe. If both are non-null, the server rejects the request.
	Serializer interface{} `json:"serializer" yaml:"serializer"`
}

// A serializer to use for converting data to the Parquet format before storing it in Amazon S3.
//
// For more information, see [Apache Parquet](https://docs.aws.amazon.com/https://parquet.apache.org/documentation/latest/) .
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_ParquetSerDeProperty struct {
	// The Hadoop Distributed File System (HDFS) block size.
	//
	// This is useful if you intend to copy the data from Amazon S3 to HDFS before querying. The default is 256 MiB and the minimum is 64 MiB. Kinesis Data Firehose uses this value for padding calculations.
	BlockSizeBytes *float64 `json:"blockSizeBytes" yaml:"blockSizeBytes"`
	// The compression code to use over data blocks.
	//
	// The possible values are `UNCOMPRESSED` , `SNAPPY` , and `GZIP` , with the default being `SNAPPY` . Use `SNAPPY` for higher decompression speed. Use `GZIP` if the compression ratio is more important than speed.
	Compression *string `json:"compression" yaml:"compression"`
	// Indicates whether to enable dictionary compression.
	EnableDictionaryCompression interface{} `json:"enableDictionaryCompression" yaml:"enableDictionaryCompression"`
	// The maximum amount of padding to apply.
	//
	// This is useful if you intend to copy the data from Amazon S3 to HDFS before querying. The default is 0.
	MaxPaddingBytes *float64 `json:"maxPaddingBytes" yaml:"maxPaddingBytes"`
	// The Parquet page size.
	//
	// Column chunks are divided into pages. A page is conceptually an indivisible unit (in terms of compression and encoding). The minimum value is 64 KiB and the default is 1 MiB.
	PageSizeBytes *float64 `json:"pageSizeBytes" yaml:"pageSizeBytes"`
	// Indicates the version of row format to output.
	//
	// The possible values are `V1` and `V2` . The default is `V1` .
	WriterVersion *string `json:"writerVersion" yaml:"writerVersion"`
}

// The `ProcessingConfiguration` property configures data processing for an Amazon Kinesis Data Firehose delivery stream.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_ProcessingConfigurationProperty struct {
	// Indicates whether data processing is enabled (true) or disabled (false).
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// The data processors.
	Processors interface{} `json:"processors" yaml:"processors"`
}

// The `ProcessorParameter` property specifies a processor parameter in a data processor for an Amazon Kinesis Data Firehose delivery stream.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_ProcessorParameterProperty struct {
	// The name of the parameter.
	//
	// Currently the following default values are supported: 3 for `NumberOfRetries` , 60 for the `BufferIntervalInSeconds` , and 3 for the `BufferSizeInMBs` .
	ParameterName *string `json:"parameterName" yaml:"parameterName"`
	// The parameter value.
	ParameterValue *string `json:"parameterValue" yaml:"parameterValue"`
}

// The `Processor` property specifies a data processor for an Amazon Kinesis Data Firehose delivery stream.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_ProcessorProperty struct {
	// The type of processor.
	//
	// Valid values: `Lambda` .
	Type *string `json:"type" yaml:"type"`
	// The processor parameters.
	Parameters interface{} `json:"parameters" yaml:"parameters"`
}

// The `RedshiftDestinationConfiguration` property type specifies an Amazon Redshift cluster to which Amazon Kinesis Data Firehose (Kinesis Data Firehose) delivers data.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_RedshiftDestinationConfigurationProperty struct {
	// The connection string that Kinesis Data Firehose uses to connect to the Amazon Redshift cluster.
	ClusterJdbcurl *string `json:"clusterJdbcurl" yaml:"clusterJdbcurl"`
	// Configures the Amazon Redshift `COPY` command that Kinesis Data Firehose uses to load data into the cluster from the Amazon S3 bucket.
	CopyCommand interface{} `json:"copyCommand" yaml:"copyCommand"`
	// The password for the Amazon Redshift user that you specified in the `Username` property.
	Password *string `json:"password" yaml:"password"`
	// The ARN of the AWS Identity and Access Management (IAM) role that grants Kinesis Data Firehose access to your Amazon S3 bucket and AWS KMS (if you enable data encryption).
	//
	// For more information, see [Grant Kinesis Data Firehose Access to an Amazon Redshift Destination](https://docs.aws.amazon.com/firehose/latest/dev/controlling-access.html#using-iam-rs) in the *Amazon Kinesis Data Firehose Developer Guide* .
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The S3 bucket where Kinesis Data Firehose first delivers data.
	//
	// After the data is in the bucket, Kinesis Data Firehose uses the `COPY` command to load the data into the Amazon Redshift cluster. For the Amazon S3 bucket's compression format, don't specify `SNAPPY` or `ZIP` because the Amazon Redshift `COPY` command doesn't support them.
	S3Configuration interface{} `json:"s3Configuration" yaml:"s3Configuration"`
	// The Amazon Redshift user that has permission to access the Amazon Redshift cluster.
	//
	// This user must have `INSERT` privileges for copying data from the Amazon S3 bucket to the cluster.
	Username *string `json:"username" yaml:"username"`
	// The CloudWatch logging options for your delivery stream.
	CloudWatchLoggingOptions interface{} `json:"cloudWatchLoggingOptions" yaml:"cloudWatchLoggingOptions"`
	// The data processing configuration for the Kinesis Data Firehose delivery stream.
	ProcessingConfiguration interface{} `json:"processingConfiguration" yaml:"processingConfiguration"`
	// The retry behavior in case Kinesis Data Firehose is unable to deliver documents to Amazon Redshift.
	//
	// Default value is 3600 (60 minutes).
	RetryOptions interface{} `json:"retryOptions" yaml:"retryOptions"`
	// The configuration for backup in Amazon S3.
	S3BackupConfiguration interface{} `json:"s3BackupConfiguration" yaml:"s3BackupConfiguration"`
	// The Amazon S3 backup mode.
	//
	// After you create a delivery stream, you can update it to enable Amazon S3 backup if it is disabled. If backup is enabled, you can't update the delivery stream to disable it.
	S3BackupMode *string `json:"s3BackupMode" yaml:"s3BackupMode"`
}

// Configures retry behavior in case Kinesis Data Firehose is unable to deliver documents to Amazon Redshift.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_RedshiftRetryOptionsProperty struct {
	// The length of time during which Kinesis Data Firehose retries delivery after a failure, starting from the initial request and including the first attempt.
	//
	// The default value is 3600 seconds (60 minutes). Kinesis Data Firehose does not retry if the value of `DurationInSeconds` is 0 (zero) or if the first delivery attempt takes longer than the current value.
	DurationInSeconds *float64 `json:"durationInSeconds" yaml:"durationInSeconds"`
}

// Describes the retry behavior in case Kinesis Data Firehose is unable to deliver data to the specified HTTP endpoint destination, or if it doesn't receive a valid acknowledgment of receipt from the specified HTTP endpoint destination.
//
// Kinesis Firehose supports any custom HTTP endpoint or HTTP endpoints owned by supported third-party service providers, including Datadog, MongoDB, and New Relic.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_RetryOptionsProperty struct {
	// The total amount of time that Kinesis Data Firehose spends on retries.
	//
	// This duration starts after the initial attempt to send data to the custom destination via HTTPS endpoint fails. It doesn't include the periods during which Kinesis Data Firehose waits for acknowledgment from the specified destination after each attempt.
	DurationInSeconds *float64 `json:"durationInSeconds" yaml:"durationInSeconds"`
}

// The `S3DestinationConfiguration` property type specifies an Amazon Simple Storage Service (Amazon S3) destination to which Amazon Kinesis Data Firehose (Kinesis Data Firehose) delivers data.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_S3DestinationConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon S3 bucket to send data to.
	BucketArn *string `json:"bucketArn" yaml:"bucketArn"`
	// The ARN of an AWS Identity and Access Management (IAM) role that grants Kinesis Data Firehose access to your Amazon S3 bucket and AWS KMS (if you enable data encryption).
	//
	// For more information, see [Grant Kinesis Data Firehose Access to an Amazon S3 Destination](https://docs.aws.amazon.com/firehose/latest/dev/controlling-access.html#using-iam-s3) in the *Amazon Kinesis Data Firehose Developer Guide* .
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// Configures how Kinesis Data Firehose buffers incoming data while delivering it to the Amazon S3 bucket.
	BufferingHints interface{} `json:"bufferingHints" yaml:"bufferingHints"`
	// The CloudWatch logging options for your delivery stream.
	CloudWatchLoggingOptions interface{} `json:"cloudWatchLoggingOptions" yaml:"cloudWatchLoggingOptions"`
	// The type of compression that Kinesis Data Firehose uses to compress the data that it delivers to the Amazon S3 bucket.
	//
	// For valid values, see the `CompressionFormat` content for the [S3DestinationConfiguration](https://docs.aws.amazon.com/firehose/latest/APIReference/API_S3DestinationConfiguration.html) data type in the *Amazon Kinesis Data Firehose API Reference* .
	CompressionFormat *string `json:"compressionFormat" yaml:"compressionFormat"`
	// Configures Amazon Simple Storage Service (Amazon S3) server-side encryption.
	//
	// Kinesis Data Firehose uses AWS Key Management Service ( AWS KMS) to encrypt the data that it delivers to your Amazon S3 bucket.
	EncryptionConfiguration interface{} `json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3.
	//
	// This prefix appears immediately following the bucket name. For information about how to specify this prefix, see [Custom Prefixes for Amazon S3 Objects](https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html) .
	ErrorOutputPrefix *string `json:"errorOutputPrefix" yaml:"errorOutputPrefix"`
	// A prefix that Kinesis Data Firehose adds to the files that it delivers to the Amazon S3 bucket.
	//
	// The prefix helps you identify the files that Kinesis Data Firehose delivered.
	Prefix *string `json:"prefix" yaml:"prefix"`
}

// Specifies the schema to which you want Kinesis Data Firehose to configure your data before it writes it to Amazon S3.
//
// This parameter is required if `Enabled` is set to true.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_SchemaConfigurationProperty struct {
	// The ID of the AWS Glue Data Catalog.
	//
	// If you don't supply this, the AWS account ID is used by default.
	CatalogId *string `json:"catalogId" yaml:"catalogId"`
	// Specifies the name of the AWS Glue database that contains the schema for the output data.
	//
	// > If the `SchemaConfiguration` request parameter is used as part of invoking the `CreateDeliveryStream` API, then the `DatabaseName` property is required and its value must be specified.
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// If you don't specify an AWS Region, the default is the current Region.
	Region *string `json:"region" yaml:"region"`
	// The role that Kinesis Data Firehose can use to access AWS Glue.
	//
	// This role must be in the same account you use for Kinesis Data Firehose. Cross-account roles aren't allowed.
	//
	// > If the `SchemaConfiguration` request parameter is used as part of invoking the `CreateDeliveryStream` API, then the `RoleARN` property is required and its value must be specified.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// Specifies the AWS Glue table that contains the column information that constitutes your data schema.
	//
	// > If the `SchemaConfiguration` request parameter is used as part of invoking the `CreateDeliveryStream` API, then the `TableName` property is required and its value must be specified.
	TableName *string `json:"tableName" yaml:"tableName"`
	// Specifies the table version for the output data schema.
	//
	// If you don't specify this version ID, or if you set it to `LATEST` , Kinesis Data Firehose uses the most recent version. This means that any updates to the table are automatically picked up.
	VersionId *string `json:"versionId" yaml:"versionId"`
}

// The serializer that you want Kinesis Data Firehose to use to convert data to the target format before writing it to Amazon S3.
//
// Kinesis Data Firehose supports two types of serializers: the [ORC SerDe](https://docs.aws.amazon.com/https://hive.apache.org/javadocs/r1.2.2/api/org/apache/hadoop/hive/ql/io/orc/OrcSerde.html) and the [Parquet SerDe](https://docs.aws.amazon.com/https://hive.apache.org/javadocs/r1.2.2/api/org/apache/hadoop/hive/ql/io/parquet/serde/ParquetHiveSerDe.html) .
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_SerializerProperty struct {
	// A serializer to use for converting data to the ORC format before storing it in Amazon S3.
	//
	// For more information, see [Apache ORC](https://docs.aws.amazon.com/https://orc.apache.org/docs/) .
	OrcSerDe interface{} `json:"orcSerDe" yaml:"orcSerDe"`
	// A serializer to use for converting data to the Parquet format before storing it in Amazon S3.
	//
	// For more information, see [Apache Parquet](https://docs.aws.amazon.com/https://parquet.apache.org/documentation/latest/) .
	ParquetSerDe interface{} `json:"parquetSerDe" yaml:"parquetSerDe"`
}

// The `SplunkDestinationConfiguration` property type specifies the configuration of a destination in Splunk for a Kinesis Data Firehose delivery stream.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_SplunkDestinationConfigurationProperty struct {
	// The HTTP Event Collector (HEC) endpoint to which Kinesis Data Firehose sends your data.
	HecEndpoint *string `json:"hecEndpoint" yaml:"hecEndpoint"`
	// This type can be either `Raw` or `Event` .
	HecEndpointType *string `json:"hecEndpointType" yaml:"hecEndpointType"`
	// This is a GUID that you obtain from your Splunk cluster when you create a new HEC endpoint.
	HecToken *string `json:"hecToken" yaml:"hecToken"`
	// The configuration for the backup Amazon S3 location.
	S3Configuration interface{} `json:"s3Configuration" yaml:"s3Configuration"`
	// The Amazon CloudWatch logging options for your delivery stream.
	CloudWatchLoggingOptions interface{} `json:"cloudWatchLoggingOptions" yaml:"cloudWatchLoggingOptions"`
	// The amount of time that Kinesis Data Firehose waits to receive an acknowledgment from Splunk after it sends it data.
	//
	// At the end of the timeout period, Kinesis Data Firehose either tries to send the data again or considers it an error, based on your retry settings.
	HecAcknowledgmentTimeoutInSeconds *float64 `json:"hecAcknowledgmentTimeoutInSeconds" yaml:"hecAcknowledgmentTimeoutInSeconds"`
	// The data processing configuration.
	ProcessingConfiguration interface{} `json:"processingConfiguration" yaml:"processingConfiguration"`
	// The retry behavior in case Kinesis Data Firehose is unable to deliver data to Splunk, or if it doesn't receive an acknowledgment of receipt from Splunk.
	RetryOptions interface{} `json:"retryOptions" yaml:"retryOptions"`
	// Defines how documents should be delivered to Amazon S3.
	//
	// When set to `FailedEventsOnly` , Kinesis Data Firehose writes any data that could not be indexed to the configured Amazon S3 destination. When set to `AllEvents` , Kinesis Data Firehose delivers all incoming records to Amazon S3, and also writes failed documents to Amazon S3. The default value is `FailedEventsOnly` .
	//
	// You can update this backup mode from `FailedEventsOnly` to `AllEvents` . You can't update it from `AllEvents` to `FailedEventsOnly` .
	S3BackupMode *string `json:"s3BackupMode" yaml:"s3BackupMode"`
}

// The `SplunkRetryOptions` property type specifies retry behavior in case Kinesis Data Firehose is unable to deliver documents to Splunk or if it doesn't receive an acknowledgment from Splunk.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_SplunkRetryOptionsProperty struct {
	// The total amount of time that Kinesis Data Firehose spends on retries.
	//
	// This duration starts after the initial attempt to send data to Splunk fails. It doesn't include the periods during which Kinesis Data Firehose waits for acknowledgment from Splunk after each attempt.
	DurationInSeconds *float64 `json:"durationInSeconds" yaml:"durationInSeconds"`
}

// The details of the VPC of the Amazon ES destination.
//
// TODO: EXAMPLE
//
type CfnDeliveryStream_VpcConfigurationProperty struct {
	// The ARN of the IAM role that you want the delivery stream to use to create endpoints in the destination VPC.
	//
	// You can use your existing Kinesis Data Firehose delivery role or you can specify a new role. In either case, make sure that the role trusts the Kinesis Data Firehose service principal and that it grants the following permissions:
	//
	// - `ec2:DescribeVpcs`
	// - `ec2:DescribeVpcAttribute`
	// - `ec2:DescribeSubnets`
	// - `ec2:DescribeSecurityGroups`
	// - `ec2:DescribeNetworkInterfaces`
	// - `ec2:CreateNetworkInterface`
	// - `ec2:CreateNetworkInterfacePermission`
	// - `ec2:DeleteNetworkInterface`
	//
	// If you revoke these permissions after you create the delivery stream, Kinesis Data Firehose can't scale out by creating more ENIs when necessary. You might therefore see a degradation in performance.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The IDs of the security groups that you want Kinesis Data Firehose to use when it creates ENIs in the VPC of the Amazon ES destination.
	//
	// You can use the same security group that the Amazon ES domain uses or different ones. If you specify different security groups here, ensure that they allow outbound HTTPS traffic to the Amazon ES domain's security group. Also ensure that the Amazon ES domain's security group allows HTTPS traffic from the security groups specified here. If you use the same security group for both your delivery stream and the Amazon ES domain, make sure the security group inbound rule allows HTTPS traffic.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// The IDs of the subnets that Kinesis Data Firehose uses to create ENIs in the VPC of the Amazon ES destination.
	//
	// Make sure that the routing tables and inbound and outbound rules allow traffic to flow from the subnets whose IDs are specified here to the subnets that have the destination Amazon ES endpoints. Kinesis Data Firehose creates at least one ENI in each of the subnets that are specified here. Do not delete or modify these ENIs.
	//
	// The number of ENIs that Kinesis Data Firehose creates in the subnets specified here scales up and down automatically based on throughput. To enable Kinesis Data Firehose to scale up the number of ENIs to match throughput, ensure that you have sufficient quota. To help you calculate the quota you need, assume that Kinesis Data Firehose can create up to three ENIs for this delivery stream for each of the subnets specified here.
	SubnetIds *[]*string `json:"subnetIds" yaml:"subnetIds"`
}

// Properties for defining a `CfnDeliveryStream`.
//
// TODO: EXAMPLE
//
type CfnDeliveryStreamProps struct {
	// The destination in Amazon OpenSearch Service.
	//
	// You can specify only one destination.
	AmazonopensearchserviceDestinationConfiguration interface{} `json:"amazonopensearchserviceDestinationConfiguration" yaml:"amazonopensearchserviceDestinationConfiguration"`
	// Specifies the type and Amazon Resource Name (ARN) of the CMK to use for Server-Side Encryption (SSE).
	DeliveryStreamEncryptionConfigurationInput interface{} `json:"deliveryStreamEncryptionConfigurationInput" yaml:"deliveryStreamEncryptionConfigurationInput"`
	// The name of the delivery stream.
	DeliveryStreamName *string `json:"deliveryStreamName" yaml:"deliveryStreamName"`
	// The delivery stream type. This can be one of the following values:.
	//
	// - `DirectPut` : Provider applications access the delivery stream directly.
	// - `KinesisStreamAsSource` : The delivery stream uses a Kinesis data stream as a source.
	DeliveryStreamType *string `json:"deliveryStreamType" yaml:"deliveryStreamType"`
	// An Amazon ES destination for the delivery stream.
	//
	// Conditional. You must specify only one destination configuration.
	//
	// If you change the delivery stream destination from an Amazon ES destination to an Amazon S3 or Amazon Redshift destination, update requires [some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) .
	ElasticsearchDestinationConfiguration interface{} `json:"elasticsearchDestinationConfiguration" yaml:"elasticsearchDestinationConfiguration"`
	// An Amazon S3 destination for the delivery stream.
	//
	// Conditional. You must specify only one destination configuration.
	//
	// If you change the delivery stream destination from an Amazon Extended S3 destination to an Amazon ES destination, update requires [some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) .
	ExtendedS3DestinationConfiguration interface{} `json:"extendedS3DestinationConfiguration" yaml:"extendedS3DestinationConfiguration"`
	// Enables configuring Kinesis Firehose to deliver data to any HTTP endpoint destination.
	//
	// You can specify only one destination.
	HttpEndpointDestinationConfiguration interface{} `json:"httpEndpointDestinationConfiguration" yaml:"httpEndpointDestinationConfiguration"`
	// When a Kinesis stream is used as the source for the delivery stream, a [KinesisStreamSourceConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-kinesisstreamsourceconfiguration.html) containing the Kinesis stream ARN and the role ARN for the source stream.
	KinesisStreamSourceConfiguration interface{} `json:"kinesisStreamSourceConfiguration" yaml:"kinesisStreamSourceConfiguration"`
	// An Amazon Redshift destination for the delivery stream.
	//
	// Conditional. You must specify only one destination configuration.
	//
	// If you change the delivery stream destination from an Amazon Redshift destination to an Amazon ES destination, update requires [some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) .
	RedshiftDestinationConfiguration interface{} `json:"redshiftDestinationConfiguration" yaml:"redshiftDestinationConfiguration"`
	// The `S3DestinationConfiguration` property type specifies an Amazon Simple Storage Service (Amazon S3) destination to which Amazon Kinesis Data Firehose (Kinesis Data Firehose) delivers data.
	//
	// Conditional. You must specify only one destination configuration.
	//
	// If you change the delivery stream destination from an Amazon S3 destination to an Amazon ES destination, update requires [some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) .
	S3DestinationConfiguration interface{} `json:"s3DestinationConfiguration" yaml:"s3DestinationConfiguration"`
	// The configuration of a destination in Splunk for the delivery stream.
	SplunkDestinationConfiguration interface{} `json:"splunkDestinationConfiguration" yaml:"splunkDestinationConfiguration"`
	// A set of tags to assign to the delivery stream.
	//
	// A tag is a key-value pair that you can define and assign to AWS resources. Tags are metadata. For example, you can add friendly names and descriptions or other types of information that can help you distinguish the delivery stream. For more information about tags, see [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html) in the AWS Billing and Cost Management User Guide.
	//
	// You can specify up to 50 tags when creating a delivery stream.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// Options when binding a DataProcessor to a delivery stream destination.
//
// TODO: EXAMPLE
//
// Experimental.
type DataProcessorBindOptions struct {
	// The IAM role assumed by Kinesis Data Firehose to write to the destination that this DataProcessor will bind to.
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
}

// The full configuration of a data processor.
//
// TODO: EXAMPLE
//
// Experimental.
type DataProcessorConfig struct {
	// The key-value pair that identifies the underlying processor resource.
	// Experimental.
	ProcessorIdentifier *DataProcessorIdentifier `json:"processorIdentifier" yaml:"processorIdentifier"`
	// The type of the underlying processor resource.
	//
	// Must be an accepted value in `CfnDeliveryStream.ProcessorProperty.Type`.
	//
	// TODO: EXAMPLE
	//
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processor.html#cfn-kinesisfirehose-deliverystream-processor-type
	//
	// Experimental.
	ProcessorType *string `json:"processorType" yaml:"processorType"`
}

// The key-value pair that identifies the underlying processor resource.
//
// TODO: EXAMPLE
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processorparameter.html
//
// Experimental.
type DataProcessorIdentifier struct {
	// The parameter name that corresponds to the processor resource's identifier.
	//
	// Must be an accepted value in `CfnDeliveryStream.ProcessoryParameterProperty.ParameterName`.
	// Experimental.
	ParameterName *string `json:"parameterName" yaml:"parameterName"`
	// The identifier of the underlying processor resource.
	// Experimental.
	ParameterValue *string `json:"parameterValue" yaml:"parameterValue"`
}

// Configure the data processor.
//
// TODO: EXAMPLE
//
// Experimental.
type DataProcessorProps struct {
	// The length of time Kinesis Data Firehose will buffer incoming data before calling the processor.
	//
	// s
	// Experimental.
	BufferInterval awscdk.Duration `json:"bufferInterval" yaml:"bufferInterval"`
	// The amount of incoming data Kinesis Data Firehose will buffer before calling the processor.
	// Experimental.
	BufferSize awscdk.Size `json:"bufferSize" yaml:"bufferSize"`
	// The number of times Kinesis Data Firehose will retry the processor invocation after a failure due to network timeout or invocation limits.
	// Experimental.
	Retries *float64 `json:"retries" yaml:"retries"`
}

// Create a Kinesis Data Firehose delivery stream.
//
// TODO: EXAMPLE
//
// Experimental.
type DeliveryStream interface {
	awscdk.Resource
	IDeliveryStream
	Connections() awsec2.Connections
	DeliveryStreamArn() *string
	DeliveryStreamName() *string
	Env() *awscdk.ResourceEnvironment
	GrantPrincipal() awsiam.IPrincipal
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	GrantPutRecords(grantee awsiam.IGrantable) awsiam.Grant
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricBackupToS3Bytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricBackupToS3DataFreshness(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricBackupToS3Records(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricIncomingBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricIncomingRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for DeliveryStream
type jsiiProxy_DeliveryStream struct {
	internal.Type__awscdkResource
	jsiiProxy_IDeliveryStream
}

func (j *jsiiProxy_DeliveryStream) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) DeliveryStreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) DeliveryStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewDeliveryStream(scope constructs.Construct, id *string, props *DeliveryStreamProps) DeliveryStream {
	_init_.Initialize()

	j := jsiiProxy_DeliveryStream{}

	_jsii_.Create(
		"monocdk.aws_kinesisfirehose.DeliveryStream",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDeliveryStream_Override(d DeliveryStream, scope constructs.Construct, id *string, props *DeliveryStreamProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_kinesisfirehose.DeliveryStream",
		[]interface{}{scope, id, props},
		d,
	)
}

// Import an existing delivery stream from its ARN.
// Experimental.
func DeliveryStream_FromDeliveryStreamArn(scope constructs.Construct, id *string, deliveryStreamArn *string) IDeliveryStream {
	_init_.Initialize()

	var returns IDeliveryStream

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisfirehose.DeliveryStream",
		"fromDeliveryStreamArn",
		[]interface{}{scope, id, deliveryStreamArn},
		&returns,
	)

	return returns
}

// Import an existing delivery stream from its attributes.
// Experimental.
func DeliveryStream_FromDeliveryStreamAttributes(scope constructs.Construct, id *string, attrs *DeliveryStreamAttributes) IDeliveryStream {
	_init_.Initialize()

	var returns IDeliveryStream

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisfirehose.DeliveryStream",
		"fromDeliveryStreamAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import an existing delivery stream from its name.
// Experimental.
func DeliveryStream_FromDeliveryStreamName(scope constructs.Construct, id *string, deliveryStreamName *string) IDeliveryStream {
	_init_.Initialize()

	var returns IDeliveryStream

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisfirehose.DeliveryStream",
		"fromDeliveryStreamName",
		[]interface{}{scope, id, deliveryStreamName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func DeliveryStream_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisfirehose.DeliveryStream",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func DeliveryStream_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisfirehose.DeliveryStream",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
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
func (d *jsiiProxy_DeliveryStream) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (d *jsiiProxy_DeliveryStream) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
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
func (d *jsiiProxy_DeliveryStream) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
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
func (d *jsiiProxy_DeliveryStream) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant the `grantee` identity permissions to perform `actions`.
// Experimental.
func (d *jsiiProxy_DeliveryStream) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grant",
		args,
		&returns,
	)

	return returns
}

// Grant the `grantee` identity permissions to perform `firehose:PutRecord` and `firehose:PutRecordBatch` actions on this delivery stream.
// Experimental.
func (d *jsiiProxy_DeliveryStream) GrantPutRecords(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantPutRecords",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Return the given named metric for this delivery stream.
// Experimental.
func (d *jsiiProxy_DeliveryStream) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of bytes delivered to Amazon S3 for backup over the specified time period.
//
// By default, this metric will be calculated as an average over a period of 5 minutes.
// Experimental.
func (d *jsiiProxy_DeliveryStream) MetricBackupToS3Bytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricBackupToS3Bytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the age (from getting into Kinesis Data Firehose to now) of the oldest record in Kinesis Data Firehose.
//
// Any record older than this age has been delivered to the Amazon S3 bucket for backup.
//
// By default, this metric will be calculated as an average over a period of 5 minutes.
// Experimental.
func (d *jsiiProxy_DeliveryStream) MetricBackupToS3DataFreshness(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricBackupToS3DataFreshness",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of records delivered to Amazon S3 for backup over the specified time period.
//
// By default, this metric will be calculated as an average over a period of 5 minutes.
// Experimental.
func (d *jsiiProxy_DeliveryStream) MetricBackupToS3Records(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricBackupToS3Records",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of bytes ingested successfully into the delivery stream over the specified time period after throttling.
//
// By default, this metric will be calculated as an average over a period of 5 minutes.
// Experimental.
func (d *jsiiProxy_DeliveryStream) MetricIncomingBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricIncomingBytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of records ingested successfully into the delivery stream over the specified time period after throttling.
//
// By default, this metric will be calculated as an average over a period of 5 minutes.
// Experimental.
func (d *jsiiProxy_DeliveryStream) MetricIncomingRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricIncomingRecords",
		[]interface{}{props},
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
func (d *jsiiProxy_DeliveryStream) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DeliveryStream) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
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
func (d *jsiiProxy_DeliveryStream) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
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
func (d *jsiiProxy_DeliveryStream) Prepare() {
	_jsii_.InvokeVoid(
		d,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DeliveryStream) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DeliveryStream) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
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
func (d *jsiiProxy_DeliveryStream) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A full specification of a delivery stream that can be used to import it fluently into the CDK application.
//
// TODO: EXAMPLE
//
// Experimental.
type DeliveryStreamAttributes struct {
	// The ARN of the delivery stream.
	//
	// At least one of deliveryStreamArn and deliveryStreamName must be provided.
	// Experimental.
	DeliveryStreamArn *string `json:"deliveryStreamArn" yaml:"deliveryStreamArn"`
	// The name of the delivery stream.
	//
	// At least one of deliveryStreamName and deliveryStreamArn  must be provided.
	// Experimental.
	DeliveryStreamName *string `json:"deliveryStreamName" yaml:"deliveryStreamName"`
	// The IAM role associated with this delivery stream.
	//
	// Assumed by Kinesis Data Firehose to read from sources and encrypt data server-side.
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
}

// Properties for a new delivery stream.
//
// TODO: EXAMPLE
//
// Experimental.
type DeliveryStreamProps struct {
	// The destinations that this delivery stream will deliver data to.
	//
	// Only a singleton array is supported at this time.
	// Experimental.
	Destinations *[]IDestination `json:"destinations" yaml:"destinations"`
	// A name for the delivery stream.
	// Experimental.
	DeliveryStreamName *string `json:"deliveryStreamName" yaml:"deliveryStreamName"`
	// Indicates the type of customer master key (CMK) to use for server-side encryption, if any.
	// Experimental.
	Encryption StreamEncryption `json:"encryption" yaml:"encryption"`
	// Customer managed key to server-side encrypt data in the stream.
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey" yaml:"encryptionKey"`
	// The IAM role associated with this delivery stream.
	//
	// Assumed by Kinesis Data Firehose to read from sources and encrypt data server-side.
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
	// The Kinesis data stream to use as a source for this delivery stream.
	// Experimental.
	SourceStream awskinesis.IStream `json:"sourceStream" yaml:"sourceStream"`
}

// Options when binding a destination to a delivery stream.
//
// TODO: EXAMPLE
//
// Experimental.
type DestinationBindOptions struct {
}

// A Kinesis Data Firehose delivery stream destination configuration.
//
// TODO: EXAMPLE
//
// Experimental.
type DestinationConfig struct {
	// Any resources that were created by the destination when binding it to the stack that must be deployed before the delivery stream is deployed.
	// Experimental.
	Dependables *[]awscdk.IDependable `json:"dependables" yaml:"dependables"`
	// S3 destination configuration properties.
	// Experimental.
	ExtendedS3DestinationConfiguration *CfnDeliveryStream_ExtendedS3DestinationConfigurationProperty `json:"extendedS3DestinationConfiguration" yaml:"extendedS3DestinationConfiguration"`
}

// A data processor that Kinesis Data Firehose will call to transform records before delivering data.
// Experimental.
type IDataProcessor interface {
	// Binds this processor to a destination of a delivery stream.
	//
	// Implementers should use this method to grant processor invocation permissions to the provided stream and return the
	// necessary configuration to register as a processor.
	// Experimental.
	Bind(scope constructs.Construct, options *DataProcessorBindOptions) *DataProcessorConfig
	// The constructor props of the DataProcessor.
	// Experimental.
	Props() *DataProcessorProps
}

// The jsii proxy for IDataProcessor
type jsiiProxy_IDataProcessor struct {
	_ byte // padding
}

func (i *jsiiProxy_IDataProcessor) Bind(scope constructs.Construct, options *DataProcessorBindOptions) *DataProcessorConfig {
	var returns *DataProcessorConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IDataProcessor) Props() *DataProcessorProps {
	var returns *DataProcessorProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

// Represents a Kinesis Data Firehose delivery stream.
// Experimental.
type IDeliveryStream interface {
	awsec2.IConnectable
	awsiam.IGrantable
	awscdk.IResource
	// Grant the `grantee` identity permissions to perform `actions`.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the `grantee` identity permissions to perform `firehose:PutRecord` and `firehose:PutRecordBatch` actions on this delivery stream.
	// Experimental.
	GrantPutRecords(grantee awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this delivery stream.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of bytes delivered to Amazon S3 for backup over the specified time period.
	//
	// By default, this metric will be calculated as an average over a period of 5 minutes.
	// Experimental.
	MetricBackupToS3Bytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the age (from getting into Kinesis Data Firehose to now) of the oldest record in Kinesis Data Firehose.
	//
	// Any record older than this age has been delivered to the Amazon S3 bucket for backup.
	//
	// By default, this metric will be calculated as an average over a period of 5 minutes.
	// Experimental.
	MetricBackupToS3DataFreshness(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of records delivered to Amazon S3 for backup over the specified time period.
	//
	// By default, this metric will be calculated as an average over a period of 5 minutes.
	// Experimental.
	MetricBackupToS3Records(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of bytes ingested successfully into the delivery stream over the specified time period after throttling.
	//
	// By default, this metric will be calculated as an average over a period of 5 minutes.
	// Experimental.
	MetricIncomingBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of records ingested successfully into the delivery stream over the specified time period after throttling.
	//
	// By default, this metric will be calculated as an average over a period of 5 minutes.
	// Experimental.
	MetricIncomingRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The ARN of the delivery stream.
	// Experimental.
	DeliveryStreamArn() *string
	// The name of the delivery stream.
	// Experimental.
	DeliveryStreamName() *string
}

// The jsii proxy for IDeliveryStream
type jsiiProxy_IDeliveryStream struct {
	internal.Type__awsec2IConnectable
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IDeliveryStream) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	args := []interface{}{grantee}
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

func (i *jsiiProxy_IDeliveryStream) GrantPutRecords(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantPutRecords",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDeliveryStream) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDeliveryStream) MetricBackupToS3Bytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricBackupToS3Bytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDeliveryStream) MetricBackupToS3DataFreshness(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricBackupToS3DataFreshness",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDeliveryStream) MetricBackupToS3Records(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricBackupToS3Records",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDeliveryStream) MetricIncomingBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricIncomingBytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDeliveryStream) MetricIncomingRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricIncomingRecords",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDeliveryStream) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IDeliveryStream) DeliveryStreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeliveryStream) DeliveryStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeliveryStream) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeliveryStream) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeliveryStream) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeliveryStream) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeliveryStream) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// A Kinesis Data Firehose delivery stream destination.
// Experimental.
type IDestination interface {
	// Binds this destination to the Kinesis Data Firehose delivery stream.
	//
	// Implementers should use this method to bind resources to the stack and initialize values using the provided stream.
	// Experimental.
	Bind(scope constructs.Construct, options *DestinationBindOptions) *DestinationConfig
}

// The jsii proxy for IDestination
type jsiiProxy_IDestination struct {
	_ byte // padding
}

func (i *jsiiProxy_IDestination) Bind(scope constructs.Construct, options *DestinationBindOptions) *DestinationConfig {
	var returns *DestinationConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

// Use an AWS Lambda function to transform records.
//
// TODO: EXAMPLE
//
// Experimental.
type LambdaFunctionProcessor interface {
	IDataProcessor
	Props() *DataProcessorProps
	Bind(_scope constructs.Construct, options *DataProcessorBindOptions) *DataProcessorConfig
}

// The jsii proxy struct for LambdaFunctionProcessor
type jsiiProxy_LambdaFunctionProcessor struct {
	jsiiProxy_IDataProcessor
}

func (j *jsiiProxy_LambdaFunctionProcessor) Props() *DataProcessorProps {
	var returns *DataProcessorProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaFunctionProcessor(lambdaFunction awslambda.IFunction, props *DataProcessorProps) LambdaFunctionProcessor {
	_init_.Initialize()

	j := jsiiProxy_LambdaFunctionProcessor{}

	_jsii_.Create(
		"monocdk.aws_kinesisfirehose.LambdaFunctionProcessor",
		[]interface{}{lambdaFunction, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaFunctionProcessor_Override(l LambdaFunctionProcessor, lambdaFunction awslambda.IFunction, props *DataProcessorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_kinesisfirehose.LambdaFunctionProcessor",
		[]interface{}{lambdaFunction, props},
		l,
	)
}

// Binds this processor to a destination of a delivery stream.
//
// Implementers should use this method to grant processor invocation permissions to the provided stream and return the
// necessary configuration to register as a processor.
// Experimental.
func (l *jsiiProxy_LambdaFunctionProcessor) Bind(_scope constructs.Construct, options *DataProcessorBindOptions) *DataProcessorConfig {
	var returns *DataProcessorConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{_scope, options},
		&returns,
	)

	return returns
}

// Options for server-side encryption of a delivery stream.
//
// TODO: EXAMPLE
//
// Experimental.
type StreamEncryption string

const (
	StreamEncryption_UNENCRYPTED StreamEncryption = "UNENCRYPTED"
	StreamEncryption_CUSTOMER_MANAGED StreamEncryption = "CUSTOMER_MANAGED"
	StreamEncryption_AWS_OWNED StreamEncryption = "AWS_OWNED"
)


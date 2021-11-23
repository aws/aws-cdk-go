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

// TODO: EXAMPLE
//
type CfnDeliveryStream_AmazonopensearchserviceBufferingHintsProperty struct {
	// `CfnDeliveryStream.AmazonopensearchserviceBufferingHintsProperty.IntervalInSeconds`.
	IntervalInSeconds *float64 `json:"intervalInSeconds"`
	// `CfnDeliveryStream.AmazonopensearchserviceBufferingHintsProperty.SizeInMBs`.
	SizeInMBs *float64 `json:"sizeInMBs"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_AmazonopensearchserviceDestinationConfigurationProperty struct {
	// `CfnDeliveryStream.AmazonopensearchserviceDestinationConfigurationProperty.BufferingHints`.
	BufferingHints interface{} `json:"bufferingHints"`
	// `CfnDeliveryStream.AmazonopensearchserviceDestinationConfigurationProperty.CloudWatchLoggingOptions`.
	CloudWatchLoggingOptions interface{} `json:"cloudWatchLoggingOptions"`
	// `CfnDeliveryStream.AmazonopensearchserviceDestinationConfigurationProperty.ClusterEndpoint`.
	ClusterEndpoint *string `json:"clusterEndpoint"`
	// `CfnDeliveryStream.AmazonopensearchserviceDestinationConfigurationProperty.DomainARN`.
	DomainArn *string `json:"domainArn"`
	// `CfnDeliveryStream.AmazonopensearchserviceDestinationConfigurationProperty.IndexName`.
	IndexName *string `json:"indexName"`
	// `CfnDeliveryStream.AmazonopensearchserviceDestinationConfigurationProperty.IndexRotationPeriod`.
	IndexRotationPeriod *string `json:"indexRotationPeriod"`
	// `CfnDeliveryStream.AmazonopensearchserviceDestinationConfigurationProperty.ProcessingConfiguration`.
	ProcessingConfiguration interface{} `json:"processingConfiguration"`
	// `CfnDeliveryStream.AmazonopensearchserviceDestinationConfigurationProperty.RetryOptions`.
	RetryOptions interface{} `json:"retryOptions"`
	// `CfnDeliveryStream.AmazonopensearchserviceDestinationConfigurationProperty.RoleARN`.
	RoleArn *string `json:"roleArn"`
	// `CfnDeliveryStream.AmazonopensearchserviceDestinationConfigurationProperty.S3BackupMode`.
	S3BackupMode *string `json:"s3BackupMode"`
	// `CfnDeliveryStream.AmazonopensearchserviceDestinationConfigurationProperty.S3Configuration`.
	S3Configuration interface{} `json:"s3Configuration"`
	// `CfnDeliveryStream.AmazonopensearchserviceDestinationConfigurationProperty.TypeName`.
	TypeName *string `json:"typeName"`
	// `CfnDeliveryStream.AmazonopensearchserviceDestinationConfigurationProperty.VpcConfiguration`.
	VpcConfiguration interface{} `json:"vpcConfiguration"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_AmazonopensearchserviceRetryOptionsProperty struct {
	// `CfnDeliveryStream.AmazonopensearchserviceRetryOptionsProperty.DurationInSeconds`.
	DurationInSeconds *float64 `json:"durationInSeconds"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_BufferingHintsProperty struct {
	// `CfnDeliveryStream.BufferingHintsProperty.IntervalInSeconds`.
	IntervalInSeconds *float64 `json:"intervalInSeconds"`
	// `CfnDeliveryStream.BufferingHintsProperty.SizeInMBs`.
	SizeInMBs *float64 `json:"sizeInMBs"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_CloudWatchLoggingOptionsProperty struct {
	// `CfnDeliveryStream.CloudWatchLoggingOptionsProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `CfnDeliveryStream.CloudWatchLoggingOptionsProperty.LogGroupName`.
	LogGroupName *string `json:"logGroupName"`
	// `CfnDeliveryStream.CloudWatchLoggingOptionsProperty.LogStreamName`.
	LogStreamName *string `json:"logStreamName"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_CopyCommandProperty struct {
	// `CfnDeliveryStream.CopyCommandProperty.CopyOptions`.
	CopyOptions *string `json:"copyOptions"`
	// `CfnDeliveryStream.CopyCommandProperty.DataTableColumns`.
	DataTableColumns *string `json:"dataTableColumns"`
	// `CfnDeliveryStream.CopyCommandProperty.DataTableName`.
	DataTableName *string `json:"dataTableName"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_DataFormatConversionConfigurationProperty struct {
	// `CfnDeliveryStream.DataFormatConversionConfigurationProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `CfnDeliveryStream.DataFormatConversionConfigurationProperty.InputFormatConfiguration`.
	InputFormatConfiguration interface{} `json:"inputFormatConfiguration"`
	// `CfnDeliveryStream.DataFormatConversionConfigurationProperty.OutputFormatConfiguration`.
	OutputFormatConfiguration interface{} `json:"outputFormatConfiguration"`
	// `CfnDeliveryStream.DataFormatConversionConfigurationProperty.SchemaConfiguration`.
	SchemaConfiguration interface{} `json:"schemaConfiguration"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_DeliveryStreamEncryptionConfigurationInputProperty struct {
	// `CfnDeliveryStream.DeliveryStreamEncryptionConfigurationInputProperty.KeyARN`.
	KeyArn *string `json:"keyArn"`
	// `CfnDeliveryStream.DeliveryStreamEncryptionConfigurationInputProperty.KeyType`.
	KeyType *string `json:"keyType"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_DeserializerProperty struct {
	// `CfnDeliveryStream.DeserializerProperty.HiveJsonSerDe`.
	HiveJsonSerDe interface{} `json:"hiveJsonSerDe"`
	// `CfnDeliveryStream.DeserializerProperty.OpenXJsonSerDe`.
	OpenXJsonSerDe interface{} `json:"openXJsonSerDe"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_DynamicPartitioningConfigurationProperty struct {
	// `CfnDeliveryStream.DynamicPartitioningConfigurationProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `CfnDeliveryStream.DynamicPartitioningConfigurationProperty.RetryOptions`.
	RetryOptions interface{} `json:"retryOptions"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_ElasticsearchBufferingHintsProperty struct {
	// `CfnDeliveryStream.ElasticsearchBufferingHintsProperty.IntervalInSeconds`.
	IntervalInSeconds *float64 `json:"intervalInSeconds"`
	// `CfnDeliveryStream.ElasticsearchBufferingHintsProperty.SizeInMBs`.
	SizeInMBs *float64 `json:"sizeInMBs"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_ElasticsearchDestinationConfigurationProperty struct {
	// `CfnDeliveryStream.ElasticsearchDestinationConfigurationProperty.BufferingHints`.
	BufferingHints interface{} `json:"bufferingHints"`
	// `CfnDeliveryStream.ElasticsearchDestinationConfigurationProperty.CloudWatchLoggingOptions`.
	CloudWatchLoggingOptions interface{} `json:"cloudWatchLoggingOptions"`
	// `CfnDeliveryStream.ElasticsearchDestinationConfigurationProperty.ClusterEndpoint`.
	ClusterEndpoint *string `json:"clusterEndpoint"`
	// `CfnDeliveryStream.ElasticsearchDestinationConfigurationProperty.DomainARN`.
	DomainArn *string `json:"domainArn"`
	// `CfnDeliveryStream.ElasticsearchDestinationConfigurationProperty.IndexName`.
	IndexName *string `json:"indexName"`
	// `CfnDeliveryStream.ElasticsearchDestinationConfigurationProperty.IndexRotationPeriod`.
	IndexRotationPeriod *string `json:"indexRotationPeriod"`
	// `CfnDeliveryStream.ElasticsearchDestinationConfigurationProperty.ProcessingConfiguration`.
	ProcessingConfiguration interface{} `json:"processingConfiguration"`
	// `CfnDeliveryStream.ElasticsearchDestinationConfigurationProperty.RetryOptions`.
	RetryOptions interface{} `json:"retryOptions"`
	// `CfnDeliveryStream.ElasticsearchDestinationConfigurationProperty.RoleARN`.
	RoleArn *string `json:"roleArn"`
	// `CfnDeliveryStream.ElasticsearchDestinationConfigurationProperty.S3BackupMode`.
	S3BackupMode *string `json:"s3BackupMode"`
	// `CfnDeliveryStream.ElasticsearchDestinationConfigurationProperty.S3Configuration`.
	S3Configuration interface{} `json:"s3Configuration"`
	// `CfnDeliveryStream.ElasticsearchDestinationConfigurationProperty.TypeName`.
	TypeName *string `json:"typeName"`
	// `CfnDeliveryStream.ElasticsearchDestinationConfigurationProperty.VpcConfiguration`.
	VpcConfiguration interface{} `json:"vpcConfiguration"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_ElasticsearchRetryOptionsProperty struct {
	// `CfnDeliveryStream.ElasticsearchRetryOptionsProperty.DurationInSeconds`.
	DurationInSeconds *float64 `json:"durationInSeconds"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_EncryptionConfigurationProperty struct {
	// `CfnDeliveryStream.EncryptionConfigurationProperty.KMSEncryptionConfig`.
	KmsEncryptionConfig interface{} `json:"kmsEncryptionConfig"`
	// `CfnDeliveryStream.EncryptionConfigurationProperty.NoEncryptionConfig`.
	NoEncryptionConfig *string `json:"noEncryptionConfig"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_ExtendedS3DestinationConfigurationProperty struct {
	// `CfnDeliveryStream.ExtendedS3DestinationConfigurationProperty.BucketARN`.
	BucketArn *string `json:"bucketArn"`
	// `CfnDeliveryStream.ExtendedS3DestinationConfigurationProperty.BufferingHints`.
	BufferingHints interface{} `json:"bufferingHints"`
	// `CfnDeliveryStream.ExtendedS3DestinationConfigurationProperty.CloudWatchLoggingOptions`.
	CloudWatchLoggingOptions interface{} `json:"cloudWatchLoggingOptions"`
	// `CfnDeliveryStream.ExtendedS3DestinationConfigurationProperty.CompressionFormat`.
	CompressionFormat *string `json:"compressionFormat"`
	// `CfnDeliveryStream.ExtendedS3DestinationConfigurationProperty.DataFormatConversionConfiguration`.
	DataFormatConversionConfiguration interface{} `json:"dataFormatConversionConfiguration"`
	// `CfnDeliveryStream.ExtendedS3DestinationConfigurationProperty.DynamicPartitioningConfiguration`.
	DynamicPartitioningConfiguration interface{} `json:"dynamicPartitioningConfiguration"`
	// `CfnDeliveryStream.ExtendedS3DestinationConfigurationProperty.EncryptionConfiguration`.
	EncryptionConfiguration interface{} `json:"encryptionConfiguration"`
	// `CfnDeliveryStream.ExtendedS3DestinationConfigurationProperty.ErrorOutputPrefix`.
	ErrorOutputPrefix *string `json:"errorOutputPrefix"`
	// `CfnDeliveryStream.ExtendedS3DestinationConfigurationProperty.Prefix`.
	Prefix *string `json:"prefix"`
	// `CfnDeliveryStream.ExtendedS3DestinationConfigurationProperty.ProcessingConfiguration`.
	ProcessingConfiguration interface{} `json:"processingConfiguration"`
	// `CfnDeliveryStream.ExtendedS3DestinationConfigurationProperty.RoleARN`.
	RoleArn *string `json:"roleArn"`
	// `CfnDeliveryStream.ExtendedS3DestinationConfigurationProperty.S3BackupConfiguration`.
	S3BackupConfiguration interface{} `json:"s3BackupConfiguration"`
	// `CfnDeliveryStream.ExtendedS3DestinationConfigurationProperty.S3BackupMode`.
	S3BackupMode *string `json:"s3BackupMode"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_HiveJsonSerDeProperty struct {
	// `CfnDeliveryStream.HiveJsonSerDeProperty.TimestampFormats`.
	TimestampFormats *[]*string `json:"timestampFormats"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_HttpEndpointCommonAttributeProperty struct {
	// `CfnDeliveryStream.HttpEndpointCommonAttributeProperty.AttributeName`.
	AttributeName *string `json:"attributeName"`
	// `CfnDeliveryStream.HttpEndpointCommonAttributeProperty.AttributeValue`.
	AttributeValue *string `json:"attributeValue"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_HttpEndpointConfigurationProperty struct {
	// `CfnDeliveryStream.HttpEndpointConfigurationProperty.AccessKey`.
	AccessKey *string `json:"accessKey"`
	// `CfnDeliveryStream.HttpEndpointConfigurationProperty.Name`.
	Name *string `json:"name"`
	// `CfnDeliveryStream.HttpEndpointConfigurationProperty.Url`.
	Url *string `json:"url"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_HttpEndpointDestinationConfigurationProperty struct {
	// `CfnDeliveryStream.HttpEndpointDestinationConfigurationProperty.BufferingHints`.
	BufferingHints interface{} `json:"bufferingHints"`
	// `CfnDeliveryStream.HttpEndpointDestinationConfigurationProperty.CloudWatchLoggingOptions`.
	CloudWatchLoggingOptions interface{} `json:"cloudWatchLoggingOptions"`
	// `CfnDeliveryStream.HttpEndpointDestinationConfigurationProperty.EndpointConfiguration`.
	EndpointConfiguration interface{} `json:"endpointConfiguration"`
	// `CfnDeliveryStream.HttpEndpointDestinationConfigurationProperty.ProcessingConfiguration`.
	ProcessingConfiguration interface{} `json:"processingConfiguration"`
	// `CfnDeliveryStream.HttpEndpointDestinationConfigurationProperty.RequestConfiguration`.
	RequestConfiguration interface{} `json:"requestConfiguration"`
	// `CfnDeliveryStream.HttpEndpointDestinationConfigurationProperty.RetryOptions`.
	RetryOptions interface{} `json:"retryOptions"`
	// `CfnDeliveryStream.HttpEndpointDestinationConfigurationProperty.RoleARN`.
	RoleArn *string `json:"roleArn"`
	// `CfnDeliveryStream.HttpEndpointDestinationConfigurationProperty.S3BackupMode`.
	S3BackupMode *string `json:"s3BackupMode"`
	// `CfnDeliveryStream.HttpEndpointDestinationConfigurationProperty.S3Configuration`.
	S3Configuration interface{} `json:"s3Configuration"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_HttpEndpointRequestConfigurationProperty struct {
	// `CfnDeliveryStream.HttpEndpointRequestConfigurationProperty.CommonAttributes`.
	CommonAttributes interface{} `json:"commonAttributes"`
	// `CfnDeliveryStream.HttpEndpointRequestConfigurationProperty.ContentEncoding`.
	ContentEncoding *string `json:"contentEncoding"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_InputFormatConfigurationProperty struct {
	// `CfnDeliveryStream.InputFormatConfigurationProperty.Deserializer`.
	Deserializer interface{} `json:"deserializer"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_KMSEncryptionConfigProperty struct {
	// `CfnDeliveryStream.KMSEncryptionConfigProperty.AWSKMSKeyARN`.
	AwskmsKeyArn *string `json:"awskmsKeyArn"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_KinesisStreamSourceConfigurationProperty struct {
	// `CfnDeliveryStream.KinesisStreamSourceConfigurationProperty.KinesisStreamARN`.
	KinesisStreamArn *string `json:"kinesisStreamArn"`
	// `CfnDeliveryStream.KinesisStreamSourceConfigurationProperty.RoleARN`.
	RoleArn *string `json:"roleArn"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_OpenXJsonSerDeProperty struct {
	// `CfnDeliveryStream.OpenXJsonSerDeProperty.CaseInsensitive`.
	CaseInsensitive interface{} `json:"caseInsensitive"`
	// `CfnDeliveryStream.OpenXJsonSerDeProperty.ColumnToJsonKeyMappings`.
	ColumnToJsonKeyMappings interface{} `json:"columnToJsonKeyMappings"`
	// `CfnDeliveryStream.OpenXJsonSerDeProperty.ConvertDotsInJsonKeysToUnderscores`.
	ConvertDotsInJsonKeysToUnderscores interface{} `json:"convertDotsInJsonKeysToUnderscores"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_OrcSerDeProperty struct {
	// `CfnDeliveryStream.OrcSerDeProperty.BlockSizeBytes`.
	BlockSizeBytes *float64 `json:"blockSizeBytes"`
	// `CfnDeliveryStream.OrcSerDeProperty.BloomFilterColumns`.
	BloomFilterColumns *[]*string `json:"bloomFilterColumns"`
	// `CfnDeliveryStream.OrcSerDeProperty.BloomFilterFalsePositiveProbability`.
	BloomFilterFalsePositiveProbability *float64 `json:"bloomFilterFalsePositiveProbability"`
	// `CfnDeliveryStream.OrcSerDeProperty.Compression`.
	Compression *string `json:"compression"`
	// `CfnDeliveryStream.OrcSerDeProperty.DictionaryKeyThreshold`.
	DictionaryKeyThreshold *float64 `json:"dictionaryKeyThreshold"`
	// `CfnDeliveryStream.OrcSerDeProperty.EnablePadding`.
	EnablePadding interface{} `json:"enablePadding"`
	// `CfnDeliveryStream.OrcSerDeProperty.FormatVersion`.
	FormatVersion *string `json:"formatVersion"`
	// `CfnDeliveryStream.OrcSerDeProperty.PaddingTolerance`.
	PaddingTolerance *float64 `json:"paddingTolerance"`
	// `CfnDeliveryStream.OrcSerDeProperty.RowIndexStride`.
	RowIndexStride *float64 `json:"rowIndexStride"`
	// `CfnDeliveryStream.OrcSerDeProperty.StripeSizeBytes`.
	StripeSizeBytes *float64 `json:"stripeSizeBytes"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_OutputFormatConfigurationProperty struct {
	// `CfnDeliveryStream.OutputFormatConfigurationProperty.Serializer`.
	Serializer interface{} `json:"serializer"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_ParquetSerDeProperty struct {
	// `CfnDeliveryStream.ParquetSerDeProperty.BlockSizeBytes`.
	BlockSizeBytes *float64 `json:"blockSizeBytes"`
	// `CfnDeliveryStream.ParquetSerDeProperty.Compression`.
	Compression *string `json:"compression"`
	// `CfnDeliveryStream.ParquetSerDeProperty.EnableDictionaryCompression`.
	EnableDictionaryCompression interface{} `json:"enableDictionaryCompression"`
	// `CfnDeliveryStream.ParquetSerDeProperty.MaxPaddingBytes`.
	MaxPaddingBytes *float64 `json:"maxPaddingBytes"`
	// `CfnDeliveryStream.ParquetSerDeProperty.PageSizeBytes`.
	PageSizeBytes *float64 `json:"pageSizeBytes"`
	// `CfnDeliveryStream.ParquetSerDeProperty.WriterVersion`.
	WriterVersion *string `json:"writerVersion"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_ProcessingConfigurationProperty struct {
	// `CfnDeliveryStream.ProcessingConfigurationProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `CfnDeliveryStream.ProcessingConfigurationProperty.Processors`.
	Processors interface{} `json:"processors"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_ProcessorParameterProperty struct {
	// `CfnDeliveryStream.ProcessorParameterProperty.ParameterName`.
	ParameterName *string `json:"parameterName"`
	// `CfnDeliveryStream.ProcessorParameterProperty.ParameterValue`.
	ParameterValue *string `json:"parameterValue"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_ProcessorProperty struct {
	// `CfnDeliveryStream.ProcessorProperty.Parameters`.
	Parameters interface{} `json:"parameters"`
	// `CfnDeliveryStream.ProcessorProperty.Type`.
	Type *string `json:"type"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_RedshiftDestinationConfigurationProperty struct {
	// `CfnDeliveryStream.RedshiftDestinationConfigurationProperty.CloudWatchLoggingOptions`.
	CloudWatchLoggingOptions interface{} `json:"cloudWatchLoggingOptions"`
	// `CfnDeliveryStream.RedshiftDestinationConfigurationProperty.ClusterJDBCURL`.
	ClusterJdbcurl *string `json:"clusterJdbcurl"`
	// `CfnDeliveryStream.RedshiftDestinationConfigurationProperty.CopyCommand`.
	CopyCommand interface{} `json:"copyCommand"`
	// `CfnDeliveryStream.RedshiftDestinationConfigurationProperty.Password`.
	Password *string `json:"password"`
	// `CfnDeliveryStream.RedshiftDestinationConfigurationProperty.ProcessingConfiguration`.
	ProcessingConfiguration interface{} `json:"processingConfiguration"`
	// `CfnDeliveryStream.RedshiftDestinationConfigurationProperty.RetryOptions`.
	RetryOptions interface{} `json:"retryOptions"`
	// `CfnDeliveryStream.RedshiftDestinationConfigurationProperty.RoleARN`.
	RoleArn *string `json:"roleArn"`
	// `CfnDeliveryStream.RedshiftDestinationConfigurationProperty.S3BackupConfiguration`.
	S3BackupConfiguration interface{} `json:"s3BackupConfiguration"`
	// `CfnDeliveryStream.RedshiftDestinationConfigurationProperty.S3BackupMode`.
	S3BackupMode *string `json:"s3BackupMode"`
	// `CfnDeliveryStream.RedshiftDestinationConfigurationProperty.S3Configuration`.
	S3Configuration interface{} `json:"s3Configuration"`
	// `CfnDeliveryStream.RedshiftDestinationConfigurationProperty.Username`.
	Username *string `json:"username"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_RedshiftRetryOptionsProperty struct {
	// `CfnDeliveryStream.RedshiftRetryOptionsProperty.DurationInSeconds`.
	DurationInSeconds *float64 `json:"durationInSeconds"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_RetryOptionsProperty struct {
	// `CfnDeliveryStream.RetryOptionsProperty.DurationInSeconds`.
	DurationInSeconds *float64 `json:"durationInSeconds"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_S3DestinationConfigurationProperty struct {
	// `CfnDeliveryStream.S3DestinationConfigurationProperty.BucketARN`.
	BucketArn *string `json:"bucketArn"`
	// `CfnDeliveryStream.S3DestinationConfigurationProperty.BufferingHints`.
	BufferingHints interface{} `json:"bufferingHints"`
	// `CfnDeliveryStream.S3DestinationConfigurationProperty.CloudWatchLoggingOptions`.
	CloudWatchLoggingOptions interface{} `json:"cloudWatchLoggingOptions"`
	// `CfnDeliveryStream.S3DestinationConfigurationProperty.CompressionFormat`.
	CompressionFormat *string `json:"compressionFormat"`
	// `CfnDeliveryStream.S3DestinationConfigurationProperty.EncryptionConfiguration`.
	EncryptionConfiguration interface{} `json:"encryptionConfiguration"`
	// `CfnDeliveryStream.S3DestinationConfigurationProperty.ErrorOutputPrefix`.
	ErrorOutputPrefix *string `json:"errorOutputPrefix"`
	// `CfnDeliveryStream.S3DestinationConfigurationProperty.Prefix`.
	Prefix *string `json:"prefix"`
	// `CfnDeliveryStream.S3DestinationConfigurationProperty.RoleARN`.
	RoleArn *string `json:"roleArn"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_SchemaConfigurationProperty struct {
	// `CfnDeliveryStream.SchemaConfigurationProperty.CatalogId`.
	CatalogId *string `json:"catalogId"`
	// `CfnDeliveryStream.SchemaConfigurationProperty.DatabaseName`.
	DatabaseName *string `json:"databaseName"`
	// `CfnDeliveryStream.SchemaConfigurationProperty.Region`.
	Region *string `json:"region"`
	// `CfnDeliveryStream.SchemaConfigurationProperty.RoleARN`.
	RoleArn *string `json:"roleArn"`
	// `CfnDeliveryStream.SchemaConfigurationProperty.TableName`.
	TableName *string `json:"tableName"`
	// `CfnDeliveryStream.SchemaConfigurationProperty.VersionId`.
	VersionId *string `json:"versionId"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_SerializerProperty struct {
	// `CfnDeliveryStream.SerializerProperty.OrcSerDe`.
	OrcSerDe interface{} `json:"orcSerDe"`
	// `CfnDeliveryStream.SerializerProperty.ParquetSerDe`.
	ParquetSerDe interface{} `json:"parquetSerDe"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_SplunkDestinationConfigurationProperty struct {
	// `CfnDeliveryStream.SplunkDestinationConfigurationProperty.CloudWatchLoggingOptions`.
	CloudWatchLoggingOptions interface{} `json:"cloudWatchLoggingOptions"`
	// `CfnDeliveryStream.SplunkDestinationConfigurationProperty.HECAcknowledgmentTimeoutInSeconds`.
	HecAcknowledgmentTimeoutInSeconds *float64 `json:"hecAcknowledgmentTimeoutInSeconds"`
	// `CfnDeliveryStream.SplunkDestinationConfigurationProperty.HECEndpoint`.
	HecEndpoint *string `json:"hecEndpoint"`
	// `CfnDeliveryStream.SplunkDestinationConfigurationProperty.HECEndpointType`.
	HecEndpointType *string `json:"hecEndpointType"`
	// `CfnDeliveryStream.SplunkDestinationConfigurationProperty.HECToken`.
	HecToken *string `json:"hecToken"`
	// `CfnDeliveryStream.SplunkDestinationConfigurationProperty.ProcessingConfiguration`.
	ProcessingConfiguration interface{} `json:"processingConfiguration"`
	// `CfnDeliveryStream.SplunkDestinationConfigurationProperty.RetryOptions`.
	RetryOptions interface{} `json:"retryOptions"`
	// `CfnDeliveryStream.SplunkDestinationConfigurationProperty.S3BackupMode`.
	S3BackupMode *string `json:"s3BackupMode"`
	// `CfnDeliveryStream.SplunkDestinationConfigurationProperty.S3Configuration`.
	S3Configuration interface{} `json:"s3Configuration"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_SplunkRetryOptionsProperty struct {
	// `CfnDeliveryStream.SplunkRetryOptionsProperty.DurationInSeconds`.
	DurationInSeconds *float64 `json:"durationInSeconds"`
}

// TODO: EXAMPLE
//
type CfnDeliveryStream_VpcConfigurationProperty struct {
	// `CfnDeliveryStream.VpcConfigurationProperty.RoleARN`.
	RoleArn *string `json:"roleArn"`
	// `CfnDeliveryStream.VpcConfigurationProperty.SecurityGroupIds`.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// `CfnDeliveryStream.VpcConfigurationProperty.SubnetIds`.
	SubnetIds *[]*string `json:"subnetIds"`
}

// Properties for defining a `AWS::KinesisFirehose::DeliveryStream`.
//
// TODO: EXAMPLE
//
type CfnDeliveryStreamProps struct {
	// `AWS::KinesisFirehose::DeliveryStream.AmazonopensearchserviceDestinationConfiguration`.
	AmazonopensearchserviceDestinationConfiguration interface{} `json:"amazonopensearchserviceDestinationConfiguration"`
	// `AWS::KinesisFirehose::DeliveryStream.DeliveryStreamEncryptionConfigurationInput`.
	DeliveryStreamEncryptionConfigurationInput interface{} `json:"deliveryStreamEncryptionConfigurationInput"`
	// `AWS::KinesisFirehose::DeliveryStream.DeliveryStreamName`.
	DeliveryStreamName *string `json:"deliveryStreamName"`
	// `AWS::KinesisFirehose::DeliveryStream.DeliveryStreamType`.
	DeliveryStreamType *string `json:"deliveryStreamType"`
	// `AWS::KinesisFirehose::DeliveryStream.ElasticsearchDestinationConfiguration`.
	ElasticsearchDestinationConfiguration interface{} `json:"elasticsearchDestinationConfiguration"`
	// `AWS::KinesisFirehose::DeliveryStream.ExtendedS3DestinationConfiguration`.
	ExtendedS3DestinationConfiguration interface{} `json:"extendedS3DestinationConfiguration"`
	// `AWS::KinesisFirehose::DeliveryStream.HttpEndpointDestinationConfiguration`.
	HttpEndpointDestinationConfiguration interface{} `json:"httpEndpointDestinationConfiguration"`
	// `AWS::KinesisFirehose::DeliveryStream.KinesisStreamSourceConfiguration`.
	KinesisStreamSourceConfiguration interface{} `json:"kinesisStreamSourceConfiguration"`
	// `AWS::KinesisFirehose::DeliveryStream.RedshiftDestinationConfiguration`.
	RedshiftDestinationConfiguration interface{} `json:"redshiftDestinationConfiguration"`
	// `AWS::KinesisFirehose::DeliveryStream.S3DestinationConfiguration`.
	S3DestinationConfiguration interface{} `json:"s3DestinationConfiguration"`
	// `AWS::KinesisFirehose::DeliveryStream.SplunkDestinationConfiguration`.
	SplunkDestinationConfiguration interface{} `json:"splunkDestinationConfiguration"`
	// `AWS::KinesisFirehose::DeliveryStream.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// Options when binding a DataProcessor to a delivery stream destination.
//
// TODO: EXAMPLE
//
// Experimental.
type DataProcessorBindOptions struct {
	// The IAM role assumed by Kinesis Data Firehose to write to the destination that this DataProcessor will bind to.
	// Experimental.
	Role awsiam.IRole `json:"role"`
}

// The full configuration of a data processor.
//
// TODO: EXAMPLE
//
// Experimental.
type DataProcessorConfig struct {
	// The key-value pair that identifies the underlying processor resource.
	// Experimental.
	ProcessorIdentifier *DataProcessorIdentifier `json:"processorIdentifier"`
	// The type of the underlying processor resource.
	//
	// Must be an accepted value in `CfnDeliveryStream.ProcessorProperty.Type`.
	//
	// TODO: EXAMPLE
	//
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processor.html#cfn-kinesisfirehose-deliverystream-processor-type
	//
	// Experimental.
	ProcessorType *string `json:"processorType"`
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
	ParameterName *string `json:"parameterName"`
	// The identifier of the underlying processor resource.
	// Experimental.
	ParameterValue *string `json:"parameterValue"`
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
	BufferInterval awscdk.Duration `json:"bufferInterval"`
	// The amount of incoming data Kinesis Data Firehose will buffer before calling the processor.
	// Experimental.
	BufferSize awscdk.Size `json:"bufferSize"`
	// The number of times Kinesis Data Firehose will retry the processor invocation after a failure due to network timeout or invocation limits.
	// Experimental.
	Retries *float64 `json:"retries"`
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
	DeliveryStreamArn *string `json:"deliveryStreamArn"`
	// The name of the delivery stream.
	//
	// At least one of deliveryStreamName and deliveryStreamArn  must be provided.
	// Experimental.
	DeliveryStreamName *string `json:"deliveryStreamName"`
	// The IAM role associated with this delivery stream.
	//
	// Assumed by Kinesis Data Firehose to read from sources and encrypt data server-side.
	// Experimental.
	Role awsiam.IRole `json:"role"`
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
	Destinations *[]IDestination `json:"destinations"`
	// A name for the delivery stream.
	// Experimental.
	DeliveryStreamName *string `json:"deliveryStreamName"`
	// Indicates the type of customer master key (CMK) to use for server-side encryption, if any.
	// Experimental.
	Encryption StreamEncryption `json:"encryption"`
	// Customer managed key to server-side encrypt data in the stream.
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey"`
	// The IAM role associated with this delivery stream.
	//
	// Assumed by Kinesis Data Firehose to read from sources and encrypt data server-side.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// The Kinesis data stream to use as a source for this delivery stream.
	// Experimental.
	SourceStream awskinesis.IStream `json:"sourceStream"`
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
	Dependables *[]awscdk.IDependable `json:"dependables"`
	// S3 destination configuration properties.
	// Experimental.
	ExtendedS3DestinationConfiguration *CfnDeliveryStream_ExtendedS3DestinationConfigurationProperty `json:"extendedS3DestinationConfiguration"`
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


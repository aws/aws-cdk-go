package awsiotanalytics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotanalytics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::IoTAnalytics::Channel`.
//
// The AWS::IoTAnalytics::Channel resource collects data from an MQTT topic and archives the raw, unprocessed messages before publishing the data to a pipeline. For more information, see [How to Use AWS IoT Analytics](https://docs.aws.amazon.com/iotanalytics/latest/userguide/welcome.html#aws-iot-analytics-how) in the *AWS IoT Analytics User Guide* .
//
// TODO: EXAMPLE
//
type CfnChannel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ChannelName() *string
	SetChannelName(val *string)
	ChannelStorage() interface{}
	SetChannelStorage(val interface{})
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	RetentionPeriod() interface{}
	SetRetentionPeriod(val interface{})
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

// The jsii proxy struct for CfnChannel
type jsiiProxy_CfnChannel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnChannel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) ChannelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) ChannelStorage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"channelStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) RetentionPeriod() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTAnalytics::Channel`.
func NewCfnChannel(scope constructs.Construct, id *string, props *CfnChannelProps) CfnChannel {
	_init_.Initialize()

	j := jsiiProxy_CfnChannel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iotanalytics.CfnChannel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTAnalytics::Channel`.
func NewCfnChannel_Override(c CfnChannel, scope constructs.Construct, id *string, props *CfnChannelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iotanalytics.CfnChannel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnChannel) SetChannelName(val *string) {
	_jsii_.Set(
		j,
		"channelName",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetChannelStorage(val interface{}) {
	_jsii_.Set(
		j,
		"channelStorage",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetRetentionPeriod(val interface{}) {
	_jsii_.Set(
		j,
		"retentionPeriod",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnChannel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotanalytics.CfnChannel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnChannel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotanalytics.CfnChannel",
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
func CfnChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotanalytics.CfnChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnChannel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iotanalytics.CfnChannel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnChannel) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnChannel) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnChannel) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnChannel) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnChannel) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnChannel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnChannel) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnChannel) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnChannel) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnChannel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnChannel) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnChannel) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Where channel data is stored.
//
// You may choose one of `serviceManagedS3` , `customerManagedS3` storage. If not specified, the default is `serviceManagedS3` . This can't be changed after creation of the channel.
//
// TODO: EXAMPLE
//
type CfnChannel_ChannelStorageProperty struct {
	// Used to store channel data in an S3 bucket that you manage.
	//
	// If customer managed storage is selected, the `retentionPeriod` parameter is ignored. You can't change the choice of S3 storage after the data store is created.
	CustomerManagedS3 interface{} `json:"customerManagedS3"`
	// Used to store channel data in an S3 bucket managed by AWS IoT Analytics .
	//
	// You can't change the choice of S3 storage after the data store is created.
	ServiceManagedS3 interface{} `json:"serviceManagedS3"`
}

// Used to store channel data in an S3 bucket that you manage.
//
// TODO: EXAMPLE
//
type CfnChannel_CustomerManagedS3Property struct {
	// The name of the S3 bucket in which channel data is stored.
	Bucket *string `json:"bucket"`
	// The ARN of the role that grants AWS IoT Analytics permission to interact with your Amazon S3 resources.
	RoleArn *string `json:"roleArn"`
	// (Optional) The prefix used to create the keys of the channel data objects.
	//
	// Each object in an S3 bucket has a key that is its unique identifier within the bucket (each object in a bucket has exactly one key). The prefix must end with a forward slash (/).
	KeyPrefix *string `json:"keyPrefix"`
}

// How long, in days, message data is kept.
//
// TODO: EXAMPLE
//
type CfnChannel_RetentionPeriodProperty struct {
	// The number of days that message data is kept.
	//
	// The `unlimited` parameter must be false.
	NumberOfDays *float64 `json:"numberOfDays"`
	// If true, message data is kept indefinitely.
	Unlimited interface{} `json:"unlimited"`
}

// Used to store channel data in an S3 bucket managed by AWS IoT Analytics .
//
// You can't change the choice of S3 storage after the data store is created.
//
// TODO: EXAMPLE
//
type CfnChannel_ServiceManagedS3Property struct {
}

// Properties for defining a `CfnChannel`.
//
// TODO: EXAMPLE
//
type CfnChannelProps struct {
	// The name of the channel.
	ChannelName *string `json:"channelName"`
	// Where channel data is stored.
	ChannelStorage interface{} `json:"channelStorage"`
	// How long, in days, message data is kept for the channel.
	RetentionPeriod interface{} `json:"retentionPeriod"`
	// Metadata which can be used to manage the channel.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::IoTAnalytics::Dataset`.
//
// The AWS::IoTAnalytics::Dataset resource stores data retrieved from a data store by applying a `queryAction` (an SQL query) or a `containerAction` (executing a containerized application). The data set can be populated manually by calling `CreateDatasetContent` or automatically according to a `trigger` you specify. For more information, see [How to Use AWS IoT Analytics](https://docs.aws.amazon.com/iotanalytics/latest/userguide/welcome.html#aws-iot-analytics-how) in the *AWS IoT Analytics User Guide* .
//
// TODO: EXAMPLE
//
type CfnDataset interface {
	awscdk.CfnResource
	awscdk.IInspectable
	Actions() interface{}
	SetActions(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ContentDeliveryRules() interface{}
	SetContentDeliveryRules(val interface{})
	CreationStack() *[]*string
	DatasetName() *string
	SetDatasetName(val *string)
	LateDataRules() interface{}
	SetLateDataRules(val interface{})
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	RetentionPeriod() interface{}
	SetRetentionPeriod(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	Triggers() interface{}
	SetTriggers(val interface{})
	UpdatedProperites() *map[string]interface{}
	VersioningConfiguration() interface{}
	SetVersioningConfiguration(val interface{})
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

// The jsii proxy struct for CfnDataset
type jsiiProxy_CfnDataset struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDataset) Actions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) ContentDeliveryRules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentDeliveryRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) DatasetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) LateDataRules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lateDataRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) RetentionPeriod() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) Triggers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) VersioningConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versioningConfiguration",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTAnalytics::Dataset`.
func NewCfnDataset(scope constructs.Construct, id *string, props *CfnDatasetProps) CfnDataset {
	_init_.Initialize()

	j := jsiiProxy_CfnDataset{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iotanalytics.CfnDataset",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTAnalytics::Dataset`.
func NewCfnDataset_Override(c CfnDataset, scope constructs.Construct, id *string, props *CfnDatasetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iotanalytics.CfnDataset",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDataset) SetActions(val interface{}) {
	_jsii_.Set(
		j,
		"actions",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetContentDeliveryRules(val interface{}) {
	_jsii_.Set(
		j,
		"contentDeliveryRules",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetDatasetName(val *string) {
	_jsii_.Set(
		j,
		"datasetName",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetLateDataRules(val interface{}) {
	_jsii_.Set(
		j,
		"lateDataRules",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetRetentionPeriod(val interface{}) {
	_jsii_.Set(
		j,
		"retentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetTriggers(val interface{}) {
	_jsii_.Set(
		j,
		"triggers",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetVersioningConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"versioningConfiguration",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDataset_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotanalytics.CfnDataset",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDataset_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotanalytics.CfnDataset",
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
func CfnDataset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotanalytics.CfnDataset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataset_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iotanalytics.CfnDataset",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnDataset) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnDataset) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnDataset) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnDataset) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnDataset) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnDataset) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnDataset) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnDataset) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnDataset) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnDataset) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnDataset) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDataset) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnDataset) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnDataset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataset) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Information needed to run the "containerAction" to produce data set contents.
//
// TODO: EXAMPLE
//
type CfnDataset_ActionProperty struct {
	// The name of the data set action by which data set contents are automatically created.
	ActionName *string `json:"actionName"`
	// Information which allows the system to run a containerized application in order to create the data set contents.
	//
	// The application must be in a Docker container along with any needed support libraries.
	ContainerAction interface{} `json:"containerAction"`
	// An "SqlQueryDatasetAction" object that uses an SQL query to automatically create data set contents.
	QueryAction interface{} `json:"queryAction"`
}

// Information needed to run the "containerAction" to produce data set contents.
//
// TODO: EXAMPLE
//
type CfnDataset_ContainerActionProperty struct {
	// The ARN of the role which gives permission to the system to access needed resources in order to run the "containerAction".
	//
	// This includes, at minimum, permission to retrieve the data set contents which are the input to the containerized application.
	ExecutionRoleArn *string `json:"executionRoleArn"`
	// The ARN of the Docker container stored in your account.
	//
	// The Docker container contains an application and needed support libraries and is used to generate data set contents.
	Image *string `json:"image"`
	// Configuration of the resource which executes the "containerAction".
	ResourceConfiguration interface{} `json:"resourceConfiguration"`
	// The values of variables used within the context of the execution of the containerized application (basically, parameters passed to the application).
	//
	// Each variable must have a name and a value given by one of "stringValue", "datasetContentVersionValue", or "outputFileUriValue".
	Variables interface{} `json:"variables"`
}

// The destination to which dataset contents are delivered.
//
// TODO: EXAMPLE
//
type CfnDataset_DatasetContentDeliveryRuleDestinationProperty struct {
	// Configuration information for delivery of dataset contents to AWS IoT Events .
	IotEventsDestinationConfiguration interface{} `json:"iotEventsDestinationConfiguration"`
	// Configuration information for delivery of dataset contents to Amazon S3.
	S3DestinationConfiguration interface{} `json:"s3DestinationConfiguration"`
}

// When dataset contents are created, they are delivered to destination specified here.
//
// TODO: EXAMPLE
//
type CfnDataset_DatasetContentDeliveryRuleProperty struct {
	// The destination to which dataset contents are delivered.
	Destination interface{} `json:"destination"`
	// The name of the dataset content delivery rules entry.
	EntryName *string `json:"entryName"`
}

// The dataset whose latest contents are used as input to the notebook or application.
//
// TODO: EXAMPLE
//
type CfnDataset_DatasetContentVersionValueProperty struct {
	// The name of the dataset whose latest contents are used as input to the notebook or application.
	DatasetName *string `json:"datasetName"`
}

// Used to limit data to that which has arrived since the last execution of the action.
//
// TODO: EXAMPLE
//
type CfnDataset_DeltaTimeProperty struct {
	// The number of seconds of estimated in-flight lag time of message data.
	//
	// When you create dataset contents using message data from a specified timeframe, some message data might still be in flight when processing begins, and so do not arrive in time to be processed. Use this field to make allowances for the in flight time of your message data, so that data not processed from a previous timeframe is included with the next timeframe. Otherwise, missed message data would be excluded from processing during the next timeframe too, because its timestamp places it within the previous timeframe.
	OffsetSeconds *float64 `json:"offsetSeconds"`
	// An expression by which the time of the message data might be determined.
	//
	// This can be the name of a timestamp field or a SQL expression that is used to derive the time the message data was generated.
	TimeExpression *string `json:"timeExpression"`
}

// A structure that contains the configuration information of a delta time session window.
//
// [`DeltaTime`](https://docs.aws.amazon.com/iotanalytics/latest/APIReference/API_DeltaTime.html) specifies a time interval. You can use `DeltaTime` to create dataset contents with data that has arrived in the data store since the last execution. For an example of `DeltaTime` , see [Creating a SQL dataset with a delta window (CLI)](https://docs.aws.amazon.com/iotanalytics/latest/userguide/automate-create-dataset.html#automate-example6) in the *AWS IoT Analytics User Guide* .
//
// TODO: EXAMPLE
//
type CfnDataset_DeltaTimeSessionWindowConfigurationProperty struct {
	// A time interval.
	//
	// You can use `timeoutInMinutes` so that AWS IoT Analytics can batch up late data notifications that have been generated since the last execution. AWS IoT Analytics sends one batch of notifications to Amazon CloudWatch Events at one time.
	//
	// For more information about how to write a timestamp expression, see [Date and Time Functions and Operators](https://docs.aws.amazon.com/https://prestodb.io/docs/0.172/functions/datetime.html) , in the *Presto 0.172 Documentation* .
	TimeoutInMinutes *float64 `json:"timeoutInMinutes"`
}

// Information which is used to filter message data, to segregate it according to the time frame in which it arrives.
//
// TODO: EXAMPLE
//
type CfnDataset_FilterProperty struct {
	// Used to limit data to that which has arrived since the last execution of the action.
	DeltaTime interface{} `json:"deltaTime"`
}

// Configuration information for coordination with AWS Glue , a fully managed extract, transform and load (ETL) service.
//
// TODO: EXAMPLE
//
type CfnDataset_GlueConfigurationProperty struct {
	// The name of the database in your AWS Glue Data Catalog in which the table is located.
	//
	// An AWS Glue Data Catalog database contains metadata tables.
	DatabaseName *string `json:"databaseName"`
	// The name of the table in your AWS Glue Data Catalog that is used to perform the ETL operations.
	//
	// An AWS Glue Data Catalog table contains partitioned data and descriptions of data sources and targets.
	TableName *string `json:"tableName"`
}

// Configuration information for delivery of dataset contents to AWS IoT Events .
//
// TODO: EXAMPLE
//
type CfnDataset_IotEventsDestinationConfigurationProperty struct {
	// The name of the AWS IoT Events input to which dataset contents are delivered.
	InputName *string `json:"inputName"`
	// The ARN of the role that grants AWS IoT Analytics permission to deliver dataset contents to an AWS IoT Events input.
	RoleArn *string `json:"roleArn"`
}

// The information needed to configure a delta time session window.
//
// TODO: EXAMPLE
//
type CfnDataset_LateDataRuleConfigurationProperty struct {
	// The information needed to configure a delta time session window.
	DeltaTimeSessionWindowConfiguration interface{} `json:"deltaTimeSessionWindowConfiguration"`
}

// A structure that contains the name and configuration information of a late data rule.
//
// TODO: EXAMPLE
//
type CfnDataset_LateDataRuleProperty struct {
	// The information needed to configure the late data rule.
	RuleConfiguration interface{} `json:"ruleConfiguration"`
	// The name of the late data rule.
	RuleName *string `json:"ruleName"`
}

// The value of the variable as a structure that specifies an output file URI.
//
// TODO: EXAMPLE
//
type CfnDataset_OutputFileUriValueProperty struct {
	// The URI of the location where dataset contents are stored, usually the URI of a file in an S3 bucket.
	FileName *string `json:"fileName"`
}

// An "SqlQueryDatasetAction" object that uses an SQL query to automatically create data set contents.
//
// TODO: EXAMPLE
//
type CfnDataset_QueryActionProperty struct {
	// An "SqlQueryDatasetAction" object that uses an SQL query to automatically create data set contents.
	SqlQuery *string `json:"sqlQuery"`
	// Pre-filters applied to message data.
	Filters interface{} `json:"filters"`
}

// The configuration of the resource used to execute the `containerAction` .
//
// TODO: EXAMPLE
//
type CfnDataset_ResourceConfigurationProperty struct {
	// The type of the compute resource used to execute the `containerAction` .
	//
	// Possible values are: `ACU_1` (vCPU=4, memory=16 GiB) or `ACU_2` (vCPU=8, memory=32 GiB).
	ComputeType *string `json:"computeType"`
	// The size, in GB, of the persistent storage available to the resource instance used to execute the `containerAction` (min: 1, max: 50).
	VolumeSizeInGb *float64 `json:"volumeSizeInGb"`
}

// How long, in days, message data is kept.
//
// TODO: EXAMPLE
//
type CfnDataset_RetentionPeriodProperty struct {
	// The number of days that message data is kept.
	//
	// The `unlimited` parameter must be false.
	NumberOfDays *float64 `json:"numberOfDays"`
	// If true, message data is kept indefinitely.
	Unlimited interface{} `json:"unlimited"`
}

// Configuration information for delivery of dataset contents to Amazon Simple Storage Service (Amazon S3).
//
// TODO: EXAMPLE
//
type CfnDataset_S3DestinationConfigurationProperty struct {
	// The name of the S3 bucket to which dataset contents are delivered.
	Bucket *string `json:"bucket"`
	// The key of the dataset contents object in an S3 bucket.
	//
	// Each object has a key that is a unique identifier. Each object has exactly one key.
	//
	// You can create a unique key with the following options:
	//
	// - Use `!{iotanalytics:scheduleTime}` to insert the time of a scheduled SQL query run.
	// - Use `!{iotanalytics:versionId}` to insert a unique hash that identifies a dataset content.
	// - Use `!{iotanalytics:creationTime}` to insert the creation time of a dataset content.
	//
	// The following example creates a unique key for a CSV file: `dataset/mydataset/!{iotanalytics:scheduleTime}/!{iotanalytics:versionId}.csv`
	//
	// > If you don't use `!{iotanalytics:versionId}` to specify the key, you might get duplicate keys. For example, you might have two dataset contents with the same `scheduleTime` but different `versionId` s. This means that one dataset content overwrites the other.
	Key *string `json:"key"`
	// The ARN of the role that grants AWS IoT Analytics permission to interact with your Amazon S3 and AWS Glue resources.
	RoleArn *string `json:"roleArn"`
	// Configuration information for coordination with AWS Glue , a fully managed extract, transform and load (ETL) service.
	GlueConfiguration interface{} `json:"glueConfiguration"`
}

// The schedule for when to trigger an update.
//
// TODO: EXAMPLE
//
type CfnDataset_ScheduleProperty struct {
	// The expression that defines when to trigger an update.
	//
	// For more information, see [Schedule Expressions for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html) in the Amazon CloudWatch documentation.
	ScheduleExpression *string `json:"scheduleExpression"`
}

// The "DatasetTrigger" that specifies when the data set is automatically updated.
//
// TODO: EXAMPLE
//
type CfnDataset_TriggerProperty struct {
	// The "Schedule" when the trigger is initiated.
	Schedule interface{} `json:"schedule"`
	// Information about the data set whose content generation triggers the new data set content generation.
	TriggeringDataset interface{} `json:"triggeringDataset"`
}

// Information about the dataset whose content generation triggers the new dataset content generation.
//
// TODO: EXAMPLE
//
type CfnDataset_TriggeringDatasetProperty struct {
	// The name of the data set whose content generation triggers the new data set content generation.
	DatasetName *string `json:"datasetName"`
}

// An instance of a variable to be passed to the `containerAction` execution.
//
// Each variable must have a name and a value given by one of `stringValue` , `datasetContentVersionValue` , or `outputFileUriValue` .
//
// TODO: EXAMPLE
//
type CfnDataset_VariableProperty struct {
	// The name of the variable.
	VariableName *string `json:"variableName"`
	// The value of the variable as a structure that specifies a dataset content version.
	DatasetContentVersionValue interface{} `json:"datasetContentVersionValue"`
	// The value of the variable as a double (numeric).
	DoubleValue *float64 `json:"doubleValue"`
	// The value of the variable as a structure that specifies an output file URI.
	OutputFileUriValue interface{} `json:"outputFileUriValue"`
	// The value of the variable as a string.
	StringValue *string `json:"stringValue"`
}

// Information about the versioning of dataset contents.
//
// TODO: EXAMPLE
//
type CfnDataset_VersioningConfigurationProperty struct {
	// How many versions of dataset contents are kept.
	//
	// The `unlimited` parameter must be `false` .
	MaxVersions *float64 `json:"maxVersions"`
	// If true, unlimited versions of dataset contents are kept.
	Unlimited interface{} `json:"unlimited"`
}

// Properties for defining a `CfnDataset`.
//
// TODO: EXAMPLE
//
type CfnDatasetProps struct {
	// The `DatasetAction` objects that automatically create the dataset contents.
	Actions interface{} `json:"actions"`
	// When dataset contents are created they are delivered to destinations specified here.
	ContentDeliveryRules interface{} `json:"contentDeliveryRules"`
	// The name of the dataset.
	DatasetName *string `json:"datasetName"`
	// A list of data rules that send notifications to CloudWatch, when data arrives late.
	//
	// To specify `lateDataRules` , the dataset must use a [DeltaTimer](https://docs.aws.amazon.com/iotanalytics/latest/APIReference/API_DeltaTime.html) filter.
	LateDataRules interface{} `json:"lateDataRules"`
	// Optional.
	//
	// How long, in days, message data is kept for the dataset.
	RetentionPeriod interface{} `json:"retentionPeriod"`
	// Metadata which can be used to manage the data set.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// The `DatasetTrigger` objects that specify when the dataset is automatically updated.
	Triggers interface{} `json:"triggers"`
	// Optional.
	//
	// How many versions of dataset contents are kept. If not specified or set to null, only the latest version plus the latest succeeded version (if they are different) are kept for the time period specified by the `retentionPeriod` parameter. For more information, see [Keeping Multiple Versions of AWS IoT Analytics datasets](https://docs.aws.amazon.com/iotanalytics/latest/userguide/getting-started.html#aws-iot-analytics-dataset-versions) in the *AWS IoT Analytics User Guide* .
	VersioningConfiguration interface{} `json:"versioningConfiguration"`
}

// A CloudFormation `AWS::IoTAnalytics::Datastore`.
//
// AWS::IoTAnalytics::Datastore resource is a repository for messages. For more information, see [How to Use AWS IoT Analytics](https://docs.aws.amazon.com/iotanalytics/latest/userguide/welcome.html#aws-iot-analytics-how) in the *AWS IoT Analytics User Guide* .
//
// TODO: EXAMPLE
//
type CfnDatastore interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DatastoreName() *string
	SetDatastoreName(val *string)
	DatastorePartitions() interface{}
	SetDatastorePartitions(val interface{})
	DatastoreStorage() interface{}
	SetDatastoreStorage(val interface{})
	FileFormatConfiguration() interface{}
	SetFileFormatConfiguration(val interface{})
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	RetentionPeriod() interface{}
	SetRetentionPeriod(val interface{})
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

// The jsii proxy struct for CfnDatastore
type jsiiProxy_CfnDatastore struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDatastore) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) DatastoreName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) DatastorePartitions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datastorePartitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) DatastoreStorage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datastoreStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) FileFormatConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fileFormatConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) RetentionPeriod() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatastore) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTAnalytics::Datastore`.
func NewCfnDatastore(scope constructs.Construct, id *string, props *CfnDatastoreProps) CfnDatastore {
	_init_.Initialize()

	j := jsiiProxy_CfnDatastore{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iotanalytics.CfnDatastore",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTAnalytics::Datastore`.
func NewCfnDatastore_Override(c CfnDatastore, scope constructs.Construct, id *string, props *CfnDatastoreProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iotanalytics.CfnDatastore",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDatastore) SetDatastoreName(val *string) {
	_jsii_.Set(
		j,
		"datastoreName",
		val,
	)
}

func (j *jsiiProxy_CfnDatastore) SetDatastorePartitions(val interface{}) {
	_jsii_.Set(
		j,
		"datastorePartitions",
		val,
	)
}

func (j *jsiiProxy_CfnDatastore) SetDatastoreStorage(val interface{}) {
	_jsii_.Set(
		j,
		"datastoreStorage",
		val,
	)
}

func (j *jsiiProxy_CfnDatastore) SetFileFormatConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"fileFormatConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDatastore) SetRetentionPeriod(val interface{}) {
	_jsii_.Set(
		j,
		"retentionPeriod",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDatastore_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotanalytics.CfnDatastore",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDatastore_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotanalytics.CfnDatastore",
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
func CfnDatastore_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotanalytics.CfnDatastore",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDatastore_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iotanalytics.CfnDatastore",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnDatastore) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnDatastore) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnDatastore) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnDatastore) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnDatastore) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnDatastore) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnDatastore) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnDatastore) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnDatastore) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnDatastore) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnDatastore) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDatastore) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnDatastore) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnDatastore) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDatastore) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Contains information about a column that stores your data.
//
// TODO: EXAMPLE
//
type CfnDatastore_ColumnProperty struct {
	// The name of the column.
	Name *string `json:"name"`
	// The type of data.
	//
	// For more information about the supported data types, see [Common data types](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-common.html) in the *AWS Glue Developer Guide* .
	Type *string `json:"type"`
}

// S3-customer-managed;
//
// When you choose customer-managed storage, the `retentionPeriod` parameter is ignored. You can't change the choice of Amazon S3 storage after your data store is created.
//
// TODO: EXAMPLE
//
type CfnDatastore_CustomerManagedS3Property struct {
	// The name of the Amazon S3 bucket where your data is stored.
	Bucket *string `json:"bucket"`
	// The ARN of the role that grants AWS IoT Analytics permission to interact with your Amazon S3 resources.
	RoleArn *string `json:"roleArn"`
	// (Optional) The prefix used to create the keys of the data store data objects.
	//
	// Each object in an Amazon S3 bucket has a key that is its unique identifier in the bucket. Each object in a bucket has exactly one key. The prefix must end with a forward slash (/).
	KeyPrefix *string `json:"keyPrefix"`
}

// Amazon S3 -customer-managed;
//
// When you choose customer-managed storage, the `retentionPeriod` parameter is ignored. You can't change the choice of Amazon S3 storage after your data store is created.
//
// TODO: EXAMPLE
//
type CfnDatastore_CustomerManagedS3StorageProperty struct {
	// The name of the Amazon S3 bucket where your data is stored.
	Bucket *string `json:"bucket"`
	// (Optional) The prefix used to create the keys of the data store data objects.
	//
	// Each object in an Amazon S3 bucket has a key that is its unique identifier in the bucket. Each object in a bucket has exactly one key. The prefix must end with a forward slash (/).
	KeyPrefix *string `json:"keyPrefix"`
}

// A single dimension to partition a data store.
//
// The dimension must be an `AttributePartition` or a `TimestampPartition` .
//
// TODO: EXAMPLE
//
type CfnDatastore_DatastorePartitionProperty struct {
	// A partition dimension defined by an attribute.
	Partition interface{} `json:"partition"`
	// A partition dimension defined by a timestamp attribute.
	TimestampPartition interface{} `json:"timestampPartition"`
}

// Information about the partition dimensions in a data store.
//
// TODO: EXAMPLE
//
type CfnDatastore_DatastorePartitionsProperty struct {
	// A list of partition dimensions in a data store.
	Partitions interface{} `json:"partitions"`
}

// Where data store data is stored.
//
// TODO: EXAMPLE
//
type CfnDatastore_DatastoreStorageProperty struct {
	// Use this to store data store data in an S3 bucket that you manage.
	//
	// The choice of service-managed or customer-managed S3 storage cannot be changed after creation of the data store.
	CustomerManagedS3 interface{} `json:"customerManagedS3"`
	// Use this to store data used by AWS IoT SiteWise in an Amazon S3 bucket that you manage.
	//
	// You can't change the choice of Amazon S3 storage after your data store is created.
	IotSiteWiseMultiLayerStorage interface{} `json:"iotSiteWiseMultiLayerStorage"`
	// Use this to store data store data in an S3 bucket managed by the AWS IoT Analytics service.
	//
	// The choice of service-managed or customer-managed S3 storage cannot be changed after creation of the data store.
	ServiceManagedS3 interface{} `json:"serviceManagedS3"`
}

// Contains the configuration information of file formats. AWS IoT Analytics data stores support JSON and [Parquet](https://docs.aws.amazon.com/https://parquet.apache.org/) .
//
// The default file format is JSON. You can specify only one format.
//
// You can't change the file format after you create the data store.
//
// TODO: EXAMPLE
//
type CfnDatastore_FileFormatConfigurationProperty struct {
	// Contains the configuration information of the JSON format.
	JsonConfiguration interface{} `json:"jsonConfiguration"`
	// Contains the configuration information of the Parquet format.
	ParquetConfiguration interface{} `json:"parquetConfiguration"`
}

// Stores data used by AWS IoT SiteWise in an Amazon S3 bucket that you manage.
//
// You can't change the choice of Amazon S3 storage after your data store is created.
//
// TODO: EXAMPLE
//
type CfnDatastore_IotSiteWiseMultiLayerStorageProperty struct {
	// Stores data used by AWS IoT SiteWise in an Amazon S3 bucket that you manage.
	CustomerManagedS3Storage interface{} `json:"customerManagedS3Storage"`
}

// Contains the configuration information of the JSON format.
//
// TODO: EXAMPLE
//
type CfnDatastore_JsonConfigurationProperty struct {
}

// Contains the configuration information of the Parquet format.
//
// TODO: EXAMPLE
//
type CfnDatastore_ParquetConfigurationProperty struct {
	// Information needed to define a schema.
	SchemaDefinition interface{} `json:"schemaDefinition"`
}

// A single dimension to partition a data store.
//
// The dimension must be an `AttributePartition` or a `TimestampPartition` .
//
// TODO: EXAMPLE
//
type CfnDatastore_PartitionProperty struct {
	// The name of the attribute that defines a partition dimension.
	AttributeName *string `json:"attributeName"`
}

// How long, in days, message data is kept.
//
// TODO: EXAMPLE
//
type CfnDatastore_RetentionPeriodProperty struct {
	// The number of days that message data is kept.
	//
	// The `unlimited` parameter must be false.
	NumberOfDays *float64 `json:"numberOfDays"`
	// If true, message data is kept indefinitely.
	Unlimited interface{} `json:"unlimited"`
}

// Information needed to define a schema.
//
// TODO: EXAMPLE
//
type CfnDatastore_SchemaDefinitionProperty struct {
	// Specifies one or more columns that store your data.
	//
	// Each schema can have up to 100 columns. Each column can have up to 100 nested types.
	Columns interface{} `json:"columns"`
}

// Used to store data in an Amazon S3 bucket managed by AWS IoT Analytics .
//
// You can't change the choice of Amazon S3 storage after your data store is created.
//
// TODO: EXAMPLE
//
type CfnDatastore_ServiceManagedS3Property struct {
}

// A partition dimension defined by a timestamp attribute.
//
// TODO: EXAMPLE
//
type CfnDatastore_TimestampPartitionProperty struct {
	// The attribute name of the partition defined by a timestamp.
	AttributeName *string `json:"attributeName"`
	// The timestamp format of a partition defined by a timestamp.
	//
	// The default format is seconds since epoch (January 1, 1970 at midnight UTC time).
	TimestampFormat *string `json:"timestampFormat"`
}

// Properties for defining a `CfnDatastore`.
//
// TODO: EXAMPLE
//
type CfnDatastoreProps struct {
	// The name of the data store.
	DatastoreName *string `json:"datastoreName"`
	// Information about the partition dimensions in a data store.
	DatastorePartitions interface{} `json:"datastorePartitions"`
	// Where data store data is stored.
	DatastoreStorage interface{} `json:"datastoreStorage"`
	// Contains the configuration information of file formats. AWS IoT Analytics data stores support JSON and [Parquet](https://docs.aws.amazon.com/https://parquet.apache.org/) .
	//
	// The default file format is JSON. You can specify only one format.
	//
	// You can't change the file format after you create the data store.
	FileFormatConfiguration interface{} `json:"fileFormatConfiguration"`
	// How long, in days, message data is kept for the data store.
	//
	// When `customerManagedS3` storage is selected, this parameter is ignored.
	RetentionPeriod interface{} `json:"retentionPeriod"`
	// Metadata which can be used to manage the data store.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::IoTAnalytics::Pipeline`.
//
// The AWS::IoTAnalytics::Pipeline resource consumes messages from one or more channels and allows you to process the messages before storing them in a data store. You must specify both a `channel` and a `datastore` activity and, optionally, as many as 23 additional activities in the `pipelineActivities` array. For more information, see [How to Use AWS IoT Analytics](https://docs.aws.amazon.com/iotanalytics/latest/userguide/welcome.html#aws-iot-analytics-how) in the *AWS IoT Analytics User Guide* .
//
// TODO: EXAMPLE
//
type CfnPipeline interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	PipelineActivities() interface{}
	SetPipelineActivities(val interface{})
	PipelineName() *string
	SetPipelineName(val *string)
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

// The jsii proxy struct for CfnPipeline
type jsiiProxy_CfnPipeline struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPipeline) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) PipelineActivities() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pipelineActivities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) PipelineName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTAnalytics::Pipeline`.
func NewCfnPipeline(scope constructs.Construct, id *string, props *CfnPipelineProps) CfnPipeline {
	_init_.Initialize()

	j := jsiiProxy_CfnPipeline{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iotanalytics.CfnPipeline",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTAnalytics::Pipeline`.
func NewCfnPipeline_Override(c CfnPipeline, scope constructs.Construct, id *string, props *CfnPipelineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iotanalytics.CfnPipeline",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPipeline) SetPipelineActivities(val interface{}) {
	_jsii_.Set(
		j,
		"pipelineActivities",
		val,
	)
}

func (j *jsiiProxy_CfnPipeline) SetPipelineName(val *string) {
	_jsii_.Set(
		j,
		"pipelineName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPipeline_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotanalytics.CfnPipeline",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPipeline_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotanalytics.CfnPipeline",
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
func CfnPipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotanalytics.CfnPipeline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPipeline_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iotanalytics.CfnPipeline",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnPipeline) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnPipeline) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnPipeline) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnPipeline) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnPipeline) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnPipeline) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnPipeline) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnPipeline) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnPipeline) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnPipeline) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnPipeline) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPipeline) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnPipeline) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnPipeline) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipeline) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// An activity that performs a transformation on a message.
//
// TODO: EXAMPLE
//
type CfnPipeline_ActivityProperty struct {
	// Adds other attributes based on existing attributes in the message.
	AddAttributes interface{} `json:"addAttributes"`
	// Determines the source of the messages to be processed.
	Channel interface{} `json:"channel"`
	// Specifies where to store the processed message data.
	Datastore interface{} `json:"datastore"`
	// Adds data from the AWS IoT device registry to your message.
	DeviceRegistryEnrich interface{} `json:"deviceRegistryEnrich"`
	// Adds information from the AWS IoT Device Shadows service to a message.
	DeviceShadowEnrich interface{} `json:"deviceShadowEnrich"`
	// Filters a message based on its attributes.
	Filter interface{} `json:"filter"`
	// Runs a Lambda function to modify the message.
	Lambda interface{} `json:"lambda"`
	// Computes an arithmetic expression using the message's attributes and adds it to the message.
	Math interface{} `json:"math"`
	// Removes attributes from a message.
	RemoveAttributes interface{} `json:"removeAttributes"`
	// Creates a new message using only the specified attributes from the original message.
	SelectAttributes interface{} `json:"selectAttributes"`
}

// An activity that adds other attributes based on existing attributes in the message.
//
// TODO: EXAMPLE
//
type CfnPipeline_AddAttributesProperty struct {
	// A list of 1-50 "AttributeNameMapping" objects that map an existing attribute to a new attribute.
	//
	// > The existing attributes remain in the message, so if you want to remove the originals, use "RemoveAttributeActivity".
	Attributes interface{} `json:"attributes"`
	// The name of the 'addAttributes' activity.
	Name *string `json:"name"`
	// The next activity in the pipeline.
	Next *string `json:"next"`
}

// Determines the source of the messages to be processed.
//
// TODO: EXAMPLE
//
type CfnPipeline_ChannelProperty struct {
	// The name of the channel from which the messages are processed.
	ChannelName *string `json:"channelName"`
	// The name of the 'channel' activity.
	Name *string `json:"name"`
	// The next activity in the pipeline.
	Next *string `json:"next"`
}

// The datastore activity that specifies where to store the processed data.
//
// TODO: EXAMPLE
//
type CfnPipeline_DatastoreProperty struct {
	// The name of the data store where processed messages are stored.
	DatastoreName *string `json:"datastoreName"`
	// The name of the datastore activity.
	Name *string `json:"name"`
}

// An activity that adds data from the AWS IoT device registry to your message.
//
// TODO: EXAMPLE
//
type CfnPipeline_DeviceRegistryEnrichProperty struct {
	// The name of the attribute that is added to the message.
	Attribute *string `json:"attribute"`
	// The name of the 'deviceRegistryEnrich' activity.
	Name *string `json:"name"`
	// The next activity in the pipeline.
	Next *string `json:"next"`
	// The ARN of the role that allows access to the device's registry information.
	RoleArn *string `json:"roleArn"`
	// The name of the IoT device whose registry information is added to the message.
	ThingName *string `json:"thingName"`
}

// An activity that adds information from the AWS IoT Device Shadows service to a message.
//
// TODO: EXAMPLE
//
type CfnPipeline_DeviceShadowEnrichProperty struct {
	// The name of the attribute that is added to the message.
	Attribute *string `json:"attribute"`
	// The name of the 'deviceShadowEnrich' activity.
	Name *string `json:"name"`
	// The next activity in the pipeline.
	Next *string `json:"next"`
	// The ARN of the role that allows access to the device's shadow.
	RoleArn *string `json:"roleArn"`
	// The name of the IoT device whose shadow information is added to the message.
	ThingName *string `json:"thingName"`
}

// An activity that filters a message based on its attributes.
//
// TODO: EXAMPLE
//
type CfnPipeline_FilterProperty struct {
	// An expression that looks like an SQL WHERE clause that must return a Boolean value.
	Filter *string `json:"filter"`
	// The name of the 'filter' activity.
	Name *string `json:"name"`
	// The next activity in the pipeline.
	Next *string `json:"next"`
}

// An activity that runs a Lambda function to modify the message.
//
// TODO: EXAMPLE
//
type CfnPipeline_LambdaProperty struct {
	// The number of messages passed to the Lambda function for processing.
	//
	// The AWS Lambda function must be able to process all of these messages within five minutes, which is the maximum timeout duration for Lambda functions.
	BatchSize *float64 `json:"batchSize"`
	// The name of the Lambda function that is run on the message.
	LambdaName *string `json:"lambdaName"`
	// The name of the 'lambda' activity.
	Name *string `json:"name"`
	// The next activity in the pipeline.
	Next *string `json:"next"`
}

// An activity that computes an arithmetic expression using the message's attributes.
//
// TODO: EXAMPLE
//
type CfnPipeline_MathProperty struct {
	// The name of the attribute that contains the result of the math operation.
	Attribute *string `json:"attribute"`
	// An expression that uses one or more existing attributes and must return an integer value.
	Math *string `json:"math"`
	// The name of the 'math' activity.
	Name *string `json:"name"`
	// The next activity in the pipeline.
	Next *string `json:"next"`
}

// An activity that removes attributes from a message.
//
// TODO: EXAMPLE
//
type CfnPipeline_RemoveAttributesProperty struct {
	// A list of 1-50 attributes to remove from the message.
	Attributes *[]*string `json:"attributes"`
	// The name of the 'removeAttributes' activity.
	Name *string `json:"name"`
	// The next activity in the pipeline.
	Next *string `json:"next"`
}

// Creates a new message using only the specified attributes from the original message.
//
// TODO: EXAMPLE
//
type CfnPipeline_SelectAttributesProperty struct {
	// A list of the attributes to select from the message.
	Attributes *[]*string `json:"attributes"`
	// The name of the 'selectAttributes' activity.
	Name *string `json:"name"`
	// The next activity in the pipeline.
	Next *string `json:"next"`
}

// Properties for defining a `CfnPipeline`.
//
// TODO: EXAMPLE
//
type CfnPipelineProps struct {
	// A list of "PipelineActivity" objects.
	//
	// Activities perform transformations on your messages, such as removing, renaming or adding message attributes; filtering messages based on attribute values; invoking your Lambda functions on messages for advanced processing; or performing mathematical transformations to normalize device data.
	//
	// The list can be 2-25 *PipelineActivity* objects and must contain both a `channel` and a `datastore` activity. Each entry in the list must contain only one activity, for example:
	//
	// `pipelineActivities = [ { "channel": { ... } }, { "lambda": { ... } }, ... ]`
	PipelineActivities interface{} `json:"pipelineActivities"`
	// The name of the pipeline.
	PipelineName *string `json:"pipelineName"`
	// Metadata which can be used to manage the pipeline.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `json:"tags"`
}


package awsiotanalytics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiotanalytics/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::IoTAnalytics::Channel`.
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
	Node() awscdk.ConstructNode
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

func (j *jsiiProxy_CfnChannel) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnChannel(scope awscdk.Construct, id *string, props *CfnChannelProps) CfnChannel {
	_init_.Initialize()

	j := jsiiProxy_CfnChannel{}

	_jsii_.Create(
		"monocdk.aws_iotanalytics.CfnChannel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTAnalytics::Channel`.
func NewCfnChannel_Override(c CfnChannel, scope awscdk.Construct, id *string, props *CfnChannelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iotanalytics.CfnChannel",
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
// Experimental.
func CfnChannel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotanalytics.CfnChannel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnChannel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotanalytics.CfnChannel",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotanalytics.CfnChannel",
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
		"monocdk.aws_iotanalytics.CfnChannel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnChannel) OnPrepare() {
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
func (c *jsiiProxy_CfnChannel) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnChannel) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnChannel) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnChannel) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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
// Experimental.
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnChannel) Synthesize(session awscdk.ISynthesisSession) {
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnChannel) Validate() *[]*string {
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
func (c *jsiiProxy_CfnChannel) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnChannel_ChannelStorageProperty struct {
	// `CfnChannel.ChannelStorageProperty.CustomerManagedS3`.
	CustomerManagedS3 interface{} `json:"customerManagedS3"`
	// `CfnChannel.ChannelStorageProperty.ServiceManagedS3`.
	ServiceManagedS3 interface{} `json:"serviceManagedS3"`
}

type CfnChannel_CustomerManagedS3Property struct {
	// `CfnChannel.CustomerManagedS3Property.Bucket`.
	Bucket *string `json:"bucket"`
	// `CfnChannel.CustomerManagedS3Property.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `CfnChannel.CustomerManagedS3Property.KeyPrefix`.
	KeyPrefix *string `json:"keyPrefix"`
}

type CfnChannel_RetentionPeriodProperty struct {
	// `CfnChannel.RetentionPeriodProperty.NumberOfDays`.
	NumberOfDays *float64 `json:"numberOfDays"`
	// `CfnChannel.RetentionPeriodProperty.Unlimited`.
	Unlimited interface{} `json:"unlimited"`
}

type CfnChannel_ServiceManagedS3Property struct {
}

// Properties for defining a `AWS::IoTAnalytics::Channel`.
type CfnChannelProps struct {
	// `AWS::IoTAnalytics::Channel.ChannelName`.
	ChannelName *string `json:"channelName"`
	// `AWS::IoTAnalytics::Channel.ChannelStorage`.
	ChannelStorage interface{} `json:"channelStorage"`
	// `AWS::IoTAnalytics::Channel.RetentionPeriod`.
	RetentionPeriod interface{} `json:"retentionPeriod"`
	// `AWS::IoTAnalytics::Channel.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::IoTAnalytics::Dataset`.
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
	Node() awscdk.ConstructNode
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

func (j *jsiiProxy_CfnDataset) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnDataset(scope awscdk.Construct, id *string, props *CfnDatasetProps) CfnDataset {
	_init_.Initialize()

	j := jsiiProxy_CfnDataset{}

	_jsii_.Create(
		"monocdk.aws_iotanalytics.CfnDataset",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTAnalytics::Dataset`.
func NewCfnDataset_Override(c CfnDataset, scope awscdk.Construct, id *string, props *CfnDatasetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iotanalytics.CfnDataset",
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
// Experimental.
func CfnDataset_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotanalytics.CfnDataset",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDataset_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotanalytics.CfnDataset",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDataset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotanalytics.CfnDataset",
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
		"monocdk.aws_iotanalytics.CfnDataset",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnDataset) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnDataset) OnPrepare() {
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
func (c *jsiiProxy_CfnDataset) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnDataset) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnDataset) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnDataset) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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
// Experimental.
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnDataset) Synthesize(session awscdk.ISynthesisSession) {
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnDataset) Validate() *[]*string {
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
func (c *jsiiProxy_CfnDataset) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnDataset_ActionProperty struct {
	// `CfnDataset.ActionProperty.ActionName`.
	ActionName *string `json:"actionName"`
	// `CfnDataset.ActionProperty.ContainerAction`.
	ContainerAction interface{} `json:"containerAction"`
	// `CfnDataset.ActionProperty.QueryAction`.
	QueryAction interface{} `json:"queryAction"`
}

type CfnDataset_ContainerActionProperty struct {
	// `CfnDataset.ContainerActionProperty.ExecutionRoleArn`.
	ExecutionRoleArn *string `json:"executionRoleArn"`
	// `CfnDataset.ContainerActionProperty.Image`.
	Image *string `json:"image"`
	// `CfnDataset.ContainerActionProperty.ResourceConfiguration`.
	ResourceConfiguration interface{} `json:"resourceConfiguration"`
	// `CfnDataset.ContainerActionProperty.Variables`.
	Variables interface{} `json:"variables"`
}

type CfnDataset_DatasetContentDeliveryRuleDestinationProperty struct {
	// `CfnDataset.DatasetContentDeliveryRuleDestinationProperty.IotEventsDestinationConfiguration`.
	IotEventsDestinationConfiguration interface{} `json:"iotEventsDestinationConfiguration"`
	// `CfnDataset.DatasetContentDeliveryRuleDestinationProperty.S3DestinationConfiguration`.
	S3DestinationConfiguration interface{} `json:"s3DestinationConfiguration"`
}

type CfnDataset_DatasetContentDeliveryRuleProperty struct {
	// `CfnDataset.DatasetContentDeliveryRuleProperty.Destination`.
	Destination interface{} `json:"destination"`
	// `CfnDataset.DatasetContentDeliveryRuleProperty.EntryName`.
	EntryName *string `json:"entryName"`
}

type CfnDataset_DatasetContentVersionValueProperty struct {
	// `CfnDataset.DatasetContentVersionValueProperty.DatasetName`.
	DatasetName *string `json:"datasetName"`
}

type CfnDataset_DeltaTimeProperty struct {
	// `CfnDataset.DeltaTimeProperty.OffsetSeconds`.
	OffsetSeconds *float64 `json:"offsetSeconds"`
	// `CfnDataset.DeltaTimeProperty.TimeExpression`.
	TimeExpression *string `json:"timeExpression"`
}

type CfnDataset_DeltaTimeSessionWindowConfigurationProperty struct {
	// `CfnDataset.DeltaTimeSessionWindowConfigurationProperty.TimeoutInMinutes`.
	TimeoutInMinutes *float64 `json:"timeoutInMinutes"`
}

type CfnDataset_FilterProperty struct {
	// `CfnDataset.FilterProperty.DeltaTime`.
	DeltaTime interface{} `json:"deltaTime"`
}

type CfnDataset_GlueConfigurationProperty struct {
	// `CfnDataset.GlueConfigurationProperty.DatabaseName`.
	DatabaseName *string `json:"databaseName"`
	// `CfnDataset.GlueConfigurationProperty.TableName`.
	TableName *string `json:"tableName"`
}

type CfnDataset_IotEventsDestinationConfigurationProperty struct {
	// `CfnDataset.IotEventsDestinationConfigurationProperty.InputName`.
	InputName *string `json:"inputName"`
	// `CfnDataset.IotEventsDestinationConfigurationProperty.RoleArn`.
	RoleArn *string `json:"roleArn"`
}

type CfnDataset_LateDataRuleConfigurationProperty struct {
	// `CfnDataset.LateDataRuleConfigurationProperty.DeltaTimeSessionWindowConfiguration`.
	DeltaTimeSessionWindowConfiguration interface{} `json:"deltaTimeSessionWindowConfiguration"`
}

type CfnDataset_LateDataRuleProperty struct {
	// `CfnDataset.LateDataRuleProperty.RuleConfiguration`.
	RuleConfiguration interface{} `json:"ruleConfiguration"`
	// `CfnDataset.LateDataRuleProperty.RuleName`.
	RuleName *string `json:"ruleName"`
}

type CfnDataset_OutputFileUriValueProperty struct {
	// `CfnDataset.OutputFileUriValueProperty.FileName`.
	FileName *string `json:"fileName"`
}

type CfnDataset_QueryActionProperty struct {
	// `CfnDataset.QueryActionProperty.SqlQuery`.
	SqlQuery *string `json:"sqlQuery"`
	// `CfnDataset.QueryActionProperty.Filters`.
	Filters interface{} `json:"filters"`
}

type CfnDataset_ResourceConfigurationProperty struct {
	// `CfnDataset.ResourceConfigurationProperty.ComputeType`.
	ComputeType *string `json:"computeType"`
	// `CfnDataset.ResourceConfigurationProperty.VolumeSizeInGB`.
	VolumeSizeInGb *float64 `json:"volumeSizeInGb"`
}

type CfnDataset_RetentionPeriodProperty struct {
	// `CfnDataset.RetentionPeriodProperty.NumberOfDays`.
	NumberOfDays *float64 `json:"numberOfDays"`
	// `CfnDataset.RetentionPeriodProperty.Unlimited`.
	Unlimited interface{} `json:"unlimited"`
}

type CfnDataset_S3DestinationConfigurationProperty struct {
	// `CfnDataset.S3DestinationConfigurationProperty.Bucket`.
	Bucket *string `json:"bucket"`
	// `CfnDataset.S3DestinationConfigurationProperty.Key`.
	Key *string `json:"key"`
	// `CfnDataset.S3DestinationConfigurationProperty.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `CfnDataset.S3DestinationConfigurationProperty.GlueConfiguration`.
	GlueConfiguration interface{} `json:"glueConfiguration"`
}

type CfnDataset_ScheduleProperty struct {
	// `CfnDataset.ScheduleProperty.ScheduleExpression`.
	ScheduleExpression *string `json:"scheduleExpression"`
}

type CfnDataset_TriggerProperty struct {
	// `CfnDataset.TriggerProperty.Schedule`.
	Schedule interface{} `json:"schedule"`
	// `CfnDataset.TriggerProperty.TriggeringDataset`.
	TriggeringDataset interface{} `json:"triggeringDataset"`
}

type CfnDataset_TriggeringDatasetProperty struct {
	// `CfnDataset.TriggeringDatasetProperty.DatasetName`.
	DatasetName *string `json:"datasetName"`
}

type CfnDataset_VariableProperty struct {
	// `CfnDataset.VariableProperty.VariableName`.
	VariableName *string `json:"variableName"`
	// `CfnDataset.VariableProperty.DatasetContentVersionValue`.
	DatasetContentVersionValue interface{} `json:"datasetContentVersionValue"`
	// `CfnDataset.VariableProperty.DoubleValue`.
	DoubleValue *float64 `json:"doubleValue"`
	// `CfnDataset.VariableProperty.OutputFileUriValue`.
	OutputFileUriValue interface{} `json:"outputFileUriValue"`
	// `CfnDataset.VariableProperty.StringValue`.
	StringValue *string `json:"stringValue"`
}

type CfnDataset_VersioningConfigurationProperty struct {
	// `CfnDataset.VersioningConfigurationProperty.MaxVersions`.
	MaxVersions *float64 `json:"maxVersions"`
	// `CfnDataset.VersioningConfigurationProperty.Unlimited`.
	Unlimited interface{} `json:"unlimited"`
}

// Properties for defining a `AWS::IoTAnalytics::Dataset`.
type CfnDatasetProps struct {
	// `AWS::IoTAnalytics::Dataset.Actions`.
	Actions interface{} `json:"actions"`
	// `AWS::IoTAnalytics::Dataset.ContentDeliveryRules`.
	ContentDeliveryRules interface{} `json:"contentDeliveryRules"`
	// `AWS::IoTAnalytics::Dataset.DatasetName`.
	DatasetName *string `json:"datasetName"`
	// `AWS::IoTAnalytics::Dataset.LateDataRules`.
	LateDataRules interface{} `json:"lateDataRules"`
	// `AWS::IoTAnalytics::Dataset.RetentionPeriod`.
	RetentionPeriod interface{} `json:"retentionPeriod"`
	// `AWS::IoTAnalytics::Dataset.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::IoTAnalytics::Dataset.Triggers`.
	Triggers interface{} `json:"triggers"`
	// `AWS::IoTAnalytics::Dataset.VersioningConfiguration`.
	VersioningConfiguration interface{} `json:"versioningConfiguration"`
}

// A CloudFormation `AWS::IoTAnalytics::Datastore`.
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
	Node() awscdk.ConstructNode
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

func (j *jsiiProxy_CfnDatastore) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnDatastore(scope awscdk.Construct, id *string, props *CfnDatastoreProps) CfnDatastore {
	_init_.Initialize()

	j := jsiiProxy_CfnDatastore{}

	_jsii_.Create(
		"monocdk.aws_iotanalytics.CfnDatastore",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTAnalytics::Datastore`.
func NewCfnDatastore_Override(c CfnDatastore, scope awscdk.Construct, id *string, props *CfnDatastoreProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iotanalytics.CfnDatastore",
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
// Experimental.
func CfnDatastore_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotanalytics.CfnDatastore",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDatastore_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotanalytics.CfnDatastore",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDatastore_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotanalytics.CfnDatastore",
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
		"monocdk.aws_iotanalytics.CfnDatastore",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnDatastore) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnDatastore) OnPrepare() {
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
func (c *jsiiProxy_CfnDatastore) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnDatastore) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnDatastore) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnDatastore) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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
// Experimental.
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnDatastore) Synthesize(session awscdk.ISynthesisSession) {
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnDatastore) Validate() *[]*string {
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
func (c *jsiiProxy_CfnDatastore) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnDatastore_ColumnProperty struct {
	// `CfnDatastore.ColumnProperty.Name`.
	Name *string `json:"name"`
	// `CfnDatastore.ColumnProperty.Type`.
	Type *string `json:"type"`
}

type CfnDatastore_CustomerManagedS3Property struct {
	// `CfnDatastore.CustomerManagedS3Property.Bucket`.
	Bucket *string `json:"bucket"`
	// `CfnDatastore.CustomerManagedS3Property.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `CfnDatastore.CustomerManagedS3Property.KeyPrefix`.
	KeyPrefix *string `json:"keyPrefix"`
}

type CfnDatastore_CustomerManagedS3StorageProperty struct {
	// `CfnDatastore.CustomerManagedS3StorageProperty.Bucket`.
	Bucket *string `json:"bucket"`
	// `CfnDatastore.CustomerManagedS3StorageProperty.KeyPrefix`.
	KeyPrefix *string `json:"keyPrefix"`
}

type CfnDatastore_DatastorePartitionProperty struct {
	// `CfnDatastore.DatastorePartitionProperty.Partition`.
	Partition interface{} `json:"partition"`
	// `CfnDatastore.DatastorePartitionProperty.TimestampPartition`.
	TimestampPartition interface{} `json:"timestampPartition"`
}

type CfnDatastore_DatastorePartitionsProperty struct {
	// `CfnDatastore.DatastorePartitionsProperty.Partitions`.
	Partitions interface{} `json:"partitions"`
}

type CfnDatastore_DatastoreStorageProperty struct {
	// `CfnDatastore.DatastoreStorageProperty.CustomerManagedS3`.
	CustomerManagedS3 interface{} `json:"customerManagedS3"`
	// `CfnDatastore.DatastoreStorageProperty.IotSiteWiseMultiLayerStorage`.
	IotSiteWiseMultiLayerStorage interface{} `json:"iotSiteWiseMultiLayerStorage"`
	// `CfnDatastore.DatastoreStorageProperty.ServiceManagedS3`.
	ServiceManagedS3 interface{} `json:"serviceManagedS3"`
}

type CfnDatastore_FileFormatConfigurationProperty struct {
	// `CfnDatastore.FileFormatConfigurationProperty.JsonConfiguration`.
	JsonConfiguration interface{} `json:"jsonConfiguration"`
	// `CfnDatastore.FileFormatConfigurationProperty.ParquetConfiguration`.
	ParquetConfiguration interface{} `json:"parquetConfiguration"`
}

type CfnDatastore_IotSiteWiseMultiLayerStorageProperty struct {
	// `CfnDatastore.IotSiteWiseMultiLayerStorageProperty.CustomerManagedS3Storage`.
	CustomerManagedS3Storage interface{} `json:"customerManagedS3Storage"`
}

type CfnDatastore_JsonConfigurationProperty struct {
}

type CfnDatastore_ParquetConfigurationProperty struct {
	// `CfnDatastore.ParquetConfigurationProperty.SchemaDefinition`.
	SchemaDefinition interface{} `json:"schemaDefinition"`
}

type CfnDatastore_PartitionProperty struct {
	// `CfnDatastore.PartitionProperty.AttributeName`.
	AttributeName *string `json:"attributeName"`
}

type CfnDatastore_RetentionPeriodProperty struct {
	// `CfnDatastore.RetentionPeriodProperty.NumberOfDays`.
	NumberOfDays *float64 `json:"numberOfDays"`
	// `CfnDatastore.RetentionPeriodProperty.Unlimited`.
	Unlimited interface{} `json:"unlimited"`
}

type CfnDatastore_SchemaDefinitionProperty struct {
	// `CfnDatastore.SchemaDefinitionProperty.Columns`.
	Columns interface{} `json:"columns"`
}

type CfnDatastore_ServiceManagedS3Property struct {
}

type CfnDatastore_TimestampPartitionProperty struct {
	// `CfnDatastore.TimestampPartitionProperty.AttributeName`.
	AttributeName *string `json:"attributeName"`
	// `CfnDatastore.TimestampPartitionProperty.TimestampFormat`.
	TimestampFormat *string `json:"timestampFormat"`
}

// Properties for defining a `AWS::IoTAnalytics::Datastore`.
type CfnDatastoreProps struct {
	// `AWS::IoTAnalytics::Datastore.DatastoreName`.
	DatastoreName *string `json:"datastoreName"`
	// `AWS::IoTAnalytics::Datastore.DatastorePartitions`.
	DatastorePartitions interface{} `json:"datastorePartitions"`
	// `AWS::IoTAnalytics::Datastore.DatastoreStorage`.
	DatastoreStorage interface{} `json:"datastoreStorage"`
	// `AWS::IoTAnalytics::Datastore.FileFormatConfiguration`.
	FileFormatConfiguration interface{} `json:"fileFormatConfiguration"`
	// `AWS::IoTAnalytics::Datastore.RetentionPeriod`.
	RetentionPeriod interface{} `json:"retentionPeriod"`
	// `AWS::IoTAnalytics::Datastore.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::IoTAnalytics::Pipeline`.
type CfnPipeline interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() awscdk.ConstructNode
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

// The jsii proxy struct for CfnPipeline
type jsiiProxy_CfnPipeline struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
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

func (j *jsiiProxy_CfnPipeline) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnPipeline(scope awscdk.Construct, id *string, props *CfnPipelineProps) CfnPipeline {
	_init_.Initialize()

	j := jsiiProxy_CfnPipeline{}

	_jsii_.Create(
		"monocdk.aws_iotanalytics.CfnPipeline",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTAnalytics::Pipeline`.
func NewCfnPipeline_Override(c CfnPipeline, scope awscdk.Construct, id *string, props *CfnPipelineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iotanalytics.CfnPipeline",
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
// Experimental.
func CfnPipeline_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotanalytics.CfnPipeline",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnPipeline_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotanalytics.CfnPipeline",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnPipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotanalytics.CfnPipeline",
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
		"monocdk.aws_iotanalytics.CfnPipeline",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnPipeline) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnPipeline) OnPrepare() {
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
func (c *jsiiProxy_CfnPipeline) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnPipeline) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnPipeline) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnPipeline) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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
// Experimental.
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnPipeline) Synthesize(session awscdk.ISynthesisSession) {
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnPipeline) Validate() *[]*string {
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
func (c *jsiiProxy_CfnPipeline) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnPipeline_ActivityProperty struct {
	// `CfnPipeline.ActivityProperty.AddAttributes`.
	AddAttributes interface{} `json:"addAttributes"`
	// `CfnPipeline.ActivityProperty.Channel`.
	Channel interface{} `json:"channel"`
	// `CfnPipeline.ActivityProperty.Datastore`.
	Datastore interface{} `json:"datastore"`
	// `CfnPipeline.ActivityProperty.DeviceRegistryEnrich`.
	DeviceRegistryEnrich interface{} `json:"deviceRegistryEnrich"`
	// `CfnPipeline.ActivityProperty.DeviceShadowEnrich`.
	DeviceShadowEnrich interface{} `json:"deviceShadowEnrich"`
	// `CfnPipeline.ActivityProperty.Filter`.
	Filter interface{} `json:"filter"`
	// `CfnPipeline.ActivityProperty.Lambda`.
	Lambda interface{} `json:"lambda"`
	// `CfnPipeline.ActivityProperty.Math`.
	Math interface{} `json:"math"`
	// `CfnPipeline.ActivityProperty.RemoveAttributes`.
	RemoveAttributes interface{} `json:"removeAttributes"`
	// `CfnPipeline.ActivityProperty.SelectAttributes`.
	SelectAttributes interface{} `json:"selectAttributes"`
}

type CfnPipeline_AddAttributesProperty struct {
	// `CfnPipeline.AddAttributesProperty.Attributes`.
	Attributes interface{} `json:"attributes"`
	// `CfnPipeline.AddAttributesProperty.Name`.
	Name *string `json:"name"`
	// `CfnPipeline.AddAttributesProperty.Next`.
	Next *string `json:"next"`
}

type CfnPipeline_ChannelProperty struct {
	// `CfnPipeline.ChannelProperty.ChannelName`.
	ChannelName *string `json:"channelName"`
	// `CfnPipeline.ChannelProperty.Name`.
	Name *string `json:"name"`
	// `CfnPipeline.ChannelProperty.Next`.
	Next *string `json:"next"`
}

type CfnPipeline_DatastoreProperty struct {
	// `CfnPipeline.DatastoreProperty.DatastoreName`.
	DatastoreName *string `json:"datastoreName"`
	// `CfnPipeline.DatastoreProperty.Name`.
	Name *string `json:"name"`
}

type CfnPipeline_DeviceRegistryEnrichProperty struct {
	// `CfnPipeline.DeviceRegistryEnrichProperty.Attribute`.
	Attribute *string `json:"attribute"`
	// `CfnPipeline.DeviceRegistryEnrichProperty.Name`.
	Name *string `json:"name"`
	// `CfnPipeline.DeviceRegistryEnrichProperty.Next`.
	Next *string `json:"next"`
	// `CfnPipeline.DeviceRegistryEnrichProperty.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `CfnPipeline.DeviceRegistryEnrichProperty.ThingName`.
	ThingName *string `json:"thingName"`
}

type CfnPipeline_DeviceShadowEnrichProperty struct {
	// `CfnPipeline.DeviceShadowEnrichProperty.Attribute`.
	Attribute *string `json:"attribute"`
	// `CfnPipeline.DeviceShadowEnrichProperty.Name`.
	Name *string `json:"name"`
	// `CfnPipeline.DeviceShadowEnrichProperty.Next`.
	Next *string `json:"next"`
	// `CfnPipeline.DeviceShadowEnrichProperty.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `CfnPipeline.DeviceShadowEnrichProperty.ThingName`.
	ThingName *string `json:"thingName"`
}

type CfnPipeline_FilterProperty struct {
	// `CfnPipeline.FilterProperty.Filter`.
	Filter *string `json:"filter"`
	// `CfnPipeline.FilterProperty.Name`.
	Name *string `json:"name"`
	// `CfnPipeline.FilterProperty.Next`.
	Next *string `json:"next"`
}

type CfnPipeline_LambdaProperty struct {
	// `CfnPipeline.LambdaProperty.BatchSize`.
	BatchSize *float64 `json:"batchSize"`
	// `CfnPipeline.LambdaProperty.LambdaName`.
	LambdaName *string `json:"lambdaName"`
	// `CfnPipeline.LambdaProperty.Name`.
	Name *string `json:"name"`
	// `CfnPipeline.LambdaProperty.Next`.
	Next *string `json:"next"`
}

type CfnPipeline_MathProperty struct {
	// `CfnPipeline.MathProperty.Attribute`.
	Attribute *string `json:"attribute"`
	// `CfnPipeline.MathProperty.Math`.
	Math *string `json:"math"`
	// `CfnPipeline.MathProperty.Name`.
	Name *string `json:"name"`
	// `CfnPipeline.MathProperty.Next`.
	Next *string `json:"next"`
}

type CfnPipeline_RemoveAttributesProperty struct {
	// `CfnPipeline.RemoveAttributesProperty.Attributes`.
	Attributes *[]*string `json:"attributes"`
	// `CfnPipeline.RemoveAttributesProperty.Name`.
	Name *string `json:"name"`
	// `CfnPipeline.RemoveAttributesProperty.Next`.
	Next *string `json:"next"`
}

type CfnPipeline_SelectAttributesProperty struct {
	// `CfnPipeline.SelectAttributesProperty.Attributes`.
	Attributes *[]*string `json:"attributes"`
	// `CfnPipeline.SelectAttributesProperty.Name`.
	Name *string `json:"name"`
	// `CfnPipeline.SelectAttributesProperty.Next`.
	Next *string `json:"next"`
}

// Properties for defining a `AWS::IoTAnalytics::Pipeline`.
type CfnPipelineProps struct {
	// `AWS::IoTAnalytics::Pipeline.PipelineActivities`.
	PipelineActivities interface{} `json:"pipelineActivities"`
	// `AWS::IoTAnalytics::Pipeline.PipelineName`.
	PipelineName *string `json:"pipelineName"`
	// `AWS::IoTAnalytics::Pipeline.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}


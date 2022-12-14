package awsmediaconvert

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediaconvert/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::MediaConvert::JobTemplate`.
//
// The AWS::MediaConvert::JobTemplate resource is an AWS Elemental MediaConvert resource type that you can use to generate transcoding jobs.
//
// When you declare this entity in your AWS CloudFormation template, you pass in your transcoding job settings in JSON or YAML format. This settings specification must be formed in a particular way that conforms to AWS Elemental MediaConvert job validation. For more information about creating a job template model for the `SettingsJson` property, see the Remarks section later in this topic.
//
// For information about job templates, see [Working with AWS Elemental MediaConvert Job Templates](https://docs.aws.amazon.com/mediaconvert/latest/ug/working-with-job-templates.html) in the ** .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var settingsJson interface{}
//   var tags interface{}
//
//   cfnJobTemplate := awscdk.Aws_mediaconvert.NewCfnJobTemplate(this, jsii.String("MyCfnJobTemplate"), &cfnJobTemplateProps{
//   	settingsJson: settingsJson,
//
//   	// the properties below are optional
//   	accelerationSettings: &accelerationSettingsProperty{
//   		mode: jsii.String("mode"),
//   	},
//   	category: jsii.String("category"),
//   	description: jsii.String("description"),
//   	hopDestinations: []interface{}{
//   		&hopDestinationProperty{
//   			priority: jsii.Number(123),
//   			queue: jsii.String("queue"),
//   			waitMinutes: jsii.Number(123),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	priority: jsii.Number(123),
//   	queue: jsii.String("queue"),
//   	statusUpdateInterval: jsii.String("statusUpdateInterval"),
//   	tags: tags,
//   })
//
type CfnJobTemplate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Accelerated transcoding can significantly speed up jobs with long, visually complex content.
	//
	// Outputs that use this feature incur pro-tier pricing. For information about feature limitations, For more information, see [Job Limitations for Accelerated Transcoding in AWS Elemental MediaConvert](https://docs.aws.amazon.com/mediaconvert/latest/ug/job-requirements.html) in the *AWS Elemental MediaConvert User Guide* .
	AccelerationSettings() interface{}
	SetAccelerationSettings(val interface{})
	// The Amazon Resource Name (ARN) of the job template, such as `arn:aws:mediaconvert:us-west-2:123456789012` .
	AttrArn() *string
	// The name of the job template, such as `Streaming stack DASH` .
	AttrName() *string
	// Optional.
	//
	// A category for the job template you are creating.
	Category() *string
	SetCategory(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Optional.
	//
	// A description of the job template you are creating.
	Description() *string
	SetDescription(val *string)
	// Optional.
	//
	// Configuration for a destination queue to which the job can hop once a customer-defined minimum wait time has passed. For more information, see [Setting Up Queue Hopping to Avoid Long Waits](https://docs.aws.amazon.com/mediaconvert/latest/ug/setting-up-queue-hopping-to-avoid-long-waits.html) in the *AWS Elemental MediaConvert User Guide* .
	HopDestinations() interface{}
	SetHopDestinations(val interface{})
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
	// The name of the job template you are creating.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Specify the relative priority for this job.
	//
	// In any given queue, the service begins processing the job with the highest value first. When more than one job has the same priority, the service begins processing the job that you submitted first. If you don't specify a priority, the service uses the default value 0. Minimum: -50 Maximum: 50
	Priority() *float64
	SetPriority(val *float64)
	// Optional.
	//
	// The queue that jobs created from this template are assigned to. Specify the Amazon Resource Name (ARN) of the queue. For example, arn:aws:mediaconvert:us-west-2:505474453218:queues/Default. If you don't specify this, jobs will go to the default queue.
	Queue() *string
	SetQueue(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Specify, in JSON format, the transcoding job settings for this job template.
	//
	// This specification must conform to the AWS Elemental MediaConvert job validation. For information about forming this specification, see the Remarks section later in this topic.
	//
	// For more information about MediaConvert job templates, see [Working with AWS Elemental MediaConvert Job Templates](https://docs.aws.amazon.com/mediaconvert/latest/ug/working-with-job-templates.html) in the ** .
	SettingsJson() interface{}
	SetSettingsJson(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Specify how often MediaConvert sends STATUS_UPDATE events to Amazon CloudWatch Events.
	//
	// Set the interval, in seconds, between status updates. MediaConvert sends an update at this interval from the time the service begins processing your job to the time it completes the transcode or encounters an error.
	//
	// Specify one of the following enums:
	//
	// SECONDS_10
	//
	// SECONDS_12
	//
	// SECONDS_15
	//
	// SECONDS_20
	//
	// SECONDS_30
	//
	// SECONDS_60
	//
	// SECONDS_120
	//
	// SECONDS_180
	//
	// SECONDS_240
	//
	// SECONDS_300
	//
	// SECONDS_360
	//
	// SECONDS_420
	//
	// SECONDS_480
	//
	// SECONDS_540
	//
	// SECONDS_600.
	StatusUpdateInterval() *string
	SetStatusUpdateInterval(val *string)
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
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
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
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

// The jsii proxy struct for CfnJobTemplate
type jsiiProxy_CfnJobTemplate struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnJobTemplate) AccelerationSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accelerationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) Category() *string {
	var returns *string
	_jsii_.Get(
		j,
		"category",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) HopDestinations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hopDestinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) Queue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) SettingsJson() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"settingsJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) StatusUpdateInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusUpdateInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplate) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaConvert::JobTemplate`.
func NewCfnJobTemplate(scope constructs.Construct, id *string, props *CfnJobTemplateProps) CfnJobTemplate {
	_init_.Initialize()

	if err := validateNewCfnJobTemplateParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnJobTemplate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_mediaconvert.CfnJobTemplate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaConvert::JobTemplate`.
func NewCfnJobTemplate_Override(c CfnJobTemplate, scope constructs.Construct, id *string, props *CfnJobTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_mediaconvert.CfnJobTemplate",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnJobTemplate)SetAccelerationSettings(val interface{}) {
	if err := j.validateSetAccelerationSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accelerationSettings",
		val,
	)
}

func (j *jsiiProxy_CfnJobTemplate)SetCategory(val *string) {
	_jsii_.Set(
		j,
		"category",
		val,
	)
}

func (j *jsiiProxy_CfnJobTemplate)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnJobTemplate)SetHopDestinations(val interface{}) {
	if err := j.validateSetHopDestinationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hopDestinations",
		val,
	)
}

func (j *jsiiProxy_CfnJobTemplate)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnJobTemplate)SetPriority(val *float64) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_CfnJobTemplate)SetQueue(val *string) {
	_jsii_.Set(
		j,
		"queue",
		val,
	)
}

func (j *jsiiProxy_CfnJobTemplate)SetSettingsJson(val interface{}) {
	if err := j.validateSetSettingsJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"settingsJson",
		val,
	)
}

func (j *jsiiProxy_CfnJobTemplate)SetStatusUpdateInterval(val *string) {
	_jsii_.Set(
		j,
		"statusUpdateInterval",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnJobTemplate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnJobTemplate_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediaconvert.CfnJobTemplate",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnJobTemplate_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnJobTemplate_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediaconvert.CfnJobTemplate",
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
func CfnJobTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnJobTemplate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediaconvert.CfnJobTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnJobTemplate_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_mediaconvert.CfnJobTemplate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnJobTemplate) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnJobTemplate) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnJobTemplate) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnJobTemplate) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnJobTemplate) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnJobTemplate) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnJobTemplate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnJobTemplate) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJobTemplate) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnJobTemplate) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnJobTemplate) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnJobTemplate) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnJobTemplate) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJobTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJobTemplate) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}


package awsimagebuilder

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsimagebuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An image pipeline is the automation configuration for building secure OS images on AWS .
//
// The Image Builder image pipeline is associated with an image recipe that defines the build, validation, and test phases for an image build lifecycle. An image pipeline can be associated with an infrastructure configuration that defines where your image is built. You can define attributes, such as instance types, a subnet for your VPC, security groups, logging, and other infrastructure-related configurations. You can also associate your image pipeline with a distribution configuration to define how you would like to deploy your image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnImagePipeline := awscdk.Aws_imagebuilder.NewCfnImagePipeline(this, jsii.String("MyCfnImagePipeline"), &CfnImagePipelineProps{
//   	InfrastructureConfigurationArn: jsii.String("infrastructureConfigurationArn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ContainerRecipeArn: jsii.String("containerRecipeArn"),
//   	Description: jsii.String("description"),
//   	DistributionConfigurationArn: jsii.String("distributionConfigurationArn"),
//   	EnhancedImageMetadataEnabled: jsii.Boolean(false),
//   	ExecutionRole: jsii.String("executionRole"),
//   	ImageRecipeArn: jsii.String("imageRecipeArn"),
//   	ImageScanningConfiguration: &ImageScanningConfigurationProperty{
//   		EcrConfiguration: &EcrConfigurationProperty{
//   			ContainerTags: []*string{
//   				jsii.String("containerTags"),
//   			},
//   			RepositoryName: jsii.String("repositoryName"),
//   		},
//   		ImageScanningEnabled: jsii.Boolean(false),
//   	},
//   	ImageTestsConfiguration: &ImageTestsConfigurationProperty{
//   		ImageTestsEnabled: jsii.Boolean(false),
//   		TimeoutMinutes: jsii.Number(123),
//   	},
//   	Schedule: &ScheduleProperty{
//   		PipelineExecutionStartCondition: jsii.String("pipelineExecutionStartCondition"),
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   	Status: jsii.String("status"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Workflows: []interface{}{
//   		&WorkflowConfigurationProperty{
//   			OnFailure: jsii.String("onFailure"),
//   			ParallelGroup: jsii.String("parallelGroup"),
//   			Parameters: []interface{}{
//   				&WorkflowParameterProperty{
//   					Name: jsii.String("name"),
//   					Value: []*string{
//   						jsii.String("value"),
//   					},
//   				},
//   			},
//   			WorkflowArn: jsii.String("workflowArn"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-imagepipeline.html
//
type CfnImagePipeline interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// Returns the Amazon Resource Name (ARN) of the image pipeline.
	//
	// For example, `arn:aws:imagebuilder:us-west-2:123456789012:image-pipeline/mywindows2016pipeline` .
	AttrArn() *string
	// Returns the name of the image pipeline.
	AttrName() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The Amazon Resource Name (ARN) of the container recipe that is used for this pipeline.
	ContainerRecipeArn() *string
	SetContainerRecipeArn(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of this image pipeline.
	Description() *string
	SetDescription(val *string)
	// The Amazon Resource Name (ARN) of the distribution configuration associated with this image pipeline.
	DistributionConfigurationArn() *string
	SetDistributionConfigurationArn(val *string)
	// Collects additional information about the image being created, including the operating system (OS) version and package list.
	EnhancedImageMetadataEnabled() interface{}
	SetEnhancedImageMetadataEnabled(val interface{})
	// The name or Amazon Resource Name (ARN) for the IAM role you create that grants Image Builder access to perform workflow actions.
	ExecutionRole() *string
	SetExecutionRole(val *string)
	// The Amazon Resource Name (ARN) of the image recipe associated with this image pipeline.
	ImageRecipeArn() *string
	SetImageRecipeArn(val *string)
	// Contains settings for vulnerability scans.
	ImageScanningConfiguration() interface{}
	SetImageScanningConfiguration(val interface{})
	// The configuration of the image tests that run after image creation to ensure the quality of the image that was created.
	ImageTestsConfiguration() interface{}
	SetImageTestsConfiguration(val interface{})
	// The Amazon Resource Name (ARN) of the infrastructure configuration associated with this image pipeline.
	InfrastructureConfigurationArn() *string
	SetInfrastructureConfigurationArn(val *string)
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
	// The name of the image pipeline.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The schedule of the image pipeline.
	Schedule() interface{}
	SetSchedule(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The status of the image pipeline.
	Status() *string
	SetStatus(val *string)
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// The tags of this image pipeline.
	TagsRaw() *map[string]*string
	SetTagsRaw(val *map[string]*string)
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
	// Contains the workflows that run for the image pipeline.
	Workflows() interface{}
	SetWorkflows(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
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
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
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
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
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

// The jsii proxy struct for CfnImagePipeline
type jsiiProxy_CfnImagePipeline struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnImagePipeline) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) ContainerRecipeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRecipeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) DistributionConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) EnhancedImageMetadataEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedImageMetadataEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) ImageRecipeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageRecipeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) ImageScanningConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageScanningConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) ImageTestsConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageTestsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) InfrastructureConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) Schedule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) TagsRaw() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipeline) Workflows() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workflows",
		&returns,
	)
	return returns
}


func NewCfnImagePipeline(scope constructs.Construct, id *string, props *CfnImagePipelineProps) CfnImagePipeline {
	_init_.Initialize()

	if err := validateNewCfnImagePipelineParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnImagePipeline{}

	_jsii_.Create(
		"aws-cdk-lib.aws_imagebuilder.CfnImagePipeline",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnImagePipeline_Override(c CfnImagePipeline, scope constructs.Construct, id *string, props *CfnImagePipelineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_imagebuilder.CfnImagePipeline",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnImagePipeline)SetContainerRecipeArn(val *string) {
	_jsii_.Set(
		j,
		"containerRecipeArn",
		val,
	)
}

func (j *jsiiProxy_CfnImagePipeline)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnImagePipeline)SetDistributionConfigurationArn(val *string) {
	_jsii_.Set(
		j,
		"distributionConfigurationArn",
		val,
	)
}

func (j *jsiiProxy_CfnImagePipeline)SetEnhancedImageMetadataEnabled(val interface{}) {
	if err := j.validateSetEnhancedImageMetadataEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enhancedImageMetadataEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnImagePipeline)SetExecutionRole(val *string) {
	_jsii_.Set(
		j,
		"executionRole",
		val,
	)
}

func (j *jsiiProxy_CfnImagePipeline)SetImageRecipeArn(val *string) {
	_jsii_.Set(
		j,
		"imageRecipeArn",
		val,
	)
}

func (j *jsiiProxy_CfnImagePipeline)SetImageScanningConfiguration(val interface{}) {
	if err := j.validateSetImageScanningConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageScanningConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnImagePipeline)SetImageTestsConfiguration(val interface{}) {
	if err := j.validateSetImageTestsConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageTestsConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnImagePipeline)SetInfrastructureConfigurationArn(val *string) {
	if err := j.validateSetInfrastructureConfigurationArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"infrastructureConfigurationArn",
		val,
	)
}

func (j *jsiiProxy_CfnImagePipeline)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnImagePipeline)SetSchedule(val interface{}) {
	if err := j.validateSetScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_CfnImagePipeline)SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_CfnImagePipeline)SetTagsRaw(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnImagePipeline)SetWorkflows(val interface{}) {
	if err := j.validateSetWorkflowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workflows",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnImagePipeline_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnImagePipeline_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnImagePipeline",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnImagePipeline_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnImagePipeline_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnImagePipeline",
		"isCfnResource",
		[]interface{}{x},
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
func CfnImagePipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnImagePipeline_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnImagePipeline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnImagePipeline_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_imagebuilder.CfnImagePipeline",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnImagePipeline) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnImagePipeline) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnImagePipeline) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnImagePipeline) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnImagePipeline) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnImagePipeline) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnImagePipeline) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnImagePipeline) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnImagePipeline) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnImagePipeline) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnImagePipeline) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnImagePipeline) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImagePipeline) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImagePipeline) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnImagePipeline) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnImagePipeline) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnImagePipeline) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnImagePipeline) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImagePipeline) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImagePipeline) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}


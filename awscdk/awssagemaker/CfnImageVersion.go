package awssagemaker

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a version of the SageMaker image specified by `ImageName` .
//
// The version represents the Amazon Container Registry (ECR) container image specified by `BaseImage` .
//
// > You can use the `DependsOn` attribute to specify that the creation of a specific resource follows another. You can use it for the following use cases. For more information, see [`DependsOn` attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) .
// >
// > 1. `DependsOn` can be used to establish a parent/child relationship between `ImageVersion` and `Image` where the `ImageVersion` `DependsOn` the `Image` .
// >
// > 2. `DependsOn` can be used to establish order among `ImageVersion` s within the same `Image` namespace. For example, if ImageVersionB `DependsOn` ImageVersionA and both share the same parent `Image` , then ImageVersionA is version N and ImageVersionB is N+1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnImageVersion := awscdk.Aws_sagemaker.NewCfnImageVersion(this, jsii.String("MyCfnImageVersion"), &CfnImageVersionProps{
//   	BaseImage: jsii.String("baseImage"),
//   	ImageName: jsii.String("imageName"),
//
//   	// the properties below are optional
//   	Alias: jsii.String("alias"),
//   	Aliases: []*string{
//   		jsii.String("aliases"),
//   	},
//   	Horovod: jsii.Boolean(false),
//   	JobType: jsii.String("jobType"),
//   	MlFramework: jsii.String("mlFramework"),
//   	Processor: jsii.String("processor"),
//   	ProgrammingLang: jsii.String("programmingLang"),
//   	ReleaseNotes: jsii.String("releaseNotes"),
//   	VendorGuidance: jsii.String("vendorGuidance"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html
//
type CfnImageVersion interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The alias of the image version.
	Alias() *string
	SetAlias(val *string)
	// List of aliases for the image version.
	Aliases() *[]*string
	SetAliases(val *[]*string)
	// The URI of the container image version referenced by ImageVersion.
	AttrContainerImage() *string
	// The Amazon Resource Name (ARN) of the parent Image.
	AttrImageArn() *string
	// The Amazon Resource Name (ARN) of the image version.
	//
	// *Type* : String
	//
	// *Length Constraints* : Maximum length of 256.
	//
	// *Pattern* : `^arn:aws(-[\w]+)*:sagemaker:.+:[0-9]{12}:image-version/[a-z0-9]([-.]?[a-z0-9])* /[0-9]+$`
	AttrImageVersionArn() *string
	// The version of the image.
	AttrVersion() *float64
	// The container image that the SageMaker image version is based on.
	BaseImage() *string
	SetBaseImage(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Indicates Horovod compatibility.
	Horovod() interface{}
	SetHorovod(val interface{})
	// The name of the parent image.
	ImageName() *string
	SetImageName(val *string)
	// Indicates SageMaker job type compatibility.
	JobType() *string
	SetJobType(val *string)
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
	// The machine learning framework vended in the image version.
	MlFramework() *string
	SetMlFramework(val *string)
	// The tree node.
	Node() constructs.Node
	// Indicates CPU or GPU compatibility.
	Processor() *string
	SetProcessor(val *string)
	// The supported programming language and its version.
	ProgrammingLang() *string
	SetProgrammingLang(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The maintainer description of the image version.
	ReleaseNotes() *string
	SetReleaseNotes(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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
	// The availability of the image version specified by the maintainer.
	VendorGuidance() *string
	SetVendorGuidance(val *string)
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

// The jsii proxy struct for CfnImageVersion
type jsiiProxy_CfnImageVersion struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnImageVersion) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) Aliases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) AttrContainerImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrContainerImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) AttrImageArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrImageArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) AttrImageVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrImageVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) AttrVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) BaseImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) Horovod() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"horovod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) JobType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) MlFramework() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlFramework",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) Processor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) ProgrammingLang() *string {
	var returns *string
	_jsii_.Get(
		j,
		"programmingLang",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) ReleaseNotes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseNotes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageVersion) VendorGuidance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vendorGuidance",
		&returns,
	)
	return returns
}


func NewCfnImageVersion(scope constructs.Construct, id *string, props *CfnImageVersionProps) CfnImageVersion {
	_init_.Initialize()

	if err := validateNewCfnImageVersionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnImageVersion{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sagemaker.CfnImageVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnImageVersion_Override(c CfnImageVersion, scope constructs.Construct, id *string, props *CfnImageVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sagemaker.CfnImageVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnImageVersion)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_CfnImageVersion)SetAliases(val *[]*string) {
	_jsii_.Set(
		j,
		"aliases",
		val,
	)
}

func (j *jsiiProxy_CfnImageVersion)SetBaseImage(val *string) {
	if err := j.validateSetBaseImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseImage",
		val,
	)
}

func (j *jsiiProxy_CfnImageVersion)SetHorovod(val interface{}) {
	if err := j.validateSetHorovodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"horovod",
		val,
	)
}

func (j *jsiiProxy_CfnImageVersion)SetImageName(val *string) {
	if err := j.validateSetImageNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageName",
		val,
	)
}

func (j *jsiiProxy_CfnImageVersion)SetJobType(val *string) {
	_jsii_.Set(
		j,
		"jobType",
		val,
	)
}

func (j *jsiiProxy_CfnImageVersion)SetMlFramework(val *string) {
	_jsii_.Set(
		j,
		"mlFramework",
		val,
	)
}

func (j *jsiiProxy_CfnImageVersion)SetProcessor(val *string) {
	_jsii_.Set(
		j,
		"processor",
		val,
	)
}

func (j *jsiiProxy_CfnImageVersion)SetProgrammingLang(val *string) {
	_jsii_.Set(
		j,
		"programmingLang",
		val,
	)
}

func (j *jsiiProxy_CfnImageVersion)SetReleaseNotes(val *string) {
	_jsii_.Set(
		j,
		"releaseNotes",
		val,
	)
}

func (j *jsiiProxy_CfnImageVersion)SetVendorGuidance(val *string) {
	_jsii_.Set(
		j,
		"vendorGuidance",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnImageVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnImageVersion_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sagemaker.CfnImageVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnImageVersion_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnImageVersion_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sagemaker.CfnImageVersion",
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
func CfnImageVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnImageVersion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sagemaker.CfnImageVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnImageVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_sagemaker.CfnImageVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnImageVersion) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnImageVersion) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnImageVersion) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnImageVersion) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnImageVersion) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnImageVersion) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnImageVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnImageVersion) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnImageVersion) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnImageVersion) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnImageVersion) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnImageVersion) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImageVersion) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImageVersion) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnImageVersion) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnImageVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnImageVersion) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnImageVersion) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImageVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImageVersion) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}


package awsimagebuilder

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsimagebuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::ImageBuilder::Component`.
//
// Creates a new component that can be used to build, validate, test, and assess your image. The component is based on a YAML document that you specify using exactly one of the following methods:
//
// - Inline, using the `data` property in the request body.
// - A URL that points to a YAML document file stored in Amazon S3, using the `uri` property in the request body.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnComponent := awscdk.Aws_imagebuilder.NewCfnComponent(this, jsii.String("MyCfnComponent"), &cfnComponentProps{
//   	name: jsii.String("name"),
//   	platform: jsii.String("platform"),
//   	version: jsii.String("version"),
//
//   	// the properties below are optional
//   	changeDescription: jsii.String("changeDescription"),
//   	data: jsii.String("data"),
//   	description: jsii.String("description"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	supportedOsVersions: []*string{
//   		jsii.String("supportedOsVersions"),
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	uri: jsii.String("uri"),
//   })
//
type CfnComponent interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Returns the Amazon Resource Name (ARN) of the component.
	//
	// The following pattern is applied: `^arn:aws[^:]*:imagebuilder:[^:]+:(?:\d{12}|aws):(?:image-recipe|infrastructure-configuration|distribution-configuration|component|image|image-pipeline)/[a-z0-9-_]+(?:/(?:(?:x|\d+)\.(?:x|\d+)\.(?:x|\d+))(?:/\d+)?)?$` .
	AttrArn() *string
	// Returns the encryption status of the component.
	//
	// For example `true` or `false` .
	AttrEncrypted() awscdk.IResolvable
	// Returns the name of the component.
	AttrName() *string
	// Image Builder determines the component type based on the phases that are defined in the component document.
	//
	// If there is only one phase, and its name is "test", then the type is `TEST` . For all other components, the type is `BUILD` .
	AttrType() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The change description of the component.
	//
	// Describes what change has been made in this version, or what makes this version different from other versions of this component.
	ChangeDescription() *string
	SetChangeDescription(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Component `data` contains inline YAML document content for the component.
	//
	// Alternatively, you can specify the `uri` of a YAML document file stored in Amazon S3. However, you cannot specify both properties.
	Data() *string
	SetData(val *string)
	// The description of the component.
	//
	// Describes the contents of the component.
	Description() *string
	SetDescription(val *string)
	// The ID of the KMS key that should be used to encrypt this component.
	KmsKeyId() *string
	SetKmsKeyId(val *string)
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
	// The name of the component.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The platform of the component.
	Platform() *string
	SetPlatform(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The operating system (OS) version supported by the component.
	//
	// If the OS information is available, a prefix match is performed against the base image OS version during image recipe creation.
	SupportedOsVersions() *[]*string
	SetSupportedOsVersions(val *[]*string)
	// The tags of the component.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// The `uri` of a YAML component document file.
	//
	// This must be an S3 URL ( `s3://bucket/key` ), and the requester must have permission to access the S3 bucket it points to. If you use Amazon S3, you can specify component content up to your service quota.
	//
	// Alternatively, you can specify the YAML document inline, using the component `data` property. You cannot specify both properties.
	Uri() *string
	SetUri(val *string)
	// The component version.
	//
	// For example, `1.0.0` .
	Version() *string
	SetVersion(val *string)
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
	GetAtt(attributeName *string) awscdk.Reference
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

// The jsii proxy struct for CfnComponent
type jsiiProxy_CfnComponent struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnComponent) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) AttrEncrypted() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) AttrType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) ChangeDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changeDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Data() *string {
	var returns *string
	_jsii_.Get(
		j,
		"data",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) SupportedOsVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedOsVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Create a new `AWS::ImageBuilder::Component`.
func NewCfnComponent(scope constructs.Construct, id *string, props *CfnComponentProps) CfnComponent {
	_init_.Initialize()

	j := jsiiProxy_CfnComponent{}

	_jsii_.Create(
		"aws-cdk-lib.aws_imagebuilder.CfnComponent",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ImageBuilder::Component`.
func NewCfnComponent_Override(c CfnComponent, scope constructs.Construct, id *string, props *CfnComponentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_imagebuilder.CfnComponent",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnComponent) SetChangeDescription(val *string) {
	_jsii_.Set(
		j,
		"changeDescription",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetData(val *string) {
	_jsii_.Set(
		j,
		"data",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetPlatform(val *string) {
	_jsii_.Set(
		j,
		"platform",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetSupportedOsVersions(val *[]*string) {
	_jsii_.Set(
		j,
		"supportedOsVersions",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetUri(val *string) {
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

func (j *jsiiProxy_CfnComponent) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnComponent_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnComponent",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnComponent_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnComponent",
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
func CfnComponent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnComponent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnComponent_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_imagebuilder.CfnComponent",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnComponent) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnComponent) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnComponent) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnComponent) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnComponent) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnComponent) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnComponent) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnComponent) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponent) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponent) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnComponent) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnComponent) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponent) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponent) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnComponent`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnComponentProps := &cfnComponentProps{
//   	name: jsii.String("name"),
//   	platform: jsii.String("platform"),
//   	version: jsii.String("version"),
//
//   	// the properties below are optional
//   	changeDescription: jsii.String("changeDescription"),
//   	data: jsii.String("data"),
//   	description: jsii.String("description"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	supportedOsVersions: []*string{
//   		jsii.String("supportedOsVersions"),
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	uri: jsii.String("uri"),
//   }
//
type CfnComponentProps struct {
	// The name of the component.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The platform of the component.
	Platform *string `field:"required" json:"platform" yaml:"platform"`
	// The component version.
	//
	// For example, `1.0.0` .
	Version *string `field:"required" json:"version" yaml:"version"`
	// The change description of the component.
	//
	// Describes what change has been made in this version, or what makes this version different from other versions of this component.
	ChangeDescription *string `field:"optional" json:"changeDescription" yaml:"changeDescription"`
	// Component `data` contains inline YAML document content for the component.
	//
	// Alternatively, you can specify the `uri` of a YAML document file stored in Amazon S3. However, you cannot specify both properties.
	Data *string `field:"optional" json:"data" yaml:"data"`
	// The description of the component.
	//
	// Describes the contents of the component.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the KMS key that should be used to encrypt this component.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The operating system (OS) version supported by the component.
	//
	// If the OS information is available, a prefix match is performed against the base image OS version during image recipe creation.
	SupportedOsVersions *[]*string `field:"optional" json:"supportedOsVersions" yaml:"supportedOsVersions"`
	// The tags of the component.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The `uri` of a YAML component document file.
	//
	// This must be an S3 URL ( `s3://bucket/key` ), and the requester must have permission to access the S3 bucket it points to. If you use Amazon S3, you can specify component content up to your service quota.
	//
	// Alternatively, you can specify the YAML document inline, using the component `data` property. You cannot specify both properties.
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

// A CloudFormation `AWS::ImageBuilder::ContainerRecipe`.
//
// Creates a new container recipe. Container recipes define how images are configured, tested, and assessed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContainerRecipe := awscdk.Aws_imagebuilder.NewCfnContainerRecipe(this, jsii.String("MyCfnContainerRecipe"), &cfnContainerRecipeProps{
//   	components: []interface{}{
//   		&componentConfigurationProperty{
//   			componentArn: jsii.String("componentArn"),
//   		},
//   	},
//   	containerType: jsii.String("containerType"),
//   	name: jsii.String("name"),
//   	parentImage: jsii.String("parentImage"),
//   	targetRepository: &targetContainerRepositoryProperty{
//   		repositoryName: jsii.String("repositoryName"),
//   		service: jsii.String("service"),
//   	},
//   	version: jsii.String("version"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	dockerfileTemplateData: jsii.String("dockerfileTemplateData"),
//   	dockerfileTemplateUri: jsii.String("dockerfileTemplateUri"),
//   	imageOsVersionOverride: jsii.String("imageOsVersionOverride"),
//   	instanceConfiguration: &instanceConfigurationProperty{
//   		blockDeviceMappings: []interface{}{
//   			&instanceBlockDeviceMappingProperty{
//   				deviceName: jsii.String("deviceName"),
//   				ebs: &ebsInstanceBlockDeviceSpecificationProperty{
//   					deleteOnTermination: jsii.Boolean(false),
//   					encrypted: jsii.Boolean(false),
//   					iops: jsii.Number(123),
//   					kmsKeyId: jsii.String("kmsKeyId"),
//   					snapshotId: jsii.String("snapshotId"),
//   					throughput: jsii.Number(123),
//   					volumeSize: jsii.Number(123),
//   					volumeType: jsii.String("volumeType"),
//   				},
//   				noDevice: jsii.String("noDevice"),
//   				virtualName: jsii.String("virtualName"),
//   			},
//   		},
//   		image: jsii.String("image"),
//   	},
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	platformOverride: jsii.String("platformOverride"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	workingDirectory: jsii.String("workingDirectory"),
//   })
//
type CfnContainerRecipe interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Returns the Amazon Resource Name (ARN) of the container recipe.
	//
	// For example, `arn:aws:imagebuilder:us-east-1:123456789012:container-recipe/mybasicrecipe/2020.12.17` .
	AttrArn() *string
	// Returns the name of the container recipe.
	AttrName() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Build and test components that are included in the container recipe.
	//
	// Recipes require a minimum of one build component, and can have a maximum of 20 build and test components in any combination.
	Components() interface{}
	SetComponents(val interface{})
	// Specifies the type of container, such as Docker.
	ContainerType() *string
	SetContainerType(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the container recipe.
	Description() *string
	SetDescription(val *string)
	// Dockerfiles are text documents that are used to build Docker containers, and ensure that they contain all of the elements required by the application running inside.
	//
	// The template data consists of contextual variables where Image Builder places build information or scripts, based on your container image recipe.
	DockerfileTemplateData() *string
	SetDockerfileTemplateData(val *string)
	// The S3 URI for the Dockerfile that will be used to build your container image.
	DockerfileTemplateUri() *string
	SetDockerfileTemplateUri(val *string)
	// Specifies the operating system version for the base image.
	ImageOsVersionOverride() *string
	SetImageOsVersionOverride(val *string)
	// A group of options that can be used to configure an instance for building and testing container images.
	InstanceConfiguration() interface{}
	SetInstanceConfiguration(val interface{})
	// Identifies which KMS key is used to encrypt the container image for distribution to the target Region.
	KmsKeyId() *string
	SetKmsKeyId(val *string)
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
	// The name of the container recipe.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The base image for the container recipe.
	ParentImage() *string
	SetParentImage(val *string)
	// Specifies the operating system platform when you use a custom base image.
	PlatformOverride() *string
	SetPlatformOverride(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tags that are attached to the container recipe.
	Tags() awscdk.TagManager
	// The destination repository for the container image.
	TargetRepository() interface{}
	SetTargetRepository(val interface{})
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// The semantic version of the container recipe.
	//
	// > The semantic version has four nodes: <major>.<minor>.<patch>/<build>. You can assign values for the first three, and can filter on all of them.
	// >
	// > *Assignment:* For the first three nodes you can assign any positive integer value, including zero, with an upper limit of 2^30-1, or 1073741823 for each node. Image Builder automatically assigns the build number to the fourth node.
	// >
	// > *Patterns:* You can use any numeric pattern that adheres to the assignment requirements for the nodes that you can assign. For example, you might choose a software version pattern, such as 1.0.0, or a date, such as 2021.01.01.
	// >
	// > *Filtering:* With semantic versioning, you have the flexibility to use wildcards (x) to specify the most recent versions or nodes when selecting the base image or components for your recipe. When you use a wildcard in any node, all nodes to the right of the first wildcard must also be wildcards.
	Version() *string
	SetVersion(val *string)
	// The working directory for use during build and test workflows.
	WorkingDirectory() *string
	SetWorkingDirectory(val *string)
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
	GetAtt(attributeName *string) awscdk.Reference
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

// The jsii proxy struct for CfnContainerRecipe
type jsiiProxy_CfnContainerRecipe struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnContainerRecipe) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) Components() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) ContainerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) DockerfileTemplateData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfileTemplateData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) DockerfileTemplateUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfileTemplateUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) ImageOsVersionOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageOsVersionOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) InstanceConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) ParentImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) PlatformOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) TargetRepository() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) WorkingDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirectory",
		&returns,
	)
	return returns
}


// Create a new `AWS::ImageBuilder::ContainerRecipe`.
func NewCfnContainerRecipe(scope constructs.Construct, id *string, props *CfnContainerRecipeProps) CfnContainerRecipe {
	_init_.Initialize()

	j := jsiiProxy_CfnContainerRecipe{}

	_jsii_.Create(
		"aws-cdk-lib.aws_imagebuilder.CfnContainerRecipe",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ImageBuilder::ContainerRecipe`.
func NewCfnContainerRecipe_Override(c CfnContainerRecipe, scope constructs.Construct, id *string, props *CfnContainerRecipeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_imagebuilder.CfnContainerRecipe",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnContainerRecipe) SetComponents(val interface{}) {
	_jsii_.Set(
		j,
		"components",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe) SetContainerType(val *string) {
	_jsii_.Set(
		j,
		"containerType",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe) SetDockerfileTemplateData(val *string) {
	_jsii_.Set(
		j,
		"dockerfileTemplateData",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe) SetDockerfileTemplateUri(val *string) {
	_jsii_.Set(
		j,
		"dockerfileTemplateUri",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe) SetImageOsVersionOverride(val *string) {
	_jsii_.Set(
		j,
		"imageOsVersionOverride",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe) SetInstanceConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"instanceConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe) SetParentImage(val *string) {
	_jsii_.Set(
		j,
		"parentImage",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe) SetPlatformOverride(val *string) {
	_jsii_.Set(
		j,
		"platformOverride",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe) SetTargetRepository(val interface{}) {
	_jsii_.Set(
		j,
		"targetRepository",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe) SetWorkingDirectory(val *string) {
	_jsii_.Set(
		j,
		"workingDirectory",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnContainerRecipe_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnContainerRecipe",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnContainerRecipe_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnContainerRecipe",
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
func CfnContainerRecipe_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnContainerRecipe",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnContainerRecipe_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_imagebuilder.CfnContainerRecipe",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnContainerRecipe) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnContainerRecipe) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnContainerRecipe) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnContainerRecipe) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnContainerRecipe) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnContainerRecipe) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnContainerRecipe) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnContainerRecipe) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnContainerRecipe) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnContainerRecipe) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnContainerRecipe) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnContainerRecipe) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnContainerRecipe) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnContainerRecipe) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnContainerRecipe) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Configuration details of the component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   componentConfigurationProperty := &componentConfigurationProperty{
//   	componentArn: jsii.String("componentArn"),
//   }
//
type CfnContainerRecipe_ComponentConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the component.
	ComponentArn *string `field:"optional" json:"componentArn" yaml:"componentArn"`
}

// Amazon EBS-specific block device mapping specifications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ebsInstanceBlockDeviceSpecificationProperty := &ebsInstanceBlockDeviceSpecificationProperty{
//   	deleteOnTermination: jsii.Boolean(false),
//   	encrypted: jsii.Boolean(false),
//   	iops: jsii.Number(123),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	snapshotId: jsii.String("snapshotId"),
//   	throughput: jsii.Number(123),
//   	volumeSize: jsii.Number(123),
//   	volumeType: jsii.String("volumeType"),
//   }
//
type CfnContainerRecipe_EbsInstanceBlockDeviceSpecificationProperty struct {
	// Use to configure delete on termination of the associated device.
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Use to configure device encryption.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// Use to configure device IOPS.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Use to configure the KMS key to use when encrypting the device.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The snapshot that defines the device contents.
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// *For GP3 volumes only* – The throughput in MiB/s that the volume supports.
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// Use to override the device's volume size.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// Use to override the device's volume type.
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

// Defines block device mappings for the instance used to configure your image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceBlockDeviceMappingProperty := &instanceBlockDeviceMappingProperty{
//   	deviceName: jsii.String("deviceName"),
//   	ebs: &ebsInstanceBlockDeviceSpecificationProperty{
//   		deleteOnTermination: jsii.Boolean(false),
//   		encrypted: jsii.Boolean(false),
//   		iops: jsii.Number(123),
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   		snapshotId: jsii.String("snapshotId"),
//   		throughput: jsii.Number(123),
//   		volumeSize: jsii.Number(123),
//   		volumeType: jsii.String("volumeType"),
//   	},
//   	noDevice: jsii.String("noDevice"),
//   	virtualName: jsii.String("virtualName"),
//   }
//
type CfnContainerRecipe_InstanceBlockDeviceMappingProperty struct {
	// The device to which these mappings apply.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// Use to manage Amazon EBS-specific configuration for this mapping.
	Ebs interface{} `field:"optional" json:"ebs" yaml:"ebs"`
	// Use to remove a mapping from the base image.
	NoDevice *string `field:"optional" json:"noDevice" yaml:"noDevice"`
	// Use to manage instance ephemeral devices.
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}

// Defines a custom base AMI and block device mapping configurations of an instance used for building and testing container images.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceConfigurationProperty := &instanceConfigurationProperty{
//   	blockDeviceMappings: []interface{}{
//   		&instanceBlockDeviceMappingProperty{
//   			deviceName: jsii.String("deviceName"),
//   			ebs: &ebsInstanceBlockDeviceSpecificationProperty{
//   				deleteOnTermination: jsii.Boolean(false),
//   				encrypted: jsii.Boolean(false),
//   				iops: jsii.Number(123),
//   				kmsKeyId: jsii.String("kmsKeyId"),
//   				snapshotId: jsii.String("snapshotId"),
//   				throughput: jsii.Number(123),
//   				volumeSize: jsii.Number(123),
//   				volumeType: jsii.String("volumeType"),
//   			},
//   			noDevice: jsii.String("noDevice"),
//   			virtualName: jsii.String("virtualName"),
//   		},
//   	},
//   	image: jsii.String("image"),
//   }
//
type CfnContainerRecipe_InstanceConfigurationProperty struct {
	// Defines the block devices to attach for building an instance from this Image Builder AMI.
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// The AMI ID to use as the base image for a container build and test instance.
	//
	// If not specified, Image Builder will use the appropriate ECS-optimized AMI as a base image.
	Image *string `field:"optional" json:"image" yaml:"image"`
}

// The container repository where the output container image is stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetContainerRepositoryProperty := &targetContainerRepositoryProperty{
//   	repositoryName: jsii.String("repositoryName"),
//   	service: jsii.String("service"),
//   }
//
type CfnContainerRecipe_TargetContainerRepositoryProperty struct {
	// The name of the container repository where the output container image is stored.
	//
	// This name is prefixed by the repository location.
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// Specifies the service in which this image was registered.
	Service *string `field:"optional" json:"service" yaml:"service"`
}

// Properties for defining a `CfnContainerRecipe`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContainerRecipeProps := &cfnContainerRecipeProps{
//   	components: []interface{}{
//   		&componentConfigurationProperty{
//   			componentArn: jsii.String("componentArn"),
//   		},
//   	},
//   	containerType: jsii.String("containerType"),
//   	name: jsii.String("name"),
//   	parentImage: jsii.String("parentImage"),
//   	targetRepository: &targetContainerRepositoryProperty{
//   		repositoryName: jsii.String("repositoryName"),
//   		service: jsii.String("service"),
//   	},
//   	version: jsii.String("version"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	dockerfileTemplateData: jsii.String("dockerfileTemplateData"),
//   	dockerfileTemplateUri: jsii.String("dockerfileTemplateUri"),
//   	imageOsVersionOverride: jsii.String("imageOsVersionOverride"),
//   	instanceConfiguration: &instanceConfigurationProperty{
//   		blockDeviceMappings: []interface{}{
//   			&instanceBlockDeviceMappingProperty{
//   				deviceName: jsii.String("deviceName"),
//   				ebs: &ebsInstanceBlockDeviceSpecificationProperty{
//   					deleteOnTermination: jsii.Boolean(false),
//   					encrypted: jsii.Boolean(false),
//   					iops: jsii.Number(123),
//   					kmsKeyId: jsii.String("kmsKeyId"),
//   					snapshotId: jsii.String("snapshotId"),
//   					throughput: jsii.Number(123),
//   					volumeSize: jsii.Number(123),
//   					volumeType: jsii.String("volumeType"),
//   				},
//   				noDevice: jsii.String("noDevice"),
//   				virtualName: jsii.String("virtualName"),
//   			},
//   		},
//   		image: jsii.String("image"),
//   	},
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	platformOverride: jsii.String("platformOverride"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	workingDirectory: jsii.String("workingDirectory"),
//   }
//
type CfnContainerRecipeProps struct {
	// Build and test components that are included in the container recipe.
	//
	// Recipes require a minimum of one build component, and can have a maximum of 20 build and test components in any combination.
	Components interface{} `field:"required" json:"components" yaml:"components"`
	// Specifies the type of container, such as Docker.
	ContainerType *string `field:"required" json:"containerType" yaml:"containerType"`
	// The name of the container recipe.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The base image for the container recipe.
	ParentImage *string `field:"required" json:"parentImage" yaml:"parentImage"`
	// The destination repository for the container image.
	TargetRepository interface{} `field:"required" json:"targetRepository" yaml:"targetRepository"`
	// The semantic version of the container recipe.
	//
	// > The semantic version has four nodes: <major>.<minor>.<patch>/<build>. You can assign values for the first three, and can filter on all of them.
	// >
	// > *Assignment:* For the first three nodes you can assign any positive integer value, including zero, with an upper limit of 2^30-1, or 1073741823 for each node. Image Builder automatically assigns the build number to the fourth node.
	// >
	// > *Patterns:* You can use any numeric pattern that adheres to the assignment requirements for the nodes that you can assign. For example, you might choose a software version pattern, such as 1.0.0, or a date, such as 2021.01.01.
	// >
	// > *Filtering:* With semantic versioning, you have the flexibility to use wildcards (x) to specify the most recent versions or nodes when selecting the base image or components for your recipe. When you use a wildcard in any node, all nodes to the right of the first wildcard must also be wildcards.
	Version *string `field:"required" json:"version" yaml:"version"`
	// The description of the container recipe.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Dockerfiles are text documents that are used to build Docker containers, and ensure that they contain all of the elements required by the application running inside.
	//
	// The template data consists of contextual variables where Image Builder places build information or scripts, based on your container image recipe.
	DockerfileTemplateData *string `field:"optional" json:"dockerfileTemplateData" yaml:"dockerfileTemplateData"`
	// The S3 URI for the Dockerfile that will be used to build your container image.
	DockerfileTemplateUri *string `field:"optional" json:"dockerfileTemplateUri" yaml:"dockerfileTemplateUri"`
	// Specifies the operating system version for the base image.
	ImageOsVersionOverride *string `field:"optional" json:"imageOsVersionOverride" yaml:"imageOsVersionOverride"`
	// A group of options that can be used to configure an instance for building and testing container images.
	InstanceConfiguration interface{} `field:"optional" json:"instanceConfiguration" yaml:"instanceConfiguration"`
	// Identifies which KMS key is used to encrypt the container image for distribution to the target Region.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies the operating system platform when you use a custom base image.
	PlatformOverride *string `field:"optional" json:"platformOverride" yaml:"platformOverride"`
	// Tags that are attached to the container recipe.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The working directory for use during build and test workflows.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

// A CloudFormation `AWS::ImageBuilder::DistributionConfiguration`.
//
// A distribution configuration allows you to specify the name and description of your output AMI, authorize other AWS account s to launch the AMI, and replicate the AMI to other AWS Regions . It also allows you to export the AMI to Amazon S3 .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var amiDistributionConfiguration interface{}
//   var containerDistributionConfiguration interface{}
//
//   cfnDistributionConfiguration := awscdk.Aws_imagebuilder.NewCfnDistributionConfiguration(this, jsii.String("MyCfnDistributionConfiguration"), &cfnDistributionConfigurationProps{
//   	distributions: []interface{}{
//   		&distributionProperty{
//   			region: jsii.String("region"),
//
//   			// the properties below are optional
//   			amiDistributionConfiguration: amiDistributionConfiguration,
//   			containerDistributionConfiguration: containerDistributionConfiguration,
//   			fastLaunchConfigurations: []interface{}{
//   				&fastLaunchConfigurationProperty{
//   					accountId: jsii.String("accountId"),
//   					enabled: jsii.Boolean(false),
//   					launchTemplate: &fastLaunchLaunchTemplateSpecificationProperty{
//   						launchTemplateId: jsii.String("launchTemplateId"),
//   						launchTemplateName: jsii.String("launchTemplateName"),
//   						launchTemplateVersion: jsii.String("launchTemplateVersion"),
//   					},
//   					maxParallelLaunches: jsii.Number(123),
//   					snapshotConfiguration: &fastLaunchSnapshotConfigurationProperty{
//   						targetResourceCount: jsii.Number(123),
//   					},
//   				},
//   			},
//   			launchTemplateConfigurations: []interface{}{
//   				&launchTemplateConfigurationProperty{
//   					accountId: jsii.String("accountId"),
//   					launchTemplateId: jsii.String("launchTemplateId"),
//   					setDefaultVersion: jsii.Boolean(false),
//   				},
//   			},
//   			licenseConfigurationArns: []*string{
//   				jsii.String("licenseConfigurationArns"),
//   			},
//   		},
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
type CfnDistributionConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Returns the Amazon Resource Name (ARN) of this distribution configuration.
	//
	// The following pattern is applied: `^arn:aws[^:]*:imagebuilder:[^:]+:(?:\d{12}|aws):(?:image-recipe|infrastructure-configuration|distribution-configuration|component|image|image-pipeline)/[a-z0-9-_]+(?:/(?:(?:x|\d+)\.(?:x|\d+)\.(?:x|\d+))(?:/\d+)?)?$` .
	AttrArn() *string
	// Returns the name of the distribution configuration.
	AttrName() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of this distribution configuration.
	Description() *string
	SetDescription(val *string)
	// The distributions of this distribution configuration formatted as an array of Distribution objects.
	Distributions() interface{}
	SetDistributions(val interface{})
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
	// The name of this distribution configuration.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags of this distribution configuration.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
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
	GetAtt(attributeName *string) awscdk.Reference
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

// The jsii proxy struct for CfnDistributionConfiguration
type jsiiProxy_CfnDistributionConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDistributionConfiguration) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistributionConfiguration) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistributionConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistributionConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistributionConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistributionConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistributionConfiguration) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistributionConfiguration) Distributions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"distributions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistributionConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistributionConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistributionConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistributionConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistributionConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistributionConfiguration) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistributionConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ImageBuilder::DistributionConfiguration`.
func NewCfnDistributionConfiguration(scope constructs.Construct, id *string, props *CfnDistributionConfigurationProps) CfnDistributionConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnDistributionConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_imagebuilder.CfnDistributionConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ImageBuilder::DistributionConfiguration`.
func NewCfnDistributionConfiguration_Override(c CfnDistributionConfiguration, scope constructs.Construct, id *string, props *CfnDistributionConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_imagebuilder.CfnDistributionConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDistributionConfiguration) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnDistributionConfiguration) SetDistributions(val interface{}) {
	_jsii_.Set(
		j,
		"distributions",
		val,
	)
}

func (j *jsiiProxy_CfnDistributionConfiguration) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDistributionConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnDistributionConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDistributionConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnDistributionConfiguration",
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
func CfnDistributionConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnDistributionConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDistributionConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_imagebuilder.CfnDistributionConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDistributionConfiguration) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDistributionConfiguration) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDistributionConfiguration) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDistributionConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDistributionConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDistributionConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDistributionConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDistributionConfiguration) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDistributionConfiguration) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDistributionConfiguration) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDistributionConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDistributionConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDistributionConfiguration) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDistributionConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDistributionConfiguration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Define and configure the output AMIs of the pipeline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   amiDistributionConfigurationProperty := &amiDistributionConfigurationProperty{
//   	amiTags: map[string]*string{
//   		"amiTagsKey": jsii.String("amiTags"),
//   	},
//   	description: jsii.String("description"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	launchPermissionConfiguration: &launchPermissionConfigurationProperty{
//   		organizationalUnitArns: []*string{
//   			jsii.String("organizationalUnitArns"),
//   		},
//   		organizationArns: []*string{
//   			jsii.String("organizationArns"),
//   		},
//   		userGroups: []*string{
//   			jsii.String("userGroups"),
//   		},
//   		userIds: []*string{
//   			jsii.String("userIds"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	targetAccountIds: []*string{
//   		jsii.String("targetAccountIds"),
//   	},
//   }
//
type CfnDistributionConfiguration_AmiDistributionConfigurationProperty struct {
	// The tags to apply to AMIs distributed to this Region.
	AmiTags interface{} `field:"optional" json:"amiTags" yaml:"amiTags"`
	// The description of the AMI distribution configuration.
	//
	// Minimum and maximum length are in characters.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The KMS key identifier used to encrypt the distributed image.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Launch permissions can be used to configure which AWS account s can use the AMI to launch instances.
	LaunchPermissionConfiguration interface{} `field:"optional" json:"launchPermissionConfiguration" yaml:"launchPermissionConfiguration"`
	// The name of the output AMI.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ID of an account to which you want to distribute an image.
	TargetAccountIds *[]*string `field:"optional" json:"targetAccountIds" yaml:"targetAccountIds"`
}

// Container distribution settings for encryption, licensing, and sharing in a specific Region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerDistributionConfigurationProperty := &containerDistributionConfigurationProperty{
//   	containerTags: []*string{
//   		jsii.String("containerTags"),
//   	},
//   	description: jsii.String("description"),
//   	targetRepository: &targetContainerRepositoryProperty{
//   		repositoryName: jsii.String("repositoryName"),
//   		service: jsii.String("service"),
//   	},
//   }
//
type CfnDistributionConfiguration_ContainerDistributionConfigurationProperty struct {
	// Tags that are attached to the container distribution configuration.
	ContainerTags *[]*string `field:"optional" json:"containerTags" yaml:"containerTags"`
	// The description of the container distribution configuration.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The destination repository for the container distribution configuration.
	TargetRepository interface{} `field:"optional" json:"targetRepository" yaml:"targetRepository"`
}

// The distribution configuration distribution defines the settings for a specific Region in the Distribution Configuration.
//
// You must specify whether the distribution is for an AMI or a container image. To do so, include exactly one of the following data types for your distribution:
//
// - amiDistributionConfiguration
// - containerDistributionConfiguration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var amiDistributionConfiguration interface{}
//   var containerDistributionConfiguration interface{}
//
//   distributionProperty := &distributionProperty{
//   	region: jsii.String("region"),
//
//   	// the properties below are optional
//   	amiDistributionConfiguration: amiDistributionConfiguration,
//   	containerDistributionConfiguration: containerDistributionConfiguration,
//   	fastLaunchConfigurations: []interface{}{
//   		&fastLaunchConfigurationProperty{
//   			accountId: jsii.String("accountId"),
//   			enabled: jsii.Boolean(false),
//   			launchTemplate: &fastLaunchLaunchTemplateSpecificationProperty{
//   				launchTemplateId: jsii.String("launchTemplateId"),
//   				launchTemplateName: jsii.String("launchTemplateName"),
//   				launchTemplateVersion: jsii.String("launchTemplateVersion"),
//   			},
//   			maxParallelLaunches: jsii.Number(123),
//   			snapshotConfiguration: &fastLaunchSnapshotConfigurationProperty{
//   				targetResourceCount: jsii.Number(123),
//   			},
//   		},
//   	},
//   	launchTemplateConfigurations: []interface{}{
//   		&launchTemplateConfigurationProperty{
//   			accountId: jsii.String("accountId"),
//   			launchTemplateId: jsii.String("launchTemplateId"),
//   			setDefaultVersion: jsii.Boolean(false),
//   		},
//   	},
//   	licenseConfigurationArns: []*string{
//   		jsii.String("licenseConfigurationArns"),
//   	},
//   }
//
type CfnDistributionConfiguration_DistributionProperty struct {
	// The target Region for the Distribution Configuration.
	//
	// For example, `eu-west-1` .
	Region *string `field:"required" json:"region" yaml:"region"`
	// The specific AMI settings, such as launch permissions and AMI tags.
	//
	// For details, see example schema below.
	AmiDistributionConfiguration interface{} `field:"optional" json:"amiDistributionConfiguration" yaml:"amiDistributionConfiguration"`
	// Container distribution settings for encryption, licensing, and sharing in a specific Region.
	//
	// For details, see example schema below.
	ContainerDistributionConfiguration interface{} `field:"optional" json:"containerDistributionConfiguration" yaml:"containerDistributionConfiguration"`
	// `CfnDistributionConfiguration.DistributionProperty.FastLaunchConfigurations`.
	FastLaunchConfigurations interface{} `field:"optional" json:"fastLaunchConfigurations" yaml:"fastLaunchConfigurations"`
	// A group of launchTemplateConfiguration settings that apply to image distribution for specified accounts.
	LaunchTemplateConfigurations interface{} `field:"optional" json:"launchTemplateConfigurations" yaml:"launchTemplateConfigurations"`
	// The License Manager Configuration to associate with the AMI in the specified Region.
	//
	// For more information, see the [LicenseConfiguration API](https://docs.aws.amazon.com/license-manager/latest/APIReference/API_LicenseConfiguration.html) .
	LicenseConfigurationArns *[]*string `field:"optional" json:"licenseConfigurationArns" yaml:"licenseConfigurationArns"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fastLaunchConfigurationProperty := &fastLaunchConfigurationProperty{
//   	accountId: jsii.String("accountId"),
//   	enabled: jsii.Boolean(false),
//   	launchTemplate: &fastLaunchLaunchTemplateSpecificationProperty{
//   		launchTemplateId: jsii.String("launchTemplateId"),
//   		launchTemplateName: jsii.String("launchTemplateName"),
//   		launchTemplateVersion: jsii.String("launchTemplateVersion"),
//   	},
//   	maxParallelLaunches: jsii.Number(123),
//   	snapshotConfiguration: &fastLaunchSnapshotConfigurationProperty{
//   		targetResourceCount: jsii.Number(123),
//   	},
//   }
//
type CfnDistributionConfiguration_FastLaunchConfigurationProperty struct {
	// `CfnDistributionConfiguration.FastLaunchConfigurationProperty.AccountId`.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// `CfnDistributionConfiguration.FastLaunchConfigurationProperty.Enabled`.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// `CfnDistributionConfiguration.FastLaunchConfigurationProperty.LaunchTemplate`.
	LaunchTemplate interface{} `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// `CfnDistributionConfiguration.FastLaunchConfigurationProperty.MaxParallelLaunches`.
	MaxParallelLaunches *float64 `field:"optional" json:"maxParallelLaunches" yaml:"maxParallelLaunches"`
	// `CfnDistributionConfiguration.FastLaunchConfigurationProperty.SnapshotConfiguration`.
	SnapshotConfiguration interface{} `field:"optional" json:"snapshotConfiguration" yaml:"snapshotConfiguration"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fastLaunchLaunchTemplateSpecificationProperty := &fastLaunchLaunchTemplateSpecificationProperty{
//   	launchTemplateId: jsii.String("launchTemplateId"),
//   	launchTemplateName: jsii.String("launchTemplateName"),
//   	launchTemplateVersion: jsii.String("launchTemplateVersion"),
//   }
//
type CfnDistributionConfiguration_FastLaunchLaunchTemplateSpecificationProperty struct {
	// `CfnDistributionConfiguration.FastLaunchLaunchTemplateSpecificationProperty.LaunchTemplateId`.
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// `CfnDistributionConfiguration.FastLaunchLaunchTemplateSpecificationProperty.LaunchTemplateName`.
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
	// `CfnDistributionConfiguration.FastLaunchLaunchTemplateSpecificationProperty.LaunchTemplateVersion`.
	LaunchTemplateVersion *string `field:"optional" json:"launchTemplateVersion" yaml:"launchTemplateVersion"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fastLaunchSnapshotConfigurationProperty := &fastLaunchSnapshotConfigurationProperty{
//   	targetResourceCount: jsii.Number(123),
//   }
//
type CfnDistributionConfiguration_FastLaunchSnapshotConfigurationProperty struct {
	// `CfnDistributionConfiguration.FastLaunchSnapshotConfigurationProperty.TargetResourceCount`.
	TargetResourceCount *float64 `field:"optional" json:"targetResourceCount" yaml:"targetResourceCount"`
}

// Describes the configuration for a launch permission.
//
// The launch permission modification request is sent to the [Amazon EC2 ModifyImageAttribute](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyImageAttribute.html) API on behalf of the user for each Region they have selected to distribute the AMI. To make an AMI public, set the launch permission authorized accounts to `all` . See the examples for making an AMI public at [Amazon EC2 ModifyImageAttribute](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyImageAttribute.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchPermissionConfigurationProperty := &launchPermissionConfigurationProperty{
//   	organizationalUnitArns: []*string{
//   		jsii.String("organizationalUnitArns"),
//   	},
//   	organizationArns: []*string{
//   		jsii.String("organizationArns"),
//   	},
//   	userGroups: []*string{
//   		jsii.String("userGroups"),
//   	},
//   	userIds: []*string{
//   		jsii.String("userIds"),
//   	},
//   }
//
type CfnDistributionConfiguration_LaunchPermissionConfigurationProperty struct {
	// The ARN for an AWS Organizations organizational unit (OU) that you want to share your AMI with.
	//
	// For more information about key concepts for AWS Organizations , see [AWS Organizations terminology and concepts](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html) .
	OrganizationalUnitArns *[]*string `field:"optional" json:"organizationalUnitArns" yaml:"organizationalUnitArns"`
	// The ARN for an AWS Organization that you want to share your AMI with.
	//
	// For more information, see [What is AWS Organizations ?](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_introduction.html) .
	OrganizationArns *[]*string `field:"optional" json:"organizationArns" yaml:"organizationArns"`
	// The name of the group.
	UserGroups *[]*string `field:"optional" json:"userGroups" yaml:"userGroups"`
	// The AWS account ID.
	UserIds *[]*string `field:"optional" json:"userIds" yaml:"userIds"`
}

// Identifies an Amazon EC2 launch template to use for a specific account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateConfigurationProperty := &launchTemplateConfigurationProperty{
//   	accountId: jsii.String("accountId"),
//   	launchTemplateId: jsii.String("launchTemplateId"),
//   	setDefaultVersion: jsii.Boolean(false),
//   }
//
type CfnDistributionConfiguration_LaunchTemplateConfigurationProperty struct {
	// The account ID that this configuration applies to.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Identifies the Amazon EC2 launch template to use.
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// Set the specified Amazon EC2 launch template as the default launch template for the specified account.
	SetDefaultVersion interface{} `field:"optional" json:"setDefaultVersion" yaml:"setDefaultVersion"`
}

// The container repository where the output container image is stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetContainerRepositoryProperty := &targetContainerRepositoryProperty{
//   	repositoryName: jsii.String("repositoryName"),
//   	service: jsii.String("service"),
//   }
//
type CfnDistributionConfiguration_TargetContainerRepositoryProperty struct {
	// The name of the container repository where the output container image is stored.
	//
	// This name is prefixed by the repository location.
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// Specifies the service in which this image was registered.
	Service *string `field:"optional" json:"service" yaml:"service"`
}

// Properties for defining a `CfnDistributionConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var amiDistributionConfiguration interface{}
//   var containerDistributionConfiguration interface{}
//
//   cfnDistributionConfigurationProps := &cfnDistributionConfigurationProps{
//   	distributions: []interface{}{
//   		&distributionProperty{
//   			region: jsii.String("region"),
//
//   			// the properties below are optional
//   			amiDistributionConfiguration: amiDistributionConfiguration,
//   			containerDistributionConfiguration: containerDistributionConfiguration,
//   			fastLaunchConfigurations: []interface{}{
//   				&fastLaunchConfigurationProperty{
//   					accountId: jsii.String("accountId"),
//   					enabled: jsii.Boolean(false),
//   					launchTemplate: &fastLaunchLaunchTemplateSpecificationProperty{
//   						launchTemplateId: jsii.String("launchTemplateId"),
//   						launchTemplateName: jsii.String("launchTemplateName"),
//   						launchTemplateVersion: jsii.String("launchTemplateVersion"),
//   					},
//   					maxParallelLaunches: jsii.Number(123),
//   					snapshotConfiguration: &fastLaunchSnapshotConfigurationProperty{
//   						targetResourceCount: jsii.Number(123),
//   					},
//   				},
//   			},
//   			launchTemplateConfigurations: []interface{}{
//   				&launchTemplateConfigurationProperty{
//   					accountId: jsii.String("accountId"),
//   					launchTemplateId: jsii.String("launchTemplateId"),
//   					setDefaultVersion: jsii.Boolean(false),
//   				},
//   			},
//   			licenseConfigurationArns: []*string{
//   				jsii.String("licenseConfigurationArns"),
//   			},
//   		},
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnDistributionConfigurationProps struct {
	// The distributions of this distribution configuration formatted as an array of Distribution objects.
	Distributions interface{} `field:"required" json:"distributions" yaml:"distributions"`
	// The name of this distribution configuration.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of this distribution configuration.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags of this distribution configuration.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::ImageBuilder::Image`.
//
// An image build version. An image is a customized, secure, and up-to-date “golden” server image that is pre-installed and pre-configured with software and settings to meet specific IT standards.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnImage := awscdk.Aws_imagebuilder.NewCfnImage(this, jsii.String("MyCfnImage"), &cfnImageProps{
//   	infrastructureConfigurationArn: jsii.String("infrastructureConfigurationArn"),
//
//   	// the properties below are optional
//   	containerRecipeArn: jsii.String("containerRecipeArn"),
//   	distributionConfigurationArn: jsii.String("distributionConfigurationArn"),
//   	enhancedImageMetadataEnabled: jsii.Boolean(false),
//   	imageRecipeArn: jsii.String("imageRecipeArn"),
//   	imageTestsConfiguration: &imageTestsConfigurationProperty{
//   		imageTestsEnabled: jsii.Boolean(false),
//   		timeoutMinutes: jsii.Number(123),
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
type CfnImage interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Returns the Amazon Resource Name (ARN) of the image.
	//
	// For example, `arn:aws:imagebuilder:us-west-2:123456789012:image/mybasicrecipe/2019.12.03/1` .
	AttrArn() *string
	// Returns the AMI ID of the Amazon EC2 AMI in the Region in which you are using Image Builder.
	//
	// Values are returned only for AMIs, and not for container images.
	AttrImageId() *string
	// Returns a list of URIs for container images created in the context Region.
	//
	// Values are returned only for container images, and not for AMIs.
	AttrImageUri() *string
	// Returns the name of the image.
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
	// The Amazon Resource Name (ARN) of the distribution configuration.
	DistributionConfigurationArn() *string
	SetDistributionConfigurationArn(val *string)
	// Collects additional information about the image being created, including the operating system (OS) version and package list.
	//
	// This information is used to enhance the overall experience of using EC2 Image Builder. Enabled by default.
	EnhancedImageMetadataEnabled() interface{}
	SetEnhancedImageMetadataEnabled(val interface{})
	// The Amazon Resource Name (ARN) of the image recipe.
	ImageRecipeArn() *string
	SetImageRecipeArn(val *string)
	// The configuration settings for your image test components, which includes a toggle that allows you to turn off tests, and a timeout setting.
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
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags of the image.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
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
	GetAtt(attributeName *string) awscdk.Reference
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

// The jsii proxy struct for CfnImage
type jsiiProxy_CfnImage struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnImage) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) AttrImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) AttrImageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrImageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) ContainerRecipeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRecipeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) DistributionConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) EnhancedImageMetadataEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedImageMetadataEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) ImageRecipeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageRecipeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) ImageTestsConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageTestsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) InfrastructureConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImage) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ImageBuilder::Image`.
func NewCfnImage(scope constructs.Construct, id *string, props *CfnImageProps) CfnImage {
	_init_.Initialize()

	j := jsiiProxy_CfnImage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_imagebuilder.CfnImage",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ImageBuilder::Image`.
func NewCfnImage_Override(c CfnImage, scope constructs.Construct, id *string, props *CfnImageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_imagebuilder.CfnImage",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnImage) SetContainerRecipeArn(val *string) {
	_jsii_.Set(
		j,
		"containerRecipeArn",
		val,
	)
}

func (j *jsiiProxy_CfnImage) SetDistributionConfigurationArn(val *string) {
	_jsii_.Set(
		j,
		"distributionConfigurationArn",
		val,
	)
}

func (j *jsiiProxy_CfnImage) SetEnhancedImageMetadataEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enhancedImageMetadataEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnImage) SetImageRecipeArn(val *string) {
	_jsii_.Set(
		j,
		"imageRecipeArn",
		val,
	)
}

func (j *jsiiProxy_CfnImage) SetImageTestsConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"imageTestsConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnImage) SetInfrastructureConfigurationArn(val *string) {
	_jsii_.Set(
		j,
		"infrastructureConfigurationArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnImage_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnImage",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnImage_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnImage",
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
func CfnImage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnImage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnImage_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_imagebuilder.CfnImage",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnImage) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnImage) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnImage) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnImage) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnImage) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnImage) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnImage) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnImage) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImage) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImage) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnImage) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnImage) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImage) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImage) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// When you create an image or container recipe with Image Builder , you can add the build or test components that are used to create the final image.
//
// You must have at least one build component to create a recipe, but test components are not required. If you have added tests, they run after the image is created, to ensure that the target image is functional and can be used reliably for launching Amazon EC2 instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageTestsConfigurationProperty := &imageTestsConfigurationProperty{
//   	imageTestsEnabled: jsii.Boolean(false),
//   	timeoutMinutes: jsii.Number(123),
//   }
//
type CfnImage_ImageTestsConfigurationProperty struct {
	// Determines if tests should run after building the image.
	//
	// Image Builder defaults to enable tests to run following the image build, before image distribution.
	ImageTestsEnabled interface{} `field:"optional" json:"imageTestsEnabled" yaml:"imageTestsEnabled"`
	// The maximum time in minutes that tests are permitted to run.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
}

// A CloudFormation `AWS::ImageBuilder::ImagePipeline`.
//
// An image pipeline is the automation configuration for building secure OS images on AWS . The Image Builder image pipeline is associated with an image recipe that defines the build, validation, and test phases for an image build lifecycle. An image pipeline can be associated with an infrastructure configuration that defines where your image is built. You can define attributes, such as instance type, subnets, security groups, logging, and other infrastructure-related configurations. You can also associate your image pipeline with a distribution configuration to define how you would like to deploy your image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnImagePipeline := awscdk.Aws_imagebuilder.NewCfnImagePipeline(this, jsii.String("MyCfnImagePipeline"), &cfnImagePipelineProps{
//   	infrastructureConfigurationArn: jsii.String("infrastructureConfigurationArn"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	containerRecipeArn: jsii.String("containerRecipeArn"),
//   	description: jsii.String("description"),
//   	distributionConfigurationArn: jsii.String("distributionConfigurationArn"),
//   	enhancedImageMetadataEnabled: jsii.Boolean(false),
//   	imageRecipeArn: jsii.String("imageRecipeArn"),
//   	imageTestsConfiguration: &imageTestsConfigurationProperty{
//   		imageTestsEnabled: jsii.Boolean(false),
//   		timeoutMinutes: jsii.Number(123),
//   	},
//   	schedule: &scheduleProperty{
//   		pipelineExecutionStartCondition: jsii.String("pipelineExecutionStartCondition"),
//   		scheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   	status: jsii.String("status"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
type CfnImagePipeline interface {
	awscdk.CfnResource
	awscdk.IInspectable
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
	//
	// This information is used to enhance the overall experience of using EC2 Image Builder. Enabled by default.
	EnhancedImageMetadataEnabled() interface{}
	SetEnhancedImageMetadataEnabled(val interface{})
	// The Amazon Resource Name (ARN) of the image recipe associated with this image pipeline.
	ImageRecipeArn() *string
	SetImageRecipeArn(val *string)
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
	//
	// A schedule configures how often and when a pipeline automatically creates a new image.
	Schedule() interface{}
	SetSchedule(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The status of the image pipeline.
	Status() *string
	SetStatus(val *string)
	// The tags of this image pipeline.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
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
	GetAtt(attributeName *string) awscdk.Reference
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

// The jsii proxy struct for CfnImagePipeline
type jsiiProxy_CfnImagePipeline struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
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

func (j *jsiiProxy_CfnImagePipeline) ImageRecipeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageRecipeArn",
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

func (j *jsiiProxy_CfnImagePipeline) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ImageBuilder::ImagePipeline`.
func NewCfnImagePipeline(scope constructs.Construct, id *string, props *CfnImagePipelineProps) CfnImagePipeline {
	_init_.Initialize()

	j := jsiiProxy_CfnImagePipeline{}

	_jsii_.Create(
		"aws-cdk-lib.aws_imagebuilder.CfnImagePipeline",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ImageBuilder::ImagePipeline`.
func NewCfnImagePipeline_Override(c CfnImagePipeline, scope constructs.Construct, id *string, props *CfnImagePipelineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_imagebuilder.CfnImagePipeline",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnImagePipeline) SetContainerRecipeArn(val *string) {
	_jsii_.Set(
		j,
		"containerRecipeArn",
		val,
	)
}

func (j *jsiiProxy_CfnImagePipeline) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnImagePipeline) SetDistributionConfigurationArn(val *string) {
	_jsii_.Set(
		j,
		"distributionConfigurationArn",
		val,
	)
}

func (j *jsiiProxy_CfnImagePipeline) SetEnhancedImageMetadataEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enhancedImageMetadataEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnImagePipeline) SetImageRecipeArn(val *string) {
	_jsii_.Set(
		j,
		"imageRecipeArn",
		val,
	)
}

func (j *jsiiProxy_CfnImagePipeline) SetImageTestsConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"imageTestsConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnImagePipeline) SetInfrastructureConfigurationArn(val *string) {
	_jsii_.Set(
		j,
		"infrastructureConfigurationArn",
		val,
	)
}

func (j *jsiiProxy_CfnImagePipeline) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnImagePipeline) SetSchedule(val interface{}) {
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_CfnImagePipeline) SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
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

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnImagePipeline",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnImagePipeline_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnImagePipeline",
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
func CfnImagePipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

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
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnImagePipeline) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnImagePipeline) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnImagePipeline) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnImagePipeline) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnImagePipeline) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnImagePipeline) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnImagePipeline) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImagePipeline) GetMetadata(key *string) interface{} {
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
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnImagePipeline) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnImagePipeline) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
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
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// When you create an image or container recipe with Image Builder , you can add the build or test components that your image pipeline uses to create the final image.
//
// You must have at least one build component to create a recipe, but test components are not required. Your pipeline runs tests after it builds the image, to ensure that the target image is functional and can be used reliably for launching Amazon EC2 instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageTestsConfigurationProperty := &imageTestsConfigurationProperty{
//   	imageTestsEnabled: jsii.Boolean(false),
//   	timeoutMinutes: jsii.Number(123),
//   }
//
type CfnImagePipeline_ImageTestsConfigurationProperty struct {
	// Defines if tests should be executed when building this image.
	//
	// For example, `true` or `false` .
	ImageTestsEnabled interface{} `field:"optional" json:"imageTestsEnabled" yaml:"imageTestsEnabled"`
	// The maximum time in minutes that tests are permitted to run.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
}

// A schedule configures how often and when a pipeline will automatically create a new image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleProperty := &scheduleProperty{
//   	pipelineExecutionStartCondition: jsii.String("pipelineExecutionStartCondition"),
//   	scheduleExpression: jsii.String("scheduleExpression"),
//   }
//
type CfnImagePipeline_ScheduleProperty struct {
	// The condition configures when the pipeline should trigger a new image build.
	//
	// When the `pipelineExecutionStartCondition` is set to `EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE` , and you use semantic version filters on the base image or components in your image recipe, Image Builder will build a new image only when there are new versions of the image or components in your recipe that match the semantic version filter. When it is set to `EXPRESSION_MATCH_ONLY` , it will build a new image every time the CRON expression matches the current time. For semantic version syntax, see [CreateComponent](https://docs.aws.amazon.com/imagebuilder/latest/APIReference/API_CreateComponent.html) in the *Image Builder API Reference* .
	PipelineExecutionStartCondition *string `field:"optional" json:"pipelineExecutionStartCondition" yaml:"pipelineExecutionStartCondition"`
	// The cron expression determines how often EC2 Image Builder evaluates your `pipelineExecutionStartCondition` .
	//
	// For information on how to format a cron expression in Image Builder, see [Use cron expressions in EC2 Image Builder](https://docs.aws.amazon.com/imagebuilder/latest/userguide/image-builder-cron.html) .
	ScheduleExpression *string `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
}

// Properties for defining a `CfnImagePipeline`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnImagePipelineProps := &cfnImagePipelineProps{
//   	infrastructureConfigurationArn: jsii.String("infrastructureConfigurationArn"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	containerRecipeArn: jsii.String("containerRecipeArn"),
//   	description: jsii.String("description"),
//   	distributionConfigurationArn: jsii.String("distributionConfigurationArn"),
//   	enhancedImageMetadataEnabled: jsii.Boolean(false),
//   	imageRecipeArn: jsii.String("imageRecipeArn"),
//   	imageTestsConfiguration: &imageTestsConfigurationProperty{
//   		imageTestsEnabled: jsii.Boolean(false),
//   		timeoutMinutes: jsii.Number(123),
//   	},
//   	schedule: &scheduleProperty{
//   		pipelineExecutionStartCondition: jsii.String("pipelineExecutionStartCondition"),
//   		scheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   	status: jsii.String("status"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnImagePipelineProps struct {
	// The Amazon Resource Name (ARN) of the infrastructure configuration associated with this image pipeline.
	InfrastructureConfigurationArn *string `field:"required" json:"infrastructureConfigurationArn" yaml:"infrastructureConfigurationArn"`
	// The name of the image pipeline.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the container recipe that is used for this pipeline.
	ContainerRecipeArn *string `field:"optional" json:"containerRecipeArn" yaml:"containerRecipeArn"`
	// The description of this image pipeline.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of the distribution configuration associated with this image pipeline.
	DistributionConfigurationArn *string `field:"optional" json:"distributionConfigurationArn" yaml:"distributionConfigurationArn"`
	// Collects additional information about the image being created, including the operating system (OS) version and package list.
	//
	// This information is used to enhance the overall experience of using EC2 Image Builder. Enabled by default.
	EnhancedImageMetadataEnabled interface{} `field:"optional" json:"enhancedImageMetadataEnabled" yaml:"enhancedImageMetadataEnabled"`
	// The Amazon Resource Name (ARN) of the image recipe associated with this image pipeline.
	ImageRecipeArn *string `field:"optional" json:"imageRecipeArn" yaml:"imageRecipeArn"`
	// The configuration of the image tests that run after image creation to ensure the quality of the image that was created.
	ImageTestsConfiguration interface{} `field:"optional" json:"imageTestsConfiguration" yaml:"imageTestsConfiguration"`
	// The schedule of the image pipeline.
	//
	// A schedule configures how often and when a pipeline automatically creates a new image.
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// The status of the image pipeline.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The tags of this image pipeline.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

// Properties for defining a `CfnImage`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnImageProps := &cfnImageProps{
//   	infrastructureConfigurationArn: jsii.String("infrastructureConfigurationArn"),
//
//   	// the properties below are optional
//   	containerRecipeArn: jsii.String("containerRecipeArn"),
//   	distributionConfigurationArn: jsii.String("distributionConfigurationArn"),
//   	enhancedImageMetadataEnabled: jsii.Boolean(false),
//   	imageRecipeArn: jsii.String("imageRecipeArn"),
//   	imageTestsConfiguration: &imageTestsConfigurationProperty{
//   		imageTestsEnabled: jsii.Boolean(false),
//   		timeoutMinutes: jsii.Number(123),
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnImageProps struct {
	// The Amazon Resource Name (ARN) of the infrastructure configuration associated with this image pipeline.
	InfrastructureConfigurationArn *string `field:"required" json:"infrastructureConfigurationArn" yaml:"infrastructureConfigurationArn"`
	// The Amazon Resource Name (ARN) of the container recipe that is used for this pipeline.
	ContainerRecipeArn *string `field:"optional" json:"containerRecipeArn" yaml:"containerRecipeArn"`
	// The Amazon Resource Name (ARN) of the distribution configuration.
	DistributionConfigurationArn *string `field:"optional" json:"distributionConfigurationArn" yaml:"distributionConfigurationArn"`
	// Collects additional information about the image being created, including the operating system (OS) version and package list.
	//
	// This information is used to enhance the overall experience of using EC2 Image Builder. Enabled by default.
	EnhancedImageMetadataEnabled interface{} `field:"optional" json:"enhancedImageMetadataEnabled" yaml:"enhancedImageMetadataEnabled"`
	// The Amazon Resource Name (ARN) of the image recipe.
	ImageRecipeArn *string `field:"optional" json:"imageRecipeArn" yaml:"imageRecipeArn"`
	// The configuration settings for your image test components, which includes a toggle that allows you to turn off tests, and a timeout setting.
	ImageTestsConfiguration interface{} `field:"optional" json:"imageTestsConfiguration" yaml:"imageTestsConfiguration"`
	// The tags of the image.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::ImageBuilder::ImageRecipe`.
//
// An Image Builder image recipe is a document that defines the base image and the components to be applied to the base image to produce the desired configuration for the output image. You can use an image recipe to duplicate builds. Image Builder image recipes can be shared, branched, and edited using the console wizard, the AWS CLI , or the API. You can use image recipes with your version control software to maintain shareable versioned image recipes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnImageRecipe := awscdk.Aws_imagebuilder.NewCfnImageRecipe(this, jsii.String("MyCfnImageRecipe"), &cfnImageRecipeProps{
//   	components: []interface{}{
//   		&componentConfigurationProperty{
//   			componentArn: jsii.String("componentArn"),
//   			parameters: []interface{}{
//   				&componentParameterProperty{
//   					name: jsii.String("name"),
//   					value: []*string{
//   						jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	name: jsii.String("name"),
//   	parentImage: jsii.String("parentImage"),
//   	version: jsii.String("version"),
//
//   	// the properties below are optional
//   	additionalInstanceConfiguration: &additionalInstanceConfigurationProperty{
//   		systemsManagerAgent: &systemsManagerAgentProperty{
//   			uninstallAfterBuild: jsii.Boolean(false),
//   		},
//   		userDataOverride: jsii.String("userDataOverride"),
//   	},
//   	blockDeviceMappings: []interface{}{
//   		&instanceBlockDeviceMappingProperty{
//   			deviceName: jsii.String("deviceName"),
//   			ebs: &ebsInstanceBlockDeviceSpecificationProperty{
//   				deleteOnTermination: jsii.Boolean(false),
//   				encrypted: jsii.Boolean(false),
//   				iops: jsii.Number(123),
//   				kmsKeyId: jsii.String("kmsKeyId"),
//   				snapshotId: jsii.String("snapshotId"),
//   				throughput: jsii.Number(123),
//   				volumeSize: jsii.Number(123),
//   				volumeType: jsii.String("volumeType"),
//   			},
//   			noDevice: jsii.String("noDevice"),
//   			virtualName: jsii.String("virtualName"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	workingDirectory: jsii.String("workingDirectory"),
//   })
//
type CfnImageRecipe interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Before you create a new AMI, Image Builder launches temporary Amazon EC2 instances to build and test your image configuration.
	//
	// Instance configuration adds a layer of control over those instances. You can define settings and add scripts to run when an instance is launched from your AMI.
	AdditionalInstanceConfiguration() interface{}
	SetAdditionalInstanceConfiguration(val interface{})
	// Returns the Amazon Resource Name (ARN) of the image recipe.
	//
	// For example, `arn:aws:imagebuilder:us-east-1:123456789012:image-recipe/mybasicrecipe/2019.12.03` .
	AttrArn() *string
	// The name of the image recipe.
	AttrName() *string
	// The block device mappings to apply when creating images from this recipe.
	BlockDeviceMappings() interface{}
	SetBlockDeviceMappings(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The components of the image recipe.
	//
	// Components are orchestration documents that define a sequence of steps for downloading, installing, configuring, and testing software packages. They also define validation and security hardening steps. A component is defined using a YAML document format.
	Components() interface{}
	SetComponents(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the image recipe.
	Description() *string
	SetDescription(val *string)
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
	// The name of the image recipe.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The parent image of the image recipe.
	//
	// The string must be either an Image ARN or an AMI ID.
	ParentImage() *string
	SetParentImage(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags of the image recipe.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// The semantic version of the image recipe.
	Version() *string
	SetVersion(val *string)
	// The working directory to be used during build and test workflows.
	WorkingDirectory() *string
	SetWorkingDirectory(val *string)
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
	GetAtt(attributeName *string) awscdk.Reference
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

// The jsii proxy struct for CfnImageRecipe
type jsiiProxy_CfnImageRecipe struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnImageRecipe) AdditionalInstanceConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalInstanceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageRecipe) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageRecipe) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageRecipe) BlockDeviceMappings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockDeviceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageRecipe) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageRecipe) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageRecipe) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageRecipe) Components() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageRecipe) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageRecipe) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageRecipe) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageRecipe) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageRecipe) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageRecipe) ParentImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageRecipe) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageRecipe) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageRecipe) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageRecipe) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageRecipe) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImageRecipe) WorkingDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirectory",
		&returns,
	)
	return returns
}


// Create a new `AWS::ImageBuilder::ImageRecipe`.
func NewCfnImageRecipe(scope constructs.Construct, id *string, props *CfnImageRecipeProps) CfnImageRecipe {
	_init_.Initialize()

	j := jsiiProxy_CfnImageRecipe{}

	_jsii_.Create(
		"aws-cdk-lib.aws_imagebuilder.CfnImageRecipe",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ImageBuilder::ImageRecipe`.
func NewCfnImageRecipe_Override(c CfnImageRecipe, scope constructs.Construct, id *string, props *CfnImageRecipeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_imagebuilder.CfnImageRecipe",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnImageRecipe) SetAdditionalInstanceConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"additionalInstanceConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnImageRecipe) SetBlockDeviceMappings(val interface{}) {
	_jsii_.Set(
		j,
		"blockDeviceMappings",
		val,
	)
}

func (j *jsiiProxy_CfnImageRecipe) SetComponents(val interface{}) {
	_jsii_.Set(
		j,
		"components",
		val,
	)
}

func (j *jsiiProxy_CfnImageRecipe) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnImageRecipe) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnImageRecipe) SetParentImage(val *string) {
	_jsii_.Set(
		j,
		"parentImage",
		val,
	)
}

func (j *jsiiProxy_CfnImageRecipe) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_CfnImageRecipe) SetWorkingDirectory(val *string) {
	_jsii_.Set(
		j,
		"workingDirectory",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnImageRecipe_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnImageRecipe",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnImageRecipe_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnImageRecipe",
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
func CfnImageRecipe_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnImageRecipe",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnImageRecipe_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_imagebuilder.CfnImageRecipe",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnImageRecipe) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnImageRecipe) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnImageRecipe) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnImageRecipe) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnImageRecipe) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnImageRecipe) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnImageRecipe) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnImageRecipe) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImageRecipe) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImageRecipe) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnImageRecipe) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnImageRecipe) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImageRecipe) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImageRecipe) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnImageRecipe) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// In addition to your infrastructure configuration, these settings provide an extra layer of control over your build instances.
//
// You can also specify commands to run on launch for all of your build instances.
//
// Image Builder does not automatically install the Systems Manager agent on Windows instances. If your base image includes the Systems Manager agent, then the AMI that you create will also include the agent. For Linux instances, if the base image does not already include the Systems Manager agent, Image Builder installs it. For Linux instances where Image Builder installs the Systems Manager agent, you can choose whether to keep it for the AMI that you create.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   additionalInstanceConfigurationProperty := &additionalInstanceConfigurationProperty{
//   	systemsManagerAgent: &systemsManagerAgentProperty{
//   		uninstallAfterBuild: jsii.Boolean(false),
//   	},
//   	userDataOverride: jsii.String("userDataOverride"),
//   }
//
type CfnImageRecipe_AdditionalInstanceConfigurationProperty struct {
	// Contains settings for the Systems Manager agent on your build instance.
	SystemsManagerAgent interface{} `field:"optional" json:"systemsManagerAgent" yaml:"systemsManagerAgent"`
	// Use this property to provide commands or a command script to run when you launch your build instance.
	//
	// The userDataOverride property replaces any commands that Image Builder might have added to ensure that Systems Manager is installed on your Linux build instance. If you override the user data, make sure that you add commands to install Systems Manager, if it is not pre-installed on your base image.
	//
	// > The user data is always base 64 encoded. For example, the following commands are encoded as `IyEvYmluL2Jhc2gKbWtkaXIgLXAgL3Zhci9iYi8KdG91Y2ggL3Zhci$` :
	// >
	// > *#!/bin/bash*
	// >
	// > mkdir -p /var/bb/
	// >
	// > touch /var.
	UserDataOverride *string `field:"optional" json:"userDataOverride" yaml:"userDataOverride"`
}

// Configuration details of the component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   componentConfigurationProperty := &componentConfigurationProperty{
//   	componentArn: jsii.String("componentArn"),
//   	parameters: []interface{}{
//   		&componentParameterProperty{
//   			name: jsii.String("name"),
//   			value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnImageRecipe_ComponentConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the component.
	ComponentArn *string `field:"optional" json:"componentArn" yaml:"componentArn"`
	// A group of parameter settings that are used to configure the component for a specific recipe.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

// Contains a key/value pair that sets the named component parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   componentParameterProperty := &componentParameterProperty{
//   	name: jsii.String("name"),
//   	value: []*string{
//   		jsii.String("value"),
//   	},
//   }
//
type CfnImageRecipe_ComponentParameterProperty struct {
	// The name of the component parameter to set.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Sets the value for the named component parameter.
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

// The image recipe EBS instance block device specification includes the Amazon EBS-specific block device mapping specifications for the image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ebsInstanceBlockDeviceSpecificationProperty := &ebsInstanceBlockDeviceSpecificationProperty{
//   	deleteOnTermination: jsii.Boolean(false),
//   	encrypted: jsii.Boolean(false),
//   	iops: jsii.Number(123),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	snapshotId: jsii.String("snapshotId"),
//   	throughput: jsii.Number(123),
//   	volumeSize: jsii.Number(123),
//   	volumeType: jsii.String("volumeType"),
//   }
//
type CfnImageRecipe_EbsInstanceBlockDeviceSpecificationProperty struct {
	// Configures delete on termination of the associated device.
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Use to configure device encryption.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// Use to configure device IOPS.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Use to configure the KMS key to use when encrypting the device.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The snapshot that defines the device contents.
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// *For GP3 volumes only* – The throughput in MiB/s that the volume supports.
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// Overrides the volume size of the device.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// Overrides the volume type of the device.
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

// Defines block device mappings for the instance used to configure your image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceBlockDeviceMappingProperty := &instanceBlockDeviceMappingProperty{
//   	deviceName: jsii.String("deviceName"),
//   	ebs: &ebsInstanceBlockDeviceSpecificationProperty{
//   		deleteOnTermination: jsii.Boolean(false),
//   		encrypted: jsii.Boolean(false),
//   		iops: jsii.Number(123),
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   		snapshotId: jsii.String("snapshotId"),
//   		throughput: jsii.Number(123),
//   		volumeSize: jsii.Number(123),
//   		volumeType: jsii.String("volumeType"),
//   	},
//   	noDevice: jsii.String("noDevice"),
//   	virtualName: jsii.String("virtualName"),
//   }
//
type CfnImageRecipe_InstanceBlockDeviceMappingProperty struct {
	// The device to which these mappings apply.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// Use to manage Amazon EBS-specific configuration for this mapping.
	Ebs interface{} `field:"optional" json:"ebs" yaml:"ebs"`
	// Enter an empty string to remove a mapping from the parent image.
	//
	// The following is an example of an empty string value in the `NoDevice` field.
	//
	// `NoDevice:""`.
	NoDevice *string `field:"optional" json:"noDevice" yaml:"noDevice"`
	// Manages the instance ephemeral devices.
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}

// Contains settings for the Systems Manager agent on your build instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   systemsManagerAgentProperty := &systemsManagerAgentProperty{
//   	uninstallAfterBuild: jsii.Boolean(false),
//   }
//
type CfnImageRecipe_SystemsManagerAgentProperty struct {
	// Controls whether the Systems Manager agent is removed from your final build image, prior to creating the new AMI.
	//
	// If this is set to true, then the agent is removed from the final image. If it's set to false, then the agent is left in, so that it is included in the new AMI. The default value is false.
	UninstallAfterBuild interface{} `field:"optional" json:"uninstallAfterBuild" yaml:"uninstallAfterBuild"`
}

// Properties for defining a `CfnImageRecipe`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnImageRecipeProps := &cfnImageRecipeProps{
//   	components: []interface{}{
//   		&componentConfigurationProperty{
//   			componentArn: jsii.String("componentArn"),
//   			parameters: []interface{}{
//   				&componentParameterProperty{
//   					name: jsii.String("name"),
//   					value: []*string{
//   						jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	name: jsii.String("name"),
//   	parentImage: jsii.String("parentImage"),
//   	version: jsii.String("version"),
//
//   	// the properties below are optional
//   	additionalInstanceConfiguration: &additionalInstanceConfigurationProperty{
//   		systemsManagerAgent: &systemsManagerAgentProperty{
//   			uninstallAfterBuild: jsii.Boolean(false),
//   		},
//   		userDataOverride: jsii.String("userDataOverride"),
//   	},
//   	blockDeviceMappings: []interface{}{
//   		&instanceBlockDeviceMappingProperty{
//   			deviceName: jsii.String("deviceName"),
//   			ebs: &ebsInstanceBlockDeviceSpecificationProperty{
//   				deleteOnTermination: jsii.Boolean(false),
//   				encrypted: jsii.Boolean(false),
//   				iops: jsii.Number(123),
//   				kmsKeyId: jsii.String("kmsKeyId"),
//   				snapshotId: jsii.String("snapshotId"),
//   				throughput: jsii.Number(123),
//   				volumeSize: jsii.Number(123),
//   				volumeType: jsii.String("volumeType"),
//   			},
//   			noDevice: jsii.String("noDevice"),
//   			virtualName: jsii.String("virtualName"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	workingDirectory: jsii.String("workingDirectory"),
//   }
//
type CfnImageRecipeProps struct {
	// The components of the image recipe.
	//
	// Components are orchestration documents that define a sequence of steps for downloading, installing, configuring, and testing software packages. They also define validation and security hardening steps. A component is defined using a YAML document format.
	Components interface{} `field:"required" json:"components" yaml:"components"`
	// The name of the image recipe.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The parent image of the image recipe.
	//
	// The string must be either an Image ARN or an AMI ID.
	ParentImage *string `field:"required" json:"parentImage" yaml:"parentImage"`
	// The semantic version of the image recipe.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Before you create a new AMI, Image Builder launches temporary Amazon EC2 instances to build and test your image configuration.
	//
	// Instance configuration adds a layer of control over those instances. You can define settings and add scripts to run when an instance is launched from your AMI.
	AdditionalInstanceConfiguration interface{} `field:"optional" json:"additionalInstanceConfiguration" yaml:"additionalInstanceConfiguration"`
	// The block device mappings to apply when creating images from this recipe.
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// The description of the image recipe.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags of the image recipe.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The working directory to be used during build and test workflows.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

// A CloudFormation `AWS::ImageBuilder::InfrastructureConfiguration`.
//
// The infrastructure configuration allows you to specify the infrastructure within which to build and test your image. In the infrastructure configuration, you can specify instance types, subnets, and security groups to associate with your instance. You can also associate an Amazon EC2 key pair with the instance used to build your image. This allows you to log on to your instance to troubleshoot if your build fails and you set terminateInstanceOnFailure to false.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInfrastructureConfiguration := awscdk.Aws_imagebuilder.NewCfnInfrastructureConfiguration(this, jsii.String("MyCfnInfrastructureConfiguration"), &cfnInfrastructureConfigurationProps{
//   	instanceProfileName: jsii.String("instanceProfileName"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	instanceMetadataOptions: &instanceMetadataOptionsProperty{
//   		httpPutResponseHopLimit: jsii.Number(123),
//   		httpTokens: jsii.String("httpTokens"),
//   	},
//   	instanceTypes: []*string{
//   		jsii.String("instanceTypes"),
//   	},
//   	keyPair: jsii.String("keyPair"),
//   	logging: &loggingProperty{
//   		s3Logs: &s3LogsProperty{
//   			s3BucketName: jsii.String("s3BucketName"),
//   			s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   		},
//   	},
//   	resourceTags: map[string]*string{
//   		"resourceTagsKey": jsii.String("resourceTags"),
//   	},
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	snsTopicArn: jsii.String("snsTopicArn"),
//   	subnetId: jsii.String("subnetId"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	terminateInstanceOnFailure: jsii.Boolean(false),
//   })
//
type CfnInfrastructureConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Returns the Amazon Resource Name (ARN) of the infrastructure configuration.
	//
	// The following pattern is applied: `^arn:aws[^:]*:imagebuilder:[^:]+:(?:\d{12}|aws):(?:image-recipe|infrastructure-configuration|distribution-configuration|component|image|image-pipeline)/[a-z0-9-_]+(?:/(?:(?:x|\d+)\.(?:x|\d+)\.(?:x|\d+))(?:/\d+)?)?$` .
	AttrArn() *string
	// The name of the infrastructure configuration.
	AttrName() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the infrastructure configuration.
	Description() *string
	SetDescription(val *string)
	// The instance metadata option settings for the infrastructure configuration.
	InstanceMetadataOptions() interface{}
	SetInstanceMetadataOptions(val interface{})
	// The instance profile of the infrastructure configuration.
	InstanceProfileName() *string
	SetInstanceProfileName(val *string)
	// The instance types of the infrastructure configuration.
	InstanceTypes() *[]*string
	SetInstanceTypes(val *[]*string)
	// The Amazon EC2 key pair of the infrastructure configuration.
	KeyPair() *string
	SetKeyPair(val *string)
	// The logging configuration defines where Image Builder uploads your logs.
	Logging() interface{}
	SetLogging(val interface{})
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
	// The name of the infrastructure configuration.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The tags attached to the resource created by Image Builder.
	ResourceTags() interface{}
	SetResourceTags(val interface{})
	// The security group IDs of the infrastructure configuration.
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	// The Amazon Resource Name (ARN) of the SNS topic for the infrastructure configuration.
	SnsTopicArn() *string
	SetSnsTopicArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The subnet ID of the infrastructure configuration.
	SubnetId() *string
	SetSubnetId(val *string)
	// The tags of the infrastructure configuration.
	Tags() awscdk.TagManager
	// The terminate instance on failure configuration of the infrastructure configuration.
	TerminateInstanceOnFailure() interface{}
	SetTerminateInstanceOnFailure(val interface{})
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
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
	GetAtt(attributeName *string) awscdk.Reference
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

// The jsii proxy struct for CfnInfrastructureConfiguration
type jsiiProxy_CfnInfrastructureConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) InstanceMetadataOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceMetadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) InstanceProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) InstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) KeyPair() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPair",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) Logging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) ResourceTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) SnsTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) TerminateInstanceOnFailure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateInstanceOnFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ImageBuilder::InfrastructureConfiguration`.
func NewCfnInfrastructureConfiguration(scope constructs.Construct, id *string, props *CfnInfrastructureConfigurationProps) CfnInfrastructureConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnInfrastructureConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_imagebuilder.CfnInfrastructureConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ImageBuilder::InfrastructureConfiguration`.
func NewCfnInfrastructureConfiguration_Override(c CfnInfrastructureConfiguration, scope constructs.Construct, id *string, props *CfnInfrastructureConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_imagebuilder.CfnInfrastructureConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) SetInstanceMetadataOptions(val interface{}) {
	_jsii_.Set(
		j,
		"instanceMetadataOptions",
		val,
	)
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) SetInstanceProfileName(val *string) {
	_jsii_.Set(
		j,
		"instanceProfileName",
		val,
	)
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) SetInstanceTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"instanceTypes",
		val,
	)
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) SetKeyPair(val *string) {
	_jsii_.Set(
		j,
		"keyPair",
		val,
	)
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) SetLogging(val interface{}) {
	_jsii_.Set(
		j,
		"logging",
		val,
	)
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) SetResourceTags(val interface{}) {
	_jsii_.Set(
		j,
		"resourceTags",
		val,
	)
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) SetSnsTopicArn(val *string) {
	_jsii_.Set(
		j,
		"snsTopicArn",
		val,
	)
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_CfnInfrastructureConfiguration) SetTerminateInstanceOnFailure(val interface{}) {
	_jsii_.Set(
		j,
		"terminateInstanceOnFailure",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnInfrastructureConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnInfrastructureConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnInfrastructureConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnInfrastructureConfiguration",
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
func CfnInfrastructureConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnInfrastructureConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInfrastructureConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_imagebuilder.CfnInfrastructureConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInfrastructureConfiguration) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnInfrastructureConfiguration) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnInfrastructureConfiguration) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnInfrastructureConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnInfrastructureConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnInfrastructureConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnInfrastructureConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnInfrastructureConfiguration) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInfrastructureConfiguration) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInfrastructureConfiguration) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnInfrastructureConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnInfrastructureConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInfrastructureConfiguration) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInfrastructureConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInfrastructureConfiguration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The instance metadata options that apply to the HTTP requests that pipeline builds use to launch EC2 build and test instances.
//
// For more information about instance metadata options, see [Configure the instance metadata options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-instance-metadata-options.html) in the **Amazon EC2 User Guide** for Linux instances, or [Configure the instance metadata options](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/configuring-instance-metadata-options.html) in the **Amazon EC2 Windows Guide** for Windows instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceMetadataOptionsProperty := &instanceMetadataOptionsProperty{
//   	httpPutResponseHopLimit: jsii.Number(123),
//   	httpTokens: jsii.String("httpTokens"),
//   }
//
type CfnInfrastructureConfiguration_InstanceMetadataOptionsProperty struct {
	// Limit the number of hops that an instance metadata request can traverse to reach its destination.
	//
	// The default is one hop. However, if HTTP tokens are required, container image builds need a minimum of two hops.
	HttpPutResponseHopLimit *float64 `field:"optional" json:"httpPutResponseHopLimit" yaml:"httpPutResponseHopLimit"`
	// Indicates whether a signed token header is required for instance metadata retrieval requests.
	//
	// The values affect the response as follows:
	//
	// - *required* – When you retrieve the IAM role credentials, version 2.0 credentials are returned in all cases.
	// - *optional* – You can include a signed token header in your request to retrieve instance metadata, or you can leave it out. If you include it, version 2.0 credentials are returned for the IAM role. Otherwise, version 1.0 credentials are returned.
	//
	// The default setting is *optional* .
	HttpTokens *string `field:"optional" json:"httpTokens" yaml:"httpTokens"`
}

// Logging configuration defines where Image Builder uploads your logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingProperty := &loggingProperty{
//   	s3Logs: &s3LogsProperty{
//   		s3BucketName: jsii.String("s3BucketName"),
//   		s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   	},
//   }
//
type CfnInfrastructureConfiguration_LoggingProperty struct {
	// The Amazon S3 logging configuration.
	S3Logs interface{} `field:"optional" json:"s3Logs" yaml:"s3Logs"`
}

// Amazon S3 logging configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LogsProperty := &s3LogsProperty{
//   	s3BucketName: jsii.String("s3BucketName"),
//   	s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   }
//
type CfnInfrastructureConfiguration_S3LogsProperty struct {
	// The S3 bucket in which to store the logs.
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
	// The Amazon S3 path to the bucket where the logs are stored.
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
}

// Properties for defining a `CfnInfrastructureConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInfrastructureConfigurationProps := &cfnInfrastructureConfigurationProps{
//   	instanceProfileName: jsii.String("instanceProfileName"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	instanceMetadataOptions: &instanceMetadataOptionsProperty{
//   		httpPutResponseHopLimit: jsii.Number(123),
//   		httpTokens: jsii.String("httpTokens"),
//   	},
//   	instanceTypes: []*string{
//   		jsii.String("instanceTypes"),
//   	},
//   	keyPair: jsii.String("keyPair"),
//   	logging: &loggingProperty{
//   		s3Logs: &s3LogsProperty{
//   			s3BucketName: jsii.String("s3BucketName"),
//   			s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   		},
//   	},
//   	resourceTags: map[string]*string{
//   		"resourceTagsKey": jsii.String("resourceTags"),
//   	},
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	snsTopicArn: jsii.String("snsTopicArn"),
//   	subnetId: jsii.String("subnetId"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	terminateInstanceOnFailure: jsii.Boolean(false),
//   }
//
type CfnInfrastructureConfigurationProps struct {
	// The instance profile of the infrastructure configuration.
	InstanceProfileName *string `field:"required" json:"instanceProfileName" yaml:"instanceProfileName"`
	// The name of the infrastructure configuration.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the infrastructure configuration.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The instance metadata option settings for the infrastructure configuration.
	InstanceMetadataOptions interface{} `field:"optional" json:"instanceMetadataOptions" yaml:"instanceMetadataOptions"`
	// The instance types of the infrastructure configuration.
	InstanceTypes *[]*string `field:"optional" json:"instanceTypes" yaml:"instanceTypes"`
	// The Amazon EC2 key pair of the infrastructure configuration.
	KeyPair *string `field:"optional" json:"keyPair" yaml:"keyPair"`
	// The logging configuration defines where Image Builder uploads your logs.
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
	// The tags attached to the resource created by Image Builder.
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// The security group IDs of the infrastructure configuration.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The Amazon Resource Name (ARN) of the SNS topic for the infrastructure configuration.
	SnsTopicArn *string `field:"optional" json:"snsTopicArn" yaml:"snsTopicArn"`
	// The subnet ID of the infrastructure configuration.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// The tags of the infrastructure configuration.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The terminate instance on failure configuration of the infrastructure configuration.
	TerminateInstanceOnFailure interface{} `field:"optional" json:"terminateInstanceOnFailure" yaml:"terminateInstanceOnFailure"`
}


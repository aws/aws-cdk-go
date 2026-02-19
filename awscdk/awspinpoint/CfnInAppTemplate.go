package awspinpoint

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawspinpoint"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a message template that you can use to send in-app messages.
//
// A message template is a set of content and settings that you can define, save, and reuse in messages for any of your Amazon Pinpoint applications. The In-App channel is unavailable in AWS GovCloud (US).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var customConfig interface{}
//   var tags interface{}
//
//   cfnInAppTemplate := awscdk.Aws_pinpoint.NewCfnInAppTemplate(this, jsii.String("MyCfnInAppTemplate"), &CfnInAppTemplateProps{
//   	TemplateName: jsii.String("templateName"),
//
//   	// the properties below are optional
//   	Content: []interface{}{
//   		&InAppMessageContentProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			BodyConfig: &BodyConfigProperty{
//   				Alignment: jsii.String("alignment"),
//   				Body: jsii.String("body"),
//   				TextColor: jsii.String("textColor"),
//   			},
//   			HeaderConfig: &HeaderConfigProperty{
//   				Alignment: jsii.String("alignment"),
//   				Header: jsii.String("header"),
//   				TextColor: jsii.String("textColor"),
//   			},
//   			ImageUrl: jsii.String("imageUrl"),
//   			PrimaryBtn: &ButtonConfigProperty{
//   				Android: &OverrideButtonConfigurationProperty{
//   					ButtonAction: jsii.String("buttonAction"),
//   					Link: jsii.String("link"),
//   				},
//   				DefaultConfig: &DefaultButtonConfigurationProperty{
//   					BackgroundColor: jsii.String("backgroundColor"),
//   					BorderRadius: jsii.Number(123),
//   					ButtonAction: jsii.String("buttonAction"),
//   					Link: jsii.String("link"),
//   					Text: jsii.String("text"),
//   					TextColor: jsii.String("textColor"),
//   				},
//   				Ios: &OverrideButtonConfigurationProperty{
//   					ButtonAction: jsii.String("buttonAction"),
//   					Link: jsii.String("link"),
//   				},
//   				Web: &OverrideButtonConfigurationProperty{
//   					ButtonAction: jsii.String("buttonAction"),
//   					Link: jsii.String("link"),
//   				},
//   			},
//   			SecondaryBtn: &ButtonConfigProperty{
//   				Android: &OverrideButtonConfigurationProperty{
//   					ButtonAction: jsii.String("buttonAction"),
//   					Link: jsii.String("link"),
//   				},
//   				DefaultConfig: &DefaultButtonConfigurationProperty{
//   					BackgroundColor: jsii.String("backgroundColor"),
//   					BorderRadius: jsii.Number(123),
//   					ButtonAction: jsii.String("buttonAction"),
//   					Link: jsii.String("link"),
//   					Text: jsii.String("text"),
//   					TextColor: jsii.String("textColor"),
//   				},
//   				Ios: &OverrideButtonConfigurationProperty{
//   					ButtonAction: jsii.String("buttonAction"),
//   					Link: jsii.String("link"),
//   				},
//   				Web: &OverrideButtonConfigurationProperty{
//   					ButtonAction: jsii.String("buttonAction"),
//   					Link: jsii.String("link"),
//   				},
//   			},
//   		},
//   	},
//   	CustomConfig: customConfig,
//   	Layout: jsii.String("layout"),
//   	Tags: tags,
//   	TemplateDescription: jsii.String("templateDescription"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-inapptemplate.html
//
type CfnInAppTemplate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawspinpoint.IInAppTemplateRef
	awscdk.ITaggable
	// The Amazon Resource Name (ARN) of the message template.
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// An object that contains information about the content of an in-app message, including its title and body text, text colors, background colors, images, buttons, and behaviors.
	Content() interface{}
	SetContent(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Custom data, in the form of key-value pairs, that is included in an in-app messaging payload.
	CustomConfig() interface{}
	SetCustomConfig(val interface{})
	Env() *interfaces.ResourceEnvironment
	// A reference to a InAppTemplate resource.
	InAppTemplateRef() *interfacesawspinpoint.InAppTemplateReference
	// A string that determines the appearance of the in-app message.
	//
	// You can specify one of the following:.
	Layout() *string
	SetLayout(val *string)
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
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// An array of key-value pairs to apply to this resource.
	TagsRaw() interface{}
	SetTagsRaw(val interface{})
	// An optional description of the in-app template.
	TemplateDescription() *string
	SetTemplateDescription(val *string)
	// The name of the in-app message template.
	TemplateName() *string
	SetTemplateName(val *string)
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
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for CfnInAppTemplate
type jsiiProxy_CfnInAppTemplate struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawspinpointIInAppTemplateRef
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnInAppTemplate) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) Content() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) CustomConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) InAppTemplateRef() *interfacesawspinpoint.InAppTemplateReference {
	var returns *interfacesawspinpoint.InAppTemplateReference
	_jsii_.Get(
		j,
		"inAppTemplateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) Layout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) TagsRaw() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) TemplateDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) TemplateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInAppTemplate) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::Pinpoint::InAppTemplate`.
func NewCfnInAppTemplate(scope constructs.Construct, id *string, props *CfnInAppTemplateProps) CfnInAppTemplate {
	_init_.Initialize()

	if err := validateNewCfnInAppTemplateParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInAppTemplate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnInAppTemplate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Pinpoint::InAppTemplate`.
func NewCfnInAppTemplate_Override(c CfnInAppTemplate, scope constructs.Construct, id *string, props *CfnInAppTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnInAppTemplate",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnInAppTemplate)SetContent(val interface{}) {
	if err := j.validateSetContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_CfnInAppTemplate)SetCustomConfig(val interface{}) {
	_jsii_.Set(
		j,
		"customConfig",
		val,
	)
}

func (j *jsiiProxy_CfnInAppTemplate)SetLayout(val *string) {
	_jsii_.Set(
		j,
		"layout",
		val,
	)
}

func (j *jsiiProxy_CfnInAppTemplate)SetTagsRaw(val interface{}) {
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnInAppTemplate)SetTemplateDescription(val *string) {
	_jsii_.Set(
		j,
		"templateDescription",
		val,
	)
}

func (j *jsiiProxy_CfnInAppTemplate)SetTemplateName(val *string) {
	if err := j.validateSetTemplateNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateName",
		val,
	)
}

func CfnInAppTemplate_ArnForInAppTemplate(resource interfacesawspinpoint.IInAppTemplateRef) *string {
	_init_.Initialize()

	if err := validateCfnInAppTemplate_ArnForInAppTemplateParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnInAppTemplate",
		"arnForInAppTemplate",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Creates a new IInAppTemplateRef from an ARN.
func CfnInAppTemplate_FromInAppTemplateArn(scope constructs.Construct, id *string, arn *string) interfacesawspinpoint.IInAppTemplateRef {
	_init_.Initialize()

	if err := validateCfnInAppTemplate_FromInAppTemplateArnParameters(scope, id, arn); err != nil {
		panic(err)
	}
	var returns interfacesawspinpoint.IInAppTemplateRef

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnInAppTemplate",
		"fromInAppTemplateArn",
		[]interface{}{scope, id, arn},
		&returns,
	)

	return returns
}

// Creates a new IInAppTemplateRef from a templateName.
func CfnInAppTemplate_FromTemplateName(scope constructs.Construct, id *string, templateName *string) interfacesawspinpoint.IInAppTemplateRef {
	_init_.Initialize()

	if err := validateCfnInAppTemplate_FromTemplateNameParameters(scope, id, templateName); err != nil {
		panic(err)
	}
	var returns interfacesawspinpoint.IInAppTemplateRef

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnInAppTemplate",
		"fromTemplateName",
		[]interface{}{scope, id, templateName},
		&returns,
	)

	return returns
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnInAppTemplate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInAppTemplate_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnInAppTemplate",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnInAppTemplate.
func CfnInAppTemplate_IsCfnInAppTemplate(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInAppTemplate_IsCfnInAppTemplateParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnInAppTemplate",
		"isCfnInAppTemplate",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnInAppTemplate_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInAppTemplate_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnInAppTemplate",
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
func CfnInAppTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInAppTemplate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnInAppTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInAppTemplate_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpoint.CfnInAppTemplate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInAppTemplate) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnInAppTemplate) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnInAppTemplate) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnInAppTemplate) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnInAppTemplate) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnInAppTemplate) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnInAppTemplate) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnInAppTemplate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnInAppTemplate) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnInAppTemplate) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnInAppTemplate) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnInAppTemplate) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInAppTemplate) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInAppTemplate) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnInAppTemplate) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnInAppTemplate) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnInAppTemplate) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnInAppTemplate) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInAppTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInAppTemplate) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnInAppTemplate) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}


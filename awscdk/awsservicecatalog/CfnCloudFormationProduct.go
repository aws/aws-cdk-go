package awsservicecatalog

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsservicecatalog"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a product.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var info interface{}
//
//   cfnCloudFormationProduct := awscdk.Aws_servicecatalog.NewCfnCloudFormationProduct(this, jsii.String("MyCfnCloudFormationProduct"), &CfnCloudFormationProductProps{
//   	Name: jsii.String("name"),
//   	Owner: jsii.String("owner"),
//
//   	// the properties below are optional
//   	AcceptLanguage: jsii.String("acceptLanguage"),
//   	Description: jsii.String("description"),
//   	Distributor: jsii.String("distributor"),
//   	ProductType: jsii.String("productType"),
//   	ProvisioningArtifactParameters: []interface{}{
//   		&ProvisioningArtifactPropertiesProperty{
//   			Info: info,
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   			DisableTemplateValidation: jsii.Boolean(false),
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	ReplaceProvisioningArtifacts: jsii.Boolean(false),
//   	SourceConnection: &SourceConnectionProperty{
//   		ConnectionParameters: &ConnectionParametersProperty{
//   			CodeStar: &CodeStarParametersProperty{
//   				ArtifactPath: jsii.String("artifactPath"),
//   				Branch: jsii.String("branch"),
//   				ConnectionArn: jsii.String("connectionArn"),
//   				Repository: jsii.String("repository"),
//   			},
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	SupportDescription: jsii.String("supportDescription"),
//   	SupportEmail: jsii.String("supportEmail"),
//   	SupportUrl: jsii.String("supportUrl"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html
//
type CfnCloudFormationProduct interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsservicecatalog.ICloudFormationProductRef
	awscdk.ITaggable
	// The language code.
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	AttrId() *string
	// The name of the product.
	AttrProductName() *string
	// The IDs of the provisioning artifacts.
	AttrProvisioningArtifactIds() *string
	// The names of the provisioning artifacts.
	AttrProvisioningArtifactNames() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// A reference to a CloudFormationProduct resource.
	CloudFormationProductRef() *interfacesawsservicecatalog.CloudFormationProductReference
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the product.
	Description() *string
	SetDescription(val *string)
	// The distributor of the product.
	Distributor() *string
	SetDistributor(val *string)
	Env() *interfaces.ResourceEnvironment
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
	// The name of the product.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The owner of the product.
	Owner() *string
	SetOwner(val *string)
	// The type of product.
	ProductType() *string
	SetProductType(val *string)
	// The configuration of the provisioning artifact (also known as a version).
	ProvisioningArtifactParameters() interface{}
	SetProvisioningArtifactParameters(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// This property is turned off by default.
	ReplaceProvisioningArtifacts() interface{}
	SetReplaceProvisioningArtifacts(val interface{})
	// A top level `ProductViewDetail` response containing details about the product’s connection.
	SourceConnection() interface{}
	SetSourceConnection(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The support information about the product.
	SupportDescription() *string
	SetSupportDescription(val *string)
	// The contact email for product support.
	SupportEmail() *string
	SetSupportEmail(val *string)
	// The contact URL for product support.
	SupportUrl() *string
	SetSupportUrl(val *string)
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// One or more tags.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
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

// The jsii proxy struct for CfnCloudFormationProduct
type jsiiProxy_CfnCloudFormationProduct struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsservicecatalogICloudFormationProductRef
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnCloudFormationProduct) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) AttrProductName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrProductName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) AttrProvisioningArtifactIds() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrProvisioningArtifactIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) AttrProvisioningArtifactNames() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrProvisioningArtifactNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) CloudFormationProductRef() *interfacesawsservicecatalog.CloudFormationProductReference {
	var returns *interfacesawsservicecatalog.CloudFormationProductReference
	_jsii_.Get(
		j,
		"cloudFormationProductRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) Distributor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) ProductType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) ProvisioningArtifactParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisioningArtifactParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) ReplaceProvisioningArtifacts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceProvisioningArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) SourceConnection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) SupportDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) SupportEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) SupportUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProduct) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::ServiceCatalog::CloudFormationProduct`.
func NewCfnCloudFormationProduct(scope constructs.Construct, id *string, props *CfnCloudFormationProductProps) CfnCloudFormationProduct {
	_init_.Initialize()

	if err := validateNewCfnCloudFormationProductParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCloudFormationProduct{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnCloudFormationProduct",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ServiceCatalog::CloudFormationProduct`.
func NewCfnCloudFormationProduct_Override(c CfnCloudFormationProduct, scope constructs.Construct, id *string, props *CfnCloudFormationProductProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnCloudFormationProduct",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCloudFormationProduct)SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProduct)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProduct)SetDistributor(val *string) {
	_jsii_.Set(
		j,
		"distributor",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProduct)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProduct)SetOwner(val *string) {
	if err := j.validateSetOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProduct)SetProductType(val *string) {
	_jsii_.Set(
		j,
		"productType",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProduct)SetProvisioningArtifactParameters(val interface{}) {
	if err := j.validateSetProvisioningArtifactParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioningArtifactParameters",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProduct)SetReplaceProvisioningArtifacts(val interface{}) {
	if err := j.validateSetReplaceProvisioningArtifactsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceProvisioningArtifacts",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProduct)SetSourceConnection(val interface{}) {
	if err := j.validateSetSourceConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceConnection",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProduct)SetSupportDescription(val *string) {
	_jsii_.Set(
		j,
		"supportDescription",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProduct)SetSupportEmail(val *string) {
	_jsii_.Set(
		j,
		"supportEmail",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProduct)SetSupportUrl(val *string) {
	_jsii_.Set(
		j,
		"supportUrl",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProduct)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

// Checks whether the given object is a CfnCloudFormationProduct.
func CfnCloudFormationProduct_IsCfnCloudFormationProduct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCloudFormationProduct_IsCfnCloudFormationProductParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnCloudFormationProduct",
		"isCfnCloudFormationProduct",
		[]interface{}{x},
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
func CfnCloudFormationProduct_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCloudFormationProduct_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnCloudFormationProduct",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnCloudFormationProduct_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCloudFormationProduct_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnCloudFormationProduct",
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
func CfnCloudFormationProduct_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCloudFormationProduct_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnCloudFormationProduct",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCloudFormationProduct_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_servicecatalog.CfnCloudFormationProduct",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCloudFormationProduct) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCloudFormationProduct) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCloudFormationProduct) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCloudFormationProduct) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCloudFormationProduct) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCloudFormationProduct) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCloudFormationProduct) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCloudFormationProduct) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCloudFormationProduct) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnCloudFormationProduct) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnCloudFormationProduct) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCloudFormationProduct) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCloudFormationProduct) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCloudFormationProduct) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCloudFormationProduct) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCloudFormationProduct) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnCloudFormationProduct) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnCloudFormationProduct) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCloudFormationProduct) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCloudFormationProduct) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnCloudFormationProduct) With(mixins ...constructs.IMixin) constructs.IConstruct {
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


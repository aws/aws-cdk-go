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

// Provisions the specified product.
//
// A provisioned product is a resourced instance of a product. For example, provisioning a product based on a AWS CloudFormation template launches a AWS CloudFormation stack and its underlying resources. You can check the status of this request using [DescribeRecord](https://docs.aws.amazon.com/servicecatalog/latest/dg/API_DescribeRecord.html) .
//
// If the request contains a tag key with an empty list of values, there is a tag conflict for that key. Do not include conflicted keys as tags, or this causes the error "Parameter validation failed: Missing required parameter in Tags[ *N* ]: *Value* ".
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCloudFormationProvisionedProduct := awscdk.Aws_servicecatalog.NewCfnCloudFormationProvisionedProduct(this, jsii.String("MyCfnCloudFormationProvisionedProduct"), &CfnCloudFormationProvisionedProductProps{
//   	AcceptLanguage: jsii.String("acceptLanguage"),
//   	NotificationArns: []*string{
//   		jsii.String("notificationArns"),
//   	},
//   	PathId: jsii.String("pathId"),
//   	PathName: jsii.String("pathName"),
//   	ProductId: jsii.String("productId"),
//   	ProductName: jsii.String("productName"),
//   	ProvisionedProductName: jsii.String("provisionedProductName"),
//   	ProvisioningArtifactId: jsii.String("provisioningArtifactId"),
//   	ProvisioningArtifactName: jsii.String("provisioningArtifactName"),
//   	ProvisioningParameters: []interface{}{
//   		&ProvisioningParameterProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ProvisioningPreferences: &ProvisioningPreferencesProperty{
//   		StackSetAccounts: []*string{
//   			jsii.String("stackSetAccounts"),
//   		},
//   		StackSetFailureToleranceCount: jsii.Number(123),
//   		StackSetFailureTolerancePercentage: jsii.Number(123),
//   		StackSetMaxConcurrencyCount: jsii.Number(123),
//   		StackSetMaxConcurrencyPercentage: jsii.Number(123),
//   		StackSetOperationType: jsii.String("stackSetOperationType"),
//   		StackSetRegions: []*string{
//   			jsii.String("stackSetRegions"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html
//
type CfnCloudFormationProvisionedProduct interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsservicecatalog.ICloudFormationProvisionedProductRef
	awscdk.ITaggable
	// The language code.
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	AttrCloudformationStackArn() *string
	// List of key-value pair outputs.
	AttrOutputs() awscdk.IResolvable
	// The ID of the provisioned product.
	AttrProvisionedProductId() *string
	// The ID of the record, such as `rec-rjeatvy434trk` .
	AttrRecordId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// A reference to a CloudFormationProvisionedProduct resource.
	CloudFormationProvisionedProductRef() *interfacesawsservicecatalog.CloudFormationProvisionedProductReference
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// The tree node.
	Node() constructs.Node
	// Passed to AWS CloudFormation .
	NotificationArns() *[]*string
	SetNotificationArns(val *[]*string)
	// The path identifier of the product.
	PathId() *string
	SetPathId(val *string)
	// The name of the path.
	PathName() *string
	SetPathName(val *string)
	// The product identifier.
	ProductId() *string
	SetProductId(val *string)
	// The name of the Service Catalog product.
	ProductName() *string
	SetProductName(val *string)
	// A user-friendly name for the provisioned product.
	ProvisionedProductName() *string
	SetProvisionedProductName(val *string)
	// The identifier of the provisioning artifact (also known as a version).
	ProvisioningArtifactId() *string
	SetProvisioningArtifactId(val *string)
	// The name of the provisioning artifact (also known as a version) for the product.
	ProvisioningArtifactName() *string
	SetProvisioningArtifactName(val *string)
	// Parameters specified by the administrator that are required for provisioning the product.
	ProvisioningParameters() interface{}
	SetProvisioningParameters(val interface{})
	// StackSet preferences that are required for provisioning the product or updating a provisioned product.
	ProvisioningPreferences() interface{}
	SetProvisioningPreferences(val interface{})
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
}

// The jsii proxy struct for CfnCloudFormationProvisionedProduct
type jsiiProxy_CfnCloudFormationProvisionedProduct struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsservicecatalogICloudFormationProvisionedProductRef
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) AttrCloudformationStackArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCloudformationStackArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) AttrOutputs() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrOutputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) AttrProvisionedProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrProvisionedProductId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) AttrRecordId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRecordId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) CloudFormationProvisionedProductRef() *interfacesawsservicecatalog.CloudFormationProvisionedProductReference {
	var returns *interfacesawsservicecatalog.CloudFormationProvisionedProductReference
	_jsii_.Get(
		j,
		"cloudFormationProvisionedProductRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) NotificationArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) PathId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) PathName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) ProductName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) ProvisionedProductName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisionedProductName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) ProvisioningArtifactId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningArtifactId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) ProvisioningArtifactName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningArtifactName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) ProvisioningParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisioningParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) ProvisioningPreferences() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisioningPreferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::ServiceCatalog::CloudFormationProvisionedProduct`.
func NewCfnCloudFormationProvisionedProduct(scope constructs.Construct, id *string, props *CfnCloudFormationProvisionedProductProps) CfnCloudFormationProvisionedProduct {
	_init_.Initialize()

	if err := validateNewCfnCloudFormationProvisionedProductParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCloudFormationProvisionedProduct{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnCloudFormationProvisionedProduct",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ServiceCatalog::CloudFormationProvisionedProduct`.
func NewCfnCloudFormationProvisionedProduct_Override(c CfnCloudFormationProvisionedProduct, scope constructs.Construct, id *string, props *CfnCloudFormationProvisionedProductProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnCloudFormationProvisionedProduct",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct)SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct)SetNotificationArns(val *[]*string) {
	_jsii_.Set(
		j,
		"notificationArns",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct)SetPathId(val *string) {
	_jsii_.Set(
		j,
		"pathId",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct)SetPathName(val *string) {
	_jsii_.Set(
		j,
		"pathName",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct)SetProductId(val *string) {
	_jsii_.Set(
		j,
		"productId",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct)SetProductName(val *string) {
	_jsii_.Set(
		j,
		"productName",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct)SetProvisionedProductName(val *string) {
	_jsii_.Set(
		j,
		"provisionedProductName",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct)SetProvisioningArtifactId(val *string) {
	_jsii_.Set(
		j,
		"provisioningArtifactId",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct)SetProvisioningArtifactName(val *string) {
	_jsii_.Set(
		j,
		"provisioningArtifactName",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct)SetProvisioningParameters(val interface{}) {
	if err := j.validateSetProvisioningParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioningParameters",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct)SetProvisioningPreferences(val interface{}) {
	if err := j.validateSetProvisioningPreferencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioningPreferences",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCloudFormationProvisionedProduct_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCloudFormationProvisionedProduct_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnCloudFormationProvisionedProduct",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnCloudFormationProvisionedProduct_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCloudFormationProvisionedProduct_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnCloudFormationProvisionedProduct",
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
func CfnCloudFormationProvisionedProduct_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCloudFormationProvisionedProduct_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnCloudFormationProvisionedProduct",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCloudFormationProvisionedProduct_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_servicecatalog.CfnCloudFormationProvisionedProduct",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}


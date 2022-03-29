package awsservicecatalog

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::ServiceCatalog::AcceptedPortfolioShare`.
//
// Accepts an offer to share the specified portfolio.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnAcceptedPortfolioShare := servicecatalog.NewCfnAcceptedPortfolioShare(this, jsii.String("MyCfnAcceptedPortfolioShare"), &cfnAcceptedPortfolioShareProps{
//   	portfolioId: jsii.String("portfolioId"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   })
//
type CfnAcceptedPortfolioShare interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// The portfolio identifier.
	PortfolioId() *string
	SetPortfolioId(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnAcceptedPortfolioShare
type jsiiProxy_CfnAcceptedPortfolioShare struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAcceptedPortfolioShare) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAcceptedPortfolioShare) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAcceptedPortfolioShare) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAcceptedPortfolioShare) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAcceptedPortfolioShare) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAcceptedPortfolioShare) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAcceptedPortfolioShare) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAcceptedPortfolioShare) PortfolioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAcceptedPortfolioShare) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAcceptedPortfolioShare) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAcceptedPortfolioShare) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ServiceCatalog::AcceptedPortfolioShare`.
func NewCfnAcceptedPortfolioShare(scope constructs.Construct, id *string, props *CfnAcceptedPortfolioShareProps) CfnAcceptedPortfolioShare {
	_init_.Initialize()

	j := jsiiProxy_CfnAcceptedPortfolioShare{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnAcceptedPortfolioShare",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ServiceCatalog::AcceptedPortfolioShare`.
func NewCfnAcceptedPortfolioShare_Override(c CfnAcceptedPortfolioShare, scope constructs.Construct, id *string, props *CfnAcceptedPortfolioShareProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnAcceptedPortfolioShare",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAcceptedPortfolioShare) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_CfnAcceptedPortfolioShare) SetPortfolioId(val *string) {
	_jsii_.Set(
		j,
		"portfolioId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAcceptedPortfolioShare_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnAcceptedPortfolioShare",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAcceptedPortfolioShare_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnAcceptedPortfolioShare",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnAcceptedPortfolioShare_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnAcceptedPortfolioShare",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAcceptedPortfolioShare_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_servicecatalog.CfnAcceptedPortfolioShare",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAcceptedPortfolioShare) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAcceptedPortfolioShare) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAcceptedPortfolioShare) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAcceptedPortfolioShare) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAcceptedPortfolioShare) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAcceptedPortfolioShare) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAcceptedPortfolioShare) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAcceptedPortfolioShare) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAcceptedPortfolioShare) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAcceptedPortfolioShare) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAcceptedPortfolioShare) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAcceptedPortfolioShare) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAcceptedPortfolioShare) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAcceptedPortfolioShare) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAcceptedPortfolioShare) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnAcceptedPortfolioShare`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnAcceptedPortfolioShareProps := &cfnAcceptedPortfolioShareProps{
//   	portfolioId: jsii.String("portfolioId"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   }
//
type CfnAcceptedPortfolioShareProps struct {
	// The portfolio identifier.
	PortfolioId *string `json:"portfolioId" yaml:"portfolioId"`
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
}

// A CloudFormation `AWS::ServiceCatalog::CloudFormationProduct`.
//
// Specifies a product.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//
//   var info interface{}
//   cfnCloudFormationProduct := servicecatalog.NewCfnCloudFormationProduct(this, jsii.String("MyCfnCloudFormationProduct"), &cfnCloudFormationProductProps{
//   	name: jsii.String("name"),
//   	owner: jsii.String("owner"),
//   	provisioningArtifactParameters: []interface{}{
//   		&provisioningArtifactPropertiesProperty{
//   			info: info,
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
//   			disableTemplateValidation: jsii.Boolean(false),
//   			name: jsii.String("name"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	description: jsii.String("description"),
//   	distributor: jsii.String("distributor"),
//   	replaceProvisioningArtifacts: jsii.Boolean(false),
//   	supportDescription: jsii.String("supportDescription"),
//   	supportEmail: jsii.String("supportEmail"),
//   	supportUrl: jsii.String("supportUrl"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnCloudFormationProduct interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
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
	// The configuration of the provisioning artifact (also known as a version).
	ProvisioningArtifactParameters() interface{}
	SetProvisioningArtifactParameters(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// This property is turned off by default.
	//
	// If turned off, you can update provisioning artifacts or product attributes (such as description, distributor, name, owner, and more) and the associated provisioning artifacts will retain the same unique identifier. Provisioning artifacts are matched within the CloudFormationProduct resource, and only those that have been updated will be changed. Provisioning artifacts are matched by a combinaton of provisioning artifact template URL and name.
	//
	// If turned on, provisioning artifacts will be given a new unique identifier when you update the product or provisioning artifacts.
	ReplaceProvisioningArtifacts() interface{}
	SetReplaceProvisioningArtifacts(val interface{})
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
	//
	// `^https?:\/\//` / is the pattern used to validate SupportUrl.
	SupportUrl() *string
	SetSupportUrl(val *string)
	// One or more tags.
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnCloudFormationProduct
type jsiiProxy_CfnCloudFormationProduct struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
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

func (j *jsiiProxy_CfnCloudFormationProduct) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ServiceCatalog::CloudFormationProduct`.
func NewCfnCloudFormationProduct(scope constructs.Construct, id *string, props *CfnCloudFormationProductProps) CfnCloudFormationProduct {
	_init_.Initialize()

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

func (j *jsiiProxy_CfnCloudFormationProduct) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProduct) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProduct) SetDistributor(val *string) {
	_jsii_.Set(
		j,
		"distributor",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProduct) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProduct) SetOwner(val *string) {
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProduct) SetProvisioningArtifactParameters(val interface{}) {
	_jsii_.Set(
		j,
		"provisioningArtifactParameters",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProduct) SetReplaceProvisioningArtifacts(val interface{}) {
	_jsii_.Set(
		j,
		"replaceProvisioningArtifacts",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProduct) SetSupportDescription(val *string) {
	_jsii_.Set(
		j,
		"supportDescription",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProduct) SetSupportEmail(val *string) {
	_jsii_.Set(
		j,
		"supportEmail",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProduct) SetSupportUrl(val *string) {
	_jsii_.Set(
		j,
		"supportUrl",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCloudFormationProduct_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnCloudFormationProduct",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCloudFormationProduct_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnCloudFormationProduct",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnCloudFormationProduct_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

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
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCloudFormationProduct) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCloudFormationProduct) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCloudFormationProduct) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCloudFormationProduct) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCloudFormationProduct) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCloudFormationProduct) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCloudFormationProduct) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCloudFormationProduct) GetMetadata(key *string) interface{} {
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
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCloudFormationProduct) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCloudFormationProduct) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
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
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Information about a provisioning artifact (also known as a version) for a product.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//
//   var info interface{}
//   provisioningArtifactPropertiesProperty := &provisioningArtifactPropertiesProperty{
//   	info: info,
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	disableTemplateValidation: jsii.Boolean(false),
//   	name: jsii.String("name"),
//   }
//
type CfnCloudFormationProduct_ProvisioningArtifactPropertiesProperty struct {
	// Specify the template source with one of the following options, but not both.
	//
	// Keys accepted: [ `LoadTemplateFromURL` , `ImportFromPhysicalId` ]
	//
	// The URL of the AWS CloudFormation template in Amazon S3, AWS CodeCommit, or GitHub in JSON format. Specify the URL in JSON format as follows:
	//
	// `"LoadTemplateFromURL": "https://s3.amazonaws.com/cf-templates-ozkq9d3hgiq2-us-east-1/..."`
	//
	// `ImportFromPhysicalId` : The physical id of the resource that contains the template. Currently only supports AWS CloudFormation stack arn. Specify the physical id in JSON format as follows: `ImportFromPhysicalId: “arn:aws:cloudformation:[us-east-1]:[accountId]:stack/[StackName]/[resourceId]`
	Info interface{} `json:"info" yaml:"info"`
	// The description of the provisioning artifact, including how it differs from the previous provisioning artifact.
	Description *string `json:"description" yaml:"description"`
	// If set to true, AWS Service Catalog stops validating the specified provisioning artifact even if it is invalid.
	DisableTemplateValidation interface{} `json:"disableTemplateValidation" yaml:"disableTemplateValidation"`
	// The name of the provisioning artifact (for example, v1 v2beta).
	//
	// No spaces are allowed.
	Name *string `json:"name" yaml:"name"`
}

// Properties for defining a `CfnCloudFormationProduct`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//
//   var info interface{}
//   cfnCloudFormationProductProps := &cfnCloudFormationProductProps{
//   	name: jsii.String("name"),
//   	owner: jsii.String("owner"),
//   	provisioningArtifactParameters: []interface{}{
//   		&provisioningArtifactPropertiesProperty{
//   			info: info,
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
//   			disableTemplateValidation: jsii.Boolean(false),
//   			name: jsii.String("name"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	description: jsii.String("description"),
//   	distributor: jsii.String("distributor"),
//   	replaceProvisioningArtifacts: jsii.Boolean(false),
//   	supportDescription: jsii.String("supportDescription"),
//   	supportEmail: jsii.String("supportEmail"),
//   	supportUrl: jsii.String("supportUrl"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnCloudFormationProductProps struct {
	// The name of the product.
	Name *string `json:"name" yaml:"name"`
	// The owner of the product.
	Owner *string `json:"owner" yaml:"owner"`
	// The configuration of the provisioning artifact (also known as a version).
	ProvisioningArtifactParameters interface{} `json:"provisioningArtifactParameters" yaml:"provisioningArtifactParameters"`
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
	// The description of the product.
	Description *string `json:"description" yaml:"description"`
	// The distributor of the product.
	Distributor *string `json:"distributor" yaml:"distributor"`
	// This property is turned off by default.
	//
	// If turned off, you can update provisioning artifacts or product attributes (such as description, distributor, name, owner, and more) and the associated provisioning artifacts will retain the same unique identifier. Provisioning artifacts are matched within the CloudFormationProduct resource, and only those that have been updated will be changed. Provisioning artifacts are matched by a combinaton of provisioning artifact template URL and name.
	//
	// If turned on, provisioning artifacts will be given a new unique identifier when you update the product or provisioning artifacts.
	ReplaceProvisioningArtifacts interface{} `json:"replaceProvisioningArtifacts" yaml:"replaceProvisioningArtifacts"`
	// The support information about the product.
	SupportDescription *string `json:"supportDescription" yaml:"supportDescription"`
	// The contact email for product support.
	SupportEmail *string `json:"supportEmail" yaml:"supportEmail"`
	// The contact URL for product support.
	//
	// `^https?:\/\//` / is the pattern used to validate SupportUrl.
	SupportUrl *string `json:"supportUrl" yaml:"supportUrl"`
	// One or more tags.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::ServiceCatalog::CloudFormationProvisionedProduct`.
//
// Provisions the specified product.
//
// A provisioned product is a resourced instance of a product. For example, provisioning a product based on a AWS CloudFormation template launches a AWS CloudFormation stack and its underlying resources. You can check the status of this request using [DescribeRecord](https://docs.aws.amazon.com/servicecatalog/latest/dg/API_DescribeRecord.html) .
//
// If the request contains a tag key with an empty list of values, there is a tag conflict for that key. Do not include conflicted keys as tags, or this causes the error "Parameter validation failed: Missing required parameter in Tags[ *N* ]: *Value* ".
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnCloudFormationProvisionedProduct := servicecatalog.NewCfnCloudFormationProvisionedProduct(this, jsii.String("MyCfnCloudFormationProvisionedProduct"), &cfnCloudFormationProvisionedProductProps{
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	notificationArns: []*string{
//   		jsii.String("notificationArns"),
//   	},
//   	pathId: jsii.String("pathId"),
//   	pathName: jsii.String("pathName"),
//   	productId: jsii.String("productId"),
//   	productName: jsii.String("productName"),
//   	provisionedProductName: jsii.String("provisionedProductName"),
//   	provisioningArtifactId: jsii.String("provisioningArtifactId"),
//   	provisioningArtifactName: jsii.String("provisioningArtifactName"),
//   	provisioningParameters: []interface{}{
//   		&provisioningParameterProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	provisioningPreferences: &provisioningPreferencesProperty{
//   		stackSetAccounts: []*string{
//   			jsii.String("stackSetAccounts"),
//   		},
//   		stackSetFailureToleranceCount: jsii.Number(123),
//   		stackSetFailureTolerancePercentage: jsii.Number(123),
//   		stackSetMaxConcurrencyCount: jsii.Number(123),
//   		stackSetMaxConcurrencyPercentage: jsii.Number(123),
//   		stackSetOperationType: jsii.String("stackSetOperationType"),
//   		stackSetRegions: []*string{
//   			jsii.String("stackSetRegions"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnCloudFormationProvisionedProduct interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	// The Amazon Resource Name (ARN) of the CloudFormation stack, such as `arn:aws:cloudformation:eu-west-1:123456789012:stack/SC-499278721343-pp-hfyszaotincww/8f3df460-346a-11e8-9444-503abe701c29` .
	AttrCloudformationStackArn() *string
	// The ID of the provisioned product.
	AttrProvisionedProductId() *string
	// The ID of the record, such as `rec-rjeatvy434trk` .
	AttrRecordId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	//
	// The SNS topic ARNs to which to publish stack-related events.
	NotificationArns() *[]*string
	SetNotificationArns(val *[]*string)
	// The path identifier of the product.
	//
	// This value is optional if the product has a default path, and required if the product has more than one path. To list the paths for a product, use [ListLaunchPaths](https://docs.aws.amazon.com/servicecatalog/latest/dg/API_ListLaunchPaths.html) .
	//
	// > You must provide the name or ID, but not both.
	PathId() *string
	SetPathId(val *string)
	// The name of the path.
	//
	// This value is optional if the product has a default path, and required if the product has more than one path. To list the paths for a product, use [ListLaunchPaths](https://docs.aws.amazon.com/servicecatalog/latest/dg/API_ListLaunchPaths.html) .
	//
	// > You must provide the name or ID, but not both.
	PathName() *string
	SetPathName(val *string)
	// The product identifier.
	//
	// > You must specify either the ID or the name of the product, but not both.
	ProductId() *string
	SetProductId(val *string)
	// A user-friendly name for the provisioned product.
	//
	// This value must be unique for the AWS account and cannot be updated after the product is provisioned.
	//
	// Each time a stack is created or updated, if `ProductName` is provided it will successfully resolve to `ProductId` as long as only one product exists in the account or Region with that `ProductName` .
	//
	// > You must specify either the name or the ID of the product, but not both.
	ProductName() *string
	SetProductName(val *string)
	// A user-friendly name for the provisioned product.
	//
	// This value must be unique for the AWS account and cannot be updated after the product is provisioned.
	ProvisionedProductName() *string
	SetProvisionedProductName(val *string)
	// The identifier of the provisioning artifact (also known as a version).
	//
	// > You must specify either the ID or the name of the provisioning artifact, but not both.
	ProvisioningArtifactId() *string
	SetProvisioningArtifactId(val *string)
	// The name of the provisioning artifact (also known as a version) for the product.
	//
	// This name must be unique for the product.
	//
	// > You must specify either the name or the ID of the provisioning artifact, but not both. You must also specify either the name or the ID of the product, but not both.
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
	// One or more tags.
	//
	// > Requires the provisioned product to have an [ResourceUpdateConstraint](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-resourceupdateconstraint.html) resource with `TagUpdatesOnProvisionedProduct` set to `ALLOWED` to allow tag updates. If `RESOURCE_UPDATE` constraint is not present, tags updates are ignored.
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnCloudFormationProvisionedProduct
type jsiiProxy_CfnCloudFormationProvisionedProduct struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
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

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
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

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ServiceCatalog::CloudFormationProvisionedProduct`.
func NewCfnCloudFormationProvisionedProduct(scope constructs.Construct, id *string, props *CfnCloudFormationProvisionedProductProps) CfnCloudFormationProvisionedProduct {
	_init_.Initialize()

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

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) SetNotificationArns(val *[]*string) {
	_jsii_.Set(
		j,
		"notificationArns",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) SetPathId(val *string) {
	_jsii_.Set(
		j,
		"pathId",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) SetPathName(val *string) {
	_jsii_.Set(
		j,
		"pathName",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) SetProductId(val *string) {
	_jsii_.Set(
		j,
		"productId",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) SetProductName(val *string) {
	_jsii_.Set(
		j,
		"productName",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) SetProvisionedProductName(val *string) {
	_jsii_.Set(
		j,
		"provisionedProductName",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) SetProvisioningArtifactId(val *string) {
	_jsii_.Set(
		j,
		"provisioningArtifactId",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) SetProvisioningArtifactName(val *string) {
	_jsii_.Set(
		j,
		"provisioningArtifactName",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) SetProvisioningParameters(val interface{}) {
	_jsii_.Set(
		j,
		"provisioningParameters",
		val,
	)
}

func (j *jsiiProxy_CfnCloudFormationProvisionedProduct) SetProvisioningPreferences(val interface{}) {
	_jsii_.Set(
		j,
		"provisioningPreferences",
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

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnCloudFormationProvisionedProduct",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCloudFormationProvisionedProduct_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnCloudFormationProvisionedProduct",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnCloudFormationProvisionedProduct_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

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
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) GetMetadata(key *string) interface{} {
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
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCloudFormationProvisionedProduct) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
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
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Information about a parameter used to provision a product.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   provisioningParameterProperty := &provisioningParameterProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnCloudFormationProvisionedProduct_ProvisioningParameterProperty struct {
	// The parameter key.
	Key *string `json:"key" yaml:"key"`
	// The parameter value.
	Value *string `json:"value" yaml:"value"`
}

// The user-defined preferences that will be applied when updating a provisioned product.
//
// Not all preferences are applicable to all provisioned product type
//
// One or more AWS accounts that will have access to the provisioned product.
//
// Applicable only to a `CFN_STACKSET` provisioned product type.
//
// The AWS accounts specified should be within the list of accounts in the `STACKSET` constraint. To get the list of accounts in the `STACKSET` constraint, use the `DescribeProvisioningParameters` operation.
//
// If no values are specified, the default value is all accounts from the `STACKSET` constraint.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   provisioningPreferencesProperty := &provisioningPreferencesProperty{
//   	stackSetAccounts: []*string{
//   		jsii.String("stackSetAccounts"),
//   	},
//   	stackSetFailureToleranceCount: jsii.Number(123),
//   	stackSetFailureTolerancePercentage: jsii.Number(123),
//   	stackSetMaxConcurrencyCount: jsii.Number(123),
//   	stackSetMaxConcurrencyPercentage: jsii.Number(123),
//   	stackSetOperationType: jsii.String("stackSetOperationType"),
//   	stackSetRegions: []*string{
//   		jsii.String("stackSetRegions"),
//   	},
//   }
//
type CfnCloudFormationProvisionedProduct_ProvisioningPreferencesProperty struct {
	// One or more AWS accounts where the provisioned product will be available.
	//
	// Applicable only to a `CFN_STACKSET` provisioned product type.
	//
	// The specified accounts should be within the list of accounts from the `STACKSET` constraint. To get the list of accounts in the `STACKSET` constraint, use the `DescribeProvisioningParameters` operation.
	//
	// If no values are specified, the default value is all acounts from the `STACKSET` constraint.
	StackSetAccounts *[]*string `json:"stackSetAccounts" yaml:"stackSetAccounts"`
	// The number of accounts, per Region, for which this operation can fail before AWS Service Catalog stops the operation in that Region.
	//
	// If the operation is stopped in a Region, AWS Service Catalog doesn't attempt the operation in any subsequent Regions.
	//
	// Applicable only to a `CFN_STACKSET` provisioned product type.
	//
	// Conditional: You must specify either `StackSetFailureToleranceCount` or `StackSetFailureTolerancePercentage` , but not both.
	//
	// The default value is `0` if no value is specified.
	StackSetFailureToleranceCount *float64 `json:"stackSetFailureToleranceCount" yaml:"stackSetFailureToleranceCount"`
	// The percentage of accounts, per Region, for which this stack operation can fail before AWS Service Catalog stops the operation in that Region.
	//
	// If the operation is stopped in a Region, AWS Service Catalog doesn't attempt the operation in any subsequent Regions.
	//
	// When calculating the number of accounts based on the specified percentage, AWS Service Catalog rounds down to the next whole number.
	//
	// Applicable only to a `CFN_STACKSET` provisioned product type.
	//
	// Conditional: You must specify either `StackSetFailureToleranceCount` or `StackSetFailureTolerancePercentage` , but not both.
	StackSetFailureTolerancePercentage *float64 `json:"stackSetFailureTolerancePercentage" yaml:"stackSetFailureTolerancePercentage"`
	// The maximum number of accounts in which to perform this operation at one time.
	//
	// This is dependent on the value of `StackSetFailureToleranceCount` . `StackSetMaxConcurrentCount` is at most one more than the `StackSetFailureToleranceCount` .
	//
	// Note that this setting lets you specify the maximum for operations. For large deployments, under certain circumstances the actual number of accounts acted upon concurrently may be lower due to service throttling.
	//
	// Applicable only to a `CFN_STACKSET` provisioned product type.
	//
	// Conditional: You must specify either `StackSetMaxConcurrentCount` or `StackSetMaxConcurrentPercentage` , but not both.
	StackSetMaxConcurrencyCount *float64 `json:"stackSetMaxConcurrencyCount" yaml:"stackSetMaxConcurrencyCount"`
	// The maximum percentage of accounts in which to perform this operation at one time.
	//
	// When calculating the number of accounts based on the specified percentage, AWS Service Catalog rounds down to the next whole number. This is true except in cases where rounding down would result is zero. In this case, AWS Service Catalog sets the number as `1` instead.
	//
	// Note that this setting lets you specify the maximum for operations. For large deployments, under certain circumstances the actual number of accounts acted upon concurrently may be lower due to service throttling.
	//
	// Applicable only to a `CFN_STACKSET` provisioned product type.
	//
	// Conditional: You must specify either `StackSetMaxConcurrentCount` or `StackSetMaxConcurrentPercentage` , but not both.
	StackSetMaxConcurrencyPercentage *float64 `json:"stackSetMaxConcurrencyPercentage" yaml:"stackSetMaxConcurrencyPercentage"`
	// Determines what action AWS Service Catalog performs to a stack set or a stack instance represented by the provisioned product.
	//
	// The default value is `UPDATE` if nothing is specified.
	//
	// Applicable only to a `CFN_STACKSET` provisioned product type.
	//
	// - **CREATE** - Creates a new stack instance in the stack set represented by the provisioned product. In this case, only new stack instances are created based on accounts and Regions; if new ProductId or ProvisioningArtifactID are passed, they will be ignored.
	// - **UPDATE** - Updates the stack set represented by the provisioned product and also its stack instances.
	// - **DELETE** - Deletes a stack instance in the stack set represented by the provisioned product.
	StackSetOperationType *string `json:"stackSetOperationType" yaml:"stackSetOperationType"`
	// One or more AWS Regions where the provisioned product will be available.
	//
	// Applicable only to a `CFN_STACKSET` provisioned product type.
	//
	// The specified Regions should be within the list of Regions from the `STACKSET` constraint. To get the list of Regions in the `STACKSET` constraint, use the `DescribeProvisioningParameters` operation.
	//
	// If no values are specified, the default value is all Regions from the `STACKSET` constraint.
	StackSetRegions *[]*string `json:"stackSetRegions" yaml:"stackSetRegions"`
}

// Properties for defining a `CfnCloudFormationProvisionedProduct`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnCloudFormationProvisionedProductProps := &cfnCloudFormationProvisionedProductProps{
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	notificationArns: []*string{
//   		jsii.String("notificationArns"),
//   	},
//   	pathId: jsii.String("pathId"),
//   	pathName: jsii.String("pathName"),
//   	productId: jsii.String("productId"),
//   	productName: jsii.String("productName"),
//   	provisionedProductName: jsii.String("provisionedProductName"),
//   	provisioningArtifactId: jsii.String("provisioningArtifactId"),
//   	provisioningArtifactName: jsii.String("provisioningArtifactName"),
//   	provisioningParameters: []interface{}{
//   		&provisioningParameterProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	provisioningPreferences: &provisioningPreferencesProperty{
//   		stackSetAccounts: []*string{
//   			jsii.String("stackSetAccounts"),
//   		},
//   		stackSetFailureToleranceCount: jsii.Number(123),
//   		stackSetFailureTolerancePercentage: jsii.Number(123),
//   		stackSetMaxConcurrencyCount: jsii.Number(123),
//   		stackSetMaxConcurrencyPercentage: jsii.Number(123),
//   		stackSetOperationType: jsii.String("stackSetOperationType"),
//   		stackSetRegions: []*string{
//   			jsii.String("stackSetRegions"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnCloudFormationProvisionedProductProps struct {
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
	// Passed to AWS CloudFormation .
	//
	// The SNS topic ARNs to which to publish stack-related events.
	NotificationArns *[]*string `json:"notificationArns" yaml:"notificationArns"`
	// The path identifier of the product.
	//
	// This value is optional if the product has a default path, and required if the product has more than one path. To list the paths for a product, use [ListLaunchPaths](https://docs.aws.amazon.com/servicecatalog/latest/dg/API_ListLaunchPaths.html) .
	//
	// > You must provide the name or ID, but not both.
	PathId *string `json:"pathId" yaml:"pathId"`
	// The name of the path.
	//
	// This value is optional if the product has a default path, and required if the product has more than one path. To list the paths for a product, use [ListLaunchPaths](https://docs.aws.amazon.com/servicecatalog/latest/dg/API_ListLaunchPaths.html) .
	//
	// > You must provide the name or ID, but not both.
	PathName *string `json:"pathName" yaml:"pathName"`
	// The product identifier.
	//
	// > You must specify either the ID or the name of the product, but not both.
	ProductId *string `json:"productId" yaml:"productId"`
	// A user-friendly name for the provisioned product.
	//
	// This value must be unique for the AWS account and cannot be updated after the product is provisioned.
	//
	// Each time a stack is created or updated, if `ProductName` is provided it will successfully resolve to `ProductId` as long as only one product exists in the account or Region with that `ProductName` .
	//
	// > You must specify either the name or the ID of the product, but not both.
	ProductName *string `json:"productName" yaml:"productName"`
	// A user-friendly name for the provisioned product.
	//
	// This value must be unique for the AWS account and cannot be updated after the product is provisioned.
	ProvisionedProductName *string `json:"provisionedProductName" yaml:"provisionedProductName"`
	// The identifier of the provisioning artifact (also known as a version).
	//
	// > You must specify either the ID or the name of the provisioning artifact, but not both.
	ProvisioningArtifactId *string `json:"provisioningArtifactId" yaml:"provisioningArtifactId"`
	// The name of the provisioning artifact (also known as a version) for the product.
	//
	// This name must be unique for the product.
	//
	// > You must specify either the name or the ID of the provisioning artifact, but not both. You must also specify either the name or the ID of the product, but not both.
	ProvisioningArtifactName *string `json:"provisioningArtifactName" yaml:"provisioningArtifactName"`
	// Parameters specified by the administrator that are required for provisioning the product.
	ProvisioningParameters interface{} `json:"provisioningParameters" yaml:"provisioningParameters"`
	// StackSet preferences that are required for provisioning the product or updating a provisioned product.
	ProvisioningPreferences interface{} `json:"provisioningPreferences" yaml:"provisioningPreferences"`
	// One or more tags.
	//
	// > Requires the provisioned product to have an [ResourceUpdateConstraint](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-resourceupdateconstraint.html) resource with `TagUpdatesOnProvisionedProduct` set to `ALLOWED` to allow tag updates. If `RESOURCE_UPDATE` constraint is not present, tags updates are ignored.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::ServiceCatalog::LaunchNotificationConstraint`.
//
// Specifies a notification constraint.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnLaunchNotificationConstraint := servicecatalog.NewCfnLaunchNotificationConstraint(this, jsii.String("MyCfnLaunchNotificationConstraint"), &cfnLaunchNotificationConstraintProps{
//   	notificationArns: []*string{
//   		jsii.String("notificationArns"),
//   	},
//   	portfolioId: jsii.String("portfolioId"),
//   	productId: jsii.String("productId"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	description: jsii.String("description"),
//   })
//
type CfnLaunchNotificationConstraint interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the constraint.
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
	// The tree node.
	Node() constructs.Node
	// The notification ARNs.
	NotificationArns() *[]*string
	SetNotificationArns(val *[]*string)
	// The portfolio identifier.
	PortfolioId() *string
	SetPortfolioId(val *string)
	// The product identifier.
	ProductId() *string
	SetProductId(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnLaunchNotificationConstraint
type jsiiProxy_CfnLaunchNotificationConstraint struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLaunchNotificationConstraint) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchNotificationConstraint) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchNotificationConstraint) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchNotificationConstraint) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchNotificationConstraint) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchNotificationConstraint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchNotificationConstraint) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchNotificationConstraint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchNotificationConstraint) NotificationArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchNotificationConstraint) PortfolioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchNotificationConstraint) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchNotificationConstraint) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchNotificationConstraint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchNotificationConstraint) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ServiceCatalog::LaunchNotificationConstraint`.
func NewCfnLaunchNotificationConstraint(scope constructs.Construct, id *string, props *CfnLaunchNotificationConstraintProps) CfnLaunchNotificationConstraint {
	_init_.Initialize()

	j := jsiiProxy_CfnLaunchNotificationConstraint{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnLaunchNotificationConstraint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ServiceCatalog::LaunchNotificationConstraint`.
func NewCfnLaunchNotificationConstraint_Override(c CfnLaunchNotificationConstraint, scope constructs.Construct, id *string, props *CfnLaunchNotificationConstraintProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnLaunchNotificationConstraint",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLaunchNotificationConstraint) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchNotificationConstraint) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchNotificationConstraint) SetNotificationArns(val *[]*string) {
	_jsii_.Set(
		j,
		"notificationArns",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchNotificationConstraint) SetPortfolioId(val *string) {
	_jsii_.Set(
		j,
		"portfolioId",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchNotificationConstraint) SetProductId(val *string) {
	_jsii_.Set(
		j,
		"productId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLaunchNotificationConstraint_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnLaunchNotificationConstraint",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLaunchNotificationConstraint_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnLaunchNotificationConstraint",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnLaunchNotificationConstraint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnLaunchNotificationConstraint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLaunchNotificationConstraint_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_servicecatalog.CfnLaunchNotificationConstraint",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLaunchNotificationConstraint) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnLaunchNotificationConstraint) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLaunchNotificationConstraint) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnLaunchNotificationConstraint) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnLaunchNotificationConstraint) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnLaunchNotificationConstraint) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnLaunchNotificationConstraint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnLaunchNotificationConstraint) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunchNotificationConstraint) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunchNotificationConstraint) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnLaunchNotificationConstraint) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLaunchNotificationConstraint) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunchNotificationConstraint) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunchNotificationConstraint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunchNotificationConstraint) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnLaunchNotificationConstraint`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnLaunchNotificationConstraintProps := &cfnLaunchNotificationConstraintProps{
//   	notificationArns: []*string{
//   		jsii.String("notificationArns"),
//   	},
//   	portfolioId: jsii.String("portfolioId"),
//   	productId: jsii.String("productId"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	description: jsii.String("description"),
//   }
//
type CfnLaunchNotificationConstraintProps struct {
	// The notification ARNs.
	NotificationArns *[]*string `json:"notificationArns" yaml:"notificationArns"`
	// The portfolio identifier.
	PortfolioId *string `json:"portfolioId" yaml:"portfolioId"`
	// The product identifier.
	ProductId *string `json:"productId" yaml:"productId"`
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
	// The description of the constraint.
	Description *string `json:"description" yaml:"description"`
}

// A CloudFormation `AWS::ServiceCatalog::LaunchRoleConstraint`.
//
// Specifies a launch constraint.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnLaunchRoleConstraint := servicecatalog.NewCfnLaunchRoleConstraint(this, jsii.String("MyCfnLaunchRoleConstraint"), &cfnLaunchRoleConstraintProps{
//   	portfolioId: jsii.String("portfolioId"),
//   	productId: jsii.String("productId"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	description: jsii.String("description"),
//   	localRoleName: jsii.String("localRoleName"),
//   	roleArn: jsii.String("roleArn"),
//   })
//
type CfnLaunchRoleConstraint interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the constraint.
	Description() *string
	SetDescription(val *string)
	// You are required to specify either the `RoleArn` or the `LocalRoleName` but can't use both.
	//
	// If you specify the `LocalRoleName` property, when an account uses the launch constraint, the IAM role with that name in the account will be used. This allows launch-role constraints to be account-agnostic so the administrator can create fewer resources per shared account.
	//
	// The given role name must exist in the account used to create the launch constraint and the account of the user who launches a product with this launch constraint.
	LocalRoleName() *string
	SetLocalRoleName(val *string)
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
	// The portfolio identifier.
	PortfolioId() *string
	SetPortfolioId(val *string)
	// The product identifier.
	ProductId() *string
	SetProductId(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The ARN of the launch role.
	//
	// You are required to specify `RoleArn` or `LocalRoleName` but can't use both.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnLaunchRoleConstraint
type jsiiProxy_CfnLaunchRoleConstraint struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLaunchRoleConstraint) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchRoleConstraint) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchRoleConstraint) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchRoleConstraint) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchRoleConstraint) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchRoleConstraint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchRoleConstraint) LocalRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchRoleConstraint) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchRoleConstraint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchRoleConstraint) PortfolioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchRoleConstraint) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchRoleConstraint) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchRoleConstraint) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchRoleConstraint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchRoleConstraint) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ServiceCatalog::LaunchRoleConstraint`.
func NewCfnLaunchRoleConstraint(scope constructs.Construct, id *string, props *CfnLaunchRoleConstraintProps) CfnLaunchRoleConstraint {
	_init_.Initialize()

	j := jsiiProxy_CfnLaunchRoleConstraint{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnLaunchRoleConstraint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ServiceCatalog::LaunchRoleConstraint`.
func NewCfnLaunchRoleConstraint_Override(c CfnLaunchRoleConstraint, scope constructs.Construct, id *string, props *CfnLaunchRoleConstraintProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnLaunchRoleConstraint",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLaunchRoleConstraint) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchRoleConstraint) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchRoleConstraint) SetLocalRoleName(val *string) {
	_jsii_.Set(
		j,
		"localRoleName",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchRoleConstraint) SetPortfolioId(val *string) {
	_jsii_.Set(
		j,
		"portfolioId",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchRoleConstraint) SetProductId(val *string) {
	_jsii_.Set(
		j,
		"productId",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchRoleConstraint) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLaunchRoleConstraint_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnLaunchRoleConstraint",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLaunchRoleConstraint_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnLaunchRoleConstraint",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnLaunchRoleConstraint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnLaunchRoleConstraint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLaunchRoleConstraint_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_servicecatalog.CfnLaunchRoleConstraint",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLaunchRoleConstraint) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnLaunchRoleConstraint) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLaunchRoleConstraint) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnLaunchRoleConstraint) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnLaunchRoleConstraint) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnLaunchRoleConstraint) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnLaunchRoleConstraint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnLaunchRoleConstraint) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunchRoleConstraint) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunchRoleConstraint) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnLaunchRoleConstraint) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLaunchRoleConstraint) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunchRoleConstraint) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunchRoleConstraint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunchRoleConstraint) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnLaunchRoleConstraint`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnLaunchRoleConstraintProps := &cfnLaunchRoleConstraintProps{
//   	portfolioId: jsii.String("portfolioId"),
//   	productId: jsii.String("productId"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	description: jsii.String("description"),
//   	localRoleName: jsii.String("localRoleName"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnLaunchRoleConstraintProps struct {
	// The portfolio identifier.
	PortfolioId *string `json:"portfolioId" yaml:"portfolioId"`
	// The product identifier.
	ProductId *string `json:"productId" yaml:"productId"`
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
	// The description of the constraint.
	Description *string `json:"description" yaml:"description"`
	// You are required to specify either the `RoleArn` or the `LocalRoleName` but can't use both.
	//
	// If you specify the `LocalRoleName` property, when an account uses the launch constraint, the IAM role with that name in the account will be used. This allows launch-role constraints to be account-agnostic so the administrator can create fewer resources per shared account.
	//
	// The given role name must exist in the account used to create the launch constraint and the account of the user who launches a product with this launch constraint.
	LocalRoleName *string `json:"localRoleName" yaml:"localRoleName"`
	// The ARN of the launch role.
	//
	// You are required to specify `RoleArn` or `LocalRoleName` but can't use both.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
}

// A CloudFormation `AWS::ServiceCatalog::LaunchTemplateConstraint`.
//
// Specifies a template constraint.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnLaunchTemplateConstraint := servicecatalog.NewCfnLaunchTemplateConstraint(this, jsii.String("MyCfnLaunchTemplateConstraint"), &cfnLaunchTemplateConstraintProps{
//   	portfolioId: jsii.String("portfolioId"),
//   	productId: jsii.String("productId"),
//   	rules: jsii.String("rules"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	description: jsii.String("description"),
//   })
//
type CfnLaunchTemplateConstraint interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the constraint.
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
	// The tree node.
	Node() constructs.Node
	// The portfolio identifier.
	PortfolioId() *string
	SetPortfolioId(val *string)
	// The product identifier.
	ProductId() *string
	SetProductId(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The constraint rules.
	Rules() *string
	SetRules(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnLaunchTemplateConstraint
type jsiiProxy_CfnLaunchTemplateConstraint struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLaunchTemplateConstraint) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchTemplateConstraint) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchTemplateConstraint) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchTemplateConstraint) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchTemplateConstraint) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchTemplateConstraint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchTemplateConstraint) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchTemplateConstraint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchTemplateConstraint) PortfolioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchTemplateConstraint) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchTemplateConstraint) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchTemplateConstraint) Rules() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchTemplateConstraint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchTemplateConstraint) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ServiceCatalog::LaunchTemplateConstraint`.
func NewCfnLaunchTemplateConstraint(scope constructs.Construct, id *string, props *CfnLaunchTemplateConstraintProps) CfnLaunchTemplateConstraint {
	_init_.Initialize()

	j := jsiiProxy_CfnLaunchTemplateConstraint{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnLaunchTemplateConstraint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ServiceCatalog::LaunchTemplateConstraint`.
func NewCfnLaunchTemplateConstraint_Override(c CfnLaunchTemplateConstraint, scope constructs.Construct, id *string, props *CfnLaunchTemplateConstraintProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnLaunchTemplateConstraint",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLaunchTemplateConstraint) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchTemplateConstraint) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchTemplateConstraint) SetPortfolioId(val *string) {
	_jsii_.Set(
		j,
		"portfolioId",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchTemplateConstraint) SetProductId(val *string) {
	_jsii_.Set(
		j,
		"productId",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchTemplateConstraint) SetRules(val *string) {
	_jsii_.Set(
		j,
		"rules",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLaunchTemplateConstraint_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnLaunchTemplateConstraint",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLaunchTemplateConstraint_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnLaunchTemplateConstraint",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnLaunchTemplateConstraint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnLaunchTemplateConstraint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLaunchTemplateConstraint_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_servicecatalog.CfnLaunchTemplateConstraint",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLaunchTemplateConstraint) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnLaunchTemplateConstraint) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLaunchTemplateConstraint) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnLaunchTemplateConstraint) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnLaunchTemplateConstraint) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnLaunchTemplateConstraint) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnLaunchTemplateConstraint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnLaunchTemplateConstraint) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunchTemplateConstraint) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunchTemplateConstraint) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnLaunchTemplateConstraint) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLaunchTemplateConstraint) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunchTemplateConstraint) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunchTemplateConstraint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunchTemplateConstraint) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnLaunchTemplateConstraint`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnLaunchTemplateConstraintProps := &cfnLaunchTemplateConstraintProps{
//   	portfolioId: jsii.String("portfolioId"),
//   	productId: jsii.String("productId"),
//   	rules: jsii.String("rules"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	description: jsii.String("description"),
//   }
//
type CfnLaunchTemplateConstraintProps struct {
	// The portfolio identifier.
	PortfolioId *string `json:"portfolioId" yaml:"portfolioId"`
	// The product identifier.
	ProductId *string `json:"productId" yaml:"productId"`
	// The constraint rules.
	Rules *string `json:"rules" yaml:"rules"`
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
	// The description of the constraint.
	Description *string `json:"description" yaml:"description"`
}

// A CloudFormation `AWS::ServiceCatalog::Portfolio`.
//
// Specifies a portfolio.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnPortfolio := servicecatalog.NewCfnPortfolio(this, jsii.String("MyCfnPortfolio"), &cfnPortfolioProps{
//   	displayName: jsii.String("displayName"),
//   	providerName: jsii.String("providerName"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnPortfolio interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	// The name of the portfolio.
	AttrPortfolioName() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the portfolio.
	Description() *string
	SetDescription(val *string)
	// The name to use for display purposes.
	DisplayName() *string
	SetDisplayName(val *string)
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
	// The name of the portfolio provider.
	ProviderName() *string
	SetProviderName(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// One or more tags.
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnPortfolio
type jsiiProxy_CfnPortfolio struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPortfolio) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolio) AttrPortfolioName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPortfolioName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolio) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolio) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolio) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolio) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolio) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolio) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolio) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolio) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolio) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolio) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolio) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolio) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolio) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ServiceCatalog::Portfolio`.
func NewCfnPortfolio(scope constructs.Construct, id *string, props *CfnPortfolioProps) CfnPortfolio {
	_init_.Initialize()

	j := jsiiProxy_CfnPortfolio{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnPortfolio",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ServiceCatalog::Portfolio`.
func NewCfnPortfolio_Override(c CfnPortfolio, scope constructs.Construct, id *string, props *CfnPortfolioProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnPortfolio",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPortfolio) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_CfnPortfolio) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnPortfolio) SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_CfnPortfolio) SetProviderName(val *string) {
	_jsii_.Set(
		j,
		"providerName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPortfolio_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnPortfolio",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPortfolio_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnPortfolio",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnPortfolio_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnPortfolio",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPortfolio_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_servicecatalog.CfnPortfolio",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPortfolio) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPortfolio) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPortfolio) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPortfolio) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPortfolio) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPortfolio) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPortfolio) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPortfolio) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortfolio) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortfolio) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPortfolio) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPortfolio) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortfolio) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortfolio) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortfolio) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A CloudFormation `AWS::ServiceCatalog::PortfolioPrincipalAssociation`.
//
// Associates the specified principal ARN with the specified portfolio.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnPortfolioPrincipalAssociation := servicecatalog.NewCfnPortfolioPrincipalAssociation(this, jsii.String("MyCfnPortfolioPrincipalAssociation"), &cfnPortfolioPrincipalAssociationProps{
//   	portfolioId: jsii.String("portfolioId"),
//   	principalArn: jsii.String("principalArn"),
//   	principalType: jsii.String("principalType"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   })
//
type CfnPortfolioPrincipalAssociation interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// The portfolio identifier.
	PortfolioId() *string
	SetPortfolioId(val *string)
	// The ARN of the principal (IAM user, role, or group).
	PrincipalArn() *string
	SetPrincipalArn(val *string)
	// The principal type.
	//
	// The supported value is `IAM` .
	PrincipalType() *string
	SetPrincipalType(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnPortfolioPrincipalAssociation
type jsiiProxy_CfnPortfolioPrincipalAssociation struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPortfolioPrincipalAssociation) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioPrincipalAssociation) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioPrincipalAssociation) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioPrincipalAssociation) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioPrincipalAssociation) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioPrincipalAssociation) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioPrincipalAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioPrincipalAssociation) PortfolioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioPrincipalAssociation) PrincipalArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioPrincipalAssociation) PrincipalType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioPrincipalAssociation) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioPrincipalAssociation) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioPrincipalAssociation) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ServiceCatalog::PortfolioPrincipalAssociation`.
func NewCfnPortfolioPrincipalAssociation(scope constructs.Construct, id *string, props *CfnPortfolioPrincipalAssociationProps) CfnPortfolioPrincipalAssociation {
	_init_.Initialize()

	j := jsiiProxy_CfnPortfolioPrincipalAssociation{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnPortfolioPrincipalAssociation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ServiceCatalog::PortfolioPrincipalAssociation`.
func NewCfnPortfolioPrincipalAssociation_Override(c CfnPortfolioPrincipalAssociation, scope constructs.Construct, id *string, props *CfnPortfolioPrincipalAssociationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnPortfolioPrincipalAssociation",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPortfolioPrincipalAssociation) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_CfnPortfolioPrincipalAssociation) SetPortfolioId(val *string) {
	_jsii_.Set(
		j,
		"portfolioId",
		val,
	)
}

func (j *jsiiProxy_CfnPortfolioPrincipalAssociation) SetPrincipalArn(val *string) {
	_jsii_.Set(
		j,
		"principalArn",
		val,
	)
}

func (j *jsiiProxy_CfnPortfolioPrincipalAssociation) SetPrincipalType(val *string) {
	_jsii_.Set(
		j,
		"principalType",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPortfolioPrincipalAssociation_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnPortfolioPrincipalAssociation",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPortfolioPrincipalAssociation_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnPortfolioPrincipalAssociation",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnPortfolioPrincipalAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnPortfolioPrincipalAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPortfolioPrincipalAssociation_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_servicecatalog.CfnPortfolioPrincipalAssociation",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPortfolioPrincipalAssociation) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPortfolioPrincipalAssociation) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPortfolioPrincipalAssociation) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPortfolioPrincipalAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPortfolioPrincipalAssociation) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPortfolioPrincipalAssociation) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPortfolioPrincipalAssociation) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPortfolioPrincipalAssociation) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortfolioPrincipalAssociation) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortfolioPrincipalAssociation) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPortfolioPrincipalAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPortfolioPrincipalAssociation) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortfolioPrincipalAssociation) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortfolioPrincipalAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortfolioPrincipalAssociation) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnPortfolioPrincipalAssociation`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnPortfolioPrincipalAssociationProps := &cfnPortfolioPrincipalAssociationProps{
//   	portfolioId: jsii.String("portfolioId"),
//   	principalArn: jsii.String("principalArn"),
//   	principalType: jsii.String("principalType"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   }
//
type CfnPortfolioPrincipalAssociationProps struct {
	// The portfolio identifier.
	PortfolioId *string `json:"portfolioId" yaml:"portfolioId"`
	// The ARN of the principal (IAM user, role, or group).
	PrincipalArn *string `json:"principalArn" yaml:"principalArn"`
	// The principal type.
	//
	// The supported value is `IAM` .
	PrincipalType *string `json:"principalType" yaml:"principalType"`
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
}

// A CloudFormation `AWS::ServiceCatalog::PortfolioProductAssociation`.
//
// Associates the specified product with the specified portfolio.
//
// A delegated admin is authorized to invoke this command.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnPortfolioProductAssociation := servicecatalog.NewCfnPortfolioProductAssociation(this, jsii.String("MyCfnPortfolioProductAssociation"), &cfnPortfolioProductAssociationProps{
//   	portfolioId: jsii.String("portfolioId"),
//   	productId: jsii.String("productId"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	sourcePortfolioId: jsii.String("sourcePortfolioId"),
//   })
//
type CfnPortfolioProductAssociation interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// The portfolio identifier.
	PortfolioId() *string
	SetPortfolioId(val *string)
	// The product identifier.
	ProductId() *string
	SetProductId(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The identifier of the source portfolio.
	SourcePortfolioId() *string
	SetSourcePortfolioId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnPortfolioProductAssociation
type jsiiProxy_CfnPortfolioProductAssociation struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPortfolioProductAssociation) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioProductAssociation) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioProductAssociation) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioProductAssociation) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioProductAssociation) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioProductAssociation) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioProductAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioProductAssociation) PortfolioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioProductAssociation) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioProductAssociation) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioProductAssociation) SourcePortfolioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePortfolioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioProductAssociation) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioProductAssociation) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ServiceCatalog::PortfolioProductAssociation`.
func NewCfnPortfolioProductAssociation(scope constructs.Construct, id *string, props *CfnPortfolioProductAssociationProps) CfnPortfolioProductAssociation {
	_init_.Initialize()

	j := jsiiProxy_CfnPortfolioProductAssociation{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnPortfolioProductAssociation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ServiceCatalog::PortfolioProductAssociation`.
func NewCfnPortfolioProductAssociation_Override(c CfnPortfolioProductAssociation, scope constructs.Construct, id *string, props *CfnPortfolioProductAssociationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnPortfolioProductAssociation",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPortfolioProductAssociation) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_CfnPortfolioProductAssociation) SetPortfolioId(val *string) {
	_jsii_.Set(
		j,
		"portfolioId",
		val,
	)
}

func (j *jsiiProxy_CfnPortfolioProductAssociation) SetProductId(val *string) {
	_jsii_.Set(
		j,
		"productId",
		val,
	)
}

func (j *jsiiProxy_CfnPortfolioProductAssociation) SetSourcePortfolioId(val *string) {
	_jsii_.Set(
		j,
		"sourcePortfolioId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPortfolioProductAssociation_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnPortfolioProductAssociation",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPortfolioProductAssociation_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnPortfolioProductAssociation",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnPortfolioProductAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnPortfolioProductAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPortfolioProductAssociation_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_servicecatalog.CfnPortfolioProductAssociation",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPortfolioProductAssociation) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPortfolioProductAssociation) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPortfolioProductAssociation) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPortfolioProductAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPortfolioProductAssociation) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPortfolioProductAssociation) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPortfolioProductAssociation) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPortfolioProductAssociation) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortfolioProductAssociation) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortfolioProductAssociation) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPortfolioProductAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPortfolioProductAssociation) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortfolioProductAssociation) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortfolioProductAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortfolioProductAssociation) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnPortfolioProductAssociation`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnPortfolioProductAssociationProps := &cfnPortfolioProductAssociationProps{
//   	portfolioId: jsii.String("portfolioId"),
//   	productId: jsii.String("productId"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	sourcePortfolioId: jsii.String("sourcePortfolioId"),
//   }
//
type CfnPortfolioProductAssociationProps struct {
	// The portfolio identifier.
	PortfolioId *string `json:"portfolioId" yaml:"portfolioId"`
	// The product identifier.
	ProductId *string `json:"productId" yaml:"productId"`
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
	// The identifier of the source portfolio.
	SourcePortfolioId *string `json:"sourcePortfolioId" yaml:"sourcePortfolioId"`
}

// Properties for defining a `CfnPortfolio`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnPortfolioProps := &cfnPortfolioProps{
//   	displayName: jsii.String("displayName"),
//   	providerName: jsii.String("providerName"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPortfolioProps struct {
	// The name to use for display purposes.
	DisplayName *string `json:"displayName" yaml:"displayName"`
	// The name of the portfolio provider.
	ProviderName *string `json:"providerName" yaml:"providerName"`
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
	// The description of the portfolio.
	Description *string `json:"description" yaml:"description"`
	// One or more tags.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::ServiceCatalog::PortfolioShare`.
//
// Shares the specified portfolio with the specified account.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnPortfolioShare := servicecatalog.NewCfnPortfolioShare(this, jsii.String("MyCfnPortfolioShare"), &cfnPortfolioShareProps{
//   	accountId: jsii.String("accountId"),
//   	portfolioId: jsii.String("portfolioId"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	shareTagOptions: jsii.Boolean(false),
//   })
//
type CfnPortfolioShare interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	// The AWS account ID.
	//
	// For example, `123456789012` .
	AccountId() *string
	SetAccountId(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// The portfolio identifier.
	PortfolioId() *string
	SetPortfolioId(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Indicates whether TagOptions sharing is enabled or disabled for the portfolio share.
	ShareTagOptions() interface{}
	SetShareTagOptions(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnPortfolioShare
type jsiiProxy_CfnPortfolioShare struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPortfolioShare) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioShare) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioShare) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioShare) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioShare) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioShare) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioShare) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioShare) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioShare) PortfolioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioShare) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioShare) ShareTagOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shareTagOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioShare) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortfolioShare) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ServiceCatalog::PortfolioShare`.
func NewCfnPortfolioShare(scope constructs.Construct, id *string, props *CfnPortfolioShareProps) CfnPortfolioShare {
	_init_.Initialize()

	j := jsiiProxy_CfnPortfolioShare{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnPortfolioShare",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ServiceCatalog::PortfolioShare`.
func NewCfnPortfolioShare_Override(c CfnPortfolioShare, scope constructs.Construct, id *string, props *CfnPortfolioShareProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnPortfolioShare",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPortfolioShare) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_CfnPortfolioShare) SetAccountId(val *string) {
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_CfnPortfolioShare) SetPortfolioId(val *string) {
	_jsii_.Set(
		j,
		"portfolioId",
		val,
	)
}

func (j *jsiiProxy_CfnPortfolioShare) SetShareTagOptions(val interface{}) {
	_jsii_.Set(
		j,
		"shareTagOptions",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPortfolioShare_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnPortfolioShare",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPortfolioShare_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnPortfolioShare",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnPortfolioShare_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnPortfolioShare",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPortfolioShare_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_servicecatalog.CfnPortfolioShare",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPortfolioShare) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPortfolioShare) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPortfolioShare) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPortfolioShare) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPortfolioShare) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPortfolioShare) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPortfolioShare) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPortfolioShare) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortfolioShare) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortfolioShare) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPortfolioShare) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPortfolioShare) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortfolioShare) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortfolioShare) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPortfolioShare) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnPortfolioShare`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnPortfolioShareProps := &cfnPortfolioShareProps{
//   	accountId: jsii.String("accountId"),
//   	portfolioId: jsii.String("portfolioId"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	shareTagOptions: jsii.Boolean(false),
//   }
//
type CfnPortfolioShareProps struct {
	// The AWS account ID.
	//
	// For example, `123456789012` .
	AccountId *string `json:"accountId" yaml:"accountId"`
	// The portfolio identifier.
	PortfolioId *string `json:"portfolioId" yaml:"portfolioId"`
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
	// Indicates whether TagOptions sharing is enabled or disabled for the portfolio share.
	ShareTagOptions interface{} `json:"shareTagOptions" yaml:"shareTagOptions"`
}

// A CloudFormation `AWS::ServiceCatalog::ResourceUpdateConstraint`.
//
// Specifies a `RESOURCE_UPDATE` constraint.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnResourceUpdateConstraint := servicecatalog.NewCfnResourceUpdateConstraint(this, jsii.String("MyCfnResourceUpdateConstraint"), &cfnResourceUpdateConstraintProps{
//   	portfolioId: jsii.String("portfolioId"),
//   	productId: jsii.String("productId"),
//   	tagUpdateOnProvisionedProduct: jsii.String("tagUpdateOnProvisionedProduct"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	description: jsii.String("description"),
//   })
//
type CfnResourceUpdateConstraint interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the constraint.
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
	// The tree node.
	Node() constructs.Node
	// The portfolio identifier.
	PortfolioId() *string
	SetPortfolioId(val *string)
	// The product identifier.
	ProductId() *string
	SetProductId(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// If set to `ALLOWED` , lets users change tags in a [CloudFormationProvisionedProduct](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html) resource.
	//
	// If set to `NOT_ALLOWED` , prevents users from changing tags in a [CloudFormationProvisionedProduct](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html) resource.
	TagUpdateOnProvisionedProduct() *string
	SetTagUpdateOnProvisionedProduct(val *string)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnResourceUpdateConstraint
type jsiiProxy_CfnResourceUpdateConstraint struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnResourceUpdateConstraint) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceUpdateConstraint) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceUpdateConstraint) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceUpdateConstraint) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceUpdateConstraint) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceUpdateConstraint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceUpdateConstraint) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceUpdateConstraint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceUpdateConstraint) PortfolioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceUpdateConstraint) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceUpdateConstraint) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceUpdateConstraint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceUpdateConstraint) TagUpdateOnProvisionedProduct() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagUpdateOnProvisionedProduct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceUpdateConstraint) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ServiceCatalog::ResourceUpdateConstraint`.
func NewCfnResourceUpdateConstraint(scope constructs.Construct, id *string, props *CfnResourceUpdateConstraintProps) CfnResourceUpdateConstraint {
	_init_.Initialize()

	j := jsiiProxy_CfnResourceUpdateConstraint{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnResourceUpdateConstraint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ServiceCatalog::ResourceUpdateConstraint`.
func NewCfnResourceUpdateConstraint_Override(c CfnResourceUpdateConstraint, scope constructs.Construct, id *string, props *CfnResourceUpdateConstraintProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnResourceUpdateConstraint",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResourceUpdateConstraint) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_CfnResourceUpdateConstraint) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnResourceUpdateConstraint) SetPortfolioId(val *string) {
	_jsii_.Set(
		j,
		"portfolioId",
		val,
	)
}

func (j *jsiiProxy_CfnResourceUpdateConstraint) SetProductId(val *string) {
	_jsii_.Set(
		j,
		"productId",
		val,
	)
}

func (j *jsiiProxy_CfnResourceUpdateConstraint) SetTagUpdateOnProvisionedProduct(val *string) {
	_jsii_.Set(
		j,
		"tagUpdateOnProvisionedProduct",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnResourceUpdateConstraint_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnResourceUpdateConstraint",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnResourceUpdateConstraint_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnResourceUpdateConstraint",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnResourceUpdateConstraint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnResourceUpdateConstraint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResourceUpdateConstraint_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_servicecatalog.CfnResourceUpdateConstraint",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResourceUpdateConstraint) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnResourceUpdateConstraint) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnResourceUpdateConstraint) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnResourceUpdateConstraint) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnResourceUpdateConstraint) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnResourceUpdateConstraint) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnResourceUpdateConstraint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnResourceUpdateConstraint) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceUpdateConstraint) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceUpdateConstraint) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnResourceUpdateConstraint) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResourceUpdateConstraint) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceUpdateConstraint) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceUpdateConstraint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourceUpdateConstraint) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnResourceUpdateConstraint`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnResourceUpdateConstraintProps := &cfnResourceUpdateConstraintProps{
//   	portfolioId: jsii.String("portfolioId"),
//   	productId: jsii.String("productId"),
//   	tagUpdateOnProvisionedProduct: jsii.String("tagUpdateOnProvisionedProduct"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	description: jsii.String("description"),
//   }
//
type CfnResourceUpdateConstraintProps struct {
	// The portfolio identifier.
	PortfolioId *string `json:"portfolioId" yaml:"portfolioId"`
	// The product identifier.
	ProductId *string `json:"productId" yaml:"productId"`
	// If set to `ALLOWED` , lets users change tags in a [CloudFormationProvisionedProduct](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html) resource.
	//
	// If set to `NOT_ALLOWED` , prevents users from changing tags in a [CloudFormationProvisionedProduct](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html) resource.
	TagUpdateOnProvisionedProduct *string `json:"tagUpdateOnProvisionedProduct" yaml:"tagUpdateOnProvisionedProduct"`
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
	// The description of the constraint.
	Description *string `json:"description" yaml:"description"`
}

// A CloudFormation `AWS::ServiceCatalog::ServiceAction`.
//
// Creates a self-service action.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnServiceAction := servicecatalog.NewCfnServiceAction(this, jsii.String("MyCfnServiceAction"), &cfnServiceActionProps{
//   	definition: []interface{}{
//   		&definitionParameterProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	definitionType: jsii.String("definitionType"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	description: jsii.String("description"),
//   })
//
type CfnServiceAction interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	// The self-service action identifier.
	//
	// For example, `act-fs7abcd89wxyz` .
	AttrId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A map that defines the self-service action.
	Definition() interface{}
	SetDefinition(val interface{})
	// The self-service action definition type.
	//
	// For example, `SSM_AUTOMATION` .
	DefinitionType() *string
	SetDefinitionType(val *string)
	// The self-service action description.
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
	// The self-service action name.
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnServiceAction
type jsiiProxy_CfnServiceAction struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnServiceAction) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceAction) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceAction) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceAction) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceAction) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceAction) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceAction) Definition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceAction) DefinitionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definitionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceAction) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceAction) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceAction) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceAction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceAction) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceAction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceAction) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ServiceCatalog::ServiceAction`.
func NewCfnServiceAction(scope constructs.Construct, id *string, props *CfnServiceActionProps) CfnServiceAction {
	_init_.Initialize()

	j := jsiiProxy_CfnServiceAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnServiceAction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ServiceCatalog::ServiceAction`.
func NewCfnServiceAction_Override(c CfnServiceAction, scope constructs.Construct, id *string, props *CfnServiceActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnServiceAction",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnServiceAction) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_CfnServiceAction) SetDefinition(val interface{}) {
	_jsii_.Set(
		j,
		"definition",
		val,
	)
}

func (j *jsiiProxy_CfnServiceAction) SetDefinitionType(val *string) {
	_jsii_.Set(
		j,
		"definitionType",
		val,
	)
}

func (j *jsiiProxy_CfnServiceAction) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnServiceAction) SetName(val *string) {
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
func CfnServiceAction_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnServiceAction",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnServiceAction_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnServiceAction",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnServiceAction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnServiceAction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnServiceAction_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_servicecatalog.CfnServiceAction",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnServiceAction) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnServiceAction) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnServiceAction) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnServiceAction) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnServiceAction) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnServiceAction) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnServiceAction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnServiceAction) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServiceAction) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServiceAction) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnServiceAction) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnServiceAction) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServiceAction) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServiceAction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServiceAction) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The list of parameters in JSON format.
//
// For example: `[{\"Name\":\"InstanceId\",\"Type\":\"TARGET\"}] or [{\"Name\":\"InstanceId\",\"Type\":\"TEXT_VALUE\"}]` .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   definitionParameterProperty := &definitionParameterProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnServiceAction_DefinitionParameterProperty struct {
	// The parameter key.
	Key *string `json:"key" yaml:"key"`
	// The value of the parameter.
	Value *string `json:"value" yaml:"value"`
}

// A CloudFormation `AWS::ServiceCatalog::ServiceActionAssociation`.
//
// A self-service action association consisting of the Action ID, the Product ID, and the Provisioning Artifact ID.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnServiceActionAssociation := servicecatalog.NewCfnServiceActionAssociation(this, jsii.String("MyCfnServiceActionAssociation"), &cfnServiceActionAssociationProps{
//   	productId: jsii.String("productId"),
//   	provisioningArtifactId: jsii.String("provisioningArtifactId"),
//   	serviceActionId: jsii.String("serviceActionId"),
//   })
//
type CfnServiceActionAssociation interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// The product identifier.
	//
	// For example, `prod-abcdzk7xy33qa` .
	ProductId() *string
	SetProductId(val *string)
	// The identifier of the provisioning artifact.
	//
	// For example, `pa-4abcdjnxjj6ne` .
	ProvisioningArtifactId() *string
	SetProvisioningArtifactId(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The self-service action identifier.
	//
	// For example, `act-fs7abcd89wxyz` .
	ServiceActionId() *string
	SetServiceActionId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnServiceActionAssociation
type jsiiProxy_CfnServiceActionAssociation struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnServiceActionAssociation) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceActionAssociation) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceActionAssociation) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceActionAssociation) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceActionAssociation) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceActionAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceActionAssociation) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceActionAssociation) ProvisioningArtifactId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningArtifactId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceActionAssociation) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceActionAssociation) ServiceActionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceActionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceActionAssociation) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceActionAssociation) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ServiceCatalog::ServiceActionAssociation`.
func NewCfnServiceActionAssociation(scope constructs.Construct, id *string, props *CfnServiceActionAssociationProps) CfnServiceActionAssociation {
	_init_.Initialize()

	j := jsiiProxy_CfnServiceActionAssociation{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnServiceActionAssociation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ServiceCatalog::ServiceActionAssociation`.
func NewCfnServiceActionAssociation_Override(c CfnServiceActionAssociation, scope constructs.Construct, id *string, props *CfnServiceActionAssociationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnServiceActionAssociation",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnServiceActionAssociation) SetProductId(val *string) {
	_jsii_.Set(
		j,
		"productId",
		val,
	)
}

func (j *jsiiProxy_CfnServiceActionAssociation) SetProvisioningArtifactId(val *string) {
	_jsii_.Set(
		j,
		"provisioningArtifactId",
		val,
	)
}

func (j *jsiiProxy_CfnServiceActionAssociation) SetServiceActionId(val *string) {
	_jsii_.Set(
		j,
		"serviceActionId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnServiceActionAssociation_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnServiceActionAssociation",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnServiceActionAssociation_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnServiceActionAssociation",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnServiceActionAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnServiceActionAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnServiceActionAssociation_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_servicecatalog.CfnServiceActionAssociation",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnServiceActionAssociation) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnServiceActionAssociation) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnServiceActionAssociation) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnServiceActionAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnServiceActionAssociation) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnServiceActionAssociation) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnServiceActionAssociation) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnServiceActionAssociation) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServiceActionAssociation) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServiceActionAssociation) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnServiceActionAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnServiceActionAssociation) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServiceActionAssociation) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServiceActionAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServiceActionAssociation) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnServiceActionAssociation`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnServiceActionAssociationProps := &cfnServiceActionAssociationProps{
//   	productId: jsii.String("productId"),
//   	provisioningArtifactId: jsii.String("provisioningArtifactId"),
//   	serviceActionId: jsii.String("serviceActionId"),
//   }
//
type CfnServiceActionAssociationProps struct {
	// The product identifier.
	//
	// For example, `prod-abcdzk7xy33qa` .
	ProductId *string `json:"productId" yaml:"productId"`
	// The identifier of the provisioning artifact.
	//
	// For example, `pa-4abcdjnxjj6ne` .
	ProvisioningArtifactId *string `json:"provisioningArtifactId" yaml:"provisioningArtifactId"`
	// The self-service action identifier.
	//
	// For example, `act-fs7abcd89wxyz` .
	ServiceActionId *string `json:"serviceActionId" yaml:"serviceActionId"`
}

// Properties for defining a `CfnServiceAction`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnServiceActionProps := &cfnServiceActionProps{
//   	definition: []interface{}{
//   		&definitionParameterProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	definitionType: jsii.String("definitionType"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	description: jsii.String("description"),
//   }
//
type CfnServiceActionProps struct {
	// A map that defines the self-service action.
	Definition interface{} `json:"definition" yaml:"definition"`
	// The self-service action definition type.
	//
	// For example, `SSM_AUTOMATION` .
	DefinitionType *string `json:"definitionType" yaml:"definitionType"`
	// The self-service action name.
	Name *string `json:"name" yaml:"name"`
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
	// The self-service action description.
	Description *string `json:"description" yaml:"description"`
}

// A CloudFormation `AWS::ServiceCatalog::StackSetConstraint`.
//
// Specifies a StackSet constraint.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnStackSetConstraint := servicecatalog.NewCfnStackSetConstraint(this, jsii.String("MyCfnStackSetConstraint"), &cfnStackSetConstraintProps{
//   	accountList: []*string{
//   		jsii.String("accountList"),
//   	},
//   	adminRole: jsii.String("adminRole"),
//   	description: jsii.String("description"),
//   	executionRole: jsii.String("executionRole"),
//   	portfolioId: jsii.String("portfolioId"),
//   	productId: jsii.String("productId"),
//   	regionList: []*string{
//   		jsii.String("regionList"),
//   	},
//   	stackInstanceControl: jsii.String("stackInstanceControl"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   })
//
type CfnStackSetConstraint interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage() *string
	SetAcceptLanguage(val *string)
	// One or more AWS accounts that will have access to the provisioned product.
	AccountList() *[]*string
	SetAccountList(val *[]*string)
	// AdminRole ARN.
	AdminRole() *string
	SetAdminRole(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the constraint.
	Description() *string
	SetDescription(val *string)
	// ExecutionRole name.
	ExecutionRole() *string
	SetExecutionRole(val *string)
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
	// The portfolio identifier.
	PortfolioId() *string
	SetPortfolioId(val *string)
	// The product identifier.
	ProductId() *string
	SetProductId(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// One or more AWS Regions where the provisioned product will be available.
	//
	// Applicable only to a `CFN_STACKSET` provisioned product type.
	//
	// The specified Regions should be within the list of Regions from the `STACKSET` constraint. To get the list of Regions in the `STACKSET` constraint, use the `DescribeProvisioningParameters` operation.
	//
	// If no values are specified, the default value is all Regions from the `STACKSET` constraint.
	RegionList() *[]*string
	SetRegionList(val *[]*string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Permission to create, update, and delete stack instances.
	//
	// Choose from ALLOWED and NOT_ALLOWED.
	StackInstanceControl() *string
	SetStackInstanceControl(val *string)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnStackSetConstraint
type jsiiProxy_CfnStackSetConstraint struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnStackSetConstraint) AcceptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSetConstraint) AccountList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSetConstraint) AdminRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSetConstraint) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSetConstraint) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSetConstraint) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSetConstraint) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSetConstraint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSetConstraint) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSetConstraint) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSetConstraint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSetConstraint) PortfolioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSetConstraint) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSetConstraint) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSetConstraint) RegionList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSetConstraint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSetConstraint) StackInstanceControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackInstanceControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSetConstraint) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ServiceCatalog::StackSetConstraint`.
func NewCfnStackSetConstraint(scope constructs.Construct, id *string, props *CfnStackSetConstraintProps) CfnStackSetConstraint {
	_init_.Initialize()

	j := jsiiProxy_CfnStackSetConstraint{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnStackSetConstraint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ServiceCatalog::StackSetConstraint`.
func NewCfnStackSetConstraint_Override(c CfnStackSetConstraint, scope constructs.Construct, id *string, props *CfnStackSetConstraintProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnStackSetConstraint",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStackSetConstraint) SetAcceptLanguage(val *string) {
	_jsii_.Set(
		j,
		"acceptLanguage",
		val,
	)
}

func (j *jsiiProxy_CfnStackSetConstraint) SetAccountList(val *[]*string) {
	_jsii_.Set(
		j,
		"accountList",
		val,
	)
}

func (j *jsiiProxy_CfnStackSetConstraint) SetAdminRole(val *string) {
	_jsii_.Set(
		j,
		"adminRole",
		val,
	)
}

func (j *jsiiProxy_CfnStackSetConstraint) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnStackSetConstraint) SetExecutionRole(val *string) {
	_jsii_.Set(
		j,
		"executionRole",
		val,
	)
}

func (j *jsiiProxy_CfnStackSetConstraint) SetPortfolioId(val *string) {
	_jsii_.Set(
		j,
		"portfolioId",
		val,
	)
}

func (j *jsiiProxy_CfnStackSetConstraint) SetProductId(val *string) {
	_jsii_.Set(
		j,
		"productId",
		val,
	)
}

func (j *jsiiProxy_CfnStackSetConstraint) SetRegionList(val *[]*string) {
	_jsii_.Set(
		j,
		"regionList",
		val,
	)
}

func (j *jsiiProxy_CfnStackSetConstraint) SetStackInstanceControl(val *string) {
	_jsii_.Set(
		j,
		"stackInstanceControl",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnStackSetConstraint_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnStackSetConstraint",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnStackSetConstraint_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnStackSetConstraint",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnStackSetConstraint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnStackSetConstraint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStackSetConstraint_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_servicecatalog.CfnStackSetConstraint",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStackSetConstraint) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnStackSetConstraint) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnStackSetConstraint) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnStackSetConstraint) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnStackSetConstraint) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnStackSetConstraint) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnStackSetConstraint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnStackSetConstraint) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStackSetConstraint) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStackSetConstraint) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnStackSetConstraint) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStackSetConstraint) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStackSetConstraint) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStackSetConstraint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStackSetConstraint) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnStackSetConstraint`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnStackSetConstraintProps := &cfnStackSetConstraintProps{
//   	accountList: []*string{
//   		jsii.String("accountList"),
//   	},
//   	adminRole: jsii.String("adminRole"),
//   	description: jsii.String("description"),
//   	executionRole: jsii.String("executionRole"),
//   	portfolioId: jsii.String("portfolioId"),
//   	productId: jsii.String("productId"),
//   	regionList: []*string{
//   		jsii.String("regionList"),
//   	},
//   	stackInstanceControl: jsii.String("stackInstanceControl"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   }
//
type CfnStackSetConstraintProps struct {
	// One or more AWS accounts that will have access to the provisioned product.
	AccountList *[]*string `json:"accountList" yaml:"accountList"`
	// AdminRole ARN.
	AdminRole *string `json:"adminRole" yaml:"adminRole"`
	// The description of the constraint.
	Description *string `json:"description" yaml:"description"`
	// ExecutionRole name.
	ExecutionRole *string `json:"executionRole" yaml:"executionRole"`
	// The portfolio identifier.
	PortfolioId *string `json:"portfolioId" yaml:"portfolioId"`
	// The product identifier.
	ProductId *string `json:"productId" yaml:"productId"`
	// One or more AWS Regions where the provisioned product will be available.
	//
	// Applicable only to a `CFN_STACKSET` provisioned product type.
	//
	// The specified Regions should be within the list of Regions from the `STACKSET` constraint. To get the list of Regions in the `STACKSET` constraint, use the `DescribeProvisioningParameters` operation.
	//
	// If no values are specified, the default value is all Regions from the `STACKSET` constraint.
	RegionList *[]*string `json:"regionList" yaml:"regionList"`
	// Permission to create, update, and delete stack instances.
	//
	// Choose from ALLOWED and NOT_ALLOWED.
	StackInstanceControl *string `json:"stackInstanceControl" yaml:"stackInstanceControl"`
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `json:"acceptLanguage" yaml:"acceptLanguage"`
}

// A CloudFormation `AWS::ServiceCatalog::TagOption`.
//
// Specifies a TagOption. A TagOption is a key-value pair managed by AWS Service Catalog that serves as a template for creating an AWS tag.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnTagOption := servicecatalog.NewCfnTagOption(this, jsii.String("MyCfnTagOption"), &cfnTagOptionProps{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//
//   	// the properties below are optional
//   	active: jsii.Boolean(false),
//   })
//
type CfnTagOption interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The TagOption active state.
	Active() interface{}
	SetActive(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The TagOption key.
	Key() *string
	SetKey(val *string)
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
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// The TagOption value.
	Value() *string
	SetValue(val *string)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnTagOption
type jsiiProxy_CfnTagOption struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTagOption) Active() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"active",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagOption) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagOption) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagOption) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagOption) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagOption) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagOption) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagOption) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagOption) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagOption) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagOption) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagOption) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Create a new `AWS::ServiceCatalog::TagOption`.
func NewCfnTagOption(scope constructs.Construct, id *string, props *CfnTagOptionProps) CfnTagOption {
	_init_.Initialize()

	j := jsiiProxy_CfnTagOption{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnTagOption",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ServiceCatalog::TagOption`.
func NewCfnTagOption_Override(c CfnTagOption, scope constructs.Construct, id *string, props *CfnTagOptionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnTagOption",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTagOption) SetActive(val interface{}) {
	_jsii_.Set(
		j,
		"active",
		val,
	)
}

func (j *jsiiProxy_CfnTagOption) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_CfnTagOption) SetValue(val *string) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnTagOption_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnTagOption",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnTagOption_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnTagOption",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnTagOption_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnTagOption",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTagOption_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_servicecatalog.CfnTagOption",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTagOption) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTagOption) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTagOption) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTagOption) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTagOption) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTagOption) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTagOption) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTagOption) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTagOption) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTagOption) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTagOption) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTagOption) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTagOption) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTagOption) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTagOption) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A CloudFormation `AWS::ServiceCatalog::TagOptionAssociation`.
//
// Associate the specified TagOption with the specified portfolio or product.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnTagOptionAssociation := servicecatalog.NewCfnTagOptionAssociation(this, jsii.String("MyCfnTagOptionAssociation"), &cfnTagOptionAssociationProps{
//   	resourceId: jsii.String("resourceId"),
//   	tagOptionId: jsii.String("tagOptionId"),
//   })
//
type CfnTagOptionAssociation interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// The resource identifier.
	ResourceId() *string
	SetResourceId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The TagOption identifier.
	TagOptionId() *string
	SetTagOptionId(val *string)
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// The jsii proxy struct for CfnTagOptionAssociation
type jsiiProxy_CfnTagOptionAssociation struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTagOptionAssociation) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagOptionAssociation) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagOptionAssociation) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagOptionAssociation) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagOptionAssociation) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagOptionAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagOptionAssociation) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagOptionAssociation) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagOptionAssociation) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagOptionAssociation) TagOptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagOptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagOptionAssociation) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ServiceCatalog::TagOptionAssociation`.
func NewCfnTagOptionAssociation(scope constructs.Construct, id *string, props *CfnTagOptionAssociationProps) CfnTagOptionAssociation {
	_init_.Initialize()

	j := jsiiProxy_CfnTagOptionAssociation{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnTagOptionAssociation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ServiceCatalog::TagOptionAssociation`.
func NewCfnTagOptionAssociation_Override(c CfnTagOptionAssociation, scope constructs.Construct, id *string, props *CfnTagOptionAssociationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CfnTagOptionAssociation",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTagOptionAssociation) SetResourceId(val *string) {
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_CfnTagOptionAssociation) SetTagOptionId(val *string) {
	_jsii_.Set(
		j,
		"tagOptionId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnTagOptionAssociation_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnTagOptionAssociation",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnTagOptionAssociation_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnTagOptionAssociation",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnTagOptionAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CfnTagOptionAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTagOptionAssociation_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_servicecatalog.CfnTagOptionAssociation",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTagOptionAssociation) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTagOptionAssociation) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTagOptionAssociation) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTagOptionAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTagOptionAssociation) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTagOptionAssociation) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTagOptionAssociation) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTagOptionAssociation) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTagOptionAssociation) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTagOptionAssociation) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTagOptionAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTagOptionAssociation) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTagOptionAssociation) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTagOptionAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTagOptionAssociation) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnTagOptionAssociation`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnTagOptionAssociationProps := &cfnTagOptionAssociationProps{
//   	resourceId: jsii.String("resourceId"),
//   	tagOptionId: jsii.String("tagOptionId"),
//   }
//
type CfnTagOptionAssociationProps struct {
	// The resource identifier.
	ResourceId *string `json:"resourceId" yaml:"resourceId"`
	// The TagOption identifier.
	TagOptionId *string `json:"tagOptionId" yaml:"tagOptionId"`
}

// Properties for defining a `CfnTagOption`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import servicecatalog "github.com/aws/aws-cdk-go/awscdk/aws_servicecatalog"
//   cfnTagOptionProps := &cfnTagOptionProps{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//
//   	// the properties below are optional
//   	active: jsii.Boolean(false),
//   }
//
type CfnTagOptionProps struct {
	// The TagOption key.
	Key *string `json:"key" yaml:"key"`
	// The TagOption value.
	Value *string `json:"value" yaml:"value"`
	// The TagOption active state.
	Active interface{} `json:"active" yaml:"active"`
}


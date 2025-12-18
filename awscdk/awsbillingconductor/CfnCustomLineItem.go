package awsbillingconductor

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbillingconductor/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbillingconductor"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a custom line item that can be used to create a one-time or recurring, fixed or percentage-based charge that you can apply to a single billing group.
//
// You can apply custom line items to the current or previous billing period. You can create either a fee or a discount custom line item.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCustomLineItem := awscdk.Aws_billingconductor.NewCfnCustomLineItem(this, jsii.String("MyCfnCustomLineItem"), &CfnCustomLineItemProps{
//   	BillingGroupArn: jsii.String("billingGroupArn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AccountId: jsii.String("accountId"),
//   	BillingPeriodRange: &BillingPeriodRangeProperty{
//   		ExclusiveEndBillingPeriod: jsii.String("exclusiveEndBillingPeriod"),
//   		InclusiveStartBillingPeriod: jsii.String("inclusiveStartBillingPeriod"),
//   	},
//   	ComputationRule: jsii.String("computationRule"),
//   	CustomLineItemChargeDetails: &CustomLineItemChargeDetailsProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Flat: &CustomLineItemFlatChargeDetailsProperty{
//   			ChargeValue: jsii.Number(123),
//   		},
//   		LineItemFilters: []interface{}{
//   			&LineItemFilterProperty{
//   				Attribute: jsii.String("attribute"),
//   				MatchOption: jsii.String("matchOption"),
//
//   				// the properties below are optional
//   				AttributeValues: []*string{
//   					jsii.String("attributeValues"),
//   				},
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		Percentage: &CustomLineItemPercentageChargeDetailsProperty{
//   			PercentageValue: jsii.Number(123),
//
//   			// the properties below are optional
//   			ChildAssociatedResources: []*string{
//   				jsii.String("childAssociatedResources"),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	PresentationDetails: &PresentationDetailsProperty{
//   		Service: jsii.String("service"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-customlineitem.html
//
type CfnCustomLineItem interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsbillingconductor.ICustomLineItemRef
	awscdk.ITaggable
	// The AWS account in which this custom line item will be applied to.
	AccountId() *string
	SetAccountId(val *string)
	// The Amazon Resource Name (ARN) that references the billing group where the custom line item applies to.
	AttrArn() *string
	// The number of resources that are associated to the custom line item.
	AttrAssociationSize() *float64
	// The time created.
	AttrCreationTime() *float64
	// The custom line item's charge value currency.
	//
	// Only one of the valid values can be used.
	AttrCurrencyCode() *string
	// The most recent time the custom line item was modified.
	AttrLastModifiedTime() *float64
	// The product code associated with the custom line item.
	AttrProductCode() *string
	// The Amazon Resource Name (ARN) that references the billing group where the custom line item applies to.
	BillingGroupArn() *string
	SetBillingGroupArn(val *string)
	// A time range for which the custom line item is effective.
	BillingPeriodRange() interface{}
	SetBillingPeriodRange(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The computation rule that determines how the custom line item charges are computed and reflected in the bill.
	ComputationRule() *string
	SetComputationRule(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The charge details of a custom line item.
	CustomLineItemChargeDetails() interface{}
	SetCustomLineItemChargeDetails(val interface{})
	// A reference to a CustomLineItem resource.
	CustomLineItemRef() *interfacesawsbillingconductor.CustomLineItemReference
	// The custom line item's description.
	Description() *string
	SetDescription(val *string)
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
	// The custom line item's name.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Configuration details specifying how the custom line item charges are presented, including which service the charges are shown under.
	PresentationDetails() interface{}
	SetPresentationDetails(val interface{})
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
	// A map that contains tag keys and tag values that are attached to a custom line item.
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

// The jsii proxy struct for CfnCustomLineItem
type jsiiProxy_CfnCustomLineItem struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsbillingconductorICustomLineItemRef
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnCustomLineItem) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) AttrAssociationSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrAssociationSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) AttrCreationTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) AttrCurrencyCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCurrencyCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) AttrLastModifiedTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrLastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) AttrProductCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrProductCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) BillingGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) BillingPeriodRange() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"billingPeriodRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) ComputationRule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computationRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) CustomLineItemChargeDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customLineItemChargeDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) CustomLineItemRef() *interfacesawsbillingconductor.CustomLineItemReference {
	var returns *interfacesawsbillingconductor.CustomLineItemReference
	_jsii_.Get(
		j,
		"customLineItemRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) PresentationDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"presentationDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomLineItem) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::BillingConductor::CustomLineItem`.
func NewCfnCustomLineItem(scope constructs.Construct, id *string, props *CfnCustomLineItemProps) CfnCustomLineItem {
	_init_.Initialize()

	if err := validateNewCfnCustomLineItemParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCustomLineItem{}

	_jsii_.Create(
		"aws-cdk-lib.aws_billingconductor.CfnCustomLineItem",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::BillingConductor::CustomLineItem`.
func NewCfnCustomLineItem_Override(c CfnCustomLineItem, scope constructs.Construct, id *string, props *CfnCustomLineItemProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_billingconductor.CfnCustomLineItem",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCustomLineItem)SetAccountId(val *string) {
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_CfnCustomLineItem)SetBillingGroupArn(val *string) {
	if err := j.validateSetBillingGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingGroupArn",
		val,
	)
}

func (j *jsiiProxy_CfnCustomLineItem)SetBillingPeriodRange(val interface{}) {
	if err := j.validateSetBillingPeriodRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingPeriodRange",
		val,
	)
}

func (j *jsiiProxy_CfnCustomLineItem)SetComputationRule(val *string) {
	_jsii_.Set(
		j,
		"computationRule",
		val,
	)
}

func (j *jsiiProxy_CfnCustomLineItem)SetCustomLineItemChargeDetails(val interface{}) {
	if err := j.validateSetCustomLineItemChargeDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customLineItemChargeDetails",
		val,
	)
}

func (j *jsiiProxy_CfnCustomLineItem)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnCustomLineItem)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnCustomLineItem)SetPresentationDetails(val interface{}) {
	if err := j.validateSetPresentationDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"presentationDetails",
		val,
	)
}

func (j *jsiiProxy_CfnCustomLineItem)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func CfnCustomLineItem_ArnForCustomLineItem(resource interfacesawsbillingconductor.ICustomLineItemRef) *string {
	_init_.Initialize()

	if err := validateCfnCustomLineItem_ArnForCustomLineItemParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_billingconductor.CfnCustomLineItem",
		"arnForCustomLineItem",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnCustomLineItem.
func CfnCustomLineItem_IsCfnCustomLineItem(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCustomLineItem_IsCfnCustomLineItemParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_billingconductor.CfnCustomLineItem",
		"isCfnCustomLineItem",
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
func CfnCustomLineItem_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCustomLineItem_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_billingconductor.CfnCustomLineItem",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnCustomLineItem_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCustomLineItem_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_billingconductor.CfnCustomLineItem",
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
func CfnCustomLineItem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCustomLineItem_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_billingconductor.CfnCustomLineItem",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCustomLineItem_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_billingconductor.CfnCustomLineItem",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCustomLineItem) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCustomLineItem) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCustomLineItem) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCustomLineItem) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCustomLineItem) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCustomLineItem) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCustomLineItem) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCustomLineItem) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCustomLineItem) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnCustomLineItem) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnCustomLineItem) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCustomLineItem) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCustomLineItem) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCustomLineItem) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCustomLineItem) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCustomLineItem) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnCustomLineItem) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnCustomLineItem) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCustomLineItem) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCustomLineItem) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}


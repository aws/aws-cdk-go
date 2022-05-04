package awsbillingconductor

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsbillingconductor/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::BillingConductor::BillingGroup`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import billingconductor "github.com/aws/aws-cdk-go/awscdk/aws_billingconductor"
//   cfnBillingGroup := billingconductor.NewCfnBillingGroup(this, jsii.String("MyCfnBillingGroup"), &cfnBillingGroupProps{
//   	accountGrouping: &accountGroupingProperty{
//   		linkedAccountIds: []*string{
//   			jsii.String("linkedAccountIds"),
//   		},
//   	},
//   	computationPreference: &computationPreferenceProperty{
//   		pricingPlanArn: jsii.String("pricingPlanArn"),
//   	},
//   	name: jsii.String("name"),
//   	primaryAccountId: jsii.String("primaryAccountId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnBillingGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// `AWS::BillingConductor::BillingGroup.AccountGrouping`.
	AccountGrouping() interface{}
	SetAccountGrouping(val interface{})
	AttrArn() *string
	AttrCreationTime() *float64
	AttrLastModifiedTime() *float64
	AttrSize() *float64
	AttrStatus() *string
	AttrStatusReason() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// `AWS::BillingConductor::BillingGroup.ComputationPreference`.
	ComputationPreference() interface{}
	SetComputationPreference(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// `AWS::BillingConductor::BillingGroup.Description`.
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
	// Experimental.
	LogicalId() *string
	// `AWS::BillingConductor::BillingGroup.Name`.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// `AWS::BillingConductor::BillingGroup.PrimaryAccountId`.
	PrimaryAccountId() *string
	SetPrimaryAccountId(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::BillingConductor::BillingGroup.Tags`.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnBillingGroup
type jsiiProxy_CfnBillingGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnBillingGroup) AccountGrouping() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountGrouping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBillingGroup) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBillingGroup) AttrCreationTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBillingGroup) AttrLastModifiedTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrLastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBillingGroup) AttrSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBillingGroup) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBillingGroup) AttrStatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBillingGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBillingGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBillingGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBillingGroup) ComputationPreference() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"computationPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBillingGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBillingGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBillingGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBillingGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBillingGroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBillingGroup) PrimaryAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBillingGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBillingGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBillingGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBillingGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::BillingConductor::BillingGroup`.
func NewCfnBillingGroup(scope awscdk.Construct, id *string, props *CfnBillingGroupProps) CfnBillingGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnBillingGroup{}

	_jsii_.Create(
		"monocdk.aws_billingconductor.CfnBillingGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::BillingConductor::BillingGroup`.
func NewCfnBillingGroup_Override(c CfnBillingGroup, scope awscdk.Construct, id *string, props *CfnBillingGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_billingconductor.CfnBillingGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnBillingGroup) SetAccountGrouping(val interface{}) {
	_jsii_.Set(
		j,
		"accountGrouping",
		val,
	)
}

func (j *jsiiProxy_CfnBillingGroup) SetComputationPreference(val interface{}) {
	_jsii_.Set(
		j,
		"computationPreference",
		val,
	)
}

func (j *jsiiProxy_CfnBillingGroup) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnBillingGroup) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnBillingGroup) SetPrimaryAccountId(val *string) {
	_jsii_.Set(
		j,
		"primaryAccountId",
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
func CfnBillingGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_billingconductor.CfnBillingGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnBillingGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_billingconductor.CfnBillingGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnBillingGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_billingconductor.CfnBillingGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBillingGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_billingconductor.CfnBillingGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBillingGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnBillingGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnBillingGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnBillingGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnBillingGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnBillingGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnBillingGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnBillingGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBillingGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBillingGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnBillingGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnBillingGroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnBillingGroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBillingGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnBillingGroup) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnBillingGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBillingGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBillingGroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnBillingGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBillingGroup) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBillingGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import billingconductor "github.com/aws/aws-cdk-go/awscdk/aws_billingconductor"
//   accountGroupingProperty := &accountGroupingProperty{
//   	linkedAccountIds: []*string{
//   		jsii.String("linkedAccountIds"),
//   	},
//   }
//
type CfnBillingGroup_AccountGroupingProperty struct {
	// `CfnBillingGroup.AccountGroupingProperty.LinkedAccountIds`.
	LinkedAccountIds *[]*string `json:"linkedAccountIds" yaml:"linkedAccountIds"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import billingconductor "github.com/aws/aws-cdk-go/awscdk/aws_billingconductor"
//   computationPreferenceProperty := &computationPreferenceProperty{
//   	pricingPlanArn: jsii.String("pricingPlanArn"),
//   }
//
type CfnBillingGroup_ComputationPreferenceProperty struct {
	// `CfnBillingGroup.ComputationPreferenceProperty.PricingPlanArn`.
	PricingPlanArn *string `json:"pricingPlanArn" yaml:"pricingPlanArn"`
}

// Properties for defining a `CfnBillingGroup`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import billingconductor "github.com/aws/aws-cdk-go/awscdk/aws_billingconductor"
//   cfnBillingGroupProps := &cfnBillingGroupProps{
//   	accountGrouping: &accountGroupingProperty{
//   		linkedAccountIds: []*string{
//   			jsii.String("linkedAccountIds"),
//   		},
//   	},
//   	computationPreference: &computationPreferenceProperty{
//   		pricingPlanArn: jsii.String("pricingPlanArn"),
//   	},
//   	name: jsii.String("name"),
//   	primaryAccountId: jsii.String("primaryAccountId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnBillingGroupProps struct {
	// `AWS::BillingConductor::BillingGroup.AccountGrouping`.
	AccountGrouping interface{} `json:"accountGrouping" yaml:"accountGrouping"`
	// `AWS::BillingConductor::BillingGroup.ComputationPreference`.
	ComputationPreference interface{} `json:"computationPreference" yaml:"computationPreference"`
	// `AWS::BillingConductor::BillingGroup.Name`.
	Name *string `json:"name" yaml:"name"`
	// `AWS::BillingConductor::BillingGroup.PrimaryAccountId`.
	PrimaryAccountId *string `json:"primaryAccountId" yaml:"primaryAccountId"`
	// `AWS::BillingConductor::BillingGroup.Description`.
	Description *string `json:"description" yaml:"description"`
	// `AWS::BillingConductor::BillingGroup.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::BillingConductor::CustomLineItem`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import billingconductor "github.com/aws/aws-cdk-go/awscdk/aws_billingconductor"
//   cfnCustomLineItem := billingconductor.NewCfnCustomLineItem(this, jsii.String("MyCfnCustomLineItem"), &cfnCustomLineItemProps{
//   	billingGroupArn: jsii.String("billingGroupArn"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	billingPeriodRange: &billingPeriodRangeProperty{
//   		exclusiveEndBillingPeriod: jsii.String("exclusiveEndBillingPeriod"),
//   		inclusiveStartBillingPeriod: jsii.String("inclusiveStartBillingPeriod"),
//   	},
//   	customLineItemChargeDetails: &customLineItemChargeDetailsProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		flat: &customLineItemFlatChargeDetailsProperty{
//   			chargeValue: jsii.Number(123),
//   		},
//   		percentage: &customLineItemPercentageChargeDetailsProperty{
//   			percentageValue: jsii.Number(123),
//
//   			// the properties below are optional
//   			childAssociatedResources: []*string{
//   				jsii.String("childAssociatedResources"),
//   			},
//   		},
//   	},
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnCustomLineItem interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrAssociationSize() *float64
	AttrCreationTime() *float64
	AttrCurrencyCode() *string
	AttrLastModifiedTime() *float64
	AttrProductCode() *string
	// `AWS::BillingConductor::CustomLineItem.BillingGroupArn`.
	BillingGroupArn() *string
	SetBillingGroupArn(val *string)
	// `AWS::BillingConductor::CustomLineItem.BillingPeriodRange`.
	BillingPeriodRange() interface{}
	SetBillingPeriodRange(val interface{})
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// `AWS::BillingConductor::CustomLineItem.CustomLineItemChargeDetails`.
	CustomLineItemChargeDetails() interface{}
	SetCustomLineItemChargeDetails(val interface{})
	// `AWS::BillingConductor::CustomLineItem.Description`.
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
	// Experimental.
	LogicalId() *string
	// `AWS::BillingConductor::CustomLineItem.Name`.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::BillingConductor::CustomLineItem.Tags`.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnCustomLineItem
type jsiiProxy_CfnCustomLineItem struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
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

func (j *jsiiProxy_CfnCustomLineItem) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
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

func (j *jsiiProxy_CfnCustomLineItem) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
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

func (j *jsiiProxy_CfnCustomLineItem) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::BillingConductor::CustomLineItem`.
func NewCfnCustomLineItem(scope awscdk.Construct, id *string, props *CfnCustomLineItemProps) CfnCustomLineItem {
	_init_.Initialize()

	j := jsiiProxy_CfnCustomLineItem{}

	_jsii_.Create(
		"monocdk.aws_billingconductor.CfnCustomLineItem",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::BillingConductor::CustomLineItem`.
func NewCfnCustomLineItem_Override(c CfnCustomLineItem, scope awscdk.Construct, id *string, props *CfnCustomLineItemProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_billingconductor.CfnCustomLineItem",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCustomLineItem) SetBillingGroupArn(val *string) {
	_jsii_.Set(
		j,
		"billingGroupArn",
		val,
	)
}

func (j *jsiiProxy_CfnCustomLineItem) SetBillingPeriodRange(val interface{}) {
	_jsii_.Set(
		j,
		"billingPeriodRange",
		val,
	)
}

func (j *jsiiProxy_CfnCustomLineItem) SetCustomLineItemChargeDetails(val interface{}) {
	_jsii_.Set(
		j,
		"customLineItemChargeDetails",
		val,
	)
}

func (j *jsiiProxy_CfnCustomLineItem) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnCustomLineItem) SetName(val *string) {
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
// Experimental.
func CfnCustomLineItem_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_billingconductor.CfnCustomLineItem",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnCustomLineItem_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_billingconductor.CfnCustomLineItem",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnCustomLineItem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_billingconductor.CfnCustomLineItem",
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
		"monocdk.aws_billingconductor.CfnCustomLineItem",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCustomLineItem) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCustomLineItem) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCustomLineItem) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCustomLineItem) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCustomLineItem) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCustomLineItem) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCustomLineItem) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCustomLineItem) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCustomLineItem) GetMetadata(key *string) interface{} {
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
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCustomLineItem) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnCustomLineItem) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnCustomLineItem) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCustomLineItem) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCustomLineItem) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnCustomLineItem) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
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

func (c *jsiiProxy_CfnCustomLineItem) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

func (c *jsiiProxy_CfnCustomLineItem) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCustomLineItem) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import billingconductor "github.com/aws/aws-cdk-go/awscdk/aws_billingconductor"
//   billingPeriodRangeProperty := &billingPeriodRangeProperty{
//   	exclusiveEndBillingPeriod: jsii.String("exclusiveEndBillingPeriod"),
//   	inclusiveStartBillingPeriod: jsii.String("inclusiveStartBillingPeriod"),
//   }
//
type CfnCustomLineItem_BillingPeriodRangeProperty struct {
	// `CfnCustomLineItem.BillingPeriodRangeProperty.ExclusiveEndBillingPeriod`.
	ExclusiveEndBillingPeriod *string `json:"exclusiveEndBillingPeriod" yaml:"exclusiveEndBillingPeriod"`
	// `CfnCustomLineItem.BillingPeriodRangeProperty.InclusiveStartBillingPeriod`.
	InclusiveStartBillingPeriod *string `json:"inclusiveStartBillingPeriod" yaml:"inclusiveStartBillingPeriod"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import billingconductor "github.com/aws/aws-cdk-go/awscdk/aws_billingconductor"
//   customLineItemChargeDetailsProperty := &customLineItemChargeDetailsProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	flat: &customLineItemFlatChargeDetailsProperty{
//   		chargeValue: jsii.Number(123),
//   	},
//   	percentage: &customLineItemPercentageChargeDetailsProperty{
//   		percentageValue: jsii.Number(123),
//
//   		// the properties below are optional
//   		childAssociatedResources: []*string{
//   			jsii.String("childAssociatedResources"),
//   		},
//   	},
//   }
//
type CfnCustomLineItem_CustomLineItemChargeDetailsProperty struct {
	// `CfnCustomLineItem.CustomLineItemChargeDetailsProperty.Type`.
	Type *string `json:"type" yaml:"type"`
	// `CfnCustomLineItem.CustomLineItemChargeDetailsProperty.Flat`.
	Flat interface{} `json:"flat" yaml:"flat"`
	// `CfnCustomLineItem.CustomLineItemChargeDetailsProperty.Percentage`.
	Percentage interface{} `json:"percentage" yaml:"percentage"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import billingconductor "github.com/aws/aws-cdk-go/awscdk/aws_billingconductor"
//   customLineItemFlatChargeDetailsProperty := &customLineItemFlatChargeDetailsProperty{
//   	chargeValue: jsii.Number(123),
//   }
//
type CfnCustomLineItem_CustomLineItemFlatChargeDetailsProperty struct {
	// `CfnCustomLineItem.CustomLineItemFlatChargeDetailsProperty.ChargeValue`.
	ChargeValue *float64 `json:"chargeValue" yaml:"chargeValue"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import billingconductor "github.com/aws/aws-cdk-go/awscdk/aws_billingconductor"
//   customLineItemPercentageChargeDetailsProperty := &customLineItemPercentageChargeDetailsProperty{
//   	percentageValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	childAssociatedResources: []*string{
//   		jsii.String("childAssociatedResources"),
//   	},
//   }
//
type CfnCustomLineItem_CustomLineItemPercentageChargeDetailsProperty struct {
	// `CfnCustomLineItem.CustomLineItemPercentageChargeDetailsProperty.PercentageValue`.
	PercentageValue *float64 `json:"percentageValue" yaml:"percentageValue"`
	// `CfnCustomLineItem.CustomLineItemPercentageChargeDetailsProperty.ChildAssociatedResources`.
	ChildAssociatedResources *[]*string `json:"childAssociatedResources" yaml:"childAssociatedResources"`
}

// Properties for defining a `CfnCustomLineItem`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import billingconductor "github.com/aws/aws-cdk-go/awscdk/aws_billingconductor"
//   cfnCustomLineItemProps := &cfnCustomLineItemProps{
//   	billingGroupArn: jsii.String("billingGroupArn"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	billingPeriodRange: &billingPeriodRangeProperty{
//   		exclusiveEndBillingPeriod: jsii.String("exclusiveEndBillingPeriod"),
//   		inclusiveStartBillingPeriod: jsii.String("inclusiveStartBillingPeriod"),
//   	},
//   	customLineItemChargeDetails: &customLineItemChargeDetailsProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		flat: &customLineItemFlatChargeDetailsProperty{
//   			chargeValue: jsii.Number(123),
//   		},
//   		percentage: &customLineItemPercentageChargeDetailsProperty{
//   			percentageValue: jsii.Number(123),
//
//   			// the properties below are optional
//   			childAssociatedResources: []*string{
//   				jsii.String("childAssociatedResources"),
//   			},
//   		},
//   	},
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnCustomLineItemProps struct {
	// `AWS::BillingConductor::CustomLineItem.BillingGroupArn`.
	BillingGroupArn *string `json:"billingGroupArn" yaml:"billingGroupArn"`
	// `AWS::BillingConductor::CustomLineItem.Name`.
	Name *string `json:"name" yaml:"name"`
	// `AWS::BillingConductor::CustomLineItem.BillingPeriodRange`.
	BillingPeriodRange interface{} `json:"billingPeriodRange" yaml:"billingPeriodRange"`
	// `AWS::BillingConductor::CustomLineItem.CustomLineItemChargeDetails`.
	CustomLineItemChargeDetails interface{} `json:"customLineItemChargeDetails" yaml:"customLineItemChargeDetails"`
	// `AWS::BillingConductor::CustomLineItem.Description`.
	Description *string `json:"description" yaml:"description"`
	// `AWS::BillingConductor::CustomLineItem.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::BillingConductor::PricingPlan`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import billingconductor "github.com/aws/aws-cdk-go/awscdk/aws_billingconductor"
//   cfnPricingPlan := billingconductor.NewCfnPricingPlan(this, jsii.String("MyCfnPricingPlan"), &cfnPricingPlanProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	pricingRuleArns: []*string{
//   		jsii.String("pricingRuleArns"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnPricingPlan interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrCreationTime() *float64
	AttrLastModifiedTime() *float64
	AttrSize() *float64
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// `AWS::BillingConductor::PricingPlan.Description`.
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
	// Experimental.
	LogicalId() *string
	// `AWS::BillingConductor::PricingPlan.Name`.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// `AWS::BillingConductor::PricingPlan.PricingRuleArns`.
	PricingRuleArns() *[]*string
	SetPricingRuleArns(val *[]*string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::BillingConductor::PricingPlan.Tags`.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnPricingPlan
type jsiiProxy_CfnPricingPlan struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPricingPlan) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingPlan) AttrCreationTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingPlan) AttrLastModifiedTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrLastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingPlan) AttrSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingPlan) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingPlan) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingPlan) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingPlan) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingPlan) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingPlan) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingPlan) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingPlan) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingPlan) PricingRuleArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pricingRuleArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingPlan) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingPlan) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingPlan) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingPlan) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::BillingConductor::PricingPlan`.
func NewCfnPricingPlan(scope awscdk.Construct, id *string, props *CfnPricingPlanProps) CfnPricingPlan {
	_init_.Initialize()

	j := jsiiProxy_CfnPricingPlan{}

	_jsii_.Create(
		"monocdk.aws_billingconductor.CfnPricingPlan",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::BillingConductor::PricingPlan`.
func NewCfnPricingPlan_Override(c CfnPricingPlan, scope awscdk.Construct, id *string, props *CfnPricingPlanProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_billingconductor.CfnPricingPlan",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPricingPlan) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnPricingPlan) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnPricingPlan) SetPricingRuleArns(val *[]*string) {
	_jsii_.Set(
		j,
		"pricingRuleArns",
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
func CfnPricingPlan_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_billingconductor.CfnPricingPlan",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnPricingPlan_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_billingconductor.CfnPricingPlan",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnPricingPlan_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_billingconductor.CfnPricingPlan",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPricingPlan_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_billingconductor.CfnPricingPlan",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPricingPlan) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPricingPlan) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPricingPlan) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPricingPlan) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPricingPlan) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPricingPlan) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPricingPlan) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPricingPlan) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPricingPlan) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPricingPlan) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPricingPlan) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnPricingPlan) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnPricingPlan) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPricingPlan) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPricingPlan) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnPricingPlan) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPricingPlan) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPricingPlan) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnPricingPlan) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPricingPlan) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPricingPlan) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnPricingPlan`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import billingconductor "github.com/aws/aws-cdk-go/awscdk/aws_billingconductor"
//   cfnPricingPlanProps := &cfnPricingPlanProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	pricingRuleArns: []*string{
//   		jsii.String("pricingRuleArns"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPricingPlanProps struct {
	// `AWS::BillingConductor::PricingPlan.Name`.
	Name *string `json:"name" yaml:"name"`
	// `AWS::BillingConductor::PricingPlan.Description`.
	Description *string `json:"description" yaml:"description"`
	// `AWS::BillingConductor::PricingPlan.PricingRuleArns`.
	PricingRuleArns *[]*string `json:"pricingRuleArns" yaml:"pricingRuleArns"`
	// `AWS::BillingConductor::PricingPlan.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::BillingConductor::PricingRule`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import billingconductor "github.com/aws/aws-cdk-go/awscdk/aws_billingconductor"
//   cfnPricingRule := billingconductor.NewCfnPricingRule(this, jsii.String("MyCfnPricingRule"), &cfnPricingRuleProps{
//   	modifierPercentage: jsii.Number(123),
//   	name: jsii.String("name"),
//   	scope: jsii.String("scope"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	service: jsii.String("service"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnPricingRule interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrAssociatedPricingPlanCount() *float64
	AttrCreationTime() *float64
	AttrLastModifiedTime() *float64
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// `AWS::BillingConductor::PricingRule.Description`.
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
	// Experimental.
	LogicalId() *string
	// `AWS::BillingConductor::PricingRule.ModifierPercentage`.
	ModifierPercentage() *float64
	SetModifierPercentage(val *float64)
	// `AWS::BillingConductor::PricingRule.Name`.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// `AWS::BillingConductor::PricingRule.Scope`.
	Scope() *string
	SetScope(val *string)
	// `AWS::BillingConductor::PricingRule.Service`.
	Service() *string
	SetService(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::BillingConductor::PricingRule.Tags`.
	Tags() awscdk.TagManager
	// `AWS::BillingConductor::PricingRule.Type`.
	Type() *string
	SetType(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnPricingRule
type jsiiProxy_CfnPricingRule struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPricingRule) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingRule) AttrAssociatedPricingPlanCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrAssociatedPricingPlanCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingRule) AttrCreationTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingRule) AttrLastModifiedTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrLastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingRule) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingRule) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingRule) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingRule) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingRule) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingRule) ModifierPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"modifierPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingRule) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingRule) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingRule) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingRule) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingRule) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingRule) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPricingRule) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::BillingConductor::PricingRule`.
func NewCfnPricingRule(scope awscdk.Construct, id *string, props *CfnPricingRuleProps) CfnPricingRule {
	_init_.Initialize()

	j := jsiiProxy_CfnPricingRule{}

	_jsii_.Create(
		"monocdk.aws_billingconductor.CfnPricingRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::BillingConductor::PricingRule`.
func NewCfnPricingRule_Override(c CfnPricingRule, scope awscdk.Construct, id *string, props *CfnPricingRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_billingconductor.CfnPricingRule",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPricingRule) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnPricingRule) SetModifierPercentage(val *float64) {
	_jsii_.Set(
		j,
		"modifierPercentage",
		val,
	)
}

func (j *jsiiProxy_CfnPricingRule) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnPricingRule) SetScope(val *string) {
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_CfnPricingRule) SetService(val *string) {
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_CfnPricingRule) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
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
func CfnPricingRule_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_billingconductor.CfnPricingRule",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnPricingRule_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_billingconductor.CfnPricingRule",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnPricingRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_billingconductor.CfnPricingRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPricingRule_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_billingconductor.CfnPricingRule",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPricingRule) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPricingRule) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPricingRule) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPricingRule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPricingRule) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPricingRule) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPricingRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPricingRule) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPricingRule) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPricingRule) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPricingRule) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnPricingRule) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnPricingRule) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPricingRule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPricingRule) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnPricingRule) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPricingRule) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPricingRule) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnPricingRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPricingRule) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPricingRule) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnPricingRule`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import billingconductor "github.com/aws/aws-cdk-go/awscdk/aws_billingconductor"
//   cfnPricingRuleProps := &cfnPricingRuleProps{
//   	modifierPercentage: jsii.Number(123),
//   	name: jsii.String("name"),
//   	scope: jsii.String("scope"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	service: jsii.String("service"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPricingRuleProps struct {
	// `AWS::BillingConductor::PricingRule.ModifierPercentage`.
	ModifierPercentage *float64 `json:"modifierPercentage" yaml:"modifierPercentage"`
	// `AWS::BillingConductor::PricingRule.Name`.
	Name *string `json:"name" yaml:"name"`
	// `AWS::BillingConductor::PricingRule.Scope`.
	Scope *string `json:"scope" yaml:"scope"`
	// `AWS::BillingConductor::PricingRule.Type`.
	Type *string `json:"type" yaml:"type"`
	// `AWS::BillingConductor::PricingRule.Description`.
	Description *string `json:"description" yaml:"description"`
	// `AWS::BillingConductor::PricingRule.Service`.
	Service *string `json:"service" yaml:"service"`
	// `AWS::BillingConductor::PricingRule.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}


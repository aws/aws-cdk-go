package awsconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsconfig/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
	"github.com/aws/constructs-go/constructs/v3"
)

// Checks whether the active access keys are rotated within the number of days specified in `maxAge`.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/config/latest/developerguide/access-keys-rotated.html
//
// Experimental.
type AccessKeysRotated interface {
	ManagedRule
	ConfigRuleArn() *string
	ConfigRuleComplianceType() *string
	ConfigRuleId() *string
	ConfigRuleName() *string
	Env() *awscdk.ResourceEnvironment
	IsCustomWithChanges() *bool
	SetIsCustomWithChanges(val *bool)
	IsManaged() *bool
	SetIsManaged(val *bool)
	Node() awscdk.ConstructNode
	PhysicalName() *string
	RuleScope() RuleScope
	SetRuleScope(val RuleScope)
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnComplianceChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnPrepare()
	OnReEvaluationStatus(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for AccessKeysRotated
type jsiiProxy_AccessKeysRotated struct {
	jsiiProxy_ManagedRule
}

func (j *jsiiProxy_AccessKeysRotated) ConfigRuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessKeysRotated) ConfigRuleComplianceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleComplianceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessKeysRotated) ConfigRuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessKeysRotated) ConfigRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessKeysRotated) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessKeysRotated) IsCustomWithChanges() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isCustomWithChanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessKeysRotated) IsManaged() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isManaged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessKeysRotated) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessKeysRotated) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessKeysRotated) RuleScope() RuleScope {
	var returns RuleScope
	_jsii_.Get(
		j,
		"ruleScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessKeysRotated) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewAccessKeysRotated(scope constructs.Construct, id *string, props *AccessKeysRotatedProps) AccessKeysRotated {
	_init_.Initialize()

	j := jsiiProxy_AccessKeysRotated{}

	_jsii_.Create(
		"monocdk.aws_config.AccessKeysRotated",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAccessKeysRotated_Override(a AccessKeysRotated, scope constructs.Construct, id *string, props *AccessKeysRotatedProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_config.AccessKeysRotated",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_AccessKeysRotated) SetIsCustomWithChanges(val *bool) {
	_jsii_.Set(
		j,
		"isCustomWithChanges",
		val,
	)
}

func (j *jsiiProxy_AccessKeysRotated) SetIsManaged(val *bool) {
	_jsii_.Set(
		j,
		"isManaged",
		val,
	)
}

func (j *jsiiProxy_AccessKeysRotated) SetRuleScope(val RuleScope) {
	_jsii_.Set(
		j,
		"ruleScope",
		val,
	)
}

// Imports an existing rule.
// Experimental.
func AccessKeysRotated_FromConfigRuleName(scope constructs.Construct, id *string, configRuleName *string) IRule {
	_init_.Initialize()

	var returns IRule

	_jsii_.StaticInvoke(
		"monocdk.aws_config.AccessKeysRotated",
		"fromConfigRuleName",
		[]interface{}{scope, id, configRuleName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func AccessKeysRotated_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.AccessKeysRotated",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func AccessKeysRotated_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.AccessKeysRotated",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (a *jsiiProxy_AccessKeysRotated) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (a *jsiiProxy_AccessKeysRotated) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (a *jsiiProxy_AccessKeysRotated) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (a *jsiiProxy_AccessKeysRotated) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Defines an EventBridge event rule which triggers for rule compliance events.
// Experimental.
func (a *jsiiProxy_AccessKeysRotated) OnComplianceChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		a,
		"onComplianceChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines an EventBridge event rule which triggers for rule events.
//
// Use
// `rule.addEventPattern(pattern)` to specify a filter.
// Experimental.
func (a *jsiiProxy_AccessKeysRotated) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		a,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (a *jsiiProxy_AccessKeysRotated) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

// Defines an EventBridge event rule which triggers for rule re-evaluation status events.
// Experimental.
func (a *jsiiProxy_AccessKeysRotated) OnReEvaluationStatus(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		a,
		"onReEvaluationStatus",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_AccessKeysRotated) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (a *jsiiProxy_AccessKeysRotated) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (a *jsiiProxy_AccessKeysRotated) Prepare() {
	_jsii_.InvokeVoid(
		a,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_AccessKeysRotated) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_AccessKeysRotated) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (a *jsiiProxy_AccessKeysRotated) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a AccessKeysRotated.
//
// TODO: EXAMPLE
//
// Experimental.
type AccessKeysRotatedProps struct {
	// A name for the AWS Config rule.
	// Experimental.
	ConfigRuleName *string `json:"configRuleName" yaml:"configRuleName"`
	// A description about this AWS Config rule.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// Input parameter values that are passed to the AWS Config rule.
	// Experimental.
	InputParameters *map[string]interface{} `json:"inputParameters" yaml:"inputParameters"`
	// The maximum frequency at which the AWS Config rule runs evaluations.
	// Experimental.
	MaximumExecutionFrequency MaximumExecutionFrequency `json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// Defines which resources trigger an evaluation for an AWS Config rule.
	// Experimental.
	RuleScope RuleScope `json:"ruleScope" yaml:"ruleScope"`
	// The maximum number of days within which the access keys must be rotated.
	// Experimental.
	MaxAge awscdk.Duration `json:"maxAge" yaml:"maxAge"`
}

// A CloudFormation `AWS::Config::AggregationAuthorization`.
//
// An object that represents the authorizations granted to aggregator accounts and regions.
//
// TODO: EXAMPLE
//
type CfnAggregationAuthorization interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrAggregationAuthorizationArn() *string
	AuthorizedAccountId() *string
	SetAuthorizedAccountId(val *string)
	AuthorizedAwsRegion() *string
	SetAuthorizedAwsRegion(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() awscdk.ConstructNode
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnAggregationAuthorization
type jsiiProxy_CfnAggregationAuthorization struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAggregationAuthorization) AttrAggregationAuthorizationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAggregationAuthorizationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAggregationAuthorization) AuthorizedAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizedAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAggregationAuthorization) AuthorizedAwsRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizedAwsRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAggregationAuthorization) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAggregationAuthorization) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAggregationAuthorization) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAggregationAuthorization) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAggregationAuthorization) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAggregationAuthorization) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAggregationAuthorization) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAggregationAuthorization) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAggregationAuthorization) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAggregationAuthorization) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Config::AggregationAuthorization`.
func NewCfnAggregationAuthorization(scope awscdk.Construct, id *string, props *CfnAggregationAuthorizationProps) CfnAggregationAuthorization {
	_init_.Initialize()

	j := jsiiProxy_CfnAggregationAuthorization{}

	_jsii_.Create(
		"monocdk.aws_config.CfnAggregationAuthorization",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Config::AggregationAuthorization`.
func NewCfnAggregationAuthorization_Override(c CfnAggregationAuthorization, scope awscdk.Construct, id *string, props *CfnAggregationAuthorizationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_config.CfnAggregationAuthorization",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAggregationAuthorization) SetAuthorizedAccountId(val *string) {
	_jsii_.Set(
		j,
		"authorizedAccountId",
		val,
	)
}

func (j *jsiiProxy_CfnAggregationAuthorization) SetAuthorizedAwsRegion(val *string) {
	_jsii_.Set(
		j,
		"authorizedAwsRegion",
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
func CfnAggregationAuthorization_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnAggregationAuthorization",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnAggregationAuthorization_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnAggregationAuthorization",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnAggregationAuthorization_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnAggregationAuthorization",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAggregationAuthorization_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.CfnAggregationAuthorization",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnAggregationAuthorization) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnAggregationAuthorization) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnAggregationAuthorization) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
// Experimental.
func (c *jsiiProxy_CfnAggregationAuthorization) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnAggregationAuthorization) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnAggregationAuthorization) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

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
func (c *jsiiProxy_CfnAggregationAuthorization) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnAggregationAuthorization) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnAggregationAuthorization) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnAggregationAuthorization) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnAggregationAuthorization) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnAggregationAuthorization) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnAggregationAuthorization) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnAggregationAuthorization) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnAggregationAuthorization) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnAggregationAuthorization) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnAggregationAuthorization) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnAggregationAuthorization) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnAggregationAuthorization) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnAggregationAuthorization) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnAggregationAuthorization) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnAggregationAuthorization`.
//
// TODO: EXAMPLE
//
type CfnAggregationAuthorizationProps struct {
	// The 12-digit account ID of the account authorized to aggregate data.
	AuthorizedAccountId *string `json:"authorizedAccountId" yaml:"authorizedAccountId"`
	// The region authorized to collect aggregated data.
	AuthorizedAwsRegion *string `json:"authorizedAwsRegion" yaml:"authorizedAwsRegion"`
	// An array of tag object.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Config::ConfigRule`.
//
// Specifies an AWS Config rule for evaluating whether your AWS resources comply with your desired configurations.
//
// You can use this action for custom AWS Config rules and AWS managed Config rules. A custom AWS Config rule is a rule that you develop and maintain. An AWS managed Config rule is a customizable, predefined rule that AWS Config provides.
//
// If you are adding a new custom AWS Config rule, you must first create the AWS Lambda function that the rule invokes to evaluate your resources. When you use the `PutConfigRule` action to add the rule to AWS Config , you must specify the Amazon Resource Name (ARN) that AWS Lambda assigns to the function. Specify the ARN for the `SourceIdentifier` key. This key is part of the `Source` object, which is part of the `ConfigRule` object.
//
// If you are adding an AWS managed Config rule, specify the rule's identifier for the `SourceIdentifier` key. To reference AWS managed Config rule identifiers, see [About AWS Managed Config Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html) .
//
// For any new rule that you add, specify the `ConfigRuleName` in the `ConfigRule` object. Do not specify the `ConfigRuleArn` or the `ConfigRuleId` . These values are generated by AWS Config for new rules.
//
// If you are updating a rule that you added previously, you can specify the rule by `ConfigRuleName` , `ConfigRuleId` , or `ConfigRuleArn` in the `ConfigRule` data type that you use in this request.
//
// The maximum number of rules that AWS Config supports is 150.
//
// For information about requesting a rule limit increase, see [AWS Config Limits](https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html#limits_config) in the *AWS General Reference Guide* .
//
// For more information about developing and using AWS Config rules, see [Evaluating AWS Resource Configurations with AWS Config](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config.html) in the *AWS Config Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnConfigRule interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrComplianceType() *string
	AttrConfigRuleId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ConfigRuleName() *string
	SetConfigRuleName(val *string)
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	InputParameters() interface{}
	SetInputParameters(val interface{})
	LogicalId() *string
	MaximumExecutionFrequency() *string
	SetMaximumExecutionFrequency(val *string)
	Node() awscdk.ConstructNode
	Ref() *string
	Scope() interface{}
	SetScope(val interface{})
	Source() interface{}
	SetSource(val interface{})
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnConfigRule
type jsiiProxy_CfnConfigRule struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnConfigRule) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigRule) AttrComplianceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrComplianceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigRule) AttrConfigRuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrConfigRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigRule) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigRule) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigRule) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigRule) ConfigRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigRule) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigRule) InputParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigRule) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigRule) MaximumExecutionFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumExecutionFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigRule) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigRule) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigRule) Scope() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigRule) Source() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigRule) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Config::ConfigRule`.
func NewCfnConfigRule(scope awscdk.Construct, id *string, props *CfnConfigRuleProps) CfnConfigRule {
	_init_.Initialize()

	j := jsiiProxy_CfnConfigRule{}

	_jsii_.Create(
		"monocdk.aws_config.CfnConfigRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Config::ConfigRule`.
func NewCfnConfigRule_Override(c CfnConfigRule, scope awscdk.Construct, id *string, props *CfnConfigRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_config.CfnConfigRule",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnConfigRule) SetConfigRuleName(val *string) {
	_jsii_.Set(
		j,
		"configRuleName",
		val,
	)
}

func (j *jsiiProxy_CfnConfigRule) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnConfigRule) SetInputParameters(val interface{}) {
	_jsii_.Set(
		j,
		"inputParameters",
		val,
	)
}

func (j *jsiiProxy_CfnConfigRule) SetMaximumExecutionFrequency(val *string) {
	_jsii_.Set(
		j,
		"maximumExecutionFrequency",
		val,
	)
}

func (j *jsiiProxy_CfnConfigRule) SetScope(val interface{}) {
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_CfnConfigRule) SetSource(val interface{}) {
	_jsii_.Set(
		j,
		"source",
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
func CfnConfigRule_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnConfigRule",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnConfigRule_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnConfigRule",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnConfigRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnConfigRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConfigRule_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.CfnConfigRule",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnConfigRule) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnConfigRule) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnConfigRule) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
// Experimental.
func (c *jsiiProxy_CfnConfigRule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnConfigRule) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnConfigRule) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

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
func (c *jsiiProxy_CfnConfigRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnConfigRule) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnConfigRule) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnConfigRule) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnConfigRule) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnConfigRule) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnConfigRule) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnConfigRule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnConfigRule) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnConfigRule) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnConfigRule) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnConfigRule) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnConfigRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnConfigRule) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnConfigRule) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Defines which resources trigger an evaluation for an AWS Config rule.
//
// The scope can include one or more resource types, a combination of a tag key and value, or a combination of one resource type and one resource ID. Specify a scope to constrain which resources trigger an evaluation for a rule. Otherwise, evaluations for the rule are triggered when any resource in your recording group changes in configuration.
//
// TODO: EXAMPLE
//
type CfnConfigRule_ScopeProperty struct {
	// The ID of the only AWS resource that you want to trigger an evaluation for the rule.
	//
	// If you specify a resource ID, you must specify one resource type for `ComplianceResourceTypes` .
	ComplianceResourceId *string `json:"complianceResourceId" yaml:"complianceResourceId"`
	// The resource types of only those AWS resources that you want to trigger an evaluation for the rule.
	//
	// You can only specify one type if you also specify a resource ID for `ComplianceResourceId` .
	ComplianceResourceTypes *[]*string `json:"complianceResourceTypes" yaml:"complianceResourceTypes"`
	// The tag key that is applied to only those AWS resources that you want to trigger an evaluation for the rule.
	TagKey *string `json:"tagKey" yaml:"tagKey"`
	// The tag value applied to only those AWS resources that you want to trigger an evaluation for the rule.
	//
	// If you specify a value for `TagValue` , you must also specify a value for `TagKey` .
	TagValue *string `json:"tagValue" yaml:"tagValue"`
}

// Provides the source and the message types that trigger AWS Config to evaluate your AWS resources against a rule.
//
// It also provides the frequency with which you want AWS Config to run evaluations for the rule if the trigger type is periodic. You can specify the parameter values for `SourceDetail` only for custom rules.
//
// TODO: EXAMPLE
//
type CfnConfigRule_SourceDetailProperty struct {
	// The source of the event, such as an AWS service, that triggers AWS Config to evaluate your AWS resources.
	EventSource *string `json:"eventSource" yaml:"eventSource"`
	// The type of notification that triggers AWS Config to run an evaluation for a rule.
	//
	// You can specify the following notification types:
	//
	// - `ConfigurationItemChangeNotification` - Triggers an evaluation when AWS Config delivers a configuration item as a result of a resource change.
	// - `OversizedConfigurationItemChangeNotification` - Triggers an evaluation when AWS Config delivers an oversized configuration item. AWS Config may generate this notification type when a resource changes and the notification exceeds the maximum size allowed by Amazon SNS.
	// - `ScheduledNotification` - Triggers a periodic evaluation at the frequency specified for `MaximumExecutionFrequency` .
	// - `ConfigurationSnapshotDeliveryCompleted` - Triggers a periodic evaluation when AWS Config delivers a configuration snapshot.
	//
	// If you want your custom rule to be triggered by configuration changes, specify two SourceDetail objects, one for `ConfigurationItemChangeNotification` and one for `OversizedConfigurationItemChangeNotification` .
	MessageType *string `json:"messageType" yaml:"messageType"`
	// The frequency at which you want AWS Config to run evaluations for a custom rule with a periodic trigger.
	//
	// If you specify a value for `MaximumExecutionFrequency` , then `MessageType` must use the `ScheduledNotification` value.
	//
	// > By default, rules with a periodic trigger are evaluated every 24 hours. To change the frequency, specify a valid value for the `MaximumExecutionFrequency` parameter.
	// >
	// > Based on the valid value you choose, AWS Config runs evaluations once for each valid value. For example, if you choose `Three_Hours` , AWS Config runs evaluations once every three hours. In this case, `Three_Hours` is the frequency of this rule.
	MaximumExecutionFrequency *string `json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
}

// Provides the AWS Config rule owner ( AWS or customer), the rule identifier, and the events that trigger the evaluation of your AWS resources.
//
// TODO: EXAMPLE
//
type CfnConfigRule_SourceProperty struct {
	// Indicates whether AWS or the customer owns and manages the AWS Config rule.
	Owner *string `json:"owner" yaml:"owner"`
	// For AWS Config managed rules, a predefined identifier from a list.
	//
	// For example, `IAM_PASSWORD_POLICY` is a managed rule. To reference a managed rule, see [Using AWS Config managed rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html) .
	//
	// For custom rules, the identifier is the Amazon Resource Name (ARN) of the rule's AWS Lambda function, such as `arn:aws:lambda:us-east-2:123456789012:function:custom_rule_name` .
	SourceIdentifier *string `json:"sourceIdentifier" yaml:"sourceIdentifier"`
	// Provides the source and type of the event that causes AWS Config to evaluate your AWS resources.
	SourceDetails interface{} `json:"sourceDetails" yaml:"sourceDetails"`
}

// Properties for defining a `CfnConfigRule`.
//
// TODO: EXAMPLE
//
type CfnConfigRuleProps struct {
	// Provides the rule owner ( AWS or customer), the rule identifier, and the notifications that cause the function to evaluate your AWS resources.
	Source interface{} `json:"source" yaml:"source"`
	// A name for the AWS Config rule.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the rule name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	ConfigRuleName *string `json:"configRuleName" yaml:"configRuleName"`
	// The description that you provide for the AWS Config rule.
	Description *string `json:"description" yaml:"description"`
	// A string, in JSON format, that is passed to the AWS Config rule Lambda function.
	InputParameters interface{} `json:"inputParameters" yaml:"inputParameters"`
	// The maximum frequency with which AWS Config runs evaluations for a rule.
	//
	// You can specify a value for `MaximumExecutionFrequency` when:
	//
	// - You are using an AWS managed rule that is triggered at a periodic frequency.
	// - Your custom rule is triggered when AWS Config delivers the configuration snapshot. For more information, see [ConfigSnapshotDeliveryProperties](https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigSnapshotDeliveryProperties.html) .
	//
	// > By default, rules with a periodic trigger are evaluated every 24 hours. To change the frequency, specify a valid value for the `MaximumExecutionFrequency` parameter.
	MaximumExecutionFrequency *string `json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// Defines which resources can trigger an evaluation for the rule.
	//
	// The scope can include one or more resource types, a combination of one resource type and one resource ID, or a combination of a tag key and value. Specify a scope to constrain the resources that can trigger an evaluation for the rule. If you do not specify a scope, evaluations are triggered when any resource in the recording group changes.
	//
	// > The scope can be empty.
	Scope interface{} `json:"scope" yaml:"scope"`
}

// A CloudFormation `AWS::Config::ConfigurationAggregator`.
//
// The details about the configuration aggregator, including information about source accounts, regions, and metadata of the aggregator.
//
// TODO: EXAMPLE
//
type CfnConfigurationAggregator interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AccountAggregationSources() interface{}
	SetAccountAggregationSources(val interface{})
	AttrConfigurationAggregatorArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ConfigurationAggregatorName() *string
	SetConfigurationAggregatorName(val *string)
	CreationStack() *[]*string
	LogicalId() *string
	Node() awscdk.ConstructNode
	OrganizationAggregationSource() interface{}
	SetOrganizationAggregationSource(val interface{})
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnConfigurationAggregator
type jsiiProxy_CfnConfigurationAggregator struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnConfigurationAggregator) AccountAggregationSources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountAggregationSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationAggregator) AttrConfigurationAggregatorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrConfigurationAggregatorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationAggregator) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationAggregator) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationAggregator) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationAggregator) ConfigurationAggregatorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationAggregatorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationAggregator) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationAggregator) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationAggregator) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationAggregator) OrganizationAggregationSource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"organizationAggregationSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationAggregator) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationAggregator) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationAggregator) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationAggregator) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Config::ConfigurationAggregator`.
func NewCfnConfigurationAggregator(scope awscdk.Construct, id *string, props *CfnConfigurationAggregatorProps) CfnConfigurationAggregator {
	_init_.Initialize()

	j := jsiiProxy_CfnConfigurationAggregator{}

	_jsii_.Create(
		"monocdk.aws_config.CfnConfigurationAggregator",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Config::ConfigurationAggregator`.
func NewCfnConfigurationAggregator_Override(c CfnConfigurationAggregator, scope awscdk.Construct, id *string, props *CfnConfigurationAggregatorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_config.CfnConfigurationAggregator",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnConfigurationAggregator) SetAccountAggregationSources(val interface{}) {
	_jsii_.Set(
		j,
		"accountAggregationSources",
		val,
	)
}

func (j *jsiiProxy_CfnConfigurationAggregator) SetConfigurationAggregatorName(val *string) {
	_jsii_.Set(
		j,
		"configurationAggregatorName",
		val,
	)
}

func (j *jsiiProxy_CfnConfigurationAggregator) SetOrganizationAggregationSource(val interface{}) {
	_jsii_.Set(
		j,
		"organizationAggregationSource",
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
func CfnConfigurationAggregator_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnConfigurationAggregator",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnConfigurationAggregator_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnConfigurationAggregator",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnConfigurationAggregator_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnConfigurationAggregator",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConfigurationAggregator_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.CfnConfigurationAggregator",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnConfigurationAggregator) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnConfigurationAggregator) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnConfigurationAggregator) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
// Experimental.
func (c *jsiiProxy_CfnConfigurationAggregator) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnConfigurationAggregator) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnConfigurationAggregator) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

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
func (c *jsiiProxy_CfnConfigurationAggregator) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnConfigurationAggregator) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnConfigurationAggregator) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnConfigurationAggregator) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnConfigurationAggregator) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnConfigurationAggregator) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnConfigurationAggregator) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnConfigurationAggregator) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnConfigurationAggregator) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnConfigurationAggregator) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnConfigurationAggregator) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnConfigurationAggregator) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnConfigurationAggregator) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnConfigurationAggregator) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnConfigurationAggregator) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A collection of accounts and regions.
//
// TODO: EXAMPLE
//
type CfnConfigurationAggregator_AccountAggregationSourceProperty struct {
	// The 12-digit account ID of the account being aggregated.
	AccountIds *[]*string `json:"accountIds" yaml:"accountIds"`
	// If true, aggregate existing AWS Config regions and future regions.
	AllAwsRegions interface{} `json:"allAwsRegions" yaml:"allAwsRegions"`
	// The source regions being aggregated.
	AwsRegions *[]*string `json:"awsRegions" yaml:"awsRegions"`
}

// This object contains regions to set up the aggregator and an IAM role to retrieve organization details.
//
// TODO: EXAMPLE
//
type CfnConfigurationAggregator_OrganizationAggregationSourceProperty struct {
	// ARN of the IAM role used to retrieve AWS Organizations details associated with the aggregator account.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// If true, aggregate existing AWS Config regions and future regions.
	AllAwsRegions interface{} `json:"allAwsRegions" yaml:"allAwsRegions"`
	// The source regions being aggregated.
	AwsRegions *[]*string `json:"awsRegions" yaml:"awsRegions"`
}

// Properties for defining a `CfnConfigurationAggregator`.
//
// TODO: EXAMPLE
//
type CfnConfigurationAggregatorProps struct {
	// Provides a list of source accounts and regions to be aggregated.
	AccountAggregationSources interface{} `json:"accountAggregationSources" yaml:"accountAggregationSources"`
	// The name of the aggregator.
	ConfigurationAggregatorName *string `json:"configurationAggregatorName" yaml:"configurationAggregatorName"`
	// Provides an organization and list of regions to be aggregated.
	OrganizationAggregationSource interface{} `json:"organizationAggregationSource" yaml:"organizationAggregationSource"`
	// An array of tag object.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Config::ConfigurationRecorder`.
//
// The AWS::Config::ConfigurationRecorder resource describes the AWS resource types for which AWS Config records configuration changes. The configuration recorder stores the configurations of the supported resources in your account as configuration items.
//
// > To enable AWS Config , you must create a configuration recorder and a delivery channel. AWS Config uses the delivery channel to deliver the configuration changes to your Amazon S3 bucket or Amazon SNS topic. For more information, see [AWS::Config::DeliveryChannel](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html) .
//
// AWS CloudFormation starts the recorder as soon as the delivery channel is available.
//
// To stop the recorder and delete it, delete the configuration recorder from your stack. To stop the recorder without deleting it, call the [StopConfigurationRecorder](https://docs.aws.amazon.com/config/latest/APIReference/API_StopConfigurationRecorder.html) action of the AWS Config API directly.
//
// For more information, see [Configuration Recorder](https://docs.aws.amazon.com/config/latest/developerguide/config-concepts.html#config-recorder) in the AWS Config Developer Guide.
//
// TODO: EXAMPLE
//
type CfnConfigurationRecorder interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() awscdk.ConstructNode
	RecordingGroup() interface{}
	SetRecordingGroup(val interface{})
	Ref() *string
	RoleArn() *string
	SetRoleArn(val *string)
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnConfigurationRecorder
type jsiiProxy_CfnConfigurationRecorder struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnConfigurationRecorder) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationRecorder) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationRecorder) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationRecorder) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationRecorder) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationRecorder) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationRecorder) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationRecorder) RecordingGroup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordingGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationRecorder) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationRecorder) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationRecorder) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfigurationRecorder) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Config::ConfigurationRecorder`.
func NewCfnConfigurationRecorder(scope awscdk.Construct, id *string, props *CfnConfigurationRecorderProps) CfnConfigurationRecorder {
	_init_.Initialize()

	j := jsiiProxy_CfnConfigurationRecorder{}

	_jsii_.Create(
		"monocdk.aws_config.CfnConfigurationRecorder",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Config::ConfigurationRecorder`.
func NewCfnConfigurationRecorder_Override(c CfnConfigurationRecorder, scope awscdk.Construct, id *string, props *CfnConfigurationRecorderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_config.CfnConfigurationRecorder",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnConfigurationRecorder) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnConfigurationRecorder) SetRecordingGroup(val interface{}) {
	_jsii_.Set(
		j,
		"recordingGroup",
		val,
	)
}

func (j *jsiiProxy_CfnConfigurationRecorder) SetRoleArn(val *string) {
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
// Experimental.
func CfnConfigurationRecorder_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnConfigurationRecorder",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnConfigurationRecorder_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnConfigurationRecorder",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnConfigurationRecorder_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnConfigurationRecorder",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConfigurationRecorder_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.CfnConfigurationRecorder",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnConfigurationRecorder) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnConfigurationRecorder) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnConfigurationRecorder) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
// Experimental.
func (c *jsiiProxy_CfnConfigurationRecorder) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnConfigurationRecorder) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnConfigurationRecorder) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

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
func (c *jsiiProxy_CfnConfigurationRecorder) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnConfigurationRecorder) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnConfigurationRecorder) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnConfigurationRecorder) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnConfigurationRecorder) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnConfigurationRecorder) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnConfigurationRecorder) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnConfigurationRecorder) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnConfigurationRecorder) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnConfigurationRecorder) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnConfigurationRecorder) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnConfigurationRecorder) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnConfigurationRecorder) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnConfigurationRecorder) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnConfigurationRecorder) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies the types of AWS resource for which AWS Config records configuration changes.
//
// In the recording group, you specify whether all supported types or specific types of resources are recorded.
//
// By default, AWS Config records configuration changes for all supported types of regional resources that AWS Config discovers in the region in which it is running. Regional resources are tied to a region and can be used only in that region. Examples of regional resources are EC2 instances and EBS volumes.
//
// You can also have AWS Config record configuration changes for supported types of global resources (for example, IAM resources). Global resources are not tied to an individual region and can be used in all regions.
//
// > The configuration details for any global resource are the same in all regions. If you customize AWS Config in multiple regions to record global resources, it will create multiple configuration items each time a global resource changes: one configuration item for each region. These configuration items will contain identical data. To prevent duplicate configuration items, you should consider customizing AWS Config in only one region to record global resources, unless you want the configuration items to be available in multiple regions.
//
// If you don't want AWS Config to record all resources, you can specify which types of resources it will record with the `resourceTypes` parameter.
//
// For a list of supported resource types, see [Supported Resource Types](https://docs.aws.amazon.com/config/latest/developerguide/resource-config-reference.html#supported-resources) .
//
// For more information, see [Selecting Which Resources AWS Config Records](https://docs.aws.amazon.com/config/latest/developerguide/select-resources.html) .
//
// TODO: EXAMPLE
//
type CfnConfigurationRecorder_RecordingGroupProperty struct {
	// Specifies whether AWS Config records configuration changes for every supported type of regional resource.
	//
	// If you set this option to `true` , when AWS Config adds support for a new type of regional resource, it starts recording resources of that type automatically.
	//
	// If you set this option to `true` , you cannot enumerate a list of `resourceTypes` .
	AllSupported interface{} `json:"allSupported" yaml:"allSupported"`
	// Specifies whether AWS Config includes all supported types of global resources (for example, IAM resources) with the resources that it records.
	//
	// Before you can set this option to `true` , you must set the `AllSupported` option to `true` .
	//
	// If you set this option to `true` , when AWS Config adds support for a new type of global resource, it starts recording resources of that type automatically.
	//
	// The configuration details for any global resource are the same in all regions. To prevent duplicate configuration items, you should consider customizing AWS Config in only one region to record global resources.
	IncludeGlobalResourceTypes interface{} `json:"includeGlobalResourceTypes" yaml:"includeGlobalResourceTypes"`
	// A comma-separated list that specifies the types of AWS resources for which AWS Config records configuration changes (for example, `AWS::EC2::Instance` or `AWS::CloudTrail::Trail` ).
	//
	// To record all configuration changes, you must set the `AllSupported` option to `false` .
	//
	// If you set this option to `true` , when AWS Config adds support for a new type of resource, it will not record resources of that type unless you manually add that type to your recording group.
	//
	// For a list of valid `resourceTypes` values, see the *resourceType Value* column in [Supported AWS Resource Types](https://docs.aws.amazon.com/config/latest/developerguide/resource-config-reference.html#supported-resources) .
	ResourceTypes *[]*string `json:"resourceTypes" yaml:"resourceTypes"`
}

// Properties for defining a `CfnConfigurationRecorder`.
//
// TODO: EXAMPLE
//
type CfnConfigurationRecorderProps struct {
	// The Amazon Resource Name (ARN) of the IAM (IAM) role that is used to make read or write requests to the delivery channel that you specify and to get configuration details for supported AWS resources.
	//
	// For more information, see [Permissions for the IAM Role Assigned](https://docs.aws.amazon.com/config/latest/developerguide/iamrole-permissions.html) to AWS Config in the AWS Config Developer Guide.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// A name for the configuration recorder.
	//
	// If you don't specify a name, AWS CloudFormation CloudFormation generates a unique physical ID and uses that ID for the configuration recorder name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > After you create a configuration recorder, you cannot rename it. If you don't want a name that AWS CloudFormation generates, specify a value for this property.
	//
	// Updates are not supported.
	Name *string `json:"name" yaml:"name"`
	// Indicates whether to record configurations for all supported resources or for a list of resource types.
	//
	// The resource types that you list must be supported by AWS Config .
	RecordingGroup interface{} `json:"recordingGroup" yaml:"recordingGroup"`
}

// A CloudFormation `AWS::Config::ConformancePack`.
//
// A conformance pack is a collection of AWS Config rules and remediation actions that can be easily deployed in an account and a region. ConformancePack creates a service linked role in your account. The service linked role is created only when the role does not exist in your account.
//
// TODO: EXAMPLE
//
type CfnConformancePack interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ConformancePackInputParameters() interface{}
	SetConformancePackInputParameters(val interface{})
	ConformancePackName() *string
	SetConformancePackName(val *string)
	CreationStack() *[]*string
	DeliveryS3Bucket() *string
	SetDeliveryS3Bucket(val *string)
	DeliveryS3KeyPrefix() *string
	SetDeliveryS3KeyPrefix(val *string)
	LogicalId() *string
	Node() awscdk.ConstructNode
	Ref() *string
	Stack() awscdk.Stack
	TemplateBody() *string
	SetTemplateBody(val *string)
	TemplateS3Uri() *string
	SetTemplateS3Uri(val *string)
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnConformancePack
type jsiiProxy_CfnConformancePack struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnConformancePack) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConformancePack) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConformancePack) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConformancePack) ConformancePackInputParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conformancePackInputParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConformancePack) ConformancePackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conformancePackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConformancePack) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConformancePack) DeliveryS3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryS3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConformancePack) DeliveryS3KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryS3KeyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConformancePack) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConformancePack) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConformancePack) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConformancePack) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConformancePack) TemplateBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConformancePack) TemplateS3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateS3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConformancePack) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Config::ConformancePack`.
func NewCfnConformancePack(scope awscdk.Construct, id *string, props *CfnConformancePackProps) CfnConformancePack {
	_init_.Initialize()

	j := jsiiProxy_CfnConformancePack{}

	_jsii_.Create(
		"monocdk.aws_config.CfnConformancePack",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Config::ConformancePack`.
func NewCfnConformancePack_Override(c CfnConformancePack, scope awscdk.Construct, id *string, props *CfnConformancePackProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_config.CfnConformancePack",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnConformancePack) SetConformancePackInputParameters(val interface{}) {
	_jsii_.Set(
		j,
		"conformancePackInputParameters",
		val,
	)
}

func (j *jsiiProxy_CfnConformancePack) SetConformancePackName(val *string) {
	_jsii_.Set(
		j,
		"conformancePackName",
		val,
	)
}

func (j *jsiiProxy_CfnConformancePack) SetDeliveryS3Bucket(val *string) {
	_jsii_.Set(
		j,
		"deliveryS3Bucket",
		val,
	)
}

func (j *jsiiProxy_CfnConformancePack) SetDeliveryS3KeyPrefix(val *string) {
	_jsii_.Set(
		j,
		"deliveryS3KeyPrefix",
		val,
	)
}

func (j *jsiiProxy_CfnConformancePack) SetTemplateBody(val *string) {
	_jsii_.Set(
		j,
		"templateBody",
		val,
	)
}

func (j *jsiiProxy_CfnConformancePack) SetTemplateS3Uri(val *string) {
	_jsii_.Set(
		j,
		"templateS3Uri",
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
func CfnConformancePack_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnConformancePack",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnConformancePack_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnConformancePack",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnConformancePack_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnConformancePack",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConformancePack_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.CfnConformancePack",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnConformancePack) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnConformancePack) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnConformancePack) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
// Experimental.
func (c *jsiiProxy_CfnConformancePack) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnConformancePack) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnConformancePack) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

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
func (c *jsiiProxy_CfnConformancePack) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnConformancePack) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnConformancePack) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnConformancePack) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnConformancePack) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnConformancePack) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnConformancePack) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnConformancePack) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnConformancePack) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnConformancePack) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnConformancePack) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnConformancePack) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnConformancePack) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnConformancePack) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnConformancePack) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Input parameters in the form of key-value pairs for the conformance pack, both of which you define.
//
// Keys can have a maximum character length of 255 characters, and values can have a maximum length of 4096 characters.
//
// TODO: EXAMPLE
//
type CfnConformancePack_ConformancePackInputParameterProperty struct {
	// One part of a key-value pair.
	ParameterName *string `json:"parameterName" yaml:"parameterName"`
	// Another part of the key-value pair.
	ParameterValue *string `json:"parameterValue" yaml:"parameterValue"`
}

// Properties for defining a `CfnConformancePack`.
//
// TODO: EXAMPLE
//
type CfnConformancePackProps struct {
	// Name of the conformance pack you want to create.
	ConformancePackName *string `json:"conformancePackName" yaml:"conformancePackName"`
	// A list of ConformancePackInputParameter objects.
	ConformancePackInputParameters interface{} `json:"conformancePackInputParameters" yaml:"conformancePackInputParameters"`
	// The name of the Amazon S3 bucket where AWS Config stores conformance pack templates.
	DeliveryS3Bucket *string `json:"deliveryS3Bucket" yaml:"deliveryS3Bucket"`
	// The prefix for the Amazon S3 bucket.
	DeliveryS3KeyPrefix *string `json:"deliveryS3KeyPrefix" yaml:"deliveryS3KeyPrefix"`
	// A string containing full conformance pack template body.
	//
	// Structure containing the template body with a minimum length of 1 byte and a maximum length of 51,200 bytes.
	//
	// > You can only use a YAML template with two resource types: config rule ( `AWS::Config::ConfigRule` ) and a remediation action ( `AWS::Config::RemediationConfiguration` ).
	TemplateBody *string `json:"templateBody" yaml:"templateBody"`
	// Location of file containing the template body (s3://bucketname/prefix).
	//
	// The uri must point to the conformance pack template (max size: 300 KB) that is located in an Amazon S3 bucket.
	//
	// > You must have access to read Amazon S3 bucket.
	TemplateS3Uri *string `json:"templateS3Uri" yaml:"templateS3Uri"`
}

// A CloudFormation `AWS::Config::DeliveryChannel`.
//
// Specifies a delivery channel object to deliver configuration information to an Amazon S3 bucket and Amazon SNS topic.
//
// Before you can create a delivery channel, you must create a configuration recorder. You can use this action to change the Amazon S3 bucket or an Amazon SNS topic of the existing delivery channel. To change the Amazon S3 bucket or an Amazon SNS topic, call this action and specify the changed values for the S3 bucket and the SNS topic. If you specify a different value for either the S3 bucket or the SNS topic, this action will keep the existing value for the parameter that is not changed.
//
// > In the China (Beijing) Region, when you call this action, the Amazon S3 bucket must also be in the China (Beijing) Region. In all the other regions, AWS Config supports cross-region and cross-account delivery channels.
//
// You can have only one delivery channel per region per AWS account, and the delivery channel is required to use AWS Config .
//
// > AWS Config does not support the delivery channel to an Amazon S3 bucket bucket where object lock is enabled. For more information, see [How S3 Object Lock works](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-overview.html) .
//
// When you create the delivery channel, you can specify; how often AWS Config delivers configuration snapshots to your Amazon S3 bucket (for example, 24 hours), the S3 bucket to which AWS Config sends configuration snapshots and configuration history files, and the Amazon SNS topic to which AWS Config sends notifications about configuration changes, such as updated resources, AWS Config rule evaluations, and when AWS Config delivers the configuration snapshot to your S3 bucket. For more information, see [Deliver Configuration Items](https://docs.aws.amazon.com/config/latest/developerguide/how-does-config-work.html#delivery-channel) in the AWS Config Developer Guide.
//
// > To enable AWS Config , you must create a configuration recorder and a delivery channel. If you want to create the resources separately, you must create a configuration recorder before you can create a delivery channel. AWS Config uses the configuration recorder to capture configuration changes to your resources. For more information, see [AWS::Config::ConfigurationRecorder](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationrecorder.html) .
//
// For more information, see [Managing the Delivery Channel](https://docs.aws.amazon.com/config/latest/developerguide/manage-delivery-channel.html) in the AWS Config Developer Guide.
//
// TODO: EXAMPLE
//
type CfnDeliveryChannel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ConfigSnapshotDeliveryProperties() interface{}
	SetConfigSnapshotDeliveryProperties(val interface{})
	CreationStack() *[]*string
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() awscdk.ConstructNode
	Ref() *string
	S3BucketName() *string
	SetS3BucketName(val *string)
	S3KeyPrefix() *string
	SetS3KeyPrefix(val *string)
	S3KmsKeyArn() *string
	SetS3KmsKeyArn(val *string)
	SnsTopicArn() *string
	SetSnsTopicArn(val *string)
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDeliveryChannel
type jsiiProxy_CfnDeliveryChannel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDeliveryChannel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryChannel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryChannel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryChannel) ConfigSnapshotDeliveryProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configSnapshotDeliveryProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryChannel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryChannel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryChannel) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryChannel) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryChannel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryChannel) S3BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryChannel) S3KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryChannel) S3KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryChannel) SnsTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryChannel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryChannel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Config::DeliveryChannel`.
func NewCfnDeliveryChannel(scope awscdk.Construct, id *string, props *CfnDeliveryChannelProps) CfnDeliveryChannel {
	_init_.Initialize()

	j := jsiiProxy_CfnDeliveryChannel{}

	_jsii_.Create(
		"monocdk.aws_config.CfnDeliveryChannel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Config::DeliveryChannel`.
func NewCfnDeliveryChannel_Override(c CfnDeliveryChannel, scope awscdk.Construct, id *string, props *CfnDeliveryChannelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_config.CfnDeliveryChannel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDeliveryChannel) SetConfigSnapshotDeliveryProperties(val interface{}) {
	_jsii_.Set(
		j,
		"configSnapshotDeliveryProperties",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryChannel) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryChannel) SetS3BucketName(val *string) {
	_jsii_.Set(
		j,
		"s3BucketName",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryChannel) SetS3KeyPrefix(val *string) {
	_jsii_.Set(
		j,
		"s3KeyPrefix",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryChannel) SetS3KmsKeyArn(val *string) {
	_jsii_.Set(
		j,
		"s3KmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryChannel) SetSnsTopicArn(val *string) {
	_jsii_.Set(
		j,
		"snsTopicArn",
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
func CfnDeliveryChannel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnDeliveryChannel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDeliveryChannel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnDeliveryChannel",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDeliveryChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnDeliveryChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDeliveryChannel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.CfnDeliveryChannel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnDeliveryChannel) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnDeliveryChannel) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnDeliveryChannel) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
// Experimental.
func (c *jsiiProxy_CfnDeliveryChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnDeliveryChannel) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnDeliveryChannel) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

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
func (c *jsiiProxy_CfnDeliveryChannel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnDeliveryChannel) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnDeliveryChannel) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnDeliveryChannel) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnDeliveryChannel) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnDeliveryChannel) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnDeliveryChannel) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnDeliveryChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnDeliveryChannel) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDeliveryChannel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnDeliveryChannel) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnDeliveryChannel) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnDeliveryChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnDeliveryChannel) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnDeliveryChannel) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Provides options for how often AWS Config delivers configuration snapshots to the Amazon S3 bucket in your delivery channel.
//
// > If you want to create a rule that triggers evaluations for your resources when AWS Config delivers the configuration snapshot, see the following:
//
// The frequency for a rule that triggers evaluations for your resources when AWS Config delivers the configuration snapshot is set by one of two values, depending on which is less frequent:
//
// - The value for the `deliveryFrequency` parameter within the delivery channel configuration, which sets how often AWS Config delivers configuration snapshots. This value also sets how often AWS Config invokes evaluations for AWS Config rules.
// - The value for the `MaximumExecutionFrequency` parameter, which sets the maximum frequency with which AWS Config invokes evaluations for the rule. For more information, see [ConfigRule](https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigRule.html) .
//
// If the `deliveryFrequency` value is less frequent than the `MaximumExecutionFrequency` value for a rule, AWS Config invokes the rule only as often as the `deliveryFrequency` value.
//
// - For example, you want your rule to run evaluations when AWS Config delivers the configuration snapshot.
// - You specify the `MaximumExecutionFrequency` value for `Six_Hours` .
// - You then specify the delivery channel `deliveryFrequency` value for `TwentyFour_Hours` .
// - Because the value for `deliveryFrequency` is less frequent than `MaximumExecutionFrequency` , AWS Config invokes evaluations for the rule every 24 hours.
//
// You should set the `MaximumExecutionFrequency` value to be at least as frequent as the `deliveryFrequency` value. You can view the `deliveryFrequency` value by using the `DescribeDeliveryChannnels` action.
//
// To update the `deliveryFrequency` with which AWS Config delivers your configuration snapshots, use the `PutDeliveryChannel` action.
//
// TODO: EXAMPLE
//
type CfnDeliveryChannel_ConfigSnapshotDeliveryPropertiesProperty struct {
	// The frequency with which AWS Config delivers configuration snapshots.
	DeliveryFrequency *string `json:"deliveryFrequency" yaml:"deliveryFrequency"`
}

// Properties for defining a `CfnDeliveryChannel`.
//
// TODO: EXAMPLE
//
type CfnDeliveryChannelProps struct {
	// The name of the Amazon S3 bucket to which AWS Config delivers configuration snapshots and configuration history files.
	//
	// If you specify a bucket that belongs to another AWS account , that bucket must have policies that grant access permissions to AWS Config . For more information, see [Permissions for the Amazon S3 Bucket](https://docs.aws.amazon.com/config/latest/developerguide/s3-bucket-policy.html) in the AWS Config Developer Guide.
	S3BucketName *string `json:"s3BucketName" yaml:"s3BucketName"`
	// The options for how often AWS Config delivers configuration snapshots to the Amazon S3 bucket.
	ConfigSnapshotDeliveryProperties interface{} `json:"configSnapshotDeliveryProperties" yaml:"configSnapshotDeliveryProperties"`
	// A name for the delivery channel.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the delivery channel name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// Updates are not supported. To change the name, you must run two separate updates. In the first update, delete this resource, and then recreate it with a new name in the second update.
	Name *string `json:"name" yaml:"name"`
	// The prefix for the specified Amazon S3 bucket.
	S3KeyPrefix *string `json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
	// The Amazon Resource Name (ARN) of the AWS Key Management Service ( AWS KMS ) AWS KMS key (KMS key) used to encrypt objects delivered by AWS Config .
	//
	// Must belong to the same Region as the destination S3 bucket.
	S3KmsKeyArn *string `json:"s3KmsKeyArn" yaml:"s3KmsKeyArn"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to which AWS Config sends notifications about configuration changes.
	//
	// If you choose a topic from another account, the topic must have policies that grant access permissions to AWS Config . For more information, see [Permissions for the Amazon SNS Topic](https://docs.aws.amazon.com/config/latest/developerguide/sns-topic-policy.html) in the AWS Config Developer Guide.
	SnsTopicArn *string `json:"snsTopicArn" yaml:"snsTopicArn"`
}

// A CloudFormation `AWS::Config::OrganizationConfigRule`.
//
// An organization config rule that has information about config rules that AWS Config creates in member accounts.Only a master account and a delegated administrator can create or update an organization config rule.
//
// `OrganizationConfigRule` resource enables organization service access through `EnableAWSServiceAccess` action and creates a service linked role in the master account of your organization. The service linked role is created only when the role does not exist in the master account. AWS Config verifies the existence of role with `GetRole` action.
//
// When creating custom organization config rules using a centralized Lambda function, you will need to allow Lambda permissions to sub-accounts and you will need to create an IAM role will to pass to the Lambda function. For more information, see [How to Centrally Manage AWS Config Rules across Multiple AWS Accounts](https://docs.aws.amazon.com/devops/how-to-centrally-manage-aws-config-rules-across-multiple-aws-accounts/) .
//
// TODO: EXAMPLE
//
type CfnOrganizationConfigRule interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	ExcludedAccounts() *[]*string
	SetExcludedAccounts(val *[]*string)
	LogicalId() *string
	Node() awscdk.ConstructNode
	OrganizationConfigRuleName() *string
	SetOrganizationConfigRuleName(val *string)
	OrganizationCustomCodeRuleMetadata() interface{}
	SetOrganizationCustomCodeRuleMetadata(val interface{})
	OrganizationCustomRuleMetadata() interface{}
	SetOrganizationCustomRuleMetadata(val interface{})
	OrganizationManagedRuleMetadata() interface{}
	SetOrganizationManagedRuleMetadata(val interface{})
	Ref() *string
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnOrganizationConfigRule
type jsiiProxy_CfnOrganizationConfigRule struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnOrganizationConfigRule) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConfigRule) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConfigRule) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConfigRule) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConfigRule) ExcludedAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConfigRule) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConfigRule) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConfigRule) OrganizationConfigRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationConfigRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConfigRule) OrganizationCustomCodeRuleMetadata() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"organizationCustomCodeRuleMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConfigRule) OrganizationCustomRuleMetadata() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"organizationCustomRuleMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConfigRule) OrganizationManagedRuleMetadata() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"organizationManagedRuleMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConfigRule) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConfigRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConfigRule) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Config::OrganizationConfigRule`.
func NewCfnOrganizationConfigRule(scope awscdk.Construct, id *string, props *CfnOrganizationConfigRuleProps) CfnOrganizationConfigRule {
	_init_.Initialize()

	j := jsiiProxy_CfnOrganizationConfigRule{}

	_jsii_.Create(
		"monocdk.aws_config.CfnOrganizationConfigRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Config::OrganizationConfigRule`.
func NewCfnOrganizationConfigRule_Override(c CfnOrganizationConfigRule, scope awscdk.Construct, id *string, props *CfnOrganizationConfigRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_config.CfnOrganizationConfigRule",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnOrganizationConfigRule) SetExcludedAccounts(val *[]*string) {
	_jsii_.Set(
		j,
		"excludedAccounts",
		val,
	)
}

func (j *jsiiProxy_CfnOrganizationConfigRule) SetOrganizationConfigRuleName(val *string) {
	_jsii_.Set(
		j,
		"organizationConfigRuleName",
		val,
	)
}

func (j *jsiiProxy_CfnOrganizationConfigRule) SetOrganizationCustomCodeRuleMetadata(val interface{}) {
	_jsii_.Set(
		j,
		"organizationCustomCodeRuleMetadata",
		val,
	)
}

func (j *jsiiProxy_CfnOrganizationConfigRule) SetOrganizationCustomRuleMetadata(val interface{}) {
	_jsii_.Set(
		j,
		"organizationCustomRuleMetadata",
		val,
	)
}

func (j *jsiiProxy_CfnOrganizationConfigRule) SetOrganizationManagedRuleMetadata(val interface{}) {
	_jsii_.Set(
		j,
		"organizationManagedRuleMetadata",
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
func CfnOrganizationConfigRule_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnOrganizationConfigRule",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnOrganizationConfigRule_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnOrganizationConfigRule",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnOrganizationConfigRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnOrganizationConfigRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOrganizationConfigRule_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.CfnOrganizationConfigRule",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConfigRule) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConfigRule) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnOrganizationConfigRule) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
// Experimental.
func (c *jsiiProxy_CfnOrganizationConfigRule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConfigRule) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConfigRule) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

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
func (c *jsiiProxy_CfnOrganizationConfigRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConfigRule) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnOrganizationConfigRule) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnOrganizationConfigRule) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConfigRule) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConfigRule) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConfigRule) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConfigRule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConfigRule) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnOrganizationConfigRule) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConfigRule) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConfigRule) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnOrganizationConfigRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConfigRule) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnOrganizationConfigRule) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnOrganizationConfigRule_OrganizationCustomCodeRuleMetadataProperty struct {
	// `CfnOrganizationConfigRule.OrganizationCustomCodeRuleMetadataProperty.CodeText`.
	CodeText *string `json:"codeText" yaml:"codeText"`
	// `CfnOrganizationConfigRule.OrganizationCustomCodeRuleMetadataProperty.Runtime`.
	Runtime *string `json:"runtime" yaml:"runtime"`
	// `CfnOrganizationConfigRule.OrganizationCustomCodeRuleMetadataProperty.DebugLogDeliveryAccounts`.
	DebugLogDeliveryAccounts *[]*string `json:"debugLogDeliveryAccounts" yaml:"debugLogDeliveryAccounts"`
	// `CfnOrganizationConfigRule.OrganizationCustomCodeRuleMetadataProperty.Description`.
	Description *string `json:"description" yaml:"description"`
	// `CfnOrganizationConfigRule.OrganizationCustomCodeRuleMetadataProperty.InputParameters`.
	InputParameters *string `json:"inputParameters" yaml:"inputParameters"`
	// `CfnOrganizationConfigRule.OrganizationCustomCodeRuleMetadataProperty.MaximumExecutionFrequency`.
	MaximumExecutionFrequency *string `json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// `CfnOrganizationConfigRule.OrganizationCustomCodeRuleMetadataProperty.OrganizationConfigRuleTriggerTypes`.
	OrganizationConfigRuleTriggerTypes *[]*string `json:"organizationConfigRuleTriggerTypes" yaml:"organizationConfigRuleTriggerTypes"`
	// `CfnOrganizationConfigRule.OrganizationCustomCodeRuleMetadataProperty.ResourceIdScope`.
	ResourceIdScope *string `json:"resourceIdScope" yaml:"resourceIdScope"`
	// `CfnOrganizationConfigRule.OrganizationCustomCodeRuleMetadataProperty.ResourceTypesScope`.
	ResourceTypesScope *[]*string `json:"resourceTypesScope" yaml:"resourceTypesScope"`
	// `CfnOrganizationConfigRule.OrganizationCustomCodeRuleMetadataProperty.TagKeyScope`.
	TagKeyScope *string `json:"tagKeyScope" yaml:"tagKeyScope"`
	// `CfnOrganizationConfigRule.OrganizationCustomCodeRuleMetadataProperty.TagValueScope`.
	TagValueScope *string `json:"tagValueScope" yaml:"tagValueScope"`
}

// An object that specifies organization custom rule metadata such as resource type, resource ID of AWS resource, Lambda function ARN, and organization trigger types that trigger AWS Config to evaluate your AWS resources against a rule.
//
// It also provides the frequency with which you want AWS Config to run evaluations for the rule if the trigger type is periodic.
//
// TODO: EXAMPLE
//
type CfnOrganizationConfigRule_OrganizationCustomRuleMetadataProperty struct {
	// The lambda function ARN.
	LambdaFunctionArn *string `json:"lambdaFunctionArn" yaml:"lambdaFunctionArn"`
	// The type of notification that triggers AWS Config to run an evaluation for a rule.
	//
	// You can specify the following notification types:
	//
	// - `ConfigurationItemChangeNotification` - Triggers an evaluation when AWS Config delivers a configuration item as a result of a resource change.
	// - `OversizedConfigurationItemChangeNotification` - Triggers an evaluation when AWS Config delivers an oversized configuration item. AWS Config may generate this notification type when a resource changes and the notification exceeds the maximum size allowed by Amazon SNS.
	// - `ScheduledNotification` - Triggers a periodic evaluation at the frequency specified for `MaximumExecutionFrequency` .
	OrganizationConfigRuleTriggerTypes *[]*string `json:"organizationConfigRuleTriggerTypes" yaml:"organizationConfigRuleTriggerTypes"`
	// The description that you provide for organization config rule.
	Description *string `json:"description" yaml:"description"`
	// A string, in JSON format, that is passed to organization config rule Lambda function.
	InputParameters *string `json:"inputParameters" yaml:"inputParameters"`
	// The maximum frequency with which AWS Config runs evaluations for a rule.
	//
	// Your custom rule is triggered when AWS Config delivers the configuration snapshot. For more information, see `ConfigSnapshotDeliveryProperties` .
	//
	// > By default, rules with a periodic trigger are evaluated every 24 hours. To change the frequency, specify a valid value for the `MaximumExecutionFrequency` parameter.
	MaximumExecutionFrequency *string `json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// The ID of the AWS resource that was evaluated.
	ResourceIdScope *string `json:"resourceIdScope" yaml:"resourceIdScope"`
	// The type of the AWS resource that was evaluated.
	ResourceTypesScope *[]*string `json:"resourceTypesScope" yaml:"resourceTypesScope"`
	// One part of a key-value pair that make up a tag.
	//
	// A key is a general label that acts like a category for more specific tag values.
	TagKeyScope *string `json:"tagKeyScope" yaml:"tagKeyScope"`
	// The optional part of a key-value pair that make up a tag.
	//
	// A value acts as a descriptor within a tag category (key).
	TagValueScope *string `json:"tagValueScope" yaml:"tagValueScope"`
}

// An object that specifies organization managed rule metadata such as resource type and ID of AWS resource along with the rule identifier.
//
// It also provides the frequency with which you want AWS Config to run evaluations for the rule if the trigger type is periodic.
//
// TODO: EXAMPLE
//
type CfnOrganizationConfigRule_OrganizationManagedRuleMetadataProperty struct {
	// For organization config managed rules, a predefined identifier from a list.
	//
	// For example, `IAM_PASSWORD_POLICY` is a managed rule. To reference a managed rule, see [Using AWS Config managed rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html) .
	RuleIdentifier *string `json:"ruleIdentifier" yaml:"ruleIdentifier"`
	// The description that you provide for organization config rule.
	Description *string `json:"description" yaml:"description"`
	// A string, in JSON format, that is passed to organization config rule Lambda function.
	InputParameters *string `json:"inputParameters" yaml:"inputParameters"`
	// The maximum frequency with which AWS Config runs evaluations for a rule.
	//
	// You are using an AWS Config managed rule that is triggered at a periodic frequency.
	//
	// > By default, rules with a periodic trigger are evaluated every 24 hours. To change the frequency, specify a valid value for the `MaximumExecutionFrequency` parameter.
	MaximumExecutionFrequency *string `json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// The ID of the AWS resource that was evaluated.
	ResourceIdScope *string `json:"resourceIdScope" yaml:"resourceIdScope"`
	// The type of the AWS resource that was evaluated.
	ResourceTypesScope *[]*string `json:"resourceTypesScope" yaml:"resourceTypesScope"`
	// One part of a key-value pair that make up a tag.
	//
	// A key is a general label that acts like a category for more specific tag values.
	TagKeyScope *string `json:"tagKeyScope" yaml:"tagKeyScope"`
	// The optional part of a key-value pair that make up a tag.
	//
	// A value acts as a descriptor within a tag category (key).
	TagValueScope *string `json:"tagValueScope" yaml:"tagValueScope"`
}

// Properties for defining a `CfnOrganizationConfigRule`.
//
// TODO: EXAMPLE
//
type CfnOrganizationConfigRuleProps struct {
	// The name that you assign to organization config rule.
	OrganizationConfigRuleName *string `json:"organizationConfigRuleName" yaml:"organizationConfigRuleName"`
	// A comma-separated list of accounts excluded from organization config rule.
	ExcludedAccounts *[]*string `json:"excludedAccounts" yaml:"excludedAccounts"`
	// `AWS::Config::OrganizationConfigRule.OrganizationCustomCodeRuleMetadata`.
	OrganizationCustomCodeRuleMetadata interface{} `json:"organizationCustomCodeRuleMetadata" yaml:"organizationCustomCodeRuleMetadata"`
	// An `OrganizationCustomRuleMetadata` object.
	OrganizationCustomRuleMetadata interface{} `json:"organizationCustomRuleMetadata" yaml:"organizationCustomRuleMetadata"`
	// An `OrganizationManagedRuleMetadata` object.
	OrganizationManagedRuleMetadata interface{} `json:"organizationManagedRuleMetadata" yaml:"organizationManagedRuleMetadata"`
}

// A CloudFormation `AWS::Config::OrganizationConformancePack`.
//
// OrganizationConformancePack deploys conformance packs across member accounts in an AWS Organizations . OrganizationConformancePack enables organization service access for `config-multiaccountsetup.amazonaws.com` through the `EnableAWSServiceAccess` action and creates a service linked role in the master account of your organization. The service linked role is created only when the role does not exist in the master account.
//
// TODO: EXAMPLE
//
type CfnOrganizationConformancePack interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ConformancePackInputParameters() interface{}
	SetConformancePackInputParameters(val interface{})
	CreationStack() *[]*string
	DeliveryS3Bucket() *string
	SetDeliveryS3Bucket(val *string)
	DeliveryS3KeyPrefix() *string
	SetDeliveryS3KeyPrefix(val *string)
	ExcludedAccounts() *[]*string
	SetExcludedAccounts(val *[]*string)
	LogicalId() *string
	Node() awscdk.ConstructNode
	OrganizationConformancePackName() *string
	SetOrganizationConformancePackName(val *string)
	Ref() *string
	Stack() awscdk.Stack
	TemplateBody() *string
	SetTemplateBody(val *string)
	TemplateS3Uri() *string
	SetTemplateS3Uri(val *string)
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnOrganizationConformancePack
type jsiiProxy_CfnOrganizationConformancePack struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnOrganizationConformancePack) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConformancePack) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConformancePack) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConformancePack) ConformancePackInputParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conformancePackInputParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConformancePack) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConformancePack) DeliveryS3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryS3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConformancePack) DeliveryS3KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryS3KeyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConformancePack) ExcludedAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConformancePack) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConformancePack) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConformancePack) OrganizationConformancePackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationConformancePackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConformancePack) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConformancePack) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConformancePack) TemplateBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConformancePack) TemplateS3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateS3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConformancePack) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Config::OrganizationConformancePack`.
func NewCfnOrganizationConformancePack(scope awscdk.Construct, id *string, props *CfnOrganizationConformancePackProps) CfnOrganizationConformancePack {
	_init_.Initialize()

	j := jsiiProxy_CfnOrganizationConformancePack{}

	_jsii_.Create(
		"monocdk.aws_config.CfnOrganizationConformancePack",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Config::OrganizationConformancePack`.
func NewCfnOrganizationConformancePack_Override(c CfnOrganizationConformancePack, scope awscdk.Construct, id *string, props *CfnOrganizationConformancePackProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_config.CfnOrganizationConformancePack",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnOrganizationConformancePack) SetConformancePackInputParameters(val interface{}) {
	_jsii_.Set(
		j,
		"conformancePackInputParameters",
		val,
	)
}

func (j *jsiiProxy_CfnOrganizationConformancePack) SetDeliveryS3Bucket(val *string) {
	_jsii_.Set(
		j,
		"deliveryS3Bucket",
		val,
	)
}

func (j *jsiiProxy_CfnOrganizationConformancePack) SetDeliveryS3KeyPrefix(val *string) {
	_jsii_.Set(
		j,
		"deliveryS3KeyPrefix",
		val,
	)
}

func (j *jsiiProxy_CfnOrganizationConformancePack) SetExcludedAccounts(val *[]*string) {
	_jsii_.Set(
		j,
		"excludedAccounts",
		val,
	)
}

func (j *jsiiProxy_CfnOrganizationConformancePack) SetOrganizationConformancePackName(val *string) {
	_jsii_.Set(
		j,
		"organizationConformancePackName",
		val,
	)
}

func (j *jsiiProxy_CfnOrganizationConformancePack) SetTemplateBody(val *string) {
	_jsii_.Set(
		j,
		"templateBody",
		val,
	)
}

func (j *jsiiProxy_CfnOrganizationConformancePack) SetTemplateS3Uri(val *string) {
	_jsii_.Set(
		j,
		"templateS3Uri",
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
func CfnOrganizationConformancePack_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnOrganizationConformancePack",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnOrganizationConformancePack_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnOrganizationConformancePack",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnOrganizationConformancePack_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnOrganizationConformancePack",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOrganizationConformancePack_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.CfnOrganizationConformancePack",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConformancePack) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConformancePack) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnOrganizationConformancePack) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
// Experimental.
func (c *jsiiProxy_CfnOrganizationConformancePack) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConformancePack) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConformancePack) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

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
func (c *jsiiProxy_CfnOrganizationConformancePack) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConformancePack) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnOrganizationConformancePack) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnOrganizationConformancePack) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConformancePack) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConformancePack) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConformancePack) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConformancePack) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConformancePack) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnOrganizationConformancePack) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConformancePack) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConformancePack) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnOrganizationConformancePack) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnOrganizationConformancePack) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnOrganizationConformancePack) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Input parameters in the form of key-value pairs for the conformance pack, both of which you define.
//
// Keys can have a maximum character length of 255 characters, and values can have a maximum length of 4096 characters.
//
// TODO: EXAMPLE
//
type CfnOrganizationConformancePack_ConformancePackInputParameterProperty struct {
	// One part of a key-value pair.
	ParameterName *string `json:"parameterName" yaml:"parameterName"`
	// One part of a key-value pair.
	ParameterValue *string `json:"parameterValue" yaml:"parameterValue"`
}

// Properties for defining a `CfnOrganizationConformancePack`.
//
// TODO: EXAMPLE
//
type CfnOrganizationConformancePackProps struct {
	// The name you assign to an organization conformance pack.
	OrganizationConformancePackName *string `json:"organizationConformancePackName" yaml:"organizationConformancePackName"`
	// A list of `ConformancePackInputParameter` objects.
	ConformancePackInputParameters interface{} `json:"conformancePackInputParameters" yaml:"conformancePackInputParameters"`
	// The name of the Amazon S3 bucket where AWS Config stores conformance pack templates.
	//
	// > This field is optional.
	DeliveryS3Bucket *string `json:"deliveryS3Bucket" yaml:"deliveryS3Bucket"`
	// Any folder structure you want to add to an Amazon S3 bucket.
	//
	// > This field is optional.
	DeliveryS3KeyPrefix *string `json:"deliveryS3KeyPrefix" yaml:"deliveryS3KeyPrefix"`
	// A comma-separated list of accounts excluded from organization conformance pack.
	ExcludedAccounts *[]*string `json:"excludedAccounts" yaml:"excludedAccounts"`
	// A string containing full conformance pack template body.
	//
	// Structure containing the template body with a minimum length of 1 byte and a maximum length of 51,200 bytes.
	TemplateBody *string `json:"templateBody" yaml:"templateBody"`
	// Location of file containing the template body.
	//
	// The uri must point to the conformance pack template (max size: 300 KB).
	TemplateS3Uri *string `json:"templateS3Uri" yaml:"templateS3Uri"`
}

// A CloudFormation `AWS::Config::RemediationConfiguration`.
//
// An object that represents the details about the remediation configuration that includes the remediation action, parameters, and data to execute the action.
//
// TODO: EXAMPLE
//
type CfnRemediationConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	Automatic() interface{}
	SetAutomatic(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ConfigRuleName() *string
	SetConfigRuleName(val *string)
	CreationStack() *[]*string
	ExecutionControls() interface{}
	SetExecutionControls(val interface{})
	LogicalId() *string
	MaximumAutomaticAttempts() *float64
	SetMaximumAutomaticAttempts(val *float64)
	Node() awscdk.ConstructNode
	Parameters() interface{}
	SetParameters(val interface{})
	Ref() *string
	ResourceType() *string
	SetResourceType(val *string)
	RetryAttemptSeconds() *float64
	SetRetryAttemptSeconds(val *float64)
	Stack() awscdk.Stack
	TargetId() *string
	SetTargetId(val *string)
	TargetType() *string
	SetTargetType(val *string)
	TargetVersion() *string
	SetTargetVersion(val *string)
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnRemediationConfiguration
type jsiiProxy_CfnRemediationConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRemediationConfiguration) Automatic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automatic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRemediationConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRemediationConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRemediationConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRemediationConfiguration) ConfigRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRemediationConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRemediationConfiguration) ExecutionControls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"executionControls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRemediationConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRemediationConfiguration) MaximumAutomaticAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumAutomaticAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRemediationConfiguration) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRemediationConfiguration) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRemediationConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRemediationConfiguration) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRemediationConfiguration) RetryAttemptSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryAttemptSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRemediationConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRemediationConfiguration) TargetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRemediationConfiguration) TargetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRemediationConfiguration) TargetVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRemediationConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Config::RemediationConfiguration`.
func NewCfnRemediationConfiguration(scope awscdk.Construct, id *string, props *CfnRemediationConfigurationProps) CfnRemediationConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnRemediationConfiguration{}

	_jsii_.Create(
		"monocdk.aws_config.CfnRemediationConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Config::RemediationConfiguration`.
func NewCfnRemediationConfiguration_Override(c CfnRemediationConfiguration, scope awscdk.Construct, id *string, props *CfnRemediationConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_config.CfnRemediationConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRemediationConfiguration) SetAutomatic(val interface{}) {
	_jsii_.Set(
		j,
		"automatic",
		val,
	)
}

func (j *jsiiProxy_CfnRemediationConfiguration) SetConfigRuleName(val *string) {
	_jsii_.Set(
		j,
		"configRuleName",
		val,
	)
}

func (j *jsiiProxy_CfnRemediationConfiguration) SetExecutionControls(val interface{}) {
	_jsii_.Set(
		j,
		"executionControls",
		val,
	)
}

func (j *jsiiProxy_CfnRemediationConfiguration) SetMaximumAutomaticAttempts(val *float64) {
	_jsii_.Set(
		j,
		"maximumAutomaticAttempts",
		val,
	)
}

func (j *jsiiProxy_CfnRemediationConfiguration) SetParameters(val interface{}) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_CfnRemediationConfiguration) SetResourceType(val *string) {
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_CfnRemediationConfiguration) SetRetryAttemptSeconds(val *float64) {
	_jsii_.Set(
		j,
		"retryAttemptSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnRemediationConfiguration) SetTargetId(val *string) {
	_jsii_.Set(
		j,
		"targetId",
		val,
	)
}

func (j *jsiiProxy_CfnRemediationConfiguration) SetTargetType(val *string) {
	_jsii_.Set(
		j,
		"targetType",
		val,
	)
}

func (j *jsiiProxy_CfnRemediationConfiguration) SetTargetVersion(val *string) {
	_jsii_.Set(
		j,
		"targetVersion",
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
func CfnRemediationConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnRemediationConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnRemediationConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnRemediationConfiguration",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnRemediationConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnRemediationConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRemediationConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.CfnRemediationConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnRemediationConfiguration) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnRemediationConfiguration) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnRemediationConfiguration) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
// Experimental.
func (c *jsiiProxy_CfnRemediationConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnRemediationConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnRemediationConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

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
func (c *jsiiProxy_CfnRemediationConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnRemediationConfiguration) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnRemediationConfiguration) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnRemediationConfiguration) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnRemediationConfiguration) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnRemediationConfiguration) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnRemediationConfiguration) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnRemediationConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnRemediationConfiguration) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnRemediationConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnRemediationConfiguration) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnRemediationConfiguration) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnRemediationConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnRemediationConfiguration) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnRemediationConfiguration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// An ExecutionControls object.
//
// TODO: EXAMPLE
//
type CfnRemediationConfiguration_ExecutionControlsProperty struct {
	// A SsmControls object.
	SsmControls interface{} `json:"ssmControls" yaml:"ssmControls"`
}

// The value is either a dynamic (resource) value or a static value.
//
// You must select either a dynamic value or a static value.
//
// TODO: EXAMPLE
//
type CfnRemediationConfiguration_RemediationParameterValueProperty struct {
	// The value is dynamic and changes at run-time.
	ResourceValue interface{} `json:"resourceValue" yaml:"resourceValue"`
	// The value is static and does not change at run-time.
	StaticValue interface{} `json:"staticValue" yaml:"staticValue"`
}

// The dynamic value of the resource.
//
// TODO: EXAMPLE
//
type CfnRemediationConfiguration_ResourceValueProperty struct {
	// The value is a resource ID.
	Value *string `json:"value" yaml:"value"`
}

// AWS Systems Manager (SSM) specific remediation controls.
//
// TODO: EXAMPLE
//
type CfnRemediationConfiguration_SsmControlsProperty struct {
	// The maximum percentage of remediation actions allowed to run in parallel on the non-compliant resources for that specific rule.
	//
	// You can specify a percentage, such as 10%. The default value is 10.
	ConcurrentExecutionRatePercentage *float64 `json:"concurrentExecutionRatePercentage" yaml:"concurrentExecutionRatePercentage"`
	// The percentage of errors that are allowed before SSM stops running automations on non-compliant resources for that specific rule.
	//
	// You can specify a percentage of errors, for example 10%. If you do not specifiy a percentage, the default is 50%. For example, if you set the ErrorPercentage to 40% for 10 non-compliant resources, then SSM stops running the automations when the fifth error is received.
	ErrorPercentage *float64 `json:"errorPercentage" yaml:"errorPercentage"`
}

// The static value of the resource.
//
// TODO: EXAMPLE
//
type CfnRemediationConfiguration_StaticValueProperty struct {
	// A list of values.
	//
	// For example, the ARN of the assumed role.
	Values *[]*string `json:"values" yaml:"values"`
}

// Properties for defining a `CfnRemediationConfiguration`.
//
// TODO: EXAMPLE
//
type CfnRemediationConfigurationProps struct {
	// The name of the AWS Config rule.
	ConfigRuleName *string `json:"configRuleName" yaml:"configRuleName"`
	// Target ID is the name of the public document.
	TargetId *string `json:"targetId" yaml:"targetId"`
	// The type of the target.
	//
	// Target executes remediation. For example, SSM document.
	TargetType *string `json:"targetType" yaml:"targetType"`
	// The remediation is triggered automatically.
	Automatic interface{} `json:"automatic" yaml:"automatic"`
	// An ExecutionControls object.
	ExecutionControls interface{} `json:"executionControls" yaml:"executionControls"`
	// The maximum number of failed attempts for auto-remediation. If you do not select a number, the default is 5.
	//
	// For example, if you specify MaximumAutomaticAttempts as 5 with RetryAttemptSeconds as 50 seconds, AWS Config will put a RemediationException on your behalf for the failing resource after the 5th failed attempt within 50 seconds.
	MaximumAutomaticAttempts *float64 `json:"maximumAutomaticAttempts" yaml:"maximumAutomaticAttempts"`
	// An object of the RemediationParameterValue.
	//
	// > The type is a map of strings to RemediationParameterValue.
	Parameters interface{} `json:"parameters" yaml:"parameters"`
	// The type of a resource.
	ResourceType *string `json:"resourceType" yaml:"resourceType"`
	// Maximum time in seconds that AWS Config runs auto-remediation.
	//
	// If you do not select a number, the default is 60 seconds.
	//
	// For example, if you specify RetryAttemptSeconds as 50 seconds and MaximumAutomaticAttempts as 5, AWS Config will run auto-remediations 5 times within 50 seconds before throwing an exception.
	RetryAttemptSeconds *float64 `json:"retryAttemptSeconds" yaml:"retryAttemptSeconds"`
	// Version of the target. For example, version of the SSM document.
	//
	// > If you make backward incompatible changes to the SSM document, you must call PutRemediationConfiguration API again to ensure the remediations can run.
	TargetVersion *string `json:"targetVersion" yaml:"targetVersion"`
}

// A CloudFormation `AWS::Config::StoredQuery`.
//
// Provides the details of a stored query.
//
// TODO: EXAMPLE
//
type CfnStoredQuery interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrQueryArn() *string
	AttrQueryId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() awscdk.ConstructNode
	QueryDescription() *string
	SetQueryDescription(val *string)
	QueryExpression() *string
	SetQueryExpression(val *string)
	QueryName() *string
	SetQueryName(val *string)
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnStoredQuery
type jsiiProxy_CfnStoredQuery struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnStoredQuery) AttrQueryArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrQueryArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStoredQuery) AttrQueryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrQueryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStoredQuery) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStoredQuery) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStoredQuery) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStoredQuery) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStoredQuery) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStoredQuery) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStoredQuery) QueryDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStoredQuery) QueryExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStoredQuery) QueryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStoredQuery) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStoredQuery) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStoredQuery) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStoredQuery) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Config::StoredQuery`.
func NewCfnStoredQuery(scope awscdk.Construct, id *string, props *CfnStoredQueryProps) CfnStoredQuery {
	_init_.Initialize()

	j := jsiiProxy_CfnStoredQuery{}

	_jsii_.Create(
		"monocdk.aws_config.CfnStoredQuery",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Config::StoredQuery`.
func NewCfnStoredQuery_Override(c CfnStoredQuery, scope awscdk.Construct, id *string, props *CfnStoredQueryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_config.CfnStoredQuery",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStoredQuery) SetQueryDescription(val *string) {
	_jsii_.Set(
		j,
		"queryDescription",
		val,
	)
}

func (j *jsiiProxy_CfnStoredQuery) SetQueryExpression(val *string) {
	_jsii_.Set(
		j,
		"queryExpression",
		val,
	)
}

func (j *jsiiProxy_CfnStoredQuery) SetQueryName(val *string) {
	_jsii_.Set(
		j,
		"queryName",
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
func CfnStoredQuery_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnStoredQuery",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnStoredQuery_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnStoredQuery",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnStoredQuery_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CfnStoredQuery",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStoredQuery_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.CfnStoredQuery",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnStoredQuery) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnStoredQuery) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnStoredQuery) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
// Experimental.
func (c *jsiiProxy_CfnStoredQuery) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnStoredQuery) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnStoredQuery) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

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
func (c *jsiiProxy_CfnStoredQuery) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnStoredQuery) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnStoredQuery) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnStoredQuery) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnStoredQuery) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnStoredQuery) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnStoredQuery) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnStoredQuery) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnStoredQuery) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnStoredQuery) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnStoredQuery) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnStoredQuery) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnStoredQuery) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnStoredQuery) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnStoredQuery) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnStoredQuery`.
//
// TODO: EXAMPLE
//
type CfnStoredQueryProps struct {
	// The expression of the query.
	//
	// For example, `SELECT resourceId, resourceType, supplementaryConfiguration.BucketVersioningConfiguration.status WHERE resourceType = 'AWS::S3::Bucket' AND supplementaryConfiguration.BucketVersioningConfiguration.status = 'Off'.`
	QueryExpression *string `json:"queryExpression" yaml:"queryExpression"`
	// The name of the query.
	QueryName *string `json:"queryName" yaml:"queryName"`
	// A unique description for the query.
	QueryDescription *string `json:"queryDescription" yaml:"queryDescription"`
	// An array of key-value pairs to apply to this resource.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// Checks whether your CloudFormation stacks' actual configuration differs, or has drifted, from its expected configuration.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/config/latest/developerguide/cloudformation-stack-drift-detection-check.html
//
// Experimental.
type CloudFormationStackDriftDetectionCheck interface {
	ManagedRule
	ConfigRuleArn() *string
	ConfigRuleComplianceType() *string
	ConfigRuleId() *string
	ConfigRuleName() *string
	Env() *awscdk.ResourceEnvironment
	IsCustomWithChanges() *bool
	SetIsCustomWithChanges(val *bool)
	IsManaged() *bool
	SetIsManaged(val *bool)
	Node() awscdk.ConstructNode
	PhysicalName() *string
	RuleScope() RuleScope
	SetRuleScope(val RuleScope)
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnComplianceChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnPrepare()
	OnReEvaluationStatus(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for CloudFormationStackDriftDetectionCheck
type jsiiProxy_CloudFormationStackDriftDetectionCheck struct {
	jsiiProxy_ManagedRule
}

func (j *jsiiProxy_CloudFormationStackDriftDetectionCheck) ConfigRuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackDriftDetectionCheck) ConfigRuleComplianceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleComplianceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackDriftDetectionCheck) ConfigRuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackDriftDetectionCheck) ConfigRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackDriftDetectionCheck) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackDriftDetectionCheck) IsCustomWithChanges() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isCustomWithChanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackDriftDetectionCheck) IsManaged() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isManaged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackDriftDetectionCheck) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackDriftDetectionCheck) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackDriftDetectionCheck) RuleScope() RuleScope {
	var returns RuleScope
	_jsii_.Get(
		j,
		"ruleScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackDriftDetectionCheck) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewCloudFormationStackDriftDetectionCheck(scope constructs.Construct, id *string, props *CloudFormationStackDriftDetectionCheckProps) CloudFormationStackDriftDetectionCheck {
	_init_.Initialize()

	j := jsiiProxy_CloudFormationStackDriftDetectionCheck{}

	_jsii_.Create(
		"monocdk.aws_config.CloudFormationStackDriftDetectionCheck",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudFormationStackDriftDetectionCheck_Override(c CloudFormationStackDriftDetectionCheck, scope constructs.Construct, id *string, props *CloudFormationStackDriftDetectionCheckProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_config.CloudFormationStackDriftDetectionCheck",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CloudFormationStackDriftDetectionCheck) SetIsCustomWithChanges(val *bool) {
	_jsii_.Set(
		j,
		"isCustomWithChanges",
		val,
	)
}

func (j *jsiiProxy_CloudFormationStackDriftDetectionCheck) SetIsManaged(val *bool) {
	_jsii_.Set(
		j,
		"isManaged",
		val,
	)
}

func (j *jsiiProxy_CloudFormationStackDriftDetectionCheck) SetRuleScope(val RuleScope) {
	_jsii_.Set(
		j,
		"ruleScope",
		val,
	)
}

// Imports an existing rule.
// Experimental.
func CloudFormationStackDriftDetectionCheck_FromConfigRuleName(scope constructs.Construct, id *string, configRuleName *string) IRule {
	_init_.Initialize()

	var returns IRule

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CloudFormationStackDriftDetectionCheck",
		"fromConfigRuleName",
		[]interface{}{scope, id, configRuleName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CloudFormationStackDriftDetectionCheck_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CloudFormationStackDriftDetectionCheck",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func CloudFormationStackDriftDetectionCheck_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CloudFormationStackDriftDetectionCheck",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (c *jsiiProxy_CloudFormationStackDriftDetectionCheck) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (c *jsiiProxy_CloudFormationStackDriftDetectionCheck) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (c *jsiiProxy_CloudFormationStackDriftDetectionCheck) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (c *jsiiProxy_CloudFormationStackDriftDetectionCheck) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Defines an EventBridge event rule which triggers for rule compliance events.
// Experimental.
func (c *jsiiProxy_CloudFormationStackDriftDetectionCheck) OnComplianceChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onComplianceChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines an EventBridge event rule which triggers for rule events.
//
// Use
// `rule.addEventPattern(pattern)` to specify a filter.
// Experimental.
func (c *jsiiProxy_CloudFormationStackDriftDetectionCheck) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CloudFormationStackDriftDetectionCheck) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Defines an EventBridge event rule which triggers for rule re-evaluation status events.
// Experimental.
func (c *jsiiProxy_CloudFormationStackDriftDetectionCheck) OnReEvaluationStatus(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onReEvaluationStatus",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CloudFormationStackDriftDetectionCheck) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CloudFormationStackDriftDetectionCheck) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CloudFormationStackDriftDetectionCheck) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CloudFormationStackDriftDetectionCheck) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (c *jsiiProxy_CloudFormationStackDriftDetectionCheck) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CloudFormationStackDriftDetectionCheck) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a CloudFormationStackDriftDetectionCheck.
//
// TODO: EXAMPLE
//
// Experimental.
type CloudFormationStackDriftDetectionCheckProps struct {
	// A name for the AWS Config rule.
	// Experimental.
	ConfigRuleName *string `json:"configRuleName" yaml:"configRuleName"`
	// A description about this AWS Config rule.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// Input parameter values that are passed to the AWS Config rule.
	// Experimental.
	InputParameters *map[string]interface{} `json:"inputParameters" yaml:"inputParameters"`
	// The maximum frequency at which the AWS Config rule runs evaluations.
	// Experimental.
	MaximumExecutionFrequency MaximumExecutionFrequency `json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// Defines which resources trigger an evaluation for an AWS Config rule.
	// Experimental.
	RuleScope RuleScope `json:"ruleScope" yaml:"ruleScope"`
	// Whether to check only the stack where this rule is deployed.
	// Experimental.
	OwnStackOnly *bool `json:"ownStackOnly" yaml:"ownStackOnly"`
	// The IAM role to use for this rule.
	//
	// It must have permissions to detect drift
	// for AWS CloudFormation stacks. Ensure to attach `config.amazonaws.com` trusted
	// permissions and `ReadOnlyAccess` policy permissions. For specific policy permissions,
	// refer to https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift.html.
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
}

// Checks whether your CloudFormation stacks are sending event notifications to a SNS topic.
//
// Optionally checks whether specified SNS topics are used.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/config/latest/developerguide/cloudformation-stack-notification-check.html
//
// Experimental.
type CloudFormationStackNotificationCheck interface {
	ManagedRule
	ConfigRuleArn() *string
	ConfigRuleComplianceType() *string
	ConfigRuleId() *string
	ConfigRuleName() *string
	Env() *awscdk.ResourceEnvironment
	IsCustomWithChanges() *bool
	SetIsCustomWithChanges(val *bool)
	IsManaged() *bool
	SetIsManaged(val *bool)
	Node() awscdk.ConstructNode
	PhysicalName() *string
	RuleScope() RuleScope
	SetRuleScope(val RuleScope)
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnComplianceChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnPrepare()
	OnReEvaluationStatus(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for CloudFormationStackNotificationCheck
type jsiiProxy_CloudFormationStackNotificationCheck struct {
	jsiiProxy_ManagedRule
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck) ConfigRuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck) ConfigRuleComplianceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleComplianceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck) ConfigRuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck) ConfigRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck) IsCustomWithChanges() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isCustomWithChanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck) IsManaged() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isManaged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck) RuleScope() RuleScope {
	var returns RuleScope
	_jsii_.Get(
		j,
		"ruleScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewCloudFormationStackNotificationCheck(scope constructs.Construct, id *string, props *CloudFormationStackNotificationCheckProps) CloudFormationStackNotificationCheck {
	_init_.Initialize()

	j := jsiiProxy_CloudFormationStackNotificationCheck{}

	_jsii_.Create(
		"monocdk.aws_config.CloudFormationStackNotificationCheck",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudFormationStackNotificationCheck_Override(c CloudFormationStackNotificationCheck, scope constructs.Construct, id *string, props *CloudFormationStackNotificationCheckProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_config.CloudFormationStackNotificationCheck",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck) SetIsCustomWithChanges(val *bool) {
	_jsii_.Set(
		j,
		"isCustomWithChanges",
		val,
	)
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck) SetIsManaged(val *bool) {
	_jsii_.Set(
		j,
		"isManaged",
		val,
	)
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck) SetRuleScope(val RuleScope) {
	_jsii_.Set(
		j,
		"ruleScope",
		val,
	)
}

// Imports an existing rule.
// Experimental.
func CloudFormationStackNotificationCheck_FromConfigRuleName(scope constructs.Construct, id *string, configRuleName *string) IRule {
	_init_.Initialize()

	var returns IRule

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CloudFormationStackNotificationCheck",
		"fromConfigRuleName",
		[]interface{}{scope, id, configRuleName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CloudFormationStackNotificationCheck_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CloudFormationStackNotificationCheck",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func CloudFormationStackNotificationCheck_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CloudFormationStackNotificationCheck",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (c *jsiiProxy_CloudFormationStackNotificationCheck) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (c *jsiiProxy_CloudFormationStackNotificationCheck) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (c *jsiiProxy_CloudFormationStackNotificationCheck) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (c *jsiiProxy_CloudFormationStackNotificationCheck) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Defines an EventBridge event rule which triggers for rule compliance events.
// Experimental.
func (c *jsiiProxy_CloudFormationStackNotificationCheck) OnComplianceChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onComplianceChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines an EventBridge event rule which triggers for rule events.
//
// Use
// `rule.addEventPattern(pattern)` to specify a filter.
// Experimental.
func (c *jsiiProxy_CloudFormationStackNotificationCheck) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CloudFormationStackNotificationCheck) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Defines an EventBridge event rule which triggers for rule re-evaluation status events.
// Experimental.
func (c *jsiiProxy_CloudFormationStackNotificationCheck) OnReEvaluationStatus(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onReEvaluationStatus",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CloudFormationStackNotificationCheck) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CloudFormationStackNotificationCheck) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CloudFormationStackNotificationCheck) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CloudFormationStackNotificationCheck) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (c *jsiiProxy_CloudFormationStackNotificationCheck) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CloudFormationStackNotificationCheck) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a CloudFormationStackNotificationCheck.
//
// TODO: EXAMPLE
//
// Experimental.
type CloudFormationStackNotificationCheckProps struct {
	// A name for the AWS Config rule.
	// Experimental.
	ConfigRuleName *string `json:"configRuleName" yaml:"configRuleName"`
	// A description about this AWS Config rule.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// Input parameter values that are passed to the AWS Config rule.
	// Experimental.
	InputParameters *map[string]interface{} `json:"inputParameters" yaml:"inputParameters"`
	// The maximum frequency at which the AWS Config rule runs evaluations.
	// Experimental.
	MaximumExecutionFrequency MaximumExecutionFrequency `json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// Defines which resources trigger an evaluation for an AWS Config rule.
	// Experimental.
	RuleScope RuleScope `json:"ruleScope" yaml:"ruleScope"`
	// A list of allowed topics.
	//
	// At most 5 topics.
	// Experimental.
	Topics *[]awssns.ITopic `json:"topics" yaml:"topics"`
}

// A new custom rule.
//
// TODO: EXAMPLE
//
// Experimental.
type CustomRule interface {
	awscdk.Resource
	IRule
	ConfigRuleArn() *string
	ConfigRuleComplianceType() *string
	ConfigRuleId() *string
	ConfigRuleName() *string
	Env() *awscdk.ResourceEnvironment
	IsCustomWithChanges() *bool
	SetIsCustomWithChanges(val *bool)
	IsManaged() *bool
	SetIsManaged(val *bool)
	Node() awscdk.ConstructNode
	PhysicalName() *string
	RuleScope() RuleScope
	SetRuleScope(val RuleScope)
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnComplianceChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnPrepare()
	OnReEvaluationStatus(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for CustomRule
type jsiiProxy_CustomRule struct {
	internal.Type__awscdkResource
	jsiiProxy_IRule
}

func (j *jsiiProxy_CustomRule) ConfigRuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) ConfigRuleComplianceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleComplianceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) ConfigRuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) ConfigRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) IsCustomWithChanges() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isCustomWithChanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) IsManaged() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isManaged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) RuleScope() RuleScope {
	var returns RuleScope
	_jsii_.Get(
		j,
		"ruleScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewCustomRule(scope constructs.Construct, id *string, props *CustomRuleProps) CustomRule {
	_init_.Initialize()

	j := jsiiProxy_CustomRule{}

	_jsii_.Create(
		"monocdk.aws_config.CustomRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCustomRule_Override(c CustomRule, scope constructs.Construct, id *string, props *CustomRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_config.CustomRule",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CustomRule) SetIsCustomWithChanges(val *bool) {
	_jsii_.Set(
		j,
		"isCustomWithChanges",
		val,
	)
}

func (j *jsiiProxy_CustomRule) SetIsManaged(val *bool) {
	_jsii_.Set(
		j,
		"isManaged",
		val,
	)
}

func (j *jsiiProxy_CustomRule) SetRuleScope(val RuleScope) {
	_jsii_.Set(
		j,
		"ruleScope",
		val,
	)
}

// Imports an existing rule.
// Experimental.
func CustomRule_FromConfigRuleName(scope constructs.Construct, id *string, configRuleName *string) IRule {
	_init_.Initialize()

	var returns IRule

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CustomRule",
		"fromConfigRuleName",
		[]interface{}{scope, id, configRuleName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CustomRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CustomRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func CustomRule_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CustomRule",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (c *jsiiProxy_CustomRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (c *jsiiProxy_CustomRule) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (c *jsiiProxy_CustomRule) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (c *jsiiProxy_CustomRule) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Defines an EventBridge event rule which triggers for rule compliance events.
// Experimental.
func (c *jsiiProxy_CustomRule) OnComplianceChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onComplianceChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines an EventBridge event rule which triggers for rule events.
//
// Use
// `rule.addEventPattern(pattern)` to specify a filter.
// Experimental.
func (c *jsiiProxy_CustomRule) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CustomRule) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Defines an EventBridge event rule which triggers for rule re-evaluation status events.
// Experimental.
func (c *jsiiProxy_CustomRule) OnReEvaluationStatus(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onReEvaluationStatus",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CustomRule) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CustomRule) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CustomRule) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CustomRule) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (c *jsiiProxy_CustomRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CustomRule) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a CustomRule.
//
// TODO: EXAMPLE
//
// Experimental.
type CustomRuleProps struct {
	// A name for the AWS Config rule.
	// Experimental.
	ConfigRuleName *string `json:"configRuleName" yaml:"configRuleName"`
	// A description about this AWS Config rule.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// Input parameter values that are passed to the AWS Config rule.
	// Experimental.
	InputParameters *map[string]interface{} `json:"inputParameters" yaml:"inputParameters"`
	// The maximum frequency at which the AWS Config rule runs evaluations.
	// Experimental.
	MaximumExecutionFrequency MaximumExecutionFrequency `json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// Defines which resources trigger an evaluation for an AWS Config rule.
	// Experimental.
	RuleScope RuleScope `json:"ruleScope" yaml:"ruleScope"`
	// The Lambda function to run.
	// Experimental.
	LambdaFunction awslambda.IFunction `json:"lambdaFunction" yaml:"lambdaFunction"`
	// Whether to run the rule on configuration changes.
	// Experimental.
	ConfigurationChanges *bool `json:"configurationChanges" yaml:"configurationChanges"`
	// Whether to run the rule on a fixed frequency.
	// Experimental.
	Periodic *bool `json:"periodic" yaml:"periodic"`
}

// Interface representing an AWS Config rule.
// Experimental.
type IRule interface {
	awscdk.IResource
	// Defines a EventBridge event rule which triggers for rule compliance events.
	// Experimental.
	OnComplianceChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an EventBridge event rule which triggers for rule events.
	//
	// Use
	// `rule.addEventPattern(pattern)` to specify a filter.
	// Experimental.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a EventBridge event rule which triggers for rule re-evaluation status events.
	// Experimental.
	OnReEvaluationStatus(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// The name of the rule.
	// Experimental.
	ConfigRuleName() *string
}

// The jsii proxy for IRule
type jsiiProxy_IRule struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IRule) OnComplianceChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onComplianceChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRule) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRule) OnReEvaluationStatus(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onReEvaluationStatus",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IRule) ConfigRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleName",
		&returns,
	)
	return returns
}

// A new managed rule.
//
// TODO: EXAMPLE
//
// Experimental.
type ManagedRule interface {
	awscdk.Resource
	IRule
	ConfigRuleArn() *string
	ConfigRuleComplianceType() *string
	ConfigRuleId() *string
	ConfigRuleName() *string
	Env() *awscdk.ResourceEnvironment
	IsCustomWithChanges() *bool
	SetIsCustomWithChanges(val *bool)
	IsManaged() *bool
	SetIsManaged(val *bool)
	Node() awscdk.ConstructNode
	PhysicalName() *string
	RuleScope() RuleScope
	SetRuleScope(val RuleScope)
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnComplianceChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnPrepare()
	OnReEvaluationStatus(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for ManagedRule
type jsiiProxy_ManagedRule struct {
	internal.Type__awscdkResource
	jsiiProxy_IRule
}

func (j *jsiiProxy_ManagedRule) ConfigRuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRule) ConfigRuleComplianceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleComplianceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRule) ConfigRuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRule) ConfigRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRule) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRule) IsCustomWithChanges() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isCustomWithChanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRule) IsManaged() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isManaged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRule) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRule) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRule) RuleScope() RuleScope {
	var returns RuleScope
	_jsii_.Get(
		j,
		"ruleScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewManagedRule(scope constructs.Construct, id *string, props *ManagedRuleProps) ManagedRule {
	_init_.Initialize()

	j := jsiiProxy_ManagedRule{}

	_jsii_.Create(
		"monocdk.aws_config.ManagedRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewManagedRule_Override(m ManagedRule, scope constructs.Construct, id *string, props *ManagedRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_config.ManagedRule",
		[]interface{}{scope, id, props},
		m,
	)
}

func (j *jsiiProxy_ManagedRule) SetIsCustomWithChanges(val *bool) {
	_jsii_.Set(
		j,
		"isCustomWithChanges",
		val,
	)
}

func (j *jsiiProxy_ManagedRule) SetIsManaged(val *bool) {
	_jsii_.Set(
		j,
		"isManaged",
		val,
	)
}

func (j *jsiiProxy_ManagedRule) SetRuleScope(val RuleScope) {
	_jsii_.Set(
		j,
		"ruleScope",
		val,
	)
}

// Imports an existing rule.
// Experimental.
func ManagedRule_FromConfigRuleName(scope constructs.Construct, id *string, configRuleName *string) IRule {
	_init_.Initialize()

	var returns IRule

	_jsii_.StaticInvoke(
		"monocdk.aws_config.ManagedRule",
		"fromConfigRuleName",
		[]interface{}{scope, id, configRuleName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func ManagedRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.ManagedRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func ManagedRule_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.ManagedRule",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (m *jsiiProxy_ManagedRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		m,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (m *jsiiProxy_ManagedRule) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (m *jsiiProxy_ManagedRule) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (m *jsiiProxy_ManagedRule) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Defines an EventBridge event rule which triggers for rule compliance events.
// Experimental.
func (m *jsiiProxy_ManagedRule) OnComplianceChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		m,
		"onComplianceChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines an EventBridge event rule which triggers for rule events.
//
// Use
// `rule.addEventPattern(pattern)` to specify a filter.
// Experimental.
func (m *jsiiProxy_ManagedRule) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		m,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (m *jsiiProxy_ManagedRule) OnPrepare() {
	_jsii_.InvokeVoid(
		m,
		"onPrepare",
		nil, // no parameters
	)
}

// Defines an EventBridge event rule which triggers for rule re-evaluation status events.
// Experimental.
func (m *jsiiProxy_ManagedRule) OnReEvaluationStatus(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		m,
		"onReEvaluationStatus",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (m *jsiiProxy_ManagedRule) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		m,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (m *jsiiProxy_ManagedRule) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (m *jsiiProxy_ManagedRule) Prepare() {
	_jsii_.InvokeVoid(
		m,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (m *jsiiProxy_ManagedRule) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		m,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (m *jsiiProxy_ManagedRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (m *jsiiProxy_ManagedRule) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Managed rules that are supported by AWS Config.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html
//
// Experimental.
type ManagedRuleIdentifiers interface {
}

// The jsii proxy struct for ManagedRuleIdentifiers
type jsiiProxy_ManagedRuleIdentifiers struct {
	_ byte // padding
}

func ManagedRuleIdentifiers_ACCESS_KEYS_ROTATED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"ACCESS_KEYS_ROTATED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ACCOUNT_PART_OF_ORGANIZATIONS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"ACCOUNT_PART_OF_ORGANIZATIONS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ACM_CERTIFICATE_EXPIRATION_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"ACM_CERTIFICATE_EXPIRATION_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ALB_HTTP_DROP_INVALID_HEADER_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"ALB_HTTP_DROP_INVALID_HEADER_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ALB_HTTP_TO_HTTPS_REDIRECTION_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"ALB_HTTP_TO_HTTPS_REDIRECTION_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ALB_WAF_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"ALB_WAF_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_API_GW_CACHE_ENABLED_AND_ENCRYPTED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"API_GW_CACHE_ENABLED_AND_ENCRYPTED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_API_GW_ENDPOINT_TYPE_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"API_GW_ENDPOINT_TYPE_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_API_GW_EXECUTION_LOGGING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"API_GW_EXECUTION_LOGGING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_APPROVED_AMIS_BY_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"APPROVED_AMIS_BY_ID",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_APPROVED_AMIS_BY_TAG() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"APPROVED_AMIS_BY_TAG",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_AUTOSCALING_GROUP_ELB_HEALTHCHECK_REQUIRED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"AUTOSCALING_GROUP_ELB_HEALTHCHECK_REQUIRED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUD_TRAIL_CLOUD_WATCH_LOGS_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"CLOUD_TRAIL_CLOUD_WATCH_LOGS_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUD_TRAIL_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"CLOUD_TRAIL_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUD_TRAIL_ENCRYPTION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"CLOUD_TRAIL_ENCRYPTION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUD_TRAIL_LOG_FILE_VALIDATION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"CLOUD_TRAIL_LOG_FILE_VALIDATION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDFORMATION_STACK_DRIFT_DETECTION_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"CLOUDFORMATION_STACK_DRIFT_DETECTION_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDFORMATION_STACK_NOTIFICATION_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"CLOUDFORMATION_STACK_NOTIFICATION_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDFRONT_DEFAULT_ROOT_OBJECT_CONFIGURED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"CLOUDFRONT_DEFAULT_ROOT_OBJECT_CONFIGURED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDFRONT_ORIGIN_ACCESS_IDENTITY_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"CLOUDFRONT_ORIGIN_ACCESS_IDENTITY_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDFRONT_ORIGIN_FAILOVER_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"CLOUDFRONT_ORIGIN_FAILOVER_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDFRONT_SNI_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"CLOUDFRONT_SNI_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDFRONT_VIEWER_POLICY_HTTPS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"CLOUDFRONT_VIEWER_POLICY_HTTPS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDTRAIL_MULTI_REGION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"CLOUDTRAIL_MULTI_REGION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDTRAIL_S3_DATAEVENTS_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"CLOUDTRAIL_S3_DATAEVENTS_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDTRAIL_SECURITY_TRAIL_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"CLOUDTRAIL_SECURITY_TRAIL_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDWATCH_ALARM_ACTION_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"CLOUDWATCH_ALARM_ACTION_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDWATCH_ALARM_RESOURCE_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"CLOUDWATCH_ALARM_RESOURCE_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDWATCH_ALARM_SETTINGS_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"CLOUDWATCH_ALARM_SETTINGS_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CLOUDWATCH_LOG_GROUP_ENCRYPTED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"CLOUDWATCH_LOG_GROUP_ENCRYPTED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CMK_BACKING_KEY_ROTATION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"CMK_BACKING_KEY_ROTATION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CODEBUILD_PROJECT_ENVVAR_AWSCRED_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"CODEBUILD_PROJECT_ENVVAR_AWSCRED_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CODEBUILD_PROJECT_SOURCE_REPO_URL_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"CODEBUILD_PROJECT_SOURCE_REPO_URL_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CODEPIPELINE_DEPLOYMENT_COUNT_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"CODEPIPELINE_DEPLOYMENT_COUNT_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CODEPIPELINE_REGION_FANOUT_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"CODEPIPELINE_REGION_FANOUT_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_CW_LOGGROUP_RETENTION_PERIOD_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"CW_LOGGROUP_RETENTION_PERIOD_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_DAX_ENCRYPTION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"DAX_ENCRYPTION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_DMS_REPLICATION_NOT_PUBLIC() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"DMS_REPLICATION_NOT_PUBLIC",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_DYNAMODB_AUTOSCALING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"DYNAMODB_AUTOSCALING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_DYNAMODB_IN_BACKUP_PLAN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"DYNAMODB_IN_BACKUP_PLAN",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_DYNAMODB_PITR_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"DYNAMODB_PITR_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_DYNAMODB_TABLE_ENCRYPTED_KMS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"DYNAMODB_TABLE_ENCRYPTED_KMS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_DYNAMODB_TABLE_ENCRYPTION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"DYNAMODB_TABLE_ENCRYPTION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_DYNAMODB_THROUGHPUT_LIMIT_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"DYNAMODB_THROUGHPUT_LIMIT_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EBS_ENCRYPTED_VOLUMES() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EBS_ENCRYPTED_VOLUMES",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EBS_IN_BACKUP_PLAN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EBS_IN_BACKUP_PLAN",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EBS_OPTIMIZED_INSTANCE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EBS_OPTIMIZED_INSTANCE",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EBS_SNAPSHOT_PUBLIC_RESTORABLE_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EBS_SNAPSHOT_PUBLIC_RESTORABLE_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_DESIRED_INSTANCE_TENANCY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EC2_DESIRED_INSTANCE_TENANCY",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_DESIRED_INSTANCE_TYPE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EC2_DESIRED_INSTANCE_TYPE",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_EBS_ENCRYPTION_BY_DEFAULT() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EC2_EBS_ENCRYPTION_BY_DEFAULT",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_IMDSV2_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EC2_IMDSV2_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_INSTANCE_DETAILED_MONITORING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EC2_INSTANCE_DETAILED_MONITORING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_INSTANCE_MANAGED_BY_SSM() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EC2_INSTANCE_MANAGED_BY_SSM",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_INSTANCE_NO_PUBLIC_IP() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EC2_INSTANCE_NO_PUBLIC_IP",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_INSTANCE_PROFILE_ATTACHED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EC2_INSTANCE_PROFILE_ATTACHED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_INSTANCES_IN_VPC() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EC2_INSTANCES_IN_VPC",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_MANAGED_INSTANCE_APPLICATIONS_BLOCKED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EC2_MANAGED_INSTANCE_APPLICATIONS_BLOCKED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_MANAGED_INSTANCE_APPLICATIONS_REQUIRED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EC2_MANAGED_INSTANCE_APPLICATIONS_REQUIRED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_MANAGED_INSTANCE_ASSOCIATION_COMPLIANCE_STATUS_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EC2_MANAGED_INSTANCE_ASSOCIATION_COMPLIANCE_STATUS_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_MANAGED_INSTANCE_INVENTORY_BLOCKED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EC2_MANAGED_INSTANCE_INVENTORY_BLOCKED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_MANAGED_INSTANCE_PATCH_COMPLIANCE_STATUS_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EC2_MANAGED_INSTANCE_PATCH_COMPLIANCE_STATUS_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_MANAGED_INSTANCE_PLATFORM_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EC2_MANAGED_INSTANCE_PLATFORM_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_SECURITY_GROUP_ATTACHED_TO_ENI() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EC2_SECURITY_GROUP_ATTACHED_TO_ENI",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_SECURITY_GROUPS_INCOMING_SSH_DISABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EC2_SECURITY_GROUPS_INCOMING_SSH_DISABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_SECURITY_GROUPS_RESTRICTED_INCOMING_TRAFFIC() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EC2_SECURITY_GROUPS_RESTRICTED_INCOMING_TRAFFIC",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_STOPPED_INSTANCE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EC2_STOPPED_INSTANCE",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EC2_VOLUME_INUSE_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EC2_VOLUME_INUSE_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EFS_ENCRYPTED_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EFS_ENCRYPTED_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EFS_IN_BACKUP_PLAN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EFS_IN_BACKUP_PLAN",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EIP_ATTACHED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EIP_ATTACHED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EKS_ENDPOINT_NO_PUBLIC_ACCESS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EKS_ENDPOINT_NO_PUBLIC_ACCESS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EKS_SECRETS_ENCRYPTED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EKS_SECRETS_ENCRYPTED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ELASTICACHE_REDIS_CLUSTER_AUTOMATIC_BACKUP_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"ELASTICACHE_REDIS_CLUSTER_AUTOMATIC_BACKUP_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ELASTICSEARCH_ENCRYPTED_AT_REST() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"ELASTICSEARCH_ENCRYPTED_AT_REST",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ELASTICSEARCH_IN_VPC_ONLY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"ELASTICSEARCH_IN_VPC_ONLY",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ELASTICSEARCH_NODE_TO_NODE_ENCRYPTION_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"ELASTICSEARCH_NODE_TO_NODE_ENCRYPTION_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ELB_ACM_CERTIFICATE_REQUIRED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"ELB_ACM_CERTIFICATE_REQUIRED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ELB_CROSS_ZONE_LOAD_BALANCING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"ELB_CROSS_ZONE_LOAD_BALANCING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ELB_CUSTOM_SECURITY_POLICY_SSL_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"ELB_CUSTOM_SECURITY_POLICY_SSL_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ELB_DELETION_PROTECTION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"ELB_DELETION_PROTECTION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ELB_LOGGING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"ELB_LOGGING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ELB_PREDEFINED_SECURITY_POLICY_SSL_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"ELB_PREDEFINED_SECURITY_POLICY_SSL_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ELB_TLS_HTTPS_LISTENERS_ONLY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"ELB_TLS_HTTPS_LISTENERS_ONLY",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EMR_KERBEROS_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EMR_KERBEROS_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_EMR_MASTER_NO_PUBLIC_IP() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"EMR_MASTER_NO_PUBLIC_IP",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_FMS_SECURITY_GROUP_AUDIT_POLICY_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"FMS_SECURITY_GROUP_AUDIT_POLICY_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_FMS_SECURITY_GROUP_CONTENT_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"FMS_SECURITY_GROUP_CONTENT_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_FMS_SECURITY_GROUP_RESOURCE_ASSOCIATION_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"FMS_SECURITY_GROUP_RESOURCE_ASSOCIATION_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_FMS_SHIELD_RESOURCE_POLICY_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"FMS_SHIELD_RESOURCE_POLICY_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_FMS_WEBACL_RESOURCE_POLICY_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"FMS_WEBACL_RESOURCE_POLICY_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_FMS_WEBACL_RULEGROUP_ASSOCIATION_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"FMS_WEBACL_RULEGROUP_ASSOCIATION_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_GUARDDUTY_ENABLED_CENTRALIZED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"GUARDDUTY_ENABLED_CENTRALIZED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_GUARDDUTY_NON_ARCHIVED_FINDINGS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"GUARDDUTY_NON_ARCHIVED_FINDINGS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_CUSTOMER_POLICY_BLOCKED_KMS_ACTIONS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"IAM_CUSTOMER_POLICY_BLOCKED_KMS_ACTIONS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_GROUP_HAS_USERS_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"IAM_GROUP_HAS_USERS_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_INLINE_POLICY_BLOCKED_KMS_ACTIONS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"IAM_INLINE_POLICY_BLOCKED_KMS_ACTIONS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_NO_INLINE_POLICY_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"IAM_NO_INLINE_POLICY_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_PASSWORD_POLICY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"IAM_PASSWORD_POLICY",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_POLICY_BLOCKED_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"IAM_POLICY_BLOCKED_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_POLICY_IN_USE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"IAM_POLICY_IN_USE",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_POLICY_NO_STATEMENTS_WITH_ADMIN_ACCESS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"IAM_POLICY_NO_STATEMENTS_WITH_ADMIN_ACCESS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_ROLE_MANAGED_POLICY_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"IAM_ROLE_MANAGED_POLICY_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_ROOT_ACCESS_KEY_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"IAM_ROOT_ACCESS_KEY_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_USER_GROUP_MEMBERSHIP_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"IAM_USER_GROUP_MEMBERSHIP_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_USER_MFA_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"IAM_USER_MFA_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_USER_NO_POLICIES_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"IAM_USER_NO_POLICIES_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_IAM_USER_UNUSED_CREDENTIALS_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"IAM_USER_UNUSED_CREDENTIALS_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_INTERNET_GATEWAY_AUTHORIZED_VPC_ONLY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"INTERNET_GATEWAY_AUTHORIZED_VPC_ONLY",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_KMS_CMK_NOT_SCHEDULED_FOR_DELETION() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"KMS_CMK_NOT_SCHEDULED_FOR_DELETION",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_LAMBDA_CONCURRENCY_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"LAMBDA_CONCURRENCY_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_LAMBDA_DLQ_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"LAMBDA_DLQ_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_LAMBDA_FUNCTION_PUBLIC_ACCESS_PROHIBITED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"LAMBDA_FUNCTION_PUBLIC_ACCESS_PROHIBITED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_LAMBDA_FUNCTION_SETTINGS_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"LAMBDA_FUNCTION_SETTINGS_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_LAMBDA_INSIDE_VPC() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"LAMBDA_INSIDE_VPC",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_MFA_ENABLED_FOR_IAM_CONSOLE_ACCESS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"MFA_ENABLED_FOR_IAM_CONSOLE_ACCESS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_CLUSTER_DELETION_PROTECTION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"RDS_CLUSTER_DELETION_PROTECTION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_DB_INSTANCE_BACKUP_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"RDS_DB_INSTANCE_BACKUP_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_ENHANCED_MONITORING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"RDS_ENHANCED_MONITORING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_IN_BACKUP_PLAN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"RDS_IN_BACKUP_PLAN",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_INSTANCE_DELETION_PROTECTION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"RDS_INSTANCE_DELETION_PROTECTION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_INSTANCE_IAM_AUTHENTICATION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"RDS_INSTANCE_IAM_AUTHENTICATION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_INSTANCE_PUBLIC_ACCESS_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"RDS_INSTANCE_PUBLIC_ACCESS_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_LOGGING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"RDS_LOGGING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_MULTI_AZ_SUPPORT() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"RDS_MULTI_AZ_SUPPORT",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_SNAPSHOT_ENCRYPTED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"RDS_SNAPSHOT_ENCRYPTED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_SNAPSHOTS_PUBLIC_PROHIBITED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"RDS_SNAPSHOTS_PUBLIC_PROHIBITED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_RDS_STORAGE_ENCRYPTED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"RDS_STORAGE_ENCRYPTED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_REDSHIFT_BACKUP_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"REDSHIFT_BACKUP_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_REDSHIFT_CLUSTER_CONFIGURATION_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"REDSHIFT_CLUSTER_CONFIGURATION_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_REDSHIFT_CLUSTER_MAINTENANCE_SETTINGS_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"REDSHIFT_CLUSTER_MAINTENANCE_SETTINGS_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_REDSHIFT_CLUSTER_PUBLIC_ACCESS_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"REDSHIFT_CLUSTER_PUBLIC_ACCESS_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_REDSHIFT_REQUIRE_TLS_SSL() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"REDSHIFT_REQUIRE_TLS_SSL",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_REQUIRED_TAGS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"REQUIRED_TAGS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ROOT_ACCOUNT_HARDWARE_MFA_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"ROOT_ACCOUNT_HARDWARE_MFA_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_ROOT_ACCOUNT_MFA_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"ROOT_ACCOUNT_MFA_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_ACCOUNT_LEVEL_PUBLIC_ACCESS_BLOCKS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"S3_ACCOUNT_LEVEL_PUBLIC_ACCESS_BLOCKS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_BUCKET_BLOCKED_ACTIONS_PROHIBITED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"S3_BUCKET_BLOCKED_ACTIONS_PROHIBITED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_BUCKET_DEFAULT_LOCK_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"S3_BUCKET_DEFAULT_LOCK_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_BUCKET_LOGGING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"S3_BUCKET_LOGGING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_BUCKET_POLICY_GRANTEE_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"S3_BUCKET_POLICY_GRANTEE_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_BUCKET_POLICY_NOT_MORE_PERMISSIVE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"S3_BUCKET_POLICY_NOT_MORE_PERMISSIVE",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_BUCKET_PUBLIC_READ_PROHIBITED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"S3_BUCKET_PUBLIC_READ_PROHIBITED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_BUCKET_PUBLIC_WRITE_PROHIBITED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"S3_BUCKET_PUBLIC_WRITE_PROHIBITED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_BUCKET_REPLICATION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"S3_BUCKET_REPLICATION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_BUCKET_SERVER_SIDE_ENCRYPTION_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"S3_BUCKET_SERVER_SIDE_ENCRYPTION_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_BUCKET_SSL_REQUESTS_ONLY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"S3_BUCKET_SSL_REQUESTS_ONLY",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_BUCKET_VERSIONING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"S3_BUCKET_VERSIONING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_S3_DEFAULT_ENCRYPTION_KMS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"S3_DEFAULT_ENCRYPTION_KMS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SAGEMAKER_ENDPOINT_CONFIGURATION_KMS_KEY_CONFIGURED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"SAGEMAKER_ENDPOINT_CONFIGURATION_KMS_KEY_CONFIGURED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SAGEMAKER_NOTEBOOK_INSTANCE_KMS_KEY_CONFIGURED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"SAGEMAKER_NOTEBOOK_INSTANCE_KMS_KEY_CONFIGURED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SAGEMAKER_NOTEBOOK_NO_DIRECT_INTERNET_ACCESS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"SAGEMAKER_NOTEBOOK_NO_DIRECT_INTERNET_ACCESS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SECRETSMANAGER_ROTATION_ENABLED_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"SECRETSMANAGER_ROTATION_ENABLED_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SECRETSMANAGER_SCHEDULED_ROTATION_SUCCESS_CHECK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"SECRETSMANAGER_SCHEDULED_ROTATION_SUCCESS_CHECK",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SECURITYHUB_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"SECURITYHUB_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SERVICE_VPC_ENDPOINT_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"SERVICE_VPC_ENDPOINT_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SHIELD_ADVANCED_ENABLED_AUTO_RENEW() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"SHIELD_ADVANCED_ENABLED_AUTO_RENEW",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SHIELD_DRT_ACCESS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"SHIELD_DRT_ACCESS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_SNS_ENCRYPTED_KMS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"SNS_ENCRYPTED_KMS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_VPC_DEFAULT_SECURITY_GROUP_CLOSED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"VPC_DEFAULT_SECURITY_GROUP_CLOSED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_VPC_FLOW_LOGS_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"VPC_FLOW_LOGS_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_VPC_SG_OPEN_ONLY_TO_AUTHORIZED_PORTS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"VPC_SG_OPEN_ONLY_TO_AUTHORIZED_PORTS",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_VPC_VPN_2_TUNNELS_UP() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"VPC_VPN_2_TUNNELS_UP",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_WAF_CLASSIC_LOGGING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"WAF_CLASSIC_LOGGING_ENABLED",
		&returns,
	)
	return returns
}

func ManagedRuleIdentifiers_WAFV2_LOGGING_ENABLED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_config.ManagedRuleIdentifiers",
		"WAFV2_LOGGING_ENABLED",
		&returns,
	)
	return returns
}

// Construction properties for a ManagedRule.
//
// TODO: EXAMPLE
//
// Experimental.
type ManagedRuleProps struct {
	// A name for the AWS Config rule.
	// Experimental.
	ConfigRuleName *string `json:"configRuleName" yaml:"configRuleName"`
	// A description about this AWS Config rule.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// Input parameter values that are passed to the AWS Config rule.
	// Experimental.
	InputParameters *map[string]interface{} `json:"inputParameters" yaml:"inputParameters"`
	// The maximum frequency at which the AWS Config rule runs evaluations.
	// Experimental.
	MaximumExecutionFrequency MaximumExecutionFrequency `json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// Defines which resources trigger an evaluation for an AWS Config rule.
	// Experimental.
	RuleScope RuleScope `json:"ruleScope" yaml:"ruleScope"`
	// The identifier of the AWS managed rule.
	// See: https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html
	//
	// Experimental.
	Identifier *string `json:"identifier" yaml:"identifier"`
}

// The maximum frequency at which the AWS Config rule runs evaluations.
//
// TODO: EXAMPLE
//
// Experimental.
type MaximumExecutionFrequency string

const (
	MaximumExecutionFrequency_ONE_HOUR MaximumExecutionFrequency = "ONE_HOUR"
	MaximumExecutionFrequency_THREE_HOURS MaximumExecutionFrequency = "THREE_HOURS"
	MaximumExecutionFrequency_SIX_HOURS MaximumExecutionFrequency = "SIX_HOURS"
	MaximumExecutionFrequency_TWELVE_HOURS MaximumExecutionFrequency = "TWELVE_HOURS"
	MaximumExecutionFrequency_TWENTY_FOUR_HOURS MaximumExecutionFrequency = "TWENTY_FOUR_HOURS"
)

// Resources types that are supported by AWS Config.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/config/latest/developerguide/resource-config-reference.html
//
// Experimental.
type ResourceType interface {
	ComplianceResourceType() *string
}

// The jsii proxy struct for ResourceType
type jsiiProxy_ResourceType struct {
	_ byte // padding
}

func (j *jsiiProxy_ResourceType) ComplianceResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceResourceType",
		&returns,
	)
	return returns
}


// A custom resource type to support future cases.
// Experimental.
func ResourceType_Of(type_ *string) ResourceType {
	_init_.Initialize()

	var returns ResourceType

	_jsii_.StaticInvoke(
		"monocdk.aws_config.ResourceType",
		"of",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

func ResourceType_ACM_CERTIFICATE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"ACM_CERTIFICATE",
		&returns,
	)
	return returns
}

func ResourceType_APIGATEWAY_REST_API() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"APIGATEWAY_REST_API",
		&returns,
	)
	return returns
}

func ResourceType_APIGATEWAY_STAGE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"APIGATEWAY_STAGE",
		&returns,
	)
	return returns
}

func ResourceType_APIGATEWAYV2_API() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"APIGATEWAYV2_API",
		&returns,
	)
	return returns
}

func ResourceType_APIGATEWAYV2_STAGE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"APIGATEWAYV2_STAGE",
		&returns,
	)
	return returns
}

func ResourceType_AUTO_SCALING_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"AUTO_SCALING_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_AUTO_SCALING_LAUNCH_CONFIGURATION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"AUTO_SCALING_LAUNCH_CONFIGURATION",
		&returns,
	)
	return returns
}

func ResourceType_AUTO_SCALING_POLICY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"AUTO_SCALING_POLICY",
		&returns,
	)
	return returns
}

func ResourceType_AUTO_SCALING_SCHEDULED_ACTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"AUTO_SCALING_SCHEDULED_ACTION",
		&returns,
	)
	return returns
}

func ResourceType_CLOUDFORMATION_STACK() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"CLOUDFORMATION_STACK",
		&returns,
	)
	return returns
}

func ResourceType_CLOUDFRONT_DISTRIBUTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"CLOUDFRONT_DISTRIBUTION",
		&returns,
	)
	return returns
}

func ResourceType_CLOUDFRONT_STREAMING_DISTRIBUTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"CLOUDFRONT_STREAMING_DISTRIBUTION",
		&returns,
	)
	return returns
}

func ResourceType_CLOUDTRAIL_TRAIL() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"CLOUDTRAIL_TRAIL",
		&returns,
	)
	return returns
}

func ResourceType_CLOUDWATCH_ALARM() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"CLOUDWATCH_ALARM",
		&returns,
	)
	return returns
}

func ResourceType_CODEBUILD_PROJECT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"CODEBUILD_PROJECT",
		&returns,
	)
	return returns
}

func ResourceType_CODEPIPELINE_PIPELINE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"CODEPIPELINE_PIPELINE",
		&returns,
	)
	return returns
}

func ResourceType_DYNAMODB_TABLE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"DYNAMODB_TABLE",
		&returns,
	)
	return returns
}

func ResourceType_EBS_VOLUME() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EBS_VOLUME",
		&returns,
	)
	return returns
}

func ResourceType_EC2_CUSTOMER_GATEWAY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_CUSTOMER_GATEWAY",
		&returns,
	)
	return returns
}

func ResourceType_EC2_EGRESS_ONLY_INTERNET_GATEWAY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_EGRESS_ONLY_INTERNET_GATEWAY",
		&returns,
	)
	return returns
}

func ResourceType_EC2_EIP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_EIP",
		&returns,
	)
	return returns
}

func ResourceType_EC2_FLOW_LOG() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_FLOW_LOG",
		&returns,
	)
	return returns
}

func ResourceType_EC2_HOST() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_HOST",
		&returns,
	)
	return returns
}

func ResourceType_EC2_INSTANCE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_INSTANCE",
		&returns,
	)
	return returns
}

func ResourceType_EC2_INTERNET_GATEWAY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_INTERNET_GATEWAY",
		&returns,
	)
	return returns
}

func ResourceType_EC2_NAT_GATEWAY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_NAT_GATEWAY",
		&returns,
	)
	return returns
}

func ResourceType_EC2_NETWORK_ACL() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_NETWORK_ACL",
		&returns,
	)
	return returns
}

func ResourceType_EC2_ROUTE_TABLE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_ROUTE_TABLE",
		&returns,
	)
	return returns
}

func ResourceType_EC2_SECURITY_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_SECURITY_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_EC2_SUBNET() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_SUBNET",
		&returns,
	)
	return returns
}

func ResourceType_EC2_VPC() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_VPC",
		&returns,
	)
	return returns
}

func ResourceType_EC2_VPC_ENDPOINT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_VPC_ENDPOINT",
		&returns,
	)
	return returns
}

func ResourceType_EC2_VPC_ENDPOINT_SERVICE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_VPC_ENDPOINT_SERVICE",
		&returns,
	)
	return returns
}

func ResourceType_EC2_VPC_PEERING_CONNECTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_VPC_PEERING_CONNECTION",
		&returns,
	)
	return returns
}

func ResourceType_EC2_VPN_CONNECTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_VPN_CONNECTION",
		&returns,
	)
	return returns
}

func ResourceType_EC2_VPN_GATEWAY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"EC2_VPN_GATEWAY",
		&returns,
	)
	return returns
}

func ResourceType_ELASTIC_BEANSTALK_APPLICATION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"ELASTIC_BEANSTALK_APPLICATION",
		&returns,
	)
	return returns
}

func ResourceType_ELASTIC_BEANSTALK_APPLICATION_VERSION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"ELASTIC_BEANSTALK_APPLICATION_VERSION",
		&returns,
	)
	return returns
}

func ResourceType_ELASTIC_BEANSTALK_ENVIRONMENT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"ELASTIC_BEANSTALK_ENVIRONMENT",
		&returns,
	)
	return returns
}

func ResourceType_ELASTICSEARCH_DOMAIN() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"ELASTICSEARCH_DOMAIN",
		&returns,
	)
	return returns
}

func ResourceType_ELB_LOAD_BALANCER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"ELB_LOAD_BALANCER",
		&returns,
	)
	return returns
}

func ResourceType_ELBV2_LOAD_BALANCER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"ELBV2_LOAD_BALANCER",
		&returns,
	)
	return returns
}

func ResourceType_IAM_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"IAM_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_IAM_POLICY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"IAM_POLICY",
		&returns,
	)
	return returns
}

func ResourceType_IAM_ROLE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"IAM_ROLE",
		&returns,
	)
	return returns
}

func ResourceType_IAM_USER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"IAM_USER",
		&returns,
	)
	return returns
}

func ResourceType_KMS_KEY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"KMS_KEY",
		&returns,
	)
	return returns
}

func ResourceType_LAMBDA_FUNCTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"LAMBDA_FUNCTION",
		&returns,
	)
	return returns
}

func ResourceType_QLDB_LEDGER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"QLDB_LEDGER",
		&returns,
	)
	return returns
}

func ResourceType_RDS_DB_CLUSTER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"RDS_DB_CLUSTER",
		&returns,
	)
	return returns
}

func ResourceType_RDS_DB_CLUSTER_SNAPSHOT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"RDS_DB_CLUSTER_SNAPSHOT",
		&returns,
	)
	return returns
}

func ResourceType_RDS_DB_INSTANCE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"RDS_DB_INSTANCE",
		&returns,
	)
	return returns
}

func ResourceType_RDS_DB_SECURITY_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"RDS_DB_SECURITY_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_RDS_DB_SNAPSHOT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"RDS_DB_SNAPSHOT",
		&returns,
	)
	return returns
}

func ResourceType_RDS_DB_SUBNET_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"RDS_DB_SUBNET_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_RDS_EVENT_SUBSCRIPTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"RDS_EVENT_SUBSCRIPTION",
		&returns,
	)
	return returns
}

func ResourceType_REDSHIFT_CLUSTER() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"REDSHIFT_CLUSTER",
		&returns,
	)
	return returns
}

func ResourceType_REDSHIFT_CLUSTER_PARAMETER_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"REDSHIFT_CLUSTER_PARAMETER_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_REDSHIFT_CLUSTER_SECURITY_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"REDSHIFT_CLUSTER_SECURITY_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_REDSHIFT_CLUSTER_SNAPSHOT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"REDSHIFT_CLUSTER_SNAPSHOT",
		&returns,
	)
	return returns
}

func ResourceType_REDSHIFT_CLUSTER_SUBNET_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"REDSHIFT_CLUSTER_SUBNET_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_REDSHIFT_EVENT_SUBSCRIPTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"REDSHIFT_EVENT_SUBSCRIPTION",
		&returns,
	)
	return returns
}

func ResourceType_S3_ACCOUNT_PUBLIC_ACCESS_BLOCK() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"S3_ACCOUNT_PUBLIC_ACCESS_BLOCK",
		&returns,
	)
	return returns
}

func ResourceType_S3_BUCKET() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"S3_BUCKET",
		&returns,
	)
	return returns
}

func ResourceType_SECRETS_MANAGER_SECRET() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"SECRETS_MANAGER_SECRET",
		&returns,
	)
	return returns
}

func ResourceType_SERVICE_CATALOG_CLOUDFORMATION_PRODUCT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"SERVICE_CATALOG_CLOUDFORMATION_PRODUCT",
		&returns,
	)
	return returns
}

func ResourceType_SERVICE_CATALOG_CLOUDFORMATION_PROVISIONED_PRODUCT() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"SERVICE_CATALOG_CLOUDFORMATION_PROVISIONED_PRODUCT",
		&returns,
	)
	return returns
}

func ResourceType_SERVICE_CATALOG_PORTFOLIO() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"SERVICE_CATALOG_PORTFOLIO",
		&returns,
	)
	return returns
}

func ResourceType_SHIELD_PROTECTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"SHIELD_PROTECTION",
		&returns,
	)
	return returns
}

func ResourceType_SHIELD_REGIONAL_PROTECTION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"SHIELD_REGIONAL_PROTECTION",
		&returns,
	)
	return returns
}

func ResourceType_SNS_TOPIC() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"SNS_TOPIC",
		&returns,
	)
	return returns
}

func ResourceType_SQS_QUEUE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"SQS_QUEUE",
		&returns,
	)
	return returns
}

func ResourceType_SYSTEMS_MANAGER_ASSOCIATION_COMPLIANCE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"SYSTEMS_MANAGER_ASSOCIATION_COMPLIANCE",
		&returns,
	)
	return returns
}

func ResourceType_SYSTEMS_MANAGER_FILE_DATA() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"SYSTEMS_MANAGER_FILE_DATA",
		&returns,
	)
	return returns
}

func ResourceType_SYSTEMS_MANAGER_MANAGED_INSTANCE_INVENTORY() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"SYSTEMS_MANAGER_MANAGED_INSTANCE_INVENTORY",
		&returns,
	)
	return returns
}

func ResourceType_SYSTEMS_MANAGER_PATCH_COMPLIANCE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"SYSTEMS_MANAGER_PATCH_COMPLIANCE",
		&returns,
	)
	return returns
}

func ResourceType_WAF_RATE_BASED_RULE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"WAF_RATE_BASED_RULE",
		&returns,
	)
	return returns
}

func ResourceType_WAF_REGIONAL_RATE_BASED_RULE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"WAF_REGIONAL_RATE_BASED_RULE",
		&returns,
	)
	return returns
}

func ResourceType_WAF_REGIONAL_RULE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"WAF_REGIONAL_RULE",
		&returns,
	)
	return returns
}

func ResourceType_WAF_REGIONAL_RULE_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"WAF_REGIONAL_RULE_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_WAF_REGIONAL_WEB_ACL() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"WAF_REGIONAL_WEB_ACL",
		&returns,
	)
	return returns
}

func ResourceType_WAF_RULE() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"WAF_RULE",
		&returns,
	)
	return returns
}

func ResourceType_WAF_RULE_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"WAF_RULE_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_WAF_WEB_ACL() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"WAF_WEB_ACL",
		&returns,
	)
	return returns
}

func ResourceType_WAFV2_MANAGED_RULE_SET() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"WAFV2_MANAGED_RULE_SET",
		&returns,
	)
	return returns
}

func ResourceType_WAFV2_RULE_GROUP() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"WAFV2_RULE_GROUP",
		&returns,
	)
	return returns
}

func ResourceType_WAFV2_WEB_ACL() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"WAFV2_WEB_ACL",
		&returns,
	)
	return returns
}

func ResourceType_XRAY_ENCRYPTION_CONFIGURATION() ResourceType {
	_init_.Initialize()
	var returns ResourceType
	_jsii_.StaticGet(
		"monocdk.aws_config.ResourceType",
		"XRAY_ENCRYPTION_CONFIGURATION",
		&returns,
	)
	return returns
}

// Construction properties for a new rule.
//
// TODO: EXAMPLE
//
// Experimental.
type RuleProps struct {
	// A name for the AWS Config rule.
	// Experimental.
	ConfigRuleName *string `json:"configRuleName" yaml:"configRuleName"`
	// A description about this AWS Config rule.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// Input parameter values that are passed to the AWS Config rule.
	// Experimental.
	InputParameters *map[string]interface{} `json:"inputParameters" yaml:"inputParameters"`
	// The maximum frequency at which the AWS Config rule runs evaluations.
	// Experimental.
	MaximumExecutionFrequency MaximumExecutionFrequency `json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// Defines which resources trigger an evaluation for an AWS Config rule.
	// Experimental.
	RuleScope RuleScope `json:"ruleScope" yaml:"ruleScope"`
}

// Determines which resources trigger an evaluation of an AWS Config rule.
//
// TODO: EXAMPLE
//
// Experimental.
type RuleScope interface {
	Key() *string
	ResourceId() *string
	ResourceTypes() *[]ResourceType
	Value() *string
}

// The jsii proxy struct for RuleScope
type jsiiProxy_RuleScope struct {
	_ byte // padding
}

func (j *jsiiProxy_RuleScope) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RuleScope) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RuleScope) ResourceTypes() *[]ResourceType {
	var returns *[]ResourceType
	_jsii_.Get(
		j,
		"resourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RuleScope) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// restricts scope of changes to a specific resource type or resource identifier.
// Experimental.
func RuleScope_FromResource(resourceType ResourceType, resourceId *string) RuleScope {
	_init_.Initialize()

	var returns RuleScope

	_jsii_.StaticInvoke(
		"monocdk.aws_config.RuleScope",
		"fromResource",
		[]interface{}{resourceType, resourceId},
		&returns,
	)

	return returns
}

// restricts scope of changes to specific resource types.
// Experimental.
func RuleScope_FromResources(resourceTypes *[]ResourceType) RuleScope {
	_init_.Initialize()

	var returns RuleScope

	_jsii_.StaticInvoke(
		"monocdk.aws_config.RuleScope",
		"fromResources",
		[]interface{}{resourceTypes},
		&returns,
	)

	return returns
}

// restricts scope of changes to a specific tag.
// Experimental.
func RuleScope_FromTag(key *string, value *string) RuleScope {
	_init_.Initialize()

	var returns RuleScope

	_jsii_.StaticInvoke(
		"monocdk.aws_config.RuleScope",
		"fromTag",
		[]interface{}{key, value},
		&returns,
	)

	return returns
}


package awsconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Config::OrganizationConfigRule`.
//
// Adds or updates an AWS Config rule for your entire organization to evaluate if your AWS resources comply with your desired configurations. For information on how many organization AWS Config rules you can have per account, see [*Service Limits*](https://docs.aws.amazon.com/config/latest/developerguide/configlimits.html) in the *AWS Config Developer Guide* .
//
// Only a management account and a delegated administrator can create or update an organization AWS Config rule. When calling the `OrganizationConfigRule` resource with a delegated administrator, you must ensure AWS Organizations `ListDelegatedAdministrator` permissions are added. An organization can have up to 3 delegated administrators.
//
// The `OrganizationConfigRule` resource enables organization service access through the `EnableAWSServiceAccess` action and creates a service-linked role `AWSServiceRoleForConfigMultiAccountSetup` in the management or delegated administrator account of your organization. The service-linked role is created only when the role does not exist in the caller account. AWS Config verifies the existence of role with `GetRole` action.
//
// To use the `OrganizationConfigRule` resource with delegated administrator, register a delegated administrator by calling AWS Organization `register-delegated-administrator` for `config-multiaccountsetup.amazonaws.com` .
//
// There are two types of rules: *AWS Config Managed Rules* and *AWS Config Custom Rules* . You can use `PutOrganizationConfigRule` to create both AWS Config Managed Rules and AWS Config Custom Rules.
//
// AWS Config Managed Rules are predefined, customizable rules created by AWS Config . For a list of managed rules, see [List of AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html) . If you are adding an AWS Config managed rule, you must specify the rule's identifier for the `RuleIdentifier` key.
//
// AWS Config Custom Rules are rules that you create from scratch. There are two ways to create AWS Config custom rules: with Lambda functions ( [AWS Lambda Developer Guide](https://docs.aws.amazon.com/config/latest/developerguide/gettingstarted-concepts.html#gettingstarted-concepts-function) ) and with Guard ( [Guard GitHub Repository](https://docs.aws.amazon.com/https://github.com/aws-cloudformation/cloudformation-guard) ), a policy-as-code language. AWS Config custom rules created with AWS Lambda are called *AWS Config Custom Lambda Rules* and AWS Config custom rules created with Guard are called *AWS Config Custom Policy Rules* .
//
// If you are adding a new AWS Config Custom Lambda rule, you first need to create an AWS Lambda function in the management account or a delegated administrator that the rule invokes to evaluate your resources. You also need to create an IAM role in the managed account that can be assumed by the Lambda function. When you use `PutOrganizationConfigRule` to add a Custom Lambda rule to AWS Config , you must specify the Amazon Resource Name (ARN) that AWS Lambda assigns to the function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOrganizationConfigRule := awscdk.Aws_config.NewCfnOrganizationConfigRule(this, jsii.String("MyCfnOrganizationConfigRule"), &cfnOrganizationConfigRuleProps{
//   	organizationConfigRuleName: jsii.String("organizationConfigRuleName"),
//
//   	// the properties below are optional
//   	excludedAccounts: []*string{
//   		jsii.String("excludedAccounts"),
//   	},
//   	organizationCustomCodeRuleMetadata: &organizationCustomCodeRuleMetadataProperty{
//   		codeText: jsii.String("codeText"),
//   		runtime: jsii.String("runtime"),
//
//   		// the properties below are optional
//   		debugLogDeliveryAccounts: []*string{
//   			jsii.String("debugLogDeliveryAccounts"),
//   		},
//   		description: jsii.String("description"),
//   		inputParameters: jsii.String("inputParameters"),
//   		maximumExecutionFrequency: jsii.String("maximumExecutionFrequency"),
//   		organizationConfigRuleTriggerTypes: []*string{
//   			jsii.String("organizationConfigRuleTriggerTypes"),
//   		},
//   		resourceIdScope: jsii.String("resourceIdScope"),
//   		resourceTypesScope: []*string{
//   			jsii.String("resourceTypesScope"),
//   		},
//   		tagKeyScope: jsii.String("tagKeyScope"),
//   		tagValueScope: jsii.String("tagValueScope"),
//   	},
//   	organizationCustomRuleMetadata: &organizationCustomRuleMetadataProperty{
//   		lambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   		organizationConfigRuleTriggerTypes: []*string{
//   			jsii.String("organizationConfigRuleTriggerTypes"),
//   		},
//
//   		// the properties below are optional
//   		description: jsii.String("description"),
//   		inputParameters: jsii.String("inputParameters"),
//   		maximumExecutionFrequency: jsii.String("maximumExecutionFrequency"),
//   		resourceIdScope: jsii.String("resourceIdScope"),
//   		resourceTypesScope: []*string{
//   			jsii.String("resourceTypesScope"),
//   		},
//   		tagKeyScope: jsii.String("tagKeyScope"),
//   		tagValueScope: jsii.String("tagValueScope"),
//   	},
//   	organizationManagedRuleMetadata: &organizationManagedRuleMetadataProperty{
//   		ruleIdentifier: jsii.String("ruleIdentifier"),
//
//   		// the properties below are optional
//   		description: jsii.String("description"),
//   		inputParameters: jsii.String("inputParameters"),
//   		maximumExecutionFrequency: jsii.String("maximumExecutionFrequency"),
//   		resourceIdScope: jsii.String("resourceIdScope"),
//   		resourceTypesScope: []*string{
//   			jsii.String("resourceTypesScope"),
//   		},
//   		tagKeyScope: jsii.String("tagKeyScope"),
//   		tagValueScope: jsii.String("tagValueScope"),
//   	},
//   })
//
type CfnOrganizationConfigRule interface {
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
	// A comma-separated list of accounts excluded from organization AWS Config rule.
	ExcludedAccounts() *[]*string
	SetExcludedAccounts(val *[]*string)
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
	// The name that you assign to organization AWS Config rule.
	OrganizationConfigRuleName() *string
	SetOrganizationConfigRuleName(val *string)
	// `AWS::Config::OrganizationConfigRule.OrganizationCustomCodeRuleMetadata`.
	OrganizationCustomCodeRuleMetadata() interface{}
	SetOrganizationCustomCodeRuleMetadata(val interface{})
	// An `OrganizationCustomRuleMetadata` object.
	OrganizationCustomRuleMetadata() interface{}
	SetOrganizationCustomRuleMetadata(val interface{})
	// An `OrganizationManagedRuleMetadata` object.
	OrganizationManagedRuleMetadata() interface{}
	SetOrganizationManagedRuleMetadata(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
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

func (j *jsiiProxy_CfnOrganizationConfigRule) Node() constructs.Node {
	var returns constructs.Node
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

func (j *jsiiProxy_CfnOrganizationConfigRule) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::Config::OrganizationConfigRule`.
func NewCfnOrganizationConfigRule(scope constructs.Construct, id *string, props *CfnOrganizationConfigRuleProps) CfnOrganizationConfigRule {
	_init_.Initialize()

	if err := validateNewCfnOrganizationConfigRuleParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnOrganizationConfigRule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_config.CfnOrganizationConfigRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Config::OrganizationConfigRule`.
func NewCfnOrganizationConfigRule_Override(c CfnOrganizationConfigRule, scope constructs.Construct, id *string, props *CfnOrganizationConfigRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_config.CfnOrganizationConfigRule",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnOrganizationConfigRule)SetExcludedAccounts(val *[]*string) {
	_jsii_.Set(
		j,
		"excludedAccounts",
		val,
	)
}

func (j *jsiiProxy_CfnOrganizationConfigRule)SetOrganizationConfigRuleName(val *string) {
	if err := j.validateSetOrganizationConfigRuleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationConfigRuleName",
		val,
	)
}

func (j *jsiiProxy_CfnOrganizationConfigRule)SetOrganizationCustomCodeRuleMetadata(val interface{}) {
	if err := j.validateSetOrganizationCustomCodeRuleMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationCustomCodeRuleMetadata",
		val,
	)
}

func (j *jsiiProxy_CfnOrganizationConfigRule)SetOrganizationCustomRuleMetadata(val interface{}) {
	if err := j.validateSetOrganizationCustomRuleMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationCustomRuleMetadata",
		val,
	)
}

func (j *jsiiProxy_CfnOrganizationConfigRule)SetOrganizationManagedRuleMetadata(val interface{}) {
	if err := j.validateSetOrganizationManagedRuleMetadataParameters(val); err != nil {
		panic(err)
	}
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
func CfnOrganizationConfigRule_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOrganizationConfigRule_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_config.CfnOrganizationConfigRule",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnOrganizationConfigRule_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnOrganizationConfigRule_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_config.CfnOrganizationConfigRule",
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
func CfnOrganizationConfigRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOrganizationConfigRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_config.CfnOrganizationConfigRule",
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
		"aws-cdk-lib.aws_config.CfnOrganizationConfigRule",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnOrganizationConfigRule) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnOrganizationConfigRule) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnOrganizationConfigRule) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnOrganizationConfigRule) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnOrganizationConfigRule) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnOrganizationConfigRule) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnOrganizationConfigRule) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnOrganizationConfigRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnOrganizationConfigRule) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnOrganizationConfigRule) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnOrganizationConfigRule) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnOrganizationConfigRule) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOrganizationConfigRule) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOrganizationConfigRule) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnOrganizationConfigRule) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnOrganizationConfigRule) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnOrganizationConfigRule) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

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

func (c *jsiiProxy_CfnOrganizationConfigRule) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}


package awsdlm

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdlm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::DLM::LifecyclePolicy`.
//
// Specifies a lifecycle policy, which is used to automate operations on Amazon EBS resources.
//
// The properties are required when you add a lifecycle policy and optional when you update a lifecycle policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLifecyclePolicy := awscdk.Aws_dlm.NewCfnLifecyclePolicy(this, jsii.String("MyCfnLifecyclePolicy"), &cfnLifecyclePolicyProps{
//   	description: jsii.String("description"),
//   	executionRoleArn: jsii.String("executionRoleArn"),
//   	policyDetails: &policyDetailsProperty{
//   		actions: []interface{}{
//   			&actionProperty{
//   				crossRegionCopy: []interface{}{
//   					&crossRegionCopyActionProperty{
//   						encryptionConfiguration: &encryptionConfigurationProperty{
//   							encrypted: jsii.Boolean(false),
//
//   							// the properties below are optional
//   							cmkArn: jsii.String("cmkArn"),
//   						},
//   						target: jsii.String("target"),
//
//   						// the properties below are optional
//   						retainRule: &crossRegionCopyRetainRuleProperty{
//   							interval: jsii.Number(123),
//   							intervalUnit: jsii.String("intervalUnit"),
//   						},
//   					},
//   				},
//   				name: jsii.String("name"),
//   			},
//   		},
//   		eventSource: &eventSourceProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			parameters: &eventParametersProperty{
//   				eventType: jsii.String("eventType"),
//   				snapshotOwner: []*string{
//   					jsii.String("snapshotOwner"),
//   				},
//
//   				// the properties below are optional
//   				descriptionRegex: jsii.String("descriptionRegex"),
//   			},
//   		},
//   		parameters: &parametersProperty{
//   			excludeBootVolume: jsii.Boolean(false),
//   			noReboot: jsii.Boolean(false),
//   		},
//   		policyType: jsii.String("policyType"),
//   		resourceLocations: []*string{
//   			jsii.String("resourceLocations"),
//   		},
//   		resourceTypes: []*string{
//   			jsii.String("resourceTypes"),
//   		},
//   		schedules: []interface{}{
//   			&scheduleProperty{
//   				copyTags: jsii.Boolean(false),
//   				createRule: &createRuleProperty{
//   					cronExpression: jsii.String("cronExpression"),
//   					interval: jsii.Number(123),
//   					intervalUnit: jsii.String("intervalUnit"),
//   					location: jsii.String("location"),
//   					times: []*string{
//   						jsii.String("times"),
//   					},
//   				},
//   				crossRegionCopyRules: []interface{}{
//   					&crossRegionCopyRuleProperty{
//   						encrypted: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						cmkArn: jsii.String("cmkArn"),
//   						copyTags: jsii.Boolean(false),
//   						deprecateRule: &crossRegionCopyDeprecateRuleProperty{
//   							interval: jsii.Number(123),
//   							intervalUnit: jsii.String("intervalUnit"),
//   						},
//   						retainRule: &crossRegionCopyRetainRuleProperty{
//   							interval: jsii.Number(123),
//   							intervalUnit: jsii.String("intervalUnit"),
//   						},
//   						target: jsii.String("target"),
//   						targetRegion: jsii.String("targetRegion"),
//   					},
//   				},
//   				deprecateRule: &deprecateRuleProperty{
//   					count: jsii.Number(123),
//   					interval: jsii.Number(123),
//   					intervalUnit: jsii.String("intervalUnit"),
//   				},
//   				fastRestoreRule: &fastRestoreRuleProperty{
//   					availabilityZones: []*string{
//   						jsii.String("availabilityZones"),
//   					},
//   					count: jsii.Number(123),
//   					interval: jsii.Number(123),
//   					intervalUnit: jsii.String("intervalUnit"),
//   				},
//   				name: jsii.String("name"),
//   				retainRule: &retainRuleProperty{
//   					count: jsii.Number(123),
//   					interval: jsii.Number(123),
//   					intervalUnit: jsii.String("intervalUnit"),
//   				},
//   				shareRules: []interface{}{
//   					&shareRuleProperty{
//   						targetAccounts: []*string{
//   							jsii.String("targetAccounts"),
//   						},
//   						unshareInterval: jsii.Number(123),
//   						unshareIntervalUnit: jsii.String("unshareIntervalUnit"),
//   					},
//   				},
//   				tagsToAdd: []interface{}{
//   					&cfnTag{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				variableTags: []interface{}{
//   					&cfnTag{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		targetTags: []interface{}{
//   			&cfnTag{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	state: jsii.String("state"),
//   	tags: []*cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnLifecyclePolicy interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the lifecycle policy.
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A description of the lifecycle policy.
	//
	// The characters ^[0-9A-Za-z _-]+$ are supported.
	Description() *string
	SetDescription(val *string)
	// The Amazon Resource Name (ARN) of the IAM role used to run the operations specified by the lifecycle policy.
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
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
	// The configuration details of the lifecycle policy.
	PolicyDetails() interface{}
	SetPolicyDetails(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The activation state of the lifecycle policy.
	State() *string
	SetState(val *string)
	// The tags to apply to the lifecycle policy during creation.
	Tags() awscdk.TagManager
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

// The jsii proxy struct for CfnLifecyclePolicy
type jsiiProxy_CfnLifecyclePolicy struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLifecyclePolicy) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) PolicyDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::DLM::LifecyclePolicy`.
func NewCfnLifecyclePolicy(scope constructs.Construct, id *string, props *CfnLifecyclePolicyProps) CfnLifecyclePolicy {
	_init_.Initialize()

	j := jsiiProxy_CfnLifecyclePolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DLM::LifecyclePolicy`.
func NewCfnLifecyclePolicy_Override(c CfnLifecyclePolicy, scope constructs.Construct, id *string, props *CfnLifecyclePolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLifecyclePolicy) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnLifecyclePolicy) SetExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnLifecyclePolicy) SetPolicyDetails(val interface{}) {
	_jsii_.Set(
		j,
		"policyDetails",
		val,
	)
}

func (j *jsiiProxy_CfnLifecyclePolicy) SetState(val *string) {
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLifecyclePolicy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLifecyclePolicy_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy",
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
func CfnLifecyclePolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLifecyclePolicy_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLifecyclePolicy) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnLifecyclePolicy) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLifecyclePolicy) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnLifecyclePolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnLifecyclePolicy) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnLifecyclePolicy) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnLifecyclePolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnLifecyclePolicy) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLifecyclePolicy) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLifecyclePolicy) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnLifecyclePolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLifecyclePolicy) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLifecyclePolicy) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLifecyclePolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLifecyclePolicy) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies an action for an event-based policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &actionProperty{
//   	crossRegionCopy: []interface{}{
//   		&crossRegionCopyActionProperty{
//   			encryptionConfiguration: &encryptionConfigurationProperty{
//   				encrypted: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				cmkArn: jsii.String("cmkArn"),
//   			},
//   			target: jsii.String("target"),
//
//   			// the properties below are optional
//   			retainRule: &crossRegionCopyRetainRuleProperty{
//   				interval: jsii.Number(123),
//   				intervalUnit: jsii.String("intervalUnit"),
//   			},
//   		},
//   	},
//   	name: jsii.String("name"),
//   }
//
type CfnLifecyclePolicy_ActionProperty struct {
	// The rule for copying shared snapshots across Regions.
	CrossRegionCopy interface{} `field:"required" json:"crossRegionCopy" yaml:"crossRegionCopy"`
	// A descriptive name for the action.
	Name *string `field:"required" json:"name" yaml:"name"`
}

// Specifies when to create snapshots of EBS volumes.
//
// You must specify either a Cron expression or an interval, interval unit, and start time. You cannot specify both.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   createRuleProperty := &createRuleProperty{
//   	cronExpression: jsii.String("cronExpression"),
//   	interval: jsii.Number(123),
//   	intervalUnit: jsii.String("intervalUnit"),
//   	location: jsii.String("location"),
//   	times: []*string{
//   		jsii.String("times"),
//   	},
//   }
//
type CfnLifecyclePolicy_CreateRuleProperty struct {
	// The schedule, as a Cron expression.
	//
	// The schedule interval must be between 1 hour and 1 year. For more information, see [Cron expressions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html#CronExpressions) in the *Amazon CloudWatch User Guide* .
	CronExpression *string `field:"optional" json:"cronExpression" yaml:"cronExpression"`
	// The interval between snapshots.
	//
	// The supported values are 1, 2, 3, 4, 6, 8, 12, and 24.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The interval unit.
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
	// Specifies the destination for snapshots created by the policy.
	//
	// To create snapshots in the same Region as the source resource, specify `CLOUD` . To create snapshots on the same Outpost as the source resource, specify `OUTPOST_LOCAL` . If you omit this parameter, `CLOUD` is used by default.
	//
	// If the policy targets resources in an AWS Region , then you must create snapshots in the same Region as the source resource.
	//
	// If the policy targets resources on an Outpost, then you can create snapshots on the same Outpost as the source resource, or in the Region of that Outpost.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The time, in UTC, to start the operation. The supported format is hh:mm.
	//
	// The operation occurs within a one-hour window following the specified time. If you do not specify a time, Amazon DLM selects a time within the next 24 hours.
	Times *[]*string `field:"optional" json:"times" yaml:"times"`
}

// Specifies a rule for copying shared snapshots across Regions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   crossRegionCopyActionProperty := &crossRegionCopyActionProperty{
//   	encryptionConfiguration: &encryptionConfigurationProperty{
//   		encrypted: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		cmkArn: jsii.String("cmkArn"),
//   	},
//   	target: jsii.String("target"),
//
//   	// the properties below are optional
//   	retainRule: &crossRegionCopyRetainRuleProperty{
//   		interval: jsii.Number(123),
//   		intervalUnit: jsii.String("intervalUnit"),
//   	},
//   }
//
type CfnLifecyclePolicy_CrossRegionCopyActionProperty struct {
	// The encryption settings for the copied snapshot.
	EncryptionConfiguration interface{} `field:"required" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The target Region.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Specifies the retention rule for cross-Region snapshot copies.
	RetainRule interface{} `field:"optional" json:"retainRule" yaml:"retainRule"`
}

// Specifies an AMI deprecation rule for cross-Region AMI copies created by a cross-Region copy rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   crossRegionCopyDeprecateRuleProperty := &crossRegionCopyDeprecateRuleProperty{
//   	interval: jsii.Number(123),
//   	intervalUnit: jsii.String("intervalUnit"),
//   }
//
type CfnLifecyclePolicy_CrossRegionCopyDeprecateRuleProperty struct {
	// The period after which to deprecate the cross-Region AMI copies.
	//
	// The period must be less than or equal to the cross-Region AMI copy retention period, and it can't be greater than 10 years. This is equivalent to 120 months, 520 weeks, or 3650 days.
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// The unit of time in which to measure the *Interval* .
	IntervalUnit *string `field:"required" json:"intervalUnit" yaml:"intervalUnit"`
}

// Specifies the retention rule for cross-Region snapshot copies.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   crossRegionCopyRetainRuleProperty := &crossRegionCopyRetainRuleProperty{
//   	interval: jsii.Number(123),
//   	intervalUnit: jsii.String("intervalUnit"),
//   }
//
type CfnLifecyclePolicy_CrossRegionCopyRetainRuleProperty struct {
	// The amount of time to retain each snapshot.
	//
	// The maximum is 100 years. This is equivalent to 1200 months, 5200 weeks, or 36500 days.
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// The unit of time for time-based retention.
	IntervalUnit *string `field:"required" json:"intervalUnit" yaml:"intervalUnit"`
}

// Specifies a rule for cross-Region snapshot copies.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   crossRegionCopyRuleProperty := &crossRegionCopyRuleProperty{
//   	encrypted: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	cmkArn: jsii.String("cmkArn"),
//   	copyTags: jsii.Boolean(false),
//   	deprecateRule: &crossRegionCopyDeprecateRuleProperty{
//   		interval: jsii.Number(123),
//   		intervalUnit: jsii.String("intervalUnit"),
//   	},
//   	retainRule: &crossRegionCopyRetainRuleProperty{
//   		interval: jsii.Number(123),
//   		intervalUnit: jsii.String("intervalUnit"),
//   	},
//   	target: jsii.String("target"),
//   	targetRegion: jsii.String("targetRegion"),
//   }
//
type CfnLifecyclePolicy_CrossRegionCopyRuleProperty struct {
	// To encrypt a copy of an unencrypted snapshot if encryption by default is not enabled, enable encryption using this parameter.
	//
	// Copies of encrypted snapshots are encrypted, even if this parameter is false or if encryption by default is not enabled.
	Encrypted interface{} `field:"required" json:"encrypted" yaml:"encrypted"`
	// The Amazon Resource Name (ARN) of the AWS KMS key to use for EBS encryption.
	//
	// If this parameter is not specified, the default KMS key for the account is used.
	CmkArn *string `field:"optional" json:"cmkArn" yaml:"cmkArn"`
	// Indicates whether to copy all user-defined tags from the source snapshot to the cross-Region snapshot copy.
	CopyTags interface{} `field:"optional" json:"copyTags" yaml:"copyTags"`
	// The AMI deprecation rule for cross-Region AMI copies created by the rule.
	DeprecateRule interface{} `field:"optional" json:"deprecateRule" yaml:"deprecateRule"`
	// The retention rule that indicates how long snapshot copies are to be retained in the destination Region.
	RetainRule interface{} `field:"optional" json:"retainRule" yaml:"retainRule"`
	// The target Region or the Amazon Resource Name (ARN) of the target Outpost for the snapshot copies.
	//
	// Use this parameter instead of *TargetRegion* . Do not specify both.
	Target *string `field:"optional" json:"target" yaml:"target"`
	// Avoid using this parameter when creating new policies.
	//
	// Instead, use *Target* to specify a target Region or a target Outpost for snapshot copies.
	//
	// For policies created before the *Target* parameter was introduced, this parameter indicates the target Region for snapshot copies.
	TargetRegion *string `field:"optional" json:"targetRegion" yaml:"targetRegion"`
}

// Specifies an AMI deprecation rule for a schedule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deprecateRuleProperty := &deprecateRuleProperty{
//   	count: jsii.Number(123),
//   	interval: jsii.Number(123),
//   	intervalUnit: jsii.String("intervalUnit"),
//   }
//
type CfnLifecyclePolicy_DeprecateRuleProperty struct {
	// If the schedule has a count-based retention rule, this parameter specifies the number of oldest AMIs to deprecate.
	//
	// The count must be less than or equal to the schedule's retention count, and it can't be greater than 1000.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// If the schedule has an age-based retention rule, this parameter specifies the period after which to deprecate AMIs created by the schedule.
	//
	// The period must be less than or equal to the schedule's retention period, and it can't be greater than 10 years. This is equivalent to 120 months, 520 weeks, or 3650 days.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The unit of time in which to measure the *Interval* .
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
}

// Specifies the encryption settings for shared snapshots that are copied across Regions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigurationProperty := &encryptionConfigurationProperty{
//   	encrypted: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	cmkArn: jsii.String("cmkArn"),
//   }
//
type CfnLifecyclePolicy_EncryptionConfigurationProperty struct {
	// To encrypt a copy of an unencrypted snapshot when encryption by default is not enabled, enable encryption using this parameter.
	//
	// Copies of encrypted snapshots are encrypted, even if this parameter is false or when encryption by default is not enabled.
	Encrypted interface{} `field:"required" json:"encrypted" yaml:"encrypted"`
	// The Amazon Resource Name (ARN) of the AWS KMS key to use for EBS encryption.
	//
	// If this parameter is not specified, the default KMS key for the account is used.
	CmkArn *string `field:"optional" json:"cmkArn" yaml:"cmkArn"`
}

// Specifies an event that triggers an event-based policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventParametersProperty := &eventParametersProperty{
//   	eventType: jsii.String("eventType"),
//   	snapshotOwner: []*string{
//   		jsii.String("snapshotOwner"),
//   	},
//
//   	// the properties below are optional
//   	descriptionRegex: jsii.String("descriptionRegex"),
//   }
//
type CfnLifecyclePolicy_EventParametersProperty struct {
	// The type of event.
	//
	// Currently, only snapshot sharing events are supported.
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// The IDs of the AWS accounts that can trigger policy by sharing snapshots with your account.
	//
	// The policy only runs if one of the specified AWS accounts shares a snapshot with your account.
	SnapshotOwner *[]*string `field:"required" json:"snapshotOwner" yaml:"snapshotOwner"`
	// The snapshot description that can trigger the policy.
	//
	// The description pattern is specified using a regular expression. The policy runs only if a snapshot with a description that matches the specified pattern is shared with your account.
	//
	// For example, specifying `^.*Created for policy: policy-1234567890abcdef0.*$` configures the policy to run only if snapshots created by policy `policy-1234567890abcdef0` are shared with your account.
	DescriptionRegex *string `field:"optional" json:"descriptionRegex" yaml:"descriptionRegex"`
}

// Specifies an event that triggers an event-based policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventSourceProperty := &eventSourceProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	parameters: &eventParametersProperty{
//   		eventType: jsii.String("eventType"),
//   		snapshotOwner: []*string{
//   			jsii.String("snapshotOwner"),
//   		},
//
//   		// the properties below are optional
//   		descriptionRegex: jsii.String("descriptionRegex"),
//   	},
//   }
//
type CfnLifecyclePolicy_EventSourceProperty struct {
	// The source of the event.
	//
	// Currently only managed CloudWatch Events rules are supported.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Information about the event.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

// Specifies a rule for enabling fast snapshot restore.
//
// You can enable fast snapshot restore based on either a count or a time interval.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fastRestoreRuleProperty := &fastRestoreRuleProperty{
//   	availabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	count: jsii.Number(123),
//   	interval: jsii.Number(123),
//   	intervalUnit: jsii.String("intervalUnit"),
//   }
//
type CfnLifecyclePolicy_FastRestoreRuleProperty struct {
	// The Availability Zones in which to enable fast snapshot restore.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// The number of snapshots to be enabled with fast snapshot restore.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// The amount of time to enable fast snapshot restore.
	//
	// The maximum is 100 years. This is equivalent to 1200 months, 5200 weeks, or 36500 days.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The unit of time for enabling fast snapshot restore.
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
}

// Specifies optional parameters to add to a policy.
//
// The set of valid parameters depends on the combination of policy type and resource type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parametersProperty := &parametersProperty{
//   	excludeBootVolume: jsii.Boolean(false),
//   	noReboot: jsii.Boolean(false),
//   }
//
type CfnLifecyclePolicy_ParametersProperty struct {
	// [EBS Snapshot Management – Instance policies only] Indicates whether to exclude the root volume from snapshots created using [CreateSnapshots](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateSnapshots.html) . The default is false.
	ExcludeBootVolume interface{} `field:"optional" json:"excludeBootVolume" yaml:"excludeBootVolume"`
	// Applies to AMI lifecycle policies only.
	//
	// Indicates whether targeted instances are rebooted when the lifecycle policy runs. `true` indicates that targeted instances are not rebooted when the policy runs. `false` indicates that target instances are rebooted when the policy runs. The default is `true` (instances are not rebooted).
	NoReboot interface{} `field:"optional" json:"noReboot" yaml:"noReboot"`
}

// Specifies the configuration of a lifecycle policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyDetailsProperty := &policyDetailsProperty{
//   	actions: []interface{}{
//   		&actionProperty{
//   			crossRegionCopy: []interface{}{
//   				&crossRegionCopyActionProperty{
//   					encryptionConfiguration: &encryptionConfigurationProperty{
//   						encrypted: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						cmkArn: jsii.String("cmkArn"),
//   					},
//   					target: jsii.String("target"),
//
//   					// the properties below are optional
//   					retainRule: &crossRegionCopyRetainRuleProperty{
//   						interval: jsii.Number(123),
//   						intervalUnit: jsii.String("intervalUnit"),
//   					},
//   				},
//   			},
//   			name: jsii.String("name"),
//   		},
//   	},
//   	eventSource: &eventSourceProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		parameters: &eventParametersProperty{
//   			eventType: jsii.String("eventType"),
//   			snapshotOwner: []*string{
//   				jsii.String("snapshotOwner"),
//   			},
//
//   			// the properties below are optional
//   			descriptionRegex: jsii.String("descriptionRegex"),
//   		},
//   	},
//   	parameters: &parametersProperty{
//   		excludeBootVolume: jsii.Boolean(false),
//   		noReboot: jsii.Boolean(false),
//   	},
//   	policyType: jsii.String("policyType"),
//   	resourceLocations: []*string{
//   		jsii.String("resourceLocations"),
//   	},
//   	resourceTypes: []*string{
//   		jsii.String("resourceTypes"),
//   	},
//   	schedules: []interface{}{
//   		&scheduleProperty{
//   			copyTags: jsii.Boolean(false),
//   			createRule: &createRuleProperty{
//   				cronExpression: jsii.String("cronExpression"),
//   				interval: jsii.Number(123),
//   				intervalUnit: jsii.String("intervalUnit"),
//   				location: jsii.String("location"),
//   				times: []*string{
//   					jsii.String("times"),
//   				},
//   			},
//   			crossRegionCopyRules: []interface{}{
//   				&crossRegionCopyRuleProperty{
//   					encrypted: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					cmkArn: jsii.String("cmkArn"),
//   					copyTags: jsii.Boolean(false),
//   					deprecateRule: &crossRegionCopyDeprecateRuleProperty{
//   						interval: jsii.Number(123),
//   						intervalUnit: jsii.String("intervalUnit"),
//   					},
//   					retainRule: &crossRegionCopyRetainRuleProperty{
//   						interval: jsii.Number(123),
//   						intervalUnit: jsii.String("intervalUnit"),
//   					},
//   					target: jsii.String("target"),
//   					targetRegion: jsii.String("targetRegion"),
//   				},
//   			},
//   			deprecateRule: &deprecateRuleProperty{
//   				count: jsii.Number(123),
//   				interval: jsii.Number(123),
//   				intervalUnit: jsii.String("intervalUnit"),
//   			},
//   			fastRestoreRule: &fastRestoreRuleProperty{
//   				availabilityZones: []*string{
//   					jsii.String("availabilityZones"),
//   				},
//   				count: jsii.Number(123),
//   				interval: jsii.Number(123),
//   				intervalUnit: jsii.String("intervalUnit"),
//   			},
//   			name: jsii.String("name"),
//   			retainRule: &retainRuleProperty{
//   				count: jsii.Number(123),
//   				interval: jsii.Number(123),
//   				intervalUnit: jsii.String("intervalUnit"),
//   			},
//   			shareRules: []interface{}{
//   				&shareRuleProperty{
//   					targetAccounts: []*string{
//   						jsii.String("targetAccounts"),
//   					},
//   					unshareInterval: jsii.Number(123),
//   					unshareIntervalUnit: jsii.String("unshareIntervalUnit"),
//   				},
//   			},
//   			tagsToAdd: []interface{}{
//   				&cfnTag{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			variableTags: []interface{}{
//   				&cfnTag{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	targetTags: []interface{}{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLifecyclePolicy_PolicyDetailsProperty struct {
	// The actions to be performed when the event-based policy is triggered. You can specify only one action per policy.
	//
	// This parameter is required for event-based policies only. If you are creating a snapshot or AMI policy, omit this parameter.
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The event that triggers the event-based policy.
	//
	// This parameter is required for event-based policies only. If you are creating a snapshot or AMI policy, omit this parameter.
	EventSource interface{} `field:"optional" json:"eventSource" yaml:"eventSource"`
	// A set of optional parameters for snapshot and AMI lifecycle policies.
	//
	// This parameter is required for snapshot and AMI policies only. If you are creating an event-based policy, omit this parameter.
	//
	// > If you are modifying a policy that was created or previously modified using the Amazon Data Lifecycle Manager console, then you must include this parameter and specify either the default values or the new values that you require. You can't omit this parameter or set its values to null.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The valid target resource types and actions a policy can manage.
	//
	// Specify `EBS_SNAPSHOT_MANAGEMENT` to create a lifecycle policy that manages the lifecycle of Amazon EBS snapshots. Specify `IMAGE_MANAGEMENT` to create a lifecycle policy that manages the lifecycle of EBS-backed AMIs. Specify `EVENT_BASED_POLICY` to create an event-based policy that performs specific actions when a defined event occurs in your AWS account .
	//
	// The default is `EBS_SNAPSHOT_MANAGEMENT` .
	PolicyType *string `field:"optional" json:"policyType" yaml:"policyType"`
	// The location of the resources to backup.
	//
	// If the source resources are located in an AWS Region , specify `CLOUD` . If the source resources are located on an Outpost in your account, specify `OUTPOST` .
	//
	// If you specify `OUTPOST` , Amazon Data Lifecycle Manager backs up all resources of the specified type with matching target tags across all of the Outposts in your account.
	ResourceLocations *[]*string `field:"optional" json:"resourceLocations" yaml:"resourceLocations"`
	// The target resource type for snapshot and AMI lifecycle policies.
	//
	// Use `VOLUME` to create snapshots of individual volumes or use `INSTANCE` to create multi-volume snapshots from the volumes for an instance.
	//
	// This parameter is required for snapshot and AMI policies only. If you are creating an event-based policy, omit this parameter.
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
	// The schedules of policy-defined actions for snapshot and AMI lifecycle policies.
	//
	// A policy can have up to four schedules—one mandatory schedule and up to three optional schedules.
	//
	// This parameter is required for snapshot and AMI policies only. If you are creating an event-based policy, omit this parameter.
	Schedules interface{} `field:"optional" json:"schedules" yaml:"schedules"`
	// The single tag that identifies targeted resources for this policy.
	//
	// This parameter is required for snapshot and AMI policies only. If you are creating an event-based policy, omit this parameter.
	TargetTags interface{} `field:"optional" json:"targetTags" yaml:"targetTags"`
}

// Specifies the retention rule for a lifecycle policy.
//
// You can retain snapshots based on either a count or a time interval.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   retainRuleProperty := &retainRuleProperty{
//   	count: jsii.Number(123),
//   	interval: jsii.Number(123),
//   	intervalUnit: jsii.String("intervalUnit"),
//   }
//
type CfnLifecyclePolicy_RetainRuleProperty struct {
	// The number of snapshots to retain for each volume, up to a maximum of 1000.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// The amount of time to retain each snapshot.
	//
	// The maximum is 100 years. This is equivalent to 1200 months, 5200 weeks, or 36500 days.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The unit of time for time-based retention.
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
}

// Specifies a backup schedule for a snapshot or AMI lifecycle policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleProperty := &scheduleProperty{
//   	copyTags: jsii.Boolean(false),
//   	createRule: &createRuleProperty{
//   		cronExpression: jsii.String("cronExpression"),
//   		interval: jsii.Number(123),
//   		intervalUnit: jsii.String("intervalUnit"),
//   		location: jsii.String("location"),
//   		times: []*string{
//   			jsii.String("times"),
//   		},
//   	},
//   	crossRegionCopyRules: []interface{}{
//   		&crossRegionCopyRuleProperty{
//   			encrypted: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			cmkArn: jsii.String("cmkArn"),
//   			copyTags: jsii.Boolean(false),
//   			deprecateRule: &crossRegionCopyDeprecateRuleProperty{
//   				interval: jsii.Number(123),
//   				intervalUnit: jsii.String("intervalUnit"),
//   			},
//   			retainRule: &crossRegionCopyRetainRuleProperty{
//   				interval: jsii.Number(123),
//   				intervalUnit: jsii.String("intervalUnit"),
//   			},
//   			target: jsii.String("target"),
//   			targetRegion: jsii.String("targetRegion"),
//   		},
//   	},
//   	deprecateRule: &deprecateRuleProperty{
//   		count: jsii.Number(123),
//   		interval: jsii.Number(123),
//   		intervalUnit: jsii.String("intervalUnit"),
//   	},
//   	fastRestoreRule: &fastRestoreRuleProperty{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		count: jsii.Number(123),
//   		interval: jsii.Number(123),
//   		intervalUnit: jsii.String("intervalUnit"),
//   	},
//   	name: jsii.String("name"),
//   	retainRule: &retainRuleProperty{
//   		count: jsii.Number(123),
//   		interval: jsii.Number(123),
//   		intervalUnit: jsii.String("intervalUnit"),
//   	},
//   	shareRules: []interface{}{
//   		&shareRuleProperty{
//   			targetAccounts: []*string{
//   				jsii.String("targetAccounts"),
//   			},
//   			unshareInterval: jsii.Number(123),
//   			unshareIntervalUnit: jsii.String("unshareIntervalUnit"),
//   		},
//   	},
//   	tagsToAdd: []interface{}{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	variableTags: []interface{}{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLifecyclePolicy_ScheduleProperty struct {
	// Copy all user-defined tags on a source volume to snapshots of the volume created by this policy.
	CopyTags interface{} `field:"optional" json:"copyTags" yaml:"copyTags"`
	// The creation rule.
	CreateRule interface{} `field:"optional" json:"createRule" yaml:"createRule"`
	// The rule for cross-Region snapshot copies.
	//
	// You can only specify cross-Region copy rules for policies that create snapshots in a Region. If the policy creates snapshots on an Outpost, then you cannot copy the snapshots to a Region or to an Outpost. If the policy creates snapshots in a Region, then snapshots can be copied to up to three Regions or Outposts.
	CrossRegionCopyRules interface{} `field:"optional" json:"crossRegionCopyRules" yaml:"crossRegionCopyRules"`
	// The AMI deprecation rule for the schedule.
	DeprecateRule interface{} `field:"optional" json:"deprecateRule" yaml:"deprecateRule"`
	// The rule for enabling fast snapshot restore.
	FastRestoreRule interface{} `field:"optional" json:"fastRestoreRule" yaml:"fastRestoreRule"`
	// The name of the schedule.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The retention rule.
	RetainRule interface{} `field:"optional" json:"retainRule" yaml:"retainRule"`
	// The rule for sharing snapshots with other AWS accounts .
	ShareRules interface{} `field:"optional" json:"shareRules" yaml:"shareRules"`
	// The tags to apply to policy-created resources.
	//
	// These user-defined tags are in addition to the AWS -added lifecycle tags.
	TagsToAdd interface{} `field:"optional" json:"tagsToAdd" yaml:"tagsToAdd"`
	// A collection of key/value pairs with values determined dynamically when the policy is executed.
	//
	// Keys may be any valid Amazon EC2 tag key. Values must be in one of the two following formats: `$(instance-id)` or `$(timestamp)` . Variable tags are only valid for EBS Snapshot Management – Instance policies.
	VariableTags interface{} `field:"optional" json:"variableTags" yaml:"variableTags"`
}

// Specifies a rule for sharing snapshots across AWS accounts .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   shareRuleProperty := &shareRuleProperty{
//   	targetAccounts: []*string{
//   		jsii.String("targetAccounts"),
//   	},
//   	unshareInterval: jsii.Number(123),
//   	unshareIntervalUnit: jsii.String("unshareIntervalUnit"),
//   }
//
type CfnLifecyclePolicy_ShareRuleProperty struct {
	// The IDs of the AWS accounts with which to share the snapshots.
	TargetAccounts *[]*string `field:"optional" json:"targetAccounts" yaml:"targetAccounts"`
	// The period after which snapshots that are shared with other AWS accounts are automatically unshared.
	UnshareInterval *float64 `field:"optional" json:"unshareInterval" yaml:"unshareInterval"`
	// The unit of time for the automatic unsharing interval.
	UnshareIntervalUnit *string `field:"optional" json:"unshareIntervalUnit" yaml:"unshareIntervalUnit"`
}

// Properties for defining a `CfnLifecyclePolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLifecyclePolicyProps := &cfnLifecyclePolicyProps{
//   	description: jsii.String("description"),
//   	executionRoleArn: jsii.String("executionRoleArn"),
//   	policyDetails: &policyDetailsProperty{
//   		actions: []interface{}{
//   			&actionProperty{
//   				crossRegionCopy: []interface{}{
//   					&crossRegionCopyActionProperty{
//   						encryptionConfiguration: &encryptionConfigurationProperty{
//   							encrypted: jsii.Boolean(false),
//
//   							// the properties below are optional
//   							cmkArn: jsii.String("cmkArn"),
//   						},
//   						target: jsii.String("target"),
//
//   						// the properties below are optional
//   						retainRule: &crossRegionCopyRetainRuleProperty{
//   							interval: jsii.Number(123),
//   							intervalUnit: jsii.String("intervalUnit"),
//   						},
//   					},
//   				},
//   				name: jsii.String("name"),
//   			},
//   		},
//   		eventSource: &eventSourceProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			parameters: &eventParametersProperty{
//   				eventType: jsii.String("eventType"),
//   				snapshotOwner: []*string{
//   					jsii.String("snapshotOwner"),
//   				},
//
//   				// the properties below are optional
//   				descriptionRegex: jsii.String("descriptionRegex"),
//   			},
//   		},
//   		parameters: &parametersProperty{
//   			excludeBootVolume: jsii.Boolean(false),
//   			noReboot: jsii.Boolean(false),
//   		},
//   		policyType: jsii.String("policyType"),
//   		resourceLocations: []*string{
//   			jsii.String("resourceLocations"),
//   		},
//   		resourceTypes: []*string{
//   			jsii.String("resourceTypes"),
//   		},
//   		schedules: []interface{}{
//   			&scheduleProperty{
//   				copyTags: jsii.Boolean(false),
//   				createRule: &createRuleProperty{
//   					cronExpression: jsii.String("cronExpression"),
//   					interval: jsii.Number(123),
//   					intervalUnit: jsii.String("intervalUnit"),
//   					location: jsii.String("location"),
//   					times: []*string{
//   						jsii.String("times"),
//   					},
//   				},
//   				crossRegionCopyRules: []interface{}{
//   					&crossRegionCopyRuleProperty{
//   						encrypted: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						cmkArn: jsii.String("cmkArn"),
//   						copyTags: jsii.Boolean(false),
//   						deprecateRule: &crossRegionCopyDeprecateRuleProperty{
//   							interval: jsii.Number(123),
//   							intervalUnit: jsii.String("intervalUnit"),
//   						},
//   						retainRule: &crossRegionCopyRetainRuleProperty{
//   							interval: jsii.Number(123),
//   							intervalUnit: jsii.String("intervalUnit"),
//   						},
//   						target: jsii.String("target"),
//   						targetRegion: jsii.String("targetRegion"),
//   					},
//   				},
//   				deprecateRule: &deprecateRuleProperty{
//   					count: jsii.Number(123),
//   					interval: jsii.Number(123),
//   					intervalUnit: jsii.String("intervalUnit"),
//   				},
//   				fastRestoreRule: &fastRestoreRuleProperty{
//   					availabilityZones: []*string{
//   						jsii.String("availabilityZones"),
//   					},
//   					count: jsii.Number(123),
//   					interval: jsii.Number(123),
//   					intervalUnit: jsii.String("intervalUnit"),
//   				},
//   				name: jsii.String("name"),
//   				retainRule: &retainRuleProperty{
//   					count: jsii.Number(123),
//   					interval: jsii.Number(123),
//   					intervalUnit: jsii.String("intervalUnit"),
//   				},
//   				shareRules: []interface{}{
//   					&shareRuleProperty{
//   						targetAccounts: []*string{
//   							jsii.String("targetAccounts"),
//   						},
//   						unshareInterval: jsii.Number(123),
//   						unshareIntervalUnit: jsii.String("unshareIntervalUnit"),
//   					},
//   				},
//   				tagsToAdd: []interface{}{
//   					&cfnTag{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				variableTags: []interface{}{
//   					&cfnTag{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		targetTags: []interface{}{
//   			&cfnTag{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	state: jsii.String("state"),
//   	tags: []*cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLifecyclePolicyProps struct {
	// A description of the lifecycle policy.
	//
	// The characters ^[0-9A-Za-z _-]+$ are supported.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of the IAM role used to run the operations specified by the lifecycle policy.
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The configuration details of the lifecycle policy.
	PolicyDetails interface{} `field:"optional" json:"policyDetails" yaml:"policyDetails"`
	// The activation state of the lifecycle policy.
	State *string `field:"optional" json:"state" yaml:"state"`
	// The tags to apply to the lifecycle policy during creation.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


package awsxray

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsxray/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::XRay::Group`.
//
// Use the `AWS::XRay::Group` resource to specify a group with a name and a filter expression.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnGroup := awscdk.Aws_xray.NewCfnGroup(this, jsii.String("MyCfnGroup"), &cfnGroupProps{
//   	filterExpression: jsii.String("filterExpression"),
//   	groupName: jsii.String("groupName"),
//   	insightsConfiguration: &insightsConfigurationProperty{
//   		insightsEnabled: jsii.Boolean(false),
//   		notificationsEnabled: jsii.Boolean(false),
//   	},
//   	tags: []interface{}{
//   		tags,
//   	},
//   })
//
type CfnGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The group ARN that was created or updated.
	AttrGroupArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The filter expression defining the parameters to include traces.
	FilterExpression() *string
	SetFilterExpression(val *string)
	// The unique case-sensitive name of the group.
	GroupName() *string
	SetGroupName(val *string)
	// The structure containing configurations related to insights.
	//
	// - The InsightsEnabled boolean can be set to true to enable insights for the group or false to disable insights for the group.
	// - The NotificationsEnabled boolean can be set to true to enable insights notifications through Amazon EventBridge for the group.
	InsightsConfiguration() interface{}
	SetInsightsConfiguration(val interface{})
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
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() *[]interface{}
	SetTags(val *[]interface{})
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

// The jsii proxy struct for CfnGroup
type jsiiProxy_CfnGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnGroup) AttrGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) FilterExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) GroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) InsightsConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insightsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) Tags() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGroup) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::XRay::Group`.
func NewCfnGroup(scope constructs.Construct, id *string, props *CfnGroupProps) CfnGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_xray.CfnGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::XRay::Group`.
func NewCfnGroup_Override(c CfnGroup, scope constructs.Construct, id *string, props *CfnGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_xray.CfnGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnGroup) SetFilterExpression(val *string) {
	_jsii_.Set(
		j,
		"filterExpression",
		val,
	)
}

func (j *jsiiProxy_CfnGroup) SetGroupName(val *string) {
	_jsii_.Set(
		j,
		"groupName",
		val,
	)
}

func (j *jsiiProxy_CfnGroup) SetInsightsConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"insightsConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnGroup) SetTags(val *[]interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_xray.CfnGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_xray.CfnGroup",
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
func CfnGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_xray.CfnGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_xray.CfnGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The structure containing configurations related to insights.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   insightsConfigurationProperty := &insightsConfigurationProperty{
//   	insightsEnabled: jsii.Boolean(false),
//   	notificationsEnabled: jsii.Boolean(false),
//   }
//
type CfnGroup_InsightsConfigurationProperty struct {
	// Set the InsightsEnabled value to true to enable insights or false to disable insights.
	InsightsEnabled interface{} `field:"optional" json:"insightsEnabled" yaml:"insightsEnabled"`
	// Set the NotificationsEnabled value to true to enable insights notifications.
	//
	// Notifications can only be enabled on a group with InsightsEnabled set to true.
	NotificationsEnabled interface{} `field:"optional" json:"notificationsEnabled" yaml:"notificationsEnabled"`
}

// Properties for defining a `CfnGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnGroupProps := &cfnGroupProps{
//   	filterExpression: jsii.String("filterExpression"),
//   	groupName: jsii.String("groupName"),
//   	insightsConfiguration: &insightsConfigurationProperty{
//   		insightsEnabled: jsii.Boolean(false),
//   		notificationsEnabled: jsii.Boolean(false),
//   	},
//   	tags: []interface{}{
//   		tags,
//   	},
//   }
//
type CfnGroupProps struct {
	// The filter expression defining the parameters to include traces.
	FilterExpression *string `field:"optional" json:"filterExpression" yaml:"filterExpression"`
	// The unique case-sensitive name of the group.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// The structure containing configurations related to insights.
	//
	// - The InsightsEnabled boolean can be set to true to enable insights for the group or false to disable insights for the group.
	// - The NotificationsEnabled boolean can be set to true to enable insights notifications through Amazon EventBridge for the group.
	InsightsConfiguration interface{} `field:"optional" json:"insightsConfiguration" yaml:"insightsConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]interface{} `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::XRay::SamplingRule`.
//
// Use the `AWS::XRay::SamplingRule` resource to specify a sampling rule, which controls sampling behavior for instrumented applications. A new sampling rule is created by specifying a `SamplingRule` . To change the configuration of an existing sampling rule, specify a `SamplingRuleUpdate` .
//
// Services retrieve rules with [GetSamplingRules](https://docs.aws.amazon.com//xray/latest/api/API_GetSamplingRules.html) , and evaluate each rule in ascending order of *priority* for each request. If a rule matches, the service records a trace, borrowing it from the reservoir size. After 10 seconds, the service reports back to X-Ray with [GetSamplingTargets](https://docs.aws.amazon.com//xray/latest/api/API_GetSamplingTargets.html) to get updated versions of each in-use rule. The updated rule contains a trace quota that the service can use instead of borrowing from the reservoir.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnSamplingRule := awscdk.Aws_xray.NewCfnSamplingRule(this, jsii.String("MyCfnSamplingRule"), &cfnSamplingRuleProps{
//   	ruleName: jsii.String("ruleName"),
//   	samplingRule: &samplingRuleProperty{
//   		attributes: map[string]*string{
//   			"attributesKey": jsii.String("attributes"),
//   		},
//   		fixedRate: jsii.Number(123),
//   		host: jsii.String("host"),
//   		httpMethod: jsii.String("httpMethod"),
//   		priority: jsii.Number(123),
//   		reservoirSize: jsii.Number(123),
//   		resourceArn: jsii.String("resourceArn"),
//   		ruleArn: jsii.String("ruleArn"),
//   		ruleName: jsii.String("ruleName"),
//   		serviceName: jsii.String("serviceName"),
//   		serviceType: jsii.String("serviceType"),
//   		urlPath: jsii.String("urlPath"),
//   		version: jsii.Number(123),
//   	},
//   	samplingRuleRecord: &samplingRuleRecordProperty{
//   		createdAt: jsii.String("createdAt"),
//   		modifiedAt: jsii.String("modifiedAt"),
//   		samplingRule: &samplingRuleProperty{
//   			attributes: map[string]*string{
//   				"attributesKey": jsii.String("attributes"),
//   			},
//   			fixedRate: jsii.Number(123),
//   			host: jsii.String("host"),
//   			httpMethod: jsii.String("httpMethod"),
//   			priority: jsii.Number(123),
//   			reservoirSize: jsii.Number(123),
//   			resourceArn: jsii.String("resourceArn"),
//   			ruleArn: jsii.String("ruleArn"),
//   			ruleName: jsii.String("ruleName"),
//   			serviceName: jsii.String("serviceName"),
//   			serviceType: jsii.String("serviceType"),
//   			urlPath: jsii.String("urlPath"),
//   			version: jsii.Number(123),
//   		},
//   	},
//   	samplingRuleUpdate: &samplingRuleUpdateProperty{
//   		attributes: map[string]*string{
//   			"attributesKey": jsii.String("attributes"),
//   		},
//   		fixedRate: jsii.Number(123),
//   		host: jsii.String("host"),
//   		httpMethod: jsii.String("httpMethod"),
//   		priority: jsii.Number(123),
//   		reservoirSize: jsii.Number(123),
//   		resourceArn: jsii.String("resourceArn"),
//   		ruleArn: jsii.String("ruleArn"),
//   		ruleName: jsii.String("ruleName"),
//   		serviceName: jsii.String("serviceName"),
//   		serviceType: jsii.String("serviceType"),
//   		urlPath: jsii.String("urlPath"),
//   	},
//   	tags: []interface{}{
//   		tags,
//   	},
//   })
//
type CfnSamplingRule interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The sampling rule ARN that was created or updated.
	AttrRuleArn() *string
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
	// The name of the sampling rule.
	//
	// Specify a rule by either name or ARN, but not both. Used only when deleting a sampling rule. When creating or updating a sampling rule, use the `RuleName` or `RuleARN` properties within `SamplingRule` or `SamplingRuleUpdate` .
	RuleName() *string
	SetRuleName(val *string)
	// The sampling rule to be created.
	//
	// Must be provided if creating a new sampling rule. Not valid when updating an existing sampling rule.
	SamplingRule() interface{}
	SetSamplingRule(val interface{})
	// `AWS::XRay::SamplingRule.SamplingRuleRecord`.
	SamplingRuleRecord() interface{}
	SetSamplingRuleRecord(val interface{})
	// A document specifying changes to a sampling rule's configuration.
	//
	// Must be provided if updating an existing sampling rule. Not valid when creating a new sampling rule.
	//
	// > The `Version` of a sampling rule cannot be updated, and is not part of `SamplingRuleUpdate` .
	SamplingRuleUpdate() interface{}
	SetSamplingRuleUpdate(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() *[]interface{}
	SetTags(val *[]interface{})
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

// The jsii proxy struct for CfnSamplingRule
type jsiiProxy_CfnSamplingRule struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSamplingRule) AttrRuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRuleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSamplingRule) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSamplingRule) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSamplingRule) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSamplingRule) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSamplingRule) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSamplingRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSamplingRule) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSamplingRule) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSamplingRule) SamplingRule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"samplingRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSamplingRule) SamplingRuleRecord() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"samplingRuleRecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSamplingRule) SamplingRuleUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"samplingRuleUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSamplingRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSamplingRule) Tags() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSamplingRule) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSamplingRule) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::XRay::SamplingRule`.
func NewCfnSamplingRule(scope constructs.Construct, id *string, props *CfnSamplingRuleProps) CfnSamplingRule {
	_init_.Initialize()

	j := jsiiProxy_CfnSamplingRule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_xray.CfnSamplingRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::XRay::SamplingRule`.
func NewCfnSamplingRule_Override(c CfnSamplingRule, scope constructs.Construct, id *string, props *CfnSamplingRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_xray.CfnSamplingRule",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSamplingRule) SetRuleName(val *string) {
	_jsii_.Set(
		j,
		"ruleName",
		val,
	)
}

func (j *jsiiProxy_CfnSamplingRule) SetSamplingRule(val interface{}) {
	_jsii_.Set(
		j,
		"samplingRule",
		val,
	)
}

func (j *jsiiProxy_CfnSamplingRule) SetSamplingRuleRecord(val interface{}) {
	_jsii_.Set(
		j,
		"samplingRuleRecord",
		val,
	)
}

func (j *jsiiProxy_CfnSamplingRule) SetSamplingRuleUpdate(val interface{}) {
	_jsii_.Set(
		j,
		"samplingRuleUpdate",
		val,
	)
}

func (j *jsiiProxy_CfnSamplingRule) SetTags(val *[]interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnSamplingRule_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_xray.CfnSamplingRule",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnSamplingRule_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_xray.CfnSamplingRule",
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
func CfnSamplingRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_xray.CfnSamplingRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSamplingRule_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_xray.CfnSamplingRule",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSamplingRule) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnSamplingRule) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSamplingRule) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnSamplingRule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnSamplingRule) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnSamplingRule) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnSamplingRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnSamplingRule) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSamplingRule) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSamplingRule) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnSamplingRule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSamplingRule) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSamplingRule) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSamplingRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSamplingRule) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A sampling rule that services use to decide whether to instrument a request.
//
// Rule fields can match properties of the service, or properties of a request. The service can ignore rules that don't match its properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   samplingRuleProperty := &samplingRuleProperty{
//   	attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	fixedRate: jsii.Number(123),
//   	host: jsii.String("host"),
//   	httpMethod: jsii.String("httpMethod"),
//   	priority: jsii.Number(123),
//   	reservoirSize: jsii.Number(123),
//   	resourceArn: jsii.String("resourceArn"),
//   	ruleArn: jsii.String("ruleArn"),
//   	ruleName: jsii.String("ruleName"),
//   	serviceName: jsii.String("serviceName"),
//   	serviceType: jsii.String("serviceType"),
//   	urlPath: jsii.String("urlPath"),
//   	version: jsii.Number(123),
//   }
//
type CfnSamplingRule_SamplingRuleProperty struct {
	// Matches attributes derived from the request.
	//
	// *Map Entries:* Maximum number of 5 items.
	//
	// *Key Length Constraints:* Minimum length of 1. Maximum length of 32.
	//
	// *Value Length Constraints:* Minimum length of 1. Maximum length of 32.
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// The percentage of matching requests to instrument, after the reservoir is exhausted.
	FixedRate *float64 `field:"optional" json:"fixedRate" yaml:"fixedRate"`
	// Matches the hostname from a request URL.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Matches the HTTP method of a request.
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// The priority of the sampling rule.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// A fixed number of matching requests to instrument per second, prior to applying the fixed rate.
	//
	// The reservoir is not used directly by services, but applies to all services using the rule collectively.
	ReservoirSize *float64 `field:"optional" json:"reservoirSize" yaml:"reservoirSize"`
	// Matches the ARN of the AWS resource on which the service runs.
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
	// The ARN of the sampling rule.
	//
	// You must specify either RuleARN or RuleName, but not both.
	RuleArn *string `field:"optional" json:"ruleArn" yaml:"ruleArn"`
	// The name of the sampling rule.
	//
	// You must specify either RuleARN or RuleName, but not both.
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// Matches the `name` that the service uses to identify itself in segments.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Matches the `origin` that the service uses to identify its type in segments.
	ServiceType *string `field:"optional" json:"serviceType" yaml:"serviceType"`
	// Matches the path from a request URL.
	UrlPath *string `field:"optional" json:"urlPath" yaml:"urlPath"`
	// The version of the sampling rule format ( `1` ).
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

// A [SamplingRule](https://docs.aws.amazon.com//xray/latest/api/API_SamplingRule.html) and its metadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   samplingRuleRecordProperty := &samplingRuleRecordProperty{
//   	createdAt: jsii.String("createdAt"),
//   	modifiedAt: jsii.String("modifiedAt"),
//   	samplingRule: &samplingRuleProperty{
//   		attributes: map[string]*string{
//   			"attributesKey": jsii.String("attributes"),
//   		},
//   		fixedRate: jsii.Number(123),
//   		host: jsii.String("host"),
//   		httpMethod: jsii.String("httpMethod"),
//   		priority: jsii.Number(123),
//   		reservoirSize: jsii.Number(123),
//   		resourceArn: jsii.String("resourceArn"),
//   		ruleArn: jsii.String("ruleArn"),
//   		ruleName: jsii.String("ruleName"),
//   		serviceName: jsii.String("serviceName"),
//   		serviceType: jsii.String("serviceType"),
//   		urlPath: jsii.String("urlPath"),
//   		version: jsii.Number(123),
//   	},
//   }
//
type CfnSamplingRule_SamplingRuleRecordProperty struct {
	// When the rule was created, in Unix time seconds.
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// When the rule was last modified, in Unix time seconds.
	ModifiedAt *string `field:"optional" json:"modifiedAt" yaml:"modifiedAt"`
	// The sampling rule.
	SamplingRule interface{} `field:"optional" json:"samplingRule" yaml:"samplingRule"`
}

// A document specifying changes to a sampling rule's configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   samplingRuleUpdateProperty := &samplingRuleUpdateProperty{
//   	attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	fixedRate: jsii.Number(123),
//   	host: jsii.String("host"),
//   	httpMethod: jsii.String("httpMethod"),
//   	priority: jsii.Number(123),
//   	reservoirSize: jsii.Number(123),
//   	resourceArn: jsii.String("resourceArn"),
//   	ruleArn: jsii.String("ruleArn"),
//   	ruleName: jsii.String("ruleName"),
//   	serviceName: jsii.String("serviceName"),
//   	serviceType: jsii.String("serviceType"),
//   	urlPath: jsii.String("urlPath"),
//   }
//
type CfnSamplingRule_SamplingRuleUpdateProperty struct {
	// Matches attributes derived from the request.
	//
	// *Map Entries:* Maximum number of 5 items.
	//
	// *Key Length Constraints:* Minimum length of 1. Maximum length of 32.
	//
	// *Value Length Constraints:* Minimum length of 1. Maximum length of 32.
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// The percentage of matching requests to instrument, after the reservoir is exhausted.
	FixedRate *float64 `field:"optional" json:"fixedRate" yaml:"fixedRate"`
	// Matches the hostname from a request URL.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Matches the HTTP method of a request.
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// The priority of the sampling rule.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// A fixed number of matching requests to instrument per second, prior to applying the fixed rate.
	//
	// The reservoir is not used directly by services, but applies to all services using the rule collectively.
	ReservoirSize *float64 `field:"optional" json:"reservoirSize" yaml:"reservoirSize"`
	// Matches the ARN of the AWS resource on which the service runs.
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
	// The ARN of the sampling rule.
	//
	// You must specify either RuleARN or RuleName, but not both.
	RuleArn *string `field:"optional" json:"ruleArn" yaml:"ruleArn"`
	// The name of the sampling rule.
	//
	// You must specify either RuleARN or RuleName, but not both.
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// Matches the `name` that the service uses to identify itself in segments.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Matches the `origin` that the service uses to identify its type in segments.
	ServiceType *string `field:"optional" json:"serviceType" yaml:"serviceType"`
	// Matches the path from a request URL.
	UrlPath *string `field:"optional" json:"urlPath" yaml:"urlPath"`
}

// Properties for defining a `CfnSamplingRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnSamplingRuleProps := &cfnSamplingRuleProps{
//   	ruleName: jsii.String("ruleName"),
//   	samplingRule: &samplingRuleProperty{
//   		attributes: map[string]*string{
//   			"attributesKey": jsii.String("attributes"),
//   		},
//   		fixedRate: jsii.Number(123),
//   		host: jsii.String("host"),
//   		httpMethod: jsii.String("httpMethod"),
//   		priority: jsii.Number(123),
//   		reservoirSize: jsii.Number(123),
//   		resourceArn: jsii.String("resourceArn"),
//   		ruleArn: jsii.String("ruleArn"),
//   		ruleName: jsii.String("ruleName"),
//   		serviceName: jsii.String("serviceName"),
//   		serviceType: jsii.String("serviceType"),
//   		urlPath: jsii.String("urlPath"),
//   		version: jsii.Number(123),
//   	},
//   	samplingRuleRecord: &samplingRuleRecordProperty{
//   		createdAt: jsii.String("createdAt"),
//   		modifiedAt: jsii.String("modifiedAt"),
//   		samplingRule: &samplingRuleProperty{
//   			attributes: map[string]*string{
//   				"attributesKey": jsii.String("attributes"),
//   			},
//   			fixedRate: jsii.Number(123),
//   			host: jsii.String("host"),
//   			httpMethod: jsii.String("httpMethod"),
//   			priority: jsii.Number(123),
//   			reservoirSize: jsii.Number(123),
//   			resourceArn: jsii.String("resourceArn"),
//   			ruleArn: jsii.String("ruleArn"),
//   			ruleName: jsii.String("ruleName"),
//   			serviceName: jsii.String("serviceName"),
//   			serviceType: jsii.String("serviceType"),
//   			urlPath: jsii.String("urlPath"),
//   			version: jsii.Number(123),
//   		},
//   	},
//   	samplingRuleUpdate: &samplingRuleUpdateProperty{
//   		attributes: map[string]*string{
//   			"attributesKey": jsii.String("attributes"),
//   		},
//   		fixedRate: jsii.Number(123),
//   		host: jsii.String("host"),
//   		httpMethod: jsii.String("httpMethod"),
//   		priority: jsii.Number(123),
//   		reservoirSize: jsii.Number(123),
//   		resourceArn: jsii.String("resourceArn"),
//   		ruleArn: jsii.String("ruleArn"),
//   		ruleName: jsii.String("ruleName"),
//   		serviceName: jsii.String("serviceName"),
//   		serviceType: jsii.String("serviceType"),
//   		urlPath: jsii.String("urlPath"),
//   	},
//   	tags: []interface{}{
//   		tags,
//   	},
//   }
//
type CfnSamplingRuleProps struct {
	// The name of the sampling rule.
	//
	// Specify a rule by either name or ARN, but not both. Used only when deleting a sampling rule. When creating or updating a sampling rule, use the `RuleName` or `RuleARN` properties within `SamplingRule` or `SamplingRuleUpdate` .
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// The sampling rule to be created.
	//
	// Must be provided if creating a new sampling rule. Not valid when updating an existing sampling rule.
	SamplingRule interface{} `field:"optional" json:"samplingRule" yaml:"samplingRule"`
	// `AWS::XRay::SamplingRule.SamplingRuleRecord`.
	SamplingRuleRecord interface{} `field:"optional" json:"samplingRuleRecord" yaml:"samplingRuleRecord"`
	// A document specifying changes to a sampling rule's configuration.
	//
	// Must be provided if updating an existing sampling rule. Not valid when creating a new sampling rule.
	//
	// > The `Version` of a sampling rule cannot be updated, and is not part of `SamplingRuleUpdate` .
	SamplingRuleUpdate interface{} `field:"optional" json:"samplingRuleUpdate" yaml:"samplingRuleUpdate"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]interface{} `field:"optional" json:"tags" yaml:"tags"`
}


package awschatbot

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awschatbot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Chatbot::MicrosoftTeamsChannelConfiguration`.
//
// The `AWS::Chatbot::MicrosoftTeamsChannelConfiguration` resource configures a Microsoft Teams channel to allow users to use AWS Chatbot with AWS CloudFormation templates.
//
// This resource requires some setup to be done in the AWS Chatbot console. To provide the required Microsoft Teams team and tenant IDs, you must perform the initial authorization flow with Microsoft Teams in the AWS Chatbot console, then copy and paste the IDs from the console. For more details, see steps 1-4 in [Setting Up AWS Chatbot with Microsoft Teams](https://docs.aws.amazon.com/chatbot/latest/adminguide/teams-setup.html#teams-client-setup) in the *AWS Chatbot Administrator Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMicrosoftTeamsChannelConfiguration := awscdk.Aws_chatbot.NewCfnMicrosoftTeamsChannelConfiguration(this, jsii.String("MyCfnMicrosoftTeamsChannelConfiguration"), &CfnMicrosoftTeamsChannelConfigurationProps{
//   	ConfigurationName: jsii.String("configurationName"),
//   	IamRoleArn: jsii.String("iamRoleArn"),
//   	TeamId: jsii.String("teamId"),
//   	TeamsChannelId: jsii.String("teamsChannelId"),
//   	TeamsTenantId: jsii.String("teamsTenantId"),
//
//   	// the properties below are optional
//   	GuardrailPolicies: []*string{
//   		jsii.String("guardrailPolicies"),
//   	},
//   	LoggingLevel: jsii.String("loggingLevel"),
//   	SnsTopicArns: []*string{
//   		jsii.String("snsTopicArns"),
//   	},
//   	UserRoleRequired: jsii.Boolean(false),
//   })
//
type CfnMicrosoftTeamsChannelConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The name of the configuration.
	ConfigurationName() *string
	SetConfigurationName(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The list of IAM policy ARNs that are applied as channel guardrails.
	//
	// The AWS managed 'AdministratorAccess' policy is applied as a default if this is not set.
	GuardrailPolicies() *[]*string
	SetGuardrailPolicies(val *[]*string)
	// The ARN of the IAM role that defines the permissions for AWS Chatbot .
	//
	// This is a user-defined role that AWS Chatbot will assume. This is not the service-linked role. For more information, see [IAM Policies for AWS Chatbot](https://docs.aws.amazon.com/chatbot/latest/adminguide/chatbot-iam-policies.html) .
	IamRoleArn() *string
	SetIamRoleArn(val *string)
	// Specifies the logging level for this configuration. This property affects the log entries pushed to Amazon CloudWatch Logs.
	//
	// Logging levels include `ERROR` , `INFO` , or `NONE` .
	LoggingLevel() *string
	SetLoggingLevel(val *string)
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
	// The ARNs of the SNS topics that deliver notifications to AWS Chatbot .
	SnsTopicArns() *[]*string
	SetSnsTopicArns(val *[]*string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The ID of the Microsoft Team authorized with AWS Chatbot .
	//
	// To get the team ID, you must perform the initial authorization flow with Microsoft Teams in the AWS Chatbot console. Then you can copy and paste the team ID from the console. For more details, see steps 1-4 in [Get started with Microsoft Teams](https://docs.aws.amazon.com/chatbot/latest/adminguide/teams-setup.html#teams-client-setup) in the *AWS Chatbot Administrator Guide* .
	TeamId() *string
	SetTeamId(val *string)
	// The ID of the Microsoft Teams channel.
	//
	// To get the channel ID, open Microsoft Teams, right click on the channel name in the left pane, then choose Copy. An example of the channel ID syntax is: `19%3ab6ef35dc342d56ba5654e6fc6d25a071%40thread.tacv2` .
	TeamsChannelId() *string
	SetTeamsChannelId(val *string)
	// The ID of the Microsoft Teams tenant.
	//
	// To get the tenant ID, you must perform the initial authorization flow with Microsoft Teams in the AWS Chatbot console. Then you can copy and paste the tenant ID from the console. For more details, see steps 1-4 in [Get started with Microsoft Teams](https://docs.aws.amazon.com/chatbot/latest/adminguide/teams-setup.html#teams-client-setup) in the *AWS Chatbot Administrator Guide* .
	TeamsTenantId() *string
	SetTeamsTenantId(val *string)
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
	// Enables use of a user role requirement in your chat configuration.
	UserRoleRequired() interface{}
	SetUserRoleRequired(val interface{})
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

// The jsii proxy struct for CfnMicrosoftTeamsChannelConfiguration
type jsiiProxy_CfnMicrosoftTeamsChannelConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) ConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) GuardrailPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guardrailPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) IamRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) LoggingLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) SnsTopicArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"snsTopicArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) TeamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) TeamsChannelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamsChannelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) TeamsTenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamsTenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) UserRoleRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userRoleRequired",
		&returns,
	)
	return returns
}


// Create a new `AWS::Chatbot::MicrosoftTeamsChannelConfiguration`.
func NewCfnMicrosoftTeamsChannelConfiguration(scope constructs.Construct, id *string, props *CfnMicrosoftTeamsChannelConfigurationProps) CfnMicrosoftTeamsChannelConfiguration {
	_init_.Initialize()

	if err := validateNewCfnMicrosoftTeamsChannelConfigurationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMicrosoftTeamsChannelConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_chatbot.CfnMicrosoftTeamsChannelConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Chatbot::MicrosoftTeamsChannelConfiguration`.
func NewCfnMicrosoftTeamsChannelConfiguration_Override(c CfnMicrosoftTeamsChannelConfiguration, scope constructs.Construct, id *string, props *CfnMicrosoftTeamsChannelConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_chatbot.CfnMicrosoftTeamsChannelConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration)SetConfigurationName(val *string) {
	if err := j.validateSetConfigurationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationName",
		val,
	)
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration)SetGuardrailPolicies(val *[]*string) {
	_jsii_.Set(
		j,
		"guardrailPolicies",
		val,
	)
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration)SetIamRoleArn(val *string) {
	if err := j.validateSetIamRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration)SetLoggingLevel(val *string) {
	_jsii_.Set(
		j,
		"loggingLevel",
		val,
	)
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration)SetSnsTopicArns(val *[]*string) {
	_jsii_.Set(
		j,
		"snsTopicArns",
		val,
	)
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration)SetTeamId(val *string) {
	if err := j.validateSetTeamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teamId",
		val,
	)
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration)SetTeamsChannelId(val *string) {
	if err := j.validateSetTeamsChannelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teamsChannelId",
		val,
	)
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration)SetTeamsTenantId(val *string) {
	if err := j.validateSetTeamsTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teamsTenantId",
		val,
	)
}

func (j *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration)SetUserRoleRequired(val interface{}) {
	if err := j.validateSetUserRoleRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userRoleRequired",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnMicrosoftTeamsChannelConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMicrosoftTeamsChannelConfiguration_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_chatbot.CfnMicrosoftTeamsChannelConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnMicrosoftTeamsChannelConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnMicrosoftTeamsChannelConfiguration_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_chatbot.CfnMicrosoftTeamsChannelConfiguration",
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
func CfnMicrosoftTeamsChannelConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMicrosoftTeamsChannelConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_chatbot.CfnMicrosoftTeamsChannelConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMicrosoftTeamsChannelConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_chatbot.CfnMicrosoftTeamsChannelConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMicrosoftTeamsChannelConfiguration) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}


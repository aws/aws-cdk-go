package awsbedrock

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies an agent as a resource in a top-level template. Minimally, you must specify the following properties:.
//
// - AgentName – Specify a name for the agent.
// - AgentResourceRoleArn – Specify the Amazon Resource Name (ARN) of the service role with permissions to invoke API operations on the agent. For more information, see [Create a service role for Agents for Amazon Bedrock](https://docs.aws.amazon.com/bedrock/latest/userguide/agents-permissions.html) .
// - FoundationModel – Specify the model ID of a foundation model to use when invoking the agent. For more information, see [Supported regions and models for Agents for Amazon Bedrock](https://docs.aws.amazon.com//bedrock/latest/userguide/agents-supported.html) .
//
// For more information about using agents in Amazon Bedrock , see [Agents for Amazon Bedrock](https://docs.aws.amazon.com/bedrock/latest/userguide/agents.html) .
//
// See the *Properties* section below for descriptions of both the required and optional properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAgent := awscdk.Aws_bedrock.NewCfnAgent(this, jsii.String("MyCfnAgent"), &CfnAgentProps{
//   	AgentName: jsii.String("agentName"),
//
//   	// the properties below are optional
//   	ActionGroups: []interface{}{
//   		&AgentActionGroupProperty{
//   			ActionGroupName: jsii.String("actionGroupName"),
//
//   			// the properties below are optional
//   			ActionGroupExecutor: &ActionGroupExecutorProperty{
//   				CustomControl: jsii.String("customControl"),
//   				Lambda: jsii.String("lambda"),
//   			},
//   			ActionGroupState: jsii.String("actionGroupState"),
//   			ApiSchema: &APISchemaProperty{
//   				Payload: jsii.String("payload"),
//   				S3: &S3IdentifierProperty{
//   					S3BucketName: jsii.String("s3BucketName"),
//   					S3ObjectKey: jsii.String("s3ObjectKey"),
//   				},
//   			},
//   			Description: jsii.String("description"),
//   			FunctionSchema: &FunctionSchemaProperty{
//   				Functions: []interface{}{
//   					&FunctionProperty{
//   						Name: jsii.String("name"),
//
//   						// the properties below are optional
//   						Description: jsii.String("description"),
//   						Parameters: map[string]interface{}{
//   							"parametersKey": &ParameterDetailProperty{
//   								"type": jsii.String("type"),
//
//   								// the properties below are optional
//   								"description": jsii.String("description"),
//   								"required": jsii.Boolean(false),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			ParentActionGroupSignature: jsii.String("parentActionGroupSignature"),
//   			SkipResourceInUseCheckOnDelete: jsii.Boolean(false),
//   		},
//   	},
//   	AgentResourceRoleArn: jsii.String("agentResourceRoleArn"),
//   	AutoPrepare: jsii.Boolean(false),
//   	CustomerEncryptionKeyArn: jsii.String("customerEncryptionKeyArn"),
//   	Description: jsii.String("description"),
//   	FoundationModel: jsii.String("foundationModel"),
//   	GuardrailConfiguration: &GuardrailConfigurationProperty{
//   		GuardrailIdentifier: jsii.String("guardrailIdentifier"),
//   		GuardrailVersion: jsii.String("guardrailVersion"),
//   	},
//   	IdleSessionTtlInSeconds: jsii.Number(123),
//   	Instruction: jsii.String("instruction"),
//   	KnowledgeBases: []interface{}{
//   		&AgentKnowledgeBaseProperty{
//   			Description: jsii.String("description"),
//   			KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//
//   			// the properties below are optional
//   			KnowledgeBaseState: jsii.String("knowledgeBaseState"),
//   		},
//   	},
//   	PromptOverrideConfiguration: &PromptOverrideConfigurationProperty{
//   		PromptConfigurations: []interface{}{
//   			&PromptConfigurationProperty{
//   				BasePromptTemplate: jsii.String("basePromptTemplate"),
//   				InferenceConfiguration: &InferenceConfigurationProperty{
//   					MaximumLength: jsii.Number(123),
//   					StopSequences: []*string{
//   						jsii.String("stopSequences"),
//   					},
//   					Temperature: jsii.Number(123),
//   					TopK: jsii.Number(123),
//   					TopP: jsii.Number(123),
//   				},
//   				ParserMode: jsii.String("parserMode"),
//   				PromptCreationMode: jsii.String("promptCreationMode"),
//   				PromptState: jsii.String("promptState"),
//   				PromptType: jsii.String("promptType"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		OverrideLambda: jsii.String("overrideLambda"),
//   	},
//   	SkipResourceInUseCheckOnDelete: jsii.Boolean(false),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TestAliasTags: map[string]*string{
//   		"testAliasTagsKey": jsii.String("testAliasTags"),
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agent.html
//
type CfnAgent interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggableV2
	// The action groups that belong to an agent.
	ActionGroups() interface{}
	SetActionGroups(val interface{})
	// The name of the agent.
	AgentName() *string
	SetAgentName(val *string)
	// The Amazon Resource Name (ARN) of the IAM role with permissions to invoke API operations on the agent.
	AgentResourceRoleArn() *string
	SetAgentResourceRoleArn(val *string)
	// The Amazon Resource Name (ARN) of the agent.
	AttrAgentArn() *string
	// The unique identifier of the agent.
	AttrAgentId() *string
	// The status of the agent and whether it is ready for use. The following statuses are possible:.
	//
	// - CREATING – The agent is being created.
	// - PREPARING – The agent is being prepared.
	// - PREPARED – The agent is prepared and ready to be invoked.
	// - NOT_PREPARED – The agent has been created but not yet prepared.
	// - FAILED – The agent API operation failed.
	// - UPDATING – The agent is being updated.
	// - DELETING – The agent is being deleted.
	AttrAgentStatus() *string
	// The version of the agent.
	AttrAgentVersion() *string
	// The time at which the agent was created.
	AttrCreatedAt() *string
	// Contains reasons that the agent-related API that you invoked failed.
	AttrFailureReasons() *[]*string
	// The time at which the agent was last prepared.
	AttrPreparedAt() *string
	// Contains recommended actions to take for the agent-related API that you invoked to succeed.
	AttrRecommendedActions() *[]*string
	// The time at which the agent was last updated.
	AttrUpdatedAt() *string
	// Specifies whether to automatically update the `DRAFT` version of the agent after making changes to the agent.
	AutoPrepare() interface{}
	SetAutoPrepare(val interface{})
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The Amazon Resource Name (ARN) of the AWS KMS key that encrypts the agent.
	CustomerEncryptionKeyArn() *string
	SetCustomerEncryptionKeyArn(val *string)
	// The description of the agent.
	Description() *string
	SetDescription(val *string)
	// The foundation model used for orchestration by the agent.
	FoundationModel() *string
	SetFoundationModel(val *string)
	// Details about the guardrail associated with the agent.
	GuardrailConfiguration() interface{}
	SetGuardrailConfiguration(val interface{})
	// The number of seconds for which Amazon Bedrock keeps information about a user's conversation with the agent.
	IdleSessionTtlInSeconds() *float64
	SetIdleSessionTtlInSeconds(val *float64)
	// Instructions that tell the agent what it should do and how it should interact with users.
	Instruction() *string
	SetInstruction(val *string)
	// The knowledge bases associated with the agent.
	KnowledgeBases() interface{}
	SetKnowledgeBases(val interface{})
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
	// Contains configurations to override prompt templates in different parts of an agent sequence.
	PromptOverrideConfiguration() interface{}
	SetPromptOverrideConfiguration(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Specifies whether to delete the resource even if it's in use.
	SkipResourceInUseCheckOnDelete() interface{}
	SetSkipResourceInUseCheckOnDelete(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Metadata that you can assign to a resource as key-value pairs.
	//
	// For more information, see the following resources:.
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	// Metadata that you can assign to a resource as key-value pairs.
	//
	// For more information, see the following resources:.
	TestAliasTags() interface{}
	SetTestAliasTags(val interface{})
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

// The jsii proxy struct for CfnAgent
type jsiiProxy_CfnAgent struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnAgent) ActionGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) AgentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) AgentResourceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentResourceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) AttrAgentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAgentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) AttrAgentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAgentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) AttrAgentStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAgentStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) AttrAgentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAgentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) AttrCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) AttrFailureReasons() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrFailureReasons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) AttrPreparedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPreparedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) AttrRecommendedActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrRecommendedActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) AttrUpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) AutoPrepare() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoPrepare",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) CustomerEncryptionKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerEncryptionKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) FoundationModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"foundationModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) GuardrailConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guardrailConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) IdleSessionTtlInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleSessionTtlInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) Instruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) KnowledgeBases() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"knowledgeBases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) PromptOverrideConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"promptOverrideConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) SkipResourceInUseCheckOnDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipResourceInUseCheckOnDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) TestAliasTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"testAliasTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgent) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnAgent(scope constructs.Construct, id *string, props *CfnAgentProps) CfnAgent {
	_init_.Initialize()

	if err := validateNewCfnAgentParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAgent{}

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrock.CfnAgent",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnAgent_Override(c CfnAgent, scope constructs.Construct, id *string, props *CfnAgentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrock.CfnAgent",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAgent)SetActionGroups(val interface{}) {
	if err := j.validateSetActionGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionGroups",
		val,
	)
}

func (j *jsiiProxy_CfnAgent)SetAgentName(val *string) {
	if err := j.validateSetAgentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentName",
		val,
	)
}

func (j *jsiiProxy_CfnAgent)SetAgentResourceRoleArn(val *string) {
	_jsii_.Set(
		j,
		"agentResourceRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnAgent)SetAutoPrepare(val interface{}) {
	if err := j.validateSetAutoPrepareParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoPrepare",
		val,
	)
}

func (j *jsiiProxy_CfnAgent)SetCustomerEncryptionKeyArn(val *string) {
	_jsii_.Set(
		j,
		"customerEncryptionKeyArn",
		val,
	)
}

func (j *jsiiProxy_CfnAgent)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnAgent)SetFoundationModel(val *string) {
	_jsii_.Set(
		j,
		"foundationModel",
		val,
	)
}

func (j *jsiiProxy_CfnAgent)SetGuardrailConfiguration(val interface{}) {
	if err := j.validateSetGuardrailConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guardrailConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnAgent)SetIdleSessionTtlInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"idleSessionTtlInSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnAgent)SetInstruction(val *string) {
	_jsii_.Set(
		j,
		"instruction",
		val,
	)
}

func (j *jsiiProxy_CfnAgent)SetKnowledgeBases(val interface{}) {
	if err := j.validateSetKnowledgeBasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"knowledgeBases",
		val,
	)
}

func (j *jsiiProxy_CfnAgent)SetPromptOverrideConfiguration(val interface{}) {
	if err := j.validateSetPromptOverrideConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"promptOverrideConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnAgent)SetSkipResourceInUseCheckOnDelete(val interface{}) {
	if err := j.validateSetSkipResourceInUseCheckOnDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipResourceInUseCheckOnDelete",
		val,
	)
}

func (j *jsiiProxy_CfnAgent)SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CfnAgent)SetTestAliasTags(val interface{}) {
	if err := j.validateSetTestAliasTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"testAliasTags",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAgent_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAgent_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrock.CfnAgent",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnAgent_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAgent_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrock.CfnAgent",
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
func CfnAgent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAgent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrock.CfnAgent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAgent_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.CfnAgent",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAgent) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAgent) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAgent) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAgent) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAgent) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAgent) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAgent) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAgent) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAgent) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnAgent) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnAgent) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAgent) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAgent) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAgent) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAgent) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAgent) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnAgent) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnAgent) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAgent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAgent) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}


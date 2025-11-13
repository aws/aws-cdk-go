package awsbedrock

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrock"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a guardrail to detect and filter harmful content in your generative AI application.
//
// Amazon Bedrock Guardrails provides the following safeguards (also known as policies) to detect and filter harmful content:
//
// - *Content filters* - Detect and filter harmful text or image content in input prompts or model responses. Filtering is done based on detection of certain predefined harmful content categories: Hate, Insults, Sexual, Violence, Misconduct and Prompt Attack. You also can adjust the filter strength for each of these categories.
// - *Denied topics* - Define a set of topics that are undesirable in the context of your application. The filter will help block them if detected in user queries or model responses.
// - *Word filters* - Configure filters to help block undesirable words, phrases, and profanity (exact match). Such words can include offensive terms, competitor names, etc.
// - *Sensitive information filters* - Configure filters to help block or mask sensitive information, such as personally identifiable information (PII), or custom regex in user inputs and model responses. Blocking or masking is done based on probabilistic detection of sensitive information in standard formats in entities such as SSN number, Date of Birth, address, etc. This also allows configuring regular expression based detection of patterns for identifiers.
// - *Contextual grounding check* - Help detect and filter hallucinations in model responses based on grounding in a source and relevance to the user query.
//
// For more information, see [How Amazon Bedrock Guardrails works](https://docs.aws.amazon.com/bedrock/latest/userguide/guardrails-how.html) .
//
// Example:
//   import bedrockl1 "github.com/aws/aws-cdk-go/awscdk"
//
//   // Import a guardrail created through the L1 CDK CfnGuardrail construct
//   l1guardrail := bedrockl1.NewCfnGuardrail(this, jsii.String("MyCfnGuardrail"), &CfnGuardrailProps{
//   	BlockedInputMessaging: jsii.String("blockedInputMessaging"),
//   	BlockedOutputsMessaging: jsii.String("blockedOutputsMessaging"),
//   	Name: jsii.String("namemycfnguardrails"),
//   	WordPolicyConfig: &WordPolicyConfigProperty{
//   		WordsConfig: []interface{}{
//   			&WordConfigProperty{
//   				Text: jsii.String("drugs"),
//   			},
//   		},
//   	},
//   })
//
//   importedGuardrail := bedrock.Guardrail_FromCfnGuardrail(l1guardrail)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-guardrail.html
//
type CfnGuardrail interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsbedrock.IGuardrailRef
	awscdk.ITaggableV2
	// The date and time at which the guardrail was created.
	AttrCreatedAt() *string
	// Appears if the `status` of the guardrail is `FAILED` .
	//
	// A list of recommendations to carry out before retrying the request.
	AttrFailureRecommendations() *[]*string
	// The ARN of the guardrail.
	AttrGuardrailArn() *string
	// The unique identifier of the guardrail.
	AttrGuardrailId() *string
	// The status of the guardrail.
	AttrStatus() *string
	// Appears if the `status` is `FAILED` .
	//
	// A list of reasons for why the guardrail failed to be created, updated, versioned, or deleted.
	AttrStatusReasons() *[]*string
	// The date and time at which the guardrail was last updated.
	AttrUpdatedAt() *string
	// The version of the guardrail that was created.
	//
	// This value will always be `DRAFT` .
	AttrVersion() *string
	// Configuration settings for integrating Automated Reasoning policies with Amazon Bedrock Guardrails.
	AutomatedReasoningPolicyConfig() interface{}
	SetAutomatedReasoningPolicyConfig(val interface{})
	// The message to return when the guardrail blocks a prompt.
	BlockedInputMessaging() *string
	SetBlockedInputMessaging(val *string)
	// The message to return when the guardrail blocks a model response.
	BlockedOutputsMessaging() *string
	SetBlockedOutputsMessaging(val *string)
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The content filter policies to configure for the guardrail.
	ContentPolicyConfig() interface{}
	SetContentPolicyConfig(val interface{})
	// Contextual grounding policy config for a guardrail.
	ContextualGroundingPolicyConfig() interface{}
	SetContextualGroundingPolicyConfig(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The system-defined guardrail profile that you're using with your guardrail.
	CrossRegionConfig() interface{}
	SetCrossRegionConfig(val interface{})
	// A description of the guardrail.
	Description() *string
	SetDescription(val *string)
	Env() *interfaces.ResourceEnvironment
	// A reference to a Guardrail resource.
	GuardrailRef() *interfacesawsbedrock.GuardrailReference
	// The ARN of the AWS KMS key that you use to encrypt the guardrail.
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
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
	// The name of the guardrail.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The sensitive information policy to configure for the guardrail.
	SensitiveInformationPolicyConfig() interface{}
	SetSensitiveInformationPolicyConfig(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags that you want to attach to the guardrail.
	Tags() *[]*awscdk.CfnTag
	SetTags(val *[]*awscdk.CfnTag)
	// The topic policies to configure for the guardrail.
	TopicPolicyConfig() interface{}
	SetTopicPolicyConfig(val interface{})
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
	// The word policy you configure for the guardrail.
	WordPolicyConfig() interface{}
	SetWordPolicyConfig(val interface{})
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

// The jsii proxy struct for CfnGuardrail
type jsiiProxy_CfnGuardrail struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsbedrockIGuardrailRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnGuardrail) AttrCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) AttrFailureRecommendations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrFailureRecommendations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) AttrGuardrailArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrGuardrailArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) AttrGuardrailId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrGuardrailId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) AttrStatusReasons() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrStatusReasons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) AttrUpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) AttrVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) AutomatedReasoningPolicyConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automatedReasoningPolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) BlockedInputMessaging() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockedInputMessaging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) BlockedOutputsMessaging() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockedOutputsMessaging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) ContentPolicyConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentPolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) ContextualGroundingPolicyConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contextualGroundingPolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) CrossRegionConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossRegionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) GuardrailRef() *interfacesawsbedrock.GuardrailReference {
	var returns *interfacesawsbedrock.GuardrailReference
	_jsii_.Get(
		j,
		"guardrailRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) SensitiveInformationPolicyConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sensitiveInformationPolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) TopicPolicyConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"topicPolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardrail) WordPolicyConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wordPolicyConfig",
		&returns,
	)
	return returns
}


// Create a new `AWS::Bedrock::Guardrail`.
func NewCfnGuardrail(scope constructs.Construct, id *string, props *CfnGuardrailProps) CfnGuardrail {
	_init_.Initialize()

	if err := validateNewCfnGuardrailParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnGuardrail{}

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrock.CfnGuardrail",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Bedrock::Guardrail`.
func NewCfnGuardrail_Override(c CfnGuardrail, scope constructs.Construct, id *string, props *CfnGuardrailProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrock.CfnGuardrail",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnGuardrail)SetAutomatedReasoningPolicyConfig(val interface{}) {
	if err := j.validateSetAutomatedReasoningPolicyConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automatedReasoningPolicyConfig",
		val,
	)
}

func (j *jsiiProxy_CfnGuardrail)SetBlockedInputMessaging(val *string) {
	if err := j.validateSetBlockedInputMessagingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockedInputMessaging",
		val,
	)
}

func (j *jsiiProxy_CfnGuardrail)SetBlockedOutputsMessaging(val *string) {
	if err := j.validateSetBlockedOutputsMessagingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockedOutputsMessaging",
		val,
	)
}

func (j *jsiiProxy_CfnGuardrail)SetContentPolicyConfig(val interface{}) {
	if err := j.validateSetContentPolicyConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentPolicyConfig",
		val,
	)
}

func (j *jsiiProxy_CfnGuardrail)SetContextualGroundingPolicyConfig(val interface{}) {
	if err := j.validateSetContextualGroundingPolicyConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contextualGroundingPolicyConfig",
		val,
	)
}

func (j *jsiiProxy_CfnGuardrail)SetCrossRegionConfig(val interface{}) {
	if err := j.validateSetCrossRegionConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossRegionConfig",
		val,
	)
}

func (j *jsiiProxy_CfnGuardrail)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnGuardrail)SetKmsKeyArn(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_CfnGuardrail)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnGuardrail)SetSensitiveInformationPolicyConfig(val interface{}) {
	if err := j.validateSetSensitiveInformationPolicyConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sensitiveInformationPolicyConfig",
		val,
	)
}

func (j *jsiiProxy_CfnGuardrail)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CfnGuardrail)SetTopicPolicyConfig(val interface{}) {
	if err := j.validateSetTopicPolicyConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topicPolicyConfig",
		val,
	)
}

func (j *jsiiProxy_CfnGuardrail)SetWordPolicyConfig(val interface{}) {
	if err := j.validateSetWordPolicyConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wordPolicyConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnGuardrail_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGuardrail_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrock.CfnGuardrail",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnGuardrail_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGuardrail_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrock.CfnGuardrail",
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
func CfnGuardrail_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGuardrail_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrock.CfnGuardrail",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGuardrail_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrock.CfnGuardrail",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGuardrail) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnGuardrail) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnGuardrail) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnGuardrail) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnGuardrail) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnGuardrail) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnGuardrail) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnGuardrail) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnGuardrail) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnGuardrail) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnGuardrail) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnGuardrail) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGuardrail) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGuardrail) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnGuardrail) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnGuardrail) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnGuardrail) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnGuardrail) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGuardrail) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGuardrail) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}


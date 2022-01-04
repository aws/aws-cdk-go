package awslex

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslex/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::Lex::Bot`.
//
// Specifies an Amazon Lex conversational bot.
//
// You must configure an intent based on the AMAZON.FallbackIntent built-in intent. If you don't add one, creating the bot will fail.
//
// TODO: EXAMPLE
//
type CfnBot interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrId() *string
	AutoBuildBotLocales() interface{}
	SetAutoBuildBotLocales(val interface{})
	BotFileS3Location() interface{}
	SetBotFileS3Location(val interface{})
	BotLocales() interface{}
	SetBotLocales(val interface{})
	BotTags() interface{}
	SetBotTags(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DataPrivacy() interface{}
	SetDataPrivacy(val interface{})
	Description() *string
	SetDescription(val *string)
	IdleSessionTtlInSeconds() *float64
	SetIdleSessionTtlInSeconds(val *float64)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() awscdk.ConstructNode
	Ref() *string
	RoleArn() *string
	SetRoleArn(val *string)
	Stack() awscdk.Stack
	TestBotAliasTags() interface{}
	SetTestBotAliasTags(val interface{})
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

// The jsii proxy struct for CfnBot
type jsiiProxy_CfnBot struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnBot) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) AutoBuildBotLocales() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoBuildBotLocales",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) BotFileS3Location() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"botFileS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) BotLocales() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"botLocales",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) BotTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"botTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) DataPrivacy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataPrivacy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) IdleSessionTtlInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleSessionTtlInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) TestBotAliasTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"testBotAliasTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lex::Bot`.
func NewCfnBot(scope awscdk.Construct, id *string, props *CfnBotProps) CfnBot {
	_init_.Initialize()

	j := jsiiProxy_CfnBot{}

	_jsii_.Create(
		"monocdk.aws_lex.CfnBot",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lex::Bot`.
func NewCfnBot_Override(c CfnBot, scope awscdk.Construct, id *string, props *CfnBotProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lex.CfnBot",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnBot) SetAutoBuildBotLocales(val interface{}) {
	_jsii_.Set(
		j,
		"autoBuildBotLocales",
		val,
	)
}

func (j *jsiiProxy_CfnBot) SetBotFileS3Location(val interface{}) {
	_jsii_.Set(
		j,
		"botFileS3Location",
		val,
	)
}

func (j *jsiiProxy_CfnBot) SetBotLocales(val interface{}) {
	_jsii_.Set(
		j,
		"botLocales",
		val,
	)
}

func (j *jsiiProxy_CfnBot) SetBotTags(val interface{}) {
	_jsii_.Set(
		j,
		"botTags",
		val,
	)
}

func (j *jsiiProxy_CfnBot) SetDataPrivacy(val interface{}) {
	_jsii_.Set(
		j,
		"dataPrivacy",
		val,
	)
}

func (j *jsiiProxy_CfnBot) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnBot) SetIdleSessionTtlInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"idleSessionTtlInSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnBot) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnBot) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnBot) SetTestBotAliasTags(val interface{}) {
	_jsii_.Set(
		j,
		"testBotAliasTags",
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
func CfnBot_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lex.CfnBot",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnBot_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lex.CfnBot",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnBot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lex.CfnBot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBot_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_lex.CfnBot",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnBot) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnBot) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnBot) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnBot) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnBot) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnBot) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnBot) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnBot) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnBot) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnBot) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnBot) OnPrepare() {
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
func (c *jsiiProxy_CfnBot) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnBot) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnBot) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnBot) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnBot) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnBot) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnBot) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnBot) ToString() *string {
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
func (c *jsiiProxy_CfnBot) Validate() *[]*string {
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
func (c *jsiiProxy_CfnBot) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Provides configuration information for a locale.
//
// TODO: EXAMPLE
//
type CfnBot_BotLocaleProperty struct {
	// The identifier of the language and locale that the bot will be used in.
	//
	// The string must match one of the supported locales.
	LocaleId *string `json:"localeId"`
	// Determines the threshold where Amazon Lex will insert the AMAZON.FallbackIntent, AMAZON.KendraSearchIntent, or both when returning alternative intents. You must configure an AMAZON.FallbackIntent. AMAZON.KendraSearchIntent is only inserted if it is configured for the bot.
	NluConfidenceThreshold *float64 `json:"nluConfidenceThreshold"`
	// A description of the bot locale.
	//
	// Use this to help identify the bot locale in lists.
	Description *string `json:"description"`
	// One or more intents defined for the locale.
	Intents interface{} `json:"intents"`
	// One or more slot types defined for the locale.
	SlotTypes interface{} `json:"slotTypes"`
	// Identifies the Amazon Polly voice used for audio interaction with the user.
	VoiceSettings interface{} `json:"voiceSettings"`
}

// Describes a button to use on a response card used to gather slot values from a user.
//
// TODO: EXAMPLE
//
type CfnBot_ButtonProperty struct {
	// The text that appears on the button.
	//
	// Use this to tell the user the value that is returned when they choose this button.
	Text *string `json:"text"`
	// The value returned to Amazon Lex when the user chooses this button.
	//
	// This must be one of the slot values configured for the slot.
	Value *string `json:"value"`
}

// A custom response string that Amazon Lex sends to your application.
//
// You define the content and structure of the string.
//
// TODO: EXAMPLE
//
type CfnBot_CustomPayloadProperty struct {
	// The string that is sent to your application.
	Value *string `json:"value"`
}

// Specifies whether an intent uses the dialog code hook during conversations with a user.
//
// TODO: EXAMPLE
//
type CfnBot_DialogCodeHookSettingProperty struct {
	// Indicates whether an intent uses the dialog code hook during a conversation with a user.
	Enabled interface{} `json:"enabled"`
}

// TODO: EXAMPLE
//
type CfnBot_ExternalSourceSettingProperty struct {
	// `CfnBot.ExternalSourceSettingProperty.GrammarSlotTypeSetting`.
	GrammarSlotTypeSetting interface{} `json:"grammarSlotTypeSetting"`
}

// Determines if a Lambda function should be invoked for a specific intent.
//
// TODO: EXAMPLE
//
type CfnBot_FulfillmentCodeHookSettingProperty struct {
	// Indicates whether a Lambda function should be invoked for fulfill a specific intent.
	Enabled interface{} `json:"enabled"`
	// Provides settings for update messages sent to the user for long-running Lambda fulfillment functions.
	//
	// Fulfillment updates can be used only with streaming conversations.
	FulfillmentUpdatesSpecification interface{} `json:"fulfillmentUpdatesSpecification"`
	// Provides settings for messages sent to the user for after the Lambda fulfillment function completes.
	//
	// Post-fulfillment messages can be sent for both streaming and non-streaming conversations.
	PostFulfillmentStatusSpecification interface{} `json:"postFulfillmentStatusSpecification"`
}

// Provides settings for a message that is sent to the user when a fulfillment Lambda function starts running.
//
// TODO: EXAMPLE
//
type CfnBot_FulfillmentStartResponseSpecificationProperty struct {
	// The delay between when the Lambda fulfillment function starts running and the start message is played.
	//
	// If the Lambda function returns before the delay is over, the start message isn't played.
	DelayInSeconds *float64 `json:"delayInSeconds"`
	// One to 5 message groups that contain start messages.
	//
	// Amazon Lex chooses one of the messages to play to the user.
	MessageGroups interface{} `json:"messageGroups"`
	// Determines whether the user can interrupt the start message while it is playing.
	AllowInterrupt interface{} `json:"allowInterrupt"`
}

// Provides information for updating the user on the progress of fulfilling an intent.
//
// TODO: EXAMPLE
//
type CfnBot_FulfillmentUpdateResponseSpecificationProperty struct {
	// The frequency that a message is sent to the user.
	//
	// When the period ends, Amazon Lex chooses a message from the message groups and plays it to the user. If the fulfillment Lambda function returns before the first period ends, an update message is not played to the user.
	FrequencyInSeconds *float64 `json:"frequencyInSeconds"`
	// One to 5 message groups that contain update messages.
	//
	// Amazon Lex chooses one of the messages to play to the user.
	MessageGroups interface{} `json:"messageGroups"`
	// Determines whether the user can interrupt an update message while it is playing.
	AllowInterrupt interface{} `json:"allowInterrupt"`
}

// Provides information for updating the user on the progress of fulfilling an intent.
//
// TODO: EXAMPLE
//
type CfnBot_FulfillmentUpdatesSpecificationProperty struct {
	// Determines whether fulfillment updates are sent to the user. When this field is true, updates are sent.
	//
	// If the active field is set to true, the `startResponse` , `updateResponse` , and `timeoutInSeconds` fields are required.
	Active interface{} `json:"active"`
	// Provides configuration information for the message sent to users when the fulfillment Lambda functions starts running.
	StartResponse interface{} `json:"startResponse"`
	// The length of time that the fulfillment Lambda function should run before it times out.
	TimeoutInSeconds *float64 `json:"timeoutInSeconds"`
	// Provides configuration information for messages sent periodically to the user while the fulfillment Lambda function is running.
	UpdateResponse interface{} `json:"updateResponse"`
}

// TODO: EXAMPLE
//
type CfnBot_GrammarSlotTypeSettingProperty struct {
	// `CfnBot.GrammarSlotTypeSettingProperty.Source`.
	Source interface{} `json:"source"`
}

// TODO: EXAMPLE
//
type CfnBot_GrammarSlotTypeSourceProperty struct {
	// `CfnBot.GrammarSlotTypeSourceProperty.S3BucketName`.
	S3BucketName *string `json:"s3BucketName"`
	// `CfnBot.GrammarSlotTypeSourceProperty.S3ObjectKey`.
	S3ObjectKey *string `json:"s3ObjectKey"`
	// `CfnBot.GrammarSlotTypeSourceProperty.KmsKeyArn`.
	KmsKeyArn *string `json:"kmsKeyArn"`
}

// A card that is shown to the user by a messaging platform.
//
// You define the contents of the card, the card is displayed by the platform.
//
// When you use a response card, the response from the user is constrained to the text associated with a button on the card.
//
// TODO: EXAMPLE
//
type CfnBot_ImageResponseCardProperty struct {
	// The title to display on the response card.
	//
	// The format of the title is determined by the platform displaying the response card.
	Title *string `json:"title"`
	// A list of buttons that should be displayed on the response card.
	//
	// The arrangement of the buttons is determined by the platform that displays the buttons.
	Buttons interface{} `json:"buttons"`
	// The URL of an image to display on the response card.
	//
	// The image URL must be publicly available so that the platform displaying the response card has access to the image.
	ImageUrl *string `json:"imageUrl"`
	// The subtitle to display on the response card.
	//
	// The format of the subtitle is determined by the platform displaying the response card.
	Subtitle *string `json:"subtitle"`
}

// The name of a context that must be active for an intent to be selected by Amazon Lex .
//
// TODO: EXAMPLE
//
type CfnBot_InputContextProperty struct {
	// The name of the context.
	Name *string `json:"name"`
}

// Provides a statement the Amazon Lex conveys to the user when the intent is successfully fulfilled.
//
// TODO: EXAMPLE
//
type CfnBot_IntentClosingSettingProperty struct {
	// The response that Amazon Lex sends to the user when the intent is complete.
	ClosingResponse interface{} `json:"closingResponse"`
	// Specifies whether an intent's closing response is used.
	//
	// When this field is false, the closing response isn't sent to the user and no closing input from the user is used. If the IsActive field isn't specified, the default is true.
	IsActive interface{} `json:"isActive"`
}

// Provides a prompt for making sure that the user is ready for the intent to be fulfilled.
//
// TODO: EXAMPLE
//
type CfnBot_IntentConfirmationSettingProperty struct {
	// When the user answers "no" to the question defined in PromptSpecification, Amazon Lex responds with this response to acknowledge that the intent was canceled.
	DeclinationResponse interface{} `json:"declinationResponse"`
	// Prompts the user to confirm the intent.
	//
	// This question should have a yes or no answer.
	PromptSpecification interface{} `json:"promptSpecification"`
	// Specifies whether the intent's confirmation is sent to the user.
	//
	// When this field is false, confirmation and declination responses aren't sent and processing continues as if the responses aren't present. If the active field isn't specified, the default is true.
	IsActive interface{} `json:"isActive"`
}

// Represents an action that the user wants to perform.
//
// TODO: EXAMPLE
//
type CfnBot_IntentProperty struct {
	// The name of the intent.
	//
	// Intent names must be unique within the locale that contains the intent and can't match the name of any built-in intent.
	Name *string `json:"name"`
	// A description of the intent.
	//
	// Use the description to help identify the intent in lists.
	Description *string `json:"description"`
	// Specifies that Amazon Lex invokes the alias Lambda function for each user input.
	//
	// You can invoke this Lambda function to personalize user interaction.
	DialogCodeHook interface{} `json:"dialogCodeHook"`
	// Specifies that Amazon Lex invokes the alias Lambda function when the intent is ready for fulfillment.
	//
	// You can invoke this function to complete the bot's transaction with the user.
	FulfillmentCodeHook interface{} `json:"fulfillmentCodeHook"`
	// A list of contexts that must be active for this intent to be considered by Amazon Lex .
	InputContexts interface{} `json:"inputContexts"`
	// Sets the response that Amazon Lex sends to the user when the intent is closed.
	IntentClosingSetting interface{} `json:"intentClosingSetting"`
	// Provides prompts that Amazon Lex sends to the user to confirm the completion of an intent.
	//
	// If the user answers "no," the settings contain a statement that is sent to the user to end the intent.
	IntentConfirmationSetting interface{} `json:"intentConfirmationSetting"`
	// Configuration information required to use the AMAZON.KendraSearchIntent intent to connect to an Amazon Kendra index. The AMAZON.KendraSearchIntent intent is called with Amazon Lex can't determine another intent to invoke.
	KendraConfiguration interface{} `json:"kendraConfiguration"`
	// A list of contexts that the intent activates when it is fulfilled.
	OutputContexts interface{} `json:"outputContexts"`
	// A unique identifier for the built-in intent to base this intent on.
	ParentIntentSignature *string `json:"parentIntentSignature"`
	// A list of utterances that a user might say to signal the intent.
	SampleUtterances interface{} `json:"sampleUtterances"`
	// Indicates the priority for slots.
	//
	// Amazon Lex prompts the user for slot values in priority order.
	SlotPriorities interface{} `json:"slotPriorities"`
	// A list of slots that the intent requires for fulfillment.
	Slots interface{} `json:"slots"`
}

// Provides configuration information for the AMAZON.KendraSearchIntent intent. When you use this intent, Amazon Lex searches the specified Amazon Kendra index and returns documents from the index that match the user's utterance.
//
// TODO: EXAMPLE
//
type CfnBot_KendraConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon Kendra index that you want the AMAZON.KendraSearchIntent intent to search. The index must be in the same account and Region as the Amazon Lex bot.
	KendraIndex *string `json:"kendraIndex"`
	// A query filter that Amazon Lex sends to Amazon Kendra to filter the response from a query.
	//
	// The filter is in the format defined by Amazon Kendra.
	QueryFilterString *string `json:"queryFilterString"`
	// Determines whether the AMAZON.KendraSearchIntent intent uses a custom query string to query the Amazon Kendra index.
	QueryFilterStringEnabled interface{} `json:"queryFilterStringEnabled"`
}

// Provides one or more messages that Amazon Lex should send to the user.
//
// TODO: EXAMPLE
//
type CfnBot_MessageGroupProperty struct {
	// The primary message that Amazon Lex should send to the user.
	Message interface{} `json:"message"`
	// Message variations to send to the user.
	//
	// When variations are defined, Amazon Lex chooses the primary message or one of the variations to send to the user.
	Variations interface{} `json:"variations"`
}

// The object that provides message text and it's type.
//
// TODO: EXAMPLE
//
type CfnBot_MessageProperty struct {
	// A message in a custom format defined by the client application.
	CustomPayload interface{} `json:"customPayload"`
	// A message that defines a response card that the client application can show to the user.
	ImageResponseCard interface{} `json:"imageResponseCard"`
	// A message in plain text format.
	PlainTextMessage interface{} `json:"plainTextMessage"`
	// A message in Speech Synthesis Markup Language (SSML) format.
	SsmlMessage interface{} `json:"ssmlMessage"`
}

// Indicates whether a slot can return multiple values.
//
// TODO: EXAMPLE
//
type CfnBot_MultipleValuesSettingProperty struct {
	// Indicates whether a slot can return multiple values.
	//
	// When true, the slot may return more than one value in a response. When false, the slot returns only a single value. If AllowMultipleValues is not set, the default value is false.
	//
	// Multi-value slots are only available in the en-US locale.
	AllowMultipleValues interface{} `json:"allowMultipleValues"`
}

// Determines whether Amazon Lex obscures slot values in conversation logs.
//
// TODO: EXAMPLE
//
type CfnBot_ObfuscationSettingProperty struct {
	// Value that determines whether Amazon Lex obscures slot values in conversation logs.
	//
	// The default is to obscure the values.
	ObfuscationSettingType *string `json:"obfuscationSettingType"`
}

// Describes a session context that is activated when an intent is fulfilled.
//
// TODO: EXAMPLE
//
type CfnBot_OutputContextProperty struct {
	// The name of the output context.
	Name *string `json:"name"`
	// The amount of time, in seconds, that the output context should remain active.
	//
	// The time is figured from the first time the context is sent to the user.
	TimeToLiveInSeconds *float64 `json:"timeToLiveInSeconds"`
	// The number of conversation turns that the output context should remain active.
	//
	// The number of turns is counted from the first time that the context is sent to the user.
	TurnsToLive *float64 `json:"turnsToLive"`
}

// Defines an ASCII text message to send to the user.
//
// TODO: EXAMPLE
//
type CfnBot_PlainTextMessageProperty struct {
	// The message to send to the user.
	Value *string `json:"value"`
}

// Provides a setting that determines whether the post-fulfillment response is sent to the user.
//
// For more information, see [Post-fulfillment response](https://docs.aws.amazon.com/latest/dg/streaming-progress.html#progress-complete) in the *Amazon Lex developer guide* .
//
// TODO: EXAMPLE
//
type CfnBot_PostFulfillmentStatusSpecificationProperty struct {
	// Specifies a list of message groups that Amazon Lex uses to respond when fulfillment isn't successful.
	FailureResponse interface{} `json:"failureResponse"`
	// Specifies a list of message groups that Amazon Lex uses to respond when the fulfillment is successful.
	SuccessResponse interface{} `json:"successResponse"`
	// Specifies a list of message groups that Amazon Lex uses to respond when fulfillment isn't completed within the timeout period.
	TimeoutResponse interface{} `json:"timeoutResponse"`
}

// Specifies a list of message groups that Amazon Lex sends to a user to elicit a response.
//
// TODO: EXAMPLE
//
type CfnBot_PromptSpecificationProperty struct {
	// The maximum number of times the bot tries to elicit a response from the user using this prompt.
	MaxRetries *float64 `json:"maxRetries"`
	// A collection of responses that Amazon Lex can send to the user.
	//
	// Amazon Lex chooses the actual response to send at runtime.
	MessageGroupsList interface{} `json:"messageGroupsList"`
	// Indicates whether the user can interrupt a speech prompt from the bot.
	AllowInterrupt interface{} `json:"allowInterrupt"`
}

// Specifies a list of message groups that Amazon Lex uses to respond to user input.
//
// TODO: EXAMPLE
//
type CfnBot_ResponseSpecificationProperty struct {
	// A collection of responses that Amazon Lex can send to the user.
	//
	// Amazon Lex chooses the actual response to send at runtime.
	MessageGroupsList interface{} `json:"messageGroupsList"`
	// Indicates whether the user can interrupt a speech response from Amazon Lex .
	AllowInterrupt interface{} `json:"allowInterrupt"`
}

// Defines an Amazon S3 bucket location.
//
// TODO: EXAMPLE
//
type CfnBot_S3LocationProperty struct {
	// The S3 bucket name.
	S3Bucket *string `json:"s3Bucket"`
	// The path and file name to the object in the S3 bucket.
	S3ObjectKey *string `json:"s3ObjectKey"`
	// The version of the object in the S3 bucket.
	S3ObjectVersion *string `json:"s3ObjectVersion"`
}

// Defines a Speech Synthesis Markup Language (SSML) prompt.
//
// TODO: EXAMPLE
//
type CfnBot_SSMLMessageProperty struct {
	// The SSML text that defines the prompt.
	Value *string `json:"value"`
}

// A sample utterance that invokes and intent or responds to a slot elicitation prompt.
//
// TODO: EXAMPLE
//
type CfnBot_SampleUtteranceProperty struct {
	// The sample utterance that Amazon Lex uses to build its machine-learning model to recognize intents.
	Utterance *string `json:"utterance"`
}

// Defines one of the values for a slot type.
//
// TODO: EXAMPLE
//
type CfnBot_SampleValueProperty struct {
	// The value that can be used for a slot type.
	Value *string `json:"value"`
}

// Specifies the default value to use when a user doesn't provide a value for a slot.
//
// TODO: EXAMPLE
//
type CfnBot_SlotDefaultValueProperty struct {
	// The default value to use when a user doesn't provide a value for a slot.
	DefaultValue *string `json:"defaultValue"`
}

// Defines a list of values that Amazon Lex should use as the default value for a slot.
//
// TODO: EXAMPLE
//
type CfnBot_SlotDefaultValueSpecificationProperty struct {
	// A list of default values.
	//
	// Amazon Lex chooses the default value to use in the order that they are presented in the list.
	DefaultValueList interface{} `json:"defaultValueList"`
}

// Sets the priority that Amazon Lex should use when eliciting slots values from a user.
//
// TODO: EXAMPLE
//
type CfnBot_SlotPriorityProperty struct {
	// The priority that Amazon Lex should apply to the slot.
	Priority *float64 `json:"priority"`
	// The name of the slot.
	SlotName *string `json:"slotName"`
}

// Specifies the definition of a slot.
//
// Amazon Lex elicits slot values from uses to fulfill the user's intent.
//
// TODO: EXAMPLE
//
type CfnBot_SlotProperty struct {
	// The name of the slot.
	Name *string `json:"name"`
	// The name of the slot type that this slot is based on.
	//
	// The slot type defines the acceptable values for the slot.
	SlotTypeName *string `json:"slotTypeName"`
	// Determines the slot resolution strategy that Amazon Lex uses to return slot type values.
	//
	// The field can be set to one of the following values:
	//
	// - OriginalValue - Returns the value entered by the user, if the user value is similar to a slot value.
	// - TopResolution - If there is a resolution list for the slot, return the first value in the resolution list as the slot type value. If there is no resolution list, null is returned.
	//
	// If you don't specify the valueSelectionStrategy, the default is OriginalValue.
	ValueElicitationSetting interface{} `json:"valueElicitationSetting"`
	// A description of the slot type.
	Description *string `json:"description"`
	// Determines whether the slot can return multiple values to the application.
	MultipleValuesSetting interface{} `json:"multipleValuesSetting"`
	// Determines whether the contents of the slot are obfuscated in Amazon CloudWatch Logs logs.
	//
	// Use obfuscated slots to protect information such as personally identifiable information (PII) in logs.
	ObfuscationSetting interface{} `json:"obfuscationSetting"`
}

// Describes a slot type.
//
// TODO: EXAMPLE
//
type CfnBot_SlotTypeProperty struct {
	// The name of the slot type.
	//
	// A slot type name must be unique withing the account.
	Name *string `json:"name"`
	// A description of the slot type.
	//
	// Use the description to help identify the slot type in lists.
	Description *string `json:"description"`
	// `CfnBot.SlotTypeProperty.ExternalSourceSetting`.
	ExternalSourceSetting interface{} `json:"externalSourceSetting"`
	// The built-in slot type used as a parent of this slot type.
	//
	// When you define a parent slot type, the new slot type has the configuration of the parent lot type.
	//
	// Only AMAZON.AlphaNumeric is supported.
	ParentSlotTypeSignature *string `json:"parentSlotTypeSignature"`
	// A list of SlotTypeValue objects that defines the values that the slot type can take.
	//
	// Each value can have a list of synonyms, additional values that help train the machine learning model about the values that it resolves for the slot.
	SlotTypeValues interface{} `json:"slotTypeValues"`
	// Determines the slot resolution strategy that Amazon Lex uses to return slot type values.
	//
	// The field can be set to one of the following values:
	//
	// - OriginalValue - Returns the value entered by the user, if the user value is similar to a slot value.
	// - TopResolution - If there is a resolution list for the slot, return the first value in the resolution list as the slot type value. If there is no resolution list, null is returned.
	//
	// If you don't specify the valueSelectionStrategy, the default is OriginalValue.
	ValueSelectionSetting interface{} `json:"valueSelectionSetting"`
}

// Each slot type can have a set of values.
//
// The `SlotTypeValue` represents a value that the slot type can take.
//
// TODO: EXAMPLE
//
type CfnBot_SlotTypeValueProperty struct {
	// The value of the slot type entry.
	SampleValue interface{} `json:"sampleValue"`
	// Additional values related to the slot type entry.
	Synonyms interface{} `json:"synonyms"`
}

// Settings that you can use for eliciting a slot value.
//
// TODO: EXAMPLE
//
type CfnBot_SlotValueElicitationSettingProperty struct {
	// Specifies whether the slot is required or optional.
	SlotConstraint *string `json:"slotConstraint"`
	// A list of default values for a slot.
	//
	// Default values are used when Amazon Lex hasn't determined a value for a slot. You can specify default values from context variables, session attributes, and defined values.
	DefaultValueSpecification interface{} `json:"defaultValueSpecification"`
	// The prompt that Amazon Lex uses to elicit the slot value from the user.
	PromptSpecification interface{} `json:"promptSpecification"`
	// If you know a specific pattern that users might respond to an Amazon Lex request for a slot value, you can provide those utterances to improve accuracy.
	//
	// This is optional. In most cases Amazon Lex is capable of understanding user utterances.
	SampleUtterances interface{} `json:"sampleUtterances"`
	// Specifies the prompts that Amazon Lex uses while a bot is waiting for customer input.
	WaitAndContinueSpecification interface{} `json:"waitAndContinueSpecification"`
}

// Provides a regular expression used to validate the value of a slot.
//
// TODO: EXAMPLE
//
type CfnBot_SlotValueRegexFilterProperty struct {
	// A regular expression used to validate the value of a slot.
	//
	// Use a standard regular expression. Amazon Lex supports the following characters in the regular expression:
	//
	// - A-Z, a-z
	// - 0-9
	// - Unicode characters ("\ u<Unicode>")
	//
	// Represent Unicode characters with four digits, for example "]u0041" or "\ u005A".
	//
	// The following regular expression operators are not supported:
	//
	// - Infinite repeaters: *, +, or {x,} with no upper bound
	// - Wild card (.)
	Pattern *string `json:"pattern"`
}

// Contains settings used by Amazon Lex to select a slot value.
//
// TODO: EXAMPLE
//
type CfnBot_SlotValueSelectionSettingProperty struct {
	// Determines the slot resolution strategy that Amazon Lex uses to return slot type values.
	//
	// The field can be set to one of the following values:
	//
	// - OriginalValue - Returns the value entered by the user, if the user value is similar to a slot value.
	// - TopResolution - If there is a resolution list for the slot, return the first value in the resolution list as the slot type value. If there is no resolution list, null is returned.
	//
	// If you don't specify the valueSelectionStrategy, the default is OriginalValue.
	ResolutionStrategy *string `json:"resolutionStrategy"`
	// A regular expression used to validate the value of a slot.
	RegexFilter interface{} `json:"regexFilter"`
}

// Defines the messages that Amazon Lex sends to a user to remind them that the bot is waiting for a response.
//
// TODO: EXAMPLE
//
type CfnBot_StillWaitingResponseSpecificationProperty struct {
	// How often a message should be sent to the user.
	//
	// Minimum of 1 second, maximum of 5 minutes.
	FrequencyInSeconds *float64 `json:"frequencyInSeconds"`
	// A collection of responses that Amazon Lex can send to the user.
	//
	// Amazon Lex chooses the actual response to send at runtime.
	MessageGroupsList interface{} `json:"messageGroupsList"`
	// If Amazon Lex waits longer than this length of time for a response, it will stop sending messages.
	TimeoutInSeconds *float64 `json:"timeoutInSeconds"`
	// Indicates that the user can interrupt the response by speaking while the message is being played.
	AllowInterrupt interface{} `json:"allowInterrupt"`
}

// Identifies the Amazon Polly voice used for audio interaction with the user.
//
// TODO: EXAMPLE
//
type CfnBot_VoiceSettingsProperty struct {
	// The Amazon Polly voice used for voice interaction with the user.
	VoiceId *string `json:"voiceId"`
}

// Specifies the prompts that Amazon Lex uses while a bot is waiting for customer input.
//
// TODO: EXAMPLE
//
type CfnBot_WaitAndContinueSpecificationProperty struct {
	// The response that Amazon Lex sends to indicate that the bot is ready to continue the conversation.
	ContinueResponse interface{} `json:"continueResponse"`
	// The response that Amazon Lex sends to indicate that the bot is waiting for the conversation to continue.
	WaitingResponse interface{} `json:"waitingResponse"`
	// Specifies whether the bot will wait for a user to respond.
	//
	// When this field is false, wait and continue responses for a slot aren't used and the bot expects an appropriate response within the configured timeout. If the IsActive field isn't specified, the default is true.
	IsActive interface{} `json:"isActive"`
	// A response that Amazon Lex sends periodically to the user to indicate that the bot is still waiting for input from the user.
	StillWaitingResponse interface{} `json:"stillWaitingResponse"`
}

// A CloudFormation `AWS::Lex::BotAlias`.
//
// Specifies an alias for the specified version of a bot. Use an alias to enable you to change the version of a bot without updating applications that use the bot.
//
// For example, you can specify an alias called "PROD" that your applications use to call the Amazon Lex bot.
//
// TODO: EXAMPLE
//
type CfnBotAlias interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrBotAliasId() *string
	AttrBotAliasStatus() *string
	BotAliasLocaleSettings() interface{}
	SetBotAliasLocaleSettings(val interface{})
	BotAliasName() *string
	SetBotAliasName(val *string)
	BotAliasTags() interface{}
	SetBotAliasTags(val interface{})
	BotId() *string
	SetBotId(val *string)
	BotVersion() *string
	SetBotVersion(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ConversationLogSettings() interface{}
	SetConversationLogSettings(val interface{})
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	LogicalId() *string
	Node() awscdk.ConstructNode
	Ref() *string
	SentimentAnalysisSettings() interface{}
	SetSentimentAnalysisSettings(val interface{})
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

// The jsii proxy struct for CfnBotAlias
type jsiiProxy_CfnBotAlias struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnBotAlias) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) AttrBotAliasId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrBotAliasId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) AttrBotAliasStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrBotAliasStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) BotAliasLocaleSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"botAliasLocaleSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) BotAliasName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botAliasName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) BotAliasTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"botAliasTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) BotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) BotVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) ConversationLogSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conversationLogSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) SentimentAnalysisSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sentimentAnalysisSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lex::BotAlias`.
func NewCfnBotAlias(scope awscdk.Construct, id *string, props *CfnBotAliasProps) CfnBotAlias {
	_init_.Initialize()

	j := jsiiProxy_CfnBotAlias{}

	_jsii_.Create(
		"monocdk.aws_lex.CfnBotAlias",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lex::BotAlias`.
func NewCfnBotAlias_Override(c CfnBotAlias, scope awscdk.Construct, id *string, props *CfnBotAliasProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lex.CfnBotAlias",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnBotAlias) SetBotAliasLocaleSettings(val interface{}) {
	_jsii_.Set(
		j,
		"botAliasLocaleSettings",
		val,
	)
}

func (j *jsiiProxy_CfnBotAlias) SetBotAliasName(val *string) {
	_jsii_.Set(
		j,
		"botAliasName",
		val,
	)
}

func (j *jsiiProxy_CfnBotAlias) SetBotAliasTags(val interface{}) {
	_jsii_.Set(
		j,
		"botAliasTags",
		val,
	)
}

func (j *jsiiProxy_CfnBotAlias) SetBotId(val *string) {
	_jsii_.Set(
		j,
		"botId",
		val,
	)
}

func (j *jsiiProxy_CfnBotAlias) SetBotVersion(val *string) {
	_jsii_.Set(
		j,
		"botVersion",
		val,
	)
}

func (j *jsiiProxy_CfnBotAlias) SetConversationLogSettings(val interface{}) {
	_jsii_.Set(
		j,
		"conversationLogSettings",
		val,
	)
}

func (j *jsiiProxy_CfnBotAlias) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnBotAlias) SetSentimentAnalysisSettings(val interface{}) {
	_jsii_.Set(
		j,
		"sentimentAnalysisSettings",
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
func CfnBotAlias_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lex.CfnBotAlias",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnBotAlias_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lex.CfnBotAlias",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnBotAlias_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lex.CfnBotAlias",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBotAlias_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_lex.CfnBotAlias",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnBotAlias) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnBotAlias) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnBotAlias) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnBotAlias) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnBotAlias) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnBotAlias) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnBotAlias) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnBotAlias) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnBotAlias) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnBotAlias) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnBotAlias) OnPrepare() {
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
func (c *jsiiProxy_CfnBotAlias) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnBotAlias) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnBotAlias) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnBotAlias) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnBotAlias) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnBotAlias) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnBotAlias) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnBotAlias) ToString() *string {
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
func (c *jsiiProxy_CfnBotAlias) Validate() *[]*string {
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
func (c *jsiiProxy_CfnBotAlias) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies the S3 bucket location where audio logs are stored.
//
// TODO: EXAMPLE
//
type CfnBotAlias_AudioLogDestinationProperty struct {
	// The S3 bucket location where audio logs are stored.
	S3Bucket interface{} `json:"s3Bucket"`
}

// Settings for logging audio of conversations between Amazon Lex and a user.
//
// You specify whether to log audio and the Amazon S3 bucket where the audio file is stored.
//
// TODO: EXAMPLE
//
type CfnBotAlias_AudioLogSettingProperty struct {
	// The location of audio log files collected when conversation logging is enabled for a bot.
	Destination interface{} `json:"destination"`
	// Determines whether audio logging in enabled for the bot.
	Enabled interface{} `json:"enabled"`
}

// Specifies settings that are unique to a locale.
//
// For example, you can use different Lambda function depending on the bot's locale.
//
// TODO: EXAMPLE
//
type CfnBotAlias_BotAliasLocaleSettingsItemProperty struct {
	// Specifies settings that are unique to a locale.
	BotAliasLocaleSetting interface{} `json:"botAliasLocaleSetting"`
	// The unique identifier of the locale.
	LocaleId *string `json:"localeId"`
}

// Specifies settings that are unique to a locale.
//
// For example, you can use different Lambda function depending on the bot's locale.
//
// TODO: EXAMPLE
//
type CfnBotAlias_BotAliasLocaleSettingsProperty struct {
	// Determines whether the locale is enabled for the bot.
	//
	// If the value is false, the locale isn't available for use.
	Enabled interface{} `json:"enabled"`
	// Specifies the Lambda function that should be used in the locale.
	CodeHookSpecification interface{} `json:"codeHookSpecification"`
}

// The Amazon CloudWatch Logs log group where the text and metadata logs are delivered.
//
// The log group must exist before you enable logging.
//
// TODO: EXAMPLE
//
type CfnBotAlias_CloudWatchLogGroupLogDestinationProperty struct {
	// The Amazon Resource Name (ARN) of the log group where text and metadata logs are delivered.
	CloudWatchLogGroupArn *string `json:"cloudWatchLogGroupArn"`
	// The prefix of the log stream name within the log group that you specified.
	LogPrefix *string `json:"logPrefix"`
}

// Contains information about code hooks that Amazon Lex calls during a conversation.
//
// TODO: EXAMPLE
//
type CfnBotAlias_CodeHookSpecificationProperty struct {
	// Specifies a Lambda function that verifies requests to a bot or fulfills the user's request to a bot.
	LambdaCodeHook interface{} `json:"lambdaCodeHook"`
}

// Configures conversation logging that saves audio, text, and metadata for the conversations with your users.
//
// TODO: EXAMPLE
//
type CfnBotAlias_ConversationLogSettingsProperty struct {
	// The Amazon S3 settings for logging audio to an S3 bucket.
	AudioLogSettings interface{} `json:"audioLogSettings"`
	// The Amazon CloudWatch Logs settings for logging text and metadata.
	TextLogSettings interface{} `json:"textLogSettings"`
}

// Specifies a Lambda function that verifies requests to a bot or fulfills the user's request to a bot.
//
// TODO: EXAMPLE
//
type CfnBotAlias_LambdaCodeHookProperty struct {
	// The version of the request-response that you want Amazon Lex to use to invoke your Lambda function.
	CodeHookInterfaceVersion *string `json:"codeHookInterfaceVersion"`
	// The Amazon Resource Name (ARN) of the Lambda function.
	LambdaArn *string `json:"lambdaArn"`
}

// Specifies an Amazon S3 bucket for logging audio conversations.
//
// TODO: EXAMPLE
//
type CfnBotAlias_S3BucketLogDestinationProperty struct {
	// The S3 prefix to assign to audio log files.
	LogPrefix *string `json:"logPrefix"`
	// The Amazon Resource Name (ARN) of an Amazon S3 bucket where audio log files are stored.
	S3BucketArn *string `json:"s3BucketArn"`
	// The Amazon Resource Name (ARN) of an AWS Key Management Service key for encrypting audio log files stored in an S3 bucket.
	KmsKeyArn *string `json:"kmsKeyArn"`
}

// Defines the Amazon CloudWatch Logs destination log group for conversation text logs.
//
// TODO: EXAMPLE
//
type CfnBotAlias_TextLogDestinationProperty struct {
}

// Defines settings to enable conversation logs.
//
// TODO: EXAMPLE
//
type CfnBotAlias_TextLogSettingProperty struct {
	// Defines the Amazon CloudWatch Logs destination log group for conversation text logs.
	Destination interface{} `json:"destination"`
	// Determines whether conversation logs should be stored for an alias.
	Enabled interface{} `json:"enabled"`
}

// Properties for defining a `CfnBotAlias`.
//
// TODO: EXAMPLE
//
type CfnBotAliasProps struct {
	// The name of the bot alias.
	BotAliasName *string `json:"botAliasName"`
	// The unique identifier of the bot.
	BotId *string `json:"botId"`
	// Maps configuration information to a specific locale.
	//
	// You can use this parameter to specify a specific Lambda function to run different functions in different locales.
	BotAliasLocaleSettings interface{} `json:"botAliasLocaleSettings"`
	// An array of key-value pairs to apply to this resource.
	//
	// You can only add tags when you specify an alias.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	BotAliasTags interface{} `json:"botAliasTags"`
	// The version of the bot that the bot alias references.
	BotVersion *string `json:"botVersion"`
	// Specifies whether Amazon Lex logs text and audio for conversations with the bot.
	//
	// When you enable conversation logs, text logs store text input, transcripts of audio input, and associated metadata in Amazon CloudWatch logs. Audio logs store input in Amazon S3 .
	ConversationLogSettings interface{} `json:"conversationLogSettings"`
	// The description of the bot alias.
	Description *string `json:"description"`
	// Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment of user utterances.
	SentimentAnalysisSettings interface{} `json:"sentimentAnalysisSettings"`
}

// Properties for defining a `CfnBot`.
//
// TODO: EXAMPLE
//
type CfnBotProps struct {
	// Provides information on additional privacy protections Amazon Lex should use with the bot's data.
	DataPrivacy interface{} `json:"dataPrivacy"`
	// The time, in seconds, that Amazon Lex should keep information about a user's conversation with the bot.
	//
	// A user interaction remains active for the amount of time specified. If no conversation occurs during this time, the session expires and Amazon Lex deletes any data provided before the timeout.
	//
	// You can specify between 60 (1 minute) and 86,400 (24 hours) seconds.
	IdleSessionTtlInSeconds *float64 `json:"idleSessionTtlInSeconds"`
	// The name of the field to filter the list of bots.
	Name *string `json:"name"`
	// The Amazon Resource Name (ARN) of the IAM role used to build and run the bot.
	RoleArn *string `json:"roleArn"`
	// Indicates whether Amazon Lex V2 should automatically build the locales for the bot after a change.
	AutoBuildBotLocales interface{} `json:"autoBuildBotLocales"`
	// The Amazon S3 location of files used to import a bot.
	//
	// The files must be in the import format specified in [JSON format for importing and exporting](https://docs.aws.amazon.com/lexv2/latest/dg/import-export-format.html) in the *Amazon Lex developer guide.*
	BotFileS3Location interface{} `json:"botFileS3Location"`
	// A list of locales for the bot.
	BotLocales interface{} `json:"botLocales"`
	// A list of tags to add to the bot.
	//
	// You can only add tags when you import a bot. You can't use the `UpdateBot` operation to update tags. To update tags, use the `TagResource` operation.
	BotTags interface{} `json:"botTags"`
	// The description of the version.
	Description *string `json:"description"`
	// A list of tags to add to the test alias for a bot.
	//
	// You can only add tags when you import a bot. You can't use the `UpdateAlias` operation to update tags. To update tags on the test alias, use the `TagResource` operation.
	TestBotAliasTags interface{} `json:"testBotAliasTags"`
}

// A CloudFormation `AWS::Lex::BotVersion`.
//
// Specifies a new version of the bot based on the `DRAFT` version. If the `DRAFT` version of this resource hasn't changed since you created the last version, Amazon Lex doesn't create a new version, it returns the last created version.
//
// When you specify the first version of a bot, Amazon Lex sets the version to 1. Subsequent versions increment by 1.
//
// TODO: EXAMPLE
//
type CfnBotVersion interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrBotVersion() *string
	BotId() *string
	SetBotId(val *string)
	BotVersionLocaleSpecification() interface{}
	SetBotVersionLocaleSpecification(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	LogicalId() *string
	Node() awscdk.ConstructNode
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

// The jsii proxy struct for CfnBotVersion
type jsiiProxy_CfnBotVersion struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnBotVersion) AttrBotVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrBotVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotVersion) BotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotVersion) BotVersionLocaleSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"botVersionLocaleSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotVersion) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotVersion) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotVersion) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotVersion) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lex::BotVersion`.
func NewCfnBotVersion(scope awscdk.Construct, id *string, props *CfnBotVersionProps) CfnBotVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnBotVersion{}

	_jsii_.Create(
		"monocdk.aws_lex.CfnBotVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lex::BotVersion`.
func NewCfnBotVersion_Override(c CfnBotVersion, scope awscdk.Construct, id *string, props *CfnBotVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lex.CfnBotVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnBotVersion) SetBotId(val *string) {
	_jsii_.Set(
		j,
		"botId",
		val,
	)
}

func (j *jsiiProxy_CfnBotVersion) SetBotVersionLocaleSpecification(val interface{}) {
	_jsii_.Set(
		j,
		"botVersionLocaleSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnBotVersion) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
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
func CfnBotVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lex.CfnBotVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnBotVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lex.CfnBotVersion",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnBotVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lex.CfnBotVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBotVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_lex.CfnBotVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnBotVersion) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnBotVersion) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnBotVersion) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnBotVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnBotVersion) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnBotVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnBotVersion) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnBotVersion) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnBotVersion) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnBotVersion) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnBotVersion) OnPrepare() {
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
func (c *jsiiProxy_CfnBotVersion) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnBotVersion) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnBotVersion) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnBotVersion) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnBotVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnBotVersion) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnBotVersion) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnBotVersion) ToString() *string {
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
func (c *jsiiProxy_CfnBotVersion) Validate() *[]*string {
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
func (c *jsiiProxy_CfnBotVersion) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The version of a bot used for a bot locale.
//
// TODO: EXAMPLE
//
type CfnBotVersion_BotVersionLocaleDetailsProperty struct {
	// The version of a bot used for a bot locale.
	SourceBotVersion *string `json:"sourceBotVersion"`
}

// TODO: EXAMPLE
//
type CfnBotVersion_BotVersionLocaleSpecificationProperty struct {
	// `CfnBotVersion.BotVersionLocaleSpecificationProperty.BotVersionLocaleDetails`.
	BotVersionLocaleDetails interface{} `json:"botVersionLocaleDetails"`
	// `CfnBotVersion.BotVersionLocaleSpecificationProperty.LocaleId`.
	LocaleId *string `json:"localeId"`
}

// Properties for defining a `CfnBotVersion`.
//
// TODO: EXAMPLE
//
type CfnBotVersionProps struct {
	// The unique identifier of the bot.
	BotId *string `json:"botId"`
	// Specifies the locales that Amazon Lex adds to this version.
	//
	// You can choose the Draft version or any other previously published version for each locale. When you specify a source version, the locale data is copied from the source version to the new version.
	BotVersionLocaleSpecification interface{} `json:"botVersionLocaleSpecification"`
	// The description of the version.
	Description *string `json:"description"`
}

// A CloudFormation `AWS::Lex::ResourcePolicy`.
//
// Specifies a new resource policy with the specified policy statements.
//
// TODO: EXAMPLE
//
type CfnResourcePolicy interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrId() *string
	AttrRevisionId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() awscdk.ConstructNode
	Policy() interface{}
	SetPolicy(val interface{})
	Ref() *string
	ResourceArn() *string
	SetResourceArn(val *string)
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

// The jsii proxy struct for CfnResourcePolicy
type jsiiProxy_CfnResourcePolicy struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnResourcePolicy) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) AttrRevisionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRevisionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) Policy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) ResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lex::ResourcePolicy`.
func NewCfnResourcePolicy(scope awscdk.Construct, id *string, props *CfnResourcePolicyProps) CfnResourcePolicy {
	_init_.Initialize()

	j := jsiiProxy_CfnResourcePolicy{}

	_jsii_.Create(
		"monocdk.aws_lex.CfnResourcePolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lex::ResourcePolicy`.
func NewCfnResourcePolicy_Override(c CfnResourcePolicy, scope awscdk.Construct, id *string, props *CfnResourcePolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lex.CfnResourcePolicy",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResourcePolicy) SetPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_CfnResourcePolicy) SetResourceArn(val *string) {
	_jsii_.Set(
		j,
		"resourceArn",
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
func CfnResourcePolicy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lex.CfnResourcePolicy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnResourcePolicy_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lex.CfnResourcePolicy",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnResourcePolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lex.CfnResourcePolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResourcePolicy_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_lex.CfnResourcePolicy",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnResourcePolicy) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnResourcePolicy) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnResourcePolicy) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnResourcePolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnResourcePolicy) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnResourcePolicy) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnResourcePolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnResourcePolicy) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnResourcePolicy) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnResourcePolicy) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnResourcePolicy) OnPrepare() {
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
func (c *jsiiProxy_CfnResourcePolicy) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnResourcePolicy) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnResourcePolicy) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnResourcePolicy) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnResourcePolicy) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnResourcePolicy) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnResourcePolicy) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnResourcePolicy) ToString() *string {
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
func (c *jsiiProxy_CfnResourcePolicy) Validate() *[]*string {
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
func (c *jsiiProxy_CfnResourcePolicy) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnResourcePolicy_PolicyProperty struct {
}

// Properties for defining a `CfnResourcePolicy`.
//
// TODO: EXAMPLE
//
type CfnResourcePolicyProps struct {
	// A resource policy to add to the resource.
	//
	// The policy is a JSON structure that contains one or more statements that define the policy. The policy must follow IAM syntax. If the policy isn't valid, Amazon Lex returns a validation exception.
	Policy interface{} `json:"policy"`
	// The Amazon Resource Name (ARN) of the bot or bot alias that the resource policy is attached to.
	ResourceArn *string `json:"resourceArn"`
}


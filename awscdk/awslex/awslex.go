package awslex

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslex/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Lex::Bot`.
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
	Node() constructs.Node
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
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

func (j *jsiiProxy_CfnBot) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnBot(scope constructs.Construct, id *string, props *CfnBotProps) CfnBot {
	_init_.Initialize()

	j := jsiiProxy_CfnBot{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lex.CfnBot",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lex::Bot`.
func NewCfnBot_Override(c CfnBot, scope constructs.Construct, id *string, props *CfnBotProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lex.CfnBot",
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
func CfnBot_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lex.CfnBot",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnBot_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lex.CfnBot",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnBot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lex.CfnBot",
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
		"aws-cdk-lib.aws_lex.CfnBot",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
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
func (c *jsiiProxy_CfnBot) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
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

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnBot) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
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

func (c *jsiiProxy_CfnBot) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnBot_BotLocaleProperty struct {
	// `CfnBot.BotLocaleProperty.Description`.
	Description *string `json:"description"`
	// `CfnBot.BotLocaleProperty.Intents`.
	Intents interface{} `json:"intents"`
	// `CfnBot.BotLocaleProperty.LocaleId`.
	LocaleId *string `json:"localeId"`
	// `CfnBot.BotLocaleProperty.NluConfidenceThreshold`.
	NluConfidenceThreshold *float64 `json:"nluConfidenceThreshold"`
	// `CfnBot.BotLocaleProperty.SlotTypes`.
	SlotTypes interface{} `json:"slotTypes"`
	// `CfnBot.BotLocaleProperty.VoiceSettings`.
	VoiceSettings interface{} `json:"voiceSettings"`
}

// TODO: EXAMPLE
//
type CfnBot_ButtonProperty struct {
	// `CfnBot.ButtonProperty.Text`.
	Text *string `json:"text"`
	// `CfnBot.ButtonProperty.Value`.
	Value *string `json:"value"`
}

// TODO: EXAMPLE
//
type CfnBot_CustomPayloadProperty struct {
	// `CfnBot.CustomPayloadProperty.Value`.
	Value *string `json:"value"`
}

// TODO: EXAMPLE
//
type CfnBot_DialogCodeHookSettingProperty struct {
	// `CfnBot.DialogCodeHookSettingProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
}

// TODO: EXAMPLE
//
type CfnBot_ExternalSourceSettingProperty struct {
	// `CfnBot.ExternalSourceSettingProperty.GrammarSlotTypeSetting`.
	GrammarSlotTypeSetting interface{} `json:"grammarSlotTypeSetting"`
}

// TODO: EXAMPLE
//
type CfnBot_FulfillmentCodeHookSettingProperty struct {
	// `CfnBot.FulfillmentCodeHookSettingProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `CfnBot.FulfillmentCodeHookSettingProperty.FulfillmentUpdatesSpecification`.
	FulfillmentUpdatesSpecification interface{} `json:"fulfillmentUpdatesSpecification"`
	// `CfnBot.FulfillmentCodeHookSettingProperty.PostFulfillmentStatusSpecification`.
	PostFulfillmentStatusSpecification interface{} `json:"postFulfillmentStatusSpecification"`
}

// TODO: EXAMPLE
//
type CfnBot_FulfillmentStartResponseSpecificationProperty struct {
	// `CfnBot.FulfillmentStartResponseSpecificationProperty.AllowInterrupt`.
	AllowInterrupt interface{} `json:"allowInterrupt"`
	// `CfnBot.FulfillmentStartResponseSpecificationProperty.DelayInSeconds`.
	DelayInSeconds *float64 `json:"delayInSeconds"`
	// `CfnBot.FulfillmentStartResponseSpecificationProperty.MessageGroups`.
	MessageGroups interface{} `json:"messageGroups"`
}

// TODO: EXAMPLE
//
type CfnBot_FulfillmentUpdateResponseSpecificationProperty struct {
	// `CfnBot.FulfillmentUpdateResponseSpecificationProperty.AllowInterrupt`.
	AllowInterrupt interface{} `json:"allowInterrupt"`
	// `CfnBot.FulfillmentUpdateResponseSpecificationProperty.FrequencyInSeconds`.
	FrequencyInSeconds *float64 `json:"frequencyInSeconds"`
	// `CfnBot.FulfillmentUpdateResponseSpecificationProperty.MessageGroups`.
	MessageGroups interface{} `json:"messageGroups"`
}

// TODO: EXAMPLE
//
type CfnBot_FulfillmentUpdatesSpecificationProperty struct {
	// `CfnBot.FulfillmentUpdatesSpecificationProperty.Active`.
	Active interface{} `json:"active"`
	// `CfnBot.FulfillmentUpdatesSpecificationProperty.StartResponse`.
	StartResponse interface{} `json:"startResponse"`
	// `CfnBot.FulfillmentUpdatesSpecificationProperty.TimeoutInSeconds`.
	TimeoutInSeconds *float64 `json:"timeoutInSeconds"`
	// `CfnBot.FulfillmentUpdatesSpecificationProperty.UpdateResponse`.
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
	// `CfnBot.GrammarSlotTypeSourceProperty.KmsKeyArn`.
	KmsKeyArn *string `json:"kmsKeyArn"`
	// `CfnBot.GrammarSlotTypeSourceProperty.S3BucketName`.
	S3BucketName *string `json:"s3BucketName"`
	// `CfnBot.GrammarSlotTypeSourceProperty.S3ObjectKey`.
	S3ObjectKey *string `json:"s3ObjectKey"`
}

// TODO: EXAMPLE
//
type CfnBot_ImageResponseCardProperty struct {
	// `CfnBot.ImageResponseCardProperty.Buttons`.
	Buttons interface{} `json:"buttons"`
	// `CfnBot.ImageResponseCardProperty.ImageUrl`.
	ImageUrl *string `json:"imageUrl"`
	// `CfnBot.ImageResponseCardProperty.Subtitle`.
	Subtitle *string `json:"subtitle"`
	// `CfnBot.ImageResponseCardProperty.Title`.
	Title *string `json:"title"`
}

// TODO: EXAMPLE
//
type CfnBot_InputContextProperty struct {
	// `CfnBot.InputContextProperty.Name`.
	Name *string `json:"name"`
}

// TODO: EXAMPLE
//
type CfnBot_IntentClosingSettingProperty struct {
	// `CfnBot.IntentClosingSettingProperty.ClosingResponse`.
	ClosingResponse interface{} `json:"closingResponse"`
	// `CfnBot.IntentClosingSettingProperty.IsActive`.
	IsActive interface{} `json:"isActive"`
}

// TODO: EXAMPLE
//
type CfnBot_IntentConfirmationSettingProperty struct {
	// `CfnBot.IntentConfirmationSettingProperty.DeclinationResponse`.
	DeclinationResponse interface{} `json:"declinationResponse"`
	// `CfnBot.IntentConfirmationSettingProperty.IsActive`.
	IsActive interface{} `json:"isActive"`
	// `CfnBot.IntentConfirmationSettingProperty.PromptSpecification`.
	PromptSpecification interface{} `json:"promptSpecification"`
}

// TODO: EXAMPLE
//
type CfnBot_IntentProperty struct {
	// `CfnBot.IntentProperty.Description`.
	Description *string `json:"description"`
	// `CfnBot.IntentProperty.DialogCodeHook`.
	DialogCodeHook interface{} `json:"dialogCodeHook"`
	// `CfnBot.IntentProperty.FulfillmentCodeHook`.
	FulfillmentCodeHook interface{} `json:"fulfillmentCodeHook"`
	// `CfnBot.IntentProperty.InputContexts`.
	InputContexts interface{} `json:"inputContexts"`
	// `CfnBot.IntentProperty.IntentClosingSetting`.
	IntentClosingSetting interface{} `json:"intentClosingSetting"`
	// `CfnBot.IntentProperty.IntentConfirmationSetting`.
	IntentConfirmationSetting interface{} `json:"intentConfirmationSetting"`
	// `CfnBot.IntentProperty.KendraConfiguration`.
	KendraConfiguration interface{} `json:"kendraConfiguration"`
	// `CfnBot.IntentProperty.Name`.
	Name *string `json:"name"`
	// `CfnBot.IntentProperty.OutputContexts`.
	OutputContexts interface{} `json:"outputContexts"`
	// `CfnBot.IntentProperty.ParentIntentSignature`.
	ParentIntentSignature *string `json:"parentIntentSignature"`
	// `CfnBot.IntentProperty.SampleUtterances`.
	SampleUtterances interface{} `json:"sampleUtterances"`
	// `CfnBot.IntentProperty.SlotPriorities`.
	SlotPriorities interface{} `json:"slotPriorities"`
	// `CfnBot.IntentProperty.Slots`.
	Slots interface{} `json:"slots"`
}

// TODO: EXAMPLE
//
type CfnBot_KendraConfigurationProperty struct {
	// `CfnBot.KendraConfigurationProperty.KendraIndex`.
	KendraIndex *string `json:"kendraIndex"`
	// `CfnBot.KendraConfigurationProperty.QueryFilterString`.
	QueryFilterString *string `json:"queryFilterString"`
	// `CfnBot.KendraConfigurationProperty.QueryFilterStringEnabled`.
	QueryFilterStringEnabled interface{} `json:"queryFilterStringEnabled"`
}

// TODO: EXAMPLE
//
type CfnBot_MessageGroupProperty struct {
	// `CfnBot.MessageGroupProperty.Message`.
	Message interface{} `json:"message"`
	// `CfnBot.MessageGroupProperty.Variations`.
	Variations interface{} `json:"variations"`
}

// TODO: EXAMPLE
//
type CfnBot_MessageProperty struct {
	// `CfnBot.MessageProperty.CustomPayload`.
	CustomPayload interface{} `json:"customPayload"`
	// `CfnBot.MessageProperty.ImageResponseCard`.
	ImageResponseCard interface{} `json:"imageResponseCard"`
	// `CfnBot.MessageProperty.PlainTextMessage`.
	PlainTextMessage interface{} `json:"plainTextMessage"`
	// `CfnBot.MessageProperty.SSMLMessage`.
	SsmlMessage interface{} `json:"ssmlMessage"`
}

// TODO: EXAMPLE
//
type CfnBot_MultipleValuesSettingProperty struct {
	// `CfnBot.MultipleValuesSettingProperty.AllowMultipleValues`.
	AllowMultipleValues interface{} `json:"allowMultipleValues"`
}

// TODO: EXAMPLE
//
type CfnBot_ObfuscationSettingProperty struct {
	// `CfnBot.ObfuscationSettingProperty.ObfuscationSettingType`.
	ObfuscationSettingType *string `json:"obfuscationSettingType"`
}

// TODO: EXAMPLE
//
type CfnBot_OutputContextProperty struct {
	// `CfnBot.OutputContextProperty.Name`.
	Name *string `json:"name"`
	// `CfnBot.OutputContextProperty.TimeToLiveInSeconds`.
	TimeToLiveInSeconds *float64 `json:"timeToLiveInSeconds"`
	// `CfnBot.OutputContextProperty.TurnsToLive`.
	TurnsToLive *float64 `json:"turnsToLive"`
}

// TODO: EXAMPLE
//
type CfnBot_PlainTextMessageProperty struct {
	// `CfnBot.PlainTextMessageProperty.Value`.
	Value *string `json:"value"`
}

// TODO: EXAMPLE
//
type CfnBot_PostFulfillmentStatusSpecificationProperty struct {
	// `CfnBot.PostFulfillmentStatusSpecificationProperty.FailureResponse`.
	FailureResponse interface{} `json:"failureResponse"`
	// `CfnBot.PostFulfillmentStatusSpecificationProperty.SuccessResponse`.
	SuccessResponse interface{} `json:"successResponse"`
	// `CfnBot.PostFulfillmentStatusSpecificationProperty.TimeoutResponse`.
	TimeoutResponse interface{} `json:"timeoutResponse"`
}

// TODO: EXAMPLE
//
type CfnBot_PromptSpecificationProperty struct {
	// `CfnBot.PromptSpecificationProperty.AllowInterrupt`.
	AllowInterrupt interface{} `json:"allowInterrupt"`
	// `CfnBot.PromptSpecificationProperty.MaxRetries`.
	MaxRetries *float64 `json:"maxRetries"`
	// `CfnBot.PromptSpecificationProperty.MessageGroupsList`.
	MessageGroupsList interface{} `json:"messageGroupsList"`
}

// TODO: EXAMPLE
//
type CfnBot_ResponseSpecificationProperty struct {
	// `CfnBot.ResponseSpecificationProperty.AllowInterrupt`.
	AllowInterrupt interface{} `json:"allowInterrupt"`
	// `CfnBot.ResponseSpecificationProperty.MessageGroupsList`.
	MessageGroupsList interface{} `json:"messageGroupsList"`
}

// TODO: EXAMPLE
//
type CfnBot_S3LocationProperty struct {
	// `CfnBot.S3LocationProperty.S3Bucket`.
	S3Bucket *string `json:"s3Bucket"`
	// `CfnBot.S3LocationProperty.S3ObjectKey`.
	S3ObjectKey *string `json:"s3ObjectKey"`
	// `CfnBot.S3LocationProperty.S3ObjectVersion`.
	S3ObjectVersion *string `json:"s3ObjectVersion"`
}

// TODO: EXAMPLE
//
type CfnBot_SSMLMessageProperty struct {
	// `CfnBot.SSMLMessageProperty.Value`.
	Value *string `json:"value"`
}

// TODO: EXAMPLE
//
type CfnBot_SampleUtteranceProperty struct {
	// `CfnBot.SampleUtteranceProperty.Utterance`.
	Utterance *string `json:"utterance"`
}

// TODO: EXAMPLE
//
type CfnBot_SampleValueProperty struct {
	// `CfnBot.SampleValueProperty.Value`.
	Value *string `json:"value"`
}

// TODO: EXAMPLE
//
type CfnBot_SlotDefaultValueProperty struct {
	// `CfnBot.SlotDefaultValueProperty.DefaultValue`.
	DefaultValue *string `json:"defaultValue"`
}

// TODO: EXAMPLE
//
type CfnBot_SlotDefaultValueSpecificationProperty struct {
	// `CfnBot.SlotDefaultValueSpecificationProperty.DefaultValueList`.
	DefaultValueList interface{} `json:"defaultValueList"`
}

// TODO: EXAMPLE
//
type CfnBot_SlotPriorityProperty struct {
	// `CfnBot.SlotPriorityProperty.Priority`.
	Priority *float64 `json:"priority"`
	// `CfnBot.SlotPriorityProperty.SlotName`.
	SlotName *string `json:"slotName"`
}

// TODO: EXAMPLE
//
type CfnBot_SlotProperty struct {
	// `CfnBot.SlotProperty.Description`.
	Description *string `json:"description"`
	// `CfnBot.SlotProperty.MultipleValuesSetting`.
	MultipleValuesSetting interface{} `json:"multipleValuesSetting"`
	// `CfnBot.SlotProperty.Name`.
	Name *string `json:"name"`
	// `CfnBot.SlotProperty.ObfuscationSetting`.
	ObfuscationSetting interface{} `json:"obfuscationSetting"`
	// `CfnBot.SlotProperty.SlotTypeName`.
	SlotTypeName *string `json:"slotTypeName"`
	// `CfnBot.SlotProperty.ValueElicitationSetting`.
	ValueElicitationSetting interface{} `json:"valueElicitationSetting"`
}

// TODO: EXAMPLE
//
type CfnBot_SlotTypeProperty struct {
	// `CfnBot.SlotTypeProperty.Description`.
	Description *string `json:"description"`
	// `CfnBot.SlotTypeProperty.ExternalSourceSetting`.
	ExternalSourceSetting interface{} `json:"externalSourceSetting"`
	// `CfnBot.SlotTypeProperty.Name`.
	Name *string `json:"name"`
	// `CfnBot.SlotTypeProperty.ParentSlotTypeSignature`.
	ParentSlotTypeSignature *string `json:"parentSlotTypeSignature"`
	// `CfnBot.SlotTypeProperty.SlotTypeValues`.
	SlotTypeValues interface{} `json:"slotTypeValues"`
	// `CfnBot.SlotTypeProperty.ValueSelectionSetting`.
	ValueSelectionSetting interface{} `json:"valueSelectionSetting"`
}

// TODO: EXAMPLE
//
type CfnBot_SlotTypeValueProperty struct {
	// `CfnBot.SlotTypeValueProperty.SampleValue`.
	SampleValue interface{} `json:"sampleValue"`
	// `CfnBot.SlotTypeValueProperty.Synonyms`.
	Synonyms interface{} `json:"synonyms"`
}

// TODO: EXAMPLE
//
type CfnBot_SlotValueElicitationSettingProperty struct {
	// `CfnBot.SlotValueElicitationSettingProperty.DefaultValueSpecification`.
	DefaultValueSpecification interface{} `json:"defaultValueSpecification"`
	// `CfnBot.SlotValueElicitationSettingProperty.PromptSpecification`.
	PromptSpecification interface{} `json:"promptSpecification"`
	// `CfnBot.SlotValueElicitationSettingProperty.SampleUtterances`.
	SampleUtterances interface{} `json:"sampleUtterances"`
	// `CfnBot.SlotValueElicitationSettingProperty.SlotConstraint`.
	SlotConstraint *string `json:"slotConstraint"`
	// `CfnBot.SlotValueElicitationSettingProperty.WaitAndContinueSpecification`.
	WaitAndContinueSpecification interface{} `json:"waitAndContinueSpecification"`
}

// TODO: EXAMPLE
//
type CfnBot_SlotValueRegexFilterProperty struct {
	// `CfnBot.SlotValueRegexFilterProperty.Pattern`.
	Pattern *string `json:"pattern"`
}

// TODO: EXAMPLE
//
type CfnBot_SlotValueSelectionSettingProperty struct {
	// `CfnBot.SlotValueSelectionSettingProperty.RegexFilter`.
	RegexFilter interface{} `json:"regexFilter"`
	// `CfnBot.SlotValueSelectionSettingProperty.ResolutionStrategy`.
	ResolutionStrategy *string `json:"resolutionStrategy"`
}

// TODO: EXAMPLE
//
type CfnBot_StillWaitingResponseSpecificationProperty struct {
	// `CfnBot.StillWaitingResponseSpecificationProperty.AllowInterrupt`.
	AllowInterrupt interface{} `json:"allowInterrupt"`
	// `CfnBot.StillWaitingResponseSpecificationProperty.FrequencyInSeconds`.
	FrequencyInSeconds *float64 `json:"frequencyInSeconds"`
	// `CfnBot.StillWaitingResponseSpecificationProperty.MessageGroupsList`.
	MessageGroupsList interface{} `json:"messageGroupsList"`
	// `CfnBot.StillWaitingResponseSpecificationProperty.TimeoutInSeconds`.
	TimeoutInSeconds *float64 `json:"timeoutInSeconds"`
}

// TODO: EXAMPLE
//
type CfnBot_VoiceSettingsProperty struct {
	// `CfnBot.VoiceSettingsProperty.VoiceId`.
	VoiceId *string `json:"voiceId"`
}

// TODO: EXAMPLE
//
type CfnBot_WaitAndContinueSpecificationProperty struct {
	// `CfnBot.WaitAndContinueSpecificationProperty.ContinueResponse`.
	ContinueResponse interface{} `json:"continueResponse"`
	// `CfnBot.WaitAndContinueSpecificationProperty.IsActive`.
	IsActive interface{} `json:"isActive"`
	// `CfnBot.WaitAndContinueSpecificationProperty.StillWaitingResponse`.
	StillWaitingResponse interface{} `json:"stillWaitingResponse"`
	// `CfnBot.WaitAndContinueSpecificationProperty.WaitingResponse`.
	WaitingResponse interface{} `json:"waitingResponse"`
}

// A CloudFormation `AWS::Lex::BotAlias`.
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
	Node() constructs.Node
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
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

func (j *jsiiProxy_CfnBotAlias) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnBotAlias(scope constructs.Construct, id *string, props *CfnBotAliasProps) CfnBotAlias {
	_init_.Initialize()

	j := jsiiProxy_CfnBotAlias{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lex.CfnBotAlias",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lex::BotAlias`.
func NewCfnBotAlias_Override(c CfnBotAlias, scope constructs.Construct, id *string, props *CfnBotAliasProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lex.CfnBotAlias",
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
func CfnBotAlias_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lex.CfnBotAlias",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnBotAlias_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lex.CfnBotAlias",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnBotAlias_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lex.CfnBotAlias",
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
		"aws-cdk-lib.aws_lex.CfnBotAlias",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
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
func (c *jsiiProxy_CfnBotAlias) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
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

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnBotAlias) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
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

func (c *jsiiProxy_CfnBotAlias) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnBotAlias_AudioLogDestinationProperty struct {
	// `CfnBotAlias.AudioLogDestinationProperty.S3Bucket`.
	S3Bucket interface{} `json:"s3Bucket"`
}

// TODO: EXAMPLE
//
type CfnBotAlias_AudioLogSettingProperty struct {
	// `CfnBotAlias.AudioLogSettingProperty.Destination`.
	Destination interface{} `json:"destination"`
	// `CfnBotAlias.AudioLogSettingProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
}

// TODO: EXAMPLE
//
type CfnBotAlias_BotAliasLocaleSettingsItemProperty struct {
	// `CfnBotAlias.BotAliasLocaleSettingsItemProperty.BotAliasLocaleSetting`.
	BotAliasLocaleSetting interface{} `json:"botAliasLocaleSetting"`
	// `CfnBotAlias.BotAliasLocaleSettingsItemProperty.LocaleId`.
	LocaleId *string `json:"localeId"`
}

// TODO: EXAMPLE
//
type CfnBotAlias_BotAliasLocaleSettingsProperty struct {
	// `CfnBotAlias.BotAliasLocaleSettingsProperty.CodeHookSpecification`.
	CodeHookSpecification interface{} `json:"codeHookSpecification"`
	// `CfnBotAlias.BotAliasLocaleSettingsProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
}

// TODO: EXAMPLE
//
type CfnBotAlias_CloudWatchLogGroupLogDestinationProperty struct {
	// `CfnBotAlias.CloudWatchLogGroupLogDestinationProperty.CloudWatchLogGroupArn`.
	CloudWatchLogGroupArn *string `json:"cloudWatchLogGroupArn"`
	// `CfnBotAlias.CloudWatchLogGroupLogDestinationProperty.LogPrefix`.
	LogPrefix *string `json:"logPrefix"`
}

// TODO: EXAMPLE
//
type CfnBotAlias_CodeHookSpecificationProperty struct {
	// `CfnBotAlias.CodeHookSpecificationProperty.LambdaCodeHook`.
	LambdaCodeHook interface{} `json:"lambdaCodeHook"`
}

// TODO: EXAMPLE
//
type CfnBotAlias_ConversationLogSettingsProperty struct {
	// `CfnBotAlias.ConversationLogSettingsProperty.AudioLogSettings`.
	AudioLogSettings interface{} `json:"audioLogSettings"`
	// `CfnBotAlias.ConversationLogSettingsProperty.TextLogSettings`.
	TextLogSettings interface{} `json:"textLogSettings"`
}

// TODO: EXAMPLE
//
type CfnBotAlias_LambdaCodeHookProperty struct {
	// `CfnBotAlias.LambdaCodeHookProperty.CodeHookInterfaceVersion`.
	CodeHookInterfaceVersion *string `json:"codeHookInterfaceVersion"`
	// `CfnBotAlias.LambdaCodeHookProperty.LambdaArn`.
	LambdaArn *string `json:"lambdaArn"`
}

// TODO: EXAMPLE
//
type CfnBotAlias_S3BucketLogDestinationProperty struct {
	// `CfnBotAlias.S3BucketLogDestinationProperty.KmsKeyArn`.
	KmsKeyArn *string `json:"kmsKeyArn"`
	// `CfnBotAlias.S3BucketLogDestinationProperty.LogPrefix`.
	LogPrefix *string `json:"logPrefix"`
	// `CfnBotAlias.S3BucketLogDestinationProperty.S3BucketArn`.
	S3BucketArn *string `json:"s3BucketArn"`
}

// TODO: EXAMPLE
//
type CfnBotAlias_TextLogDestinationProperty struct {
}

// TODO: EXAMPLE
//
type CfnBotAlias_TextLogSettingProperty struct {
	// `CfnBotAlias.TextLogSettingProperty.Destination`.
	Destination interface{} `json:"destination"`
	// `CfnBotAlias.TextLogSettingProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
}

// Properties for defining a `AWS::Lex::BotAlias`.
//
// TODO: EXAMPLE
//
type CfnBotAliasProps struct {
	// `AWS::Lex::BotAlias.BotAliasLocaleSettings`.
	BotAliasLocaleSettings interface{} `json:"botAliasLocaleSettings"`
	// `AWS::Lex::BotAlias.BotAliasName`.
	BotAliasName *string `json:"botAliasName"`
	// `AWS::Lex::BotAlias.BotAliasTags`.
	BotAliasTags interface{} `json:"botAliasTags"`
	// `AWS::Lex::BotAlias.BotId`.
	BotId *string `json:"botId"`
	// `AWS::Lex::BotAlias.BotVersion`.
	BotVersion *string `json:"botVersion"`
	// `AWS::Lex::BotAlias.ConversationLogSettings`.
	ConversationLogSettings interface{} `json:"conversationLogSettings"`
	// `AWS::Lex::BotAlias.Description`.
	Description *string `json:"description"`
	// `AWS::Lex::BotAlias.SentimentAnalysisSettings`.
	SentimentAnalysisSettings interface{} `json:"sentimentAnalysisSettings"`
}

// Properties for defining a `AWS::Lex::Bot`.
//
// TODO: EXAMPLE
//
type CfnBotProps struct {
	// `AWS::Lex::Bot.AutoBuildBotLocales`.
	AutoBuildBotLocales interface{} `json:"autoBuildBotLocales"`
	// `AWS::Lex::Bot.BotFileS3Location`.
	BotFileS3Location interface{} `json:"botFileS3Location"`
	// `AWS::Lex::Bot.BotLocales`.
	BotLocales interface{} `json:"botLocales"`
	// `AWS::Lex::Bot.BotTags`.
	BotTags interface{} `json:"botTags"`
	// `AWS::Lex::Bot.DataPrivacy`.
	DataPrivacy interface{} `json:"dataPrivacy"`
	// `AWS::Lex::Bot.Description`.
	Description *string `json:"description"`
	// `AWS::Lex::Bot.IdleSessionTTLInSeconds`.
	IdleSessionTtlInSeconds *float64 `json:"idleSessionTtlInSeconds"`
	// `AWS::Lex::Bot.Name`.
	Name *string `json:"name"`
	// `AWS::Lex::Bot.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `AWS::Lex::Bot.TestBotAliasTags`.
	TestBotAliasTags interface{} `json:"testBotAliasTags"`
}

// A CloudFormation `AWS::Lex::BotVersion`.
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
	Node() constructs.Node
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
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

func (j *jsiiProxy_CfnBotVersion) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnBotVersion(scope constructs.Construct, id *string, props *CfnBotVersionProps) CfnBotVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnBotVersion{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lex.CfnBotVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lex::BotVersion`.
func NewCfnBotVersion_Override(c CfnBotVersion, scope constructs.Construct, id *string, props *CfnBotVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lex.CfnBotVersion",
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
func CfnBotVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lex.CfnBotVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnBotVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lex.CfnBotVersion",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnBotVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lex.CfnBotVersion",
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
		"aws-cdk-lib.aws_lex.CfnBotVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
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
func (c *jsiiProxy_CfnBotVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
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

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnBotVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
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

func (c *jsiiProxy_CfnBotVersion) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnBotVersion_BotVersionLocaleDetailsProperty struct {
	// `CfnBotVersion.BotVersionLocaleDetailsProperty.SourceBotVersion`.
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

// Properties for defining a `AWS::Lex::BotVersion`.
//
// TODO: EXAMPLE
//
type CfnBotVersionProps struct {
	// `AWS::Lex::BotVersion.BotId`.
	BotId *string `json:"botId"`
	// `AWS::Lex::BotVersion.BotVersionLocaleSpecification`.
	BotVersionLocaleSpecification interface{} `json:"botVersionLocaleSpecification"`
	// `AWS::Lex::BotVersion.Description`.
	Description *string `json:"description"`
}

// A CloudFormation `AWS::Lex::ResourcePolicy`.
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
	Node() constructs.Node
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
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

func (j *jsiiProxy_CfnResourcePolicy) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnResourcePolicy(scope constructs.Construct, id *string, props *CfnResourcePolicyProps) CfnResourcePolicy {
	_init_.Initialize()

	j := jsiiProxy_CfnResourcePolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lex.CfnResourcePolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lex::ResourcePolicy`.
func NewCfnResourcePolicy_Override(c CfnResourcePolicy, scope constructs.Construct, id *string, props *CfnResourcePolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lex.CfnResourcePolicy",
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
func CfnResourcePolicy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lex.CfnResourcePolicy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnResourcePolicy_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lex.CfnResourcePolicy",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnResourcePolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lex.CfnResourcePolicy",
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
		"aws-cdk-lib.aws_lex.CfnResourcePolicy",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
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
func (c *jsiiProxy_CfnResourcePolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
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

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnResourcePolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
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

// Properties for defining a `AWS::Lex::ResourcePolicy`.
//
// TODO: EXAMPLE
//
type CfnResourcePolicyProps struct {
	// `AWS::Lex::ResourcePolicy.Policy`.
	Policy interface{} `json:"policy"`
	// `AWS::Lex::ResourcePolicy.ResourceArn`.
	ResourceArn *string `json:"resourceArn"`
}


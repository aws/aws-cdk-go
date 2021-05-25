package awsiotevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiotevents/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::IoTEvents::DetectorModel`.
type CfnDetectorModel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DetectorModelDefinition() interface{}
	SetDetectorModelDefinition(val interface{})
	DetectorModelDescription() *string
	SetDetectorModelDescription(val *string)
	DetectorModelName() *string
	SetDetectorModelName(val *string)
	EvaluationMethod() *string
	SetEvaluationMethod(val *string)
	Key() *string
	SetKey(val *string)
	LogicalId() *string
	Node() awscdk.ConstructNode
	Ref() *string
	RoleArn() *string
	SetRoleArn(val *string)
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

// The jsii proxy struct for CfnDetectorModel
type jsiiProxy_CfnDetectorModel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDetectorModel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) DetectorModelDefinition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detectorModelDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) DetectorModelDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detectorModelDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) DetectorModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detectorModelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) EvaluationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDetectorModel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTEvents::DetectorModel`.
func NewCfnDetectorModel(scope awscdk.Construct, id *string, props *CfnDetectorModelProps) CfnDetectorModel {
	_init_.Initialize()

	j := jsiiProxy_CfnDetectorModel{}

	_jsii_.Create(
		"monocdk.aws_iotevents.CfnDetectorModel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTEvents::DetectorModel`.
func NewCfnDetectorModel_Override(c CfnDetectorModel, scope awscdk.Construct, id *string, props *CfnDetectorModelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iotevents.CfnDetectorModel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDetectorModel) SetDetectorModelDefinition(val interface{}) {
	_jsii_.Set(
		j,
		"detectorModelDefinition",
		val,
	)
}

func (j *jsiiProxy_CfnDetectorModel) SetDetectorModelDescription(val *string) {
	_jsii_.Set(
		j,
		"detectorModelDescription",
		val,
	)
}

func (j *jsiiProxy_CfnDetectorModel) SetDetectorModelName(val *string) {
	_jsii_.Set(
		j,
		"detectorModelName",
		val,
	)
}

func (j *jsiiProxy_CfnDetectorModel) SetEvaluationMethod(val *string) {
	_jsii_.Set(
		j,
		"evaluationMethod",
		val,
	)
}

func (j *jsiiProxy_CfnDetectorModel) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_CfnDetectorModel) SetRoleArn(val *string) {
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
func CfnDetectorModel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotevents.CfnDetectorModel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDetectorModel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotevents.CfnDetectorModel",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDetectorModel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotevents.CfnDetectorModel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDetectorModel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_iotevents.CfnDetectorModel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnDetectorModel) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnDetectorModel) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnDetectorModel) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnDetectorModel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnDetectorModel) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnDetectorModel) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnDetectorModel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnDetectorModel) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnDetectorModel) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnDetectorModel) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnDetectorModel) OnPrepare() {
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
func (c *jsiiProxy_CfnDetectorModel) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnDetectorModel) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnDetectorModel) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnDetectorModel) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDetectorModel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnDetectorModel) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnDetectorModel) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnDetectorModel) ToString() *string {
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
func (c *jsiiProxy_CfnDetectorModel) Validate() *[]*string {
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
func (c *jsiiProxy_CfnDetectorModel) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnDetectorModel_ActionProperty struct {
	// `CfnDetectorModel.ActionProperty.ClearTimer`.
	ClearTimer interface{} `json:"clearTimer"`
	// `CfnDetectorModel.ActionProperty.DynamoDB`.
	DynamoDb interface{} `json:"dynamoDb"`
	// `CfnDetectorModel.ActionProperty.DynamoDBv2`.
	DynamoDBv2 interface{} `json:"dynamoDBv2"`
	// `CfnDetectorModel.ActionProperty.Firehose`.
	Firehose interface{} `json:"firehose"`
	// `CfnDetectorModel.ActionProperty.IotEvents`.
	IotEvents interface{} `json:"iotEvents"`
	// `CfnDetectorModel.ActionProperty.IotSiteWise`.
	IotSiteWise interface{} `json:"iotSiteWise"`
	// `CfnDetectorModel.ActionProperty.IotTopicPublish`.
	IotTopicPublish interface{} `json:"iotTopicPublish"`
	// `CfnDetectorModel.ActionProperty.Lambda`.
	Lambda interface{} `json:"lambda"`
	// `CfnDetectorModel.ActionProperty.ResetTimer`.
	ResetTimer interface{} `json:"resetTimer"`
	// `CfnDetectorModel.ActionProperty.SetTimer`.
	SetTimer interface{} `json:"setTimer"`
	// `CfnDetectorModel.ActionProperty.SetVariable`.
	SetVariable interface{} `json:"setVariable"`
	// `CfnDetectorModel.ActionProperty.Sns`.
	Sns interface{} `json:"sns"`
	// `CfnDetectorModel.ActionProperty.Sqs`.
	Sqs interface{} `json:"sqs"`
}

type CfnDetectorModel_AssetPropertyTimestampProperty struct {
	// `CfnDetectorModel.AssetPropertyTimestampProperty.TimeInSeconds`.
	TimeInSeconds *string `json:"timeInSeconds"`
	// `CfnDetectorModel.AssetPropertyTimestampProperty.OffsetInNanos`.
	OffsetInNanos *string `json:"offsetInNanos"`
}

type CfnDetectorModel_AssetPropertyValueProperty struct {
	// `CfnDetectorModel.AssetPropertyValueProperty.Value`.
	Value interface{} `json:"value"`
	// `CfnDetectorModel.AssetPropertyValueProperty.Quality`.
	Quality *string `json:"quality"`
	// `CfnDetectorModel.AssetPropertyValueProperty.Timestamp`.
	Timestamp interface{} `json:"timestamp"`
}

type CfnDetectorModel_AssetPropertyVariantProperty struct {
	// `CfnDetectorModel.AssetPropertyVariantProperty.BooleanValue`.
	BooleanValue *string `json:"booleanValue"`
	// `CfnDetectorModel.AssetPropertyVariantProperty.DoubleValue`.
	DoubleValue *string `json:"doubleValue"`
	// `CfnDetectorModel.AssetPropertyVariantProperty.IntegerValue`.
	IntegerValue *string `json:"integerValue"`
	// `CfnDetectorModel.AssetPropertyVariantProperty.StringValue`.
	StringValue *string `json:"stringValue"`
}

type CfnDetectorModel_ClearTimerProperty struct {
	// `CfnDetectorModel.ClearTimerProperty.TimerName`.
	TimerName *string `json:"timerName"`
}

type CfnDetectorModel_DetectorModelDefinitionProperty struct {
	// `CfnDetectorModel.DetectorModelDefinitionProperty.InitialStateName`.
	InitialStateName *string `json:"initialStateName"`
	// `CfnDetectorModel.DetectorModelDefinitionProperty.States`.
	States interface{} `json:"states"`
}

type CfnDetectorModel_DynamoDBProperty struct {
	// `CfnDetectorModel.DynamoDBProperty.HashKeyField`.
	HashKeyField *string `json:"hashKeyField"`
	// `CfnDetectorModel.DynamoDBProperty.HashKeyValue`.
	HashKeyValue *string `json:"hashKeyValue"`
	// `CfnDetectorModel.DynamoDBProperty.TableName`.
	TableName *string `json:"tableName"`
	// `CfnDetectorModel.DynamoDBProperty.HashKeyType`.
	HashKeyType *string `json:"hashKeyType"`
	// `CfnDetectorModel.DynamoDBProperty.Operation`.
	Operation *string `json:"operation"`
	// `CfnDetectorModel.DynamoDBProperty.Payload`.
	Payload interface{} `json:"payload"`
	// `CfnDetectorModel.DynamoDBProperty.PayloadField`.
	PayloadField *string `json:"payloadField"`
	// `CfnDetectorModel.DynamoDBProperty.RangeKeyField`.
	RangeKeyField *string `json:"rangeKeyField"`
	// `CfnDetectorModel.DynamoDBProperty.RangeKeyType`.
	RangeKeyType *string `json:"rangeKeyType"`
	// `CfnDetectorModel.DynamoDBProperty.RangeKeyValue`.
	RangeKeyValue *string `json:"rangeKeyValue"`
}

type CfnDetectorModel_DynamoDBv2Property struct {
	// `CfnDetectorModel.DynamoDBv2Property.TableName`.
	TableName *string `json:"tableName"`
	// `CfnDetectorModel.DynamoDBv2Property.Payload`.
	Payload interface{} `json:"payload"`
}

type CfnDetectorModel_EventProperty struct {
	// `CfnDetectorModel.EventProperty.EventName`.
	EventName *string `json:"eventName"`
	// `CfnDetectorModel.EventProperty.Actions`.
	Actions interface{} `json:"actions"`
	// `CfnDetectorModel.EventProperty.Condition`.
	Condition *string `json:"condition"`
}

type CfnDetectorModel_FirehoseProperty struct {
	// `CfnDetectorModel.FirehoseProperty.DeliveryStreamName`.
	DeliveryStreamName *string `json:"deliveryStreamName"`
	// `CfnDetectorModel.FirehoseProperty.Payload`.
	Payload interface{} `json:"payload"`
	// `CfnDetectorModel.FirehoseProperty.Separator`.
	Separator *string `json:"separator"`
}

type CfnDetectorModel_IotEventsProperty struct {
	// `CfnDetectorModel.IotEventsProperty.InputName`.
	InputName *string `json:"inputName"`
	// `CfnDetectorModel.IotEventsProperty.Payload`.
	Payload interface{} `json:"payload"`
}

type CfnDetectorModel_IotSiteWiseProperty struct {
	// `CfnDetectorModel.IotSiteWiseProperty.PropertyValue`.
	PropertyValue interface{} `json:"propertyValue"`
	// `CfnDetectorModel.IotSiteWiseProperty.AssetId`.
	AssetId *string `json:"assetId"`
	// `CfnDetectorModel.IotSiteWiseProperty.EntryId`.
	EntryId *string `json:"entryId"`
	// `CfnDetectorModel.IotSiteWiseProperty.PropertyAlias`.
	PropertyAlias *string `json:"propertyAlias"`
	// `CfnDetectorModel.IotSiteWiseProperty.PropertyId`.
	PropertyId *string `json:"propertyId"`
}

type CfnDetectorModel_IotTopicPublishProperty struct {
	// `CfnDetectorModel.IotTopicPublishProperty.MqttTopic`.
	MqttTopic *string `json:"mqttTopic"`
	// `CfnDetectorModel.IotTopicPublishProperty.Payload`.
	Payload interface{} `json:"payload"`
}

type CfnDetectorModel_LambdaProperty struct {
	// `CfnDetectorModel.LambdaProperty.FunctionArn`.
	FunctionArn *string `json:"functionArn"`
	// `CfnDetectorModel.LambdaProperty.Payload`.
	Payload interface{} `json:"payload"`
}

type CfnDetectorModel_OnEnterProperty struct {
	// `CfnDetectorModel.OnEnterProperty.Events`.
	Events interface{} `json:"events"`
}

type CfnDetectorModel_OnExitProperty struct {
	// `CfnDetectorModel.OnExitProperty.Events`.
	Events interface{} `json:"events"`
}

type CfnDetectorModel_OnInputProperty struct {
	// `CfnDetectorModel.OnInputProperty.Events`.
	Events interface{} `json:"events"`
	// `CfnDetectorModel.OnInputProperty.TransitionEvents`.
	TransitionEvents interface{} `json:"transitionEvents"`
}

type CfnDetectorModel_PayloadProperty struct {
	// `CfnDetectorModel.PayloadProperty.ContentExpression`.
	ContentExpression *string `json:"contentExpression"`
	// `CfnDetectorModel.PayloadProperty.Type`.
	Type *string `json:"type"`
}

type CfnDetectorModel_ResetTimerProperty struct {
	// `CfnDetectorModel.ResetTimerProperty.TimerName`.
	TimerName *string `json:"timerName"`
}

type CfnDetectorModel_SetTimerProperty struct {
	// `CfnDetectorModel.SetTimerProperty.TimerName`.
	TimerName *string `json:"timerName"`
	// `CfnDetectorModel.SetTimerProperty.DurationExpression`.
	DurationExpression *string `json:"durationExpression"`
	// `CfnDetectorModel.SetTimerProperty.Seconds`.
	Seconds *float64 `json:"seconds"`
}

type CfnDetectorModel_SetVariableProperty struct {
	// `CfnDetectorModel.SetVariableProperty.Value`.
	Value *string `json:"value"`
	// `CfnDetectorModel.SetVariableProperty.VariableName`.
	VariableName *string `json:"variableName"`
}

type CfnDetectorModel_SnsProperty struct {
	// `CfnDetectorModel.SnsProperty.TargetArn`.
	TargetArn *string `json:"targetArn"`
	// `CfnDetectorModel.SnsProperty.Payload`.
	Payload interface{} `json:"payload"`
}

type CfnDetectorModel_SqsProperty struct {
	// `CfnDetectorModel.SqsProperty.QueueUrl`.
	QueueUrl *string `json:"queueUrl"`
	// `CfnDetectorModel.SqsProperty.Payload`.
	Payload interface{} `json:"payload"`
	// `CfnDetectorModel.SqsProperty.UseBase64`.
	UseBase64 interface{} `json:"useBase64"`
}

type CfnDetectorModel_StateProperty struct {
	// `CfnDetectorModel.StateProperty.StateName`.
	StateName *string `json:"stateName"`
	// `CfnDetectorModel.StateProperty.OnEnter`.
	OnEnter interface{} `json:"onEnter"`
	// `CfnDetectorModel.StateProperty.OnExit`.
	OnExit interface{} `json:"onExit"`
	// `CfnDetectorModel.StateProperty.OnInput`.
	OnInput interface{} `json:"onInput"`
}

type CfnDetectorModel_TransitionEventProperty struct {
	// `CfnDetectorModel.TransitionEventProperty.Condition`.
	Condition *string `json:"condition"`
	// `CfnDetectorModel.TransitionEventProperty.EventName`.
	EventName *string `json:"eventName"`
	// `CfnDetectorModel.TransitionEventProperty.NextState`.
	NextState *string `json:"nextState"`
	// `CfnDetectorModel.TransitionEventProperty.Actions`.
	Actions interface{} `json:"actions"`
}

// Properties for defining a `AWS::IoTEvents::DetectorModel`.
type CfnDetectorModelProps struct {
	// `AWS::IoTEvents::DetectorModel.DetectorModelDefinition`.
	DetectorModelDefinition interface{} `json:"detectorModelDefinition"`
	// `AWS::IoTEvents::DetectorModel.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `AWS::IoTEvents::DetectorModel.DetectorModelDescription`.
	DetectorModelDescription *string `json:"detectorModelDescription"`
	// `AWS::IoTEvents::DetectorModel.DetectorModelName`.
	DetectorModelName *string `json:"detectorModelName"`
	// `AWS::IoTEvents::DetectorModel.EvaluationMethod`.
	EvaluationMethod *string `json:"evaluationMethod"`
	// `AWS::IoTEvents::DetectorModel.Key`.
	Key *string `json:"key"`
	// `AWS::IoTEvents::DetectorModel.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::IoTEvents::Input`.
type CfnInput interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	InputDefinition() interface{}
	SetInputDefinition(val interface{})
	InputDescription() *string
	SetInputDescription(val *string)
	InputName() *string
	SetInputName(val *string)
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

// The jsii proxy struct for CfnInput
type jsiiProxy_CfnInput struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnInput) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) InputDefinition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) InputDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) InputName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTEvents::Input`.
func NewCfnInput(scope awscdk.Construct, id *string, props *CfnInputProps) CfnInput {
	_init_.Initialize()

	j := jsiiProxy_CfnInput{}

	_jsii_.Create(
		"monocdk.aws_iotevents.CfnInput",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTEvents::Input`.
func NewCfnInput_Override(c CfnInput, scope awscdk.Construct, id *string, props *CfnInputProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iotevents.CfnInput",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnInput) SetInputDefinition(val interface{}) {
	_jsii_.Set(
		j,
		"inputDefinition",
		val,
	)
}

func (j *jsiiProxy_CfnInput) SetInputDescription(val *string) {
	_jsii_.Set(
		j,
		"inputDescription",
		val,
	)
}

func (j *jsiiProxy_CfnInput) SetInputName(val *string) {
	_jsii_.Set(
		j,
		"inputName",
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
func CfnInput_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotevents.CfnInput",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnInput_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotevents.CfnInput",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnInput_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotevents.CfnInput",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInput_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_iotevents.CfnInput",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnInput) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnInput) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnInput) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnInput) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnInput) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnInput) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnInput) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnInput) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnInput) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnInput) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnInput) OnPrepare() {
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
func (c *jsiiProxy_CfnInput) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnInput) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnInput) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnInput) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnInput) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnInput) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnInput) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnInput) ToString() *string {
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
func (c *jsiiProxy_CfnInput) Validate() *[]*string {
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
func (c *jsiiProxy_CfnInput) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnInput_AttributeProperty struct {
	// `CfnInput.AttributeProperty.JsonPath`.
	JsonPath *string `json:"jsonPath"`
}

type CfnInput_InputDefinitionProperty struct {
	// `CfnInput.InputDefinitionProperty.Attributes`.
	Attributes interface{} `json:"attributes"`
}

// Properties for defining a `AWS::IoTEvents::Input`.
type CfnInputProps struct {
	// `AWS::IoTEvents::Input.InputDefinition`.
	InputDefinition interface{} `json:"inputDefinition"`
	// `AWS::IoTEvents::Input.InputDescription`.
	InputDescription *string `json:"inputDescription"`
	// `AWS::IoTEvents::Input.InputName`.
	InputName *string `json:"inputName"`
	// `AWS::IoTEvents::Input.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}


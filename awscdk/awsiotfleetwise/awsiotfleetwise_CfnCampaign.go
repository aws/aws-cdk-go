package awsiotfleetwise

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiotfleetwise/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::IoTFleetWise::Campaign`.
//
// Creates an orchestration of data collection rules. The AWS IoT FleetWise Edge Agent software running in vehicles uses campaigns to decide how to collect and transfer data to the cloud. You create campaigns in the cloud. After you or your team approve campaigns, AWS IoT FleetWise automatically deploys them to vehicles.
//
// For more information, see [Collect and transfer data with campaigns](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/campaigns.html) in the *AWS IoT FleetWise Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCampaign := awscdk.Aws_iotfleetwise.NewCfnCampaign(this, jsii.String("MyCfnCampaign"), &CfnCampaignProps{
//   	Action: jsii.String("action"),
//   	CollectionScheme: &CollectionSchemeProperty{
//   		ConditionBasedCollectionScheme: &ConditionBasedCollectionSchemeProperty{
//   			Expression: jsii.String("expression"),
//
//   			// the properties below are optional
//   			ConditionLanguageVersion: jsii.Number(123),
//   			MinimumTriggerIntervalMs: jsii.Number(123),
//   			TriggerMode: jsii.String("triggerMode"),
//   		},
//   		TimeBasedCollectionScheme: &TimeBasedCollectionSchemeProperty{
//   			PeriodMs: jsii.Number(123),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	SignalCatalogArn: jsii.String("signalCatalogArn"),
//   	TargetArn: jsii.String("targetArn"),
//
//   	// the properties below are optional
//   	Compression: jsii.String("compression"),
//   	DataExtraDimensions: []*string{
//   		jsii.String("dataExtraDimensions"),
//   	},
//   	Description: jsii.String("description"),
//   	DiagnosticsMode: jsii.String("diagnosticsMode"),
//   	ExpiryTime: jsii.String("expiryTime"),
//   	PostTriggerCollectionDuration: jsii.Number(123),
//   	Priority: jsii.Number(123),
//   	SignalsToCollect: []interface{}{
//   		&SignalInformationProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			MaxSampleCount: jsii.Number(123),
//   			MinimumSamplingIntervalMs: jsii.Number(123),
//   		},
//   	},
//   	SpoolingMode: jsii.String("spoolingMode"),
//   	StartTime: jsii.String("startTime"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnCampaign interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// `AWS::IoTFleetWise::Campaign.Action`.
	Action() *string
	SetAction(val *string)
	AttrArn() *string
	AttrCreationTime() *string
	AttrLastModificationTime() *string
	AttrStatus() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// `AWS::IoTFleetWise::Campaign.CollectionScheme`.
	CollectionScheme() interface{}
	SetCollectionScheme(val interface{})
	// `AWS::IoTFleetWise::Campaign.Compression`.
	Compression() *string
	SetCompression(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// `AWS::IoTFleetWise::Campaign.DataExtraDimensions`.
	DataExtraDimensions() *[]*string
	SetDataExtraDimensions(val *[]*string)
	// The description of the campaign.
	Description() *string
	SetDescription(val *string)
	// `AWS::IoTFleetWise::Campaign.DiagnosticsMode`.
	DiagnosticsMode() *string
	SetDiagnosticsMode(val *string)
	// `AWS::IoTFleetWise::Campaign.ExpiryTime`.
	ExpiryTime() *string
	SetExpiryTime(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The name of a campaign.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// `AWS::IoTFleetWise::Campaign.PostTriggerCollectionDuration`.
	PostTriggerCollectionDuration() *float64
	SetPostTriggerCollectionDuration(val *float64)
	// `AWS::IoTFleetWise::Campaign.Priority`.
	Priority() *float64
	SetPriority(val *float64)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The ARN of the signal catalog associated with the campaign.
	SignalCatalogArn() *string
	SetSignalCatalogArn(val *string)
	// `AWS::IoTFleetWise::Campaign.SignalsToCollect`.
	SignalsToCollect() interface{}
	SetSignalsToCollect(val interface{})
	// `AWS::IoTFleetWise::Campaign.SpoolingMode`.
	SpoolingMode() *string
	SetSpoolingMode(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::IoTFleetWise::Campaign.StartTime`.
	StartTime() *string
	SetStartTime(val *string)
	// `AWS::IoTFleetWise::Campaign.Tags`.
	Tags() awscdk.TagManager
	// The ARN of a vehicle or fleet to which the campaign is deployed.
	TargetArn() *string
	SetTargetArn(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnCampaign
type jsiiProxy_CfnCampaign struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCampaign) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) AttrCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) AttrLastModificationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastModificationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) CollectionScheme() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"collectionScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Compression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) DataExtraDimensions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataExtraDimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) DiagnosticsMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diagnosticsMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) ExpiryTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiryTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) PostTriggerCollectionDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"postTriggerCollectionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) SignalCatalogArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signalCatalogArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) SignalsToCollect() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signalsToCollect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) SpoolingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spoolingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTFleetWise::Campaign`.
func NewCfnCampaign(scope awscdk.Construct, id *string, props *CfnCampaignProps) CfnCampaign {
	_init_.Initialize()

	if err := validateNewCfnCampaignParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCampaign{}

	_jsii_.Create(
		"monocdk.aws_iotfleetwise.CfnCampaign",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTFleetWise::Campaign`.
func NewCfnCampaign_Override(c CfnCampaign, scope awscdk.Construct, id *string, props *CfnCampaignProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iotfleetwise.CfnCampaign",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCampaign)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetCollectionScheme(val interface{}) {
	if err := j.validateSetCollectionSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collectionScheme",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetCompression(val *string) {
	_jsii_.Set(
		j,
		"compression",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetDataExtraDimensions(val *[]*string) {
	_jsii_.Set(
		j,
		"dataExtraDimensions",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetDiagnosticsMode(val *string) {
	_jsii_.Set(
		j,
		"diagnosticsMode",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetExpiryTime(val *string) {
	_jsii_.Set(
		j,
		"expiryTime",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetPostTriggerCollectionDuration(val *float64) {
	_jsii_.Set(
		j,
		"postTriggerCollectionDuration",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetPriority(val *float64) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetSignalCatalogArn(val *string) {
	if err := j.validateSetSignalCatalogArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signalCatalogArn",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetSignalsToCollect(val interface{}) {
	if err := j.validateSetSignalsToCollectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signalsToCollect",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetSpoolingMode(val *string) {
	_jsii_.Set(
		j,
		"spoolingMode",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetStartTime(val *string) {
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetTargetArn(val *string) {
	if err := j.validateSetTargetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetArn",
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
func CfnCampaign_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCampaign_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotfleetwise.CfnCampaign",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnCampaign_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnCampaign_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotfleetwise.CfnCampaign",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnCampaign_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCampaign_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iotfleetwise.CfnCampaign",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCampaign_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_iotfleetwise.CfnCampaign",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCampaign) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCampaign) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCampaign) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCampaign) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCampaign) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCampaign) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCampaign) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCampaign) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCampaign) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnCampaign) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCampaign) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnCampaign) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnCampaign) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCampaign) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCampaign) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnCampaign) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnCampaign) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCampaign) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnCampaign) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCampaign) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCampaign) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}


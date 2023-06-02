package awsiotfleetwise

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotfleetwise/internal"
	"github.com/aws/constructs-go/constructs/v10"
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
//   	DataDestinationConfigs: []interface{}{
//   		&DataDestinationConfigProperty{
//   			S3Config: &S3ConfigProperty{
//   				BucketArn: jsii.String("bucketArn"),
//
//   				// the properties below are optional
//   				DataFormat: jsii.String("dataFormat"),
//   				Prefix: jsii.String("prefix"),
//   				StorageCompressionFormat: jsii.String("storageCompressionFormat"),
//   			},
//   			TimestreamConfig: &TimestreamConfigProperty{
//   				ExecutionRoleArn: jsii.String("executionRoleArn"),
//   				TimestreamTableArn: jsii.String("timestreamTableArn"),
//   			},
//   		},
//   	},
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
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
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
	CreationStack() *[]*string
	// `AWS::IoTFleetWise::Campaign.DataDestinationConfigs`.
	DataDestinationConfigs() interface{}
	SetDataDestinationConfigs(val interface{})
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
	LogicalId() *string
	// The name of a campaign.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
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
	Stack() awscdk.Stack
	// `AWS::IoTFleetWise::Campaign.StartTime`.
	StartTime() *string
	SetStartTime(val *string)
	// `AWS::IoTFleetWise::Campaign.Tags`.
	Tags() awscdk.TagManager
	// The ARN of a vehicle or fleet to which the campaign is deployed.
	TargetArn() *string
	SetTargetArn(val *string)
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

func (j *jsiiProxy_CfnCampaign) DataDestinationConfigs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataDestinationConfigs",
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

func (j *jsiiProxy_CfnCampaign) Node() constructs.Node {
	var returns constructs.Node
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

func (j *jsiiProxy_CfnCampaign) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTFleetWise::Campaign`.
func NewCfnCampaign(scope constructs.Construct, id *string, props *CfnCampaignProps) CfnCampaign {
	_init_.Initialize()

	if err := validateNewCfnCampaignParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCampaign{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iotfleetwise.CfnCampaign",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTFleetWise::Campaign`.
func NewCfnCampaign_Override(c CfnCampaign, scope constructs.Construct, id *string, props *CfnCampaignProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iotfleetwise.CfnCampaign",
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

func (j *jsiiProxy_CfnCampaign)SetDataDestinationConfigs(val interface{}) {
	if err := j.validateSetDataDestinationConfigsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataDestinationConfigs",
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
func CfnCampaign_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCampaign_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotfleetwise.CfnCampaign",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCampaign_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnCampaign_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotfleetwise.CfnCampaign",
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
func CfnCampaign_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCampaign_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotfleetwise.CfnCampaign",
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
		"aws-cdk-lib.aws_iotfleetwise.CfnCampaign",
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

func (c *jsiiProxy_CfnCampaign) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
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

func (c *jsiiProxy_CfnCampaign) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnCampaign) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCampaign) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
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

func (c *jsiiProxy_CfnCampaign) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
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

func (c *jsiiProxy_CfnCampaign) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
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


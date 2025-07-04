package awsconnectcampaignsv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnectcampaignsv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an outbound campaign.
//
// > - For users to be able to view or edit a campaign at a later date by using the Amazon Connect user interface, you must add the instance ID as a tag. For example, `{ "tags": {"owner": "arn:aws:connect:{REGION}:{AWS_ACCOUNT_ID}:instance/{CONNECT_INSTANCE_ID}"}}` .
// > - After a campaign is created, you can't add/remove source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var agentlessConfig interface{}
//
//   cfnCampaign := awscdk.Aws_connectcampaignsv2.NewCfnCampaign(this, jsii.String("MyCfnCampaign"), &CfnCampaignProps{
//   	ChannelSubtypeConfig: &ChannelSubtypeConfigProperty{
//   		Email: &EmailChannelSubtypeConfigProperty{
//   			DefaultOutboundConfig: &EmailOutboundConfigProperty{
//   				ConnectSourceEmailAddress: jsii.String("connectSourceEmailAddress"),
//   				WisdomTemplateArn: jsii.String("wisdomTemplateArn"),
//
//   				// the properties below are optional
//   				SourceEmailAddressDisplayName: jsii.String("sourceEmailAddressDisplayName"),
//   			},
//   			OutboundMode: &EmailOutboundModeProperty{
//   				AgentlessConfig: agentlessConfig,
//   			},
//
//   			// the properties below are optional
//   			Capacity: jsii.Number(123),
//   		},
//   		Sms: &SmsChannelSubtypeConfigProperty{
//   			DefaultOutboundConfig: &SmsOutboundConfigProperty{
//   				ConnectSourcePhoneNumberArn: jsii.String("connectSourcePhoneNumberArn"),
//   				WisdomTemplateArn: jsii.String("wisdomTemplateArn"),
//   			},
//   			OutboundMode: &SmsOutboundModeProperty{
//   				AgentlessConfig: agentlessConfig,
//   			},
//
//   			// the properties below are optional
//   			Capacity: jsii.Number(123),
//   		},
//   		Telephony: &TelephonyChannelSubtypeConfigProperty{
//   			DefaultOutboundConfig: &TelephonyOutboundConfigProperty{
//   				ConnectContactFlowId: jsii.String("connectContactFlowId"),
//
//   				// the properties below are optional
//   				AnswerMachineDetectionConfig: &AnswerMachineDetectionConfigProperty{
//   					EnableAnswerMachineDetection: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					AwaitAnswerMachinePrompt: jsii.Boolean(false),
//   				},
//   				ConnectSourcePhoneNumber: jsii.String("connectSourcePhoneNumber"),
//   			},
//   			OutboundMode: &TelephonyOutboundModeProperty{
//   				AgentlessConfig: agentlessConfig,
//   				PredictiveConfig: &PredictiveConfigProperty{
//   					BandwidthAllocation: jsii.Number(123),
//   				},
//   				ProgressiveConfig: &ProgressiveConfigProperty{
//   					BandwidthAllocation: jsii.Number(123),
//   				},
//   			},
//
//   			// the properties below are optional
//   			Capacity: jsii.Number(123),
//   			ConnectQueueId: jsii.String("connectQueueId"),
//   		},
//   	},
//   	ConnectInstanceId: jsii.String("connectInstanceId"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	CommunicationLimitsOverride: &CommunicationLimitsConfigProperty{
//   		AllChannelsSubtypes: &CommunicationLimitsProperty{
//   			CommunicationLimitList: []interface{}{
//   				&CommunicationLimitProperty{
//   					Frequency: jsii.Number(123),
//   					MaxCountPerRecipient: jsii.Number(123),
//   					Unit: jsii.String("unit"),
//   				},
//   			},
//   		},
//   		InstanceLimitsHandling: jsii.String("instanceLimitsHandling"),
//   	},
//   	CommunicationTimeConfig: &CommunicationTimeConfigProperty{
//   		LocalTimeZoneConfig: &LocalTimeZoneConfigProperty{
//   			DefaultTimeZone: jsii.String("defaultTimeZone"),
//   			LocalTimeZoneDetection: []*string{
//   				jsii.String("localTimeZoneDetection"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		Email: &TimeWindowProperty{
//   			OpenHours: &OpenHoursProperty{
//   				DailyHours: []interface{}{
//   					&DailyHourProperty{
//   						Key: jsii.String("key"),
//   						Value: []interface{}{
//   							&TimeRangeProperty{
//   								EndTime: jsii.String("endTime"),
//   								StartTime: jsii.String("startTime"),
//   							},
//   						},
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			RestrictedPeriods: &RestrictedPeriodsProperty{
//   				RestrictedPeriodList: []interface{}{
//   					&RestrictedPeriodProperty{
//   						EndDate: jsii.String("endDate"),
//   						StartDate: jsii.String("startDate"),
//
//   						// the properties below are optional
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   		},
//   		Sms: &TimeWindowProperty{
//   			OpenHours: &OpenHoursProperty{
//   				DailyHours: []interface{}{
//   					&DailyHourProperty{
//   						Key: jsii.String("key"),
//   						Value: []interface{}{
//   							&TimeRangeProperty{
//   								EndTime: jsii.String("endTime"),
//   								StartTime: jsii.String("startTime"),
//   							},
//   						},
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			RestrictedPeriods: &RestrictedPeriodsProperty{
//   				RestrictedPeriodList: []interface{}{
//   					&RestrictedPeriodProperty{
//   						EndDate: jsii.String("endDate"),
//   						StartDate: jsii.String("startDate"),
//
//   						// the properties below are optional
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   		},
//   		Telephony: &TimeWindowProperty{
//   			OpenHours: &OpenHoursProperty{
//   				DailyHours: []interface{}{
//   					&DailyHourProperty{
//   						Key: jsii.String("key"),
//   						Value: []interface{}{
//   							&TimeRangeProperty{
//   								EndTime: jsii.String("endTime"),
//   								StartTime: jsii.String("startTime"),
//   							},
//   						},
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			RestrictedPeriods: &RestrictedPeriodsProperty{
//   				RestrictedPeriodList: []interface{}{
//   					&RestrictedPeriodProperty{
//   						EndDate: jsii.String("endDate"),
//   						StartDate: jsii.String("startDate"),
//
//   						// the properties below are optional
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	ConnectCampaignFlowArn: jsii.String("connectCampaignFlowArn"),
//   	Schedule: &ScheduleProperty{
//   		EndTime: jsii.String("endTime"),
//   		StartTime: jsii.String("startTime"),
//
//   		// the properties below are optional
//   		RefreshFrequency: jsii.String("refreshFrequency"),
//   	},
//   	Source: &SourceProperty{
//   		CustomerProfilesSegmentArn: jsii.String("customerProfilesSegmentArn"),
//   		EventTrigger: &EventTriggerProperty{
//   			CustomerProfilesDomainArn: jsii.String("customerProfilesDomainArn"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaignsv2-campaign.html
//
type CfnCampaign interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggableV2
	// The Amazon Resource Name (ARN).
	AttrArn() *string
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Contains channel subtype configuration for an outbound campaign.
	ChannelSubtypeConfig() interface{}
	SetChannelSubtypeConfig(val interface{})
	// Communication limits configuration for an outbound campaign.
	CommunicationLimitsOverride() interface{}
	SetCommunicationLimitsOverride(val interface{})
	// Contains communication time configuration for an outbound campaign.
	CommunicationTimeConfig() interface{}
	SetCommunicationTimeConfig(val interface{})
	// The Amazon Resource Name (ARN) of the Amazon Connect campaign flow associated with the outbound campaign.
	ConnectCampaignFlowArn() *string
	SetConnectCampaignFlowArn(val *string)
	// The identifier of the Amazon Connect instance.
	ConnectInstanceId() *string
	SetConnectInstanceId(val *string)
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
	// The name of the outbound campaign.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Contains the schedule configuration.
	Schedule() interface{}
	SetSchedule(val interface{})
	// Contains source configuration.
	Source() interface{}
	SetSource(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags used to organize, track, or control access for this resource.
	Tags() *[]*awscdk.CfnTag
	SetTags(val *[]*awscdk.CfnTag)
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
	internal.Type__awscdkITaggableV2
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

func (j *jsiiProxy_CfnCampaign) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
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

func (j *jsiiProxy_CfnCampaign) ChannelSubtypeConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"channelSubtypeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) CommunicationLimitsOverride() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"communicationLimitsOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) CommunicationTimeConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"communicationTimeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) ConnectCampaignFlowArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectCampaignFlowArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) ConnectInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectInstanceId",
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

func (j *jsiiProxy_CfnCampaign) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Schedule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Source() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"source",
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

func (j *jsiiProxy_CfnCampaign) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
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


func NewCfnCampaign(scope constructs.Construct, id *string, props *CfnCampaignProps) CfnCampaign {
	_init_.Initialize()

	if err := validateNewCfnCampaignParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCampaign{}

	_jsii_.Create(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnCampaign_Override(c CfnCampaign, scope constructs.Construct, id *string, props *CfnCampaignProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCampaign)SetChannelSubtypeConfig(val interface{}) {
	if err := j.validateSetChannelSubtypeConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelSubtypeConfig",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetCommunicationLimitsOverride(val interface{}) {
	if err := j.validateSetCommunicationLimitsOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"communicationLimitsOverride",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetCommunicationTimeConfig(val interface{}) {
	if err := j.validateSetCommunicationTimeConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"communicationTimeConfig",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetConnectCampaignFlowArn(val *string) {
	_jsii_.Set(
		j,
		"connectCampaignFlowArn",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetConnectInstanceId(val *string) {
	if err := j.validateSetConnectInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectInstanceId",
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

func (j *jsiiProxy_CfnCampaign)SetSchedule(val interface{}) {
	if err := j.validateSetScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetSource(val interface{}) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
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
func CfnCampaign_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCampaign_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnCampaign_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCampaign_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign",
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
func CfnCampaign_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCampaign_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign",
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
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign",
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


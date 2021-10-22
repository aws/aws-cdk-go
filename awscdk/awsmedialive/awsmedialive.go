package awsmedialive

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::MediaLive::Channel`.
type CfnChannel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrInputs() *[]*string
	CdiInputSpecification() interface{}
	SetCdiInputSpecification(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ChannelClass() *string
	SetChannelClass(val *string)
	CreationStack() *[]*string
	Destinations() interface{}
	SetDestinations(val interface{})
	EncoderSettings() interface{}
	SetEncoderSettings(val interface{})
	InputAttachments() interface{}
	SetInputAttachments(val interface{})
	InputSpecification() interface{}
	SetInputSpecification(val interface{})
	LogicalId() *string
	LogLevel() *string
	SetLogLevel(val *string)
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	RoleArn() *string
	SetRoleArn(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	Vpc() interface{}
	SetVpc(val interface{})
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

// The jsii proxy struct for CfnChannel
type jsiiProxy_CfnChannel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnChannel) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) AttrInputs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrInputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) CdiInputSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdiInputSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) ChannelClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Destinations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) EncoderSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encoderSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) InputAttachments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputAttachments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) InputSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannel) Vpc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaLive::Channel`.
func NewCfnChannel(scope constructs.Construct, id *string, props *CfnChannelProps) CfnChannel {
	_init_.Initialize()

	j := jsiiProxy_CfnChannel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_medialive.CfnChannel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaLive::Channel`.
func NewCfnChannel_Override(c CfnChannel, scope constructs.Construct, id *string, props *CfnChannelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_medialive.CfnChannel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnChannel) SetCdiInputSpecification(val interface{}) {
	_jsii_.Set(
		j,
		"cdiInputSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetChannelClass(val *string) {
	_jsii_.Set(
		j,
		"channelClass",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetDestinations(val interface{}) {
	_jsii_.Set(
		j,
		"destinations",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetEncoderSettings(val interface{}) {
	_jsii_.Set(
		j,
		"encoderSettings",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetInputAttachments(val interface{}) {
	_jsii_.Set(
		j,
		"inputAttachments",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetInputSpecification(val interface{}) {
	_jsii_.Set(
		j,
		"inputSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetLogLevel(val *string) {
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnChannel) SetVpc(val interface{}) {
	_jsii_.Set(
		j,
		"vpc",
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
func CfnChannel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_medialive.CfnChannel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnChannel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_medialive.CfnChannel",
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
func CfnChannel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_medialive.CfnChannel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnChannel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_medialive.CfnChannel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnChannel) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnChannel) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnChannel) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnChannel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnChannel) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnChannel) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnChannel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnChannel) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnChannel) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnChannel) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnChannel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnChannel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnChannel) ShouldSynthesize() *bool {
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
// Experimental.
func (c *jsiiProxy_CfnChannel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnChannel) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnChannel_AacSettingsProperty struct {
	// `CfnChannel.AacSettingsProperty.Bitrate`.
	Bitrate *float64 `json:"bitrate"`
	// `CfnChannel.AacSettingsProperty.CodingMode`.
	CodingMode *string `json:"codingMode"`
	// `CfnChannel.AacSettingsProperty.InputType`.
	InputType *string `json:"inputType"`
	// `CfnChannel.AacSettingsProperty.Profile`.
	Profile *string `json:"profile"`
	// `CfnChannel.AacSettingsProperty.RateControlMode`.
	RateControlMode *string `json:"rateControlMode"`
	// `CfnChannel.AacSettingsProperty.RawFormat`.
	RawFormat *string `json:"rawFormat"`
	// `CfnChannel.AacSettingsProperty.SampleRate`.
	SampleRate *float64 `json:"sampleRate"`
	// `CfnChannel.AacSettingsProperty.Spec`.
	Spec *string `json:"spec"`
	// `CfnChannel.AacSettingsProperty.VbrQuality`.
	VbrQuality *string `json:"vbrQuality"`
}

type CfnChannel_Ac3SettingsProperty struct {
	// `CfnChannel.Ac3SettingsProperty.Bitrate`.
	Bitrate *float64 `json:"bitrate"`
	// `CfnChannel.Ac3SettingsProperty.BitstreamMode`.
	BitstreamMode *string `json:"bitstreamMode"`
	// `CfnChannel.Ac3SettingsProperty.CodingMode`.
	CodingMode *string `json:"codingMode"`
	// `CfnChannel.Ac3SettingsProperty.Dialnorm`.
	Dialnorm *float64 `json:"dialnorm"`
	// `CfnChannel.Ac3SettingsProperty.DrcProfile`.
	DrcProfile *string `json:"drcProfile"`
	// `CfnChannel.Ac3SettingsProperty.LfeFilter`.
	LfeFilter *string `json:"lfeFilter"`
	// `CfnChannel.Ac3SettingsProperty.MetadataControl`.
	MetadataControl *string `json:"metadataControl"`
}

type CfnChannel_AncillarySourceSettingsProperty struct {
	// `CfnChannel.AncillarySourceSettingsProperty.SourceAncillaryChannelNumber`.
	SourceAncillaryChannelNumber *float64 `json:"sourceAncillaryChannelNumber"`
}

type CfnChannel_ArchiveCdnSettingsProperty struct {
	// `CfnChannel.ArchiveCdnSettingsProperty.ArchiveS3Settings`.
	ArchiveS3Settings interface{} `json:"archiveS3Settings"`
}

type CfnChannel_ArchiveContainerSettingsProperty struct {
	// `CfnChannel.ArchiveContainerSettingsProperty.M2tsSettings`.
	M2TsSettings interface{} `json:"m2TsSettings"`
	// `CfnChannel.ArchiveContainerSettingsProperty.RawSettings`.
	RawSettings interface{} `json:"rawSettings"`
}

type CfnChannel_ArchiveGroupSettingsProperty struct {
	// `CfnChannel.ArchiveGroupSettingsProperty.ArchiveCdnSettings`.
	ArchiveCdnSettings interface{} `json:"archiveCdnSettings"`
	// `CfnChannel.ArchiveGroupSettingsProperty.Destination`.
	Destination interface{} `json:"destination"`
	// `CfnChannel.ArchiveGroupSettingsProperty.RolloverInterval`.
	RolloverInterval *float64 `json:"rolloverInterval"`
}

type CfnChannel_ArchiveOutputSettingsProperty struct {
	// `CfnChannel.ArchiveOutputSettingsProperty.ContainerSettings`.
	ContainerSettings interface{} `json:"containerSettings"`
	// `CfnChannel.ArchiveOutputSettingsProperty.Extension`.
	Extension *string `json:"extension"`
	// `CfnChannel.ArchiveOutputSettingsProperty.NameModifier`.
	NameModifier *string `json:"nameModifier"`
}

type CfnChannel_ArchiveS3SettingsProperty struct {
	// `CfnChannel.ArchiveS3SettingsProperty.CannedAcl`.
	CannedAcl *string `json:"cannedAcl"`
}

type CfnChannel_AribDestinationSettingsProperty struct {
}

type CfnChannel_AribSourceSettingsProperty struct {
}

type CfnChannel_AudioChannelMappingProperty struct {
	// `CfnChannel.AudioChannelMappingProperty.InputChannelLevels`.
	InputChannelLevels interface{} `json:"inputChannelLevels"`
	// `CfnChannel.AudioChannelMappingProperty.OutputChannel`.
	OutputChannel *float64 `json:"outputChannel"`
}

type CfnChannel_AudioCodecSettingsProperty struct {
	// `CfnChannel.AudioCodecSettingsProperty.AacSettings`.
	AacSettings interface{} `json:"aacSettings"`
	// `CfnChannel.AudioCodecSettingsProperty.Ac3Settings`.
	Ac3Settings interface{} `json:"ac3Settings"`
	// `CfnChannel.AudioCodecSettingsProperty.Eac3Settings`.
	Eac3Settings interface{} `json:"eac3Settings"`
	// `CfnChannel.AudioCodecSettingsProperty.Mp2Settings`.
	Mp2Settings interface{} `json:"mp2Settings"`
	// `CfnChannel.AudioCodecSettingsProperty.PassThroughSettings`.
	PassThroughSettings interface{} `json:"passThroughSettings"`
	// `CfnChannel.AudioCodecSettingsProperty.WavSettings`.
	WavSettings interface{} `json:"wavSettings"`
}

type CfnChannel_AudioDescriptionProperty struct {
	// `CfnChannel.AudioDescriptionProperty.AudioNormalizationSettings`.
	AudioNormalizationSettings interface{} `json:"audioNormalizationSettings"`
	// `CfnChannel.AudioDescriptionProperty.AudioSelectorName`.
	AudioSelectorName *string `json:"audioSelectorName"`
	// `CfnChannel.AudioDescriptionProperty.AudioType`.
	AudioType *string `json:"audioType"`
	// `CfnChannel.AudioDescriptionProperty.AudioTypeControl`.
	AudioTypeControl *string `json:"audioTypeControl"`
	// `CfnChannel.AudioDescriptionProperty.CodecSettings`.
	CodecSettings interface{} `json:"codecSettings"`
	// `CfnChannel.AudioDescriptionProperty.LanguageCode`.
	LanguageCode *string `json:"languageCode"`
	// `CfnChannel.AudioDescriptionProperty.LanguageCodeControl`.
	LanguageCodeControl *string `json:"languageCodeControl"`
	// `CfnChannel.AudioDescriptionProperty.Name`.
	Name *string `json:"name"`
	// `CfnChannel.AudioDescriptionProperty.RemixSettings`.
	RemixSettings interface{} `json:"remixSettings"`
	// `CfnChannel.AudioDescriptionProperty.StreamName`.
	StreamName *string `json:"streamName"`
}

type CfnChannel_AudioLanguageSelectionProperty struct {
	// `CfnChannel.AudioLanguageSelectionProperty.LanguageCode`.
	LanguageCode *string `json:"languageCode"`
	// `CfnChannel.AudioLanguageSelectionProperty.LanguageSelectionPolicy`.
	LanguageSelectionPolicy *string `json:"languageSelectionPolicy"`
}

type CfnChannel_AudioNormalizationSettingsProperty struct {
	// `CfnChannel.AudioNormalizationSettingsProperty.Algorithm`.
	Algorithm *string `json:"algorithm"`
	// `CfnChannel.AudioNormalizationSettingsProperty.AlgorithmControl`.
	AlgorithmControl *string `json:"algorithmControl"`
	// `CfnChannel.AudioNormalizationSettingsProperty.TargetLkfs`.
	TargetLkfs *float64 `json:"targetLkfs"`
}

type CfnChannel_AudioOnlyHlsSettingsProperty struct {
	// `CfnChannel.AudioOnlyHlsSettingsProperty.AudioGroupId`.
	AudioGroupId *string `json:"audioGroupId"`
	// `CfnChannel.AudioOnlyHlsSettingsProperty.AudioOnlyImage`.
	AudioOnlyImage interface{} `json:"audioOnlyImage"`
	// `CfnChannel.AudioOnlyHlsSettingsProperty.AudioTrackType`.
	AudioTrackType *string `json:"audioTrackType"`
	// `CfnChannel.AudioOnlyHlsSettingsProperty.SegmentType`.
	SegmentType *string `json:"segmentType"`
}

type CfnChannel_AudioPidSelectionProperty struct {
	// `CfnChannel.AudioPidSelectionProperty.Pid`.
	Pid *float64 `json:"pid"`
}

type CfnChannel_AudioSelectorProperty struct {
	// `CfnChannel.AudioSelectorProperty.Name`.
	Name *string `json:"name"`
	// `CfnChannel.AudioSelectorProperty.SelectorSettings`.
	SelectorSettings interface{} `json:"selectorSettings"`
}

type CfnChannel_AudioSelectorSettingsProperty struct {
	// `CfnChannel.AudioSelectorSettingsProperty.AudioLanguageSelection`.
	AudioLanguageSelection interface{} `json:"audioLanguageSelection"`
	// `CfnChannel.AudioSelectorSettingsProperty.AudioPidSelection`.
	AudioPidSelection interface{} `json:"audioPidSelection"`
	// `CfnChannel.AudioSelectorSettingsProperty.AudioTrackSelection`.
	AudioTrackSelection interface{} `json:"audioTrackSelection"`
}

type CfnChannel_AudioSilenceFailoverSettingsProperty struct {
	// `CfnChannel.AudioSilenceFailoverSettingsProperty.AudioSelectorName`.
	AudioSelectorName *string `json:"audioSelectorName"`
	// `CfnChannel.AudioSilenceFailoverSettingsProperty.AudioSilenceThresholdMsec`.
	AudioSilenceThresholdMsec *float64 `json:"audioSilenceThresholdMsec"`
}

type CfnChannel_AudioTrackProperty struct {
	// `CfnChannel.AudioTrackProperty.Track`.
	Track *float64 `json:"track"`
}

type CfnChannel_AudioTrackSelectionProperty struct {
	// `CfnChannel.AudioTrackSelectionProperty.Tracks`.
	Tracks interface{} `json:"tracks"`
}

type CfnChannel_AutomaticInputFailoverSettingsProperty struct {
	// `CfnChannel.AutomaticInputFailoverSettingsProperty.ErrorClearTimeMsec`.
	ErrorClearTimeMsec *float64 `json:"errorClearTimeMsec"`
	// `CfnChannel.AutomaticInputFailoverSettingsProperty.FailoverConditions`.
	FailoverConditions interface{} `json:"failoverConditions"`
	// `CfnChannel.AutomaticInputFailoverSettingsProperty.InputPreference`.
	InputPreference *string `json:"inputPreference"`
	// `CfnChannel.AutomaticInputFailoverSettingsProperty.SecondaryInputId`.
	SecondaryInputId *string `json:"secondaryInputId"`
}

type CfnChannel_AvailBlankingProperty struct {
	// `CfnChannel.AvailBlankingProperty.AvailBlankingImage`.
	AvailBlankingImage interface{} `json:"availBlankingImage"`
	// `CfnChannel.AvailBlankingProperty.State`.
	State *string `json:"state"`
}

type CfnChannel_AvailConfigurationProperty struct {
	// `CfnChannel.AvailConfigurationProperty.AvailSettings`.
	AvailSettings interface{} `json:"availSettings"`
}

type CfnChannel_AvailSettingsProperty struct {
	// `CfnChannel.AvailSettingsProperty.Scte35SpliceInsert`.
	Scte35SpliceInsert interface{} `json:"scte35SpliceInsert"`
	// `CfnChannel.AvailSettingsProperty.Scte35TimeSignalApos`.
	Scte35TimeSignalApos interface{} `json:"scte35TimeSignalApos"`
}

type CfnChannel_BlackoutSlateProperty struct {
	// `CfnChannel.BlackoutSlateProperty.BlackoutSlateImage`.
	BlackoutSlateImage interface{} `json:"blackoutSlateImage"`
	// `CfnChannel.BlackoutSlateProperty.NetworkEndBlackout`.
	NetworkEndBlackout *string `json:"networkEndBlackout"`
	// `CfnChannel.BlackoutSlateProperty.NetworkEndBlackoutImage`.
	NetworkEndBlackoutImage interface{} `json:"networkEndBlackoutImage"`
	// `CfnChannel.BlackoutSlateProperty.NetworkId`.
	NetworkId *string `json:"networkId"`
	// `CfnChannel.BlackoutSlateProperty.State`.
	State *string `json:"state"`
}

type CfnChannel_BurnInDestinationSettingsProperty struct {
	// `CfnChannel.BurnInDestinationSettingsProperty.Alignment`.
	Alignment *string `json:"alignment"`
	// `CfnChannel.BurnInDestinationSettingsProperty.BackgroundColor`.
	BackgroundColor *string `json:"backgroundColor"`
	// `CfnChannel.BurnInDestinationSettingsProperty.BackgroundOpacity`.
	BackgroundOpacity *float64 `json:"backgroundOpacity"`
	// `CfnChannel.BurnInDestinationSettingsProperty.Font`.
	Font interface{} `json:"font"`
	// `CfnChannel.BurnInDestinationSettingsProperty.FontColor`.
	FontColor *string `json:"fontColor"`
	// `CfnChannel.BurnInDestinationSettingsProperty.FontOpacity`.
	FontOpacity *float64 `json:"fontOpacity"`
	// `CfnChannel.BurnInDestinationSettingsProperty.FontResolution`.
	FontResolution *float64 `json:"fontResolution"`
	// `CfnChannel.BurnInDestinationSettingsProperty.FontSize`.
	FontSize *string `json:"fontSize"`
	// `CfnChannel.BurnInDestinationSettingsProperty.OutlineColor`.
	OutlineColor *string `json:"outlineColor"`
	// `CfnChannel.BurnInDestinationSettingsProperty.OutlineSize`.
	OutlineSize *float64 `json:"outlineSize"`
	// `CfnChannel.BurnInDestinationSettingsProperty.ShadowColor`.
	ShadowColor *string `json:"shadowColor"`
	// `CfnChannel.BurnInDestinationSettingsProperty.ShadowOpacity`.
	ShadowOpacity *float64 `json:"shadowOpacity"`
	// `CfnChannel.BurnInDestinationSettingsProperty.ShadowXOffset`.
	ShadowXOffset *float64 `json:"shadowXOffset"`
	// `CfnChannel.BurnInDestinationSettingsProperty.ShadowYOffset`.
	ShadowYOffset *float64 `json:"shadowYOffset"`
	// `CfnChannel.BurnInDestinationSettingsProperty.TeletextGridControl`.
	TeletextGridControl *string `json:"teletextGridControl"`
	// `CfnChannel.BurnInDestinationSettingsProperty.XPosition`.
	XPosition *float64 `json:"xPosition"`
	// `CfnChannel.BurnInDestinationSettingsProperty.YPosition`.
	YPosition *float64 `json:"yPosition"`
}

type CfnChannel_CaptionDescriptionProperty struct {
	// `CfnChannel.CaptionDescriptionProperty.CaptionSelectorName`.
	CaptionSelectorName *string `json:"captionSelectorName"`
	// `CfnChannel.CaptionDescriptionProperty.DestinationSettings`.
	DestinationSettings interface{} `json:"destinationSettings"`
	// `CfnChannel.CaptionDescriptionProperty.LanguageCode`.
	LanguageCode *string `json:"languageCode"`
	// `CfnChannel.CaptionDescriptionProperty.LanguageDescription`.
	LanguageDescription *string `json:"languageDescription"`
	// `CfnChannel.CaptionDescriptionProperty.Name`.
	Name *string `json:"name"`
}

type CfnChannel_CaptionDestinationSettingsProperty struct {
	// `CfnChannel.CaptionDestinationSettingsProperty.AribDestinationSettings`.
	AribDestinationSettings interface{} `json:"aribDestinationSettings"`
	// `CfnChannel.CaptionDestinationSettingsProperty.BurnInDestinationSettings`.
	BurnInDestinationSettings interface{} `json:"burnInDestinationSettings"`
	// `CfnChannel.CaptionDestinationSettingsProperty.DvbSubDestinationSettings`.
	DvbSubDestinationSettings interface{} `json:"dvbSubDestinationSettings"`
	// `CfnChannel.CaptionDestinationSettingsProperty.EbuTtDDestinationSettings`.
	EbuTtDDestinationSettings interface{} `json:"ebuTtDDestinationSettings"`
	// `CfnChannel.CaptionDestinationSettingsProperty.EmbeddedDestinationSettings`.
	EmbeddedDestinationSettings interface{} `json:"embeddedDestinationSettings"`
	// `CfnChannel.CaptionDestinationSettingsProperty.EmbeddedPlusScte20DestinationSettings`.
	EmbeddedPlusScte20DestinationSettings interface{} `json:"embeddedPlusScte20DestinationSettings"`
	// `CfnChannel.CaptionDestinationSettingsProperty.RtmpCaptionInfoDestinationSettings`.
	RtmpCaptionInfoDestinationSettings interface{} `json:"rtmpCaptionInfoDestinationSettings"`
	// `CfnChannel.CaptionDestinationSettingsProperty.Scte20PlusEmbeddedDestinationSettings`.
	Scte20PlusEmbeddedDestinationSettings interface{} `json:"scte20PlusEmbeddedDestinationSettings"`
	// `CfnChannel.CaptionDestinationSettingsProperty.Scte27DestinationSettings`.
	Scte27DestinationSettings interface{} `json:"scte27DestinationSettings"`
	// `CfnChannel.CaptionDestinationSettingsProperty.SmpteTtDestinationSettings`.
	SmpteTtDestinationSettings interface{} `json:"smpteTtDestinationSettings"`
	// `CfnChannel.CaptionDestinationSettingsProperty.TeletextDestinationSettings`.
	TeletextDestinationSettings interface{} `json:"teletextDestinationSettings"`
	// `CfnChannel.CaptionDestinationSettingsProperty.TtmlDestinationSettings`.
	TtmlDestinationSettings interface{} `json:"ttmlDestinationSettings"`
	// `CfnChannel.CaptionDestinationSettingsProperty.WebvttDestinationSettings`.
	WebvttDestinationSettings interface{} `json:"webvttDestinationSettings"`
}

type CfnChannel_CaptionLanguageMappingProperty struct {
	// `CfnChannel.CaptionLanguageMappingProperty.CaptionChannel`.
	CaptionChannel *float64 `json:"captionChannel"`
	// `CfnChannel.CaptionLanguageMappingProperty.LanguageCode`.
	LanguageCode *string `json:"languageCode"`
	// `CfnChannel.CaptionLanguageMappingProperty.LanguageDescription`.
	LanguageDescription *string `json:"languageDescription"`
}

type CfnChannel_CaptionRectangleProperty struct {
	// `CfnChannel.CaptionRectangleProperty.Height`.
	Height *float64 `json:"height"`
	// `CfnChannel.CaptionRectangleProperty.LeftOffset`.
	LeftOffset *float64 `json:"leftOffset"`
	// `CfnChannel.CaptionRectangleProperty.TopOffset`.
	TopOffset *float64 `json:"topOffset"`
	// `CfnChannel.CaptionRectangleProperty.Width`.
	Width *float64 `json:"width"`
}

type CfnChannel_CaptionSelectorProperty struct {
	// `CfnChannel.CaptionSelectorProperty.LanguageCode`.
	LanguageCode *string `json:"languageCode"`
	// `CfnChannel.CaptionSelectorProperty.Name`.
	Name *string `json:"name"`
	// `CfnChannel.CaptionSelectorProperty.SelectorSettings`.
	SelectorSettings interface{} `json:"selectorSettings"`
}

type CfnChannel_CaptionSelectorSettingsProperty struct {
	// `CfnChannel.CaptionSelectorSettingsProperty.AncillarySourceSettings`.
	AncillarySourceSettings interface{} `json:"ancillarySourceSettings"`
	// `CfnChannel.CaptionSelectorSettingsProperty.AribSourceSettings`.
	AribSourceSettings interface{} `json:"aribSourceSettings"`
	// `CfnChannel.CaptionSelectorSettingsProperty.DvbSubSourceSettings`.
	DvbSubSourceSettings interface{} `json:"dvbSubSourceSettings"`
	// `CfnChannel.CaptionSelectorSettingsProperty.EmbeddedSourceSettings`.
	EmbeddedSourceSettings interface{} `json:"embeddedSourceSettings"`
	// `CfnChannel.CaptionSelectorSettingsProperty.Scte20SourceSettings`.
	Scte20SourceSettings interface{} `json:"scte20SourceSettings"`
	// `CfnChannel.CaptionSelectorSettingsProperty.Scte27SourceSettings`.
	Scte27SourceSettings interface{} `json:"scte27SourceSettings"`
	// `CfnChannel.CaptionSelectorSettingsProperty.TeletextSourceSettings`.
	TeletextSourceSettings interface{} `json:"teletextSourceSettings"`
}

type CfnChannel_CdiInputSpecificationProperty struct {
	// `CfnChannel.CdiInputSpecificationProperty.Resolution`.
	Resolution *string `json:"resolution"`
}

type CfnChannel_ColorSpacePassthroughSettingsProperty struct {
}

type CfnChannel_DvbNitSettingsProperty struct {
	// `CfnChannel.DvbNitSettingsProperty.NetworkId`.
	NetworkId *float64 `json:"networkId"`
	// `CfnChannel.DvbNitSettingsProperty.NetworkName`.
	NetworkName *string `json:"networkName"`
	// `CfnChannel.DvbNitSettingsProperty.RepInterval`.
	RepInterval *float64 `json:"repInterval"`
}

type CfnChannel_DvbSdtSettingsProperty struct {
	// `CfnChannel.DvbSdtSettingsProperty.OutputSdt`.
	OutputSdt *string `json:"outputSdt"`
	// `CfnChannel.DvbSdtSettingsProperty.RepInterval`.
	RepInterval *float64 `json:"repInterval"`
	// `CfnChannel.DvbSdtSettingsProperty.ServiceName`.
	ServiceName *string `json:"serviceName"`
	// `CfnChannel.DvbSdtSettingsProperty.ServiceProviderName`.
	ServiceProviderName *string `json:"serviceProviderName"`
}

type CfnChannel_DvbSubDestinationSettingsProperty struct {
	// `CfnChannel.DvbSubDestinationSettingsProperty.Alignment`.
	Alignment *string `json:"alignment"`
	// `CfnChannel.DvbSubDestinationSettingsProperty.BackgroundColor`.
	BackgroundColor *string `json:"backgroundColor"`
	// `CfnChannel.DvbSubDestinationSettingsProperty.BackgroundOpacity`.
	BackgroundOpacity *float64 `json:"backgroundOpacity"`
	// `CfnChannel.DvbSubDestinationSettingsProperty.Font`.
	Font interface{} `json:"font"`
	// `CfnChannel.DvbSubDestinationSettingsProperty.FontColor`.
	FontColor *string `json:"fontColor"`
	// `CfnChannel.DvbSubDestinationSettingsProperty.FontOpacity`.
	FontOpacity *float64 `json:"fontOpacity"`
	// `CfnChannel.DvbSubDestinationSettingsProperty.FontResolution`.
	FontResolution *float64 `json:"fontResolution"`
	// `CfnChannel.DvbSubDestinationSettingsProperty.FontSize`.
	FontSize *string `json:"fontSize"`
	// `CfnChannel.DvbSubDestinationSettingsProperty.OutlineColor`.
	OutlineColor *string `json:"outlineColor"`
	// `CfnChannel.DvbSubDestinationSettingsProperty.OutlineSize`.
	OutlineSize *float64 `json:"outlineSize"`
	// `CfnChannel.DvbSubDestinationSettingsProperty.ShadowColor`.
	ShadowColor *string `json:"shadowColor"`
	// `CfnChannel.DvbSubDestinationSettingsProperty.ShadowOpacity`.
	ShadowOpacity *float64 `json:"shadowOpacity"`
	// `CfnChannel.DvbSubDestinationSettingsProperty.ShadowXOffset`.
	ShadowXOffset *float64 `json:"shadowXOffset"`
	// `CfnChannel.DvbSubDestinationSettingsProperty.ShadowYOffset`.
	ShadowYOffset *float64 `json:"shadowYOffset"`
	// `CfnChannel.DvbSubDestinationSettingsProperty.TeletextGridControl`.
	TeletextGridControl *string `json:"teletextGridControl"`
	// `CfnChannel.DvbSubDestinationSettingsProperty.XPosition`.
	XPosition *float64 `json:"xPosition"`
	// `CfnChannel.DvbSubDestinationSettingsProperty.YPosition`.
	YPosition *float64 `json:"yPosition"`
}

type CfnChannel_DvbSubSourceSettingsProperty struct {
	// `CfnChannel.DvbSubSourceSettingsProperty.Pid`.
	Pid *float64 `json:"pid"`
}

type CfnChannel_DvbTdtSettingsProperty struct {
	// `CfnChannel.DvbTdtSettingsProperty.RepInterval`.
	RepInterval *float64 `json:"repInterval"`
}

type CfnChannel_Eac3SettingsProperty struct {
	// `CfnChannel.Eac3SettingsProperty.AttenuationControl`.
	AttenuationControl *string `json:"attenuationControl"`
	// `CfnChannel.Eac3SettingsProperty.Bitrate`.
	Bitrate *float64 `json:"bitrate"`
	// `CfnChannel.Eac3SettingsProperty.BitstreamMode`.
	BitstreamMode *string `json:"bitstreamMode"`
	// `CfnChannel.Eac3SettingsProperty.CodingMode`.
	CodingMode *string `json:"codingMode"`
	// `CfnChannel.Eac3SettingsProperty.DcFilter`.
	DcFilter *string `json:"dcFilter"`
	// `CfnChannel.Eac3SettingsProperty.Dialnorm`.
	Dialnorm *float64 `json:"dialnorm"`
	// `CfnChannel.Eac3SettingsProperty.DrcLine`.
	DrcLine *string `json:"drcLine"`
	// `CfnChannel.Eac3SettingsProperty.DrcRf`.
	DrcRf *string `json:"drcRf"`
	// `CfnChannel.Eac3SettingsProperty.LfeControl`.
	LfeControl *string `json:"lfeControl"`
	// `CfnChannel.Eac3SettingsProperty.LfeFilter`.
	LfeFilter *string `json:"lfeFilter"`
	// `CfnChannel.Eac3SettingsProperty.LoRoCenterMixLevel`.
	LoRoCenterMixLevel *float64 `json:"loRoCenterMixLevel"`
	// `CfnChannel.Eac3SettingsProperty.LoRoSurroundMixLevel`.
	LoRoSurroundMixLevel *float64 `json:"loRoSurroundMixLevel"`
	// `CfnChannel.Eac3SettingsProperty.LtRtCenterMixLevel`.
	LtRtCenterMixLevel *float64 `json:"ltRtCenterMixLevel"`
	// `CfnChannel.Eac3SettingsProperty.LtRtSurroundMixLevel`.
	LtRtSurroundMixLevel *float64 `json:"ltRtSurroundMixLevel"`
	// `CfnChannel.Eac3SettingsProperty.MetadataControl`.
	MetadataControl *string `json:"metadataControl"`
	// `CfnChannel.Eac3SettingsProperty.PassthroughControl`.
	PassthroughControl *string `json:"passthroughControl"`
	// `CfnChannel.Eac3SettingsProperty.PhaseControl`.
	PhaseControl *string `json:"phaseControl"`
	// `CfnChannel.Eac3SettingsProperty.StereoDownmix`.
	StereoDownmix *string `json:"stereoDownmix"`
	// `CfnChannel.Eac3SettingsProperty.SurroundExMode`.
	SurroundExMode *string `json:"surroundExMode"`
	// `CfnChannel.Eac3SettingsProperty.SurroundMode`.
	SurroundMode *string `json:"surroundMode"`
}

type CfnChannel_EbuTtDDestinationSettingsProperty struct {
	// `CfnChannel.EbuTtDDestinationSettingsProperty.CopyrightHolder`.
	CopyrightHolder *string `json:"copyrightHolder"`
	// `CfnChannel.EbuTtDDestinationSettingsProperty.FillLineGap`.
	FillLineGap *string `json:"fillLineGap"`
	// `CfnChannel.EbuTtDDestinationSettingsProperty.FontFamily`.
	FontFamily *string `json:"fontFamily"`
	// `CfnChannel.EbuTtDDestinationSettingsProperty.StyleControl`.
	StyleControl *string `json:"styleControl"`
}

type CfnChannel_EmbeddedDestinationSettingsProperty struct {
}

type CfnChannel_EmbeddedPlusScte20DestinationSettingsProperty struct {
}

type CfnChannel_EmbeddedSourceSettingsProperty struct {
	// `CfnChannel.EmbeddedSourceSettingsProperty.Convert608To708`.
	Convert608To708 *string `json:"convert608To708"`
	// `CfnChannel.EmbeddedSourceSettingsProperty.Scte20Detection`.
	Scte20Detection *string `json:"scte20Detection"`
	// `CfnChannel.EmbeddedSourceSettingsProperty.Source608ChannelNumber`.
	Source608ChannelNumber *float64 `json:"source608ChannelNumber"`
	// `CfnChannel.EmbeddedSourceSettingsProperty.Source608TrackNumber`.
	Source608TrackNumber *float64 `json:"source608TrackNumber"`
}

type CfnChannel_EncoderSettingsProperty struct {
	// `CfnChannel.EncoderSettingsProperty.AudioDescriptions`.
	AudioDescriptions interface{} `json:"audioDescriptions"`
	// `CfnChannel.EncoderSettingsProperty.AvailBlanking`.
	AvailBlanking interface{} `json:"availBlanking"`
	// `CfnChannel.EncoderSettingsProperty.AvailConfiguration`.
	AvailConfiguration interface{} `json:"availConfiguration"`
	// `CfnChannel.EncoderSettingsProperty.BlackoutSlate`.
	BlackoutSlate interface{} `json:"blackoutSlate"`
	// `CfnChannel.EncoderSettingsProperty.CaptionDescriptions`.
	CaptionDescriptions interface{} `json:"captionDescriptions"`
	// `CfnChannel.EncoderSettingsProperty.FeatureActivations`.
	FeatureActivations interface{} `json:"featureActivations"`
	// `CfnChannel.EncoderSettingsProperty.GlobalConfiguration`.
	GlobalConfiguration interface{} `json:"globalConfiguration"`
	// `CfnChannel.EncoderSettingsProperty.MotionGraphicsConfiguration`.
	MotionGraphicsConfiguration interface{} `json:"motionGraphicsConfiguration"`
	// `CfnChannel.EncoderSettingsProperty.NielsenConfiguration`.
	NielsenConfiguration interface{} `json:"nielsenConfiguration"`
	// `CfnChannel.EncoderSettingsProperty.OutputGroups`.
	OutputGroups interface{} `json:"outputGroups"`
	// `CfnChannel.EncoderSettingsProperty.TimecodeConfig`.
	TimecodeConfig interface{} `json:"timecodeConfig"`
	// `CfnChannel.EncoderSettingsProperty.VideoDescriptions`.
	VideoDescriptions interface{} `json:"videoDescriptions"`
}

type CfnChannel_FailoverConditionProperty struct {
	// `CfnChannel.FailoverConditionProperty.FailoverConditionSettings`.
	FailoverConditionSettings interface{} `json:"failoverConditionSettings"`
}

type CfnChannel_FailoverConditionSettingsProperty struct {
	// `CfnChannel.FailoverConditionSettingsProperty.AudioSilenceSettings`.
	AudioSilenceSettings interface{} `json:"audioSilenceSettings"`
	// `CfnChannel.FailoverConditionSettingsProperty.InputLossSettings`.
	InputLossSettings interface{} `json:"inputLossSettings"`
	// `CfnChannel.FailoverConditionSettingsProperty.VideoBlackSettings`.
	VideoBlackSettings interface{} `json:"videoBlackSettings"`
}

type CfnChannel_FeatureActivationsProperty struct {
	// `CfnChannel.FeatureActivationsProperty.InputPrepareScheduleActions`.
	InputPrepareScheduleActions *string `json:"inputPrepareScheduleActions"`
}

type CfnChannel_FecOutputSettingsProperty struct {
	// `CfnChannel.FecOutputSettingsProperty.ColumnDepth`.
	ColumnDepth *float64 `json:"columnDepth"`
	// `CfnChannel.FecOutputSettingsProperty.IncludeFec`.
	IncludeFec *string `json:"includeFec"`
	// `CfnChannel.FecOutputSettingsProperty.RowLength`.
	RowLength *float64 `json:"rowLength"`
}

type CfnChannel_Fmp4HlsSettingsProperty struct {
	// `CfnChannel.Fmp4HlsSettingsProperty.AudioRenditionSets`.
	AudioRenditionSets *string `json:"audioRenditionSets"`
	// `CfnChannel.Fmp4HlsSettingsProperty.NielsenId3Behavior`.
	NielsenId3Behavior *string `json:"nielsenId3Behavior"`
	// `CfnChannel.Fmp4HlsSettingsProperty.TimedMetadataBehavior`.
	TimedMetadataBehavior *string `json:"timedMetadataBehavior"`
}

type CfnChannel_FrameCaptureCdnSettingsProperty struct {
	// `CfnChannel.FrameCaptureCdnSettingsProperty.FrameCaptureS3Settings`.
	FrameCaptureS3Settings interface{} `json:"frameCaptureS3Settings"`
}

type CfnChannel_FrameCaptureGroupSettingsProperty struct {
	// `CfnChannel.FrameCaptureGroupSettingsProperty.Destination`.
	Destination interface{} `json:"destination"`
	// `CfnChannel.FrameCaptureGroupSettingsProperty.FrameCaptureCdnSettings`.
	FrameCaptureCdnSettings interface{} `json:"frameCaptureCdnSettings"`
}

type CfnChannel_FrameCaptureHlsSettingsProperty struct {
}

type CfnChannel_FrameCaptureOutputSettingsProperty struct {
	// `CfnChannel.FrameCaptureOutputSettingsProperty.NameModifier`.
	NameModifier *string `json:"nameModifier"`
}

type CfnChannel_FrameCaptureS3SettingsProperty struct {
	// `CfnChannel.FrameCaptureS3SettingsProperty.CannedAcl`.
	CannedAcl *string `json:"cannedAcl"`
}

type CfnChannel_FrameCaptureSettingsProperty struct {
	// `CfnChannel.FrameCaptureSettingsProperty.CaptureInterval`.
	CaptureInterval *float64 `json:"captureInterval"`
	// `CfnChannel.FrameCaptureSettingsProperty.CaptureIntervalUnits`.
	CaptureIntervalUnits *string `json:"captureIntervalUnits"`
}

type CfnChannel_GlobalConfigurationProperty struct {
	// `CfnChannel.GlobalConfigurationProperty.InitialAudioGain`.
	InitialAudioGain *float64 `json:"initialAudioGain"`
	// `CfnChannel.GlobalConfigurationProperty.InputEndAction`.
	InputEndAction *string `json:"inputEndAction"`
	// `CfnChannel.GlobalConfigurationProperty.InputLossBehavior`.
	InputLossBehavior interface{} `json:"inputLossBehavior"`
	// `CfnChannel.GlobalConfigurationProperty.OutputLockingMode`.
	OutputLockingMode *string `json:"outputLockingMode"`
	// `CfnChannel.GlobalConfigurationProperty.OutputTimingSource`.
	OutputTimingSource *string `json:"outputTimingSource"`
	// `CfnChannel.GlobalConfigurationProperty.SupportLowFramerateInputs`.
	SupportLowFramerateInputs *string `json:"supportLowFramerateInputs"`
}

type CfnChannel_H264ColorSpaceSettingsProperty struct {
	// `CfnChannel.H264ColorSpaceSettingsProperty.ColorSpacePassthroughSettings`.
	ColorSpacePassthroughSettings interface{} `json:"colorSpacePassthroughSettings"`
	// `CfnChannel.H264ColorSpaceSettingsProperty.Rec601Settings`.
	Rec601Settings interface{} `json:"rec601Settings"`
	// `CfnChannel.H264ColorSpaceSettingsProperty.Rec709Settings`.
	Rec709Settings interface{} `json:"rec709Settings"`
}

type CfnChannel_H264FilterSettingsProperty struct {
	// `CfnChannel.H264FilterSettingsProperty.TemporalFilterSettings`.
	TemporalFilterSettings interface{} `json:"temporalFilterSettings"`
}

type CfnChannel_H264SettingsProperty struct {
	// `CfnChannel.H264SettingsProperty.AdaptiveQuantization`.
	AdaptiveQuantization *string `json:"adaptiveQuantization"`
	// `CfnChannel.H264SettingsProperty.AfdSignaling`.
	AfdSignaling *string `json:"afdSignaling"`
	// `CfnChannel.H264SettingsProperty.Bitrate`.
	Bitrate *float64 `json:"bitrate"`
	// `CfnChannel.H264SettingsProperty.BufFillPct`.
	BufFillPct *float64 `json:"bufFillPct"`
	// `CfnChannel.H264SettingsProperty.BufSize`.
	BufSize *float64 `json:"bufSize"`
	// `CfnChannel.H264SettingsProperty.ColorMetadata`.
	ColorMetadata *string `json:"colorMetadata"`
	// `CfnChannel.H264SettingsProperty.ColorSpaceSettings`.
	ColorSpaceSettings interface{} `json:"colorSpaceSettings"`
	// `CfnChannel.H264SettingsProperty.EntropyEncoding`.
	EntropyEncoding *string `json:"entropyEncoding"`
	// `CfnChannel.H264SettingsProperty.FilterSettings`.
	FilterSettings interface{} `json:"filterSettings"`
	// `CfnChannel.H264SettingsProperty.FixedAfd`.
	FixedAfd *string `json:"fixedAfd"`
	// `CfnChannel.H264SettingsProperty.FlickerAq`.
	FlickerAq *string `json:"flickerAq"`
	// `CfnChannel.H264SettingsProperty.ForceFieldPictures`.
	ForceFieldPictures *string `json:"forceFieldPictures"`
	// `CfnChannel.H264SettingsProperty.FramerateControl`.
	FramerateControl *string `json:"framerateControl"`
	// `CfnChannel.H264SettingsProperty.FramerateDenominator`.
	FramerateDenominator *float64 `json:"framerateDenominator"`
	// `CfnChannel.H264SettingsProperty.FramerateNumerator`.
	FramerateNumerator *float64 `json:"framerateNumerator"`
	// `CfnChannel.H264SettingsProperty.GopBReference`.
	GopBReference *string `json:"gopBReference"`
	// `CfnChannel.H264SettingsProperty.GopClosedCadence`.
	GopClosedCadence *float64 `json:"gopClosedCadence"`
	// `CfnChannel.H264SettingsProperty.GopNumBFrames`.
	GopNumBFrames *float64 `json:"gopNumBFrames"`
	// `CfnChannel.H264SettingsProperty.GopSize`.
	GopSize *float64 `json:"gopSize"`
	// `CfnChannel.H264SettingsProperty.GopSizeUnits`.
	GopSizeUnits *string `json:"gopSizeUnits"`
	// `CfnChannel.H264SettingsProperty.Level`.
	Level *string `json:"level"`
	// `CfnChannel.H264SettingsProperty.LookAheadRateControl`.
	LookAheadRateControl *string `json:"lookAheadRateControl"`
	// `CfnChannel.H264SettingsProperty.MaxBitrate`.
	MaxBitrate *float64 `json:"maxBitrate"`
	// `CfnChannel.H264SettingsProperty.MinIInterval`.
	MinIInterval *float64 `json:"minIInterval"`
	// `CfnChannel.H264SettingsProperty.NumRefFrames`.
	NumRefFrames *float64 `json:"numRefFrames"`
	// `CfnChannel.H264SettingsProperty.ParControl`.
	ParControl *string `json:"parControl"`
	// `CfnChannel.H264SettingsProperty.ParDenominator`.
	ParDenominator *float64 `json:"parDenominator"`
	// `CfnChannel.H264SettingsProperty.ParNumerator`.
	ParNumerator *float64 `json:"parNumerator"`
	// `CfnChannel.H264SettingsProperty.Profile`.
	Profile *string `json:"profile"`
	// `CfnChannel.H264SettingsProperty.QualityLevel`.
	QualityLevel *string `json:"qualityLevel"`
	// `CfnChannel.H264SettingsProperty.QvbrQualityLevel`.
	QvbrQualityLevel *float64 `json:"qvbrQualityLevel"`
	// `CfnChannel.H264SettingsProperty.RateControlMode`.
	RateControlMode *string `json:"rateControlMode"`
	// `CfnChannel.H264SettingsProperty.ScanType`.
	ScanType *string `json:"scanType"`
	// `CfnChannel.H264SettingsProperty.SceneChangeDetect`.
	SceneChangeDetect *string `json:"sceneChangeDetect"`
	// `CfnChannel.H264SettingsProperty.Slices`.
	Slices *float64 `json:"slices"`
	// `CfnChannel.H264SettingsProperty.Softness`.
	Softness *float64 `json:"softness"`
	// `CfnChannel.H264SettingsProperty.SpatialAq`.
	SpatialAq *string `json:"spatialAq"`
	// `CfnChannel.H264SettingsProperty.SubgopLength`.
	SubgopLength *string `json:"subgopLength"`
	// `CfnChannel.H264SettingsProperty.Syntax`.
	Syntax *string `json:"syntax"`
	// `CfnChannel.H264SettingsProperty.TemporalAq`.
	TemporalAq *string `json:"temporalAq"`
	// `CfnChannel.H264SettingsProperty.TimecodeInsertion`.
	TimecodeInsertion *string `json:"timecodeInsertion"`
}

type CfnChannel_H265ColorSpaceSettingsProperty struct {
	// `CfnChannel.H265ColorSpaceSettingsProperty.ColorSpacePassthroughSettings`.
	ColorSpacePassthroughSettings interface{} `json:"colorSpacePassthroughSettings"`
	// `CfnChannel.H265ColorSpaceSettingsProperty.Hdr10Settings`.
	Hdr10Settings interface{} `json:"hdr10Settings"`
	// `CfnChannel.H265ColorSpaceSettingsProperty.Rec601Settings`.
	Rec601Settings interface{} `json:"rec601Settings"`
	// `CfnChannel.H265ColorSpaceSettingsProperty.Rec709Settings`.
	Rec709Settings interface{} `json:"rec709Settings"`
}

type CfnChannel_H265FilterSettingsProperty struct {
	// `CfnChannel.H265FilterSettingsProperty.TemporalFilterSettings`.
	TemporalFilterSettings interface{} `json:"temporalFilterSettings"`
}

type CfnChannel_H265SettingsProperty struct {
	// `CfnChannel.H265SettingsProperty.AdaptiveQuantization`.
	AdaptiveQuantization *string `json:"adaptiveQuantization"`
	// `CfnChannel.H265SettingsProperty.AfdSignaling`.
	AfdSignaling *string `json:"afdSignaling"`
	// `CfnChannel.H265SettingsProperty.AlternativeTransferFunction`.
	AlternativeTransferFunction *string `json:"alternativeTransferFunction"`
	// `CfnChannel.H265SettingsProperty.Bitrate`.
	Bitrate *float64 `json:"bitrate"`
	// `CfnChannel.H265SettingsProperty.BufSize`.
	BufSize *float64 `json:"bufSize"`
	// `CfnChannel.H265SettingsProperty.ColorMetadata`.
	ColorMetadata *string `json:"colorMetadata"`
	// `CfnChannel.H265SettingsProperty.ColorSpaceSettings`.
	ColorSpaceSettings interface{} `json:"colorSpaceSettings"`
	// `CfnChannel.H265SettingsProperty.FilterSettings`.
	FilterSettings interface{} `json:"filterSettings"`
	// `CfnChannel.H265SettingsProperty.FixedAfd`.
	FixedAfd *string `json:"fixedAfd"`
	// `CfnChannel.H265SettingsProperty.FlickerAq`.
	FlickerAq *string `json:"flickerAq"`
	// `CfnChannel.H265SettingsProperty.FramerateDenominator`.
	FramerateDenominator *float64 `json:"framerateDenominator"`
	// `CfnChannel.H265SettingsProperty.FramerateNumerator`.
	FramerateNumerator *float64 `json:"framerateNumerator"`
	// `CfnChannel.H265SettingsProperty.GopClosedCadence`.
	GopClosedCadence *float64 `json:"gopClosedCadence"`
	// `CfnChannel.H265SettingsProperty.GopSize`.
	GopSize *float64 `json:"gopSize"`
	// `CfnChannel.H265SettingsProperty.GopSizeUnits`.
	GopSizeUnits *string `json:"gopSizeUnits"`
	// `CfnChannel.H265SettingsProperty.Level`.
	Level *string `json:"level"`
	// `CfnChannel.H265SettingsProperty.LookAheadRateControl`.
	LookAheadRateControl *string `json:"lookAheadRateControl"`
	// `CfnChannel.H265SettingsProperty.MaxBitrate`.
	MaxBitrate *float64 `json:"maxBitrate"`
	// `CfnChannel.H265SettingsProperty.MinIInterval`.
	MinIInterval *float64 `json:"minIInterval"`
	// `CfnChannel.H265SettingsProperty.ParDenominator`.
	ParDenominator *float64 `json:"parDenominator"`
	// `CfnChannel.H265SettingsProperty.ParNumerator`.
	ParNumerator *float64 `json:"parNumerator"`
	// `CfnChannel.H265SettingsProperty.Profile`.
	Profile *string `json:"profile"`
	// `CfnChannel.H265SettingsProperty.QvbrQualityLevel`.
	QvbrQualityLevel *float64 `json:"qvbrQualityLevel"`
	// `CfnChannel.H265SettingsProperty.RateControlMode`.
	RateControlMode *string `json:"rateControlMode"`
	// `CfnChannel.H265SettingsProperty.ScanType`.
	ScanType *string `json:"scanType"`
	// `CfnChannel.H265SettingsProperty.SceneChangeDetect`.
	SceneChangeDetect *string `json:"sceneChangeDetect"`
	// `CfnChannel.H265SettingsProperty.Slices`.
	Slices *float64 `json:"slices"`
	// `CfnChannel.H265SettingsProperty.Tier`.
	Tier *string `json:"tier"`
	// `CfnChannel.H265SettingsProperty.TimecodeInsertion`.
	TimecodeInsertion *string `json:"timecodeInsertion"`
}

type CfnChannel_Hdr10SettingsProperty struct {
	// `CfnChannel.Hdr10SettingsProperty.MaxCll`.
	MaxCll *float64 `json:"maxCll"`
	// `CfnChannel.Hdr10SettingsProperty.MaxFall`.
	MaxFall *float64 `json:"maxFall"`
}

type CfnChannel_HlsAkamaiSettingsProperty struct {
	// `CfnChannel.HlsAkamaiSettingsProperty.ConnectionRetryInterval`.
	ConnectionRetryInterval *float64 `json:"connectionRetryInterval"`
	// `CfnChannel.HlsAkamaiSettingsProperty.FilecacheDuration`.
	FilecacheDuration *float64 `json:"filecacheDuration"`
	// `CfnChannel.HlsAkamaiSettingsProperty.HttpTransferMode`.
	HttpTransferMode *string `json:"httpTransferMode"`
	// `CfnChannel.HlsAkamaiSettingsProperty.NumRetries`.
	NumRetries *float64 `json:"numRetries"`
	// `CfnChannel.HlsAkamaiSettingsProperty.RestartDelay`.
	RestartDelay *float64 `json:"restartDelay"`
	// `CfnChannel.HlsAkamaiSettingsProperty.Salt`.
	Salt *string `json:"salt"`
	// `CfnChannel.HlsAkamaiSettingsProperty.Token`.
	Token *string `json:"token"`
}

type CfnChannel_HlsBasicPutSettingsProperty struct {
	// `CfnChannel.HlsBasicPutSettingsProperty.ConnectionRetryInterval`.
	ConnectionRetryInterval *float64 `json:"connectionRetryInterval"`
	// `CfnChannel.HlsBasicPutSettingsProperty.FilecacheDuration`.
	FilecacheDuration *float64 `json:"filecacheDuration"`
	// `CfnChannel.HlsBasicPutSettingsProperty.NumRetries`.
	NumRetries *float64 `json:"numRetries"`
	// `CfnChannel.HlsBasicPutSettingsProperty.RestartDelay`.
	RestartDelay *float64 `json:"restartDelay"`
}

type CfnChannel_HlsCdnSettingsProperty struct {
	// `CfnChannel.HlsCdnSettingsProperty.HlsAkamaiSettings`.
	HlsAkamaiSettings interface{} `json:"hlsAkamaiSettings"`
	// `CfnChannel.HlsCdnSettingsProperty.HlsBasicPutSettings`.
	HlsBasicPutSettings interface{} `json:"hlsBasicPutSettings"`
	// `CfnChannel.HlsCdnSettingsProperty.HlsMediaStoreSettings`.
	HlsMediaStoreSettings interface{} `json:"hlsMediaStoreSettings"`
	// `CfnChannel.HlsCdnSettingsProperty.HlsS3Settings`.
	HlsS3Settings interface{} `json:"hlsS3Settings"`
	// `CfnChannel.HlsCdnSettingsProperty.HlsWebdavSettings`.
	HlsWebdavSettings interface{} `json:"hlsWebdavSettings"`
}

type CfnChannel_HlsGroupSettingsProperty struct {
	// `CfnChannel.HlsGroupSettingsProperty.AdMarkers`.
	AdMarkers *[]*string `json:"adMarkers"`
	// `CfnChannel.HlsGroupSettingsProperty.BaseUrlContent`.
	BaseUrlContent *string `json:"baseUrlContent"`
	// `CfnChannel.HlsGroupSettingsProperty.BaseUrlContent1`.
	BaseUrlContent1 *string `json:"baseUrlContent1"`
	// `CfnChannel.HlsGroupSettingsProperty.BaseUrlManifest`.
	BaseUrlManifest *string `json:"baseUrlManifest"`
	// `CfnChannel.HlsGroupSettingsProperty.BaseUrlManifest1`.
	BaseUrlManifest1 *string `json:"baseUrlManifest1"`
	// `CfnChannel.HlsGroupSettingsProperty.CaptionLanguageMappings`.
	CaptionLanguageMappings interface{} `json:"captionLanguageMappings"`
	// `CfnChannel.HlsGroupSettingsProperty.CaptionLanguageSetting`.
	CaptionLanguageSetting *string `json:"captionLanguageSetting"`
	// `CfnChannel.HlsGroupSettingsProperty.ClientCache`.
	ClientCache *string `json:"clientCache"`
	// `CfnChannel.HlsGroupSettingsProperty.CodecSpecification`.
	CodecSpecification *string `json:"codecSpecification"`
	// `CfnChannel.HlsGroupSettingsProperty.ConstantIv`.
	ConstantIv *string `json:"constantIv"`
	// `CfnChannel.HlsGroupSettingsProperty.Destination`.
	Destination interface{} `json:"destination"`
	// `CfnChannel.HlsGroupSettingsProperty.DirectoryStructure`.
	DirectoryStructure *string `json:"directoryStructure"`
	// `CfnChannel.HlsGroupSettingsProperty.DiscontinuityTags`.
	DiscontinuityTags *string `json:"discontinuityTags"`
	// `CfnChannel.HlsGroupSettingsProperty.EncryptionType`.
	EncryptionType *string `json:"encryptionType"`
	// `CfnChannel.HlsGroupSettingsProperty.HlsCdnSettings`.
	HlsCdnSettings interface{} `json:"hlsCdnSettings"`
	// `CfnChannel.HlsGroupSettingsProperty.HlsId3SegmentTagging`.
	HlsId3SegmentTagging *string `json:"hlsId3SegmentTagging"`
	// `CfnChannel.HlsGroupSettingsProperty.IFrameOnlyPlaylists`.
	IFrameOnlyPlaylists *string `json:"iFrameOnlyPlaylists"`
	// `CfnChannel.HlsGroupSettingsProperty.IncompleteSegmentBehavior`.
	IncompleteSegmentBehavior *string `json:"incompleteSegmentBehavior"`
	// `CfnChannel.HlsGroupSettingsProperty.IndexNSegments`.
	IndexNSegments *float64 `json:"indexNSegments"`
	// `CfnChannel.HlsGroupSettingsProperty.InputLossAction`.
	InputLossAction *string `json:"inputLossAction"`
	// `CfnChannel.HlsGroupSettingsProperty.IvInManifest`.
	IvInManifest *string `json:"ivInManifest"`
	// `CfnChannel.HlsGroupSettingsProperty.IvSource`.
	IvSource *string `json:"ivSource"`
	// `CfnChannel.HlsGroupSettingsProperty.KeepSegments`.
	KeepSegments *float64 `json:"keepSegments"`
	// `CfnChannel.HlsGroupSettingsProperty.KeyFormat`.
	KeyFormat *string `json:"keyFormat"`
	// `CfnChannel.HlsGroupSettingsProperty.KeyFormatVersions`.
	KeyFormatVersions *string `json:"keyFormatVersions"`
	// `CfnChannel.HlsGroupSettingsProperty.KeyProviderSettings`.
	KeyProviderSettings interface{} `json:"keyProviderSettings"`
	// `CfnChannel.HlsGroupSettingsProperty.ManifestCompression`.
	ManifestCompression *string `json:"manifestCompression"`
	// `CfnChannel.HlsGroupSettingsProperty.ManifestDurationFormat`.
	ManifestDurationFormat *string `json:"manifestDurationFormat"`
	// `CfnChannel.HlsGroupSettingsProperty.MinSegmentLength`.
	MinSegmentLength *float64 `json:"minSegmentLength"`
	// `CfnChannel.HlsGroupSettingsProperty.Mode`.
	Mode *string `json:"mode"`
	// `CfnChannel.HlsGroupSettingsProperty.OutputSelection`.
	OutputSelection *string `json:"outputSelection"`
	// `CfnChannel.HlsGroupSettingsProperty.ProgramDateTime`.
	ProgramDateTime *string `json:"programDateTime"`
	// `CfnChannel.HlsGroupSettingsProperty.ProgramDateTimePeriod`.
	ProgramDateTimePeriod *float64 `json:"programDateTimePeriod"`
	// `CfnChannel.HlsGroupSettingsProperty.RedundantManifest`.
	RedundantManifest *string `json:"redundantManifest"`
	// `CfnChannel.HlsGroupSettingsProperty.SegmentationMode`.
	SegmentationMode *string `json:"segmentationMode"`
	// `CfnChannel.HlsGroupSettingsProperty.SegmentLength`.
	SegmentLength *float64 `json:"segmentLength"`
	// `CfnChannel.HlsGroupSettingsProperty.SegmentsPerSubdirectory`.
	SegmentsPerSubdirectory *float64 `json:"segmentsPerSubdirectory"`
	// `CfnChannel.HlsGroupSettingsProperty.StreamInfResolution`.
	StreamInfResolution *string `json:"streamInfResolution"`
	// `CfnChannel.HlsGroupSettingsProperty.TimedMetadataId3Frame`.
	TimedMetadataId3Frame *string `json:"timedMetadataId3Frame"`
	// `CfnChannel.HlsGroupSettingsProperty.TimedMetadataId3Period`.
	TimedMetadataId3Period *float64 `json:"timedMetadataId3Period"`
	// `CfnChannel.HlsGroupSettingsProperty.TimestampDeltaMilliseconds`.
	TimestampDeltaMilliseconds *float64 `json:"timestampDeltaMilliseconds"`
	// `CfnChannel.HlsGroupSettingsProperty.TsFileMode`.
	TsFileMode *string `json:"tsFileMode"`
}

type CfnChannel_HlsInputSettingsProperty struct {
	// `CfnChannel.HlsInputSettingsProperty.Bandwidth`.
	Bandwidth *float64 `json:"bandwidth"`
	// `CfnChannel.HlsInputSettingsProperty.BufferSegments`.
	BufferSegments *float64 `json:"bufferSegments"`
	// `CfnChannel.HlsInputSettingsProperty.Retries`.
	Retries *float64 `json:"retries"`
	// `CfnChannel.HlsInputSettingsProperty.RetryInterval`.
	RetryInterval *float64 `json:"retryInterval"`
}

type CfnChannel_HlsMediaStoreSettingsProperty struct {
	// `CfnChannel.HlsMediaStoreSettingsProperty.ConnectionRetryInterval`.
	ConnectionRetryInterval *float64 `json:"connectionRetryInterval"`
	// `CfnChannel.HlsMediaStoreSettingsProperty.FilecacheDuration`.
	FilecacheDuration *float64 `json:"filecacheDuration"`
	// `CfnChannel.HlsMediaStoreSettingsProperty.MediaStoreStorageClass`.
	MediaStoreStorageClass *string `json:"mediaStoreStorageClass"`
	// `CfnChannel.HlsMediaStoreSettingsProperty.NumRetries`.
	NumRetries *float64 `json:"numRetries"`
	// `CfnChannel.HlsMediaStoreSettingsProperty.RestartDelay`.
	RestartDelay *float64 `json:"restartDelay"`
}

type CfnChannel_HlsOutputSettingsProperty struct {
	// `CfnChannel.HlsOutputSettingsProperty.H265PackagingType`.
	H265PackagingType *string `json:"h265PackagingType"`
	// `CfnChannel.HlsOutputSettingsProperty.HlsSettings`.
	HlsSettings interface{} `json:"hlsSettings"`
	// `CfnChannel.HlsOutputSettingsProperty.NameModifier`.
	NameModifier *string `json:"nameModifier"`
	// `CfnChannel.HlsOutputSettingsProperty.SegmentModifier`.
	SegmentModifier *string `json:"segmentModifier"`
}

type CfnChannel_HlsS3SettingsProperty struct {
	// `CfnChannel.HlsS3SettingsProperty.CannedAcl`.
	CannedAcl *string `json:"cannedAcl"`
}

type CfnChannel_HlsSettingsProperty struct {
	// `CfnChannel.HlsSettingsProperty.AudioOnlyHlsSettings`.
	AudioOnlyHlsSettings interface{} `json:"audioOnlyHlsSettings"`
	// `CfnChannel.HlsSettingsProperty.Fmp4HlsSettings`.
	Fmp4HlsSettings interface{} `json:"fmp4HlsSettings"`
	// `CfnChannel.HlsSettingsProperty.FrameCaptureHlsSettings`.
	FrameCaptureHlsSettings interface{} `json:"frameCaptureHlsSettings"`
	// `CfnChannel.HlsSettingsProperty.StandardHlsSettings`.
	StandardHlsSettings interface{} `json:"standardHlsSettings"`
}

type CfnChannel_HlsWebdavSettingsProperty struct {
	// `CfnChannel.HlsWebdavSettingsProperty.ConnectionRetryInterval`.
	ConnectionRetryInterval *float64 `json:"connectionRetryInterval"`
	// `CfnChannel.HlsWebdavSettingsProperty.FilecacheDuration`.
	FilecacheDuration *float64 `json:"filecacheDuration"`
	// `CfnChannel.HlsWebdavSettingsProperty.HttpTransferMode`.
	HttpTransferMode *string `json:"httpTransferMode"`
	// `CfnChannel.HlsWebdavSettingsProperty.NumRetries`.
	NumRetries *float64 `json:"numRetries"`
	// `CfnChannel.HlsWebdavSettingsProperty.RestartDelay`.
	RestartDelay *float64 `json:"restartDelay"`
}

type CfnChannel_HtmlMotionGraphicsSettingsProperty struct {
}

type CfnChannel_InputAttachmentProperty struct {
	// `CfnChannel.InputAttachmentProperty.AutomaticInputFailoverSettings`.
	AutomaticInputFailoverSettings interface{} `json:"automaticInputFailoverSettings"`
	// `CfnChannel.InputAttachmentProperty.InputAttachmentName`.
	InputAttachmentName *string `json:"inputAttachmentName"`
	// `CfnChannel.InputAttachmentProperty.InputId`.
	InputId *string `json:"inputId"`
	// `CfnChannel.InputAttachmentProperty.InputSettings`.
	InputSettings interface{} `json:"inputSettings"`
}

type CfnChannel_InputChannelLevelProperty struct {
	// `CfnChannel.InputChannelLevelProperty.Gain`.
	Gain *float64 `json:"gain"`
	// `CfnChannel.InputChannelLevelProperty.InputChannel`.
	InputChannel *float64 `json:"inputChannel"`
}

type CfnChannel_InputLocationProperty struct {
	// `CfnChannel.InputLocationProperty.PasswordParam`.
	PasswordParam *string `json:"passwordParam"`
	// `CfnChannel.InputLocationProperty.Uri`.
	Uri *string `json:"uri"`
	// `CfnChannel.InputLocationProperty.Username`.
	Username *string `json:"username"`
}

type CfnChannel_InputLossBehaviorProperty struct {
	// `CfnChannel.InputLossBehaviorProperty.BlackFrameMsec`.
	BlackFrameMsec *float64 `json:"blackFrameMsec"`
	// `CfnChannel.InputLossBehaviorProperty.InputLossImageColor`.
	InputLossImageColor *string `json:"inputLossImageColor"`
	// `CfnChannel.InputLossBehaviorProperty.InputLossImageSlate`.
	InputLossImageSlate interface{} `json:"inputLossImageSlate"`
	// `CfnChannel.InputLossBehaviorProperty.InputLossImageType`.
	InputLossImageType *string `json:"inputLossImageType"`
	// `CfnChannel.InputLossBehaviorProperty.RepeatFrameMsec`.
	RepeatFrameMsec *float64 `json:"repeatFrameMsec"`
}

type CfnChannel_InputLossFailoverSettingsProperty struct {
	// `CfnChannel.InputLossFailoverSettingsProperty.InputLossThresholdMsec`.
	InputLossThresholdMsec *float64 `json:"inputLossThresholdMsec"`
}

type CfnChannel_InputSettingsProperty struct {
	// `CfnChannel.InputSettingsProperty.AudioSelectors`.
	AudioSelectors interface{} `json:"audioSelectors"`
	// `CfnChannel.InputSettingsProperty.CaptionSelectors`.
	CaptionSelectors interface{} `json:"captionSelectors"`
	// `CfnChannel.InputSettingsProperty.DeblockFilter`.
	DeblockFilter *string `json:"deblockFilter"`
	// `CfnChannel.InputSettingsProperty.DenoiseFilter`.
	DenoiseFilter *string `json:"denoiseFilter"`
	// `CfnChannel.InputSettingsProperty.FilterStrength`.
	FilterStrength *float64 `json:"filterStrength"`
	// `CfnChannel.InputSettingsProperty.InputFilter`.
	InputFilter *string `json:"inputFilter"`
	// `CfnChannel.InputSettingsProperty.NetworkInputSettings`.
	NetworkInputSettings interface{} `json:"networkInputSettings"`
	// `CfnChannel.InputSettingsProperty.Smpte2038DataPreference`.
	Smpte2038DataPreference *string `json:"smpte2038DataPreference"`
	// `CfnChannel.InputSettingsProperty.SourceEndBehavior`.
	SourceEndBehavior *string `json:"sourceEndBehavior"`
	// `CfnChannel.InputSettingsProperty.VideoSelector`.
	VideoSelector interface{} `json:"videoSelector"`
}

type CfnChannel_InputSpecificationProperty struct {
	// `CfnChannel.InputSpecificationProperty.Codec`.
	Codec *string `json:"codec"`
	// `CfnChannel.InputSpecificationProperty.MaximumBitrate`.
	MaximumBitrate *string `json:"maximumBitrate"`
	// `CfnChannel.InputSpecificationProperty.Resolution`.
	Resolution *string `json:"resolution"`
}

type CfnChannel_KeyProviderSettingsProperty struct {
	// `CfnChannel.KeyProviderSettingsProperty.StaticKeySettings`.
	StaticKeySettings interface{} `json:"staticKeySettings"`
}

type CfnChannel_M2tsSettingsProperty struct {
	// `CfnChannel.M2tsSettingsProperty.AbsentInputAudioBehavior`.
	AbsentInputAudioBehavior *string `json:"absentInputAudioBehavior"`
	// `CfnChannel.M2tsSettingsProperty.Arib`.
	Arib *string `json:"arib"`
	// `CfnChannel.M2tsSettingsProperty.AribCaptionsPid`.
	AribCaptionsPid *string `json:"aribCaptionsPid"`
	// `CfnChannel.M2tsSettingsProperty.AribCaptionsPidControl`.
	AribCaptionsPidControl *string `json:"aribCaptionsPidControl"`
	// `CfnChannel.M2tsSettingsProperty.AudioBufferModel`.
	AudioBufferModel *string `json:"audioBufferModel"`
	// `CfnChannel.M2tsSettingsProperty.AudioFramesPerPes`.
	AudioFramesPerPes *float64 `json:"audioFramesPerPes"`
	// `CfnChannel.M2tsSettingsProperty.AudioPids`.
	AudioPids *string `json:"audioPids"`
	// `CfnChannel.M2tsSettingsProperty.AudioStreamType`.
	AudioStreamType *string `json:"audioStreamType"`
	// `CfnChannel.M2tsSettingsProperty.Bitrate`.
	Bitrate *float64 `json:"bitrate"`
	// `CfnChannel.M2tsSettingsProperty.BufferModel`.
	BufferModel *string `json:"bufferModel"`
	// `CfnChannel.M2tsSettingsProperty.CcDescriptor`.
	CcDescriptor *string `json:"ccDescriptor"`
	// `CfnChannel.M2tsSettingsProperty.DvbNitSettings`.
	DvbNitSettings interface{} `json:"dvbNitSettings"`
	// `CfnChannel.M2tsSettingsProperty.DvbSdtSettings`.
	DvbSdtSettings interface{} `json:"dvbSdtSettings"`
	// `CfnChannel.M2tsSettingsProperty.DvbSubPids`.
	DvbSubPids *string `json:"dvbSubPids"`
	// `CfnChannel.M2tsSettingsProperty.DvbTdtSettings`.
	DvbTdtSettings interface{} `json:"dvbTdtSettings"`
	// `CfnChannel.M2tsSettingsProperty.DvbTeletextPid`.
	DvbTeletextPid *string `json:"dvbTeletextPid"`
	// `CfnChannel.M2tsSettingsProperty.Ebif`.
	Ebif *string `json:"ebif"`
	// `CfnChannel.M2tsSettingsProperty.EbpAudioInterval`.
	EbpAudioInterval *string `json:"ebpAudioInterval"`
	// `CfnChannel.M2tsSettingsProperty.EbpLookaheadMs`.
	EbpLookaheadMs *float64 `json:"ebpLookaheadMs"`
	// `CfnChannel.M2tsSettingsProperty.EbpPlacement`.
	EbpPlacement *string `json:"ebpPlacement"`
	// `CfnChannel.M2tsSettingsProperty.EcmPid`.
	EcmPid *string `json:"ecmPid"`
	// `CfnChannel.M2tsSettingsProperty.EsRateInPes`.
	EsRateInPes *string `json:"esRateInPes"`
	// `CfnChannel.M2tsSettingsProperty.EtvPlatformPid`.
	EtvPlatformPid *string `json:"etvPlatformPid"`
	// `CfnChannel.M2tsSettingsProperty.EtvSignalPid`.
	EtvSignalPid *string `json:"etvSignalPid"`
	// `CfnChannel.M2tsSettingsProperty.FragmentTime`.
	FragmentTime *float64 `json:"fragmentTime"`
	// `CfnChannel.M2tsSettingsProperty.Klv`.
	Klv *string `json:"klv"`
	// `CfnChannel.M2tsSettingsProperty.KlvDataPids`.
	KlvDataPids *string `json:"klvDataPids"`
	// `CfnChannel.M2tsSettingsProperty.NielsenId3Behavior`.
	NielsenId3Behavior *string `json:"nielsenId3Behavior"`
	// `CfnChannel.M2tsSettingsProperty.NullPacketBitrate`.
	NullPacketBitrate *float64 `json:"nullPacketBitrate"`
	// `CfnChannel.M2tsSettingsProperty.PatInterval`.
	PatInterval *float64 `json:"patInterval"`
	// `CfnChannel.M2tsSettingsProperty.PcrControl`.
	PcrControl *string `json:"pcrControl"`
	// `CfnChannel.M2tsSettingsProperty.PcrPeriod`.
	PcrPeriod *float64 `json:"pcrPeriod"`
	// `CfnChannel.M2tsSettingsProperty.PcrPid`.
	PcrPid *string `json:"pcrPid"`
	// `CfnChannel.M2tsSettingsProperty.PmtInterval`.
	PmtInterval *float64 `json:"pmtInterval"`
	// `CfnChannel.M2tsSettingsProperty.PmtPid`.
	PmtPid *string `json:"pmtPid"`
	// `CfnChannel.M2tsSettingsProperty.ProgramNum`.
	ProgramNum *float64 `json:"programNum"`
	// `CfnChannel.M2tsSettingsProperty.RateMode`.
	RateMode *string `json:"rateMode"`
	// `CfnChannel.M2tsSettingsProperty.Scte27Pids`.
	Scte27Pids *string `json:"scte27Pids"`
	// `CfnChannel.M2tsSettingsProperty.Scte35Control`.
	Scte35Control *string `json:"scte35Control"`
	// `CfnChannel.M2tsSettingsProperty.Scte35Pid`.
	Scte35Pid *string `json:"scte35Pid"`
	// `CfnChannel.M2tsSettingsProperty.SegmentationMarkers`.
	SegmentationMarkers *string `json:"segmentationMarkers"`
	// `CfnChannel.M2tsSettingsProperty.SegmentationStyle`.
	SegmentationStyle *string `json:"segmentationStyle"`
	// `CfnChannel.M2tsSettingsProperty.SegmentationTime`.
	SegmentationTime *float64 `json:"segmentationTime"`
	// `CfnChannel.M2tsSettingsProperty.TimedMetadataBehavior`.
	TimedMetadataBehavior *string `json:"timedMetadataBehavior"`
	// `CfnChannel.M2tsSettingsProperty.TimedMetadataPid`.
	TimedMetadataPid *string `json:"timedMetadataPid"`
	// `CfnChannel.M2tsSettingsProperty.TransportStreamId`.
	TransportStreamId *float64 `json:"transportStreamId"`
	// `CfnChannel.M2tsSettingsProperty.VideoPid`.
	VideoPid *string `json:"videoPid"`
}

type CfnChannel_M3u8SettingsProperty struct {
	// `CfnChannel.M3u8SettingsProperty.AudioFramesPerPes`.
	AudioFramesPerPes *float64 `json:"audioFramesPerPes"`
	// `CfnChannel.M3u8SettingsProperty.AudioPids`.
	AudioPids *string `json:"audioPids"`
	// `CfnChannel.M3u8SettingsProperty.EcmPid`.
	EcmPid *string `json:"ecmPid"`
	// `CfnChannel.M3u8SettingsProperty.NielsenId3Behavior`.
	NielsenId3Behavior *string `json:"nielsenId3Behavior"`
	// `CfnChannel.M3u8SettingsProperty.PatInterval`.
	PatInterval *float64 `json:"patInterval"`
	// `CfnChannel.M3u8SettingsProperty.PcrControl`.
	PcrControl *string `json:"pcrControl"`
	// `CfnChannel.M3u8SettingsProperty.PcrPeriod`.
	PcrPeriod *float64 `json:"pcrPeriod"`
	// `CfnChannel.M3u8SettingsProperty.PcrPid`.
	PcrPid *string `json:"pcrPid"`
	// `CfnChannel.M3u8SettingsProperty.PmtInterval`.
	PmtInterval *float64 `json:"pmtInterval"`
	// `CfnChannel.M3u8SettingsProperty.PmtPid`.
	PmtPid *string `json:"pmtPid"`
	// `CfnChannel.M3u8SettingsProperty.ProgramNum`.
	ProgramNum *float64 `json:"programNum"`
	// `CfnChannel.M3u8SettingsProperty.Scte35Behavior`.
	Scte35Behavior *string `json:"scte35Behavior"`
	// `CfnChannel.M3u8SettingsProperty.Scte35Pid`.
	Scte35Pid *string `json:"scte35Pid"`
	// `CfnChannel.M3u8SettingsProperty.TimedMetadataBehavior`.
	TimedMetadataBehavior *string `json:"timedMetadataBehavior"`
	// `CfnChannel.M3u8SettingsProperty.TimedMetadataPid`.
	TimedMetadataPid *string `json:"timedMetadataPid"`
	// `CfnChannel.M3u8SettingsProperty.TransportStreamId`.
	TransportStreamId *float64 `json:"transportStreamId"`
	// `CfnChannel.M3u8SettingsProperty.VideoPid`.
	VideoPid *string `json:"videoPid"`
}

type CfnChannel_MediaPackageGroupSettingsProperty struct {
	// `CfnChannel.MediaPackageGroupSettingsProperty.Destination`.
	Destination interface{} `json:"destination"`
}

type CfnChannel_MediaPackageOutputDestinationSettingsProperty struct {
	// `CfnChannel.MediaPackageOutputDestinationSettingsProperty.ChannelId`.
	ChannelId *string `json:"channelId"`
}

type CfnChannel_MediaPackageOutputSettingsProperty struct {
}

type CfnChannel_MotionGraphicsConfigurationProperty struct {
	// `CfnChannel.MotionGraphicsConfigurationProperty.MotionGraphicsInsertion`.
	MotionGraphicsInsertion *string `json:"motionGraphicsInsertion"`
	// `CfnChannel.MotionGraphicsConfigurationProperty.MotionGraphicsSettings`.
	MotionGraphicsSettings interface{} `json:"motionGraphicsSettings"`
}

type CfnChannel_MotionGraphicsSettingsProperty struct {
	// `CfnChannel.MotionGraphicsSettingsProperty.HtmlMotionGraphicsSettings`.
	HtmlMotionGraphicsSettings interface{} `json:"htmlMotionGraphicsSettings"`
}

type CfnChannel_Mp2SettingsProperty struct {
	// `CfnChannel.Mp2SettingsProperty.Bitrate`.
	Bitrate *float64 `json:"bitrate"`
	// `CfnChannel.Mp2SettingsProperty.CodingMode`.
	CodingMode *string `json:"codingMode"`
	// `CfnChannel.Mp2SettingsProperty.SampleRate`.
	SampleRate *float64 `json:"sampleRate"`
}

type CfnChannel_Mpeg2FilterSettingsProperty struct {
	// `CfnChannel.Mpeg2FilterSettingsProperty.TemporalFilterSettings`.
	TemporalFilterSettings interface{} `json:"temporalFilterSettings"`
}

type CfnChannel_Mpeg2SettingsProperty struct {
	// `CfnChannel.Mpeg2SettingsProperty.AdaptiveQuantization`.
	AdaptiveQuantization *string `json:"adaptiveQuantization"`
	// `CfnChannel.Mpeg2SettingsProperty.AfdSignaling`.
	AfdSignaling *string `json:"afdSignaling"`
	// `CfnChannel.Mpeg2SettingsProperty.ColorMetadata`.
	ColorMetadata *string `json:"colorMetadata"`
	// `CfnChannel.Mpeg2SettingsProperty.ColorSpace`.
	ColorSpace *string `json:"colorSpace"`
	// `CfnChannel.Mpeg2SettingsProperty.DisplayAspectRatio`.
	DisplayAspectRatio *string `json:"displayAspectRatio"`
	// `CfnChannel.Mpeg2SettingsProperty.FilterSettings`.
	FilterSettings interface{} `json:"filterSettings"`
	// `CfnChannel.Mpeg2SettingsProperty.FixedAfd`.
	FixedAfd *string `json:"fixedAfd"`
	// `CfnChannel.Mpeg2SettingsProperty.FramerateDenominator`.
	FramerateDenominator *float64 `json:"framerateDenominator"`
	// `CfnChannel.Mpeg2SettingsProperty.FramerateNumerator`.
	FramerateNumerator *float64 `json:"framerateNumerator"`
	// `CfnChannel.Mpeg2SettingsProperty.GopClosedCadence`.
	GopClosedCadence *float64 `json:"gopClosedCadence"`
	// `CfnChannel.Mpeg2SettingsProperty.GopNumBFrames`.
	GopNumBFrames *float64 `json:"gopNumBFrames"`
	// `CfnChannel.Mpeg2SettingsProperty.GopSize`.
	GopSize *float64 `json:"gopSize"`
	// `CfnChannel.Mpeg2SettingsProperty.GopSizeUnits`.
	GopSizeUnits *string `json:"gopSizeUnits"`
	// `CfnChannel.Mpeg2SettingsProperty.ScanType`.
	ScanType *string `json:"scanType"`
	// `CfnChannel.Mpeg2SettingsProperty.SubgopLength`.
	SubgopLength *string `json:"subgopLength"`
	// `CfnChannel.Mpeg2SettingsProperty.TimecodeInsertion`.
	TimecodeInsertion *string `json:"timecodeInsertion"`
}

type CfnChannel_MsSmoothGroupSettingsProperty struct {
	// `CfnChannel.MsSmoothGroupSettingsProperty.AcquisitionPointId`.
	AcquisitionPointId *string `json:"acquisitionPointId"`
	// `CfnChannel.MsSmoothGroupSettingsProperty.AudioOnlyTimecodeControl`.
	AudioOnlyTimecodeControl *string `json:"audioOnlyTimecodeControl"`
	// `CfnChannel.MsSmoothGroupSettingsProperty.CertificateMode`.
	CertificateMode *string `json:"certificateMode"`
	// `CfnChannel.MsSmoothGroupSettingsProperty.ConnectionRetryInterval`.
	ConnectionRetryInterval *float64 `json:"connectionRetryInterval"`
	// `CfnChannel.MsSmoothGroupSettingsProperty.Destination`.
	Destination interface{} `json:"destination"`
	// `CfnChannel.MsSmoothGroupSettingsProperty.EventId`.
	EventId *string `json:"eventId"`
	// `CfnChannel.MsSmoothGroupSettingsProperty.EventIdMode`.
	EventIdMode *string `json:"eventIdMode"`
	// `CfnChannel.MsSmoothGroupSettingsProperty.EventStopBehavior`.
	EventStopBehavior *string `json:"eventStopBehavior"`
	// `CfnChannel.MsSmoothGroupSettingsProperty.FilecacheDuration`.
	FilecacheDuration *float64 `json:"filecacheDuration"`
	// `CfnChannel.MsSmoothGroupSettingsProperty.FragmentLength`.
	FragmentLength *float64 `json:"fragmentLength"`
	// `CfnChannel.MsSmoothGroupSettingsProperty.InputLossAction`.
	InputLossAction *string `json:"inputLossAction"`
	// `CfnChannel.MsSmoothGroupSettingsProperty.NumRetries`.
	NumRetries *float64 `json:"numRetries"`
	// `CfnChannel.MsSmoothGroupSettingsProperty.RestartDelay`.
	RestartDelay *float64 `json:"restartDelay"`
	// `CfnChannel.MsSmoothGroupSettingsProperty.SegmentationMode`.
	SegmentationMode *string `json:"segmentationMode"`
	// `CfnChannel.MsSmoothGroupSettingsProperty.SendDelayMs`.
	SendDelayMs *float64 `json:"sendDelayMs"`
	// `CfnChannel.MsSmoothGroupSettingsProperty.SparseTrackType`.
	SparseTrackType *string `json:"sparseTrackType"`
	// `CfnChannel.MsSmoothGroupSettingsProperty.StreamManifestBehavior`.
	StreamManifestBehavior *string `json:"streamManifestBehavior"`
	// `CfnChannel.MsSmoothGroupSettingsProperty.TimestampOffset`.
	TimestampOffset *string `json:"timestampOffset"`
	// `CfnChannel.MsSmoothGroupSettingsProperty.TimestampOffsetMode`.
	TimestampOffsetMode *string `json:"timestampOffsetMode"`
}

type CfnChannel_MsSmoothOutputSettingsProperty struct {
	// `CfnChannel.MsSmoothOutputSettingsProperty.H265PackagingType`.
	H265PackagingType *string `json:"h265PackagingType"`
	// `CfnChannel.MsSmoothOutputSettingsProperty.NameModifier`.
	NameModifier *string `json:"nameModifier"`
}

type CfnChannel_MultiplexGroupSettingsProperty struct {
}

type CfnChannel_MultiplexOutputSettingsProperty struct {
	// `CfnChannel.MultiplexOutputSettingsProperty.Destination`.
	Destination interface{} `json:"destination"`
}

type CfnChannel_MultiplexProgramChannelDestinationSettingsProperty struct {
	// `CfnChannel.MultiplexProgramChannelDestinationSettingsProperty.MultiplexId`.
	MultiplexId *string `json:"multiplexId"`
	// `CfnChannel.MultiplexProgramChannelDestinationSettingsProperty.ProgramName`.
	ProgramName *string `json:"programName"`
}

type CfnChannel_NetworkInputSettingsProperty struct {
	// `CfnChannel.NetworkInputSettingsProperty.HlsInputSettings`.
	HlsInputSettings interface{} `json:"hlsInputSettings"`
	// `CfnChannel.NetworkInputSettingsProperty.ServerValidation`.
	ServerValidation *string `json:"serverValidation"`
}

type CfnChannel_NielsenConfigurationProperty struct {
	// `CfnChannel.NielsenConfigurationProperty.DistributorId`.
	DistributorId *string `json:"distributorId"`
	// `CfnChannel.NielsenConfigurationProperty.NielsenPcmToId3Tagging`.
	NielsenPcmToId3Tagging *string `json:"nielsenPcmToId3Tagging"`
}

type CfnChannel_OutputDestinationProperty struct {
	// `CfnChannel.OutputDestinationProperty.Id`.
	Id *string `json:"id"`
	// `CfnChannel.OutputDestinationProperty.MediaPackageSettings`.
	MediaPackageSettings interface{} `json:"mediaPackageSettings"`
	// `CfnChannel.OutputDestinationProperty.MultiplexSettings`.
	MultiplexSettings interface{} `json:"multiplexSettings"`
	// `CfnChannel.OutputDestinationProperty.Settings`.
	Settings interface{} `json:"settings"`
}

type CfnChannel_OutputDestinationSettingsProperty struct {
	// `CfnChannel.OutputDestinationSettingsProperty.PasswordParam`.
	PasswordParam *string `json:"passwordParam"`
	// `CfnChannel.OutputDestinationSettingsProperty.StreamName`.
	StreamName *string `json:"streamName"`
	// `CfnChannel.OutputDestinationSettingsProperty.Url`.
	Url *string `json:"url"`
	// `CfnChannel.OutputDestinationSettingsProperty.Username`.
	Username *string `json:"username"`
}

type CfnChannel_OutputGroupProperty struct {
	// `CfnChannel.OutputGroupProperty.Name`.
	Name *string `json:"name"`
	// `CfnChannel.OutputGroupProperty.OutputGroupSettings`.
	OutputGroupSettings interface{} `json:"outputGroupSettings"`
	// `CfnChannel.OutputGroupProperty.Outputs`.
	Outputs interface{} `json:"outputs"`
}

type CfnChannel_OutputGroupSettingsProperty struct {
	// `CfnChannel.OutputGroupSettingsProperty.ArchiveGroupSettings`.
	ArchiveGroupSettings interface{} `json:"archiveGroupSettings"`
	// `CfnChannel.OutputGroupSettingsProperty.FrameCaptureGroupSettings`.
	FrameCaptureGroupSettings interface{} `json:"frameCaptureGroupSettings"`
	// `CfnChannel.OutputGroupSettingsProperty.HlsGroupSettings`.
	HlsGroupSettings interface{} `json:"hlsGroupSettings"`
	// `CfnChannel.OutputGroupSettingsProperty.MediaPackageGroupSettings`.
	MediaPackageGroupSettings interface{} `json:"mediaPackageGroupSettings"`
	// `CfnChannel.OutputGroupSettingsProperty.MsSmoothGroupSettings`.
	MsSmoothGroupSettings interface{} `json:"msSmoothGroupSettings"`
	// `CfnChannel.OutputGroupSettingsProperty.MultiplexGroupSettings`.
	MultiplexGroupSettings interface{} `json:"multiplexGroupSettings"`
	// `CfnChannel.OutputGroupSettingsProperty.RtmpGroupSettings`.
	RtmpGroupSettings interface{} `json:"rtmpGroupSettings"`
	// `CfnChannel.OutputGroupSettingsProperty.UdpGroupSettings`.
	UdpGroupSettings interface{} `json:"udpGroupSettings"`
}

type CfnChannel_OutputLocationRefProperty struct {
	// `CfnChannel.OutputLocationRefProperty.DestinationRefId`.
	DestinationRefId *string `json:"destinationRefId"`
}

type CfnChannel_OutputProperty struct {
	// `CfnChannel.OutputProperty.AudioDescriptionNames`.
	AudioDescriptionNames *[]*string `json:"audioDescriptionNames"`
	// `CfnChannel.OutputProperty.CaptionDescriptionNames`.
	CaptionDescriptionNames *[]*string `json:"captionDescriptionNames"`
	// `CfnChannel.OutputProperty.OutputName`.
	OutputName *string `json:"outputName"`
	// `CfnChannel.OutputProperty.OutputSettings`.
	OutputSettings interface{} `json:"outputSettings"`
	// `CfnChannel.OutputProperty.VideoDescriptionName`.
	VideoDescriptionName *string `json:"videoDescriptionName"`
}

type CfnChannel_OutputSettingsProperty struct {
	// `CfnChannel.OutputSettingsProperty.ArchiveOutputSettings`.
	ArchiveOutputSettings interface{} `json:"archiveOutputSettings"`
	// `CfnChannel.OutputSettingsProperty.FrameCaptureOutputSettings`.
	FrameCaptureOutputSettings interface{} `json:"frameCaptureOutputSettings"`
	// `CfnChannel.OutputSettingsProperty.HlsOutputSettings`.
	HlsOutputSettings interface{} `json:"hlsOutputSettings"`
	// `CfnChannel.OutputSettingsProperty.MediaPackageOutputSettings`.
	MediaPackageOutputSettings interface{} `json:"mediaPackageOutputSettings"`
	// `CfnChannel.OutputSettingsProperty.MsSmoothOutputSettings`.
	MsSmoothOutputSettings interface{} `json:"msSmoothOutputSettings"`
	// `CfnChannel.OutputSettingsProperty.MultiplexOutputSettings`.
	MultiplexOutputSettings interface{} `json:"multiplexOutputSettings"`
	// `CfnChannel.OutputSettingsProperty.RtmpOutputSettings`.
	RtmpOutputSettings interface{} `json:"rtmpOutputSettings"`
	// `CfnChannel.OutputSettingsProperty.UdpOutputSettings`.
	UdpOutputSettings interface{} `json:"udpOutputSettings"`
}

type CfnChannel_PassThroughSettingsProperty struct {
}

type CfnChannel_RawSettingsProperty struct {
}

type CfnChannel_Rec601SettingsProperty struct {
}

type CfnChannel_Rec709SettingsProperty struct {
}

type CfnChannel_RemixSettingsProperty struct {
	// `CfnChannel.RemixSettingsProperty.ChannelMappings`.
	ChannelMappings interface{} `json:"channelMappings"`
	// `CfnChannel.RemixSettingsProperty.ChannelsIn`.
	ChannelsIn *float64 `json:"channelsIn"`
	// `CfnChannel.RemixSettingsProperty.ChannelsOut`.
	ChannelsOut *float64 `json:"channelsOut"`
}

type CfnChannel_RtmpCaptionInfoDestinationSettingsProperty struct {
}

type CfnChannel_RtmpGroupSettingsProperty struct {
	// `CfnChannel.RtmpGroupSettingsProperty.AdMarkers`.
	AdMarkers *[]*string `json:"adMarkers"`
	// `CfnChannel.RtmpGroupSettingsProperty.AuthenticationScheme`.
	AuthenticationScheme *string `json:"authenticationScheme"`
	// `CfnChannel.RtmpGroupSettingsProperty.CacheFullBehavior`.
	CacheFullBehavior *string `json:"cacheFullBehavior"`
	// `CfnChannel.RtmpGroupSettingsProperty.CacheLength`.
	CacheLength *float64 `json:"cacheLength"`
	// `CfnChannel.RtmpGroupSettingsProperty.CaptionData`.
	CaptionData *string `json:"captionData"`
	// `CfnChannel.RtmpGroupSettingsProperty.InputLossAction`.
	InputLossAction *string `json:"inputLossAction"`
	// `CfnChannel.RtmpGroupSettingsProperty.RestartDelay`.
	RestartDelay *float64 `json:"restartDelay"`
}

type CfnChannel_RtmpOutputSettingsProperty struct {
	// `CfnChannel.RtmpOutputSettingsProperty.CertificateMode`.
	CertificateMode *string `json:"certificateMode"`
	// `CfnChannel.RtmpOutputSettingsProperty.ConnectionRetryInterval`.
	ConnectionRetryInterval *float64 `json:"connectionRetryInterval"`
	// `CfnChannel.RtmpOutputSettingsProperty.Destination`.
	Destination interface{} `json:"destination"`
	// `CfnChannel.RtmpOutputSettingsProperty.NumRetries`.
	NumRetries *float64 `json:"numRetries"`
}

type CfnChannel_Scte20PlusEmbeddedDestinationSettingsProperty struct {
}

type CfnChannel_Scte20SourceSettingsProperty struct {
	// `CfnChannel.Scte20SourceSettingsProperty.Convert608To708`.
	Convert608To708 *string `json:"convert608To708"`
	// `CfnChannel.Scte20SourceSettingsProperty.Source608ChannelNumber`.
	Source608ChannelNumber *float64 `json:"source608ChannelNumber"`
}

type CfnChannel_Scte27DestinationSettingsProperty struct {
}

type CfnChannel_Scte27SourceSettingsProperty struct {
	// `CfnChannel.Scte27SourceSettingsProperty.Pid`.
	Pid *float64 `json:"pid"`
}

type CfnChannel_Scte35SpliceInsertProperty struct {
	// `CfnChannel.Scte35SpliceInsertProperty.AdAvailOffset`.
	AdAvailOffset *float64 `json:"adAvailOffset"`
	// `CfnChannel.Scte35SpliceInsertProperty.NoRegionalBlackoutFlag`.
	NoRegionalBlackoutFlag *string `json:"noRegionalBlackoutFlag"`
	// `CfnChannel.Scte35SpliceInsertProperty.WebDeliveryAllowedFlag`.
	WebDeliveryAllowedFlag *string `json:"webDeliveryAllowedFlag"`
}

type CfnChannel_Scte35TimeSignalAposProperty struct {
	// `CfnChannel.Scte35TimeSignalAposProperty.AdAvailOffset`.
	AdAvailOffset *float64 `json:"adAvailOffset"`
	// `CfnChannel.Scte35TimeSignalAposProperty.NoRegionalBlackoutFlag`.
	NoRegionalBlackoutFlag *string `json:"noRegionalBlackoutFlag"`
	// `CfnChannel.Scte35TimeSignalAposProperty.WebDeliveryAllowedFlag`.
	WebDeliveryAllowedFlag *string `json:"webDeliveryAllowedFlag"`
}

type CfnChannel_SmpteTtDestinationSettingsProperty struct {
}

type CfnChannel_StandardHlsSettingsProperty struct {
	// `CfnChannel.StandardHlsSettingsProperty.AudioRenditionSets`.
	AudioRenditionSets *string `json:"audioRenditionSets"`
	// `CfnChannel.StandardHlsSettingsProperty.M3u8Settings`.
	M3U8Settings interface{} `json:"m3U8Settings"`
}

type CfnChannel_StaticKeySettingsProperty struct {
	// `CfnChannel.StaticKeySettingsProperty.KeyProviderServer`.
	KeyProviderServer interface{} `json:"keyProviderServer"`
	// `CfnChannel.StaticKeySettingsProperty.StaticKeyValue`.
	StaticKeyValue *string `json:"staticKeyValue"`
}

type CfnChannel_TeletextDestinationSettingsProperty struct {
}

type CfnChannel_TeletextSourceSettingsProperty struct {
	// `CfnChannel.TeletextSourceSettingsProperty.OutputRectangle`.
	OutputRectangle interface{} `json:"outputRectangle"`
	// `CfnChannel.TeletextSourceSettingsProperty.PageNumber`.
	PageNumber *string `json:"pageNumber"`
}

type CfnChannel_TemporalFilterSettingsProperty struct {
	// `CfnChannel.TemporalFilterSettingsProperty.PostFilterSharpening`.
	PostFilterSharpening *string `json:"postFilterSharpening"`
	// `CfnChannel.TemporalFilterSettingsProperty.Strength`.
	Strength *string `json:"strength"`
}

type CfnChannel_TimecodeConfigProperty struct {
	// `CfnChannel.TimecodeConfigProperty.Source`.
	Source *string `json:"source"`
	// `CfnChannel.TimecodeConfigProperty.SyncThreshold`.
	SyncThreshold *float64 `json:"syncThreshold"`
}

type CfnChannel_TtmlDestinationSettingsProperty struct {
	// `CfnChannel.TtmlDestinationSettingsProperty.StyleControl`.
	StyleControl *string `json:"styleControl"`
}

type CfnChannel_UdpContainerSettingsProperty struct {
	// `CfnChannel.UdpContainerSettingsProperty.M2tsSettings`.
	M2TsSettings interface{} `json:"m2TsSettings"`
}

type CfnChannel_UdpGroupSettingsProperty struct {
	// `CfnChannel.UdpGroupSettingsProperty.InputLossAction`.
	InputLossAction *string `json:"inputLossAction"`
	// `CfnChannel.UdpGroupSettingsProperty.TimedMetadataId3Frame`.
	TimedMetadataId3Frame *string `json:"timedMetadataId3Frame"`
	// `CfnChannel.UdpGroupSettingsProperty.TimedMetadataId3Period`.
	TimedMetadataId3Period *float64 `json:"timedMetadataId3Period"`
}

type CfnChannel_UdpOutputSettingsProperty struct {
	// `CfnChannel.UdpOutputSettingsProperty.BufferMsec`.
	BufferMsec *float64 `json:"bufferMsec"`
	// `CfnChannel.UdpOutputSettingsProperty.ContainerSettings`.
	ContainerSettings interface{} `json:"containerSettings"`
	// `CfnChannel.UdpOutputSettingsProperty.Destination`.
	Destination interface{} `json:"destination"`
	// `CfnChannel.UdpOutputSettingsProperty.FecOutputSettings`.
	FecOutputSettings interface{} `json:"fecOutputSettings"`
}

type CfnChannel_VideoBlackFailoverSettingsProperty struct {
	// `CfnChannel.VideoBlackFailoverSettingsProperty.BlackDetectThreshold`.
	BlackDetectThreshold *float64 `json:"blackDetectThreshold"`
	// `CfnChannel.VideoBlackFailoverSettingsProperty.VideoBlackThresholdMsec`.
	VideoBlackThresholdMsec *float64 `json:"videoBlackThresholdMsec"`
}

type CfnChannel_VideoCodecSettingsProperty struct {
	// `CfnChannel.VideoCodecSettingsProperty.FrameCaptureSettings`.
	FrameCaptureSettings interface{} `json:"frameCaptureSettings"`
	// `CfnChannel.VideoCodecSettingsProperty.H264Settings`.
	H264Settings interface{} `json:"h264Settings"`
	// `CfnChannel.VideoCodecSettingsProperty.H265Settings`.
	H265Settings interface{} `json:"h265Settings"`
	// `CfnChannel.VideoCodecSettingsProperty.Mpeg2Settings`.
	Mpeg2Settings interface{} `json:"mpeg2Settings"`
}

type CfnChannel_VideoDescriptionProperty struct {
	// `CfnChannel.VideoDescriptionProperty.CodecSettings`.
	CodecSettings interface{} `json:"codecSettings"`
	// `CfnChannel.VideoDescriptionProperty.Height`.
	Height *float64 `json:"height"`
	// `CfnChannel.VideoDescriptionProperty.Name`.
	Name *string `json:"name"`
	// `CfnChannel.VideoDescriptionProperty.RespondToAfd`.
	RespondToAfd *string `json:"respondToAfd"`
	// `CfnChannel.VideoDescriptionProperty.ScalingBehavior`.
	ScalingBehavior *string `json:"scalingBehavior"`
	// `CfnChannel.VideoDescriptionProperty.Sharpness`.
	Sharpness *float64 `json:"sharpness"`
	// `CfnChannel.VideoDescriptionProperty.Width`.
	Width *float64 `json:"width"`
}

type CfnChannel_VideoSelectorColorSpaceSettingsProperty struct {
	// `CfnChannel.VideoSelectorColorSpaceSettingsProperty.Hdr10Settings`.
	Hdr10Settings interface{} `json:"hdr10Settings"`
}

type CfnChannel_VideoSelectorPidProperty struct {
	// `CfnChannel.VideoSelectorPidProperty.Pid`.
	Pid *float64 `json:"pid"`
}

type CfnChannel_VideoSelectorProgramIdProperty struct {
	// `CfnChannel.VideoSelectorProgramIdProperty.ProgramId`.
	ProgramId *float64 `json:"programId"`
}

type CfnChannel_VideoSelectorProperty struct {
	// `CfnChannel.VideoSelectorProperty.ColorSpace`.
	ColorSpace *string `json:"colorSpace"`
	// `CfnChannel.VideoSelectorProperty.ColorSpaceSettings`.
	ColorSpaceSettings interface{} `json:"colorSpaceSettings"`
	// `CfnChannel.VideoSelectorProperty.ColorSpaceUsage`.
	ColorSpaceUsage *string `json:"colorSpaceUsage"`
	// `CfnChannel.VideoSelectorProperty.SelectorSettings`.
	SelectorSettings interface{} `json:"selectorSettings"`
}

type CfnChannel_VideoSelectorSettingsProperty struct {
	// `CfnChannel.VideoSelectorSettingsProperty.VideoSelectorPid`.
	VideoSelectorPid interface{} `json:"videoSelectorPid"`
	// `CfnChannel.VideoSelectorSettingsProperty.VideoSelectorProgramId`.
	VideoSelectorProgramId interface{} `json:"videoSelectorProgramId"`
}

type CfnChannel_VpcOutputSettingsProperty struct {
	// `CfnChannel.VpcOutputSettingsProperty.PublicAddressAllocationIds`.
	PublicAddressAllocationIds *[]*string `json:"publicAddressAllocationIds"`
	// `CfnChannel.VpcOutputSettingsProperty.SecurityGroupIds`.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// `CfnChannel.VpcOutputSettingsProperty.SubnetIds`.
	SubnetIds *[]*string `json:"subnetIds"`
}

type CfnChannel_WavSettingsProperty struct {
	// `CfnChannel.WavSettingsProperty.BitDepth`.
	BitDepth *float64 `json:"bitDepth"`
	// `CfnChannel.WavSettingsProperty.CodingMode`.
	CodingMode *string `json:"codingMode"`
	// `CfnChannel.WavSettingsProperty.SampleRate`.
	SampleRate *float64 `json:"sampleRate"`
}

type CfnChannel_WebvttDestinationSettingsProperty struct {
}

// Properties for defining a `AWS::MediaLive::Channel`.
type CfnChannelProps struct {
	// `AWS::MediaLive::Channel.CdiInputSpecification`.
	CdiInputSpecification interface{} `json:"cdiInputSpecification"`
	// `AWS::MediaLive::Channel.ChannelClass`.
	ChannelClass *string `json:"channelClass"`
	// `AWS::MediaLive::Channel.Destinations`.
	Destinations interface{} `json:"destinations"`
	// `AWS::MediaLive::Channel.EncoderSettings`.
	EncoderSettings interface{} `json:"encoderSettings"`
	// `AWS::MediaLive::Channel.InputAttachments`.
	InputAttachments interface{} `json:"inputAttachments"`
	// `AWS::MediaLive::Channel.InputSpecification`.
	InputSpecification interface{} `json:"inputSpecification"`
	// `AWS::MediaLive::Channel.LogLevel`.
	LogLevel *string `json:"logLevel"`
	// `AWS::MediaLive::Channel.Name`.
	Name *string `json:"name"`
	// `AWS::MediaLive::Channel.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `AWS::MediaLive::Channel.Tags`.
	Tags interface{} `json:"tags"`
	// `AWS::MediaLive::Channel.Vpc`.
	Vpc interface{} `json:"vpc"`
}

// A CloudFormation `AWS::MediaLive::Input`.
type CfnInput interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrDestinations() *[]*string
	AttrSources() *[]*string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Destinations() interface{}
	SetDestinations(val interface{})
	InputDevices() interface{}
	SetInputDevices(val interface{})
	InputSecurityGroups() *[]*string
	SetInputSecurityGroups(val *[]*string)
	LogicalId() *string
	MediaConnectFlows() interface{}
	SetMediaConnectFlows(val interface{})
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	RoleArn() *string
	SetRoleArn(val *string)
	Sources() interface{}
	SetSources(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	Type() *string
	SetType(val *string)
	UpdatedProperites() *map[string]interface{}
	Vpc() interface{}
	SetVpc(val interface{})
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

// The jsii proxy struct for CfnInput
type jsiiProxy_CfnInput struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnInput) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) AttrDestinations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrDestinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) AttrSources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrSources",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_CfnInput) Destinations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) InputDevices() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputDevices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) InputSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputSecurityGroups",
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

func (j *jsiiProxy_CfnInput) MediaConnectFlows() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mediaConnectFlows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) Node() constructs.Node {
	var returns constructs.Node
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

func (j *jsiiProxy_CfnInput) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInput) Sources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sources",
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

func (j *jsiiProxy_CfnInput) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
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

func (j *jsiiProxy_CfnInput) Vpc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaLive::Input`.
func NewCfnInput(scope constructs.Construct, id *string, props *CfnInputProps) CfnInput {
	_init_.Initialize()

	j := jsiiProxy_CfnInput{}

	_jsii_.Create(
		"aws-cdk-lib.aws_medialive.CfnInput",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaLive::Input`.
func NewCfnInput_Override(c CfnInput, scope constructs.Construct, id *string, props *CfnInputProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_medialive.CfnInput",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnInput) SetDestinations(val interface{}) {
	_jsii_.Set(
		j,
		"destinations",
		val,
	)
}

func (j *jsiiProxy_CfnInput) SetInputDevices(val interface{}) {
	_jsii_.Set(
		j,
		"inputDevices",
		val,
	)
}

func (j *jsiiProxy_CfnInput) SetInputSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"inputSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_CfnInput) SetMediaConnectFlows(val interface{}) {
	_jsii_.Set(
		j,
		"mediaConnectFlows",
		val,
	)
}

func (j *jsiiProxy_CfnInput) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnInput) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnInput) SetSources(val interface{}) {
	_jsii_.Set(
		j,
		"sources",
		val,
	)
}

func (j *jsiiProxy_CfnInput) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CfnInput) SetVpc(val interface{}) {
	_jsii_.Set(
		j,
		"vpc",
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
		"aws-cdk-lib.aws_medialive.CfnInput",
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
		"aws-cdk-lib.aws_medialive.CfnInput",
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
func CfnInput_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_medialive.CfnInput",
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
		"aws-cdk-lib.aws_medialive.CfnInput",
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
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnInput) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Experimental.
func (c *jsiiProxy_CfnInput) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnInput_InputDestinationRequestProperty struct {
	// `CfnInput.InputDestinationRequestProperty.StreamName`.
	StreamName *string `json:"streamName"`
}

type CfnInput_InputDeviceRequestProperty struct {
	// `CfnInput.InputDeviceRequestProperty.Id`.
	Id *string `json:"id"`
}

type CfnInput_InputDeviceSettingsProperty struct {
	// `CfnInput.InputDeviceSettingsProperty.Id`.
	Id *string `json:"id"`
}

type CfnInput_InputSourceRequestProperty struct {
	// `CfnInput.InputSourceRequestProperty.PasswordParam`.
	PasswordParam *string `json:"passwordParam"`
	// `CfnInput.InputSourceRequestProperty.Url`.
	Url *string `json:"url"`
	// `CfnInput.InputSourceRequestProperty.Username`.
	Username *string `json:"username"`
}

type CfnInput_InputVpcRequestProperty struct {
	// `CfnInput.InputVpcRequestProperty.SecurityGroupIds`.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// `CfnInput.InputVpcRequestProperty.SubnetIds`.
	SubnetIds *[]*string `json:"subnetIds"`
}

type CfnInput_MediaConnectFlowRequestProperty struct {
	// `CfnInput.MediaConnectFlowRequestProperty.FlowArn`.
	FlowArn *string `json:"flowArn"`
}

// Properties for defining a `AWS::MediaLive::Input`.
type CfnInputProps struct {
	// `AWS::MediaLive::Input.Destinations`.
	Destinations interface{} `json:"destinations"`
	// `AWS::MediaLive::Input.InputDevices`.
	InputDevices interface{} `json:"inputDevices"`
	// `AWS::MediaLive::Input.InputSecurityGroups`.
	InputSecurityGroups *[]*string `json:"inputSecurityGroups"`
	// `AWS::MediaLive::Input.MediaConnectFlows`.
	MediaConnectFlows interface{} `json:"mediaConnectFlows"`
	// `AWS::MediaLive::Input.Name`.
	Name *string `json:"name"`
	// `AWS::MediaLive::Input.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `AWS::MediaLive::Input.Sources`.
	Sources interface{} `json:"sources"`
	// `AWS::MediaLive::Input.Tags`.
	Tags interface{} `json:"tags"`
	// `AWS::MediaLive::Input.Type`.
	Type *string `json:"type"`
	// `AWS::MediaLive::Input.Vpc`.
	Vpc interface{} `json:"vpc"`
}

// A CloudFormation `AWS::MediaLive::InputSecurityGroup`.
type CfnInputSecurityGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	WhitelistRules() interface{}
	SetWhitelistRules(val interface{})
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

// The jsii proxy struct for CfnInputSecurityGroup
type jsiiProxy_CfnInputSecurityGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnInputSecurityGroup) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInputSecurityGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInputSecurityGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInputSecurityGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInputSecurityGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInputSecurityGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInputSecurityGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInputSecurityGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInputSecurityGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInputSecurityGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInputSecurityGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInputSecurityGroup) WhitelistRules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"whitelistRules",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaLive::InputSecurityGroup`.
func NewCfnInputSecurityGroup(scope constructs.Construct, id *string, props *CfnInputSecurityGroupProps) CfnInputSecurityGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnInputSecurityGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_medialive.CfnInputSecurityGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaLive::InputSecurityGroup`.
func NewCfnInputSecurityGroup_Override(c CfnInputSecurityGroup, scope constructs.Construct, id *string, props *CfnInputSecurityGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_medialive.CfnInputSecurityGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnInputSecurityGroup) SetWhitelistRules(val interface{}) {
	_jsii_.Set(
		j,
		"whitelistRules",
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
func CfnInputSecurityGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_medialive.CfnInputSecurityGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnInputSecurityGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_medialive.CfnInputSecurityGroup",
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
func CfnInputSecurityGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_medialive.CfnInputSecurityGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInputSecurityGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_medialive.CfnInputSecurityGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnInputSecurityGroup) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnInputSecurityGroup) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnInputSecurityGroup) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnInputSecurityGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnInputSecurityGroup) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnInputSecurityGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnInputSecurityGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnInputSecurityGroup) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnInputSecurityGroup) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnInputSecurityGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnInputSecurityGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnInputSecurityGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnInputSecurityGroup) ShouldSynthesize() *bool {
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
// Experimental.
func (c *jsiiProxy_CfnInputSecurityGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnInputSecurityGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnInputSecurityGroup_InputWhitelistRuleCidrProperty struct {
	// `CfnInputSecurityGroup.InputWhitelistRuleCidrProperty.Cidr`.
	Cidr *string `json:"cidr"`
}

// Properties for defining a `AWS::MediaLive::InputSecurityGroup`.
type CfnInputSecurityGroupProps struct {
	// `AWS::MediaLive::InputSecurityGroup.Tags`.
	Tags interface{} `json:"tags"`
	// `AWS::MediaLive::InputSecurityGroup.WhitelistRules`.
	WhitelistRules interface{} `json:"whitelistRules"`
}


package awsmediatailor

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediatailor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::MediaTailor::PlaybackConfiguration`.
//
// Adds a new playback configuration to AWS Elemental MediaTailor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationAliases interface{}
//
//   cfnPlaybackConfiguration := awscdk.Aws_mediatailor.NewCfnPlaybackConfiguration(this, jsii.String("MyCfnPlaybackConfiguration"), &cfnPlaybackConfigurationProps{
//   	adDecisionServerUrl: jsii.String("adDecisionServerUrl"),
//   	name: jsii.String("name"),
//   	videoContentSourceUrl: jsii.String("videoContentSourceUrl"),
//
//   	// the properties below are optional
//   	availSuppression: &availSuppressionProperty{
//   		mode: jsii.String("mode"),
//   		value: jsii.String("value"),
//   	},
//   	bumper: &bumperProperty{
//   		endUrl: jsii.String("endUrl"),
//   		startUrl: jsii.String("startUrl"),
//   	},
//   	cdnConfiguration: &cdnConfigurationProperty{
//   		adSegmentUrlPrefix: jsii.String("adSegmentUrlPrefix"),
//   		contentSegmentUrlPrefix: jsii.String("contentSegmentUrlPrefix"),
//   	},
//   	configurationAliases: map[string]interface{}{
//   		"configurationAliasesKey": configurationAliases,
//   	},
//   	dashConfiguration: &dashConfigurationProperty{
//   		manifestEndpointPrefix: jsii.String("manifestEndpointPrefix"),
//   		mpdLocation: jsii.String("mpdLocation"),
//   		originManifestType: jsii.String("originManifestType"),
//   	},
//   	livePreRollConfiguration: &livePreRollConfigurationProperty{
//   		adDecisionServerUrl: jsii.String("adDecisionServerUrl"),
//   		maxDurationSeconds: jsii.Number(123),
//   	},
//   	manifestProcessingRules: &manifestProcessingRulesProperty{
//   		adMarkerPassthrough: &adMarkerPassthroughProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   	personalizationThresholdSeconds: jsii.Number(123),
//   	slateAdUrl: jsii.String("slateAdUrl"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	transcodeProfileName: jsii.String("transcodeProfileName"),
//   })
//
type CfnPlaybackConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The URL for the ad decision server (ADS).
	//
	// This includes the specification of static parameters and placeholders for dynamic parameters. MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing you can provide a static VAST URL. The maximum length is 25,000 characters.
	AdDecisionServerUrl() *string
	SetAdDecisionServerUrl(val *string)
	AttrDashConfigurationManifestEndpointPrefix() *string
	AttrHlsConfigurationManifestEndpointPrefix() *string
	AttrPlaybackConfigurationArn() *string
	AttrPlaybackEndpointPrefix() *string
	AttrSessionInitializationEndpointPrefix() *string
	// The configuration for avail suppression, also known as ad suppression.
	//
	// For more information about ad suppression, see [Ad Suppression](https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html) .
	AvailSuppression() interface{}
	SetAvailSuppression(val interface{})
	// The configuration for bumpers.
	//
	// Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see [Bumpers](https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html) .
	Bumper() interface{}
	SetBumper(val interface{})
	// The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.
	CdnConfiguration() interface{}
	SetCdnConfiguration(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The player parameters and aliases used as dynamic variables during session initialization.
	//
	// For more information, see [Domain Variables](https://docs.aws.amazon.com/mediatailor/latest/ug/variables-domain.html) .
	ConfigurationAliases() interface{}
	SetConfigurationAliases(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The configuration for DASH content.
	DashConfiguration() interface{}
	SetDashConfiguration(val interface{})
	// The configuration for pre-roll ad insertion.
	LivePreRollConfiguration() interface{}
	SetLivePreRollConfiguration(val interface{})
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
	// The configuration for manifest processing rules.
	//
	// Manifest processing rules enable customization of the personalized manifests created by MediaTailor.
	ManifestProcessingRules() interface{}
	SetManifestProcessingRules(val interface{})
	// The identifier for the playback configuration.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Defines the maximum duration of underfilled ad time (in seconds) allowed in an ad break.
	//
	// If the duration of underfilled ad time exceeds the personalization threshold, then the personalization of the ad break is abandoned and the underlying content is shown. This feature applies to *ad replacement* in live and VOD streams, rather than ad insertion, because it relies on an underlying content stream. For more information about ad break behavior, including ad replacement and insertion, see [Ad Behavior in MediaTailor](https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html) .
	PersonalizationThresholdSeconds() *float64
	SetPersonalizationThresholdSeconds(val *float64)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The URL for a high-quality video asset to transcode and use to fill in time that's not used by ads.
	//
	// MediaTailor shows the slate to fill in gaps in media content. Configuring the slate is optional for non-VPAID configurations. For VPAID, the slate is required because MediaTailor provides it in the slots that are designated for dynamic ad content. The slate must be a high-quality asset that contains both audio and video.
	SlateAdUrl() *string
	SetSlateAdUrl(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags to assign to the playback configuration.
	Tags() awscdk.TagManager
	// The name that is used to associate this playback configuration with a custom transcode profile.
	//
	// This overrides the dynamic transcoding defaults of MediaTailor. Use this only if you have already set up custom profiles with the help of AWS Support.
	TranscodeProfileName() *string
	SetTranscodeProfileName(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// The URL prefix for the parent manifest for the stream, minus the asset ID.
	//
	// The maximum length is 512 characters.
	VideoContentSourceUrl() *string
	SetVideoContentSourceUrl(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnPlaybackConfiguration
type jsiiProxy_CfnPlaybackConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPlaybackConfiguration) AdDecisionServerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adDecisionServerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) AttrDashConfigurationManifestEndpointPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDashConfigurationManifestEndpointPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) AttrHlsConfigurationManifestEndpointPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrHlsConfigurationManifestEndpointPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) AttrPlaybackConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPlaybackConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) AttrPlaybackEndpointPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPlaybackEndpointPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) AttrSessionInitializationEndpointPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSessionInitializationEndpointPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) AvailSuppression() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"availSuppression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) Bumper() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bumper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) CdnConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdnConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) ConfigurationAliases() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configurationAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) DashConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dashConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) LivePreRollConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"livePreRollConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) ManifestProcessingRules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manifestProcessingRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) PersonalizationThresholdSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"personalizationThresholdSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) SlateAdUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slateAdUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) TranscodeProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transcodeProfileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) VideoContentSourceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"videoContentSourceUrl",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaTailor::PlaybackConfiguration`.
func NewCfnPlaybackConfiguration(scope constructs.Construct, id *string, props *CfnPlaybackConfigurationProps) CfnPlaybackConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnPlaybackConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_mediatailor.CfnPlaybackConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaTailor::PlaybackConfiguration`.
func NewCfnPlaybackConfiguration_Override(c CfnPlaybackConfiguration, scope constructs.Construct, id *string, props *CfnPlaybackConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_mediatailor.CfnPlaybackConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration) SetAdDecisionServerUrl(val *string) {
	_jsii_.Set(
		j,
		"adDecisionServerUrl",
		val,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration) SetAvailSuppression(val interface{}) {
	_jsii_.Set(
		j,
		"availSuppression",
		val,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration) SetBumper(val interface{}) {
	_jsii_.Set(
		j,
		"bumper",
		val,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration) SetCdnConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"cdnConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration) SetConfigurationAliases(val interface{}) {
	_jsii_.Set(
		j,
		"configurationAliases",
		val,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration) SetDashConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"dashConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration) SetLivePreRollConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"livePreRollConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration) SetManifestProcessingRules(val interface{}) {
	_jsii_.Set(
		j,
		"manifestProcessingRules",
		val,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration) SetPersonalizationThresholdSeconds(val *float64) {
	_jsii_.Set(
		j,
		"personalizationThresholdSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration) SetSlateAdUrl(val *string) {
	_jsii_.Set(
		j,
		"slateAdUrl",
		val,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration) SetTranscodeProfileName(val *string) {
	_jsii_.Set(
		j,
		"transcodeProfileName",
		val,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration) SetVideoContentSourceUrl(val *string) {
	_jsii_.Set(
		j,
		"videoContentSourceUrl",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPlaybackConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediatailor.CfnPlaybackConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPlaybackConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediatailor.CfnPlaybackConfiguration",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnPlaybackConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediatailor.CfnPlaybackConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPlaybackConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_mediatailor.CfnPlaybackConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPlaybackConfiguration) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPlaybackConfiguration) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPlaybackConfiguration) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPlaybackConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPlaybackConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPlaybackConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPlaybackConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPlaybackConfiguration) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPlaybackConfiguration) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPlaybackConfiguration) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPlaybackConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPlaybackConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPlaybackConfiguration) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPlaybackConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPlaybackConfiguration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// For HLS, when set to `true` , MediaTailor passes through `EXT-X-CUE-IN` , `EXT-X-CUE-OUT` , and `EXT-X-SPLICEPOINT-SCTE35` ad markers from the origin manifest to the MediaTailor personalized manifest.
//
// No logic is applied to these ad markers. For example, if `EXT-X-CUE-OUT` has a value of `60` , but no ads are filled for that ad break, MediaTailor will not set the value to `0` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   adMarkerPassthroughProperty := &adMarkerPassthroughProperty{
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnPlaybackConfiguration_AdMarkerPassthroughProperty struct {
	// Enables ad marker passthrough for your configuration.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

// The configuration for avail suppression, also known as ad suppression.
//
// For more information about ad suppression, see [Ad Suppression](https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   availSuppressionProperty := &availSuppressionProperty{
//   	mode: jsii.String("mode"),
//   	value: jsii.String("value"),
//   }
//
type CfnPlaybackConfiguration_AvailSuppressionProperty struct {
	// Sets the ad suppression mode.
	//
	// By default, ad suppression is off and all ad breaks are filled with ads or slate. When Mode is set to BEHIND_LIVE_EDGE, ad suppression is active and MediaTailor won't fill ad breaks on or behind the ad suppression Value time in the manifest lookback window.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// A live edge offset time in HH:MM:SS.
	//
	// MediaTailor won't fill ad breaks on or behind this time in the manifest lookback window. If Value is set to 00:00:00, it is in sync with the live edge, and MediaTailor won't fill any ad breaks on or behind the live edge. If you set a Value time, MediaTailor won't fill any ad breaks on or behind this time in the manifest lookback window. For example, if you set 00:45:00, then MediaTailor will fill ad breaks that occur within 45 minutes behind the live edge, but won't fill ad breaks on or behind 45 minutes behind the live edge.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// The configuration for bumpers.
//
// Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see [Bumpers](https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bumperProperty := &bumperProperty{
//   	endUrl: jsii.String("endUrl"),
//   	startUrl: jsii.String("startUrl"),
//   }
//
type CfnPlaybackConfiguration_BumperProperty struct {
	// The URL for the end bumper asset.
	EndUrl *string `field:"optional" json:"endUrl" yaml:"endUrl"`
	// The URL for the start bumper asset.
	StartUrl *string `field:"optional" json:"startUrl" yaml:"startUrl"`
}

// The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cdnConfigurationProperty := &cdnConfigurationProperty{
//   	adSegmentUrlPrefix: jsii.String("adSegmentUrlPrefix"),
//   	contentSegmentUrlPrefix: jsii.String("contentSegmentUrlPrefix"),
//   }
//
type CfnPlaybackConfiguration_CdnConfigurationProperty struct {
	// A non-default content delivery network (CDN) to serve ad segments.
	//
	// By default, MediaTailor uses Amazon CloudFront with default cache settings as its CDN for ad segments. To set up an alternate CDN, create a rule in your CDN for the origin ads.mediatailor.&lt;region>.amazonaws.com. Then specify the rule's name in this AdSegmentUrlPrefix. When MediaTailor serves a manifest, it reports your CDN as the source for ad segments.
	AdSegmentUrlPrefix *string `field:"optional" json:"adSegmentUrlPrefix" yaml:"adSegmentUrlPrefix"`
	// A content delivery network (CDN) to cache content segments, so that content requests don’t always have to go to the origin server.
	//
	// First, create a rule in your CDN for the content segment origin server. Then specify the rule's name in this ContentSegmentUrlPrefix. When MediaTailor serves a manifest, it reports your CDN as the source for content segments.
	ContentSegmentUrlPrefix *string `field:"optional" json:"contentSegmentUrlPrefix" yaml:"contentSegmentUrlPrefix"`
}

// The configuration for DASH content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashConfigurationProperty := &dashConfigurationProperty{
//   	manifestEndpointPrefix: jsii.String("manifestEndpointPrefix"),
//   	mpdLocation: jsii.String("mpdLocation"),
//   	originManifestType: jsii.String("originManifestType"),
//   }
//
type CfnPlaybackConfiguration_DashConfigurationProperty struct {
	// The URL generated by MediaTailor to initiate a playback session.
	//
	// The session uses server-side reporting. This setting is ignored in PUT operations.
	ManifestEndpointPrefix *string `field:"optional" json:"manifestEndpointPrefix" yaml:"manifestEndpointPrefix"`
	// The setting that controls whether MediaTailor includes the Location tag in DASH manifests.
	//
	// MediaTailor populates the Location tag with the URL for manifest update requests, to be used by players that don't support sticky redirects. Disable this if you have CDN routing rules set up for accessing MediaTailor manifests, and you are either using client-side reporting or your players support sticky HTTP redirects. Valid values are DISABLED and EMT_DEFAULT. The EMT_DEFAULT setting enables the inclusion of the tag and is the default value.
	MpdLocation *string `field:"optional" json:"mpdLocation" yaml:"mpdLocation"`
	// The setting that controls whether MediaTailor handles manifests from the origin server as multi-period manifests or single-period manifests.
	//
	// If your origin server produces single-period manifests, set this to SINGLE_PERIOD. The default setting is MULTI_PERIOD. For multi-period manifests, omit this setting or set it to MULTI_PERIOD.
	OriginManifestType *string `field:"optional" json:"originManifestType" yaml:"originManifestType"`
}

// The configuration for HLS content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsConfigurationProperty := &hlsConfigurationProperty{
//   	manifestEndpointPrefix: jsii.String("manifestEndpointPrefix"),
//   }
//
type CfnPlaybackConfiguration_HlsConfigurationProperty struct {
	// The URL that is used to initiate a playback session for devices that support Apple HLS.
	//
	// The session uses server-side reporting.
	ManifestEndpointPrefix *string `field:"optional" json:"manifestEndpointPrefix" yaml:"manifestEndpointPrefix"`
}

// The configuration for pre-roll ad insertion.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   livePreRollConfigurationProperty := &livePreRollConfigurationProperty{
//   	adDecisionServerUrl: jsii.String("adDecisionServerUrl"),
//   	maxDurationSeconds: jsii.Number(123),
//   }
//
type CfnPlaybackConfiguration_LivePreRollConfigurationProperty struct {
	// The URL for the ad decision server (ADS) for pre-roll ads.
	//
	// This includes the specification of static parameters and placeholders for dynamic parameters. MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing, you can provide a static VAST URL. The maximum length is 25,000 characters.
	AdDecisionServerUrl *string `field:"optional" json:"adDecisionServerUrl" yaml:"adDecisionServerUrl"`
	// The maximum allowed duration for the pre-roll ad avail.
	//
	// MediaTailor won't play pre-roll ads to exceed this duration, regardless of the total duration of ads that the ADS returns.
	MaxDurationSeconds *float64 `field:"optional" json:"maxDurationSeconds" yaml:"maxDurationSeconds"`
}

// The configuration for manifest processing rules.
//
// Manifest processing rules enable customization of the personalized manifests created by MediaTailor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   manifestProcessingRulesProperty := &manifestProcessingRulesProperty{
//   	adMarkerPassthrough: &adMarkerPassthroughProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   }
//
type CfnPlaybackConfiguration_ManifestProcessingRulesProperty struct {
	// For HLS, when set to `true` , MediaTailor passes through `EXT-X-CUE-IN` , `EXT-X-CUE-OUT` , and `EXT-X-SPLICEPOINT-SCTE35` ad markers from the origin manifest to the MediaTailor personalized manifest.
	//
	// No logic is applied to these ad markers. For example, if `EXT-X-CUE-OUT` has a value of `60` , but no ads are filled for that ad break, MediaTailor will not set the value to `0` .
	AdMarkerPassthrough interface{} `field:"optional" json:"adMarkerPassthrough" yaml:"adMarkerPassthrough"`
}

// Properties for defining a `CfnPlaybackConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationAliases interface{}
//
//   cfnPlaybackConfigurationProps := &cfnPlaybackConfigurationProps{
//   	adDecisionServerUrl: jsii.String("adDecisionServerUrl"),
//   	name: jsii.String("name"),
//   	videoContentSourceUrl: jsii.String("videoContentSourceUrl"),
//
//   	// the properties below are optional
//   	availSuppression: &availSuppressionProperty{
//   		mode: jsii.String("mode"),
//   		value: jsii.String("value"),
//   	},
//   	bumper: &bumperProperty{
//   		endUrl: jsii.String("endUrl"),
//   		startUrl: jsii.String("startUrl"),
//   	},
//   	cdnConfiguration: &cdnConfigurationProperty{
//   		adSegmentUrlPrefix: jsii.String("adSegmentUrlPrefix"),
//   		contentSegmentUrlPrefix: jsii.String("contentSegmentUrlPrefix"),
//   	},
//   	configurationAliases: map[string]interface{}{
//   		"configurationAliasesKey": configurationAliases,
//   	},
//   	dashConfiguration: &dashConfigurationProperty{
//   		manifestEndpointPrefix: jsii.String("manifestEndpointPrefix"),
//   		mpdLocation: jsii.String("mpdLocation"),
//   		originManifestType: jsii.String("originManifestType"),
//   	},
//   	livePreRollConfiguration: &livePreRollConfigurationProperty{
//   		adDecisionServerUrl: jsii.String("adDecisionServerUrl"),
//   		maxDurationSeconds: jsii.Number(123),
//   	},
//   	manifestProcessingRules: &manifestProcessingRulesProperty{
//   		adMarkerPassthrough: &adMarkerPassthroughProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   	personalizationThresholdSeconds: jsii.Number(123),
//   	slateAdUrl: jsii.String("slateAdUrl"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	transcodeProfileName: jsii.String("transcodeProfileName"),
//   }
//
type CfnPlaybackConfigurationProps struct {
	// The URL for the ad decision server (ADS).
	//
	// This includes the specification of static parameters and placeholders for dynamic parameters. MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing you can provide a static VAST URL. The maximum length is 25,000 characters.
	AdDecisionServerUrl *string `field:"required" json:"adDecisionServerUrl" yaml:"adDecisionServerUrl"`
	// The identifier for the playback configuration.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The URL prefix for the parent manifest for the stream, minus the asset ID.
	//
	// The maximum length is 512 characters.
	VideoContentSourceUrl *string `field:"required" json:"videoContentSourceUrl" yaml:"videoContentSourceUrl"`
	// The configuration for avail suppression, also known as ad suppression.
	//
	// For more information about ad suppression, see [Ad Suppression](https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html) .
	AvailSuppression interface{} `field:"optional" json:"availSuppression" yaml:"availSuppression"`
	// The configuration for bumpers.
	//
	// Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see [Bumpers](https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html) .
	Bumper interface{} `field:"optional" json:"bumper" yaml:"bumper"`
	// The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.
	CdnConfiguration interface{} `field:"optional" json:"cdnConfiguration" yaml:"cdnConfiguration"`
	// The player parameters and aliases used as dynamic variables during session initialization.
	//
	// For more information, see [Domain Variables](https://docs.aws.amazon.com/mediatailor/latest/ug/variables-domain.html) .
	ConfigurationAliases interface{} `field:"optional" json:"configurationAliases" yaml:"configurationAliases"`
	// The configuration for DASH content.
	DashConfiguration interface{} `field:"optional" json:"dashConfiguration" yaml:"dashConfiguration"`
	// The configuration for pre-roll ad insertion.
	LivePreRollConfiguration interface{} `field:"optional" json:"livePreRollConfiguration" yaml:"livePreRollConfiguration"`
	// The configuration for manifest processing rules.
	//
	// Manifest processing rules enable customization of the personalized manifests created by MediaTailor.
	ManifestProcessingRules interface{} `field:"optional" json:"manifestProcessingRules" yaml:"manifestProcessingRules"`
	// Defines the maximum duration of underfilled ad time (in seconds) allowed in an ad break.
	//
	// If the duration of underfilled ad time exceeds the personalization threshold, then the personalization of the ad break is abandoned and the underlying content is shown. This feature applies to *ad replacement* in live and VOD streams, rather than ad insertion, because it relies on an underlying content stream. For more information about ad break behavior, including ad replacement and insertion, see [Ad Behavior in MediaTailor](https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html) .
	PersonalizationThresholdSeconds *float64 `field:"optional" json:"personalizationThresholdSeconds" yaml:"personalizationThresholdSeconds"`
	// The URL for a high-quality video asset to transcode and use to fill in time that's not used by ads.
	//
	// MediaTailor shows the slate to fill in gaps in media content. Configuring the slate is optional for non-VPAID configurations. For VPAID, the slate is required because MediaTailor provides it in the slots that are designated for dynamic ad content. The slate must be a high-quality asset that contains both audio and video.
	SlateAdUrl *string `field:"optional" json:"slateAdUrl" yaml:"slateAdUrl"`
	// The tags to assign to the playback configuration.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name that is used to associate this playback configuration with a custom transcode profile.
	//
	// This overrides the dynamic transcoding defaults of MediaTailor. Use this only if you have already set up custom profiles with the help of AWS Support.
	TranscodeProfileName *string `field:"optional" json:"transcodeProfileName" yaml:"transcodeProfileName"`
}


package awsmediatailor

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsmediatailor/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::MediaTailor::PlaybackConfiguration`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import mediatailor "github.com/aws/aws-cdk-go/awscdk/aws_mediatailor"
//
//   var configurationAliases interface{}
//   cfnPlaybackConfiguration := mediatailor.NewCfnPlaybackConfiguration(this, jsii.String("MyCfnPlaybackConfiguration"), &cfnPlaybackConfigurationProps{
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
//   	dashConfiguration: &dashConfigurationForPutProperty{
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
//   	sessionInitializationEndpointPrefix: jsii.String("sessionInitializationEndpointPrefix"),
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
	// `AWS::MediaTailor::PlaybackConfiguration.AdDecisionServerUrl`.
	AdDecisionServerUrl() *string
	SetAdDecisionServerUrl(val *string)
	// `AWS::MediaTailor::PlaybackConfiguration.AvailSuppression`.
	AvailSuppression() interface{}
	SetAvailSuppression(val interface{})
	// `AWS::MediaTailor::PlaybackConfiguration.Bumper`.
	Bumper() interface{}
	SetBumper(val interface{})
	// `AWS::MediaTailor::PlaybackConfiguration.CdnConfiguration`.
	CdnConfiguration() interface{}
	SetCdnConfiguration(val interface{})
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// `AWS::MediaTailor::PlaybackConfiguration.ConfigurationAliases`.
	ConfigurationAliases() interface{}
	SetConfigurationAliases(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// `AWS::MediaTailor::PlaybackConfiguration.DashConfiguration`.
	DashConfiguration() interface{}
	SetDashConfiguration(val interface{})
	// `AWS::MediaTailor::PlaybackConfiguration.LivePreRollConfiguration`.
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
	// Experimental.
	LogicalId() *string
	// `AWS::MediaTailor::PlaybackConfiguration.ManifestProcessingRules`.
	ManifestProcessingRules() interface{}
	SetManifestProcessingRules(val interface{})
	// `AWS::MediaTailor::PlaybackConfiguration.Name`.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// `AWS::MediaTailor::PlaybackConfiguration.PersonalizationThresholdSeconds`.
	PersonalizationThresholdSeconds() *float64
	SetPersonalizationThresholdSeconds(val *float64)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// `AWS::MediaTailor::PlaybackConfiguration.SessionInitializationEndpointPrefix`.
	SessionInitializationEndpointPrefix() *string
	SetSessionInitializationEndpointPrefix(val *string)
	// `AWS::MediaTailor::PlaybackConfiguration.SlateAdUrl`.
	SlateAdUrl() *string
	SetSlateAdUrl(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::MediaTailor::PlaybackConfiguration.Tags`.
	Tags() awscdk.TagManager
	// `AWS::MediaTailor::PlaybackConfiguration.TranscodeProfileName`.
	TranscodeProfileName() *string
	SetTranscodeProfileName(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// `AWS::MediaTailor::PlaybackConfiguration.VideoContentSourceUrl`.
	VideoContentSourceUrl() *string
	SetVideoContentSourceUrl(val *string)
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

func (j *jsiiProxy_CfnPlaybackConfiguration) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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

func (j *jsiiProxy_CfnPlaybackConfiguration) SessionInitializationEndpointPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionInitializationEndpointPrefix",
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
func NewCfnPlaybackConfiguration(scope awscdk.Construct, id *string, props *CfnPlaybackConfigurationProps) CfnPlaybackConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnPlaybackConfiguration{}

	_jsii_.Create(
		"monocdk.aws_mediatailor.CfnPlaybackConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaTailor::PlaybackConfiguration`.
func NewCfnPlaybackConfiguration_Override(c CfnPlaybackConfiguration, scope awscdk.Construct, id *string, props *CfnPlaybackConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_mediatailor.CfnPlaybackConfiguration",
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

func (j *jsiiProxy_CfnPlaybackConfiguration) SetSessionInitializationEndpointPrefix(val *string) {
	_jsii_.Set(
		j,
		"sessionInitializationEndpointPrefix",
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
// Experimental.
func CfnPlaybackConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_mediatailor.CfnPlaybackConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnPlaybackConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_mediatailor.CfnPlaybackConfiguration",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnPlaybackConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_mediatailor.CfnPlaybackConfiguration",
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
		"monocdk.aws_mediatailor.CfnPlaybackConfiguration",
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

func (c *jsiiProxy_CfnPlaybackConfiguration) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnPlaybackConfiguration) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnPlaybackConfiguration) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPlaybackConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPlaybackConfiguration) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnPlaybackConfiguration) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

func (c *jsiiProxy_CfnPlaybackConfiguration) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import mediatailor "github.com/aws/aws-cdk-go/awscdk/aws_mediatailor"
//   adMarkerPassthroughProperty := &adMarkerPassthroughProperty{
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnPlaybackConfiguration_AdMarkerPassthroughProperty struct {
	// `CfnPlaybackConfiguration.AdMarkerPassthroughProperty.Enabled`.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import mediatailor "github.com/aws/aws-cdk-go/awscdk/aws_mediatailor"
//   availSuppressionProperty := &availSuppressionProperty{
//   	mode: jsii.String("mode"),
//   	value: jsii.String("value"),
//   }
//
type CfnPlaybackConfiguration_AvailSuppressionProperty struct {
	// `CfnPlaybackConfiguration.AvailSuppressionProperty.Mode`.
	Mode *string `json:"mode" yaml:"mode"`
	// `CfnPlaybackConfiguration.AvailSuppressionProperty.Value`.
	Value *string `json:"value" yaml:"value"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import mediatailor "github.com/aws/aws-cdk-go/awscdk/aws_mediatailor"
//   bumperProperty := &bumperProperty{
//   	endUrl: jsii.String("endUrl"),
//   	startUrl: jsii.String("startUrl"),
//   }
//
type CfnPlaybackConfiguration_BumperProperty struct {
	// `CfnPlaybackConfiguration.BumperProperty.EndUrl`.
	EndUrl *string `json:"endUrl" yaml:"endUrl"`
	// `CfnPlaybackConfiguration.BumperProperty.StartUrl`.
	StartUrl *string `json:"startUrl" yaml:"startUrl"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import mediatailor "github.com/aws/aws-cdk-go/awscdk/aws_mediatailor"
//   cdnConfigurationProperty := &cdnConfigurationProperty{
//   	adSegmentUrlPrefix: jsii.String("adSegmentUrlPrefix"),
//   	contentSegmentUrlPrefix: jsii.String("contentSegmentUrlPrefix"),
//   }
//
type CfnPlaybackConfiguration_CdnConfigurationProperty struct {
	// `CfnPlaybackConfiguration.CdnConfigurationProperty.AdSegmentUrlPrefix`.
	AdSegmentUrlPrefix *string `json:"adSegmentUrlPrefix" yaml:"adSegmentUrlPrefix"`
	// `CfnPlaybackConfiguration.CdnConfigurationProperty.ContentSegmentUrlPrefix`.
	ContentSegmentUrlPrefix *string `json:"contentSegmentUrlPrefix" yaml:"contentSegmentUrlPrefix"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import mediatailor "github.com/aws/aws-cdk-go/awscdk/aws_mediatailor"
//   dashConfigurationForPutProperty := &dashConfigurationForPutProperty{
//   	mpdLocation: jsii.String("mpdLocation"),
//   	originManifestType: jsii.String("originManifestType"),
//   }
//
type CfnPlaybackConfiguration_DashConfigurationForPutProperty struct {
	// `CfnPlaybackConfiguration.DashConfigurationForPutProperty.MpdLocation`.
	MpdLocation *string `json:"mpdLocation" yaml:"mpdLocation"`
	// `CfnPlaybackConfiguration.DashConfigurationForPutProperty.OriginManifestType`.
	OriginManifestType *string `json:"originManifestType" yaml:"originManifestType"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import mediatailor "github.com/aws/aws-cdk-go/awscdk/aws_mediatailor"
//   livePreRollConfigurationProperty := &livePreRollConfigurationProperty{
//   	adDecisionServerUrl: jsii.String("adDecisionServerUrl"),
//   	maxDurationSeconds: jsii.Number(123),
//   }
//
type CfnPlaybackConfiguration_LivePreRollConfigurationProperty struct {
	// `CfnPlaybackConfiguration.LivePreRollConfigurationProperty.AdDecisionServerUrl`.
	AdDecisionServerUrl *string `json:"adDecisionServerUrl" yaml:"adDecisionServerUrl"`
	// `CfnPlaybackConfiguration.LivePreRollConfigurationProperty.MaxDurationSeconds`.
	MaxDurationSeconds *float64 `json:"maxDurationSeconds" yaml:"maxDurationSeconds"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import mediatailor "github.com/aws/aws-cdk-go/awscdk/aws_mediatailor"
//   manifestProcessingRulesProperty := &manifestProcessingRulesProperty{
//   	adMarkerPassthrough: &adMarkerPassthroughProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   }
//
type CfnPlaybackConfiguration_ManifestProcessingRulesProperty struct {
	// `CfnPlaybackConfiguration.ManifestProcessingRulesProperty.AdMarkerPassthrough`.
	AdMarkerPassthrough interface{} `json:"adMarkerPassthrough" yaml:"adMarkerPassthrough"`
}

// Properties for defining a `CfnPlaybackConfiguration`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import mediatailor "github.com/aws/aws-cdk-go/awscdk/aws_mediatailor"
//
//   var configurationAliases interface{}
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
//   	dashConfiguration: &dashConfigurationForPutProperty{
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
//   	sessionInitializationEndpointPrefix: jsii.String("sessionInitializationEndpointPrefix"),
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
	// `AWS::MediaTailor::PlaybackConfiguration.AdDecisionServerUrl`.
	AdDecisionServerUrl *string `json:"adDecisionServerUrl" yaml:"adDecisionServerUrl"`
	// `AWS::MediaTailor::PlaybackConfiguration.Name`.
	Name *string `json:"name" yaml:"name"`
	// `AWS::MediaTailor::PlaybackConfiguration.VideoContentSourceUrl`.
	VideoContentSourceUrl *string `json:"videoContentSourceUrl" yaml:"videoContentSourceUrl"`
	// `AWS::MediaTailor::PlaybackConfiguration.AvailSuppression`.
	AvailSuppression interface{} `json:"availSuppression" yaml:"availSuppression"`
	// `AWS::MediaTailor::PlaybackConfiguration.Bumper`.
	Bumper interface{} `json:"bumper" yaml:"bumper"`
	// `AWS::MediaTailor::PlaybackConfiguration.CdnConfiguration`.
	CdnConfiguration interface{} `json:"cdnConfiguration" yaml:"cdnConfiguration"`
	// `AWS::MediaTailor::PlaybackConfiguration.ConfigurationAliases`.
	ConfigurationAliases interface{} `json:"configurationAliases" yaml:"configurationAliases"`
	// `AWS::MediaTailor::PlaybackConfiguration.DashConfiguration`.
	DashConfiguration interface{} `json:"dashConfiguration" yaml:"dashConfiguration"`
	// `AWS::MediaTailor::PlaybackConfiguration.LivePreRollConfiguration`.
	LivePreRollConfiguration interface{} `json:"livePreRollConfiguration" yaml:"livePreRollConfiguration"`
	// `AWS::MediaTailor::PlaybackConfiguration.ManifestProcessingRules`.
	ManifestProcessingRules interface{} `json:"manifestProcessingRules" yaml:"manifestProcessingRules"`
	// `AWS::MediaTailor::PlaybackConfiguration.PersonalizationThresholdSeconds`.
	PersonalizationThresholdSeconds *float64 `json:"personalizationThresholdSeconds" yaml:"personalizationThresholdSeconds"`
	// `AWS::MediaTailor::PlaybackConfiguration.SessionInitializationEndpointPrefix`.
	SessionInitializationEndpointPrefix *string `json:"sessionInitializationEndpointPrefix" yaml:"sessionInitializationEndpointPrefix"`
	// `AWS::MediaTailor::PlaybackConfiguration.SlateAdUrl`.
	SlateAdUrl *string `json:"slateAdUrl" yaml:"slateAdUrl"`
	// `AWS::MediaTailor::PlaybackConfiguration.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// `AWS::MediaTailor::PlaybackConfiguration.TranscodeProfileName`.
	TranscodeProfileName *string `json:"transcodeProfileName" yaml:"transcodeProfileName"`
}


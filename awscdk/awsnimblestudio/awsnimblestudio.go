package awsnimblestudio

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnimblestudio/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::NimbleStudio::LaunchProfile`.
//
// The `AWS::NimbleStudio::LaunchProfile` resource represents access permissions for a set of studio components, including types of workstations, render farms, and shared file systems. Launch profiles are shared with studio users to give them access to the set of studio components.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLaunchProfile := awscdk.Aws_nimblestudio.NewCfnLaunchProfile(this, jsii.String("MyCfnLaunchProfile"), &cfnLaunchProfileProps{
//   	ec2SubnetIds: []*string{
//   		jsii.String("ec2SubnetIds"),
//   	},
//   	launchProfileProtocolVersions: []*string{
//   		jsii.String("launchProfileProtocolVersions"),
//   	},
//   	name: jsii.String("name"),
//   	streamConfiguration: &streamConfigurationProperty{
//   		clipboardMode: jsii.String("clipboardMode"),
//   		ec2InstanceTypes: []*string{
//   			jsii.String("ec2InstanceTypes"),
//   		},
//   		streamingImageIds: []*string{
//   			jsii.String("streamingImageIds"),
//   		},
//
//   		// the properties below are optional
//   		maxSessionLengthInMinutes: jsii.Number(123),
//   		maxStoppedSessionLengthInMinutes: jsii.Number(123),
//   		sessionStorage: &streamConfigurationSessionStorageProperty{
//   			mode: []*string{
//   				jsii.String("mode"),
//   			},
//   			root: &streamingSessionStorageRootProperty{
//   				linux: jsii.String("linux"),
//   				windows: jsii.String("windows"),
//   			},
//   		},
//   	},
//   	studioComponentIds: []*string{
//   		jsii.String("studioComponentIds"),
//   	},
//   	studioId: jsii.String("studioId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
type CfnLaunchProfile interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The unique identifier for the launch profile resource.
	AttrLaunchProfileId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A human-readable description of the launch profile.
	Description() *string
	SetDescription(val *string)
	// Unique identifiers for a collection of EC2 subnets.
	Ec2SubnetIds() *[]*string
	SetEc2SubnetIds(val *[]*string)
	// The version number of the protocol that is used by the launch profile.
	//
	// The only valid version is "2021-03-31".
	LaunchProfileProtocolVersions() *[]*string
	SetLaunchProfileProtocolVersions(val *[]*string)
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
	// A friendly name for the launch profile.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// A configuration for a streaming session.
	StreamConfiguration() interface{}
	SetStreamConfiguration(val interface{})
	// Unique identifiers for a collection of studio components that can be used with this launch profile.
	StudioComponentIds() *[]*string
	SetStudioComponentIds(val *[]*string)
	// The unique identifier for a studio resource.
	//
	// In Nimble Studio , all other resources are contained in a studio resource.
	StudioId() *string
	SetStudioId(val *string)
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
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

// The jsii proxy struct for CfnLaunchProfile
type jsiiProxy_CfnLaunchProfile struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLaunchProfile) AttrLaunchProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLaunchProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchProfile) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchProfile) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchProfile) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchProfile) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchProfile) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchProfile) Ec2SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ec2SubnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchProfile) LaunchProfileProtocolVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"launchProfileProtocolVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchProfile) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchProfile) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchProfile) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchProfile) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchProfile) StreamConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchProfile) StudioComponentIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"studioComponentIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchProfile) StudioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"studioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchProfile) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchProfile) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::NimbleStudio::LaunchProfile`.
func NewCfnLaunchProfile(scope constructs.Construct, id *string, props *CfnLaunchProfileProps) CfnLaunchProfile {
	_init_.Initialize()

	j := jsiiProxy_CfnLaunchProfile{}

	_jsii_.Create(
		"aws-cdk-lib.aws_nimblestudio.CfnLaunchProfile",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::NimbleStudio::LaunchProfile`.
func NewCfnLaunchProfile_Override(c CfnLaunchProfile, scope constructs.Construct, id *string, props *CfnLaunchProfileProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_nimblestudio.CfnLaunchProfile",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLaunchProfile) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchProfile) SetEc2SubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"ec2SubnetIds",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchProfile) SetLaunchProfileProtocolVersions(val *[]*string) {
	_jsii_.Set(
		j,
		"launchProfileProtocolVersions",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchProfile) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchProfile) SetStreamConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"streamConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchProfile) SetStudioComponentIds(val *[]*string) {
	_jsii_.Set(
		j,
		"studioComponentIds",
		val,
	)
}

func (j *jsiiProxy_CfnLaunchProfile) SetStudioId(val *string) {
	_jsii_.Set(
		j,
		"studioId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLaunchProfile_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_nimblestudio.CfnLaunchProfile",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLaunchProfile_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_nimblestudio.CfnLaunchProfile",
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
func CfnLaunchProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_nimblestudio.CfnLaunchProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLaunchProfile_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_nimblestudio.CfnLaunchProfile",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLaunchProfile) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnLaunchProfile) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLaunchProfile) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnLaunchProfile) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnLaunchProfile) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnLaunchProfile) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnLaunchProfile) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnLaunchProfile) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunchProfile) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunchProfile) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnLaunchProfile) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLaunchProfile) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunchProfile) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunchProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunchProfile) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A configuration for a streaming session.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamConfigurationProperty := &streamConfigurationProperty{
//   	clipboardMode: jsii.String("clipboardMode"),
//   	ec2InstanceTypes: []*string{
//   		jsii.String("ec2InstanceTypes"),
//   	},
//   	streamingImageIds: []*string{
//   		jsii.String("streamingImageIds"),
//   	},
//
//   	// the properties below are optional
//   	maxSessionLengthInMinutes: jsii.Number(123),
//   	maxStoppedSessionLengthInMinutes: jsii.Number(123),
//   	sessionStorage: &streamConfigurationSessionStorageProperty{
//   		mode: []*string{
//   			jsii.String("mode"),
//   		},
//   		root: &streamingSessionStorageRootProperty{
//   			linux: jsii.String("linux"),
//   			windows: jsii.String("windows"),
//   		},
//   	},
//   }
//
type CfnLaunchProfile_StreamConfigurationProperty struct {
	// Enable or disable the use of the system clipboard to copy and paste between the streaming session and streaming client.
	ClipboardMode *string `field:"required" json:"clipboardMode" yaml:"clipboardMode"`
	// The EC2 instance types that users can select from when launching a streaming session with this launch profile.
	Ec2InstanceTypes *[]*string `field:"required" json:"ec2InstanceTypes" yaml:"ec2InstanceTypes"`
	// The streaming images that users can select from when launching a streaming session with this launch profile.
	StreamingImageIds *[]*string `field:"required" json:"streamingImageIds" yaml:"streamingImageIds"`
	// The length of time, in minutes, that a streaming session can be active before it is stopped or terminated.
	//
	// After this point, Nimble Studio automatically terminates or stops the session. The default length of time is 690 minutes, and the maximum length of time is 30 days.
	MaxSessionLengthInMinutes *float64 `field:"optional" json:"maxSessionLengthInMinutes" yaml:"maxSessionLengthInMinutes"`
	// Integer that determines if you can start and stop your sessions and how long a session can stay in the STOPPED state.
	//
	// The default value is 0. The maximum value is 5760.
	//
	// If the value is missing or set to 0, your sessions can’t be stopped. If you then call `StopStreamingSession` , the session fails. If the time that a session stays in the READY state exceeds the `maxSessionLengthInMinutes` value, the session will automatically be terminated (instead of stopped).
	//
	// If the value is set to a positive number, the session can be stopped. You can call `StopStreamingSession` to stop sessions in the READY state. If the time that a session stays in the READY state exceeds the `maxSessionLengthInMinutes` value, the session will automatically be stopped (instead of terminated).
	MaxStoppedSessionLengthInMinutes *float64 `field:"optional" json:"maxStoppedSessionLengthInMinutes" yaml:"maxStoppedSessionLengthInMinutes"`
	// (Optional) The upload storage for a streaming session.
	SessionStorage interface{} `field:"optional" json:"sessionStorage" yaml:"sessionStorage"`
}

// The configuration for a streaming session’s upload storage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamConfigurationSessionStorageProperty := &streamConfigurationSessionStorageProperty{
//   	mode: []*string{
//   		jsii.String("mode"),
//   	},
//   	root: &streamingSessionStorageRootProperty{
//   		linux: jsii.String("linux"),
//   		windows: jsii.String("windows"),
//   	},
//   }
//
type CfnLaunchProfile_StreamConfigurationSessionStorageProperty struct {
	// Allows artists to upload files to their workstations.
	//
	// The only valid option is `UPLOAD` .
	Mode *[]*string `field:"optional" json:"mode" yaml:"mode"`
	// The configuration for the upload storage root of the streaming session.
	Root interface{} `field:"optional" json:"root" yaml:"root"`
}

// The upload storage root location (folder) on streaming workstations where files are uploaded.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamingSessionStorageRootProperty := &streamingSessionStorageRootProperty{
//   	linux: jsii.String("linux"),
//   	windows: jsii.String("windows"),
//   }
//
type CfnLaunchProfile_StreamingSessionStorageRootProperty struct {
	// The folder path in Linux workstations where files are uploaded.
	Linux *string `field:"optional" json:"linux" yaml:"linux"`
	// The folder path in Windows workstations where files are uploaded.
	Windows *string `field:"optional" json:"windows" yaml:"windows"`
}

// Properties for defining a `CfnLaunchProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLaunchProfileProps := &cfnLaunchProfileProps{
//   	ec2SubnetIds: []*string{
//   		jsii.String("ec2SubnetIds"),
//   	},
//   	launchProfileProtocolVersions: []*string{
//   		jsii.String("launchProfileProtocolVersions"),
//   	},
//   	name: jsii.String("name"),
//   	streamConfiguration: &streamConfigurationProperty{
//   		clipboardMode: jsii.String("clipboardMode"),
//   		ec2InstanceTypes: []*string{
//   			jsii.String("ec2InstanceTypes"),
//   		},
//   		streamingImageIds: []*string{
//   			jsii.String("streamingImageIds"),
//   		},
//
//   		// the properties below are optional
//   		maxSessionLengthInMinutes: jsii.Number(123),
//   		maxStoppedSessionLengthInMinutes: jsii.Number(123),
//   		sessionStorage: &streamConfigurationSessionStorageProperty{
//   			mode: []*string{
//   				jsii.String("mode"),
//   			},
//   			root: &streamingSessionStorageRootProperty{
//   				linux: jsii.String("linux"),
//   				windows: jsii.String("windows"),
//   			},
//   		},
//   	},
//   	studioComponentIds: []*string{
//   		jsii.String("studioComponentIds"),
//   	},
//   	studioId: jsii.String("studioId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnLaunchProfileProps struct {
	// Unique identifiers for a collection of EC2 subnets.
	Ec2SubnetIds *[]*string `field:"required" json:"ec2SubnetIds" yaml:"ec2SubnetIds"`
	// The version number of the protocol that is used by the launch profile.
	//
	// The only valid version is "2021-03-31".
	LaunchProfileProtocolVersions *[]*string `field:"required" json:"launchProfileProtocolVersions" yaml:"launchProfileProtocolVersions"`
	// A friendly name for the launch profile.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A configuration for a streaming session.
	StreamConfiguration interface{} `field:"required" json:"streamConfiguration" yaml:"streamConfiguration"`
	// Unique identifiers for a collection of studio components that can be used with this launch profile.
	StudioComponentIds *[]*string `field:"required" json:"studioComponentIds" yaml:"studioComponentIds"`
	// The unique identifier for a studio resource.
	//
	// In Nimble Studio , all other resources are contained in a studio resource.
	StudioId *string `field:"required" json:"studioId" yaml:"studioId"`
	// A human-readable description of the launch profile.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::NimbleStudio::StreamingImage`.
//
// The `AWS::NimbleStudio::StreamingImage` resource creates a streaming image in a studio. A streaming image defines the operating system and software to be used in an  streaming session.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStreamingImage := awscdk.Aws_nimblestudio.NewCfnStreamingImage(this, jsii.String("MyCfnStreamingImage"), &cfnStreamingImageProps{
//   	ec2ImageId: jsii.String("ec2ImageId"),
//   	name: jsii.String("name"),
//   	studioId: jsii.String("studioId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
type CfnStreamingImage interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The list of IDs of EULAs that must be accepted before a streaming session can be started using this streaming image.
	AttrEulaIds() *[]*string
	// The owner of the streaming image, either the studioId that contains the streaming image or 'amazon' for images that are provided by  .
	AttrOwner() *string
	// The platform of the streaming image, either WINDOWS or LINUX.
	AttrPlatform() *string
	// The unique identifier for the streaming image resource.
	AttrStreamingImageId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A human-readable description of the streaming image.
	Description() *string
	SetDescription(val *string)
	// The ID of an EC2 machine image with which to create the streaming image.
	Ec2ImageId() *string
	SetEc2ImageId(val *string)
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
	// A friendly name for a streaming image resource.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The unique identifier for a studio resource.
	//
	// In Nimble Studio , all other resources are contained in a studio resource.
	StudioId() *string
	SetStudioId(val *string)
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
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

// The jsii proxy struct for CfnStreamingImage
type jsiiProxy_CfnStreamingImage struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnStreamingImage) AttrEulaIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrEulaIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingImage) AttrOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingImage) AttrPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingImage) AttrStreamingImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStreamingImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingImage) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingImage) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingImage) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingImage) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingImage) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingImage) Ec2ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2ImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingImage) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingImage) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingImage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingImage) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingImage) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingImage) StudioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"studioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingImage) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingImage) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::NimbleStudio::StreamingImage`.
func NewCfnStreamingImage(scope constructs.Construct, id *string, props *CfnStreamingImageProps) CfnStreamingImage {
	_init_.Initialize()

	j := jsiiProxy_CfnStreamingImage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_nimblestudio.CfnStreamingImage",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::NimbleStudio::StreamingImage`.
func NewCfnStreamingImage_Override(c CfnStreamingImage, scope constructs.Construct, id *string, props *CfnStreamingImageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_nimblestudio.CfnStreamingImage",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStreamingImage) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnStreamingImage) SetEc2ImageId(val *string) {
	_jsii_.Set(
		j,
		"ec2ImageId",
		val,
	)
}

func (j *jsiiProxy_CfnStreamingImage) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnStreamingImage) SetStudioId(val *string) {
	_jsii_.Set(
		j,
		"studioId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnStreamingImage_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_nimblestudio.CfnStreamingImage",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnStreamingImage_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_nimblestudio.CfnStreamingImage",
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
func CfnStreamingImage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_nimblestudio.CfnStreamingImage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStreamingImage_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_nimblestudio.CfnStreamingImage",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStreamingImage) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnStreamingImage) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnStreamingImage) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnStreamingImage) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnStreamingImage) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnStreamingImage) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnStreamingImage) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnStreamingImage) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStreamingImage) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStreamingImage) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnStreamingImage) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStreamingImage) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStreamingImage) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStreamingImage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStreamingImage) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnStreamingImage`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStreamingImageProps := &cfnStreamingImageProps{
//   	ec2ImageId: jsii.String("ec2ImageId"),
//   	name: jsii.String("name"),
//   	studioId: jsii.String("studioId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnStreamingImageProps struct {
	// The ID of an EC2 machine image with which to create the streaming image.
	Ec2ImageId *string `field:"required" json:"ec2ImageId" yaml:"ec2ImageId"`
	// A friendly name for a streaming image resource.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The unique identifier for a studio resource.
	//
	// In Nimble Studio , all other resources are contained in a studio resource.
	StudioId *string `field:"required" json:"studioId" yaml:"studioId"`
	// A human-readable description of the streaming image.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::NimbleStudio::Studio`.
//
// The `AWS::NimbleStudio::Studio` resource creates a new studio resource. In  , all other resources are contained in a studio.
//
// When creating a studio, two IAM roles must be provided: the admin role and the user role. These roles are assumed by your users when they log in to the  portal. The user role must have the AmazonNimbleStudio-StudioUser managed policy attached for the portal to function properly. The Admin Role must have the AmazonNimbleStudio-StudioAdmin managed policy attached for the portal to function properly.
//
// You can optionally specify an AWS Key Management Service key in the StudioEncryptionConfiguration. In Nimble Studio, resource names, descriptions, initialization scripts, and other data you provide are always encrypted at rest using an AWS Key Management Service key. By default, this key is owned by AWS and managed on your behalf. You may provide your own AWS Key Management Service key when calling CreateStudio to encrypt this data using a key that you own and manage. When providing an AWS Key Management Service key during studio creation,  creates AWS Key Management Service grants in your account to provide your studio user and admin roles access to these AWS Key Management Service keys. If you delete this grant, the studio will no longer be accessible to your portal users. If you delete the studio AWS Key Management Service key, your studio will no longer be accessible.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStudio := awscdk.Aws_nimblestudio.NewCfnStudio(this, jsii.String("MyCfnStudio"), &cfnStudioProps{
//   	adminRoleArn: jsii.String("adminRoleArn"),
//   	displayName: jsii.String("displayName"),
//   	studioName: jsii.String("studioName"),
//   	userRoleArn: jsii.String("userRoleArn"),
//
//   	// the properties below are optional
//   	studioEncryptionConfiguration: &studioEncryptionConfigurationProperty{
//   		keyType: jsii.String("keyType"),
//
//   		// the properties below are optional
//   		keyArn: jsii.String("keyArn"),
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
type CfnStudio interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The IAM role that studio admins assume when logging in to the Nimble Studio portal.
	AdminRoleArn() *string
	SetAdminRoleArn(val *string)
	// The AWS Region where the studio resource is located.
	//
	// For example, `us-west-2` .
	AttrHomeRegion() *string
	// The AWS SSO application client ID that is used to integrate with AWS SSO , which enables AWS SSO users to log into the  portal.
	AttrSsoClientId() *string
	// The unique identifier for the studio resource.
	AttrStudioId() *string
	// The unique identifier for the studio resource.
	AttrStudioUrl() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A friendly name for the studio.
	DisplayName() *string
	SetDisplayName(val *string)
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
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Configuration of the encryption method that is used for the studio.
	StudioEncryptionConfiguration() interface{}
	SetStudioEncryptionConfiguration(val interface{})
	// The name of the studio, as included in the URL when accessing it in the Nimble Studio portal.
	StudioName() *string
	SetStudioName(val *string)
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// The IAM role that studio users assume when logging in to the Nimble Studio portal.
	UserRoleArn() *string
	SetUserRoleArn(val *string)
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

// The jsii proxy struct for CfnStudio
type jsiiProxy_CfnStudio struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnStudio) AdminRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) AttrHomeRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrHomeRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) AttrSsoClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSsoClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) AttrStudioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStudioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) AttrStudioUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStudioUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) StudioEncryptionConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"studioEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) StudioName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"studioName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) UserRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userRoleArn",
		&returns,
	)
	return returns
}


// Create a new `AWS::NimbleStudio::Studio`.
func NewCfnStudio(scope constructs.Construct, id *string, props *CfnStudioProps) CfnStudio {
	_init_.Initialize()

	j := jsiiProxy_CfnStudio{}

	_jsii_.Create(
		"aws-cdk-lib.aws_nimblestudio.CfnStudio",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::NimbleStudio::Studio`.
func NewCfnStudio_Override(c CfnStudio, scope constructs.Construct, id *string, props *CfnStudioProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_nimblestudio.CfnStudio",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStudio) SetAdminRoleArn(val *string) {
	_jsii_.Set(
		j,
		"adminRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnStudio) SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_CfnStudio) SetStudioEncryptionConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"studioEncryptionConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnStudio) SetStudioName(val *string) {
	_jsii_.Set(
		j,
		"studioName",
		val,
	)
}

func (j *jsiiProxy_CfnStudio) SetUserRoleArn(val *string) {
	_jsii_.Set(
		j,
		"userRoleArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnStudio_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_nimblestudio.CfnStudio",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnStudio_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_nimblestudio.CfnStudio",
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
func CfnStudio_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_nimblestudio.CfnStudio",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStudio_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_nimblestudio.CfnStudio",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStudio) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnStudio) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnStudio) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnStudio) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnStudio) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnStudio) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnStudio) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnStudio) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStudio) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStudio) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnStudio) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStudio) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStudio) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStudio) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStudio) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Configuration of the encryption method that is used for the studio.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   studioEncryptionConfigurationProperty := &studioEncryptionConfigurationProperty{
//   	keyType: jsii.String("keyType"),
//
//   	// the properties below are optional
//   	keyArn: jsii.String("keyArn"),
//   }
//
type CfnStudio_StudioEncryptionConfigurationProperty struct {
	// The type of KMS key that is used to encrypt studio data.
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// The ARN for a KMS key that is used to encrypt studio data.
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
}

// A CloudFormation `AWS::NimbleStudio::StudioComponent`.
//
// The `AWS::NimbleStudio::StudioComponent` resource represents a network resource that is used by a studio's users and workflows. A typical studio contains studio components for the following: a render farm, an Active Directory, a licensing service, and a shared file system.
//
// Access to a studio component is managed by specifying security groups for the resource, as well as its endpoint.
//
// A studio component also has a set of initialization scripts, which are returned by `GetLaunchProfileInitialization` . These initialization scripts run on streaming sessions when they start. They provide users with flexibility in controlling how studio resources are configured on a streaming session.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStudioComponent := awscdk.Aws_nimblestudio.NewCfnStudioComponent(this, jsii.String("MyCfnStudioComponent"), &cfnStudioComponentProps{
//   	name: jsii.String("name"),
//   	studioId: jsii.String("studioId"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	configuration: &studioComponentConfigurationProperty{
//   		activeDirectoryConfiguration: &activeDirectoryConfigurationProperty{
//   			computerAttributes: []interface{}{
//   				&activeDirectoryComputerAttributeProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			directoryId: jsii.String("directoryId"),
//   			organizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   		},
//   		computeFarmConfiguration: &computeFarmConfigurationProperty{
//   			activeDirectoryUser: jsii.String("activeDirectoryUser"),
//   			endpoint: jsii.String("endpoint"),
//   		},
//   		licenseServiceConfiguration: &licenseServiceConfigurationProperty{
//   			endpoint: jsii.String("endpoint"),
//   		},
//   		sharedFileSystemConfiguration: &sharedFileSystemConfigurationProperty{
//   			endpoint: jsii.String("endpoint"),
//   			fileSystemId: jsii.String("fileSystemId"),
//   			linuxMountPoint: jsii.String("linuxMountPoint"),
//   			shareName: jsii.String("shareName"),
//   			windowsMountDrive: jsii.String("windowsMountDrive"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	ec2SecurityGroupIds: []*string{
//   		jsii.String("ec2SecurityGroupIds"),
//   	},
//   	initializationScripts: []interface{}{
//   		&studioComponentInitializationScriptProperty{
//   			launchProfileProtocolVersion: jsii.String("launchProfileProtocolVersion"),
//   			platform: jsii.String("platform"),
//   			runContext: jsii.String("runContext"),
//   			script: jsii.String("script"),
//   		},
//   	},
//   	scriptParameters: []interface{}{
//   		&scriptParameterKeyValueProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	subtype: jsii.String("subtype"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
type CfnStudioComponent interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The unique identifier for the studio component resource.
	AttrStudioComponentId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The configuration of the studio component, based on component type.
	Configuration() interface{}
	SetConfiguration(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A human-readable description for the studio component resource.
	Description() *string
	SetDescription(val *string)
	// The EC2 security groups that control access to the studio component.
	Ec2SecurityGroupIds() *[]*string
	SetEc2SecurityGroupIds(val *[]*string)
	// Initialization scripts for studio components.
	InitializationScripts() interface{}
	SetInitializationScripts(val interface{})
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
	// A friendly name for the studio component resource.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Parameters for the studio component scripts.
	ScriptParameters() interface{}
	SetScriptParameters(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The unique identifier for a studio resource.
	//
	// In Nimble Studio , all other resources are contained in a studio resource.
	StudioId() *string
	SetStudioId(val *string)
	// The specific subtype of a studio component.
	Subtype() *string
	SetSubtype(val *string)
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
	// The type of the studio component.
	Type() *string
	SetType(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
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

// The jsii proxy struct for CfnStudioComponent
type jsiiProxy_CfnStudioComponent struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnStudioComponent) AttrStudioComponentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStudioComponentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioComponent) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioComponent) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioComponent) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioComponent) Configuration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioComponent) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioComponent) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioComponent) Ec2SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ec2SecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioComponent) InitializationScripts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initializationScripts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioComponent) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioComponent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioComponent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioComponent) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioComponent) ScriptParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scriptParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioComponent) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioComponent) StudioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"studioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioComponent) Subtype() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subtype",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioComponent) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioComponent) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioComponent) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::NimbleStudio::StudioComponent`.
func NewCfnStudioComponent(scope constructs.Construct, id *string, props *CfnStudioComponentProps) CfnStudioComponent {
	_init_.Initialize()

	j := jsiiProxy_CfnStudioComponent{}

	_jsii_.Create(
		"aws-cdk-lib.aws_nimblestudio.CfnStudioComponent",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::NimbleStudio::StudioComponent`.
func NewCfnStudioComponent_Override(c CfnStudioComponent, scope constructs.Construct, id *string, props *CfnStudioComponentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_nimblestudio.CfnStudioComponent",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStudioComponent) SetConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

func (j *jsiiProxy_CfnStudioComponent) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnStudioComponent) SetEc2SecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"ec2SecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_CfnStudioComponent) SetInitializationScripts(val interface{}) {
	_jsii_.Set(
		j,
		"initializationScripts",
		val,
	)
}

func (j *jsiiProxy_CfnStudioComponent) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnStudioComponent) SetScriptParameters(val interface{}) {
	_jsii_.Set(
		j,
		"scriptParameters",
		val,
	)
}

func (j *jsiiProxy_CfnStudioComponent) SetStudioId(val *string) {
	_jsii_.Set(
		j,
		"studioId",
		val,
	)
}

func (j *jsiiProxy_CfnStudioComponent) SetSubtype(val *string) {
	_jsii_.Set(
		j,
		"subtype",
		val,
	)
}

func (j *jsiiProxy_CfnStudioComponent) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnStudioComponent_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_nimblestudio.CfnStudioComponent",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnStudioComponent_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_nimblestudio.CfnStudioComponent",
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
func CfnStudioComponent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_nimblestudio.CfnStudioComponent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStudioComponent_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_nimblestudio.CfnStudioComponent",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStudioComponent) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnStudioComponent) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnStudioComponent) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnStudioComponent) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnStudioComponent) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnStudioComponent) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnStudioComponent) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnStudioComponent) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStudioComponent) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStudioComponent) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnStudioComponent) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStudioComponent) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStudioComponent) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStudioComponent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStudioComponent) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// An LDAP attribute of an Active Directory computer account, in the form of a name:value pair.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   activeDirectoryComputerAttributeProperty := &activeDirectoryComputerAttributeProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnStudioComponent_ActiveDirectoryComputerAttributeProperty struct {
	// The name for the LDAP attribute.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value for the LDAP attribute.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// The configuration for a Microsoft Active Directory (Microsoft AD) studio resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   activeDirectoryConfigurationProperty := &activeDirectoryConfigurationProperty{
//   	computerAttributes: []interface{}{
//   		&activeDirectoryComputerAttributeProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	directoryId: jsii.String("directoryId"),
//   	organizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   }
//
type CfnStudioComponent_ActiveDirectoryConfigurationProperty struct {
	// A collection of custom attributes for an Active Directory computer.
	ComputerAttributes interface{} `field:"optional" json:"computerAttributes" yaml:"computerAttributes"`
	// The directory ID of the Directory Service for Microsoft Active Directory to access using this studio component.
	DirectoryId *string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// The distinguished name (DN) and organizational unit (OU) of an Active Directory computer.
	OrganizationalUnitDistinguishedName *string `field:"optional" json:"organizationalUnitDistinguishedName" yaml:"organizationalUnitDistinguishedName"`
}

// The configuration for a render farm that is associated with a studio resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   computeFarmConfigurationProperty := &computeFarmConfigurationProperty{
//   	activeDirectoryUser: jsii.String("activeDirectoryUser"),
//   	endpoint: jsii.String("endpoint"),
//   }
//
type CfnStudioComponent_ComputeFarmConfigurationProperty struct {
	// The name of an Active Directory user that is used on ComputeFarm worker instances.
	ActiveDirectoryUser *string `field:"optional" json:"activeDirectoryUser" yaml:"activeDirectoryUser"`
	// The endpoint of the ComputeFarm that is accessed by the studio component resource.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

// The configuration for a license service that is associated with a studio resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   licenseServiceConfigurationProperty := &licenseServiceConfigurationProperty{
//   	endpoint: jsii.String("endpoint"),
//   }
//
type CfnStudioComponent_LicenseServiceConfigurationProperty struct {
	// The endpoint of the license service that is accessed by the studio component resource.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

// A parameter for a studio component script, in the form of a key:value pair.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scriptParameterKeyValueProperty := &scriptParameterKeyValueProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnStudioComponent_ScriptParameterKeyValueProperty struct {
	// A script parameter key.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A script parameter value.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// The configuration for a shared file storage system that is associated with a studio resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sharedFileSystemConfigurationProperty := &sharedFileSystemConfigurationProperty{
//   	endpoint: jsii.String("endpoint"),
//   	fileSystemId: jsii.String("fileSystemId"),
//   	linuxMountPoint: jsii.String("linuxMountPoint"),
//   	shareName: jsii.String("shareName"),
//   	windowsMountDrive: jsii.String("windowsMountDrive"),
//   }
//
type CfnStudioComponent_SharedFileSystemConfigurationProperty struct {
	// The endpoint of the shared file system that is accessed by the studio component resource.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The unique identifier for a file system.
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
	// The mount location for a shared file system on a Linux virtual workstation.
	LinuxMountPoint *string `field:"optional" json:"linuxMountPoint" yaml:"linuxMountPoint"`
	// The name of the file share.
	ShareName *string `field:"optional" json:"shareName" yaml:"shareName"`
	// The mount location for a shared file system on a Windows virtual workstation.
	WindowsMountDrive *string `field:"optional" json:"windowsMountDrive" yaml:"windowsMountDrive"`
}

// The configuration of the studio component, based on component type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   studioComponentConfigurationProperty := &studioComponentConfigurationProperty{
//   	activeDirectoryConfiguration: &activeDirectoryConfigurationProperty{
//   		computerAttributes: []interface{}{
//   			&activeDirectoryComputerAttributeProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		directoryId: jsii.String("directoryId"),
//   		organizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   	},
//   	computeFarmConfiguration: &computeFarmConfigurationProperty{
//   		activeDirectoryUser: jsii.String("activeDirectoryUser"),
//   		endpoint: jsii.String("endpoint"),
//   	},
//   	licenseServiceConfiguration: &licenseServiceConfigurationProperty{
//   		endpoint: jsii.String("endpoint"),
//   	},
//   	sharedFileSystemConfiguration: &sharedFileSystemConfigurationProperty{
//   		endpoint: jsii.String("endpoint"),
//   		fileSystemId: jsii.String("fileSystemId"),
//   		linuxMountPoint: jsii.String("linuxMountPoint"),
//   		shareName: jsii.String("shareName"),
//   		windowsMountDrive: jsii.String("windowsMountDrive"),
//   	},
//   }
//
type CfnStudioComponent_StudioComponentConfigurationProperty struct {
	// The configuration for a Microsoft Active Directory (Microsoft AD) studio resource.
	ActiveDirectoryConfiguration interface{} `field:"optional" json:"activeDirectoryConfiguration" yaml:"activeDirectoryConfiguration"`
	// The configuration for a render farm that is associated with a studio resource.
	ComputeFarmConfiguration interface{} `field:"optional" json:"computeFarmConfiguration" yaml:"computeFarmConfiguration"`
	// The configuration for a license service that is associated with a studio resource.
	LicenseServiceConfiguration interface{} `field:"optional" json:"licenseServiceConfiguration" yaml:"licenseServiceConfiguration"`
	// The configuration for a shared file storage system that is associated with a studio resource.
	SharedFileSystemConfiguration interface{} `field:"optional" json:"sharedFileSystemConfiguration" yaml:"sharedFileSystemConfiguration"`
}

// Initialization scripts for studio components.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   studioComponentInitializationScriptProperty := &studioComponentInitializationScriptProperty{
//   	launchProfileProtocolVersion: jsii.String("launchProfileProtocolVersion"),
//   	platform: jsii.String("platform"),
//   	runContext: jsii.String("runContext"),
//   	script: jsii.String("script"),
//   }
//
type CfnStudioComponent_StudioComponentInitializationScriptProperty struct {
	// The version number of the protocol that is used by the launch profile.
	//
	// The only valid version is "2021-03-31".
	LaunchProfileProtocolVersion *string `field:"optional" json:"launchProfileProtocolVersion" yaml:"launchProfileProtocolVersion"`
	// The platform of the initialization script, either WINDOWS or LINUX.
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// The method to use when running the initialization script.
	RunContext *string `field:"optional" json:"runContext" yaml:"runContext"`
	// The initialization script.
	Script *string `field:"optional" json:"script" yaml:"script"`
}

// Properties for defining a `CfnStudioComponent`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStudioComponentProps := &cfnStudioComponentProps{
//   	name: jsii.String("name"),
//   	studioId: jsii.String("studioId"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	configuration: &studioComponentConfigurationProperty{
//   		activeDirectoryConfiguration: &activeDirectoryConfigurationProperty{
//   			computerAttributes: []interface{}{
//   				&activeDirectoryComputerAttributeProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			directoryId: jsii.String("directoryId"),
//   			organizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   		},
//   		computeFarmConfiguration: &computeFarmConfigurationProperty{
//   			activeDirectoryUser: jsii.String("activeDirectoryUser"),
//   			endpoint: jsii.String("endpoint"),
//   		},
//   		licenseServiceConfiguration: &licenseServiceConfigurationProperty{
//   			endpoint: jsii.String("endpoint"),
//   		},
//   		sharedFileSystemConfiguration: &sharedFileSystemConfigurationProperty{
//   			endpoint: jsii.String("endpoint"),
//   			fileSystemId: jsii.String("fileSystemId"),
//   			linuxMountPoint: jsii.String("linuxMountPoint"),
//   			shareName: jsii.String("shareName"),
//   			windowsMountDrive: jsii.String("windowsMountDrive"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	ec2SecurityGroupIds: []*string{
//   		jsii.String("ec2SecurityGroupIds"),
//   	},
//   	initializationScripts: []interface{}{
//   		&studioComponentInitializationScriptProperty{
//   			launchProfileProtocolVersion: jsii.String("launchProfileProtocolVersion"),
//   			platform: jsii.String("platform"),
//   			runContext: jsii.String("runContext"),
//   			script: jsii.String("script"),
//   		},
//   	},
//   	scriptParameters: []interface{}{
//   		&scriptParameterKeyValueProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	subtype: jsii.String("subtype"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnStudioComponentProps struct {
	// A friendly name for the studio component resource.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The unique identifier for a studio resource.
	//
	// In Nimble Studio , all other resources are contained in a studio resource.
	StudioId *string `field:"required" json:"studioId" yaml:"studioId"`
	// The type of the studio component.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The configuration of the studio component, based on component type.
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// A human-readable description for the studio component resource.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The EC2 security groups that control access to the studio component.
	Ec2SecurityGroupIds *[]*string `field:"optional" json:"ec2SecurityGroupIds" yaml:"ec2SecurityGroupIds"`
	// Initialization scripts for studio components.
	InitializationScripts interface{} `field:"optional" json:"initializationScripts" yaml:"initializationScripts"`
	// Parameters for the studio component scripts.
	ScriptParameters interface{} `field:"optional" json:"scriptParameters" yaml:"scriptParameters"`
	// The specific subtype of a studio component.
	Subtype *string `field:"optional" json:"subtype" yaml:"subtype"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

// Properties for defining a `CfnStudio`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStudioProps := &cfnStudioProps{
//   	adminRoleArn: jsii.String("adminRoleArn"),
//   	displayName: jsii.String("displayName"),
//   	studioName: jsii.String("studioName"),
//   	userRoleArn: jsii.String("userRoleArn"),
//
//   	// the properties below are optional
//   	studioEncryptionConfiguration: &studioEncryptionConfigurationProperty{
//   		keyType: jsii.String("keyType"),
//
//   		// the properties below are optional
//   		keyArn: jsii.String("keyArn"),
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnStudioProps struct {
	// The IAM role that studio admins assume when logging in to the Nimble Studio portal.
	AdminRoleArn *string `field:"required" json:"adminRoleArn" yaml:"adminRoleArn"`
	// A friendly name for the studio.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The name of the studio, as included in the URL when accessing it in the Nimble Studio portal.
	StudioName *string `field:"required" json:"studioName" yaml:"studioName"`
	// The IAM role that studio users assume when logging in to the Nimble Studio portal.
	UserRoleArn *string `field:"required" json:"userRoleArn" yaml:"userRoleArn"`
	// Configuration of the encryption method that is used for the studio.
	StudioEncryptionConfiguration interface{} `field:"optional" json:"studioEncryptionConfiguration" yaml:"studioEncryptionConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}


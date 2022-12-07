package awsappstream

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::AppStream::Fleet`.
//
// The `AWS::AppStream::Fleet` resource creates a fleet for Amazon AppStream 2.0. A fleet consists of streaming instances that run a specified image when using Always-On or On-Demand.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFleet := awscdk.Aws_appstream.NewCfnFleet(this, jsii.String("MyCfnFleet"), &cfnFleetProps{
//   	instanceType: jsii.String("instanceType"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	computeCapacity: &computeCapacityProperty{
//   		desiredInstances: jsii.Number(123),
//   	},
//   	description: jsii.String("description"),
//   	disconnectTimeoutInSeconds: jsii.Number(123),
//   	displayName: jsii.String("displayName"),
//   	domainJoinInfo: &domainJoinInfoProperty{
//   		directoryName: jsii.String("directoryName"),
//   		organizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   	},
//   	enableDefaultInternetAccess: jsii.Boolean(false),
//   	fleetType: jsii.String("fleetType"),
//   	iamRoleArn: jsii.String("iamRoleArn"),
//   	idleDisconnectTimeoutInSeconds: jsii.Number(123),
//   	imageArn: jsii.String("imageArn"),
//   	imageName: jsii.String("imageName"),
//   	maxConcurrentSessions: jsii.Number(123),
//   	maxUserDurationInSeconds: jsii.Number(123),
//   	platform: jsii.String("platform"),
//   	sessionScriptS3Location: &s3LocationProperty{
//   		s3Bucket: jsii.String("s3Bucket"),
//   		s3Key: jsii.String("s3Key"),
//   	},
//   	streamView: jsii.String("streamView"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	usbDeviceFilterStrings: []*string{
//   		jsii.String("usbDeviceFilterStrings"),
//   	},
//   	vpcConfig: &vpcConfigProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   })
//
type CfnFleet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The desired capacity for the fleet.
	//
	// This is not allowed for Elastic fleets.
	ComputeCapacity() interface{}
	SetComputeCapacity(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description to display.
	Description() *string
	SetDescription(val *string)
	// The amount of time that a streaming session remains active after users disconnect.
	//
	// If users try to reconnect to the streaming session after a disconnection or network interruption within this time interval, they are connected to their previous session. Otherwise, they are connected to a new session with a new streaming instance.
	//
	// Specify a value between 60 and 360000.
	DisconnectTimeoutInSeconds() *float64
	SetDisconnectTimeoutInSeconds(val *float64)
	// The fleet name to display.
	DisplayName() *string
	SetDisplayName(val *string)
	// The name of the directory and organizational unit (OU) to use to join the fleet to a Microsoft Active Directory domain.
	//
	// This is not allowed for Elastic fleets.
	DomainJoinInfo() interface{}
	SetDomainJoinInfo(val interface{})
	// Enables or disables default internet access for the fleet.
	EnableDefaultInternetAccess() interface{}
	SetEnableDefaultInternetAccess(val interface{})
	// The fleet type.
	//
	// - **ALWAYS_ON** - Provides users with instant-on access to their apps. You are charged for all running instances in your fleet, even if no users are streaming apps.
	// - **ON_DEMAND** - Provide users with access to applications after they connect, which takes one to two minutes. You are charged for instance streaming when users are connected and a small hourly fee for instances that are not streaming apps.
	// - **ELASTIC** - The pool of streaming instances is managed by Amazon AppStream 2.0. When a user selects their application or desktop to launch, they will start streaming after the app block has been downloaded and mounted to a streaming instance.
	//
	// *Allowed Values* : `ALWAYS_ON` | `ELASTIC` | `ON_DEMAND`.
	FleetType() *string
	SetFleetType(val *string)
	// The ARN of the IAM role that is applied to the fleet.
	//
	// To assume a role, the fleet instance calls the AWS Security Token Service `AssumeRole` API operation and passes the ARN of the role to use. The operation creates a new session with temporary credentials. AppStream 2.0 retrieves the temporary credentials and creates the *appstream_machine_role* credential profile on the instance.
	//
	// For more information, see [Using an IAM Role to Grant Permissions to Applications and Scripts Running on AppStream 2.0 Streaming Instances](https://docs.aws.amazon.com/appstream2/latest/developerguide/using-iam-roles-to-grant-permissions-to-applications-scripts-streaming-instances.html) in the *Amazon AppStream 2.0 Administration Guide* .
	IamRoleArn() *string
	SetIamRoleArn(val *string)
	// The amount of time that users can be idle (inactive) before they are disconnected from their streaming session and the `DisconnectTimeoutInSeconds` time interval begins.
	//
	// Users are notified before they are disconnected due to inactivity. If they try to reconnect to the streaming session before the time interval specified in `DisconnectTimeoutInSeconds` elapses, they are connected to their previous session. Users are considered idle when they stop providing keyboard or mouse input during their streaming session. File uploads and downloads, audio in, audio out, and pixels changing do not qualify as user activity. If users continue to be idle after the time interval in `IdleDisconnectTimeoutInSeconds` elapses, they are disconnected.
	//
	// To prevent users from being disconnected due to inactivity, specify a value of 0. Otherwise, specify a value between 60 and 3600.
	//
	// If you enable this feature, we recommend that you specify a value that corresponds exactly to a whole number of minutes (for example, 60, 120, and 180). If you don't do this, the value is rounded to the nearest minute. For example, if you specify a value of 70, users are disconnected after 1 minute of inactivity. If you specify a value that is at the midpoint between two different minutes, the value is rounded up. For example, if you specify a value of 90, users are disconnected after 2 minutes of inactivity.
	IdleDisconnectTimeoutInSeconds() *float64
	SetIdleDisconnectTimeoutInSeconds(val *float64)
	// The ARN of the public, private, or shared image to use.
	ImageArn() *string
	SetImageArn(val *string)
	// The name of the image used to create the fleet.
	ImageName() *string
	SetImageName(val *string)
	// The instance type to use when launching fleet instances. The following instance types are available for non-Elastic fleets:.
	//
	// - stream.standard.small
	// - stream.standard.medium
	// - stream.standard.large
	// - stream.compute.large
	// - stream.compute.xlarge
	// - stream.compute.2xlarge
	// - stream.compute.4xlarge
	// - stream.compute.8xlarge
	// - stream.memory.large
	// - stream.memory.xlarge
	// - stream.memory.2xlarge
	// - stream.memory.4xlarge
	// - stream.memory.8xlarge
	// - stream.memory.z1d.large
	// - stream.memory.z1d.xlarge
	// - stream.memory.z1d.2xlarge
	// - stream.memory.z1d.3xlarge
	// - stream.memory.z1d.6xlarge
	// - stream.memory.z1d.12xlarge
	// - stream.graphics-design.large
	// - stream.graphics-design.xlarge
	// - stream.graphics-design.2xlarge
	// - stream.graphics-design.4xlarge
	// - stream.graphics-desktop.2xlarge
	// - stream.graphics.g4dn.xlarge
	// - stream.graphics.g4dn.2xlarge
	// - stream.graphics.g4dn.4xlarge
	// - stream.graphics.g4dn.8xlarge
	// - stream.graphics.g4dn.12xlarge
	// - stream.graphics.g4dn.16xlarge
	// - stream.graphics-pro.4xlarge
	// - stream.graphics-pro.8xlarge
	// - stream.graphics-pro.16xlarge
	//
	// The following instance types are available for Elastic fleets:
	//
	// - stream.standard.small
	// - stream.standard.medium
	InstanceType() *string
	SetInstanceType(val *string)
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
	// The maximum number of concurrent sessions that can be run on an Elastic fleet.
	//
	// This setting is required for Elastic fleets, but is not used for other fleet types.
	MaxConcurrentSessions() *float64
	SetMaxConcurrentSessions(val *float64)
	// The maximum amount of time that a streaming session can remain active, in seconds.
	//
	// If users are still connected to a streaming instance five minutes before this limit is reached, they are prompted to save any open documents before being disconnected. After this time elapses, the instance is terminated and replaced by a new instance.
	//
	// Specify a value between 600 and 360000.
	MaxUserDurationInSeconds() *float64
	SetMaxUserDurationInSeconds(val *float64)
	// A unique name for the fleet.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The platform of the fleet.
	//
	// Platform is a required setting for Elastic fleets, and is not used for other fleet types.
	//
	// *Allowed Values* : `WINDOWS_SERVER_2019` | `AMAZON_LINUX2`.
	Platform() *string
	SetPlatform(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The S3 location of the session scripts configuration zip file.
	//
	// This only applies to Elastic fleets.
	SessionScriptS3Location() interface{}
	SetSessionScriptS3Location(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The AppStream 2.0 view that is displayed to your users when they stream from the fleet. When `APP` is specified, only the windows of applications opened by users display. When `DESKTOP` is specified, the standard desktop that is provided by the operating system displays.
	//
	// The default value is `APP` .
	StreamView() *string
	SetStreamView(val *string)
	// An array of key-value pairs.
	Tags() awscdk.TagManager
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
	// The USB device filter strings that specify which USB devices a user can redirect to the fleet streaming session, when using the Windows native client.
	//
	// This is allowed but not required for Elastic fleets.
	UsbDeviceFilterStrings() *[]*string
	SetUsbDeviceFilterStrings(val *[]*string)
	// The VPC configuration for the fleet.
	//
	// This is required for Elastic fleets, but not required for other fleet types.
	VpcConfig() interface{}
	SetVpcConfig(val interface{})
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

// The jsii proxy struct for CfnFleet
type jsiiProxy_CfnFleet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFleet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) ComputeCapacity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"computeCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) DisconnectTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"disconnectTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) DomainJoinInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"domainJoinInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) EnableDefaultInternetAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDefaultInternetAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) FleetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) IamRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) IdleDisconnectTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleDisconnectTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) ImageArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) MaxConcurrentSessions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentSessions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) MaxUserDurationInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUserDurationInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) SessionScriptS3Location() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionScriptS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) StreamView() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamView",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) UsbDeviceFilterStrings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usbDeviceFilterStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) VpcConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppStream::Fleet`.
func NewCfnFleet(scope constructs.Construct, id *string, props *CfnFleetProps) CfnFleet {
	_init_.Initialize()

	if err := validateNewCfnFleetParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFleet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appstream.CfnFleet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppStream::Fleet`.
func NewCfnFleet_Override(c CfnFleet, scope constructs.Construct, id *string, props *CfnFleetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appstream.CfnFleet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFleet)SetComputeCapacity(val interface{}) {
	if err := j.validateSetComputeCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeCapacity",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetDisconnectTimeoutInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"disconnectTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetDomainJoinInfo(val interface{}) {
	if err := j.validateSetDomainJoinInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainJoinInfo",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetEnableDefaultInternetAccess(val interface{}) {
	if err := j.validateSetEnableDefaultInternetAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDefaultInternetAccess",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetFleetType(val *string) {
	_jsii_.Set(
		j,
		"fleetType",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetIamRoleArn(val *string) {
	_jsii_.Set(
		j,
		"iamRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetIdleDisconnectTimeoutInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"idleDisconnectTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetImageArn(val *string) {
	_jsii_.Set(
		j,
		"imageArn",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetImageName(val *string) {
	_jsii_.Set(
		j,
		"imageName",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetMaxConcurrentSessions(val *float64) {
	_jsii_.Set(
		j,
		"maxConcurrentSessions",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetMaxUserDurationInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"maxUserDurationInSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetPlatform(val *string) {
	_jsii_.Set(
		j,
		"platform",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetSessionScriptS3Location(val interface{}) {
	if err := j.validateSetSessionScriptS3LocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionScriptS3Location",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetStreamView(val *string) {
	_jsii_.Set(
		j,
		"streamView",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetUsbDeviceFilterStrings(val *[]*string) {
	_jsii_.Set(
		j,
		"usbDeviceFilterStrings",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetVpcConfig(val interface{}) {
	if err := j.validateSetVpcConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnFleet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFleet_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appstream.CfnFleet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnFleet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnFleet_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appstream.CfnFleet",
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
func CfnFleet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFleet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appstream.CfnFleet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFleet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appstream.CfnFleet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFleet) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFleet) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFleet) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFleet) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFleet) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFleet) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFleet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFleet) GetAtt(attributeName *string) awscdk.Reference {
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

func (c *jsiiProxy_CfnFleet) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnFleet) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFleet) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFleet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnFleet) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFleet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFleet) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}


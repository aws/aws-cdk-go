package awsgamelift

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::GameLift::Fleet`.
//
// The `AWS::GameLift::Fleet` resource creates an Amazon GameLift (GameLift) fleet to host game servers. A fleet is a set of EC2 instances, each of which can host multiple game sessions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFleet := awscdk.Aws_gamelift.NewCfnFleet(this, jsii.String("MyCfnFleet"), &cfnFleetProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	anywhereConfiguration: &anywhereConfigurationProperty{
//   		cost: jsii.String("cost"),
//   	},
//   	buildId: jsii.String("buildId"),
//   	certificateConfiguration: &certificateConfigurationProperty{
//   		certificateType: jsii.String("certificateType"),
//   	},
//   	computeType: jsii.String("computeType"),
//   	description: jsii.String("description"),
//   	desiredEc2Instances: jsii.Number(123),
//   	ec2InboundPermissions: []interface{}{
//   		&ipPermissionProperty{
//   			fromPort: jsii.Number(123),
//   			ipRange: jsii.String("ipRange"),
//   			protocol: jsii.String("protocol"),
//   			toPort: jsii.Number(123),
//   		},
//   	},
//   	ec2InstanceType: jsii.String("ec2InstanceType"),
//   	fleetType: jsii.String("fleetType"),
//   	instanceRoleArn: jsii.String("instanceRoleArn"),
//   	locations: []interface{}{
//   		&locationConfigurationProperty{
//   			location: jsii.String("location"),
//
//   			// the properties below are optional
//   			locationCapacity: &locationCapacityProperty{
//   				desiredEc2Instances: jsii.Number(123),
//   				maxSize: jsii.Number(123),
//   				minSize: jsii.Number(123),
//   			},
//   		},
//   	},
//   	maxSize: jsii.Number(123),
//   	metricGroups: []*string{
//   		jsii.String("metricGroups"),
//   	},
//   	minSize: jsii.Number(123),
//   	newGameSessionProtectionPolicy: jsii.String("newGameSessionProtectionPolicy"),
//   	peerVpcAwsAccountId: jsii.String("peerVpcAwsAccountId"),
//   	peerVpcId: jsii.String("peerVpcId"),
//   	resourceCreationLimitPolicy: &resourceCreationLimitPolicyProperty{
//   		newGameSessionsPerCreator: jsii.Number(123),
//   		policyPeriodInMinutes: jsii.Number(123),
//   	},
//   	runtimeConfiguration: &runtimeConfigurationProperty{
//   		gameSessionActivationTimeoutSeconds: jsii.Number(123),
//   		maxConcurrentGameSessionActivations: jsii.Number(123),
//   		serverProcesses: []interface{}{
//   			&serverProcessProperty{
//   				concurrentExecutions: jsii.Number(123),
//   				launchPath: jsii.String("launchPath"),
//
//   				// the properties below are optional
//   				parameters: jsii.String("parameters"),
//   			},
//   		},
//   	},
//   	scriptId: jsii.String("scriptId"),
//   })
//
type CfnFleet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// `AWS::GameLift::Fleet.AnywhereConfiguration`.
	AnywhereConfiguration() interface{}
	SetAnywhereConfiguration(val interface{})
	// A unique identifier for the fleet.
	AttrFleetId() *string
	// A unique identifier for a build to be deployed on the new fleet.
	//
	// If you are deploying the fleet with a custom game build, you must specify this property. The build must have been successfully uploaded to Amazon GameLift and be in a `READY` status. This fleet setting cannot be changed once the fleet is created.
	BuildId() *string
	SetBuildId(val *string)
	// Prompts GameLift to generate a TLS/SSL certificate for the fleet.
	//
	// GameLift uses the certificates to encrypt traffic between game clients and the game servers running on GameLift. By default, the `CertificateConfiguration` is `DISABLED` . You can't change this property after you create the fleet.
	//
	// AWS Certificate Manager (ACM) certificates expire after 13 months. Certificate expiration can cause fleets to fail, preventing players from connecting to instances in the fleet. We recommend you replace fleets before 13 months, consider using fleet aliases for a smooth transition.
	//
	// > ACM isn't available in all AWS regions. A fleet creation request with certificate generation enabled in an unsupported Region, fails with a 4xx error. For more information about the supported Regions, see [Supported Regions](https://docs.aws.amazon.com/acm/latest/userguide/acm-regions.html) in the *AWS Certificate Manager User Guide* .
	CertificateConfiguration() interface{}
	SetCertificateConfiguration(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The type of compute resource used to host your game servers.
	//
	// You can use your own compute resources with GameLift Anywhere or use Amazon EC2 instances with managed GameLift.
	ComputeType() *string
	SetComputeType(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A description for the fleet.
	Description() *string
	SetDescription(val *string)
	// The number of EC2 instances that you want this fleet to host.
	//
	// When creating a new fleet, GameLift automatically sets this value to "1" and initiates a single instance. Once the fleet is active, update this value to trigger GameLift to add or remove instances from the fleet.
	DesiredEc2Instances() *float64
	SetDesiredEc2Instances(val *float64)
	// The allowed IP address ranges and port settings that allow inbound traffic to access game sessions on this fleet.
	//
	// If the fleet is hosting a custom game build, this property must be set before players can connect to game sessions. For Realtime Servers fleets, GameLift automatically sets TCP and UDP ranges.
	Ec2InboundPermissions() interface{}
	SetEc2InboundPermissions(val interface{})
	// The GameLift-supported Amazon EC2 instance type to use for all fleet instances.
	//
	// Instance type determines the computing resources that will be used to host your game servers, including CPU, memory, storage, and networking capacity. See [Amazon Elastic Compute Cloud Instance Types](https://docs.aws.amazon.com/ec2/instance-types/) for detailed descriptions of Amazon EC2 instance types.
	Ec2InstanceType() *string
	SetEc2InstanceType(val *string)
	// Indicates whether to use On-Demand or Spot instances for this fleet.
	//
	// By default, this property is set to `ON_DEMAND` . Learn more about when to use [On-Demand versus Spot Instances](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-ec2-instances.html#gamelift-ec2-instances-spot) . This property cannot be changed after the fleet is created.
	FleetType() *string
	SetFleetType(val *string)
	// A unique identifier for an IAM role that manages access to your AWS services.
	//
	// With an instance role ARN set, any application that runs on an instance in this fleet can assume the role, including install scripts, server processes, and daemons (background processes). Create a role or look up a role's ARN by using the [IAM dashboard](https://docs.aws.amazon.com/iam/) in the AWS Management Console . Learn more about using on-box credentials for your game servers at [Access external resources from a game server](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-resources.html) . This property cannot be changed after the fleet is created.
	InstanceRoleArn() *string
	SetInstanceRoleArn(val *string)
	// A set of remote locations to deploy additional instances to and manage as part of the fleet.
	//
	// This parameter can only be used when creating fleets in AWS Regions that support multiple locations. You can add any GameLift-supported AWS Region as a remote location, in the form of an AWS Region code such as `us-west-2` . To create a fleet with instances in the home Region only, omit this parameter.
	Locations() interface{}
	SetLocations(val interface{})
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
	// The maximum number of instances that are allowed in the specified fleet location.
	//
	// If this parameter is not set, the default is 1.
	MaxSize() *float64
	SetMaxSize(val *float64)
	// The name of an AWS CloudWatch metric group to add this fleet to.
	//
	// A metric group is used to aggregate the metrics for multiple fleets. You can specify an existing metric group name or set a new name to create a new metric group. A fleet can be included in only one metric group at a time.
	MetricGroups() *[]*string
	SetMetricGroups(val *[]*string)
	// The minimum number of instances that are allowed in the specified fleet location.
	//
	// If this parameter is not set, the default is 0.
	MinSize() *float64
	SetMinSize(val *float64)
	// A descriptive label that is associated with a fleet.
	//
	// Fleet names do not need to be unique.
	Name() *string
	SetName(val *string)
	// The status of termination protection for active game sessions on the fleet.
	//
	// By default, this property is set to `NoProtection` .
	//
	// - *NoProtection* - Game sessions can be terminated during active gameplay as a result of a scale-down event.
	// - *FullProtection* - Game sessions in `ACTIVE` status cannot be terminated during a scale-down event.
	NewGameSessionProtectionPolicy() *string
	SetNewGameSessionProtectionPolicy(val *string)
	// The tree node.
	Node() constructs.Node
	// Used when peering your GameLift fleet with a VPC, the unique identifier for the AWS account that owns the VPC.
	//
	// You can find your account ID in the AWS Management Console under account settings.
	PeerVpcAwsAccountId() *string
	SetPeerVpcAwsAccountId(val *string)
	// A unique identifier for a VPC with resources to be accessed by your GameLift fleet.
	//
	// The VPC must be in the same Region as your fleet. To look up a VPC ID, use the [VPC Dashboard](https://docs.aws.amazon.com/vpc/) in the AWS Management Console . Learn more about VPC peering in [VPC Peering with GameLift Fleets](https://docs.aws.amazon.com/gamelift/latest/developerguide/vpc-peering.html) .
	PeerVpcId() *string
	SetPeerVpcId(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// A policy that limits the number of game sessions that an individual player can create on instances in this fleet within a specified span of time.
	ResourceCreationLimitPolicy() interface{}
	SetResourceCreationLimitPolicy(val interface{})
	// Instructions for how to launch and maintain server processes on instances in the fleet.
	//
	// The runtime configuration defines one or more server process configurations, each identifying a build executable or Realtime script file and the number of processes of that type to run concurrently.
	//
	// > The `RuntimeConfiguration` parameter is required unless the fleet is being configured using the older parameters `ServerLaunchPath` and `ServerLaunchParameters` , which are still supported for backward compatibility.
	RuntimeConfiguration() interface{}
	SetRuntimeConfiguration(val interface{})
	// The unique identifier for a Realtime configuration script to be deployed on fleet instances.
	//
	// You can use either the script ID or ARN. Scripts must be uploaded to GameLift prior to creating the fleet. This fleet property cannot be changed later.
	//
	// > You can't use the `!Ref` command to reference a script created with a CloudFormation template for the fleet property `ScriptId` . Instead, use `Fn::GetAtt Script.Arn` or `Fn::GetAtt Script.Id` to retrieve either of these properties as input for `ScriptId` . Alternatively, enter a `ScriptId` string manually.
	ScriptId() *string
	SetScriptId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnFleet
type jsiiProxy_CfnFleet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFleet) AnywhereConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anywhereConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) AttrFleetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFleetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) BuildId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) CertificateConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateConfiguration",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_CfnFleet) ComputeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeType",
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

func (j *jsiiProxy_CfnFleet) DesiredEc2Instances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredEc2Instances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) Ec2InboundPermissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2InboundPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) Ec2InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2InstanceType",
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

func (j *jsiiProxy_CfnFleet) InstanceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) Locations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"locations",
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

func (j *jsiiProxy_CfnFleet) MaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) MetricGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metricGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) MinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSize",
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

func (j *jsiiProxy_CfnFleet) NewGameSessionProtectionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newGameSessionProtectionPolicy",
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

func (j *jsiiProxy_CfnFleet) PeerVpcAwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVpcAwsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) PeerVpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVpcId",
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

func (j *jsiiProxy_CfnFleet) ResourceCreationLimitPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceCreationLimitPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) RuntimeConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runtimeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleet) ScriptId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptId",
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


// Create a new `AWS::GameLift::Fleet`.
func NewCfnFleet(scope constructs.Construct, id *string, props *CfnFleetProps) CfnFleet {
	_init_.Initialize()

	if err := validateNewCfnFleetParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFleet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_gamelift.CfnFleet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GameLift::Fleet`.
func NewCfnFleet_Override(c CfnFleet, scope constructs.Construct, id *string, props *CfnFleetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_gamelift.CfnFleet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFleet)SetAnywhereConfiguration(val interface{}) {
	if err := j.validateSetAnywhereConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anywhereConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetBuildId(val *string) {
	_jsii_.Set(
		j,
		"buildId",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetCertificateConfiguration(val interface{}) {
	if err := j.validateSetCertificateConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetComputeType(val *string) {
	_jsii_.Set(
		j,
		"computeType",
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

func (j *jsiiProxy_CfnFleet)SetDesiredEc2Instances(val *float64) {
	_jsii_.Set(
		j,
		"desiredEc2Instances",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetEc2InboundPermissions(val interface{}) {
	if err := j.validateSetEc2InboundPermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ec2InboundPermissions",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetEc2InstanceType(val *string) {
	_jsii_.Set(
		j,
		"ec2InstanceType",
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

func (j *jsiiProxy_CfnFleet)SetInstanceRoleArn(val *string) {
	_jsii_.Set(
		j,
		"instanceRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetLocations(val interface{}) {
	if err := j.validateSetLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locations",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetMaxSize(val *float64) {
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetMetricGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"metricGroups",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetMinSize(val *float64) {
	_jsii_.Set(
		j,
		"minSize",
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

func (j *jsiiProxy_CfnFleet)SetNewGameSessionProtectionPolicy(val *string) {
	_jsii_.Set(
		j,
		"newGameSessionProtectionPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetPeerVpcAwsAccountId(val *string) {
	_jsii_.Set(
		j,
		"peerVpcAwsAccountId",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetPeerVpcId(val *string) {
	_jsii_.Set(
		j,
		"peerVpcId",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetResourceCreationLimitPolicy(val interface{}) {
	if err := j.validateSetResourceCreationLimitPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceCreationLimitPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetRuntimeConfiguration(val interface{}) {
	if err := j.validateSetRuntimeConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnFleet)SetScriptId(val *string) {
	_jsii_.Set(
		j,
		"scriptId",
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
		"aws-cdk-lib.aws_gamelift.CfnFleet",
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
		"aws-cdk-lib.aws_gamelift.CfnFleet",
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
		"aws-cdk-lib.aws_gamelift.CfnFleet",
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
		"aws-cdk-lib.aws_gamelift.CfnFleet",
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

func (c *jsiiProxy_CfnFleet) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
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

func (c *jsiiProxy_CfnFleet) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnFleet) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFleet) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
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

func (c *jsiiProxy_CfnFleet) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
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

func (c *jsiiProxy_CfnFleet) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
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


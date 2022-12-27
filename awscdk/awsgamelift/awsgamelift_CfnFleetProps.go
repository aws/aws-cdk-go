package awsgamelift


// Properties for defining a `CfnFleet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFleetProps := &cfnFleetProps{
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
//   }
//
type CfnFleetProps struct {
	// A descriptive label that is associated with a fleet.
	//
	// Fleet names do not need to be unique.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::GameLift::Fleet.AnywhereConfiguration`.
	AnywhereConfiguration interface{} `field:"optional" json:"anywhereConfiguration" yaml:"anywhereConfiguration"`
	// A unique identifier for a build to be deployed on the new fleet.
	//
	// If you are deploying the fleet with a custom game build, you must specify this property. The build must have been successfully uploaded to Amazon GameLift and be in a `READY` status. This fleet setting cannot be changed once the fleet is created.
	BuildId *string `field:"optional" json:"buildId" yaml:"buildId"`
	// Prompts GameLift to generate a TLS/SSL certificate for the fleet.
	//
	// GameLift uses the certificates to encrypt traffic between game clients and the game servers running on GameLift. By default, the `CertificateConfiguration` is `DISABLED` . You can't change this property after you create the fleet.
	//
	// AWS Certificate Manager (ACM) certificates expire after 13 months. Certificate expiration can cause fleets to fail, preventing players from connecting to instances in the fleet. We recommend you replace fleets before 13 months, consider using fleet aliases for a smooth transition.
	//
	// > ACM isn't available in all AWS regions. A fleet creation request with certificate generation enabled in an unsupported Region, fails with a 4xx error. For more information about the supported Regions, see [Supported Regions](https://docs.aws.amazon.com/acm/latest/userguide/acm-regions.html) in the *AWS Certificate Manager User Guide* .
	CertificateConfiguration interface{} `field:"optional" json:"certificateConfiguration" yaml:"certificateConfiguration"`
	// `AWS::GameLift::Fleet.ComputeType`.
	ComputeType *string `field:"optional" json:"computeType" yaml:"computeType"`
	// A human-readable description of the fleet.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The number of EC2 instances that you want this fleet to host.
	//
	// When creating a new fleet, GameLift automatically sets this value to "1" and initiates a single instance. Once the fleet is active, update this value to trigger GameLift to add or remove instances from the fleet.
	DesiredEc2Instances *float64 `field:"optional" json:"desiredEc2Instances" yaml:"desiredEc2Instances"`
	// The allowed IP address ranges and port settings that allow inbound traffic to access game sessions on this fleet.
	//
	// If the fleet is hosting a custom game build, this property must be set before players can connect to game sessions. For Realtime Servers fleets, GameLift automatically sets TCP and UDP ranges.
	Ec2InboundPermissions interface{} `field:"optional" json:"ec2InboundPermissions" yaml:"ec2InboundPermissions"`
	// The GameLift-supported Amazon EC2 instance type to use for all fleet instances.
	//
	// Instance type determines the computing resources that will be used to host your game servers, including CPU, memory, storage, and networking capacity. See [Amazon Elastic Compute Cloud Instance Types](https://docs.aws.amazon.com/ec2/instance-types/) for detailed descriptions of Amazon EC2 instance types.
	Ec2InstanceType *string `field:"optional" json:"ec2InstanceType" yaml:"ec2InstanceType"`
	// Indicates whether to use On-Demand or Spot instances for this fleet.
	//
	// By default, this property is set to `ON_DEMAND` . Learn more about when to use [On-Demand versus Spot Instances](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-ec2-instances.html#gamelift-ec2-instances-spot) . This property cannot be changed after the fleet is created.
	FleetType *string `field:"optional" json:"fleetType" yaml:"fleetType"`
	// A unique identifier for an IAM role that manages access to your AWS services.
	//
	// With an instance role ARN set, any application that runs on an instance in this fleet can assume the role, including install scripts, server processes, and daemons (background processes). Create a role or look up a role's ARN by using the [IAM dashboard](https://docs.aws.amazon.com/iam/) in the AWS Management Console . Learn more about using on-box credentials for your game servers at [Access external resources from a game server](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-resources.html) . This property cannot be changed after the fleet is created.
	InstanceRoleArn *string `field:"optional" json:"instanceRoleArn" yaml:"instanceRoleArn"`
	// A set of remote locations to deploy additional instances to and manage as part of the fleet.
	//
	// This parameter can only be used when creating fleets in AWS Regions that support multiple locations. You can add any GameLift-supported AWS Region as a remote location, in the form of an AWS Region code such as `us-west-2` . To create a fleet with instances in the home Region only, omit this parameter.
	Locations interface{} `field:"optional" json:"locations" yaml:"locations"`
	// The maximum number of instances that are allowed in the specified fleet location.
	//
	// If this parameter is not set, the default is 1.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// The name of an AWS CloudWatch metric group to add this fleet to.
	//
	// A metric group is used to aggregate the metrics for multiple fleets. You can specify an existing metric group name or set a new name to create a new metric group. A fleet can be included in only one metric group at a time.
	MetricGroups *[]*string `field:"optional" json:"metricGroups" yaml:"metricGroups"`
	// The minimum number of instances that are allowed in the specified fleet location.
	//
	// If this parameter is not set, the default is 0.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// The status of termination protection for active game sessions on the fleet.
	//
	// By default, this property is set to `NoProtection` .
	//
	// - *NoProtection* - Game sessions can be terminated during active gameplay as a result of a scale-down event.
	// - *FullProtection* - Game sessions in `ACTIVE` status cannot be terminated during a scale-down event.
	NewGameSessionProtectionPolicy *string `field:"optional" json:"newGameSessionProtectionPolicy" yaml:"newGameSessionProtectionPolicy"`
	// Used when peering your GameLift fleet with a VPC, the unique identifier for the AWS account that owns the VPC.
	//
	// You can find your account ID in the AWS Management Console under account settings.
	PeerVpcAwsAccountId *string `field:"optional" json:"peerVpcAwsAccountId" yaml:"peerVpcAwsAccountId"`
	// A unique identifier for a VPC with resources to be accessed by your GameLift fleet.
	//
	// The VPC must be in the same Region as your fleet. To look up a VPC ID, use the [VPC Dashboard](https://docs.aws.amazon.com/vpc/) in the AWS Management Console . Learn more about VPC peering in [VPC Peering with GameLift Fleets](https://docs.aws.amazon.com/gamelift/latest/developerguide/vpc-peering.html) .
	PeerVpcId *string `field:"optional" json:"peerVpcId" yaml:"peerVpcId"`
	// A policy that limits the number of game sessions that an individual player can create on instances in this fleet within a specified span of time.
	ResourceCreationLimitPolicy interface{} `field:"optional" json:"resourceCreationLimitPolicy" yaml:"resourceCreationLimitPolicy"`
	// Instructions for how to launch and maintain server processes on instances in the fleet.
	//
	// The runtime configuration defines one or more server process configurations, each identifying a build executable or Realtime script file and the number of processes of that type to run concurrently.
	//
	// > The `RuntimeConfiguration` parameter is required unless the fleet is being configured using the older parameters `ServerLaunchPath` and `ServerLaunchParameters` , which are still supported for backward compatibility.
	RuntimeConfiguration interface{} `field:"optional" json:"runtimeConfiguration" yaml:"runtimeConfiguration"`
	// The unique identifier for a Realtime configuration script to be deployed on fleet instances.
	//
	// You can use either the script ID or ARN. Scripts must be uploaded to GameLift prior to creating the fleet. This fleet property cannot be changed later.
	//
	// > You can't use the `!Ref` command to reference a script created with a CloudFormation template for the fleet property `ScriptId` . Instead, use `Fn::GetAtt Script.Arn` or `Fn::GetAtt Script.Id` to retrieve either of these properties as input for `ScriptId` . Alternatively, enter a `ScriptId` string manually.
	ScriptId *string `field:"optional" json:"scriptId" yaml:"scriptId"`
}


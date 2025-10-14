package awsgamelift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnFleet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFleetProps := &CfnFleetProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AnywhereConfiguration: &AnywhereConfigurationProperty{
//   		Cost: jsii.String("cost"),
//   	},
//   	ApplyCapacity: jsii.String("applyCapacity"),
//   	BuildId: jsii.String("buildId"),
//   	CertificateConfiguration: &CertificateConfigurationProperty{
//   		CertificateType: jsii.String("certificateType"),
//   	},
//   	ComputeType: jsii.String("computeType"),
//   	Description: jsii.String("description"),
//   	DesiredEc2Instances: jsii.Number(123),
//   	Ec2InboundPermissions: []interface{}{
//   		&IpPermissionProperty{
//   			FromPort: jsii.Number(123),
//   			IpRange: jsii.String("ipRange"),
//   			Protocol: jsii.String("protocol"),
//   			ToPort: jsii.Number(123),
//   		},
//   	},
//   	Ec2InstanceType: jsii.String("ec2InstanceType"),
//   	FleetType: jsii.String("fleetType"),
//   	InstanceRoleArn: jsii.String("instanceRoleArn"),
//   	InstanceRoleCredentialsProvider: jsii.String("instanceRoleCredentialsProvider"),
//   	Locations: []interface{}{
//   		&LocationConfigurationProperty{
//   			Location: jsii.String("location"),
//
//   			// the properties below are optional
//   			LocationCapacity: &LocationCapacityProperty{
//   				MaxSize: jsii.Number(123),
//   				MinSize: jsii.Number(123),
//
//   				// the properties below are optional
//   				DesiredEc2Instances: jsii.Number(123),
//   			},
//   		},
//   	},
//   	LogPaths: []*string{
//   		jsii.String("logPaths"),
//   	},
//   	MaxSize: jsii.Number(123),
//   	MetricGroups: []*string{
//   		jsii.String("metricGroups"),
//   	},
//   	MinSize: jsii.Number(123),
//   	NewGameSessionProtectionPolicy: jsii.String("newGameSessionProtectionPolicy"),
//   	PeerVpcAwsAccountId: jsii.String("peerVpcAwsAccountId"),
//   	PeerVpcId: jsii.String("peerVpcId"),
//   	ResourceCreationLimitPolicy: &ResourceCreationLimitPolicyProperty{
//   		NewGameSessionsPerCreator: jsii.Number(123),
//   		PolicyPeriodInMinutes: jsii.Number(123),
//   	},
//   	RuntimeConfiguration: &RuntimeConfigurationProperty{
//   		GameSessionActivationTimeoutSeconds: jsii.Number(123),
//   		MaxConcurrentGameSessionActivations: jsii.Number(123),
//   		ServerProcesses: []interface{}{
//   			&ServerProcessProperty{
//   				ConcurrentExecutions: jsii.Number(123),
//   				LaunchPath: jsii.String("launchPath"),
//
//   				// the properties below are optional
//   				Parameters: jsii.String("parameters"),
//   			},
//   		},
//   	},
//   	ScalingPolicies: []interface{}{
//   		&ScalingPolicyProperty{
//   			MetricName: jsii.String("metricName"),
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			ComparisonOperator: jsii.String("comparisonOperator"),
//   			EvaluationPeriods: jsii.Number(123),
//   			Location: jsii.String("location"),
//   			PolicyType: jsii.String("policyType"),
//   			ScalingAdjustment: jsii.Number(123),
//   			ScalingAdjustmentType: jsii.String("scalingAdjustmentType"),
//   			Status: jsii.String("status"),
//   			TargetConfiguration: &TargetConfigurationProperty{
//   				TargetValue: jsii.Number(123),
//   			},
//   			Threshold: jsii.Number(123),
//   			UpdateStatus: jsii.String("updateStatus"),
//   		},
//   	},
//   	ScriptId: jsii.String("scriptId"),
//   	ServerLaunchParameters: jsii.String("serverLaunchParameters"),
//   	ServerLaunchPath: jsii.String("serverLaunchPath"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html
//
type CfnFleetProps struct {
	// A descriptive label that is associated with a fleet.
	//
	// Fleet names do not need to be unique.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Amazon GameLift Servers Anywhere configuration options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-anywhereconfiguration
	//
	AnywhereConfiguration interface{} `field:"optional" json:"anywhereConfiguration" yaml:"anywhereConfiguration"`
	// Current resource capacity settings for managed EC2 fleets and managed container fleets.
	//
	// For multi-location fleets, location values might refer to a fleet's remote location or its home Region.
	//
	// *Returned by:* [DescribeFleetCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetCapacity.html) , [DescribeFleetLocationCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetLocationCapacity.html) , [UpdateFleetCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_UpdateFleetCapacity.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-applycapacity
	//
	ApplyCapacity *string `field:"optional" json:"applyCapacity" yaml:"applyCapacity"`
	// A unique identifier for a build to be deployed on the new fleet.
	//
	// If you are deploying the fleet with a custom game build, you must specify this property. The build must have been successfully uploaded to Amazon GameLift and be in a `READY` status. This fleet setting cannot be changed once the fleet is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-buildid
	//
	BuildId *string `field:"optional" json:"buildId" yaml:"buildId"`
	// Prompts Amazon GameLift Servers to generate a TLS/SSL certificate for the fleet.
	//
	// Amazon GameLift Servers uses the certificates to encrypt traffic between game clients and the game servers running on Amazon GameLift Servers. By default, the `CertificateConfiguration` is `DISABLED` . You can't change this property after you create the fleet.
	//
	// Certificate Manager (ACM) certificates expire after 13 months. Certificate expiration can cause fleets to fail, preventing players from connecting to instances in the fleet. We recommend you replace fleets before 13 months, consider using fleet aliases for a smooth transition.
	//
	// > ACM isn't available in all AWS regions. A fleet creation request with certificate generation enabled in an unsupported Region, fails with a 4xx error. For more information about the supported Regions, see [Supported Regions](https://docs.aws.amazon.com/acm/latest/userguide/acm-regions.html) in the *Certificate Manager User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-certificateconfiguration
	//
	CertificateConfiguration interface{} `field:"optional" json:"certificateConfiguration" yaml:"certificateConfiguration"`
	// The type of compute resource used to host your game servers.
	//
	// - `EC2` – The game server build is deployed to Amazon EC2 instances for cloud hosting. This is the default setting.
	// - `ANYWHERE` – Game servers and supporting software are deployed to compute resources that you provide and manage. With this compute type, you can also set the `AnywhereConfiguration` parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-computetype
	//
	ComputeType *string `field:"optional" json:"computeType" yaml:"computeType"`
	// A description for the fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// [DEPRECATED] The number of EC2 instances that you want this fleet to host.
	//
	// When creating a new fleet, GameLift automatically sets this value to "1" and initiates a single instance. Once the fleet is active, update this value to trigger GameLift to add or remove instances from the fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-desiredec2instances
	//
	// Deprecated: this property has been deprecated.
	DesiredEc2Instances *float64 `field:"optional" json:"desiredEc2Instances" yaml:"desiredEc2Instances"`
	// The IP address ranges and port settings that allow inbound traffic to access game server processes and other processes on this fleet.
	//
	// Set this parameter for managed EC2 fleets. You can leave this parameter empty when creating the fleet, but you must call [](https://docs.aws.amazon.com/gamelift/latest/apireference/API_UpdateFleetPortSettings) to set it before players can connect to game sessions. As a best practice, we recommend opening ports for remote access only when you need them and closing them when you're finished. For Amazon GameLift Servers Realtime fleets, Amazon GameLift Servers automatically sets TCP and UDP ranges.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-ec2inboundpermissions
	//
	Ec2InboundPermissions interface{} `field:"optional" json:"ec2InboundPermissions" yaml:"ec2InboundPermissions"`
	// The Amazon GameLift Servers-supported Amazon EC2 instance type to use with managed EC2 fleets.
	//
	// Instance type determines the computing resources that will be used to host your game servers, including CPU, memory, storage, and networking capacity. See [Amazon Elastic Compute Cloud Instance Types](https://docs.aws.amazon.com/ec2/instance-types/) for detailed descriptions of Amazon EC2 instance types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-ec2instancetype
	//
	Ec2InstanceType *string `field:"optional" json:"ec2InstanceType" yaml:"ec2InstanceType"`
	// Indicates whether to use On-Demand or Spot instances for this fleet.
	//
	// By default, this property is set to `ON_DEMAND` . Learn more about when to use [On-Demand versus Spot Instances](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-ec2-instances.html#gamelift-ec2-instances-spot) . This fleet property can't be changed after the fleet is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-fleettype
	//
	FleetType *string `field:"optional" json:"fleetType" yaml:"fleetType"`
	// A unique identifier for an IAM role that manages access to your AWS services.
	//
	// With an instance role ARN set, any application that runs on an instance in this fleet can assume the role, including install scripts, server processes, and daemons (background processes). Create a role or look up a role's ARN by using the [IAM dashboard](https://docs.aws.amazon.com/iam/) in the AWS Management Console . Learn more about using on-box credentials for your game servers at [Access external resources from a game server](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-resources.html) . This attribute is used with fleets where `ComputeType` is `EC2` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-instancerolearn
	//
	InstanceRoleArn *string `field:"optional" json:"instanceRoleArn" yaml:"instanceRoleArn"`
	// Indicates that fleet instances maintain a shared credentials file for the IAM role defined in `InstanceRoleArn` .
	//
	// Shared credentials allow applications that are deployed with the game server executable to communicate with other AWS resources. This property is used only when the game server is integrated with the server SDK version 5.x. For more information about using shared credentials, see [Communicate with other AWS resources from your fleets](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-resources.html) . This attribute is used with fleets where `ComputeType` is `EC2` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-instancerolecredentialsprovider
	//
	InstanceRoleCredentialsProvider *string `field:"optional" json:"instanceRoleCredentialsProvider" yaml:"instanceRoleCredentialsProvider"`
	// A set of remote locations to deploy additional instances to and manage as a multi-location fleet.
	//
	// Use this parameter when creating a fleet in AWS Regions that support multiple locations. You can add any AWS Region or Local Zone that's supported by Amazon GameLift Servers. Provide a list of one or more AWS Region codes, such as `us-west-2` , or Local Zone names. When using this parameter, Amazon GameLift Servers requires you to include your home location in the request. For a list of supported Regions and Local Zones, see [Amazon GameLift Servers service locations](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-regions.html) for managed hosting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-locations
	//
	Locations interface{} `field:"optional" json:"locations" yaml:"locations"`
	// This parameter is no longer used.
	//
	// When hosting a custom game build, specify where Amazon GameLift should store log files using the Amazon GameLift server API call ProcessReady().
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-logpaths
	//
	// Deprecated: this property has been deprecated.
	LogPaths *[]*string `field:"optional" json:"logPaths" yaml:"logPaths"`
	// [DEPRECATED] The maximum value that is allowed for the fleet's instance count.
	//
	// When creating a new fleet, GameLift automatically sets this value to "1". Once the fleet is active, you can change this value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-maxsize
	//
	// Deprecated: this property has been deprecated.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// The name of an AWS CloudWatch metric group to add this fleet to.
	//
	// A metric group is used to aggregate the metrics for multiple fleets. You can specify an existing metric group name or set a new name to create a new metric group. A fleet can be included in only one metric group at a time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-metricgroups
	//
	MetricGroups *[]*string `field:"optional" json:"metricGroups" yaml:"metricGroups"`
	// [DEPRECATED] The minimum value allowed for the fleet's instance count.
	//
	// When creating a new fleet, GameLift automatically sets this value to "0". After the fleet is active, you can change this value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-minsize
	//
	// Deprecated: this property has been deprecated.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// The status of termination protection for active game sessions on the fleet.
	//
	// By default, this property is set to `NoProtection` .
	//
	// - *NoProtection* - Game sessions can be terminated during active gameplay as a result of a scale-down event.
	// - *FullProtection* - Game sessions in `ACTIVE` status cannot be terminated during a scale-down event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-newgamesessionprotectionpolicy
	//
	NewGameSessionProtectionPolicy *string `field:"optional" json:"newGameSessionProtectionPolicy" yaml:"newGameSessionProtectionPolicy"`
	// Used when peering your Amazon GameLift Servers fleet with a VPC, the unique identifier for the AWS account that owns the VPC.
	//
	// You can find your account ID in the AWS Management Console under account settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-peervpcawsaccountid
	//
	PeerVpcAwsAccountId *string `field:"optional" json:"peerVpcAwsAccountId" yaml:"peerVpcAwsAccountId"`
	// A unique identifier for a VPC with resources to be accessed by your Amazon GameLift Servers fleet.
	//
	// The VPC must be in the same Region as your fleet. To look up a VPC ID, use the [VPC Dashboard](https://docs.aws.amazon.com/vpc/) in the AWS Management Console . Learn more about VPC peering in [VPC Peering with Amazon GameLift Servers Fleets](https://docs.aws.amazon.com/gamelift/latest/developerguide/vpc-peering.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-peervpcid
	//
	PeerVpcId *string `field:"optional" json:"peerVpcId" yaml:"peerVpcId"`
	// A policy that limits the number of game sessions that an individual player can create on instances in this fleet within a specified span of time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-resourcecreationlimitpolicy
	//
	ResourceCreationLimitPolicy interface{} `field:"optional" json:"resourceCreationLimitPolicy" yaml:"resourceCreationLimitPolicy"`
	// Instructions for how to launch and maintain server processes on instances in the fleet.
	//
	// The runtime configuration defines one or more server process configurations, each identifying a build executable or Realtime script file and the number of processes of that type to run concurrently.
	//
	// > The `RuntimeConfiguration` parameter is required unless the fleet is being configured using the older parameters `ServerLaunchPath` and `ServerLaunchParameters` , which are still supported for backward compatibility.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-runtimeconfiguration
	//
	RuntimeConfiguration interface{} `field:"optional" json:"runtimeConfiguration" yaml:"runtimeConfiguration"`
	// Rule that controls how a fleet is scaled.
	//
	// Scaling policies are uniquely identified by the combination of name and fleet ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-scalingpolicies
	//
	ScalingPolicies interface{} `field:"optional" json:"scalingPolicies" yaml:"scalingPolicies"`
	// The unique identifier for a Realtime configuration script to be deployed on fleet instances.
	//
	// You can use either the script ID or ARN. Scripts must be uploaded to Amazon GameLift Servers prior to creating the fleet. This fleet property cannot be changed later.
	//
	// > You can't use the `!Ref` command to reference a script created with a CloudFormation template for the fleet property `ScriptId` . Instead, use `Fn::GetAtt Script.Arn` or `Fn::GetAtt Script.Id` to retrieve either of these properties as input for `ScriptId` . Alternatively, enter a `ScriptId` string manually.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-scriptid
	//
	ScriptId *string `field:"optional" json:"scriptId" yaml:"scriptId"`
	// This parameter is no longer used but is retained for backward compatibility.
	//
	// Instead, specify server launch parameters in the RuntimeConfiguration parameter. A request must specify either a runtime configuration or values for both ServerLaunchParameters and ServerLaunchPath.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-serverlaunchparameters
	//
	// Deprecated: this property has been deprecated.
	ServerLaunchParameters *string `field:"optional" json:"serverLaunchParameters" yaml:"serverLaunchParameters"`
	// This parameter is no longer used.
	//
	// Instead, specify a server launch path using the RuntimeConfiguration parameter. Requests that specify a server launch path and launch parameters instead of a runtime configuration will continue to work.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-serverlaunchpath
	//
	// Deprecated: this property has been deprecated.
	ServerLaunchPath *string `field:"optional" json:"serverLaunchPath" yaml:"serverLaunchPath"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


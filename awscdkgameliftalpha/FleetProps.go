// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for a new Gamelift fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import gamelift_alpha "github.com/aws/aws-cdk-go/awscdkgameliftalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var instanceType instanceType
//   var role role
//   var vpc vpc
//
//   fleetProps := &FleetProps{
//   	FleetName: jsii.String("fleetName"),
//   	InstanceType: instanceType,
//   	RuntimeConfiguration: &RuntimeConfiguration{
//   		ServerProcesses: []serverProcess{
//   			&serverProcess{
//   				LaunchPath: jsii.String("launchPath"),
//
//   				// the properties below are optional
//   				ConcurrentExecutions: jsii.Number(123),
//   				Parameters: jsii.String("parameters"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		GameSessionActivationTimeout: cdk.Duration_Minutes(jsii.Number(30)),
//   		MaxConcurrentGameSessionActivations: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	DesiredCapacity: jsii.Number(123),
//   	Locations: []location{
//   		&location{
//   			Region: jsii.String("region"),
//
//   			// the properties below are optional
//   			Capacity: &LocationCapacity{
//   				DesiredCapacity: jsii.Number(123),
//   				MaxSize: jsii.Number(123),
//   				MinSize: jsii.Number(123),
//   			},
//   		},
//   	},
//   	MaxSize: jsii.Number(123),
//   	MetricGroup: jsii.String("metricGroup"),
//   	MinSize: jsii.Number(123),
//   	PeerVpc: vpc,
//   	ProtectNewGameSession: jsii.Boolean(false),
//   	ResourceCreationLimitPolicy: &ResourceCreationLimitPolicy{
//   		NewGameSessionsPerCreator: jsii.Number(123),
//   		PolicyPeriod: cdk.Duration_*Minutes(jsii.Number(30)),
//   	},
//   	Role: role,
//   	UseCertificate: jsii.Boolean(false),
//   	UseSpot: jsii.Boolean(false),
//   }
//
// Experimental.
type FleetProps struct {
	// A descriptive label that is associated with a fleet.
	//
	// Fleet names do not need to be unique.
	// Experimental.
	FleetName *string `field:"required" json:"fleetName" yaml:"fleetName"`
	// The GameLift-supported Amazon EC2 instance type to use for all fleet instances.
	//
	// Instance type determines the computing resources that will be used to host your game servers, including CPU, memory, storage, and networking capacity.
	// See: http://aws.amazon.com/ec2/instance-types/ for detailed descriptions of Amazon EC2 instance types.
	//
	// Experimental.
	InstanceType awsec2.InstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
	// A collection of server process configurations that describe the set of processes to run on each instance in a fleet.
	//
	// Server processes run either an executable in a custom game build or a Realtime Servers script.
	// GameLift launches the configured processes, manages their life cycle, and replaces them as needed.
	// Each instance checks regularly for an updated runtime configuration.
	//
	// A GameLift instance is limited to 50 processes running concurrently.
	// To calculate the total number of processes in a runtime configuration, add the values of the ConcurrentExecutions parameter for each ServerProcess.
	// See: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-multiprocess.html
	//
	// Experimental.
	RuntimeConfiguration *RuntimeConfiguration `field:"required" json:"runtimeConfiguration" yaml:"runtimeConfiguration"`
	// A human-readable description of the fleet.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The number of EC2 instances that you want this fleet to host.
	//
	// When creating a new fleet, GameLift automatically sets this value to "1" and initiates a single instance.
	// Once the fleet is active, update this value to trigger GameLift to add or remove instances from the fleet.
	// Experimental.
	DesiredCapacity *float64 `field:"optional" json:"desiredCapacity" yaml:"desiredCapacity"`
	// A set of remote locations to deploy additional instances to and manage as part of the fleet.
	//
	// This parameter can only be used when creating fleets in AWS Regions that support multiple locations.
	// You can add any GameLift-supported AWS Region as a remote location, in the form of an AWS Region code such as `us-west-2`.
	// To create a fleet with instances in the home region only, omit this parameter.
	// Experimental.
	Locations *[]*Location `field:"optional" json:"locations" yaml:"locations"`
	// The maximum number of instances that are allowed in the specified fleet location.
	// Experimental.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// The name of an AWS CloudWatch metric group to add this fleet to.
	//
	// A metric group is used to aggregate the metrics for multiple fleets.
	// You can specify an existing metric group name or set a new name to create a new metric group.
	// A fleet can be included in only one metric group at a time.
	// Experimental.
	MetricGroup *string `field:"optional" json:"metricGroup" yaml:"metricGroup"`
	// The minimum number of instances that are allowed in the specified fleet location.
	// Experimental.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// A VPC peering connection between your GameLift-hosted game servers and your other non-GameLift resources.
	//
	// Use Amazon Virtual Private Cloud (VPC) peering connections to enable your game servers to communicate directly and privately with your other AWS resources, such as a web service or a repository.
	// You can establish VPC peering with any resources that run on AWS and are managed by an AWS account that you have access to.
	// The VPC must be in the same Region as your fleet.
	//
	// Warning:
	// Be sure to create a VPC Peering authorization through Gamelift Service API.
	// See: https://docs.aws.amazon.com/gamelift/latest/developerguide/vpc-peering.html
	//
	// Experimental.
	PeerVpc awsec2.IVpc `field:"optional" json:"peerVpc" yaml:"peerVpc"`
	// The status of termination protection for active game sessions on the fleet.
	//
	// By default, new game sessions are protected and cannot be terminated during a scale-down event.
	// Experimental.
	ProtectNewGameSession *bool `field:"optional" json:"protectNewGameSession" yaml:"protectNewGameSession"`
	// A policy that limits the number of game sessions that an individual player can create on instances in this fleet within a specified span of time.
	// Experimental.
	ResourceCreationLimitPolicy *ResourceCreationLimitPolicy `field:"optional" json:"resourceCreationLimitPolicy" yaml:"resourceCreationLimitPolicy"`
	// The IAM role assumed by GameLift fleet instances to access AWS ressources.
	//
	// With a role set, any application that runs on an instance in this fleet can assume the role, including install scripts, server processes, and daemons (background processes).
	// If providing a custom role, it needs to trust the GameLift service principal (gamelift.amazonaws.com).
	// No permission is required by default.
	//
	// This property cannot be changed after the fleet is created.
	// See: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-resources.html
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Prompts GameLift to generate a TLS/SSL certificate for the fleet.
	//
	// GameLift uses the certificates to encrypt traffic between game clients and the game servers running on GameLift.
	//
	// You can't change this property after you create the fleet.
	//
	// Additionnal info:
	// AWS Certificate Manager (ACM) certificates expire after 13 months.
	// Certificate expiration can cause fleets to fail, preventing players from connecting to instances in the fleet.
	// We recommend you replace fleets before 13 months, consider using fleet aliases for a smooth transition.
	// Experimental.
	UseCertificate *bool `field:"optional" json:"useCertificate" yaml:"useCertificate"`
	// Indicates whether to use On-Demand or Spot instances for this fleet. By default, fleet use on demand capacity.
	//
	// This property cannot be changed after the fleet is created.
	// See: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-ec2-instances.html#gamelift-ec2-instances-spot
	//
	// Experimental.
	UseSpot *bool `field:"optional" json:"useSpot" yaml:"useSpot"`
}


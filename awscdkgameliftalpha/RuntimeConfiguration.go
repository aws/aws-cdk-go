// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// A collection of server process configurations that describe the set of processes to run on each instance in a fleet.
//
// Server processes run either an executable in a custom game build or a Realtime Servers script.
// GameLift launches the configured processes, manages their life cycle, and replaces them as needed.
// Each instance checks regularly for an updated runtime configuration.
//
// A GameLift instance is limited to 50 processes running concurrently.
// To calculate the total number of processes in a runtime configuration, add the values of the `ConcurrentExecutions` parameter for each `ServerProcess`.
//
// Example:
//   var build build
//
//   // Server processes can be delcared in a declarative way through the constructor
//   fleet := gamelift.NewBuildFleet(this, jsii.String("Game server fleet"), &BuildFleetProps{
//   	FleetName: jsii.String("test-fleet"),
//   	Content: build,
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_C4, ec2.InstanceSize_LARGE),
//   	RuntimeConfiguration: &RuntimeConfiguration{
//   		ServerProcesses: []serverProcess{
//   			&serverProcess{
//   				LaunchPath: jsii.String("/local/game/GameLiftExampleServer.x86_64"),
//   				Parameters: jsii.String("-logFile /local/game/logs/myserver1935.log -port 1935"),
//   				ConcurrentExecutions: jsii.Number(100),
//   			},
//   		},
//   	},
//   })
//
// See: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-multiprocess.html
//
// Experimental.
type RuntimeConfiguration struct {
	// A collection of server process configurations that identify what server processes to run on each instance in a fleet.
	// Experimental.
	ServerProcesses *[]*ServerProcess `field:"required" json:"serverProcesses" yaml:"serverProcesses"`
	// The maximum amount of time allowed to launch a new game session and have it report ready to host players.
	//
	// During this time, the game session is in status `ACTIVATING`.
	//
	// If the game session does not become active before the timeout, it is ended and the game session status is changed to `TERMINATED`.
	// Experimental.
	GameSessionActivationTimeout awscdk.Duration `field:"optional" json:"gameSessionActivationTimeout" yaml:"gameSessionActivationTimeout"`
	// The number of game sessions in status `ACTIVATING` to allow on an instance.
	//
	// This setting limits the instance resources that can be used for new game activations at any one time.
	// Experimental.
	MaxConcurrentGameSessionActivations *float64 `field:"optional" json:"maxConcurrentGameSessionActivations" yaml:"maxConcurrentGameSessionActivations"`
}


package awsgamelift


// A collection of server process configurations that describe the set of processes to run on each instance in a fleet.
//
// Server processes run either an executable in a custom game build or a Realtime Servers script. GameLift launches the configured processes, manages their life cycle, and replaces them as needed. Each instance checks regularly for an updated runtime configuration.
//
// A GameLift instance is limited to 50 processes running concurrently. To calculate the total number of processes in a runtime configuration, add the values of the `ConcurrentExecutions` parameter for each ServerProcess. Learn more about [Running Multiple Processes on a Fleet](https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-multiprocess.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runtimeConfigurationProperty := &RuntimeConfigurationProperty{
//   	GameSessionActivationTimeoutSeconds: jsii.Number(123),
//   	MaxConcurrentGameSessionActivations: jsii.Number(123),
//   	ServerProcesses: []interface{}{
//   		&ServerProcessProperty{
//   			ConcurrentExecutions: jsii.Number(123),
//   			LaunchPath: jsii.String("launchPath"),
//
//   			// the properties below are optional
//   			Parameters: jsii.String("parameters"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-runtimeconfiguration.html
//
type CfnFleet_RuntimeConfigurationProperty struct {
	// The maximum amount of time (in seconds) allowed to launch a new game session and have it report ready to host players.
	//
	// During this time, the game session is in status `ACTIVATING` . If the game session does not become active before the timeout, it is ended and the game session status is changed to `TERMINATED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-runtimeconfiguration.html#cfn-gamelift-fleet-runtimeconfiguration-gamesessionactivationtimeoutseconds
	//
	GameSessionActivationTimeoutSeconds *float64 `field:"optional" json:"gameSessionActivationTimeoutSeconds" yaml:"gameSessionActivationTimeoutSeconds"`
	// The number of game sessions in status `ACTIVATING` to allow on an instance or container.
	//
	// This setting limits the instance resources that can be used for new game activations at any one time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-runtimeconfiguration.html#cfn-gamelift-fleet-runtimeconfiguration-maxconcurrentgamesessionactivations
	//
	MaxConcurrentGameSessionActivations *float64 `field:"optional" json:"maxConcurrentGameSessionActivations" yaml:"maxConcurrentGameSessionActivations"`
	// A collection of server process configurations that identify what server processes to run on fleet computes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-runtimeconfiguration.html#cfn-gamelift-fleet-runtimeconfiguration-serverprocesses
	//
	ServerProcesses interface{} `field:"optional" json:"serverProcesses" yaml:"serverProcesses"`
}


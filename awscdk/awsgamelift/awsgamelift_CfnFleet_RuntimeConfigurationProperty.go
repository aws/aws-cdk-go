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
//   runtimeConfigurationProperty := &runtimeConfigurationProperty{
//   	gameSessionActivationTimeoutSeconds: jsii.Number(123),
//   	maxConcurrentGameSessionActivations: jsii.Number(123),
//   	serverProcesses: []interface{}{
//   		&serverProcessProperty{
//   			concurrentExecutions: jsii.Number(123),
//   			launchPath: jsii.String("launchPath"),
//
//   			// the properties below are optional
//   			parameters: jsii.String("parameters"),
//   		},
//   	},
//   }
//
type CfnFleet_RuntimeConfigurationProperty struct {
	// The maximum amount of time (in seconds) allowed to launch a new game session and have it report ready to host players.
	//
	// During this time, the game session is in status `ACTIVATING` . If the game session does not become active before the timeout, it is ended and the game session status is changed to `TERMINATED` .
	GameSessionActivationTimeoutSeconds *float64 `field:"optional" json:"gameSessionActivationTimeoutSeconds" yaml:"gameSessionActivationTimeoutSeconds"`
	// The number of game sessions in status `ACTIVATING` to allow on an instance.
	//
	// This setting limits the instance resources that can be used for new game activations at any one time.
	MaxConcurrentGameSessionActivations *float64 `field:"optional" json:"maxConcurrentGameSessionActivations" yaml:"maxConcurrentGameSessionActivations"`
	// A collection of server process configurations that identify what server processes to run on each instance in a fleet.
	ServerProcesses interface{} `field:"optional" json:"serverProcesses" yaml:"serverProcesses"`
}


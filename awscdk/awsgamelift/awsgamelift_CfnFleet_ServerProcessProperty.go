package awsgamelift


// A set of instructions for launching server processes on each instance in a fleet.
//
// Server processes run either an executable in a custom game build or a Realtime Servers script.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverProcessProperty := &serverProcessProperty{
//   	concurrentExecutions: jsii.Number(123),
//   	launchPath: jsii.String("launchPath"),
//
//   	// the properties below are optional
//   	parameters: jsii.String("parameters"),
//   }
//
type CfnFleet_ServerProcessProperty struct {
	// The number of server processes using this configuration that run concurrently on each instance.
	ConcurrentExecutions *float64 `field:"required" json:"concurrentExecutions" yaml:"concurrentExecutions"`
	// The location of a game build executable or the Realtime script file that contains the `Init()` function.
	//
	// Game builds and Realtime scripts are installed on instances at the root:
	//
	// - Windows (custom game builds only): `C:\game` . Example: " `C:\game\MyGame\server.exe` "
	// - Linux: `/local/game` . Examples: " `/local/game/MyGame/server.exe` " or " `/local/game/MyRealtimeScript.js` "
	LaunchPath *string `field:"required" json:"launchPath" yaml:"launchPath"`
	// An optional list of parameters to pass to the server executable or Realtime script on launch.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}


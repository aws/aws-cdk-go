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
//   serverProcessProperty := &ServerProcessProperty{
//   	ConcurrentExecutions: jsii.Number(123),
//   	LaunchPath: jsii.String("launchPath"),
//
//   	// the properties below are optional
//   	Parameters: jsii.String("parameters"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-serverprocess.html
//
type CfnFleet_ServerProcessProperty struct {
	// The number of server processes using this configuration that run concurrently on each instance or compute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-serverprocess.html#cfn-gamelift-fleet-serverprocess-concurrentexecutions
	//
	ConcurrentExecutions *float64 `field:"required" json:"concurrentExecutions" yaml:"concurrentExecutions"`
	// The location of a game build executable or Realtime script.
	//
	// Game builds and Realtime scripts are installed on instances at the root:
	//
	// - Windows (custom game builds only): `C:\game` . Example: " `C:\game\MyGame\server.exe` "
	// - Linux: `/local/game` . Examples: " `/local/game/MyGame/server.exe` " or " `/local/game/MyRealtimeScript.js` "
	//
	// > Amazon GameLift doesn't support the use of setup scripts that launch the game executable. For custom game builds, this parameter must indicate the executable that calls the server SDK operations `initSDK()` and `ProcessReady()` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-serverprocess.html#cfn-gamelift-fleet-serverprocess-launchpath
	//
	LaunchPath *string `field:"required" json:"launchPath" yaml:"launchPath"`
	// An optional list of parameters to pass to the server executable or Realtime script on launch.
	//
	// Length Constraints: Minimum length of 1. Maximum length of 1024.
	//
	// Pattern: [A-Za-z0-9_:.+\/\\\- =@{},?'\[\]"]+
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-serverprocess.html#cfn-gamelift-fleet-serverprocess-parameters
	//
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}


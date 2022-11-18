// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha


// Configuration of a fleet server process.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import gamelift_alpha "github.com/aws/aws-cdk-go/awscdkgameliftalpha"
//
//   serverProcess := &serverProcess{
//   	launchPath: jsii.String("launchPath"),
//
//   	// the properties below are optional
//   	concurrentExecutions: jsii.Number(123),
//   	parameters: jsii.String("parameters"),
//   }
//
// Experimental.
type ServerProcess struct {
	// The location of a game build executable or the Realtime script file that contains the Init() function.
	//
	// Game builds and Realtime scripts are installed on instances at the root:
	// - Windows (custom game builds only): `C:\game`. Example: `C:\game\MyGame\server.exe`
	// - Linux: `/local/game`. Examples: `/local/game/MyGame/server.exe` or `/local/game/MyRealtimeScript.js`
	// Experimental.
	LaunchPath *string `field:"required" json:"launchPath" yaml:"launchPath"`
	// The number of server processes using this configuration that run concurrently on each instance.
	//
	// Minimum is `1`.
	// Experimental.
	ConcurrentExecutions *float64 `field:"optional" json:"concurrentExecutions" yaml:"concurrentExecutions"`
	// An optional list of parameters to pass to the server executable or Realtime script on launch.
	// Experimental.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}


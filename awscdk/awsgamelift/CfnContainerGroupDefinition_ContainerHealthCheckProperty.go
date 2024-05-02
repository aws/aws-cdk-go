package awsgamelift


// Specifies how the process manager checks the health of containers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerHealthCheckProperty := &ContainerHealthCheckProperty{
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//
//   	// the properties below are optional
//   	Interval: jsii.Number(123),
//   	Retries: jsii.Number(123),
//   	StartPeriod: jsii.Number(123),
//   	Timeout: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerhealthcheck.html
//
type CfnContainerGroupDefinition_ContainerHealthCheckProperty struct {
	// A string array representing the command that the container runs to determine if it is healthy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerhealthcheck.html#cfn-gamelift-containergroupdefinition-containerhealthcheck-command
	//
	Command *[]*string `field:"required" json:"command" yaml:"command"`
	// How often (in seconds) the health is checked.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerhealthcheck.html#cfn-gamelift-containergroupdefinition-containerhealthcheck-interval
	//
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// How many times the process manager will retry the command after a timeout.
	//
	// (The first run of the command does not count as a retry.)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerhealthcheck.html#cfn-gamelift-containergroupdefinition-containerhealthcheck-retries
	//
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// The optional grace period (in seconds) to give a container time to boostrap before teh health check is declared failed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerhealthcheck.html#cfn-gamelift-containergroupdefinition-containerhealthcheck-startperiod
	//
	StartPeriod *float64 `field:"optional" json:"startPeriod" yaml:"startPeriod"`
	// How many seconds the process manager allows the command to run before canceling it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerhealthcheck.html#cfn-gamelift-containergroupdefinition-containerhealthcheck-timeout
	//
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}


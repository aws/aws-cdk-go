package awsgamelift


// Instructions on when and how to check the health of a support container in a container fleet.
//
// These properties override any Docker health checks that are set in the container image. For more information on container health checks, see [HealthCheck command](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_HealthCheck.html#ECS-Type-HealthCheck-command) in the *Amazon Elastic Container Service API* . Game server containers don't have a health check parameter; Amazon GameLift Servers automatically handles health checks for these containers.
//
// The following example instructs the container to initiate a health check command every 60 seconds and wait 10 seconds for it to succeed. If it fails, retry the command 3 times before flagging the container as unhealthy. It also tells the container to wait 100 seconds after launch before counting failed health checks.
//
// `{"Command": [ "CMD-SHELL", "ps cax | grep "processmanager" || exit 1" ], "Interval": 60, "Timeout": 10, "Retries": 3, "StartPeriod": 100 }`
//
// *Part of:* [SupportContainerDefinition](https://docs.aws.amazon.com/gamelift/latest/apireference/API_SupportContainerDefinition.html) , [SupportContainerDefinitionInput](https://docs.aws.amazon.com/gamelift/latest/apireference/API_SupportContainerDefinitionInput.html)
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
	// A string array that specifies the command that the container runs to determine if it's healthy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerhealthcheck.html#cfn-gamelift-containergroupdefinition-containerhealthcheck-command
	//
	Command *[]*string `field:"required" json:"command" yaml:"command"`
	// The time period (in seconds) between each health check.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerhealthcheck.html#cfn-gamelift-containergroupdefinition-containerhealthcheck-interval
	//
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The number of times to retry a failed health check before flagging the container unhealthy.
	//
	// The first run of the command does not count as a retry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerhealthcheck.html#cfn-gamelift-containergroupdefinition-containerhealthcheck-retries
	//
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// The optional grace period (in seconds) to give a container time to bootstrap before the first failed health check counts toward the number of retries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerhealthcheck.html#cfn-gamelift-containergroupdefinition-containerhealthcheck-startperiod
	//
	StartPeriod *float64 `field:"optional" json:"startPeriod" yaml:"startPeriod"`
	// The time period (in seconds) to wait for a health check to succeed before counting a failed health check.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerhealthcheck.html#cfn-gamelift-containergroupdefinition-containerhealthcheck-timeout
	//
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}


package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckProperty := &HealthCheckProperty{
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//   	Interval: jsii.Number(123),
//   	Retries: jsii.Number(123),
//   	StartPeriod: jsii.Number(123),
//   	Timeout: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-healthcheck.html
//
type CfnDaemonTaskDefinition_HealthCheckProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-healthcheck.html#cfn-ecs-daemontaskdefinition-healthcheck-command
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-healthcheck.html#cfn-ecs-daemontaskdefinition-healthcheck-interval
	//
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-healthcheck.html#cfn-ecs-daemontaskdefinition-healthcheck-retries
	//
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-healthcheck.html#cfn-ecs-daemontaskdefinition-healthcheck-startperiod
	//
	StartPeriod *float64 `field:"optional" json:"startPeriod" yaml:"startPeriod"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-healthcheck.html#cfn-ecs-daemontaskdefinition-healthcheck-timeout
	//
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}


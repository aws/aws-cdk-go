package awsecs


// The CloudWatch alarm configuration for a daemon.
//
// When enabled, CloudWatch alarms determine whether a daemon deployment has failed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   daemonAlarmConfigurationProperty := &DaemonAlarmConfigurationProperty{
//   	AlarmNames: []*string{
//   		jsii.String("alarmNames"),
//   	},
//   	Enable: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemon-daemonalarmconfiguration.html
//
type CfnDaemon_DaemonAlarmConfigurationProperty struct {
	// The CloudWatch alarm names to monitor during a daemon deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemon-daemonalarmconfiguration.html#cfn-ecs-daemon-daemonalarmconfiguration-alarmnames
	//
	AlarmNames *[]*string `field:"optional" json:"alarmNames" yaml:"alarmNames"`
	// Determines whether to use the CloudWatch alarm option in the daemon deployment process.
	//
	// The default value is ``false``.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemon-daemonalarmconfiguration.html#cfn-ecs-daemon-daemonalarmconfiguration-enable
	//
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
}


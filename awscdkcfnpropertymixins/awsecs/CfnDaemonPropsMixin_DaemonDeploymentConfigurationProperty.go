package awsecs


// Optional deployment parameters that control how a daemon rolls out updates across container instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   daemonDeploymentConfigurationProperty := &DaemonDeploymentConfigurationProperty{
//   	Alarms: &DaemonAlarmConfigurationProperty{
//   		AlarmNames: []*string{
//   			jsii.String("alarmNames"),
//   		},
//   		Enable: jsii.Boolean(false),
//   	},
//   	BakeTimeInMinutes: jsii.Number(123),
//   	DrainPercent: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemon-daemondeploymentconfiguration.html
//
type CfnDaemonPropsMixin_DaemonDeploymentConfigurationProperty struct {
	// The CloudWatch alarm configuration for a daemon.
	//
	// When enabled, CloudWatch alarms determine whether a daemon deployment has failed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemon-daemondeploymentconfiguration.html#cfn-ecs-daemon-daemondeploymentconfiguration-alarms
	//
	Alarms interface{} `field:"optional" json:"alarms" yaml:"alarms"`
	// The amount of time (in minutes) to wait after a successful deployment step before proceeding.
	//
	// This allows time to monitor for issues before continuing. The default value is 0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemon-daemondeploymentconfiguration.html#cfn-ecs-daemon-daemondeploymentconfiguration-baketimeinminutes
	//
	BakeTimeInMinutes *float64 `field:"optional" json:"bakeTimeInMinutes" yaml:"bakeTimeInMinutes"`
	// The percentage of container instances to drain simultaneously during a daemon deployment.
	//
	// Valid values are between 0.0 and 100.0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemon-daemondeploymentconfiguration.html#cfn-ecs-daemon-daemondeploymentconfiguration-drainpercent
	//
	DrainPercent *float64 `field:"optional" json:"drainPercent" yaml:"drainPercent"`
}


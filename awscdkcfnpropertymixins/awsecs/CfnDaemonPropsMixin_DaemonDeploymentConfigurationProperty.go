package awsecs


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemon-daemondeploymentconfiguration.html#cfn-ecs-daemon-daemondeploymentconfiguration-alarms
	//
	Alarms interface{} `field:"optional" json:"alarms" yaml:"alarms"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemon-daemondeploymentconfiguration.html#cfn-ecs-daemon-daemondeploymentconfiguration-baketimeinminutes
	//
	BakeTimeInMinutes *float64 `field:"optional" json:"bakeTimeInMinutes" yaml:"bakeTimeInMinutes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemon-daemondeploymentconfiguration.html#cfn-ecs-daemon-daemondeploymentconfiguration-drainpercent
	//
	DrainPercent *float64 `field:"optional" json:"drainPercent" yaml:"drainPercent"`
}


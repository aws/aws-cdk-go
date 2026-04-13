package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
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
type CfnDaemonPropsMixin_DaemonAlarmConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemon-daemonalarmconfiguration.html#cfn-ecs-daemon-daemonalarmconfiguration-alarmnames
	//
	AlarmNames *[]*string `field:"optional" json:"alarmNames" yaml:"alarmNames"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemon-daemonalarmconfiguration.html#cfn-ecs-daemon-daemonalarmconfiguration-enable
	//
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
}


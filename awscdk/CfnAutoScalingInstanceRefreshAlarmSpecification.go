package awscdk


// Alarm specification for the AutoScalingInstanceRefresh update policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAutoScalingInstanceRefreshAlarmSpecification := &CfnAutoScalingInstanceRefreshAlarmSpecification{
//   	Alarms: []*string{
//   		jsii.String("alarms"),
//   	},
//   }
//
type CfnAutoScalingInstanceRefreshAlarmSpecification struct {
	// The names of the CloudWatch alarms to monitor.
	Alarms *[]*string `field:"optional" json:"alarms" yaml:"alarms"`
}


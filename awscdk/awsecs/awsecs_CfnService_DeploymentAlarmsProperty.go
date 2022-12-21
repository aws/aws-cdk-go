package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentAlarmsProperty := &deploymentAlarmsProperty{
//   	alarmNames: []*string{
//   		jsii.String("alarmNames"),
//   	},
//   	enable: jsii.Boolean(false),
//   	rollback: jsii.Boolean(false),
//   }
//
type CfnService_DeploymentAlarmsProperty struct {
	// `CfnService.DeploymentAlarmsProperty.AlarmNames`.
	AlarmNames *[]*string `field:"required" json:"alarmNames" yaml:"alarmNames"`
	// `CfnService.DeploymentAlarmsProperty.Enable`.
	Enable interface{} `field:"required" json:"enable" yaml:"enable"`
	// `CfnService.DeploymentAlarmsProperty.Rollback`.
	Rollback interface{} `field:"required" json:"rollback" yaml:"rollback"`
}


package awscdkappconfigalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Defines monitors that will be associated with an AWS AppConfig environment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appconfig_alpha "github.com/aws/aws-cdk-go/awscdkappconfigalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var alarm alarm
//   var role role
//
//   monitor := &Monitor{
//   	Alarm: alarm,
//
//   	// the properties below are optional
//   	AlarmRole: role,
//   }
//
// Experimental.
type Monitor struct {
	// The Amazon CloudWatch alarm.
	// Experimental.
	Alarm awscloudwatch.IAlarm `field:"required" json:"alarm" yaml:"alarm"`
	// The IAM role for AWS AppConfig to view the alarm state.
	// Default: - A role is generated.
	//
	// Experimental.
	AlarmRole awsiam.IRole `field:"optional" json:"alarmRole" yaml:"alarmRole"`
}


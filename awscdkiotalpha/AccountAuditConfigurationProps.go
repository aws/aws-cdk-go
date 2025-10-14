package awscdkiotalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Properties for defining AWS IoT Audit Configuration.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   iot.NewAccountAuditConfiguration(this, jsii.String("AuditConfiguration"), &AccountAuditConfigurationProps{
//   	CheckConfiguration: &CheckConfiguration{
//   		DeviceCertificateAgeCheck: jsii.Boolean(true),
//   		// The default value is 365 days
//   		// Valid values range from 30 days (minimum) to 3650 days (10 years, maximum)
//   		DeviceCertificateAgeCheckDuration: awscdk.Duration_Days(jsii.Number(365)),
//   	},
//   })
//
// Experimental.
type AccountAuditConfigurationProps struct {
	// Specifies which audit checks are enabled and disabled for this account.
	// Default: - all checks are enabled.
	//
	// Experimental.
	CheckConfiguration *CheckConfiguration `field:"optional" json:"checkConfiguration" yaml:"checkConfiguration"`
	// The target SNS topic to which audit notifications are sent.
	// Default: - no notifications are sent.
	//
	// Experimental.
	TargetTopic awssns.ITopic `field:"optional" json:"targetTopic" yaml:"targetTopic"`
}


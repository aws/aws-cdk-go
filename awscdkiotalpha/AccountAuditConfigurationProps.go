package awscdkiotalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Properties for defining AWS IoT Audit Configuration.
//
// Example:
//   // Audit notification are sent to the SNS topic
//   var targetTopic iTopic
//
//
//   iot.NewAccountAuditConfiguration(this, jsii.String("AuditConfiguration"), &AccountAuditConfigurationProps{
//   	TargetTopic: TargetTopic,
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


package awsiot


// Information about the targets to which audit notifications are sent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   auditNotificationTargetProperty := &auditNotificationTargetProperty{
//   	enabled: jsii.Boolean(false),
//   	roleArn: jsii.String("roleArn"),
//   	targetArn: jsii.String("targetArn"),
//   }
//
type CfnAccountAuditConfiguration_AuditNotificationTargetProperty struct {
	// True if notifications to the target are enabled.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The ARN of the role that grants permission to send notifications to the target.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The ARN of the target (SNS topic) to which audit notifications are sent.
	TargetArn *string `field:"optional" json:"targetArn" yaml:"targetArn"`
}


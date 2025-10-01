package awsiot


// A reference to a ScheduledAudit resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduledAuditReference := &ScheduledAuditReference{
//   	ScheduledAuditArn: jsii.String("scheduledAuditArn"),
//   	ScheduledAuditName: jsii.String("scheduledAuditName"),
//   }
//
type ScheduledAuditReference struct {
	// The ARN of the ScheduledAudit resource.
	ScheduledAuditArn *string `field:"required" json:"scheduledAuditArn" yaml:"scheduledAuditArn"`
	// The ScheduledAuditName of the ScheduledAudit resource.
	ScheduledAuditName *string `field:"required" json:"scheduledAuditName" yaml:"scheduledAuditName"`
}


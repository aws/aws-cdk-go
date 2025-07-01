package awscdkiotalpha


// Construction properties for a Scheduled Audit.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import iot_alpha "github.com/aws/aws-cdk-go/awscdkiotalpha"
//
//   scheduledAuditAttributes := &ScheduledAuditAttributes{
//   	ScheduledAuditArn: jsii.String("scheduledAuditArn"),
//   	ScheduledAuditName: jsii.String("scheduledAuditName"),
//   }
//
// Experimental.
type ScheduledAuditAttributes struct {
	// The ARN of the scheduled audit.
	// Experimental.
	ScheduledAuditArn *string `field:"required" json:"scheduledAuditArn" yaml:"scheduledAuditArn"`
	// The scheduled audit name.
	// Experimental.
	ScheduledAuditName *string `field:"required" json:"scheduledAuditName" yaml:"scheduledAuditName"`
}


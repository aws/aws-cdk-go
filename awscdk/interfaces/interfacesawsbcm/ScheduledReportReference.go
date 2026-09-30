package interfacesawsbcm


// A reference to a ScheduledReport resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduledReportReference := &ScheduledReportReference{
//   	ScheduledReportArn: jsii.String("scheduledReportArn"),
//   }
//
type ScheduledReportReference struct {
	// The Arn of the ScheduledReport resource.
	ScheduledReportArn *string `field:"required" json:"scheduledReportArn" yaml:"scheduledReportArn"`
}


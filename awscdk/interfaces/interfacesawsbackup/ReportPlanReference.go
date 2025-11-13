package interfacesawsbackup


// A reference to a ReportPlan resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   reportPlanReference := &ReportPlanReference{
//   	ReportPlanArn: jsii.String("reportPlanArn"),
//   }
//
type ReportPlanReference struct {
	// The ReportPlanArn of the ReportPlan resource.
	ReportPlanArn *string `field:"required" json:"reportPlanArn" yaml:"reportPlanArn"`
}


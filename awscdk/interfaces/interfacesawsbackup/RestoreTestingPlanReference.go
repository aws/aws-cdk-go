package interfacesawsbackup


// A reference to a RestoreTestingPlan resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   restoreTestingPlanReference := &RestoreTestingPlanReference{
//   	RestoreTestingPlanArn: jsii.String("restoreTestingPlanArn"),
//   	RestoreTestingPlanName: jsii.String("restoreTestingPlanName"),
//   }
//
type RestoreTestingPlanReference struct {
	// The ARN of the RestoreTestingPlan resource.
	RestoreTestingPlanArn *string `field:"required" json:"restoreTestingPlanArn" yaml:"restoreTestingPlanArn"`
	// The RestoreTestingPlanName of the RestoreTestingPlan resource.
	RestoreTestingPlanName *string `field:"required" json:"restoreTestingPlanName" yaml:"restoreTestingPlanName"`
}


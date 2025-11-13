package interfacesawsbackup


// A reference to a RestoreTestingSelection resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   restoreTestingSelectionReference := &RestoreTestingSelectionReference{
//   	RestoreTestingPlanName: jsii.String("restoreTestingPlanName"),
//   	RestoreTestingSelectionName: jsii.String("restoreTestingSelectionName"),
//   }
//
type RestoreTestingSelectionReference struct {
	// The RestoreTestingPlanName of the RestoreTestingSelection resource.
	RestoreTestingPlanName *string `field:"required" json:"restoreTestingPlanName" yaml:"restoreTestingPlanName"`
	// The RestoreTestingSelectionName of the RestoreTestingSelection resource.
	RestoreTestingSelectionName *string `field:"required" json:"restoreTestingSelectionName" yaml:"restoreTestingSelectionName"`
}


package interfacesawsdrs


// A reference to a RecoveryPlan resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recoveryPlanReference := &RecoveryPlanReference{
//   	RecoveryPlanArn: jsii.String("recoveryPlanArn"),
//   }
//
type RecoveryPlanReference struct {
	// The Arn of the RecoveryPlan resource.
	RecoveryPlanArn *string `field:"required" json:"recoveryPlanArn" yaml:"recoveryPlanArn"`
}


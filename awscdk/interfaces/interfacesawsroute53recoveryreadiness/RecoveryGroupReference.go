package interfacesawsroute53recoveryreadiness


// A reference to a RecoveryGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recoveryGroupReference := &RecoveryGroupReference{
//   	RecoveryGroupArn: jsii.String("recoveryGroupArn"),
//   	RecoveryGroupName: jsii.String("recoveryGroupName"),
//   }
//
type RecoveryGroupReference struct {
	// The ARN of the RecoveryGroup resource.
	RecoveryGroupArn *string `field:"required" json:"recoveryGroupArn" yaml:"recoveryGroupArn"`
	// The RecoveryGroupName of the RecoveryGroup resource.
	RecoveryGroupName *string `field:"required" json:"recoveryGroupName" yaml:"recoveryGroupName"`
}


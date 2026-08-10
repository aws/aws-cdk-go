package interfacesawsredshiftserverless


// A reference to a RecoveryPoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recoveryPointReference := &RecoveryPointReference{
//   	RecoveryPointArn: jsii.String("recoveryPointArn"),
//   }
//
type RecoveryPointReference struct {
	// The Arn of the RecoveryPoint resource.
	RecoveryPointArn *string `field:"required" json:"recoveryPointArn" yaml:"recoveryPointArn"`
}


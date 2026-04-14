package interfacesawsbraket


// A reference to a SpendingLimit resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spendingLimitReference := &SpendingLimitReference{
//   	SpendingLimitArn: jsii.String("spendingLimitArn"),
//   }
//
type SpendingLimitReference struct {
	// The SpendingLimitArn of the SpendingLimit resource.
	SpendingLimitArn *string `field:"required" json:"spendingLimitArn" yaml:"spendingLimitArn"`
}


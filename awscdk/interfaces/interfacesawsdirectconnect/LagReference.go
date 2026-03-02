package interfacesawsdirectconnect


// A reference to a Lag resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lagReference := &LagReference{
//   	LagArn: jsii.String("lagArn"),
//   }
//
type LagReference struct {
	// The LagArn of the Lag resource.
	LagArn *string `field:"required" json:"lagArn" yaml:"lagArn"`
}


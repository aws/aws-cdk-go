package interfacesawsdeadline


// A reference to a Limit resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   limitReference := &LimitReference{
//   	FarmId: jsii.String("farmId"),
//   	LimitId: jsii.String("limitId"),
//   }
//
type LimitReference struct {
	// The FarmId of the Limit resource.
	FarmId *string `field:"required" json:"farmId" yaml:"farmId"`
	// The LimitId of the Limit resource.
	LimitId *string `field:"required" json:"limitId" yaml:"limitId"`
}


package interfacesawsdeadline


// A reference to a Farm resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   farmReference := &FarmReference{
//   	FarmArn: jsii.String("farmArn"),
//   }
//
type FarmReference struct {
	// The Arn of the Farm resource.
	FarmArn *string `field:"required" json:"farmArn" yaml:"farmArn"`
}


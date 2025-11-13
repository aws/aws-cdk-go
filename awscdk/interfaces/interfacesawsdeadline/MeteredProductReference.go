package interfacesawsdeadline


// A reference to a MeteredProduct resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   meteredProductReference := &MeteredProductReference{
//   	MeteredProductArn: jsii.String("meteredProductArn"),
//   }
//
type MeteredProductReference struct {
	// The Arn of the MeteredProduct resource.
	MeteredProductArn *string `field:"required" json:"meteredProductArn" yaml:"meteredProductArn"`
}


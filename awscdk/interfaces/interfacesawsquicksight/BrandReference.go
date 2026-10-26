package interfacesawsquicksight


// A reference to a Brand resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   brandReference := &BrandReference{
//   	BrandArn: jsii.String("brandArn"),
//   }
//
type BrandReference struct {
	// The Arn of the Brand resource.
	BrandArn *string `field:"required" json:"brandArn" yaml:"brandArn"`
}


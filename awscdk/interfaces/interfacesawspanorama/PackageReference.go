package interfacesawspanorama


// A reference to a Package resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   packageReference := &PackageReference{
//   	PackageArn: jsii.String("packageArn"),
//   	PackageId: jsii.String("packageId"),
//   }
//
type PackageReference struct {
	// The ARN of the Package resource.
	PackageArn *string `field:"required" json:"packageArn" yaml:"packageArn"`
	// The PackageId of the Package resource.
	PackageId *string `field:"required" json:"packageId" yaml:"packageId"`
}


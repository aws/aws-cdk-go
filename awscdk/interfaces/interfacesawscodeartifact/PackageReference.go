package interfacesawscodeartifact


// A reference to a Package resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   packageReference := &PackageReference{
//   	PackageArn: jsii.String("packageArn"),
//   }
//
type PackageReference struct {
	// The Arn of the Package resource.
	PackageArn *string `field:"required" json:"packageArn" yaml:"packageArn"`
}


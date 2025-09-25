package awscodeartifact


// A reference to a PackageGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   packageGroupReference := &PackageGroupReference{
//   	PackageGroupArn: jsii.String("packageGroupArn"),
//   }
//
type PackageGroupReference struct {
	// The Arn of the PackageGroup resource.
	PackageGroupArn *string `field:"required" json:"packageGroupArn" yaml:"packageGroupArn"`
}


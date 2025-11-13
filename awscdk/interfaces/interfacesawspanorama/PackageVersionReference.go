package interfacesawspanorama


// A reference to a PackageVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   packageVersionReference := &PackageVersionReference{
//   	PackageId: jsii.String("packageId"),
//   	PackageVersion: jsii.String("packageVersion"),
//   	PatchVersion: jsii.String("patchVersion"),
//   }
//
type PackageVersionReference struct {
	// The PackageId of the PackageVersion resource.
	PackageId *string `field:"required" json:"packageId" yaml:"packageId"`
	// The PackageVersion of the PackageVersion resource.
	PackageVersion *string `field:"required" json:"packageVersion" yaml:"packageVersion"`
	// The PatchVersion of the PackageVersion resource.
	PatchVersion *string `field:"required" json:"patchVersion" yaml:"patchVersion"`
}


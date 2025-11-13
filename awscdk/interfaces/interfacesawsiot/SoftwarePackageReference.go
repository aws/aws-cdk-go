package interfacesawsiot


// A reference to a SoftwarePackage resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   softwarePackageReference := &SoftwarePackageReference{
//   	PackageName: jsii.String("packageName"),
//   }
//
type SoftwarePackageReference struct {
	// The PackageName of the SoftwarePackage resource.
	PackageName *string `field:"required" json:"packageName" yaml:"packageName"`
}


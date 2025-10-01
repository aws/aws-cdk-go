package awsiot


// A reference to a SoftwarePackageVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   softwarePackageVersionReference := &SoftwarePackageVersionReference{
//   	PackageName: jsii.String("packageName"),
//   	VersionName: jsii.String("versionName"),
//   }
//
type SoftwarePackageVersionReference struct {
	// The PackageName of the SoftwarePackageVersion resource.
	PackageName *string `field:"required" json:"packageName" yaml:"packageName"`
	// The VersionName of the SoftwarePackageVersion resource.
	VersionName *string `field:"required" json:"versionName" yaml:"versionName"`
}


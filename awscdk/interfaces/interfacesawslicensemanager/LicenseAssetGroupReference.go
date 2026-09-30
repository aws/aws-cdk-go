package interfacesawslicensemanager


// A reference to a LicenseAssetGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   licenseAssetGroupReference := &LicenseAssetGroupReference{
//   	LicenseAssetGroupArn: jsii.String("licenseAssetGroupArn"),
//   }
//
type LicenseAssetGroupReference struct {
	// The LicenseAssetGroupArn of the LicenseAssetGroup resource.
	LicenseAssetGroupArn *string `field:"required" json:"licenseAssetGroupArn" yaml:"licenseAssetGroupArn"`
}


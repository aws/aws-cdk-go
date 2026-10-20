package interfacesawslicensemanager


// A reference to a LicenseConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   licenseConfigurationReference := &LicenseConfigurationReference{
//   	LicenseConfigurationArn: jsii.String("licenseConfigurationArn"),
//   }
//
type LicenseConfigurationReference struct {
	// The LicenseConfigurationArn of the LicenseConfiguration resource.
	LicenseConfigurationArn *string `field:"required" json:"licenseConfigurationArn" yaml:"licenseConfigurationArn"`
}


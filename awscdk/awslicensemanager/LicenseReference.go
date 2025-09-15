package awslicensemanager


// A reference to a License resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   licenseReference := &LicenseReference{
//   	LicenseArn: jsii.String("licenseArn"),
//   }
//
type LicenseReference struct {
	// The LicenseArn of the License resource.
	LicenseArn *string `field:"required" json:"licenseArn" yaml:"licenseArn"`
}


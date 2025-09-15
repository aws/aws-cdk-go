package awsdeadline


// A reference to a LicenseEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   licenseEndpointReference := &LicenseEndpointReference{
//   	LicenseEndpointArn: jsii.String("licenseEndpointArn"),
//   }
//
type LicenseEndpointReference struct {
	// The Arn of the LicenseEndpoint resource.
	LicenseEndpointArn *string `field:"required" json:"licenseEndpointArn" yaml:"licenseEndpointArn"`
}


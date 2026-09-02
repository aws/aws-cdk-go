package interfacesawsec2


// A reference to a IpamExternalResourceVerificationToken resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ipamExternalResourceVerificationTokenReference := &IpamExternalResourceVerificationTokenReference{
//   	IpamExternalResourceVerificationTokenArn: jsii.String("ipamExternalResourceVerificationTokenArn"),
//   }
//
type IpamExternalResourceVerificationTokenReference struct {
	// The IpamExternalResourceVerificationTokenArn of the IpamExternalResourceVerificationToken resource.
	IpamExternalResourceVerificationTokenArn *string `field:"required" json:"ipamExternalResourceVerificationTokenArn" yaml:"ipamExternalResourceVerificationTokenArn"`
}


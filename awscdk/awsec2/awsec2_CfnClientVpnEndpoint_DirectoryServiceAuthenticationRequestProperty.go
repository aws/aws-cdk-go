package awsec2


// Describes the Active Directory to be used for client authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   directoryServiceAuthenticationRequestProperty := &directoryServiceAuthenticationRequestProperty{
//   	directoryId: jsii.String("directoryId"),
//   }
//
type CfnClientVpnEndpoint_DirectoryServiceAuthenticationRequestProperty struct {
	// The ID of the Active Directory to be used for authentication.
	DirectoryId *string `field:"required" json:"directoryId" yaml:"directoryId"`
}


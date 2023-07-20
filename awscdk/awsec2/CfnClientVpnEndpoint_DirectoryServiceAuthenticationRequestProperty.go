package awsec2


// Describes the Active Directory to be used for client authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   directoryServiceAuthenticationRequestProperty := &DirectoryServiceAuthenticationRequestProperty{
//   	DirectoryId: jsii.String("directoryId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-directoryserviceauthenticationrequest.html
//
type CfnClientVpnEndpoint_DirectoryServiceAuthenticationRequestProperty struct {
	// The ID of the Active Directory to be used for authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-directoryserviceauthenticationrequest.html#cfn-ec2-clientvpnendpoint-directoryserviceauthenticationrequest-directoryid
	//
	DirectoryId *string `field:"required" json:"directoryId" yaml:"directoryId"`
}


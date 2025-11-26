package previewawsec2mixins


// Describes the authentication method to be used by a Client VPN endpoint.
//
// For more information, see [Authentication](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/authentication-authrization.html#client-authentication) in the *AWS Client VPN Administrator Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   clientAuthenticationRequestProperty := &ClientAuthenticationRequestProperty{
//   	ActiveDirectory: &DirectoryServiceAuthenticationRequestProperty{
//   		DirectoryId: jsii.String("directoryId"),
//   	},
//   	FederatedAuthentication: &FederatedAuthenticationRequestProperty{
//   		SamlProviderArn: jsii.String("samlProviderArn"),
//   		SelfServiceSamlProviderArn: jsii.String("selfServiceSamlProviderArn"),
//   	},
//   	MutualAuthentication: &CertificateAuthenticationRequestProperty{
//   		ClientRootCertificateChainArn: jsii.String("clientRootCertificateChainArn"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-clientauthenticationrequest.html
//
type CfnClientVpnEndpointPropsMixin_ClientAuthenticationRequestProperty struct {
	// Information about the Active Directory to be used, if applicable.
	//
	// You must provide this information if *Type* is `directory-service-authentication` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-clientauthenticationrequest.html#cfn-ec2-clientvpnendpoint-clientauthenticationrequest-activedirectory
	//
	ActiveDirectory interface{} `field:"optional" json:"activeDirectory" yaml:"activeDirectory"`
	// Information about the IAM SAML identity provider, if applicable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-clientauthenticationrequest.html#cfn-ec2-clientvpnendpoint-clientauthenticationrequest-federatedauthentication
	//
	FederatedAuthentication interface{} `field:"optional" json:"federatedAuthentication" yaml:"federatedAuthentication"`
	// Information about the authentication certificates to be used, if applicable.
	//
	// You must provide this information if *Type* is `certificate-authentication` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-clientauthenticationrequest.html#cfn-ec2-clientvpnendpoint-clientauthenticationrequest-mutualauthentication
	//
	MutualAuthentication interface{} `field:"optional" json:"mutualAuthentication" yaml:"mutualAuthentication"`
	// The type of client authentication to be used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-clientauthenticationrequest.html#cfn-ec2-clientvpnendpoint-clientauthenticationrequest-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}


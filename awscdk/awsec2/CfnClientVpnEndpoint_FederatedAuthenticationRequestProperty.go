package awsec2


// The IAM SAML identity provider used for federated authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   federatedAuthenticationRequestProperty := &FederatedAuthenticationRequestProperty{
//   	SamlProviderArn: jsii.String("samlProviderArn"),
//
//   	// the properties below are optional
//   	SelfServiceSamlProviderArn: jsii.String("selfServiceSamlProviderArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-federatedauthenticationrequest.html
//
type CfnClientVpnEndpoint_FederatedAuthenticationRequestProperty struct {
	// The Amazon Resource Name (ARN) of the IAM SAML identity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-federatedauthenticationrequest.html#cfn-ec2-clientvpnendpoint-federatedauthenticationrequest-samlproviderarn
	//
	SamlProviderArn *string `field:"required" json:"samlProviderArn" yaml:"samlProviderArn"`
	// The Amazon Resource Name (ARN) of the IAM SAML identity provider for the self-service portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-federatedauthenticationrequest.html#cfn-ec2-clientvpnendpoint-federatedauthenticationrequest-selfservicesamlproviderarn
	//
	SelfServiceSamlProviderArn *string `field:"optional" json:"selfServiceSamlProviderArn" yaml:"selfServiceSamlProviderArn"`
}


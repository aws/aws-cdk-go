package awsappintegrations


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   authConfigProperty := &AuthConfigProperty{
//   	AudienceOverride: jsii.String("audienceOverride"),
//   	AuthType: jsii.String("authType"),
//   	CredentialProviderIdentifier: jsii.String("credentialProviderIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-application-authconfig.html
//
type CfnApplicationPropsMixin_AuthConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-application-authconfig.html#cfn-appintegrations-application-authconfig-audienceoverride
	//
	AudienceOverride *string `field:"optional" json:"audienceOverride" yaml:"audienceOverride"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-application-authconfig.html#cfn-appintegrations-application-authconfig-authtype
	//
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-application-authconfig.html#cfn-appintegrations-application-authconfig-credentialprovideridentifier
	//
	CredentialProviderIdentifier *string `field:"optional" json:"credentialProviderIdentifier" yaml:"credentialProviderIdentifier"`
}


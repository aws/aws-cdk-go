package previewawspcaconnectorscepmixins


// Contains OpenID Connect (OIDC) parameters for use with Microsoft Intune.
//
// For more information about using Connector for SCEP for Microsoft Intune, see [Using Connector for SCEP for Microsoft Intune](https://docs.aws.amazon.com/privateca/latest/userguide/scep-connector.htmlconnector-for-scep-intune.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   openIdConfigurationProperty := &OpenIdConfigurationProperty{
//   	Audience: jsii.String("audience"),
//   	Issuer: jsii.String("issuer"),
//   	Subject: jsii.String("subject"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorscep-connector-openidconfiguration.html
//
type CfnConnectorPropsMixin_OpenIdConfigurationProperty struct {
	// The audience value to copy into your Microsoft Entra app registration's OIDC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorscep-connector-openidconfiguration.html#cfn-pcaconnectorscep-connector-openidconfiguration-audience
	//
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	// The issuer value to copy into your Microsoft Entra app registration's OIDC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorscep-connector-openidconfiguration.html#cfn-pcaconnectorscep-connector-openidconfiguration-issuer
	//
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// The subject value to copy into your Microsoft Entra app registration's OIDC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorscep-connector-openidconfiguration.html#cfn-pcaconnectorscep-connector-openidconfiguration-subject
	//
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
}


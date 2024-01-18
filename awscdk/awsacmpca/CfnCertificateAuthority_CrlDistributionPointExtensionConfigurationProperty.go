package awsacmpca


// Configures the default behavior of the CRL Distribution Point extension for certificates issued by your certificate authority.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   crlDistributionPointExtensionConfigurationProperty := &CrlDistributionPointExtensionConfigurationProperty{
//   	OmitExtension: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-crldistributionpointextensionconfiguration.html
//
type CfnCertificateAuthority_CrlDistributionPointExtensionConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-crldistributionpointextensionconfiguration.html#cfn-acmpca-certificateauthority-crldistributionpointextensionconfiguration-omitextension
	//
	OmitExtension interface{} `field:"required" json:"omitExtension" yaml:"omitExtension"`
}


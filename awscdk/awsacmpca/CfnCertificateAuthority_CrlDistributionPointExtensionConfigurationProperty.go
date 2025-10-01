package awsacmpca


// Contains configuration information for the default behavior of the CRL Distribution Point (CDP) extension in certificates issued by your CA.
//
// This extension contains a link to download the CRL, so you can check whether a certificate has been revoked. To choose whether you want this extension omitted or not in certificates issued by your CA, you can set the *OmitExtension* parameter.
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
	// Configures whether the CRL Distribution Point extension should be populated with the default URL to the CRL.
	//
	// If set to `true` , then the CDP extension will not be present in any certificates issued by that CA unless otherwise specified through CSR or API passthrough.
	//
	// > Only set this if you have another way to distribute the CRL Distribution Points for certificates issued by your CA, such as the Matter Distributed Compliance Ledger.
	// >
	// > This configuration cannot be enabled with a custom CNAME set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-crldistributionpointextensionconfiguration.html#cfn-acmpca-certificateauthority-crldistributionpointextensionconfiguration-omitextension
	//
	OmitExtension interface{} `field:"required" json:"omitExtension" yaml:"omitExtension"`
}


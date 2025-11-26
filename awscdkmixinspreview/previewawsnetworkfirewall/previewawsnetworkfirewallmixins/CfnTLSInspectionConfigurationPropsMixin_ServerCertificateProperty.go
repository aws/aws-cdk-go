package previewawsnetworkfirewallmixins


// Any Certificate Manager (ACM) Secure Sockets Layer/Transport Layer Security (SSL/TLS) server certificate that's associated with a [ServerCertificateConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-tlsinspectionconfiguration-servercertificateconfiguration.html) . Used in a [TLSInspectionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-networkfirewall-tlsinspectionconfiguration.html) for inspection of inbound traffic to your firewall. You must request or import a SSL/TLS certificate into ACM for each domain Network Firewall needs to decrypt and inspect. AWS Network Firewall uses the SSL/TLS certificates to decrypt specified inbound SSL/TLS traffic going to your firewall. For information about working with certificates in Certificate Manager , see [Request a public certificate](https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-request-public.html) or [Importing certificates](https://docs.aws.amazon.com/acm/latest/userguide/import-certificate.html) in the *Certificate Manager User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   serverCertificateProperty := &ServerCertificateProperty{
//   	ResourceArn: jsii.String("resourceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-tlsinspectionconfiguration-servercertificate.html
//
type CfnTLSInspectionConfigurationPropsMixin_ServerCertificateProperty struct {
	// The Amazon Resource Name (ARN) of the Certificate Manager SSL/TLS server certificate that's used for inbound SSL/TLS inspection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-tlsinspectionconfiguration-servercertificate.html#cfn-networkfirewall-tlsinspectionconfiguration-servercertificate-resourcearn
	//
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
}


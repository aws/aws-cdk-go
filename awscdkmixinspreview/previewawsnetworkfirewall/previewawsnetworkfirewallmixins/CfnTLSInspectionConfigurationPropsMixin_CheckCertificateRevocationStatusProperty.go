package previewawsnetworkfirewallmixins


// When enabled, Network Firewall checks if the server certificate presented by the server in the SSL/TLS connection has a revoked or unkown status.
//
// If the certificate has an unknown or revoked status, you must specify the actions that Network Firewall takes on outbound traffic. To check the certificate revocation status, you must also specify a `CertificateAuthorityArn` in [ServerCertificateConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-networkfirewall-servercertificateconfiguration.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   checkCertificateRevocationStatusProperty := &CheckCertificateRevocationStatusProperty{
//   	RevokedStatusAction: jsii.String("revokedStatusAction"),
//   	UnknownStatusAction: jsii.String("unknownStatusAction"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-tlsinspectionconfiguration-checkcertificaterevocationstatus.html
//
type CfnTLSInspectionConfigurationPropsMixin_CheckCertificateRevocationStatusProperty struct {
	// Configures how Network Firewall processes traffic when it determines that the certificate presented by the server in the SSL/TLS connection has a revoked status.
	//
	// - *PASS* - Allow the connection to continue, and pass subsequent packets to the stateful engine for inspection.
	// - *DROP* - Network Firewall closes the connection and drops subsequent packets for that connection.
	// - *REJECT* - Network Firewall sends a TCP reject packet back to your client. The service closes the connection and drops subsequent packets for that connection. `REJECT` is available only for TCP traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-tlsinspectionconfiguration-checkcertificaterevocationstatus.html#cfn-networkfirewall-tlsinspectionconfiguration-checkcertificaterevocationstatus-revokedstatusaction
	//
	RevokedStatusAction *string `field:"optional" json:"revokedStatusAction" yaml:"revokedStatusAction"`
	// Configures how Network Firewall processes traffic when it determines that the certificate presented by the server in the SSL/TLS connection has an unknown status, or a status that cannot be determined for any other reason, including when the service is unable to connect to the OCSP and CRL endpoints for the certificate.
	//
	// - *PASS* - Allow the connection to continue, and pass subsequent packets to the stateful engine for inspection.
	// - *DROP* - Network Firewall closes the connection and drops subsequent packets for that connection.
	// - *REJECT* - Network Firewall sends a TCP reject packet back to your client. The service closes the connection and drops subsequent packets for that connection. `REJECT` is available only for TCP traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-tlsinspectionconfiguration-checkcertificaterevocationstatus.html#cfn-networkfirewall-tlsinspectionconfiguration-checkcertificaterevocationstatus-unknownstatusaction
	//
	UnknownStatusAction *string `field:"optional" json:"unknownStatusAction" yaml:"unknownStatusAction"`
}


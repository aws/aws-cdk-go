package awsnetworkfirewall


// Configures the AWS Certificate Manager certificates and scope that Network Firewall uses to decrypt and re-encrypt traffic using a [TLSInspectionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-tlsinspectionconfiguration.html) . You can configure `ServerCertificates` for inbound SSL/TLS inspection, a `CertificateAuthorityArn` for outbound SSL/TLS inspection, or both. For information about working with certificates for TLS inspection, see [Using SSL/TLS server certficiates with TLS inspection configurations](https://docs.aws.amazon.com/network-firewall/latest/developerguide/tls-inspection-certificate-requirements.html) in the *AWS Network Firewall Developer Guide* .
//
// > If a server certificate that's associated with your [TLSInspectionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-tlsinspectionconfiguration.html) is revoked, deleted, or expired it can result in client-side TLS errors.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverCertificateConfigurationProperty := &ServerCertificateConfigurationProperty{
//   	CertificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   	CheckCertificateRevocationStatus: &CheckCertificateRevocationStatusProperty{
//   		RevokedStatusAction: jsii.String("revokedStatusAction"),
//   		UnknownStatusAction: jsii.String("unknownStatusAction"),
//   	},
//   	Scopes: []interface{}{
//   		&ServerCertificateScopeProperty{
//   			DestinationPorts: []interface{}{
//   				&PortRangeProperty{
//   					FromPort: jsii.Number(123),
//   					ToPort: jsii.Number(123),
//   				},
//   			},
//   			Destinations: []interface{}{
//   				&AddressProperty{
//   					AddressDefinition: jsii.String("addressDefinition"),
//   				},
//   			},
//   			Protocols: []interface{}{
//   				jsii.Number(123),
//   			},
//   			SourcePorts: []interface{}{
//   				&PortRangeProperty{
//   					FromPort: jsii.Number(123),
//   					ToPort: jsii.Number(123),
//   				},
//   			},
//   			Sources: []interface{}{
//   				&AddressProperty{
//   					AddressDefinition: jsii.String("addressDefinition"),
//   				},
//   			},
//   		},
//   	},
//   	ServerCertificates: []interface{}{
//   		&ServerCertificateProperty{
//   			ResourceArn: jsii.String("resourceArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-tlsinspectionconfiguration-servercertificateconfiguration.html
//
type CfnTLSInspectionConfiguration_ServerCertificateConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the imported certificate authority (CA) certificate within AWS Certificate Manager (ACM) to use for outbound SSL/TLS inspection.
	//
	// The following limitations apply:
	//
	// - You can use CA certificates that you imported into ACM, but you can't generate CA certificates with ACM.
	// - You can't use certificates issued by AWS Private Certificate Authority .
	//
	// For more information about configuring certificates for outbound inspection, see [Using SSL/TLS certificates with TLS inspection configurations](https://docs.aws.amazon.com/network-firewall/latest/developerguide/tls-inspection-certificate-requirements.html) in the *AWS Network Firewall Developer Guide* .
	//
	// For information about working with certificates in ACM, see [Importing certificates](https://docs.aws.amazon.com/acm/latest/userguide/import-certificate.html) in the *AWS Certificate Manager User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-tlsinspectionconfiguration-servercertificateconfiguration.html#cfn-networkfirewall-tlsinspectionconfiguration-servercertificateconfiguration-certificateauthorityarn
	//
	CertificateAuthorityArn *string `field:"optional" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// When enabled, Network Firewall checks if the server certificate presented by the server in the SSL/TLS connection has a revoked or unkown status.
	//
	// If the certificate has an unknown or revoked status, you must specify the actions that Network Firewall takes on outbound traffic. To check the certificate revocation status, you must also specify a `CertificateAuthorityArn` in [ServerCertificateConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-servercertificateconfiguration.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-tlsinspectionconfiguration-servercertificateconfiguration.html#cfn-networkfirewall-tlsinspectionconfiguration-servercertificateconfiguration-checkcertificaterevocationstatus
	//
	CheckCertificateRevocationStatus interface{} `field:"optional" json:"checkCertificateRevocationStatus" yaml:"checkCertificateRevocationStatus"`
	// A list of scopes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-tlsinspectionconfiguration-servercertificateconfiguration.html#cfn-networkfirewall-tlsinspectionconfiguration-servercertificateconfiguration-scopes
	//
	Scopes interface{} `field:"optional" json:"scopes" yaml:"scopes"`
	// The list of server certificates to use for inbound SSL/TLS inspection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-tlsinspectionconfiguration-servercertificateconfiguration.html#cfn-networkfirewall-tlsinspectionconfiguration-servercertificateconfiguration-servercertificates
	//
	ServerCertificates interface{} `field:"optional" json:"serverCertificates" yaml:"serverCertificates"`
}


package awsnetworkfirewall


// The object that defines a TLS inspection configuration.
//
// This, along with `TLSInspectionConfigurationResponse` , define the TLS inspection configuration. You can retrieve all objects for a TLS inspection configuration by calling `DescribeTLSInspectionConfiguration` .
//
// AWS Network Firewall uses a TLS inspection configuration to decrypt traffic. Network Firewall re-encrypts the traffic before sending it to its destination.
//
// To use a TLS inspection configuration, you add it to a new Network Firewall firewall policy, then you apply the firewall policy to a firewall. Network Firewall acts as a proxy service to decrypt and inspect the traffic traveling through your firewalls. You can reference a TLS inspection configuration from more than one firewall policy, and you can use a firewall policy in more than one firewall. For more information about using TLS inspection configurations, see [Inspecting SSL/TLS traffic with TLS
// inspection configurations](https://docs.aws.amazon.com/network-firewall/latest/developerguide/tls-inspection.html) in the *AWS Network Firewall Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tLSInspectionConfigurationProperty := &TLSInspectionConfigurationProperty{
//   	ServerCertificateConfigurations: []interface{}{
//   		&ServerCertificateConfigurationProperty{
//   			CertificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   			CheckCertificateRevocationStatus: &CheckCertificateRevocationStatusProperty{
//   				RevokedStatusAction: jsii.String("revokedStatusAction"),
//   				UnknownStatusAction: jsii.String("unknownStatusAction"),
//   			},
//   			Scopes: []interface{}{
//   				&ServerCertificateScopeProperty{
//   					DestinationPorts: []interface{}{
//   						&PortRangeProperty{
//   							FromPort: jsii.Number(123),
//   							ToPort: jsii.Number(123),
//   						},
//   					},
//   					Destinations: []interface{}{
//   						&AddressProperty{
//   							AddressDefinition: jsii.String("addressDefinition"),
//   						},
//   					},
//   					Protocols: []interface{}{
//   						jsii.Number(123),
//   					},
//   					SourcePorts: []interface{}{
//   						&PortRangeProperty{
//   							FromPort: jsii.Number(123),
//   							ToPort: jsii.Number(123),
//   						},
//   					},
//   					Sources: []interface{}{
//   						&AddressProperty{
//   							AddressDefinition: jsii.String("addressDefinition"),
//   						},
//   					},
//   				},
//   			},
//   			ServerCertificates: []interface{}{
//   				&ServerCertificateProperty{
//   					ResourceArn: jsii.String("resourceArn"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-tlsinspectionconfiguration-tlsinspectionconfiguration.html
//
type CfnTLSInspectionConfiguration_TLSInspectionConfigurationProperty struct {
	// Lists the server certificate configurations that are associated with the TLS configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-tlsinspectionconfiguration-tlsinspectionconfiguration.html#cfn-networkfirewall-tlsinspectionconfiguration-tlsinspectionconfiguration-servercertificateconfigurations
	//
	ServerCertificateConfigurations interface{} `field:"optional" json:"serverCertificateConfigurations" yaml:"serverCertificateConfigurations"`
}


package mixinsawsnetworkfirewall

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnTLSInspectionConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTLSInspectionConfigurationMixinProps := &CfnTLSInspectionConfigurationMixinProps{
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TlsInspectionConfiguration: &TLSInspectionConfigurationProperty{
//   		ServerCertificateConfigurations: []interface{}{
//   			&ServerCertificateConfigurationProperty{
//   				CertificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   				CheckCertificateRevocationStatus: &CheckCertificateRevocationStatusProperty{
//   					RevokedStatusAction: jsii.String("revokedStatusAction"),
//   					UnknownStatusAction: jsii.String("unknownStatusAction"),
//   				},
//   				Scopes: []interface{}{
//   					&ServerCertificateScopeProperty{
//   						DestinationPorts: []interface{}{
//   							&PortRangeProperty{
//   								FromPort: jsii.Number(123),
//   								ToPort: jsii.Number(123),
//   							},
//   						},
//   						Destinations: []interface{}{
//   							&AddressProperty{
//   								AddressDefinition: jsii.String("addressDefinition"),
//   							},
//   						},
//   						Protocols: []interface{}{
//   							jsii.Number(123),
//   						},
//   						SourcePorts: []interface{}{
//   							&PortRangeProperty{
//   								FromPort: jsii.Number(123),
//   								ToPort: jsii.Number(123),
//   							},
//   						},
//   						Sources: []interface{}{
//   							&AddressProperty{
//   								AddressDefinition: jsii.String("addressDefinition"),
//   							},
//   						},
//   					},
//   				},
//   				ServerCertificates: []interface{}{
//   					&ServerCertificateProperty{
//   						ResourceArn: jsii.String("resourceArn"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	TlsInspectionConfigurationName: jsii.String("tlsInspectionConfigurationName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-tlsinspectionconfiguration.html
//
type CfnTLSInspectionConfigurationMixinProps struct {
	// A description of the TLS inspection configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-tlsinspectionconfiguration.html#cfn-networkfirewall-tlsinspectionconfiguration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The key:value pairs to associate with the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-tlsinspectionconfiguration.html#cfn-networkfirewall-tlsinspectionconfiguration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The object that defines a TLS inspection configuration.
	//
	// AWS Network Firewall uses TLS inspection configurations to decrypt your firewall's inbound and outbound SSL/TLS traffic. After decryption, AWS Network Firewall inspects the traffic according to your firewall policy's stateful rules, and then re-encrypts it before sending it to its destination. You can enable inspection of your firewall's inbound traffic, outbound traffic, or both. To use TLS inspection with your firewall, you must first import or provision certificates using Certificate Manager , create a TLS inspection configuration, add that configuration to a new firewall policy, and then associate that policy with your firewall. For more information about using TLS inspection configurations, see [Inspecting SSL/TLS traffic with TLS inspection configurations](https://docs.aws.amazon.com/network-firewall/latest/developerguide/tls-inspection.html) in the *AWS Network Firewall Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-tlsinspectionconfiguration.html#cfn-networkfirewall-tlsinspectionconfiguration-tlsinspectionconfiguration
	//
	TlsInspectionConfiguration interface{} `field:"optional" json:"tlsInspectionConfiguration" yaml:"tlsInspectionConfiguration"`
	// The descriptive name of the TLS inspection configuration.
	//
	// You can't change the name of a TLS inspection configuration after you create it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-tlsinspectionconfiguration.html#cfn-networkfirewall-tlsinspectionconfiguration-tlsinspectionconfigurationname
	//
	TlsInspectionConfigurationName *string `field:"optional" json:"tlsInspectionConfigurationName" yaml:"tlsInspectionConfigurationName"`
}


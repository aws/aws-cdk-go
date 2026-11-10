package awsmediaconnect


// The configuration settings for a standard destination endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   standardDestinationEndpointConfigurationProperty := &StandardDestinationEndpointConfigurationProperty{
//   	DestinationNetworkingConfiguration: &StandardDestinationEndpointNetworkingConfigurationProperty{
//   		Public: &PublicDestinationEndpointNetworkingConfigurationProperty{
//   			DestinationAllowRules: []interface{}{
//   				&PublicDestinationEndpointNetworkingConfigurationAllowRuleProperty{
//   					Cidr: jsii.String("cidr"),
//   				},
//   			},
//   		},
//   		Vpc: &VpcDestinationEndpointNetworkingConfigurationProperty{
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetId: jsii.String("subnetId"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	ProtocolConfiguration: &StandardDestinationEndpointProtocolConfigurationProperty{
//   		Rist: &RistStandardDestinationEndpointConfigurationProperty{
//   			DestinationAddress: jsii.String("destinationAddress"),
//   			DestinationIngestPort: jsii.Number(123),
//   		},
//   		Rtp: &RtpStandardDestinationEndpointConfigurationProperty{
//   			DestinationAddress: jsii.String("destinationAddress"),
//   			DestinationIngestPort: jsii.Number(123),
//   			ForwardErrorCorrectionState: jsii.String("forwardErrorCorrectionState"),
//   		},
//   		SrtCaller: &SrtCallerStandardDestinationEndpointConfigurationProperty{
//   			MinimumLatencyMilliseconds: jsii.Number(123),
//   			SourceListenerPort: jsii.Number(123),
//
//   			// the properties below are optional
//   			EncryptionPassphrase: &EndpointEncryptionSecretConfigurationProperty{
//   				SecretArn: jsii.String("secretArn"),
//   			},
//   		},
//   		SrtListener: &SrtListenerStandardDestinationEndpointConfigurationProperty{
//   			DestinationAddress: jsii.String("destinationAddress"),
//   			DestinationIngestPort: jsii.Number(123),
//   			MinimumLatencyMilliseconds: jsii.Number(123),
//
//   			// the properties below are optional
//   			EncryptionPassphrase: &EndpointEncryptionSecretConfigurationProperty{
//   				SecretArn: jsii.String("secretArn"),
//   			},
//   			StreamId: jsii.String("streamId"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	DestinationNetworkingType: jsii.String("destinationNetworkingType"),
//   	Protocol: jsii.String("protocol"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-standarddestinationendpointconfiguration.html
//
type CfnDestinationEndpoint_StandardDestinationEndpointConfigurationProperty struct {
	// The networking configuration for a standard destination endpoint.
	//
	// This is a union type, so exactly one of its members is set: public for a publicly accessible endpoint, or vpc for an endpoint in a virtual private cloud (VPC).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-standarddestinationendpointconfiguration.html#cfn-mediaconnect-destinationendpoint-standarddestinationendpointconfiguration-destinationnetworkingconfiguration
	//
	DestinationNetworkingConfiguration interface{} `field:"required" json:"destinationNetworkingConfiguration" yaml:"destinationNetworkingConfiguration"`
	// The name of the destination endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-standarddestinationendpointconfiguration.html#cfn-mediaconnect-destinationendpoint-standarddestinationendpointconfiguration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The protocol configuration for a standard destination endpoint.
	//
	// This is a union type, so exactly one of its members is set, depending on the protocol that the destination endpoint uses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-standarddestinationendpointconfiguration.html#cfn-mediaconnect-destinationendpoint-standarddestinationendpointconfiguration-protocolconfiguration
	//
	ProtocolConfiguration interface{} `field:"required" json:"protocolConfiguration" yaml:"protocolConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-standarddestinationendpointconfiguration.html#cfn-mediaconnect-destinationendpoint-standarddestinationendpointconfiguration-destinationnetworkingtype
	//
	DestinationNetworkingType *string `field:"optional" json:"destinationNetworkingType" yaml:"destinationNetworkingType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-standarddestinationendpointconfiguration.html#cfn-mediaconnect-destinationendpoint-standarddestinationendpointconfiguration-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}


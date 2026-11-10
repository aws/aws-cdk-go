package awsmediaconnect


// The configuration settings for a standard source endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   standardSourceEndpointConfigurationProperty := &StandardSourceEndpointConfigurationProperty{
//   	Name: jsii.String("name"),
//   	Protocol: jsii.String("protocol"),
//   	ProtocolConfiguration: &StandardSourceEndpointProtocolConfigurationProperty{
//   		Rist: &RistStandardSourceEndpointConfigurationProperty{
//   			DestinationIngestPort: jsii.Number(123),
//   			RecoveryLatencyMilliseconds: jsii.Number(123),
//   		},
//   		Rtp: &RtpStandardSourceEndpointConfigurationProperty{
//   			DestinationIngestPort: jsii.Number(123),
//   			ForwardErrorCorrectionState: jsii.String("forwardErrorCorrectionState"),
//   		},
//   		SrtCaller: &SrtCallerStandardSourceEndpointConfigurationProperty{
//   			DestinationIngestPort: jsii.Number(123),
//   			EncryptionPassphrase: &EndpointEncryptionSecretConfigurationProperty{
//   				SecretArn: jsii.String("secretArn"),
//   			},
//   			MinimumLatencyMilliseconds: jsii.Number(123),
//   		},
//   		SrtListener: &SrtListenerStandardSourceEndpointConfigurationProperty{
//   			EncryptionPassphrase: &EndpointEncryptionSecretConfigurationProperty{
//   				SecretArn: jsii.String("secretArn"),
//   			},
//   			MinimumLatencyMilliseconds: jsii.Number(123),
//   			SourceAddress: jsii.String("sourceAddress"),
//   			SourceListenerPort: jsii.Number(123),
//   			StreamId: jsii.String("streamId"),
//   		},
//   	},
//   	SourceNetworkingConfiguration: &StandardSourceEndpointNetworkingConfigurationProperty{
//   		Public: &PublicSourceEndpointNetworkingConfigurationProperty{
//   			SourceAllowRules: []interface{}{
//   				&PublicSourceEndpointNetworkingConfigurationAllowRuleProperty{
//   					Cidr: jsii.String("cidr"),
//   				},
//   			},
//   		},
//   		Vpc: &VpcSourceEndpointNetworkingConfigurationProperty{
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetId: jsii.String("subnetId"),
//   		},
//   	},
//   	SourceNetworkingType: jsii.String("sourceNetworkingType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-standardsourceendpointconfiguration.html
//
type CfnSourceEndpointPropsMixin_StandardSourceEndpointConfigurationProperty struct {
	// The name of the source endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-standardsourceendpointconfiguration.html#cfn-mediaconnect-sourceendpoint-standardsourceendpointconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-standardsourceendpointconfiguration.html#cfn-mediaconnect-sourceendpoint-standardsourceendpointconfiguration-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The protocol configuration for a standard source endpoint.
	//
	// This is a union type, so exactly one of its members is set, depending on the protocol that the source endpoint uses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-standardsourceendpointconfiguration.html#cfn-mediaconnect-sourceendpoint-standardsourceendpointconfiguration-protocolconfiguration
	//
	ProtocolConfiguration interface{} `field:"optional" json:"protocolConfiguration" yaml:"protocolConfiguration"`
	// The networking configuration for a standard source endpoint.
	//
	// This is a union type, so exactly one of its members is set: public for a publicly accessible endpoint, or vpc for an endpoint in a virtual private cloud (VPC).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-standardsourceendpointconfiguration.html#cfn-mediaconnect-sourceendpoint-standardsourceendpointconfiguration-sourcenetworkingconfiguration
	//
	SourceNetworkingConfiguration interface{} `field:"optional" json:"sourceNetworkingConfiguration" yaml:"sourceNetworkingConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-standardsourceendpointconfiguration.html#cfn-mediaconnect-sourceendpoint-standardsourceendpointconfiguration-sourcenetworkingtype
	//
	SourceNetworkingType *string `field:"optional" json:"sourceNetworkingType" yaml:"sourceNetworkingType"`
}


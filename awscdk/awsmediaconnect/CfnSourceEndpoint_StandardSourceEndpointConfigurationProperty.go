package awsmediaconnect


// The configuration settings for a standard source endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   standardSourceEndpointConfigurationProperty := &StandardSourceEndpointConfigurationProperty{
//   	Name: jsii.String("name"),
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
//   			MinimumLatencyMilliseconds: jsii.Number(123),
//
//   			// the properties below are optional
//   			EncryptionPassphrase: &EndpointEncryptionSecretConfigurationProperty{
//   				SecretArn: jsii.String("secretArn"),
//   			},
//   		},
//   		SrtListener: &SrtListenerStandardSourceEndpointConfigurationProperty{
//   			MinimumLatencyMilliseconds: jsii.Number(123),
//   			SourceAddress: jsii.String("sourceAddress"),
//   			SourceListenerPort: jsii.Number(123),
//
//   			// the properties below are optional
//   			EncryptionPassphrase: &EndpointEncryptionSecretConfigurationProperty{
//   				SecretArn: jsii.String("secretArn"),
//   			},
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
//
//   	// the properties below are optional
//   	Protocol: jsii.String("protocol"),
//   	SourceNetworkingType: jsii.String("sourceNetworkingType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-standardsourceendpointconfiguration.html
//
type CfnSourceEndpoint_StandardSourceEndpointConfigurationProperty struct {
	// The name of the source endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-standardsourceendpointconfiguration.html#cfn-mediaconnect-sourceendpoint-standardsourceendpointconfiguration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The protocol configuration for a standard source endpoint.
	//
	// This is a union type, so exactly one of its members is set, depending on the protocol that the source endpoint uses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-standardsourceendpointconfiguration.html#cfn-mediaconnect-sourceendpoint-standardsourceendpointconfiguration-protocolconfiguration
	//
	ProtocolConfiguration interface{} `field:"required" json:"protocolConfiguration" yaml:"protocolConfiguration"`
	// The networking configuration for a standard source endpoint.
	//
	// This is a union type, so exactly one of its members is set: public for a publicly accessible endpoint, or vpc for an endpoint in a virtual private cloud (VPC).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-standardsourceendpointconfiguration.html#cfn-mediaconnect-sourceendpoint-standardsourceendpointconfiguration-sourcenetworkingconfiguration
	//
	SourceNetworkingConfiguration interface{} `field:"required" json:"sourceNetworkingConfiguration" yaml:"sourceNetworkingConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-standardsourceendpointconfiguration.html#cfn-mediaconnect-sourceendpoint-standardsourceendpointconfiguration-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-standardsourceendpointconfiguration.html#cfn-mediaconnect-sourceendpoint-standardsourceendpointconfiguration-sourcenetworkingtype
	//
	SourceNetworkingType *string `field:"optional" json:"sourceNetworkingType" yaml:"sourceNetworkingType"`
}


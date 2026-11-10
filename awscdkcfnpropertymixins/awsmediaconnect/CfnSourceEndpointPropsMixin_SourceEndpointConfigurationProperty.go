package awsmediaconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   sourceEndpointConfigurationProperty := &SourceEndpointConfigurationProperty{
//   	RouterOutput: &RouterOutputSourceEndpointConfigurationProperty{
//   		RouterOutputArn: jsii.String("routerOutputArn"),
//   	},
//   	Standard: &StandardSourceEndpointConfigurationProperty{
//   		Name: jsii.String("name"),
//   		Protocol: jsii.String("protocol"),
//   		ProtocolConfiguration: &StandardSourceEndpointProtocolConfigurationProperty{
//   			Rist: &RistStandardSourceEndpointConfigurationProperty{
//   				DestinationIngestPort: jsii.Number(123),
//   				RecoveryLatencyMilliseconds: jsii.Number(123),
//   			},
//   			Rtp: &RtpStandardSourceEndpointConfigurationProperty{
//   				DestinationIngestPort: jsii.Number(123),
//   				ForwardErrorCorrectionState: jsii.String("forwardErrorCorrectionState"),
//   			},
//   			SrtCaller: &SrtCallerStandardSourceEndpointConfigurationProperty{
//   				DestinationIngestPort: jsii.Number(123),
//   				EncryptionPassphrase: &EndpointEncryptionSecretConfigurationProperty{
//   					SecretArn: jsii.String("secretArn"),
//   				},
//   				MinimumLatencyMilliseconds: jsii.Number(123),
//   			},
//   			SrtListener: &SrtListenerStandardSourceEndpointConfigurationProperty{
//   				EncryptionPassphrase: &EndpointEncryptionSecretConfigurationProperty{
//   					SecretArn: jsii.String("secretArn"),
//   				},
//   				MinimumLatencyMilliseconds: jsii.Number(123),
//   				SourceAddress: jsii.String("sourceAddress"),
//   				SourceListenerPort: jsii.Number(123),
//   				StreamId: jsii.String("streamId"),
//   			},
//   		},
//   		SourceNetworkingConfiguration: &StandardSourceEndpointNetworkingConfigurationProperty{
//   			Public: &PublicSourceEndpointNetworkingConfigurationProperty{
//   				SourceAllowRules: []interface{}{
//   					&PublicSourceEndpointNetworkingConfigurationAllowRuleProperty{
//   						Cidr: jsii.String("cidr"),
//   					},
//   				},
//   			},
//   			Vpc: &VpcSourceEndpointNetworkingConfigurationProperty{
//   				SecurityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				SubnetId: jsii.String("subnetId"),
//   			},
//   		},
//   		SourceNetworkingType: jsii.String("sourceNetworkingType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-sourceendpointconfiguration.html
//
type CfnSourceEndpointPropsMixin_SourceEndpointConfigurationProperty struct {
	// The configuration settings for a source endpoint that represents a router output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-sourceendpointconfiguration.html#cfn-mediaconnect-sourceendpoint-sourceendpointconfiguration-routeroutput
	//
	RouterOutput interface{} `field:"optional" json:"routerOutput" yaml:"routerOutput"`
	// The configuration settings for a standard source endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-sourceendpointconfiguration.html#cfn-mediaconnect-sourceendpoint-sourceendpointconfiguration-standard
	//
	Standard interface{} `field:"optional" json:"standard" yaml:"standard"`
}


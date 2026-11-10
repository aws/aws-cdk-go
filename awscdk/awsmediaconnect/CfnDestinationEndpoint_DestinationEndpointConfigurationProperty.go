package awsmediaconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationEndpointConfigurationProperty := &DestinationEndpointConfigurationProperty{
//   	RouterInput: &RouterInputDestinationEndpointConfigurationProperty{
//   		RouterInputArn: jsii.String("routerInputArn"),
//   	},
//   	Standard: &StandardDestinationEndpointConfigurationProperty{
//   		DestinationNetworkingConfiguration: &StandardDestinationEndpointNetworkingConfigurationProperty{
//   			Public: &PublicDestinationEndpointNetworkingConfigurationProperty{
//   				DestinationAllowRules: []interface{}{
//   					&PublicDestinationEndpointNetworkingConfigurationAllowRuleProperty{
//   						Cidr: jsii.String("cidr"),
//   					},
//   				},
//   			},
//   			Vpc: &VpcDestinationEndpointNetworkingConfigurationProperty{
//   				SecurityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				SubnetId: jsii.String("subnetId"),
//   			},
//   		},
//   		Name: jsii.String("name"),
//   		ProtocolConfiguration: &StandardDestinationEndpointProtocolConfigurationProperty{
//   			Rist: &RistStandardDestinationEndpointConfigurationProperty{
//   				DestinationAddress: jsii.String("destinationAddress"),
//   				DestinationIngestPort: jsii.Number(123),
//   			},
//   			Rtp: &RtpStandardDestinationEndpointConfigurationProperty{
//   				DestinationAddress: jsii.String("destinationAddress"),
//   				DestinationIngestPort: jsii.Number(123),
//   				ForwardErrorCorrectionState: jsii.String("forwardErrorCorrectionState"),
//   			},
//   			SrtCaller: &SrtCallerStandardDestinationEndpointConfigurationProperty{
//   				MinimumLatencyMilliseconds: jsii.Number(123),
//   				SourceListenerPort: jsii.Number(123),
//
//   				// the properties below are optional
//   				EncryptionPassphrase: &EndpointEncryptionSecretConfigurationProperty{
//   					SecretArn: jsii.String("secretArn"),
//   				},
//   			},
//   			SrtListener: &SrtListenerStandardDestinationEndpointConfigurationProperty{
//   				DestinationAddress: jsii.String("destinationAddress"),
//   				DestinationIngestPort: jsii.Number(123),
//   				MinimumLatencyMilliseconds: jsii.Number(123),
//
//   				// the properties below are optional
//   				EncryptionPassphrase: &EndpointEncryptionSecretConfigurationProperty{
//   					SecretArn: jsii.String("secretArn"),
//   				},
//   				StreamId: jsii.String("streamId"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		DestinationNetworkingType: jsii.String("destinationNetworkingType"),
//   		Protocol: jsii.String("protocol"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-destinationendpointconfiguration.html
//
type CfnDestinationEndpoint_DestinationEndpointConfigurationProperty struct {
	// The configuration settings for a destination endpoint that represents a router input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-destinationendpointconfiguration.html#cfn-mediaconnect-destinationendpoint-destinationendpointconfiguration-routerinput
	//
	RouterInput interface{} `field:"optional" json:"routerInput" yaml:"routerInput"`
	// The configuration settings for a standard destination endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-destinationendpointconfiguration.html#cfn-mediaconnect-destinationendpoint-destinationendpointconfiguration-standard
	//
	Standard interface{} `field:"optional" json:"standard" yaml:"standard"`
}


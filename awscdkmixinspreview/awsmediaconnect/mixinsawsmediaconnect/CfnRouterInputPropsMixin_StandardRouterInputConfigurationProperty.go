package mixinsawsmediaconnect


// The configuration settings for a standard router input, including the protocol, protocol-specific configuration, network interface, and availability zone.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   standardRouterInputConfigurationProperty := &StandardRouterInputConfigurationProperty{
//   	NetworkInterfaceArn: jsii.String("networkInterfaceArn"),
//   	Protocol: jsii.String("protocol"),
//   	ProtocolConfiguration: &RouterInputProtocolConfigurationProperty{
//   		Rist: &RistRouterInputConfigurationProperty{
//   			Port: jsii.Number(123),
//   			RecoveryLatencyMilliseconds: jsii.Number(123),
//   		},
//   		Rtp: &RtpRouterInputConfigurationProperty{
//   			ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   			Port: jsii.Number(123),
//   		},
//   		SrtCaller: &SrtCallerRouterInputConfigurationProperty{
//   			DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   				EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   					RoleArn: jsii.String("roleArn"),
//   					SecretArn: jsii.String("secretArn"),
//   				},
//   			},
//   			MinimumLatencyMilliseconds: jsii.Number(123),
//   			SourceAddress: jsii.String("sourceAddress"),
//   			SourcePort: jsii.Number(123),
//   			StreamId: jsii.String("streamId"),
//   		},
//   		SrtListener: &SrtListenerRouterInputConfigurationProperty{
//   			DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   				EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   					RoleArn: jsii.String("roleArn"),
//   					SecretArn: jsii.String("secretArn"),
//   				},
//   			},
//   			MinimumLatencyMilliseconds: jsii.Number(123),
//   			Port: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-standardrouterinputconfiguration.html
//
type CfnRouterInputPropsMixin_StandardRouterInputConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the network interface associated with the standard router input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-standardrouterinputconfiguration.html#cfn-mediaconnect-routerinput-standardrouterinputconfiguration-networkinterfacearn
	//
	NetworkInterfaceArn *string `field:"optional" json:"networkInterfaceArn" yaml:"networkInterfaceArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-standardrouterinputconfiguration.html#cfn-mediaconnect-routerinput-standardrouterinputconfiguration-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The protocol configuration settings for a router input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-standardrouterinputconfiguration.html#cfn-mediaconnect-routerinput-standardrouterinputconfiguration-protocolconfiguration
	//
	ProtocolConfiguration interface{} `field:"optional" json:"protocolConfiguration" yaml:"protocolConfiguration"`
}


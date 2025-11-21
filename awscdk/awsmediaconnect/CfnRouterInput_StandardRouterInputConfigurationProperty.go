package awsmediaconnect


// The configuration settings for a standard router input, including the protocol, protocol-specific configuration, network interface, and availability zone.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   standardRouterInputConfigurationProperty := &StandardRouterInputConfigurationProperty{
//   	NetworkInterfaceArn: jsii.String("networkInterfaceArn"),
//   	ProtocolConfiguration: &RouterInputProtocolConfigurationProperty{
//   		Rist: &RistRouterInputConfigurationProperty{
//   			Port: jsii.Number(123),
//   			RecoveryLatencyMilliseconds: jsii.Number(123),
//   		},
//   		Rtp: &RtpRouterInputConfigurationProperty{
//   			Port: jsii.Number(123),
//
//   			// the properties below are optional
//   			ForwardErrorCorrection: jsii.String("forwardErrorCorrection"),
//   		},
//   		SrtCaller: &SrtCallerRouterInputConfigurationProperty{
//   			MinimumLatencyMilliseconds: jsii.Number(123),
//   			SourceAddress: jsii.String("sourceAddress"),
//   			SourcePort: jsii.Number(123),
//
//   			// the properties below are optional
//   			DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   				EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   					RoleArn: jsii.String("roleArn"),
//   					SecretArn: jsii.String("secretArn"),
//   				},
//   			},
//   			StreamId: jsii.String("streamId"),
//   		},
//   		SrtListener: &SrtListenerRouterInputConfigurationProperty{
//   			MinimumLatencyMilliseconds: jsii.Number(123),
//   			Port: jsii.Number(123),
//
//   			// the properties below are optional
//   			DecryptionConfiguration: &SrtDecryptionConfigurationProperty{
//   				EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   					RoleArn: jsii.String("roleArn"),
//   					SecretArn: jsii.String("secretArn"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	Protocol: jsii.String("protocol"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-standardrouterinputconfiguration.html
//
type CfnRouterInput_StandardRouterInputConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the network interface associated with the standard router input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-standardrouterinputconfiguration.html#cfn-mediaconnect-routerinput-standardrouterinputconfiguration-networkinterfacearn
	//
	NetworkInterfaceArn *string `field:"required" json:"networkInterfaceArn" yaml:"networkInterfaceArn"`
	// The protocol configuration settings for a router input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-standardrouterinputconfiguration.html#cfn-mediaconnect-routerinput-standardrouterinputconfiguration-protocolconfiguration
	//
	ProtocolConfiguration interface{} `field:"required" json:"protocolConfiguration" yaml:"protocolConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-standardrouterinputconfiguration.html#cfn-mediaconnect-routerinput-standardrouterinputconfiguration-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

